package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/shawnlower/go-ltp/api"
	go_ltp "github.com/shawnlower/go-ltp/api/proto"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/kv/bolt"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/voc"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

var (
	store *cayley.Handle
)

// Return a version string to the user
func (s *Server) GetVersion(ctx context.Context, in *go_ltp.Empty) (*go_ltp.VersionResponse, error) {
	log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))

	err := status.Error(codes.OK, "OK")
	return &go_ltp.VersionResponse{VersionString: "LTP Server version 0.0.0"}, err
}

// Provide some diagnostic information (items stored, etc)
func (s *Server) GetServerInfo(ctx context.Context, in *go_ltp.Empty) (*go_ltp.ServerInfoResponse, error) {
	log.Debug(fmt.Sprintf("GetServerInfo called. ctx: %#v\n", ctx))

	info := map[string]string{
		"Server name":  "Quad Damage",
		"Quads Stored": fmt.Sprintf("%d", store.QuadStore.Size()),
	}

	err := status.Error(codes.OK, "OK")
	return &go_ltp.ServerInfoResponse{InfoItems: info}, err
}

// Pretty-print the metadata from a request
func PprintMeta(md metadata.MD) {
	var line string
	for k, items := range md {
		line += fmt.Sprintf(" %s=[", k)
		for i, v := range items {
			if i > 0 {
				line += ", "
			}
			line += fmt.Sprintf("%s", v)
		}
		line += fmt.Sprintf("]")
	}
	log.Debug("context: ", line)
}

// Return the Type objects for a given item
func (s *Server) GetType(ctx context.Context, req *go_ltp.GetTypeRequest) (*go_ltp.GetTypeResponse, error) {

	iri := req.IRI

	testIRI := "http://ltp.shawnlower.net/_test"
	if iri == testIRI {
		t := &go_ltp.Type{
			IRI: testIRI,
			Label: "This is a built-in test item type",
			Parents: nil,
			Children: nil,
		}

		resp := &go_ltp.GetTypeResponse{Type: t}

		return resp, status.Error(codes.OK, "OK")
	} else if iri == "http://ltp.shawnlower.net/_missing" {
		return nil, status.Errorf(codes.NotFound, "Type not found: %s", iri)
	}

	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Server) GetItem(ctx context.Context, req *go_ltp.GetItemRequest) (*go_ltp.GetItemResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	PprintMeta(md)
	log.Debug("GetItemRequest for IRI: ", req.IRI)

	item := &go_ltp.Item{
		IRI: req.IRI,
	}

	predicates := cayley.StartPath(store, quad.String(req.IRI)).OutPredicates()

	err := predicates.Iterate(nil).EachValue(nil, func(pred quad.Value) {
		p := cayley.StartPath(store, quad.String(req.IRI)).Out(pred)
		p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			// nativeValue := quad.NativeOf(value)
			// log.Debugf("%s: %s = %s", req.IRI, pred, nativeValue)
			item.Statements = append(item.Statements, &go_ltp.Statement{
				Subject:   item.IRI,
				Predicate: pred.String(),
				Object:    value.String(),
			})
		})
	})

	resp := &go_ltp.GetItemResponse{
		Item: item,
	}

	return resp, err
}

func (s *Server) CreateItem(ctx context.Context, req *go_ltp.CreateItemRequest) (*go_ltp.CreateItemResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	PprintMeta(md)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, api.ErrInvalidItem
	}

	item := &go_ltp.Item{
		IRI:        "http://shawnlower.net/i/" + uuid.String(),
		Types:  req.Types,
		Statements: req.Statements,
	}

	log.Debugf("api.Item: [%s]", item)

	for _, itemType := range item.Types {
		// q := quad.Make(item.IRI, "rdf:type", itemType, nil)
		q := quad.Make(
			quad.String(item.IRI),
			quad.IRI("rdf:type").Full(),
			quad.IRI(itemType),
			quad.String(""))
		log.Debug("CreateItem: adding type: ", q)
		store.AddQuad(q)
	}

	for _, statement := range item.GetStatements() {
		q := quad.Make(
			item.IRI,
			quad.IRI(statement.Predicate).Full(),
			quad.String(statement.Object),
			nil)
		log.Debug("CreateItem: creating quad: ", q)
		store.AddQuad(q)
	}

	resp := &go_ltp.CreateItemResponse{
		Item: item,
	}
	return resp, nil
}

type Server struct{}

func InitStore(storeDBPath string) error {
	var err error

	if storeDBPath == ":memory:" {
		store, err = cayley.NewMemoryGraph()
		if err != nil {
			log.Fatal("Failed to create graph: ", err)
		}
	} else if storeDBPath == "" {
		// Use a temporary dir if none specified
		if storeDBPath == "" {
			storeDBPath, err = ioutil.TempDir("", "ltp.bolt")
			if err != nil {
				log.Fatal("Unable to create temp DB: ", err)
			}
			log.Warning(fmt.Sprintf("Using tempdir for DB: `%s'", storeDBPath))
		}
		err = graph.InitQuadStore("bolt", storeDBPath, nil)
		if err != nil {
			log.Fatal("Failed to create graph: ", err)
		}
		store, err = cayley.NewGraph("bolt", storeDBPath, nil)
		if err != nil {
			log.Fatal("Failed to open DB: ", err)
		}

	}

	voc.RegisterPrefix("schema:", "https://schema.org/")
	voc.RegisterPrefix("rdf:", "http://www.w3.org/1999/02/22-rdf-syntax-ns#")

	if err != nil {
		return fmt.Errorf("Failed to initialize store: %v", err)
	}

	return err
}

func InitServer() error {

	storeName := viper.GetString("store.driver")
	storeDBPath := viper.GetString("store.dbpath")

	if storeName == "" {
		storeName = "bolt"
		storeDBPath = ""
		log.Warningf("No store defined. Using bolt in a temp dir")
	}

	// todo: store.Init(storeDBPath)
	err := InitStore(storeDBPath)

	if err != nil {
		return fmt.Errorf("Failed to initialize store.")
	}

	return nil
}

func ShutdownServer() error {

	log.Debug("Shutting down server")

	// Now we create the path, to get to our data
	p := cayley.StartPath(store, quad.String("phrase of the day")).Out(quad.String("is of course"))

	// Now we iterate over results. Arguments:
	// 1. Optional context used for cancellation.
	// 2. Flag to optimize query before execution.
	// 3. Quad store, but we can omit it because we have already built path with it.
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value) // this converts RDF values to normal Go types
		fmt.Println(nativeValue)
	})

	return err
}

// Creates a new server with mandatory mutual-TLS authentication
func NewInsecureGrpcServer(host string, port string) (*grpc.Server, chan error) {

	// Get credentials

	server := grpc.NewServer()
	go_ltp.RegisterAPIServer(server, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	done := make(chan error, 1)

	// Initialization
	err := InitServer()
	if err != nil {
		done <- err
		return nil, done
	}

	// Listen and serve
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
		err = server.Serve(lis)
		ShutdownServer()
		done <- err
	}()

	return server, done
}

// Creates a new server with mandatory mutual-TLS authentication
func NewMutualTLSGrpcServer(host string, port string, certFile string, keyFile string, caCertFile string) (*grpc.Server, chan error) {

	// Get credentials
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Unable to load keypair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("Unable to load CA cert: %s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("Unable to add CA cert: %s", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(creds))
	go_ltp.RegisterAPIServer(server, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	done := make(chan error, 1)

	// Initialization
	err = InitServer()
	if err != nil {
		done <- err
		return nil, done
	}

	// Listen and serve
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
		err = server.Serve(lis)
		ShutdownServer()
		done <- err
	}()

	return server, done
}

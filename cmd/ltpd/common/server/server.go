package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	// "sync"

	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/api/proto"

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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

var (
	store *graph.Handle
)

func (s *Server) GetVersion(ctx context.Context, in *proto.Empty) (*proto.VersionResponse, error) {
	log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))
	return &proto.VersionResponse{VersionString: "LTP Server version 0.0.0"}, nil
}

func (s *Server) GetServerInfo(ctx context.Context, in *proto.Empty) (*proto.ServerInfoResponse, error) {
	log.Debug(fmt.Sprintf("GetServerInfo called. ctx: %#v\n", ctx))

	info := map[string]string{
		"Server name":  "Quad Damage",
		"Quads Stored": fmt.Sprintf("%d", store.QuadStore.Size()),
	}

	return &proto.ServerInfoResponse{InfoItems: info}, nil
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

func (s *Server) GetItemType(ctx context.Context, req *proto.GetItemTypeRequest) (*proto.GetItemTypeResponse, *Error) {

	iri := req.IRI

	testIRI := "http://ltp.shawnlower.net/_test"
	if iri == testIRI {
		t := &proto.ItemType{
			IRI: testIRI,
			Label: "This is a fake type",
			Parents: nil,
			Children: nil,
		}

		resp := &proto.GetItemTypeResponse{ItemType: t}

		return resp, &Error{codes.OK, "ltpd: OK"}
	} else if iri == "http://ltp.shawnlower.net/_missing" {
		return nil, &Error{codes.NotFound, "Item not found"}
	}

	return nil, &Error{codes.Unimplemented, "Unimplemented."}
}

func (s *Server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	PprintMeta(md)
	log.Debug("GetItemRequest for IRI: ", req.IRI)

	item := &proto.Item{
		IRI: req.IRI,
	}

	predicates := cayley.StartPath(store, quad.String(req.IRI)).OutPredicates()

	err := predicates.Iterate(nil).EachValue(nil, func(pred quad.Value) {
		p := cayley.StartPath(store, quad.String(req.IRI)).Out(pred)
		p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			// nativeValue := quad.NativeOf(value)
			// log.Debugf("%s: %s = %s", req.IRI, pred, nativeValue)
			item.Statements = append(item.Statements, &proto.Statement{
				Subject:   item.IRI,
				Predicate: pred.String(),
				Object:    value.String(),
			})
		})
	})

	resp := &proto.GetItemResponse{
		Item: item,
	}

	return resp, err
}

func (s *Server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	PprintMeta(md)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, api.ErrInvalidItem
	}

	item := &proto.Item{
		IRI:        "http://shawnlower.net/i/" + uuid.String(),
		ItemTypes:  req.ItemTypes,
		Statements: req.Statements,
	}

	log.Debugf("api.Item: [%s]", item)

	for _, itemType := range item.ItemTypes {
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

	resp := &proto.CreateItemResponse{
		Item: item,
	}
	return resp, nil
}

type Server struct{}

func InitServer() error {

	var err error

	storeName := viper.GetString("store.driver")
	storeDBPath := viper.GetString("store.dbpath")

	if storeName == "" {
		storeName = "bolt"
		storeDBPath = ""
		log.Warningf("No store defined. Using bolt in a temp dir")
	}

	if storeName == "memory" {
		store, err = cayley.NewMemoryGraph()
		if err != nil {
			log.Fatal("Failed to create graph: ", err)
		}
	} else if storeName == "bolt" {
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
		return fmt.Errorf("Failed to initialize store.")
	}

	return err

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
	proto.RegisterAPIServer(server, &Server{})

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
	proto.RegisterAPIServer(server, &Server{})

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

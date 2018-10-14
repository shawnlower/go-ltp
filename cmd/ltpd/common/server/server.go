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
    "github.com/cayleygraph/cayley/quad"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

)

var (
    store *graph.Handle
)

func (s *Server) GetVersion(ctx context.Context, in *proto.Empty) (*proto.VersionResponse, error) {
	log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))
	return &proto.VersionResponse{VersionString: "LTP Server v0.0.0"}, nil
}

func (s *Server) GetServerInfo(ctx context.Context, in *proto.Empty) (*proto.ServerInfoResponse, error) {
	log.Debug(fmt.Sprintf("GetServerInfo called. ctx: %#v\n", ctx))

    info := map[string]string{
        "Server name": "Quad Damage",
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
    log.Debug("CreateItem context: ", line)
}

func (s *Server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	// log.Debug("CreateItem called: ", request)
    md, _ := metadata.FromIncomingContext(ctx)
    PprintMeta(md)

	uuid, err := uuid.NewUUID()
	if err != nil {
        return nil, api.ErrInvalidItem
	}

	item := &proto.Item{
		IRI:       "http://shawnlower.net/i/" + uuid.String(),
		ItemTypes: req.ItemTypes,
        Statements: req.Statements,
	}

	log.Debug("api.Item: ", item)

    for _, statement := range item.GetStatements() {
        store.AddQuad(quad.Make(statement.Subject, statement.Predicate, statement.Object, item.IRI))
    }

	resp := &proto.CreateItemResponse{
		Item: item,
	}
	return resp, nil
}

type Server struct {}

func InitServer() error {

    var err error
    store, err = cayley.NewMemoryGraph()

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
    err := p.Iterate(nil).EachValue(nil, func(value quad.Value){
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
        done <-err
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
        ClientAuth: tls.RequireAndVerifyClientCert,
        Certificates: []tls.Certificate{certificate},
        ClientCAs: certPool,
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
        done <-err
    }()

    return server, done
}

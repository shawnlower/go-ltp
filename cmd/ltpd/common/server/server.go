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

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

func (s *Server) GetVersion(ctx context.Context, in *proto.Empty) (*proto.VersionResponse, error) {
	log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))
	return &proto.VersionResponse{VersionString: "LTP Server v0.0.0"}, nil
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

func (s *Server) CreateItem(ctx context.Context, request *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	// log.Debug("CreateItem called: ", request)
    md, _ := metadata.FromIncomingContext(ctx)
    PprintMeta(md)

	uuid, err := uuid.NewUUID()
	if err != nil {
        return nil, api.ErrInvalidItem
	}

	item := &proto.Item{
		IRI:       "http://shawnlower.net/i/" + uuid.String(),
		ItemTypes: request.ItemTypes,
	}

	log.Debug("api.Item: ", item)

	resp := &proto.CreateItemResponse{
		Item: item,
	}
	return resp, nil
}

type Server struct {}

func (s *Server) init() error {

    return fmt.Errorf("Failed to initialize store.")

}

// Creates a new server with mandatory mutual-TLS authentication
func NewInsecureGrpcServer(host string, port string) (*grpc.Server, chan error) {

    // Get credentials

	server := grpc.NewServer()
	proto.RegisterAPIServer(server, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

    // Setup network listener
    done := make(chan error, 1)
    go func() {
        lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
        err = server.Serve(lis)
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

    // Setup network listener
    done := make(chan error, 1)
    go func() {
        lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
        err = server.Serve(lis)
        done <-err
    }()

    return server, done
}

package server

import (
    "crypto/tls"
    "crypto/x509"
	"fmt"
    "io/ioutil"
	"net"

	"github.com/shawnlower/go-ltp/api"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func (s *Server) GetVersion(ctx context.Context, in *api.Empty) (*api.VersionResponse, error) {
	log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))
	return &api.VersionResponse{VersionString: "LTP Server v0.0.0"}, nil
}

func (s *Server) CreateItem(ctx context.Context, request *api.CreateItemRequest) (*api.CreateItemResponse, error) {
	log.Debug("CreateItem called: ", request)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, api.ErrInvalidItem
	}

	item := &api.Item{
		Uri:       "http://shawnlower.net/i/" + uuid.String(),
		ItemTypes: request.ItemTypes,
	}

	log.Debug("api.Item: ", item)

	resp := &api.CreateItemResponse{
		Item: item,
	}
	return resp, nil
}

type Server struct {}

// Creates a new server with mandatory mutual-TLS authentication
func NewInsecureGrpcServer(host string, port string) (*grpc.Server, error) {

    // Get credentials

	server := grpc.NewServer()
	api.RegisterAPIServer(server, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

    // Setup network listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
        log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

    return server, nil
}

// Creates a new server with mandatory mutual-TLS authentication
func NewMutualTLSGrpcServer(host string, port string, certFile string, keyFile string, caCertFile string) (*grpc.Server, error) {

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
	api.RegisterAPIServer(server, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

    // Setup network listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

    return server, nil
}

package server

import (
	"fmt"
	"net"
	"net/url"

	"github.com/shawnlower/go-ltp/api"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetVersion(ctx context.Context, in *api.Empty) (*api.VersionResponse, error) {
	log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))
	return &api.VersionResponse{VersionString: "LTP Server v0.0.0"}, nil
}

func (s *server) CreateItem(ctx context.Context, request *api.CreateItemRequest) (*api.CreateItemResponse, error) {
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

func Serve(listenAddr string) {

	var (
		scheme string
		host   string
		port   string
	)

	u, err := url.Parse(listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	host = u.Hostname()

	if u.Port() == "" {
		port = "17900"
	} else {
		port = u.Port()
	}

	switch u.Scheme {
	case "grpc":
		scheme = "grpc"
	case "http":
	case "https":
		scheme = "http"

	default:
		panic(fmt.Sprintf("Invalid scheme: %s. Valid schemes are 'grpc', 'http'", u.Scheme))
	}

	log.Debug(fmt.Sprintf("Starting %s server on %s port %s", scheme, host, port))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterAPIServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

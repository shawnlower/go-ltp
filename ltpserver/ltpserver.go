package ltpserver

import (
    "fmt"
    "net"
    "net/url"

    pb "github.com/shawnlower/go-ltp/pb"

    log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {}

func (s *server) GetVersion(ctx context.Context, in *pb.Empty) (*pb.VersionResponse, error) {
    log.Debug(fmt.Sprintf("GetVersion called. ctx: %#v\n", ctx))
	return &pb.VersionResponse{VersionString: "LTP Server v0.0.0"}, nil
}

func Serve(listenAddr string) {

    var (
        scheme string
        host string
        port string
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
	pb.RegisterAPIServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


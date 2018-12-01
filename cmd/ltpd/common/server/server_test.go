package server

import (
	"regexp"
	"testing"

	pb "github.com/shawnlower/go-ltp/api/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Test that we get a valid version response
func TestGetVersion(t *testing.T) {

	req := &pb.Empty{}

	s := Server{}
	resp, err := s.GetVersion(context.Background(), req)

	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Unknown, err.Error())
	}

	if st.Code() != codes.OK {
		t.Fatalf("GetType returned error: %v", st)
	}

	expectedRegexp := "LTP Server version.*"
	m, _ := regexp.MatchString(expectedRegexp, resp.VersionString)
	if !m {
		t.Fatalf("Expected regexp: `%s', got `%s'", expectedRegexp, resp.VersionString)
	}
}

// Use a reserved IRI (http://ltp.shawnlower.net/_test) to ensure we get
// a valid type back (DB not req'd)
func TestGetType(t *testing.T) {

	iri := "http://ltp.shawnlower.net/_test"

	req := &pb.GetTypeRequest{
		IRI: iri,
	}

	s := Server{}

	resp, err := s.GetType(context.Background(), req)

	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Unknown, err.Error())
	}

	if st.Code() != codes.OK {
		t.Fatalf("GetType returned error: %v", st)
	}

	if resp.Type.IRI != iri {
		t.Fatalf("Unexpected return item type: %s (expected %s)", resp.Type.IRI, iri)
	}
}

// Use a reserved IRI (http://ltp.shawnlower.net/_missing) to ensure we get
// a valid error on missing (DB not req'd)
func TestGetTypeMissing(t *testing.T) {

	iri := "http://ltp.shawnlower.net/_missing"

	req := &pb.GetTypeRequest{
		IRI: iri,
	}

	s := Server{}

	_, err := s.GetType(context.Background(), req)

	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Unknown, err.Error())
	}

	if st.Code() != codes.NotFound {
		t.Fatalf("Expected GetType to return not found, got: %v", st)
	}
}

func TestGetServerInfo(t *testing.T) {

	s := &Server{}

	InitServer()

	_, err := s.GetServerInfo(context.Background(), &pb.Empty{})

	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Unknown, err.Error())
	}

	if st.Code() != codes.OK {
		t.Fatalf("GetType returned error: %s", err)
	}

}

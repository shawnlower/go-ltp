// +build integration

package server

import (
	"testing"

	pb "github.com/shawnlower/go-ltp/api/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func TestGetTypeKnownGood(t *testing.T) {

	iri := "http://schema.org/Book"

	req := &pb.GetTypeRequest{
		IRI: iri,
	}

	s := Server{}
	InitServer()

	_, err := s.GetType(context.Background(), req)

	if err.Code != codes.OK {
		t.Fatalf("Expected GetType to return OK, got: %s", err)
	}
}


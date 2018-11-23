package server

import (
	"regexp"
	"testing"

	pb "github.com/shawnlower/go-ltp/api/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func TestGetVersion(t *testing.T) {

	req := &pb.Empty{}

	s := Server{}
	resp, err := s.GetVersion(context.Background(), req)
	if err != nil {
		t.Fatalf("GetVersion returned error: %s", err)
	}
	expectedRegexp := "LTP Server version.*"
	m, _ := regexp.MatchString(expectedRegexp, resp.VersionString)
	if !m {
		t.Fatalf("Expected regexp: `%s', got `%s'", expectedRegexp, resp.VersionString)
	}
}

func TestGetItemType(t *testing.T) {

	iri := "http://ltp.shawnlower.net/_test"

	req := &pb.GetItemTypeRequest{
		IRI: iri,
	}

	s := Server{}

	resp, err := s.GetItemType(context.Background(), req)
	if err.Code != codes.OK {
		t.Fatalf("GetItemType returned error: %s", err)
	}

	if resp.ItemType.IRI != iri {
		t.Fatalf("Unexpected return item type: %s (expected %s)", resp.ItemType.IRI, iri)
	}
}

func TestGetItemTypeMissing(t *testing.T) {

	iri := "http://ltp.shawnlower.net/_missing"

	req := &pb.GetItemTypeRequest{
		IRI: iri,
	}

	s := Server{}

	_, err := s.GetItemType(context.Background(), req)

	if err.Code != codes.NotFound {
		t.Fatalf("Expected GetItemType to return not found, got: %s", err)
	}
}

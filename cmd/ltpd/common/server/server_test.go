package server

import (
	"regexp"
	"testing"

	pb "github.com/shawnlower/go-ltp/api/proto"

	"golang.org/x/net/context"
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

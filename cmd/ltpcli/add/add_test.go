package add

import (
	"os"
	"testing"

	"github.com/shawnlower/go-ltp/cmd/ltpcli"
)

func TestCmdAdd(t *testing.T) {
	cmd := NewAddCommand()
	os.Args = []string{"add", "/etc/passwd", "--debug"}
	if err := cmd.Execute(); err != nil {
		t.Fatal("Error executing: ", err)
	}
}

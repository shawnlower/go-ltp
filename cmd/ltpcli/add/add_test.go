package add

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/shawnlower/go-ltp/api"
)
// todo: move to integration tests
// func TestCmdAdd(t *testing.T) {
// 	cmd := NewAddCommand()
// 	os.Args = []string{"add", "/etc/passwd", "--debug"}
// 	if err := cmd.Execute(); err != nil {
// 		t.Fatal("Error executing: ", err)
// 	}
// }

func TestCmdAdd(t *testing.T) {
	typeIRI := api.IRI("http://schema.org/Book")
	_ = typeIRI

	t.Log("Viper config...")
	viper.Set("remote.auth", "insecure")
	viper.Set("remote.url", "grpc://localhost:17900")

	doAdd("test book", typeIRI, []string{"/etc/passwd"})

}



// Implements a MIME-type parser to detect the content type of
// a file or stream.
//
// Using net/http, the implementation of the 'sniffing' logic is
// implemented in https://golang.org/src/net/http/sniff.go
package mimetype

import (
	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"
	"github.com/shawnlower/go-ltp/parsers"

	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var typeMap map[string]api.IRI = map[string]api.IRI{
	"^text/plain": api.IRI("schema:TextDigitalDocument"),
	"^image/.*":   api.IRI("schema:Photograph"),
}

type MimetypeParser struct {
	Statements []api.Statement
}

func (p *MimetypeParser) GetStatements() []api.Statement {
	return p.Statements
}

func (p *MimetypeParser) GetName() string {
	return "MimetypeParser"
}

func (p *MimetypeParser) Parse(r io.Reader) (io.Reader, error) {

	if r == nil {
		return nil, errors.New("Unable to use nil input reader")
	}

	// DetectContentType() uses at most 512 bytes
	buf := make([]byte, 512)

	r.Read(buf)
	mimetype := http.DetectContentType(buf)

	// Read to EOF so we don't block later
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
	}

	p.Statements = append(p.Statements, api.Statement{
		Subject:   api.IRI(""),
		Predicate: api.IRI("schema:encodingFormat"),
		Object:    api.String(mimetype),
	})

	for pat, t := range typeMap {
		if matched, _ := regexp.MatchString(pat, mimetype); matched == true {
			log.Debug(fmt.Sprintf("Detected mime-type `%s' matched pattern for type `%s'.",
				mimetype, t))
			p.Statements = append(p.Statements, api.Statement{
				Subject:   api.IRI(""),
				Predicate: api.IRI("rdf:type"),
				Object:    t,
			})
		}
	}

	return r, nil
}

func NewMimetypeParser() models.Parser {
	return &MimetypeParser{}
}

func init() {
	parsers.RegisterParser("MIMETYPE", NewMimetypeParser)
}

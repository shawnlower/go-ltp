package counter

import (
	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"
	"github.com/shawnlower/go-ltp/parsers"

	"fmt"
	"io"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type CounterParser struct {
	Name       string
	Statements []api.Statement
}

func (p *CounterParser) GetStatements() []api.Statement {
	return p.Statements
}

func (p *CounterParser) GetName() string {
	return "CounterParser"
}

func (p *CounterParser) Parse(reader io.Reader) (io.Reader, error) {

	buf := make([]byte, 1024)
	ctr := 0
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Read error.")
			break
		}
		ctr += n
		if err == io.EOF {
			// Write any remaining data, close the writer and break
			break
		}
	}

	p.Statements = []api.Statement{
		api.Statement{
			Subject:   api.IRI(""),
			Predicate: api.IRI("ltpcli.encoding.bytes"),
			Object:    api.String(strconv.Itoa(ctr)),
		},
	}

	return nil, nil
}

func NewCounterParser() models.Parser {
	return &CounterParser{}
}

func init() {
	parsers.RegisterParser("COUNTER", NewCounterParser)
	log.Debug("Registering counter parser")
}

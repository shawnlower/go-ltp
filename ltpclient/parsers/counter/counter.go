package counter

import (
    "github.com/shawnlower/go-ltp/ltpclient/models"
    "github.com/shawnlower/go-ltp/ltpclient/parsers"

	"fmt"
	"io"

    log "github.com/sirupsen/logrus"
)


type CounterParser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *CounterParser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *CounterParser) GetName() string {
    return "CounterParser"
}

func (p *CounterParser) Parse(reader io.Reader) (io.Reader, error) {

    buf := make([]byte, 1024)
    ctr := 0
    for {
        n, err := reader.Read(buf)
        if (err != nil && err != io.EOF) {
            fmt.Println("Read error.")
            break
        }
        ctr += n
        if (err == io.EOF) {
            // Write any remaining data, close the writer and break
            break
        }
    }

    p.Metadata = []models.MetadataItem{
        {"bytes": fmt.Sprintf("%d", ctr)},
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
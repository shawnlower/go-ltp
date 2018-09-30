package sha256

import (
    "github.com/shawnlower/go-ltp/ltpclient/models"
    "github.com/shawnlower/go-ltp/ltpclient/parsers"

    "crypto/sha256"
	"fmt"
	"io"

    log "github.com/sirupsen/logrus"
)

type Sha256Parser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *Sha256Parser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *Sha256Parser) GetName() string {
    return "Sha256Parser"
}

func (p *Sha256Parser) Parse(r io.Reader) (io.Reader, error) {
    h := sha256.New()
    _, err := io.Copy(h, r)

    p.Metadata = []models.MetadataItem{
        {"hash": fmt.Sprintf("%x", h.Sum(nil))},
    }

    return nil, err
}


func NewSha256Parser() models.Parser {
    return &Sha256Parser{}
}

func init() {
    parsers.RegisterParser("SHA256", NewSha256Parser)
    log.Debug("Registering SHA256 parser")
}

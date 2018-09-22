package sha512

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"
    "github.com/shawnlower/go-ltp/go-ltp/parsers"

    "crypto/sha512"
	"fmt"
	"io"

    log "github.com/sirupsen/logrus"
)

func GetMetadataJson() string {
    return "{}"
}

type Sha512Parser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *Sha512Parser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *Sha512Parser) GetName() string {
    return "Sha512Parser"
}

func (p *Sha512Parser) Parse(r io.Reader) (io.Reader, error) {
    h := sha512.New()
    _, err := io.Copy(h, r)

    var meta models.MetadataItem
    meta.Key = "sha512sum"
    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    p.Metadata = append(p.Metadata, meta)
    return nil, err
}

func NewSha512Parser() parsers.Parser {
    return &Sha512Parser{}
}

func init() {
    parsers.RegisterParser("SHA512", NewSha512Parser)
    log.Debug("Registering SHA512 parser")
}

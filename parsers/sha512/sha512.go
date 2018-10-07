package sha512

import (
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"
	"github.com/shawnlower/go-ltp/parsers"

	"crypto/sha512"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

func GetMetadataJson() string {
	return "{}"
}

type Sha512Parser struct {
	Name     string
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

	p.Metadata = []models.MetadataItem{
		{"hash": fmt.Sprintf("%x", h.Sum(nil))},
	}
	return nil, err
}

func NewSha512Parser() models.Parser {
	return &Sha512Parser{}
}

func init() {
	parsers.RegisterParser("SHA512", NewSha512Parser)
	log.Debug("Registering SHA512 parser")
}

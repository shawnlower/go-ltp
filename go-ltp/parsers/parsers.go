package parsers

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"

    "crypto/sha256"
    "crypto/sha512"
	"fmt"
	"io"

    log "github.com/sirupsen/logrus"
)

type Parser interface {
    Parse(r models.Reader) (error)
}

type Sha512Parser struct{
    Name string
    Metadata []models.MetadataItem
}

type Sha256Parser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *Sha256Parser) Parse(r io.Reader) (error) {
    h := sha256.New()
    _, err := io.Copy(h, r)

    var meta models.MetadataItem
    meta.Key = "sha256sum"
    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    p.Metadata = append(p.Metadata, meta)
    log.Debug(fmt.Sprintf("sha256 metadata: %s", p))
    return err
}

func (p *Sha512Parser) Parse(r io.Reader) (error) {
    h := sha512.New()
    _, err := io.Copy(h, r)

    var meta models.MetadataItem
    meta.Key = "sha512sum"
    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    p.Metadata = append(p.Metadata, meta)
    return err
}


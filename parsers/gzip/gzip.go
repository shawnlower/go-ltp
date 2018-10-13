package gzip

import (
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"
	"github.com/shawnlower/go-ltp/parsers"

	"bytes"
	"compress/gzip"
	"io"
	"time"
	// log "github.com/sirupsen/logrus"
)

type GzipParser struct {
	Metadata models.Metadata
}

func (p *GzipParser) GetMetadata() models.Metadata {
	return p.Metadata
}

func (p *GzipParser) GetName() string {
	return "GzipParser"
}

func (p *GzipParser) Parse(r io.Reader) (io.Reader, error) {

	if r == nil {
		panic("GzipParser cannot compress nil input reader %s")
	}
	buf := new(bytes.Buffer)
	gzipWriter := gzip.NewWriter(buf)
	defer gzipWriter.Close()

	comment := "This is a marvelous comment."
	modtime := time.Unix(1e8, 0)
	name := "no name"

	gzipWriter.Comment = comment
	gzipWriter.ModTime = modtime
	gzipWriter.Name = name

	io.Copy(gzipWriter, r)

	// var meta models.MetadataItem
	// meta := models.MetadataItem{ "name": name }
	p.Metadata = models.Metadata{
		"name": name,
        "comment":  comment,
        "modified": modtime.String(),
	}
	return buf, nil
}
func NewGzipParser() models.Parser {
	return &GzipParser{}
}

func init() {
	parsers.RegisterParser("GZIP", NewGzipParser)
}

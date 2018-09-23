package gzip

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"
    "github.com/shawnlower/go-ltp/go-ltp/parsers"

    "bytes"
    "compress/gzip"
	"fmt"
	"io"
    "time"

    // log "github.com/sirupsen/logrus"
)


type GzipParser struct{
    Metadata []models.MetadataItem
}

func (p *GzipParser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *GzipParser) GetName() string {
    return "GzipParser"
}

func (p *GzipParser) Parse(r io.Reader) (io.Reader, error) {

    if (r == nil) {
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

    meta := []models.MetadataItem{
        { Key: "comment", Value: comment },
        { Key: "modtime", Value: fmt.Sprint(modtime) },
        { Key: "name", Value: name },
    }

    p.Metadata = meta

    return buf, nil
}
func NewGzipParser() parsers.Parser {
    return &GzipParser{}
}

func init() {
    parsers.RegisterParser("GZIP", NewGzipParser)
}

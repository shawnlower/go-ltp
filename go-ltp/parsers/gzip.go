package parsers

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"

    "bytes"
	"fmt"
	"io"

    "compress/gzip"
    log "github.com/sirupsen/logrus"
)


type GzipParser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *GzipParser) Parse(r models.Reader) (models.Reader, error) {

    buf := new(bytes.Buffer)
    gzipWriter := gzip.NewWriter(buf)

    io.Copy(gzipWriter, r)

    var meta models.MetadataItem
    meta.Key = "gzip.comment"
    meta.Value = fmt.Sprintf("%s", "Comment.")

    p.Metadata = append(p.Metadata, meta)
    log.Debug(fmt.Sprintf("gzip metadata: %s", p))

    return buf, nil
}


package parsers

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"

    "bytes"
	"fmt"
	"io"
	// "os"
    "time"

    "compress/gzip"
    log "github.com/sirupsen/logrus"
)


type GzipParser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *GzipParser) Parse(r models.Reader) (models.Reader, error) {

    if (r == nil) {
        panic("GzipParser cannot compress nil input reader %s")
    }
    buf := new(bytes.Buffer)
    gzipWriter := gzip.NewWriter(buf)
    defer gzipWriter.Close()

    gzipWriter.Comment = "comment"
    gzipWriter.Extra = []byte("extra")
    gzipWriter.ModTime = time.Unix(1e8, 0)
    gzipWriter.Name = "name"

    io.Copy(gzipWriter, r)

    var meta models.MetadataItem
    meta.Key = "gzip.comment"
    meta.Value = fmt.Sprintf("%s", "Comment.")

    p.Metadata = append(p.Metadata, meta)
    log.Debug(fmt.Sprintf("gzip metadata: %s", p))

    return buf, nil
}


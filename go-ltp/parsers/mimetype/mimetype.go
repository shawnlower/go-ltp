package mimetype

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"
    "github.com/shawnlower/go-ltp/go-ltp/parsers"

    "errors"
	"fmt"
    "net/http"
	"io"

    log "github.com/sirupsen/logrus"
)


type MimetypeParser struct{
    Metadata []models.MetadataItem
}

func (p *MimetypeParser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *MimetypeParser) GetName() string {
    return "MimetypeParser"
}

func (p *MimetypeParser) Parse(r io.Reader) (io.Reader, error) {

    if (r == nil) {
        return nil, errors.New("Unable to use nil input reader")
    }

    buf := make([]byte, 512)

    // We only need the first 512 bytes
    r.Read(buf)

    // Read out the rest of the file
    for {
        _, err := r.Read(buf)
        if (err == io.EOF) {
            break
        }
    }
    mimetype := http.DetectContentType(buf)

    p.Metadata = []models.MetadataItem{
        { "mime-type": mimetype },
    }

    log.Debug(fmt.Sprintf("%s metadata: %s", p.GetName(), p.GetMetadata()))

    return r, nil
}
func NewMimetypeParser() models.Parser {
    return &MimetypeParser{}
}

func init() {
    parsers.RegisterParser("MIMETYPE", NewMimetypeParser)
}

package mimetype

import (
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"
	"github.com/shawnlower/go-ltp/parsers"

	"errors"
	"io"
	"net/http"
)

// Implements a MIME-type parser, to detect the content type of
// a file or stream.

// Using net/http, the implementation of the 'sniffing' logic is
// implemented in <https://golang.org/src/net/http/sniff.go>

type MimetypeParser struct {
	Metadata []models.MetadataItem
}

func (p *MimetypeParser) GetMetadata() []models.MetadataItem {
	return p.Metadata
}

func (p *MimetypeParser) GetName() string {
	return "MimetypeParser"
}

func (p *MimetypeParser) Parse(r io.Reader) (io.Reader, error) {

	if r == nil {
		return nil, errors.New("Unable to use nil input reader")
	}

	// DetectContentType() uses at most 512 bytes
	buf := make([]byte, 512)

	r.Read(buf)
	mimetype := http.DetectContentType(buf)

	// Read to EOF so we don't block later
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
	}

	p.Metadata = []models.MetadataItem{
		{"mime-type": mimetype},
	}

	return r, nil
}

func NewMimetypeParser() models.Parser {
	return &MimetypeParser{}
}

func init() {
	parsers.RegisterParser("MIMETYPE", NewMimetypeParser)
}

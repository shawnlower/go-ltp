package models

import (
	"io"
	"github.com/shawnlower/go-ltp/api"
)

type Input struct {
	Name     string
	Reader   io.Reader
	Metadata []MetadataItem
    Item     *api.Item
}

type MetadataItem map[string]string

type JsonMetadata struct {
	Source struct {
		Name string `json:"name"`
	} `json:"source"`
	Metadata []JsonMetaItem `json:"metadata"`
}

type JsonMetaItem struct {
	Parser string         `json:"parser"`
	Type   string         `json:"type"`
	Items  []MetadataItem `json:"items"`
}

type Parser interface {
	Parse(r io.Reader) (io.Reader, error)
	GetName() string
	GetMetadata() []MetadataItem
}

package models

import (
    "os"
)

type MetadataItem struct {
    Key string
    Value string
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type StdinReader struct {
    Name string
    Metadata []MetadataItem
    reader Reader
}

func NewReader(options []interface{}) (*StdinReader, error) {
    r := new(StdinReader)

    r.Name = "StdinReader"
    r.reader = os.Stdin

    var meta MetadataItem
    meta.Key = "stdinKey"
    meta.Value = "stdinVal"

    r.Metadata = append(r.Metadata, meta)

    return r, nil
}

func (r *StdinReader) Read(p []byte) (n int, err error) {
    return os.Stdin.Read(p)
}

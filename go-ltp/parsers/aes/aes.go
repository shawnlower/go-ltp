package aes

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"
    "github.com/shawnlower/go-ltp/go-ltp/parsers"

	"bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"io"
    "fmt"

    log "github.com/sirupsen/logrus"
)

/*
Parser for in-line encryption.
Uses AES-256 GCM w/ 96-bit random nonce, prepended to output

To validate a gzip + aes pipeline:
$ go run ./util/decrypt.go ENCRYPTION_KEY ./outfile | zcat

References:
    https://golang.org/pkg/crypto/cipher/#NewGCM
*/

type AESParser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *AESParser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *AESParser) GetName() string {
    return "AESParser"
}

func (p *AESParser) Parse(r io.Reader) (io.Reader, error) {
	message := new(bytes.Buffer)
    message.ReadFrom(r)

    key := []byte("the-key-has-to-be-32-bytes-long!")
    log.Debug(fmt.Sprintf("Using key: %x", key))

    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    // allocate a buffer with enough storage (initially) to hold the nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    buf := make([]byte, gcm.NonceSize())

    copy(buf, nonce)

    buf = gcm.Seal(buf, buf[:gcm.NonceSize()], message.Bytes(), nil)

    log.Debug(fmt.Sprintf("Using nonce: %x", nonce))
    // log.Debug(fmt.Sprintf("ciphertext: %x", buf))
    return bytes.NewBuffer(buf), nil
}

func NewAESParser() parsers.Parser {
    return &AESParser{}
}

func init() {
    parsers.RegisterParser("AES", NewAESParser)
}

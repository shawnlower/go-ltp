package aes

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"
    "github.com/shawnlower/go-ltp/go-ltp/parsers"

	"bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
	"io"
    "fmt"

	"github.com/spf13/viper"
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

const (
    KEYSIZE = 32
)

type AESParser struct{
    Key string
    Metadata []models.MetadataItem
}

func (p *AESParser) GetMetadata() []models.MetadataItem {
    return p.Metadata
}

func (p *AESParser) GetName() string {
    return "AESParser"
}

func (p *AESParser) Parse(r io.Reader) (io.Reader, error) {
    var key []byte
    var err error

	message := new(bytes.Buffer)
    message.ReadFrom(r)

    k := viper.GetString("parsers.aes.key")
    if (k[:2] == "0x") {
        key, err = hex.DecodeString(k[2:])
        if (err != nil) {
            panic(err)
        }
    } else {
        key = []byte(k)
    }

    if (len(key) != KEYSIZE) {
        panic(fmt.Sprintf("Key must be exactly %d bytes (got %d)",
            KEYSIZE, len(key)))
    }
    log.Debug(fmt.Sprintf("Using key: %s", hex.EncodeToString(key)))

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

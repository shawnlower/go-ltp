package parsers

import (
    "github.com/shawnlower/go-ltp/go-ltp/models"

    "crypto/sha256"
    "crypto/sha512"
	"fmt"
	"io"
    "sync"

    log "github.com/sirupsen/logrus"
)

type Parser interface {
    Parse(r models.Reader) (models.Reader, error)
}

type Sha256Parser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *Sha256Parser) Parse(r models.Reader) (models.Reader, error) {
    h := sha256.New()
    _, err := io.Copy(h, r)

    var meta models.MetadataItem
    meta.Key = "sha256sum"
    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    p.Metadata = append(p.Metadata, meta)
    log.Debug(fmt.Sprintf("sha256 metadata: %s", p))
    return nil, err
}

type Sha512Parser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *Sha512Parser) Parse(r models.Reader) (models.Reader, error) {
    h := sha512.New()
    _, err := io.Copy(h, r)

    var meta models.MetadataItem
    meta.Key = "sha512sum"
    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    meta.Value = fmt.Sprintf("%x", h.Sum(nil))

    p.Metadata = append(p.Metadata, meta)
    log.Debug(fmt.Sprintf("sha256 metadata: %s", p))
    return nil, err
}

type Pipe struct {
    r *io.PipeReader;
    w *io.PipeWriter
}

type CounterParser struct{
    Name string
    Metadata []models.MetadataItem
}

func (p *CounterParser) Parse(reader models.Reader) (models.Reader, error) {

    buf := make([]byte, 1024)
    ctr := 0
    for {
        n, err := reader.Read(buf)
        if (err != nil && err != io.EOF) {
            fmt.Println("Read error.")
            break
        }
        ctr += n
        if (err == io.EOF) {
            // Write any remaining data, close the writer and break
            break
        }
    }

    var meta models.MetadataItem
    meta.Key = "CounterParser.bytes"
    meta.Value = fmt.Sprintf("%d", ctr)

    p.Metadata = append(p.Metadata, meta)
    log.Debug(fmt.Sprintf("counter: %s", p))

    return nil, nil
}




func FanoutParsers(reader io.Reader, parsers []Parser) (err error) {

    buf := make([]byte, 1024)

    var pipes []Pipe

    for i := 0; i < len(parsers); i++ {
        var p Pipe
        p.r, p.w = io.Pipe()
        pipes = append(pipes, p)
    }

    var wg sync.WaitGroup
    for i := range parsers {
        wg.Add(1)
        f := func(i int) {
            // Call our parser, with its own pipe reader
            // we'll use the writer after all goroutines are launched
            _, err := parsers[i].Parse(pipes[i].r)
            if (err != nil) {
                fmt.Println("Error writing pipe ", i, err)
            }
            wg.Done()
        }
        go f(i)
    }

    for {
        n, err := reader.Read(buf)
        if (err != nil && err != io.EOF) {
            fmt.Println("Read error.")
            break
        }
        if (err == io.EOF) {
            // Write any remaining data, close the writer and break
            for i, _ := range parsers {
                pipes[i].w.Write(buf[:n])
                pipes[i].w.Close()
            }
            break
        }
        for i, _ := range parsers {
            pipes[i].w.Write(buf[:n])
        }
    }
    wg.Wait()
    return nil
}

func SerialParsers(reader io.Reader, parsers []Parser) (r models.Reader, err error) {

    var r1, r2 models.Reader

    log.Debug(fmt.Sprintf("Received %d parsers (%s)", len(parsers), parsers))
    log.Debug(fmt.Sprintf("Connecting input %s to first pipe", reader))

    parser := parsers[0]
    r1, err = parser.Parse(reader)

    var i int
    for i = 1; i < len(parsers); i++ {
        parser = parsers[i]
        if (r1 == nil) {
            panic("Received nil reader from parser list. Cannot continue.")
        }
        r2, err = parser.Parse(r1)
        if (err != nil) {
            panic(fmt.Sprintf("Failed to parse pipe %#v", parser))
        }
        r1 = r2
    }
    log.Debug(fmt.Sprintf("Returning I/O stream from %#v", parser))
    return r1, nil
}
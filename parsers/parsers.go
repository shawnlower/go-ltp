package parsers

import (
	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"

	"fmt"
	"io"
	"sync"

	log "github.com/sirupsen/logrus"
)

var parsers = make(map[string]models.Parser)

func RegisterParser(name string, f func() models.Parser) {
	_, ok := parsers[name]
	if ok != false {
		log.Warning("Re-registering parser", name)
	}
	parsers[name] = f()
}

func GetParser(name string) models.Parser {
	parser, ok := parsers[name]
	if ok != true {
		panic(fmt.Sprintf("Unable to find parser %s. Parsers: %#v.",
			name, parsers))
	}
	return parser
}

type Pipe struct {
	r *io.PipeReader
	w *io.PipeWriter
}

func FanoutParsers(reader io.Reader, parsers []models.Parser) (err error) {

	if len(parsers) < 1 {
		return nil
	}

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
			if err != nil {
				fmt.Println("Error writing pipe ", i, err)
			}
			wg.Done()
		}
		go f(i)
	}

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Read error.")
			break
		}
		if err == io.EOF {
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

func SerialParsers(reader io.Reader, parsers []models.Parser) (r io.Reader, err error) {

	if len(parsers) < 1 {
		return nil, nil
	}

	var r1, r2 io.Reader

	parser := parsers[0]
	r1, err = parser.Parse(reader)

	var i int
	for i = 1; i < len(parsers); i++ {
		parser = parsers[i]
		if r1 == nil {
			panic("Received nil reader from parser list. Cannot continue.")
		}
		r2, err = parser.Parse(r1)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse pipe %#v", parser))
		}
		r1 = r2
	}
	log.Debug(fmt.Sprintln("Returning I/O stream from", parser.GetName()))
	return r1, nil
}

// Converts a slice of metadata map[string]string into
// ... that can be used with the semantic store
//
// Photo 
// Subject:
// Predicate: metadata

func MetadataToStatements(mm models.Metadata) {
    for _, m := range(mm) {
        s := &api.Statement{
            Subject: "",
            Predicate: "",
            Object: "",
            Scope: nil,
        }
        log.Debugf("MetadataItem(%#v) Statement(%#v)", m, s)
    }

}

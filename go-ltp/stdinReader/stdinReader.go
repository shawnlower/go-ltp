// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stdinReader

import (
    "crypto/sha256"
    "fmt"
    "io"
    "os"
    "sync"

    log "github.com/sirupsen/logrus"
)

const BUFSIZE = 1024

// Must be >= SHA block size (64 for sha256)
// const BUFSIZE = 64

func genHash(r io.Reader, wg *sync.WaitGroup) {
    h   := sha256.New()
    if _, err := io.Copy(h, r); err != nil {
        log.Error(fmt.Sprintf("Unable to generate hash: %#U\n", err))
    }
    log.Debug("Generated hash: ", h.Sum(nil))
    wg.Done()
}

func Handle(done chan bool) {

    // Use a waitgroup to synchronize procs
    var wg sync.WaitGroup

    r, w := io.Pipe()

    wg.Add(1)
    go genHash(r, &wg)

    log.Info("Reading...")
    fd := os.Stdin

    out := make(chan []byte, BUFSIZE)
    defer close(out)

    cnt := 0
    buf := make([]byte, BUFSIZE)
    for {
        n, err := fd.Read(buf)
        w.Write(buf)

        cnt += n
        if err != nil && err != io.EOF {
            log.Error(fmt.Sprintf("Error reading file. Error: %#v", err))
        }

        // log.Debug("Writing: ", string(buf))
        out <- buf

        if err == io.EOF {
            break
        }
    }
    log.Info(fmt.Sprintf("Read %d bytes.", cnt))
    testFunc(out)

    w.Close()
    wg.Wait()
    done<- true
}

func testFunc(in chan []byte) {
    for done := false; done; {
        select {
        case data := <- in:
            if len(data) == 0 {
                done = true
            }
            fmt.Printf("Received %d bytes.", len(data))
        default:
            continue
        }
    }
}

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

    // Increment waitgroup and spawn a goroutine to calculate the hash
    // of the data written into the pipe
    wg.Add(1)
    go genHash(r, &wg)

    fd := os.Stdin

    // 'out' channel receives all data that we write to the 'w' pipe above
    out := make(chan byte, BUFSIZE)

    wg.Add(1)
    go func() {
        c1 := testFunc1(out)
        c2 := testFunc2(c1)
        cnt := 0
        for c := range c2 {
            _ = c
            cnt++
        }
        log.Info(fmt.Sprintf("Pipeline finished with %d bytes.", cnt))
        wg.Done()
    }()

    cnt := 0
    buf := make([]byte, BUFSIZE)
    for {
        n, err := fd.Read(buf)
        w.Write(buf)

        cnt += n
        if err != nil && err != io.EOF {
            log.Error(fmt.Sprintf("Error reading file. Error: %#v", err))
        }

        // log.Debug(fmt.Sprintf("Writing %d bytes.", n))
        for i := 0; i < n; i++ {
            out <- buf[i]
        }

        if err == io.EOF {
            break
        }
    }
    log.Debug(fmt.Sprintf("Wrote %d bytes.", cnt))
    close(out)

    w.Close()
    wg.Wait()
    done<- true
}

func testFunc1(in <-chan byte) <-chan byte{

    log.Debug(fmt.Sprintf("testFunc1 ready.\n"))

    out := make(chan byte)

    go func() {

        cnt := 0;
        for n := range in {
            out <-n
            cnt++
        }
        close(out)
        log.Debug(fmt.Sprintf("testFunc1 read %d bytes.", cnt))
    }()
    return out
}

func testFunc2(in <-chan byte) <-chan byte{
    log.Debug(fmt.Sprintf("testFunc2 ready.\n"))

    out := make(chan byte)

    go func() {

        cnt := 0;
        for n := range in {
            out <-n
            cnt++
        }
        close(out)
        log.Debug(fmt.Sprintf("testFunc2 read %d bytes.", cnt))
    }()
    return out
}

/*
func testFunc2(in <-chan byte) <-chan byte{

    out := make(chan byte)

    cnt := 0;
    for {
        fmt.Printf("2")
        var data byte
        var more bool
        select {
        case data, more = <-in:
            out <-data
        default:
            continue
        }
        if more == false {
            break
        }
        cnt++
    }
    log.Debug(fmt.Sprintf("testFunc2 read %d bytes.", cnt))
    return out
}
*/


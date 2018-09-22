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

package cmd

import (
    "fmt"
    "io"
    "os"
    "sync"

    "github.com/shawnlower/go-ltp/go-ltp/parsers"

    log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to the repository",
    Long: `
Add an item to the repository:

New items can be created from multiple sources, such as:
  - Local files and directories
  - Dereferenceable URLs
  - Standard input (pipes, etc)

Examples:
    # add from standard-input
    xclip -o | go-ltp add -

    # add from a local file
    go-ltp add ~/Pictures/Jamaica_Bay.png
`,
	Run: func(cmd *cobra.Command, args []string) {

        // Attempt to parse each of the inputs provided
        if len(args) < 1 {
            log.Fatal("No inputs specified. Use '-' for stdin.")
            os.Exit(1)
        }

        // array for holding multiple inputs
        var readers []io.Reader

        // Setup our initial list of parsers. All inputs should support at
        // least the following
        sha512 := parsers.Sha512Parser{}
        counterParser := parsers.CounterParser{}
        parserList := []parsers.Parser{&sha512, &counterParser}

        for _, inputString := range args {
            switch inputString {
            case "-":
                log.Info("Reading from stdin...")
                readers = append(readers, os.Stdin)

            default:
                // Unknown input type
                log.Fatal("Input type not supported: " + inputString)
                os.Exit(1)
            }
        }

        /* Main loop; iterate across all of our readers and do the following:

            1) Split the input into two `io.Reader`s. The first is processed
               as part of a fanout pipeline by FanoutParsers(), each parser
               reading the stream and outputting metadata, such as a hash,
               the filesystem metadata, or the OS process metadata for the
               remote end of the stdin pipe.
            2) The second stream is a serial pipeline, which passes the data
               through a sequence of parsers.
               Example:
                 {source stream} -> {compression parser} -> {encryption parser}
            3) Finally, the output stream and metadata are written.
        */
        for idx, reader := range(readers) {

            // Create a pipe for the serial pipeline
            serialPipeR, serialPipeW := io.Pipe()
            tr := io.TeeReader(reader, serialPipeW)

            var wg sync.WaitGroup

            // Reader 1
            wg.Add(1)
            go func() {
                log.Debug(fmt.Sprintf("Running parsers: %s", parserList))
                err := parsers.FanoutParsers(serialPipeR, parserList)
                if (err != nil) {
                    log.Error("Failed to parse")
                }
                wg.Done()
            }()

            // Serial parsing pipeline ( input -> compression -> encryption )
            gzipParser := parsers.GzipParser{}
            // gzipParser2 := parsers.GzipParser{}
            // serialParsers := []parsers.Parser{&gzipParser, &gzipParser2}
            serialParsers := []parsers.Parser{&gzipParser}

            log.Debug(fmt.Sprint("Running serial parsing pipeline"))
            outReader, err := parsers.SerialParsers(tr, serialParsers)
            if (err != nil) {
                log.Error("Failed to parse")
            }

            // The pipe must be closed to allow all readers to exit
            serialPipeW.Close()
            wg.Wait()

            log.Debug(fmt.Printf("Reader %#v", reader))
            // Write out metadata
            for i, mdi := range(parserList) {
                log.Debug(fmt.Sprintf("Meta for fanout parser %d: %#v", i, mdi))
            }
            for i, mdi := range(serialParsers) {
                log.Debug(fmt.Sprintf("Meta for serial parser %d: %#v", i, mdi))
            }

            // Output pipeline (disk, network, etc)
            filename := fmt.Sprintf("outfile.%d.data", idx)
            fileWriter(outReader, filename)
        }
	},
}

func fileWriter(r io.Reader, filename string) (err error) {
    outfile, err := os.Create(filename)
    defer outfile.Close()

    if (err != nil) {
        panic(fmt.Sprintln("Unable to create output file", filename))
    }
    log.Debug(fmt.Println("Writing output to file:", filename))
    io.Copy(outfile, r)

    return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

}

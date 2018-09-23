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
    _ "github.com/shawnlower/go-ltp/go-ltp/parsers/aes"
    _ "github.com/shawnlower/go-ltp/go-ltp/parsers/counter"
    _ "github.com/shawnlower/go-ltp/go-ltp/parsers/mimetype"
    _ "github.com/shawnlower/go-ltp/go-ltp/parsers/gzip"
    _ "github.com/shawnlower/go-ltp/go-ltp/parsers/sha256"
    _ "github.com/shawnlower/go-ltp/go-ltp/parsers/sha512"

    log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

        // Setup our initial list of parsers.

        var asyncParsers []parsers.Parser
        var parserNames []string
        for _, parserName := range(viper.GetStringSlice("parsers.async")) {
            parser := parsers.GetParser(parserName)
            asyncParsers = append(asyncParsers, parser)
            parserNames = append(parserNames, parser.GetName())
        }
        log.Debug("Added async parsers: ", parserNames)

        var serialParsers []parsers.Parser
        parserNames = nil
        for _, parserName := range(viper.GetStringSlice("parsers.serial")) {
            parser := parsers.GetParser(parserName)
            serialParsers = append(serialParsers, parser)
            parserNames = append(parserNames, parser.GetName())
        }
        log.Debug("Added serial parsers: ", parserNames)

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
                err := parsers.FanoutParsers(serialPipeR, asyncParsers)
                if (err != nil) {
                    log.Error("Failed to parse")
                }
                wg.Done()
            }()

            // Serial parsing pipeline ( input -> compression -> encryption )
            outReader, err := parsers.SerialParsers(tr, serialParsers)
            if (err != nil) {
                log.Error("Failed to parse")
            }

            // The pipe must be closed to allow all readers to exit
            serialPipeW.Close()
            wg.Wait()

            // Write out metadata
            for _, parser := range(asyncParsers) {
                name := parser.GetName()
                log.Debug(fmt.Sprintln("Metadata for fanout parser", name))
                for _, mdi := range(parser.GetMetadata()) {
                    log.Debug(fmt.Sprintf("\t%12s = %s", mdi.Key, mdi.Value))
                }
            }
            for _, parser := range(serialParsers) {
                name := parser.GetName()
                log.Debug(fmt.Sprintln("Metadata for serial parser", name))
                for _, mdi := range(parser.GetMetadata()) {
                    log.Debug(fmt.Sprintf("\t%12s = %s", mdi.Key, mdi.Value))
                }
            }

            // Output pipeline (disk, network, etc)
            filename := fmt.Sprintf("outfile.%d.data", idx)
            fileWriter(outReader, filename)

            /*
            We can now handle the metadata (including any output
            meta-data, such as output filename, s3/gcs URL, etc).

            */
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

}

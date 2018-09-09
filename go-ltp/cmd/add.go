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
    "github.com/shawnlower/go-ltp/go-ltp/models"

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
        var readers []models.Reader

        for _, inputString := range args {
            switch inputString {
            case "-":
                log.Info("Reading from stdin...")
                stdin, err := models.NewReader(nil)
                if (err != nil) {
                    log.Error("Error creating stdinReader: ", err)
                }
                log.Debug("Created stdinReader: ", stdin)
                readers = append(readers, stdin)

            default:
                // Unknown input type
                log.Fatal("Input type not supported: " + inputString)
                os.Exit(1)
            }
        }

        /* Main loop; iterate across all of our readers and do the following:

        1) Split our input into 2 io.Readers : tr, serialPipeR
        2) Pass the 


        */
        for _, reader := range(readers) {

            // Add initial metadata (timestamp, etc)


            // Split reader and run fanout parsing in parallel

            // Create a pipe for the serial pipeline
            serialPipeR, serialPipeW := io.Pipe()
            tr := io.TeeReader(reader, serialPipeW)

            // Collect metadata (hash, process metadata)
            sha256 := parsers.Sha256Parser{}
            sha512 := parsers.Sha512Parser{}
            parserList := []parsers.Parser{&sha256, &sha512}

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
            wg.Add(1)
            go func() {
                gzipParser := parsers.GzipParser{}
                gzipParser2 := parsers.GzipParser{}
                counterParser := parsers.CounterParser{}
                // serialParsers := []parsers.Parser{&gzipParser, &counterParser, &sha256}
                serialParsers := []parsers.Parser{&gzipParser, &gzipParser2, &counterParser}

                log.Debug(fmt.Sprint("Running serial parsing pipeline"))
                err := parsers.SerialParsers(tr, serialParsers)
                if (err != nil) {
                    log.Error("Failed to parse")
                }

                // The pipe must be closed to allow all readers to exit
                serialPipeW.Close()
                wg.Done()
            }()

            wg.Wait()
        }

        // Output pipeline (disk, network, etc)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

}

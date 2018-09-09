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
    "os"

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

        log.Debug("Running parsers...")
        for _, reader := range(readers) {
            var p parsers.Sha256Parser
            err := p.Parse(reader)
            if (err != nil) {
                log.Error("Failed to parse")
            }
        }

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

}

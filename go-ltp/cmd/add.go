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
    "os"

    "github.com/shawnlower/go-ltp/go-ltp/stdinReader"

    log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type inputHandler func(chan bool)

type job struct {
    name    string
    h       inputHandler
    done    chan bool
}

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

        var jobs []job

        // Attempt to parse each of the inputs provided
        if len(args) < 1 {
            log.Fatal("No inputs specified. Use '-' for stdin.")
            os.Exit(1)
        }

        // Queue a job to handle each input specified as an argument
        for _, inputString := range args {
            switch inputString {
            case "-":
                log.Info("Reading from stdin...")
                var c = make(chan bool, 1)
                jobs = append(jobs, job{name: "-",
                                        h: stdinReader.Read,
                                        done: c})

            default:
                // Unknown input type
                log.Fatal("Input type not supported: " + inputString)

                os.Exit(1)
            }
        }

        // Launch each job by calling its associated handler function
        for _, job := range(jobs) {
            log.Info(fmt.Sprintf("Launching job: %#v", job))

            // Use the 'done' channel to signify job completion
            go job.h(job.done)
        }

        // Wait for each job to complete
        for _, job := range(jobs) {
            log.Info(fmt.Sprintf("Waiting for job: %#v", job))

            <-job.done
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

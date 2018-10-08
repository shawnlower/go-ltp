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

package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common/models"
	"github.com/shawnlower/go-ltp/parsers"
	_ "github.com/shawnlower/go-ltp/parsers/aes"
	_ "github.com/shawnlower/go-ltp/parsers/counter"
	_ "github.com/shawnlower/go-ltp/parsers/gzip"
	_ "github.com/shawnlower/go-ltp/parsers/mimetype"
	_ "github.com/shawnlower/go-ltp/parsers/sha256"
	_ "github.com/shawnlower/go-ltp/parsers/sha512"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

var dryRun bool
var config common.ClientConfig

// addCmd represents the add command
func NewAddCommand() *cobra.Command {
	var cmd = &cobra.Command{
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
		PreRun: func(cmd *cobra.Command, args []string) {
            dryRun = viper.GetBool("outputs.dry-run")
		},
		Run: func(cmd *cobra.Command, args []string) {

			// Attempt to parse each of the inputs provided
			if len(args) < 1 {
				log.Fatal("No inputs specified. Use '-' for stdin.")
				os.Exit(1)
			}

			// array for holding multiple inputs
			var inputs []models.Input

			// Setup our initial list of parsers.

			var asyncParsers []models.Parser
			var parserNames []string
			for _, parserName := range viper.GetStringSlice("parsers.async") {
				parser := parsers.GetParser(parserName)
				asyncParsers = append(asyncParsers, parser)
				parserNames = append(parserNames, parser.GetName())
			}
			log.Debug("Added async parsers: ", parserNames)

			var serialParsers []models.Parser
			parserNames = nil
			for _, parserName := range viper.GetStringSlice("parsers.serial") {
				parser := parsers.GetParser(parserName)
				serialParsers = append(serialParsers, parser)
				parserNames = append(parserNames, parser.GetName())
			}
			log.Debug("Added serial parsers: ", parserNames)

			for _, inputString := range args {
				if inputString == "-" {
					log.Info("Reading from stdin...")
					inputs = append(inputs,
						models.Input{Name: "stdin", Reader: os.Stdin})
				} else if m, _ := regexp.MatchString("https?://", inputString); m {
					// Call HTTP fetch module to retrieve page
					log.Fatal("URLs not yet supported.")
					os.Exit(1)
				} else {
					// Assume the argument is a file; fail if it can't be opened
					fd, err := os.Open(inputString)
					if os.IsNotExist(err) {
						log.Fatal(err)
						os.Exit(1)
					} else if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}

					// Ensure this is a file, not a directory
					st, _ := fd.Stat()
					if st.IsDir() {
						log.Fatal("Directories not yet supported.")
						os.Exit(1)
					}

					input := models.Input{Name: "file", Reader: fd}

					// Add the parsers to the input object
					input.Metadata = append(input.Metadata,
						models.MetadataItem{
							"filename": fd.Name(),
						})
					inputs = append(inputs, input)
				}
			}

			/* Main loop; iterate across all of our inputs and do the following:

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
			for _, input := range inputs {

				// Create a pipe for the serial pipeline
				serialPipeR, serialPipeW := io.Pipe()
				tr := io.TeeReader(input.Reader, serialPipeW)

				var wg sync.WaitGroup

				// Reader 1
				wg.Add(1)
				go func() {
					err := parsers.FanoutParsers(serialPipeR, asyncParsers)
					if err != nil {
						log.Error("Failed to parse")
					}
					wg.Done()
				}()

				// Serial parsing pipeline ( input -> compression -> encryption )
				outReader, err := parsers.SerialParsers(tr, serialParsers)
				if err != nil {
					log.Error("Failed to parse")
				}

				// The pipe must be closed to allow all inputs to exit
				serialPipeW.Close()
				wg.Wait()

				// Construct filename based on hash of the contents.
				var datafile, metadatafile string
				for _, parser := range asyncParsers {
					for _, mdi := range parser.GetMetadata() {
						if mdi["hash"] != "" {
							datafile = fmt.Sprintf("%s.data", mdi["hash"])
							metadatafile = fmt.Sprintf("%s.json", mdi["hash"])
							log.Debug("Got filename ", datafile)
						}
					}
				}
				if datafile == "" {
					datafile = fmt.Sprintf("output.data")
					metadatafile = fmt.Sprintf("output.json")
				}
				fileWriter(outReader, datafile)

				/*
				   We can now handle the metadata (including any output
				   meta-data, such as output filename, s3/gcs URL, etc).
				*/

				jsonDoc, err := inputToJson(&input, &asyncParsers, &serialParsers)
				fileWriter(bytes.NewReader(jsonDoc), metadatafile)

                // Setup a client
                client, ctx, err := common.GetClient(cmd)
				remoteWriter(client, ctx, bytes.NewReader(jsonDoc), metadatafile)
			}
		},
	}

	cmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Dry-run only. DO NOT write output")

	viper.BindPFlag("outputs.dry-run", cmd.PersistentFlags().Lookup("dry-run"))
	return cmd
}

func inputToJson(input *models.Input, asyncParsers *[]models.Parser,
	serialParsers *[]models.Parser) (jsonDoc []byte, err error) {

	/*
	   Return a JSON document ( []byte ) from an input, and associated parsers.
	       {
	           "source": {
	               "name": input.Name,
	           },
	           "metadata": [ {
	               "parser": "SHA256",
	               "type": "async",
	               "items": [
	                   { Key: Value },
	                   { Key: Value },
	                   { Key: Value } ]
	           }, {
	               "parser": "AES",
	               "type": "serial",
	               "items": [
	                   { Key: Value },
	                   { Key: Value },
	                   { Key: Value } ]
	         } ] } }
	*/

	// Setup our initial JSON object with the top-level keys
	jmeta := models.JsonMetadata{}

	jmeta.Source.Name = input.Name

	// Add any metadata from the input itself
	if len(input.Metadata) > 0 {
		metadata := models.JsonMetaItem{}
		metadata.Parser = input.Name
		metadata.Type = "input"

		// Add per-parser metadata
		for _, metadataItem := range input.Metadata {
			metadata.Items = append(metadata.Items, metadataItem)
			log.Debug(fmt.Sprintf("meta: input.%s %#v", input.Name,
				metadataItem))
		}
		jmeta.Metadata = append(jmeta.Metadata, metadata)
	}

	// Add async
	for _, parser := range *asyncParsers {
		metadata := models.JsonMetaItem{}
		metadata.Parser = parser.GetName()
		metadata.Type = "async"

        // Test fetching meta
        parsers.MetadataToStatements(parser.GetMetadata())
        // log.Debugf("*** METADATA(%#v) - Statements(%#v)\n", metadataItem,

		// Add per-parser metadata
		for _, metadataItem := range parser.GetMetadata() {
			metadata.Items = append(metadata.Items, metadataItem)
			log.Debug(fmt.Sprintf("meta: async.%s %#v", input.Name,
				metadataItem))
		}
		jmeta.Metadata = append(jmeta.Metadata, metadata)
	}

	// Add serial
	for _, parser := range *serialParsers {
		metadata := models.JsonMetaItem{}
		metadata.Parser = parser.GetName()
		metadata.Type = "serial"

		// Add per-parser metadata
		for _, metadataItem := range parser.GetMetadata() {
			metadata.Items = append(metadata.Items, metadataItem)
			log.Debug(fmt.Sprintf("meta: serial.%s %#v", input.Name,
				metadataItem))
		}
		jmeta.Metadata = append(jmeta.Metadata, metadata)
	}

	jsonDoc, _ = json.MarshalIndent(jmeta, "", "  ")
	return jsonDoc, nil
}

func remoteWriter(c api.APIClient, ctx context.Context, r io.Reader, f string) (err error) {
	if r == nil {
		log.Fatal("Nothing to write (parser returned empty input?)")
	}

	type Item struct {
		ItemTypeURI string
	}

    req, err := api.NewItemRequest("http://schema.org/Thing")
    if err != nil {
		log.Fatalf("Error calling CreateItemRequest: %v", err)
    }

	resp, err := c.CreateItem(ctx, req)
    if err != nil {
		log.Fatalf("Error calling CreateItem: %v", err)
	}
	log.Printf("Received: %s", resp)

	return nil
}

func fileWriter(r io.Reader, f string) (err error) {
	if r == nil {
		log.Fatal("Nothing to write (parser returned empty input?)")
	}
	basedir := viper.GetString("outputs.file.basedir")
	if strings.HasPrefix(basedir, "~") {
		u, _ := user.Current()
		basedir = filepath.Join(u.HomeDir, basedir[1:])
	}
	_, err = os.Stat(basedir)
	if err != nil {
		log.Fatal("Path error: ", err)
	}
	filename := filepath.Join(basedir, f)

	if dryRun {
		log.Info("`dry-run' specified. NOT WRITING: ", filename)
		return nil
	}

	outfile, err := os.Create(filename)
	defer outfile.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Write output
	_, err = io.Copy(outfile, r)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Wrote output to file:", filename)

	return nil
}

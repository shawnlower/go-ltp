// Copyright © 2018 Shawn Lower <shawn@shawnlower.com>
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
	"github.com/shawnlower/go-ltp/api/proto"
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
		Run: addCommand,
	}

	cmd.Flags().StringP("type", "t", "", "Type of item to create.")
	cmd.Flags().StringP("name", "n", "", "Name of item to create. This becomes the primaryLabel")

	cmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Dry-run only. DO NOT write output")

	viper.BindPFlag("outputs.dry-run", cmd.PersistentFlags().Lookup("dry-run"))
	return cmd
}

func addCommand(cmd *cobra.Command, args []string) {

    var (
        input models.Input
        parserNames []string
    )

    if len(args) == 0 {
        // Calling add with no positional arguments is allowed if
        // we're provided with both:
        //   -n: A name for the new item; this becomes the primaryLabel
        //   -t: A type for the new item
        name, _ := cmd.Flags().GetString("name")
        typeIRI, err := cmd.Flags().GetString("type")

        if name == "" || typeIRI == "" {
            log.Fatal("No inputs specified. Use '-' for stdin.")
        }

        // Create item request
        i, err := api.NewItem(api.IRI(typeIRI))
        if err != nil {
            log.Fatalf("Error creating Item: %v", err)
        }
        req, err := i.ToRequest()
        if err != nil {
            log.Fatalf("Error creating ItemRequest: %v", err)
        }

        // Get client to submit
        c, ctx, err := common.GetClient()
        if err != nil {
            log.Fatalf("did not connect: %v", err)
        }

        resp, err := c.CreateItem(ctx, req)
        if err != nil {
            log.Fatalf("Error calling CreateItem: %v", err)
        }
        log.Printf("Received: %s", resp)

        return
    }

    if len(args) > 1 {
        log.Fatal("Only a single input supported.")
    }

    inputString := args[0]
    if inputString == "-" {
        log.Info("Reading from stdin...")
        item, _ := api.NewItem(api.IRI("schema:Thing"))
        input = models.Input{
            Name: "stdin",
            Reader: os.Stdin,
            Metadata: models.Metadata{"filename": "-"},
            Item: &item,
        }
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

        item, _ := api.NewItem(api.IRI("schema:Thing"))
        input = models.Input{
            Name: "file",
            Reader: fd,
            Metadata: models.Metadata{"filename": fd.Name()},
            Item: &item,
        }
    }

    // Setup our initial list of parsers.
    for _, parserName := range viper.GetStringSlice("parsers.async") {
        parser := parsers.GetParser(parserName)
        input.AsyncParsers = append(input.AsyncParsers, parser)
        parserNames = append(parserNames, parser.GetName())
    }
    log.Debug("Added async parsers: ", parserNames)

    parserNames = nil
    for _, parserName := range viper.GetStringSlice("parsers.serial") {
        parser := parsers.GetParser(parserName)
        input.SerialParsers = append(input.SerialParsers, parser)
        parserNames = append(parserNames, parser.GetName())
    }
    log.Debug("Added serial parsers: ", parserNames)

    if err := handleInput(input); err != nil {
        log.Fatal("Parsing input failed: ", err)
    }
}

func handleInput(input models.Input) error {
    /* Main loop; handle our input

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

    // Create a pipe for the serial pipeline
    serialPipeR, serialPipeW := io.Pipe()
    tr := io.TeeReader(input.Reader, serialPipeW)

    var wg sync.WaitGroup

    // Reader 1
    wg.Add(1)
    go func() {
        err := parsers.FanoutParsers(serialPipeR, input.AsyncParsers)
        if err != nil {
            log.Fatal("Failed to parse")
        }
        wg.Done()
    }()

    // Serial parsing pipeline ( input -> compression -> encryption )
    outReader, err := parsers.SerialParsers(tr, input.SerialParsers)
    if err != nil {
        log.Fatal("Failed to parse")
    }

    // The pipe must be closed to allow all inputs to exit
    serialPipeW.Close()
    wg.Wait()

    // Get the label to use for the item. This may be either
    // - From the -n argument, or
    // - From a 'name', or primaryLabel property

    // Determine the type of the item from the parsers' outputs

    // Pass to any outputs (localfile, client, gcs, etc)

    // Submit the CreateItemRequest


    // Construct filename based on hash of the contents.
    var datafile, metadatafile string
    for _, parser := range input.AsyncParsers {
        metadata := parser.GetMetadata()
        for _, k := range metadata {
            if k == "hash" {
                datafile = fmt.Sprintf("%s.data", metadata[k])
                metadatafile = fmt.Sprintf("%s.json", metadata[k])
                log.Debug("Got filename ", datafile)
            }
        }
    }
    if datafile == "" {
        datafile = fmt.Sprintf("output.data")
        metadatafile = fmt.Sprintf("output.json")
    }
    // fileWriter(outReader, datafile)
    _ = bytes.NewBuffer
    _ = outReader
    _ = metadatafile


    /*
    We can now handle the metadata (including any output
    meta-data, such as output filename, s3/gcs URL, etc).

    'ltpcli add' is a combination of two things:

    1) Create a new Item, referring to some real-world thing,
    such as a person, a book, a restaurant, or an event.
    2) Take any semantic metadata from the parsers, and link
    it to our new item.

    */

    // jsonDoc, err := inputToJson(&input, &asyncParsers, &serialParsers)
    // fileWriter(bytes.NewReader(jsonDoc), metadatafile)

    // Setup a client
    client, ctx, err := common.GetClient()
    if err != nil {
        log.Fatal(err)
    }
    err = remoteWriter(input, client, ctx)
    return err
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
        metadata.Items = input.Metadata
        log.Debug(fmt.Sprintf("meta: input.%s %#v", input.Name,
            input.Metadata))
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
        metadata.Items = parser.GetMetadata()
        log.Debug(fmt.Sprintf("meta: async.%s %#v", input.Name,
            metadata.Items))
		jmeta.Metadata = append(jmeta.Metadata, metadata)
	}

	// Add serial
	for _, parser := range *serialParsers {
		metadata := models.JsonMetaItem{}
		metadata.Parser = parser.GetName()
		metadata.Type = "serial"

		// Add per-parser metadata
        metadata.Items = parser.GetMetadata()
        log.Debug(fmt.Sprintf("meta: serial.%s %#v", input.Name, metadata.Items))
		jmeta.Metadata = append(jmeta.Metadata, metadata)
	}

	jsonDoc, _ = json.MarshalIndent(jmeta, "", "  ")
	return jsonDoc, nil
}

func remoteWriter(in models.Input, c proto.APIClient, ctx context.Context) error {

    for _, parser := range(in.AsyncParsers) {
        metadata := parser.GetMetadata()
        for k, v := range metadata  {
            s := in.Item.IRI
            p := k
            o := v
            g := fmt.Sprintf("ltpcli.%s", parser.GetName())
            log.Debugf("%s.async: <%s> <%s> <%s>", g, s, p, o)
            prop := api.NewProperty(api.IRI(p))
            in.Item.AddProperty(*prop, api.IRI(o))
        }
    }

    for _, parser := range(in.SerialParsers) {
        metadata := parser.GetMetadata()
        for k, v := range metadata  {
            s := fmt.Sprintf("item")
            p := k
            o := v
            g := fmt.Sprintf("ltpcli.%s", parser.GetName())
            log.Debugf("%s.serial: <%s> <%s> <%s>", g, s, p, o)
            prop := api.NewProperty(api.IRI(p))
            in.Item.AddProperty(*prop, api.IRI(o))
        }
    }

    // Determine Item type
    // itemType := in.Item.GetType()
    // itemType := "http://schema.org/Thing"
    log.Warning("Unable to set item type.")

    // Determine Item label

    req, err := in.Item.ToRequest()
    if err != nil {
        log.Fatalf("Error creating ItemRequest: %v", err)
    }

	resp, err := c.CreateItem(ctx, req)
    log.Debugf("Sending request %s with %d statements", req, len(req.Statements))
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

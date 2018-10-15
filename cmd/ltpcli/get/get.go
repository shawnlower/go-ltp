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

package get

import (
	"fmt"

	"github.com/shawnlower/go-ltp/api/proto"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewGetItemCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an item from the remote store",
		Long: `Retrieves an item and any statements related to the item`,
		Run: getItemCommand,
	}

	return cmd
}

func getItemCommand(cmd *cobra.Command, args []string) {
	fmt.Println("get called with ", args)

    if len(args) < 1 {
        log.Fatalf("Required: item IRI")
    }

    IRI := args[0]

	c, ctx, err := common.GetClient()
    getResp, err := c.GetItem(ctx, &proto.GetItemRequest{IRI: IRI})
	if err != nil {
		log.Fatalf("Error in gRPC: %v", err)
	}
	log.Printf("Received via gRPC: %s", getResp.Item)

    if getResp.Item == nil || getResp.Item.Statements == nil {
        log.Fatalf("No item received.")
    }

    log.Printf("Received via gRPC: ")
    for i, statement := range getResp.Item.Statements {
        log.Printf("%d: [%+v]", int(i), statement)
    }
}


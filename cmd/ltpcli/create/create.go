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

package create

import (
    "github.com/shawnlower/go-ltp/api"
    "github.com/shawnlower/go-ltp/cmd/ltpcli/common"

    log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
    DEFAULT_TYPE = "http://schema.org/Thing"
)

// createCmd represents the create command
func NewCreateCommand() *cobra.Command {
    var cmd = &cobra.Command{
        Use:   "create",
        Short: "Create an item in the repository",
        Long: `
        Create a new item - without uploading an object.

        Example:
            $ ltpcli create -t 'http://schema.org/Book' -n 'Sapiens'
            Created http://shawnlower.net/i/Sapiens
        `,
        Run: func(cmd *cobra.Command, args []string) {

            c, ctx, err := common.GetClient()
            if err != nil {
                log.Fatalf("did not connect: %v", err)
            }

            type Item struct {
                ItemTypeURI string;
            }

            typeUri, err := cmd.Flags().GetString("type")

            itemTypes := []*api.ItemType{ }
            itemTypes = append(itemTypes, &api.ItemType{ Uri: typeUri, })

            req := &api.CreateItemRequest{
                ItemTypes: itemTypes,
            }

            resp, err := c.CreateItem(ctx, req)
            if err != nil {
                log.Fatalf("Error calling CreateItem: %v", err)
            }
            log.Printf("Received: %s", resp)
        },
    }

    cmd.PersistentFlags().String("type", DEFAULT_TYPE,
        "Type of item to create. Default: " + DEFAULT_TYPE)
    cmd.MarkFlagRequired("type")

    return cmd
}

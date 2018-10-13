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

package list

import (
	"fmt"

	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/common"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List remote objects",
		Long:  `List the objects available in the remote store.`,
		Run:   listCommand,
	}

	return cmd
}

func listCommand(cmd *cobra.Command, args []string) {
	fmt.Println("list called")

	c, ctx, err := common.GetClient()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	r, err := c.GetVersion(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("Error calling GetVersion: %v", err)
	}
	log.Printf("Received: %s", r.VersionString)
}

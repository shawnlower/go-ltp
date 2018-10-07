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

package info

import (
	"fmt"

    "github.com/shawnlower/go-ltp/api"
    "github.com/shawnlower/go-ltp/cmd/ltpcli/common"

	"github.com/spf13/cobra"
    log "github.com/sirupsen/logrus"
)

func NewInfoCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "info",
        Short: "Display client/server info",
        Long: `Provide information about both the local environment,
    as well as the remote server information, such as server version,
    number of objects, and overall health.`,
        Run:  infoCommand,
    }

    return cmd
}

func infoCommand(cmd *cobra.Command, args []string) {
    fmt.Println("info called with ", args)

    c, ctx, err := common.GetClient()
	r, err := c.GetVersion(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("Error in gRPC: %v", err)
	}
	log.Printf("Received via gRPC: %s", r.VersionString)

}

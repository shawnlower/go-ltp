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

package run

import (
	"fmt"

	"github.com/shawnlower/go-ltp/cmd/ltpd/common/server"

	"github.com/spf13/cobra"
)

var (
	listenAddr string
)

func RunServer(cmd *cobra.Command, args []string) error {
	listenAddr := cmd.Flags().Lookup("listen-addr").Value.String()
	fmt.Println("run called with ", args)
	server.Serve(listenAddr)
	return nil // TODO: Get proper return code
}

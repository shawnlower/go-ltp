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
    "net/url"

	"github.com/shawnlower/go-ltp/cmd/ltpd/common/server"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"
)

var (
	listenAddr string
)

func RunServer(cmd *cobra.Command, args []string) error {
	listenAddr := cmd.Flags().Lookup("listen-addr").Value.String()

    var (
        scheme string
        host string
        port string
        srv *grpc.Server
    )
    _ = srv

	u, err := url.Parse(listenAddr)
	if err != nil {
		log.Fatal(err)
	}

    host = u.Hostname()

	if u.Port() == "" {
        port = "17900"
	} else {
        port = u.Port()
	}

	switch u.Scheme {
	case "grpc":
        scheme = "grpc"
	case "http":
	case "https":
        scheme = "http"

	log.Debug(fmt.Sprintf("Starting %s server on %s port %s", scheme, host, port))
	default:
		panic(fmt.Sprintf("Invalid scheme: %s. Valid schemes are 'grpc', 'http'", u.Scheme))
	}

    if scheme == "grpc" {
        serverCert := cmd.Flags().Lookup("cert").Value.String()
		// log.Debug(serverCert)
        serverKey := cmd.Flags().Lookup("key").Value.String()
        caCert := cmd.Flags().Lookup("ca-cert").Value.String()

        srv, err = server.NewMutualTLSGrpcServer(host, port, serverCert, serverKey, caCert)
        // return NewInsecureGrpcServer(host, port)
    }

	return nil // TODO: Get proper return code
}

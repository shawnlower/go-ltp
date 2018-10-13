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
    "net/url"
    "strings"

	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/cmd/ltpd/common/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

var (
	listenAddr string
)

type ServerConfig struct {
	ServerUrl string;
    AuthMethod string;
    ServerCert string;
    ServerKey string;
    CACert string;
}

func RunServer(cmd *cobra.Command, args []string) error {
    config := &ServerConfig{
        ServerUrl : strings.ToLower(viper.GetString("server.listen-addr")),
        ServerCert : viper.GetString("server.cert"),
        ServerKey : viper.GetString("server.key"),
        CACert : viper.GetString("server.ca-cert"),
        AuthMethod : strings.ToLower(viper.GetString("remote.auth")),
    }

	u, err := url.Parse(config.ServerUrl)
	if err != nil {
		log.Fatal(err)
	}

    var (
        host string
        port string
    )

    host = u.Hostname()

	if u.Port() == "" {
        port = "17900"
	} else {
        port = u.Port()
	}
    log.Debug(u, config)

    if u.Scheme == "grpc" {
        if config.AuthMethod == "mutual-tls" {
            _, err = server.NewMutualTLSGrpcServer(host, port,
                config.ServerCert, config.ServerKey, config.CACert)
            return err
        } else if config.AuthMethod == "insecure" {
            _, err := server.NewInsecureGrpcServer(host, port)
            return err
        } else {
            return &api.ErrInvalidAuthMethod{Method: config.AuthMethod}
        }
    } else {
        return &api.ErrInvalidScheme{Scheme: u.Scheme}
    }
}

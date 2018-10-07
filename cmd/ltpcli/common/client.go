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

package common

import (
	"fmt"
	"net/url"
	"time"

	"github.com/shawnlower/go-ltp/api"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	serverUrl = "grpc://127.0.0.1:17900"
)

func GetClient() (c api.APIClient, ctx context.Context, err error) {

	var (
		host string
		port string
	)

	u, err := url.Parse(serverUrl)
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
		return getGrpcClient(host, port)

	default:
		panic(fmt.Sprintf("Invalid scheme: %s. Valid schemes are 'grpc', '...'", u.Scheme))
	}
}

func getGrpcClient(host string, port string) (c api.APIClient, ctx context.Context, err error) {

	var (
		address string
	)

	address = fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c = api.NewAPIClient(conn)

	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)

	return c, ctx, nil

}

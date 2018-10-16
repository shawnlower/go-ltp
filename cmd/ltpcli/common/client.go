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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/shawnlower/go-ltp/api"
	"github.com/shawnlower/go-ltp/api/proto"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientConfig struct {
	ServerUrl  string
	AuthMethod string
	ClientCert string
	ClientKey  string
	CACert     string
}

func GetClient() (proto.APIClient, context.Context, error) {
	config := &ClientConfig{
		ServerUrl:  strings.ToLower(viper.GetString("remote.url")),
		ClientCert: viper.GetString("remote.cert"),
		ClientKey:  viper.GetString("remote.key"),
		CACert:     viper.GetString("remote.ca-cert"),
		AuthMethod: strings.ToLower(viper.GetString("remote.auth")),
	}

	u, err := url.Parse(config.ServerUrl)
	if err != nil {
		log.Fatal(err)
	}

	if u.Scheme == "grpc" {
		if config.AuthMethod == "mutual-tls" {
			return getMutualTLSGrpcClient(config)
		} else if config.AuthMethod == "insecure" {
			return getInsecureGrpcClient(config)
		} else {
			return nil, nil, &api.ErrInvalidAuthMethod{Method: config.AuthMethod}
		}
	} else {
		return nil, nil, &api.ErrInvalidScheme{Scheme: u.Scheme}
	}
}

func getInsecureGrpcClient(cfg *ClientConfig) (proto.APIClient, context.Context, error) {

	var (
		host string
		port string
	)

	u, err := url.Parse(cfg.ServerUrl)
	if err != nil {
		log.Fatal(err)
	}

	host = u.Hostname()

	if u.Port() == "" {
		port = "17900"
	} else {
		port = u.Port()
	}

	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := proto.NewAPIClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	return c, ctx, nil
}

func getMutualTLSGrpcClient(cfg *ClientConfig) (proto.APIClient, context.Context, error) {

	var (
		host string
		port string
	)

	u, err := url.Parse(cfg.ServerUrl)
	if err != nil {
		log.Fatal(err)
	}

	host = u.Hostname()

	if u.Port() == "" {
		port = "17900"
	} else {
		port = u.Port()
	}

	certificate, err := tls.LoadX509KeyPair(cfg.ClientCert, cfg.ClientKey)
	if err != nil {
		log.Fatalf("Unable to load keypair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(cfg.CACert)
	if err != nil {
		log.Fatalf("Unable to load CA cert: %s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("Unable to add CA cert: %s", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   host,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := proto.NewAPIClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	return c, ctx, nil

}

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

package main

import (
	"fmt"
	"os"

	"github.com/shawnlower/go-ltp/cmd/ltpcli/add"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/create"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/info"
	"github.com/shawnlower/go-ltp/cmd/ltpcli/list"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-ltp",
	Short: "golang cli for LTP",
	Long: `Utility for persisting data to a remote store.
    LTP store persists object data and semantic meta-data and provides
    the ability to arbitrarily link any objects within the store`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// TODO: this can't be the easiest way
		if cmd.Flag("debug").Value.String() == "true" {
			log.SetLevel(log.DebugLevel)
		}
		log.Debug("Log level set to debug")
	},
}

func init() {
	// Setup logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-ltp.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
    rootCmd.PersistentFlags().String("cert", "cert.pem", "X509 Certificate identifying this server")
    rootCmd.PersistentFlags().String("auth", "mutual-tls", "Authentication type (insecure or mutual-tls)")
    rootCmd.PersistentFlags().String("key", "cert.pem", "PEM encoded private key for the certificate")
    rootCmd.PersistentFlags().String("ca-cert", "cacert.pem", "X509 Certificate of issuing certificate authority")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("remote.auth", rootCmd.PersistentFlags().Lookup("auth"))
	viper.BindPFlag("remote.ca-cert", rootCmd.PersistentFlags().Lookup("ca-cert"))
	viper.BindPFlag("remote.cert", rootCmd.PersistentFlags().Lookup("cert"))
	viper.BindPFlag("remote.key", rootCmd.PersistentFlags().Lookup("key"))

	rootCmd.AddCommand(add.NewAddCommand())
	rootCmd.AddCommand(info.NewInfoCommand())
	rootCmd.AddCommand(list.NewListCommand())
	rootCmd.AddCommand(create.NewCreateCommand())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-ltp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-ltp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("err", err)
	}
}

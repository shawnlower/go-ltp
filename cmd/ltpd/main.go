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

	"github.com/shawnlower/go-ltp/cmd/ltpd/run"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ltpserver",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("debug").Value.String() == "true" {
			log.SetLevel(log.DebugLevel)
		}
		log.Debug("Log level set to debug")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := run.RunServer(cmd, args)
		if err != nil {
			log.Fatal(err)
		}
		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Setup logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-ltp.toml)")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
	rootCmd.PersistentFlags().Bool("foreground", false, "Run in foreground (default when debug enabled)")
	rootCmd.Flags().String("listen-addr", "grpc://127.0.0.1:17900", "listen address")
	rootCmd.Flags().String("cert", "cert.pem", "X509 Certificate identifying this server")
	rootCmd.Flags().String("key", "cert.pem", "PEM encoded private key for the certificate")
	rootCmd.Flags().String("ca-cert", "cacert.pem", "X509 Certificate of issuing certificate authority")
	rootCmd.Flags().String("auth-method", "mutual-tls", "Authentication method to use (mutual-tls, insecure). Default: mutual-tls")

	viper.BindPFlag("server.listen-addr", rootCmd.Flags().Lookup("listen-addr"))
	viper.BindPFlag("server.cert", rootCmd.Flags().Lookup("cert"))
	viper.BindPFlag("server.key", rootCmd.Flags().Lookup("key"))
	viper.BindPFlag("server.ca-cert", rootCmd.Flags().Lookup("ca-cert"))
	viper.BindPFlag("remote.auth", rootCmd.Flags().Lookup("auth-method"))

	fmt.Print("Using cert: ", viper.GetString("server.cert"))

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

		viper.AddConfigPath(home)
		viper.SetConfigName(".go-ltp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config found, using defaults")
	}
}

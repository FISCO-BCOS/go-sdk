/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package console

import (
  "fmt"
  "os"

  "github.com/KasperLiu/gobcos/client"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"

)


var cfgFile string
// RPC is the client connected to the blockchain
var RPC *client.Client
// GroupID default
var GroupID uint
// URL default
var URL string
// PrivateKey default
var PrivateKey = "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58"

// GetClient is used for test, it will be init by a config file later.
func getClient(url string, groupID uint) (*client.Client) {
	// RPC API
	c, err := client.Dial(url, groupID)  // change to your RPC and groupID
	if err != nil {
    fmt.Println("can not dial to the RPC API, please check the config file gobcos_config.yaml: ", err)
    os.Exit(1)
	}
	return c
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "gobcos",
  Short: "gobcos is a RPC console for FISCO BCOS 2.0.0",
  Long: `gobcos is a Golang client for FISCO BCOS 2.0.0 and it supports the JSON-RPC 
service and the contract operations(e.g. deploying && writting contracts).

Also, gobcos can be used as a Go package for FISCO BCOS that just simply adding 
the import statement:

    import "github.com/KasperLiu/gobcos" 
or 
    import "github.com/KasperLiu/gobcos/client" # using the client package

Please access the github site for more details:
    https://github.com/KasperLiu/gobcos.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is the project directory ./gobcos_config.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    // home, err := homedir.Dir()
    // if err != nil {
    //   fmt.Println(err)
    //   os.Exit(1)
    // }

    // Search config in current directory with name "gobcos_config" (without extension).
    viper.AddConfigPath(".")
    viper.SetConfigName("gobcos_config")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    if viper.IsSet("GroupID") {
      GroupID = uint(viper.GetInt("GroupID"))
    } else {
      fmt.Println("GroupID has not been set, please check the config file gobcos_config.yaml")
      os.Exit(1)
    }
    if viper.IsSet("RPCurl") {
      URL = viper.GetString("RPCurl")
    } else {
      fmt.Println("RPCurl has not been set, please check the config file gobcos_config.yaml")
      os.Exit(1)
    }
    RPC = getClient(URL, GroupID)
  }
}
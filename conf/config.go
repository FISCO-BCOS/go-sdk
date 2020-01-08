// Package conf parse config to configuration
package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config contains configuration items for sdk
type Config struct {
	IsHTTP     bool
	ChainID    int64
	CAFile     string
	Cert       string
	Key        string
	IsSMCrypto bool
	PrivateKey string
	GroupID    int
	NodeURL    string
}

// ParseConfig parses the configuration from toml config file
func ParseConfig(cfgFile string) []Config {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetConfigType("toml")
	config := new(Config)
	var configs []Config
	viper.SetDefault("SMCrypto", false)
	viper.SetDefault("Network.Type", "rpc")
	viper.SetDefault("Network.CAFile", "ca.crt")
	viper.SetDefault("Network.Key", "sdk.key")
	viper.SetDefault("Network.Cert", "sdk.crt")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.IsSet("Chain") {
			if viper.IsSet("Chain.ChainID") {
				config.ChainID = int64(viper.GetInt("Chain.ChainID"))
			} else {
				fmt.Println("Chain.ChainID has not been set")
				os.Exit(1)
			}
			if viper.IsSet("Chain.SMCrypto") {
				config.IsSMCrypto = viper.GetBool("Chain.SMCrypto")
			} else {
				fmt.Println("SMCrypto has not been set")
				os.Exit(1)
			}
		} else {
			fmt.Println("Chain has not been set")
			os.Exit(1)
		}
		if viper.IsSet("Account") {
			// accountKeyFile := viper.GetString("Account.KeyFile")
			// FIXME: parse private key from pem file
			config.PrivateKey = "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58"
		} else {
			fmt.Println("Network has not been set")
			os.Exit(1)
		}
		if viper.IsSet("Network") {
			connectionType := viper.GetString("Network.Type")
			if strings.EqualFold(connectionType, "rpc") {
				config.IsHTTP = true
			} else if strings.EqualFold(connectionType, "channel") {
				config.IsHTTP = false
			} else {
				fmt.Printf("Network.Type %s is not supported, use channel", connectionType)
			}
			config.CAFile = viper.GetString("Network.CAFile")
			config.Key = viper.GetString("Network.Key")
			config.Cert = viper.GetString("Network.Cert")
			var connections []struct {
				GroupID int
				NodeURL string
			}
			if viper.IsSet("Network.Connection") {
				err := viper.UnmarshalKey("Network.Connection", &connections)
				if err != nil {
					fmt.Printf("Parse Network.Connection failed. err:%v", err)
					os.Exit(1)
				}
				for i := range connections {
					configs = append(configs, *config)
					configs[i].GroupID = connections[i].GroupID
					configs[i].NodeURL = connections[i].NodeURL
				}
			} else {
				fmt.Printf("Network.Connection has not been set.")
				os.Exit(1)
			}
		} else {
			fmt.Println("Network has not been set")
			os.Exit(1)
		}
	} else {
		fmt.Printf("err message is : %v", err)
	}

	// fmt.Printf("configuration is %+v\n", configs)
	return configs
}

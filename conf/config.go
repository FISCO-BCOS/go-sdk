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
	IsHTTP            bool
	GroupID           uint
	ChainID           int64
	CAFile            string
	Key               string
	Cert              string
	IsSMCrypto        bool
	AccountPrivateKey string
	URLs              []string
}

// ParseConfig parses the configuration from toml config file
func ParseConfig(cfgFile string) *Config {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetConfigType("toml")
	config := new(Config)
	viper.SetDefault("SMCrypto", false)
	viper.SetDefault("Connection.Type", "rpc")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.IsSet("Chain.GroupID") {
			config.GroupID = uint(viper.GetInt("Chain.GroupID"))
		} else {
			fmt.Println("Chain.GroupID has not been set, please check the GroupID in " + viper.ConfigFileUsed())
			os.Exit(1)
		}
		if viper.IsSet("Chain.ChainID") {
			config.ChainID = int64(viper.GetInt("Chain.ChainID"))
		} else {
			fmt.Println("Chain.ChainID has not been set, please check the ChainID in " + viper.ConfigFileUsed())
			os.Exit(1)
		}
		if viper.IsSet("Chain.SMCrypto") {
			config.IsSMCrypto = viper.GetBool("Chain.SMCrypto")
		} else {
			fmt.Println("SMCrypto has not been set, please check the SMCrypto in " + viper.ConfigFileUsed())
			os.Exit(1)
		}
		if viper.IsSet("Connection") {
			connectionType := viper.GetString("Connection.Type")
			if strings.EqualFold(connectionType, "rpc") {
				config.IsHTTP = true
			} else if strings.EqualFold(connectionType, "channel") {
				config.IsHTTP = false
			} else {
				fmt.Printf("Connection.Type %s is not supported, use channel", connectionType)
			}
			nodesURLs := viper.GetStringSlice("Connection.Nodes")
			if len(nodesURLs) == 0 {
				fmt.Println("Connection.Nodes is empty, please check the Connection in " + viper.ConfigFileUsed())
				os.Exit(1)
			} else {
				config.URLs = nodesURLs
			}
		} else {
			fmt.Println("Connection has not been set, please check the Connection in " + viper.ConfigFileUsed())
			os.Exit(1)
		}
		if viper.IsSet("Account") {
			// accountKeyFile := viper.GetString("Account.KeyFile")
			// FIXME: parse private key from pem file
			config.AccountPrivateKey = "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58"
		} else {
			fmt.Println("Connection has not been set, please check the Connection in " + viper.ConfigFileUsed())
			os.Exit(1)
		}
	} else {
		fmt.Printf("err message is : %v", err)
	}
	return config
}

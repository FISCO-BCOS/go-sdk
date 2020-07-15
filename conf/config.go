// Package conf parse config to configuration
package conf

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config contains configuration items for sdk
type Config struct {
	IsHTTP     bool
	ChainID    int64
	CAFile     string
	Key        string
	Cert       string
	IsSMCrypto bool
	PrivateKey string
	GroupID    int
	NodeURL    string
}

// ParseConfigFile parses the configuration from toml config file
func ParseConfigFile(cfgFile string) ([]Config, error) {
	file, err := os.Open(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("open file failed, err: %v", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("close file failed, err: %v", err)
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("file is not found, err: %v", err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("read file failed, err: %v", err)
	}
	return ParseConfig(buffer)
}

// ParseConfig parses the configuration from []byte
func ParseConfig(buffer []byte) ([]Config, error) {
	viper.SetConfigType("toml")
	err := viper.ReadConfig(bytes.NewBuffer(buffer))
	if err != nil {
		return nil, fmt.Errorf("viper .ReadConfig failed, err: %v", err)
	}
	config := new(Config)
	var configs []Config
	viper.SetDefault("SMCrypto", false)
	viper.SetDefault("Network.Type", "rpc")
	viper.SetDefault("Network.CAFile", "ca.crt")
	viper.SetDefault("Network.Key", "sdk.key")
	viper.SetDefault("Network.Cert", "sdk.crt")

	if viper.IsSet("Chain") {
		if viper.IsSet("Chain.ChainID") {
			config.ChainID = int64(viper.GetInt("Chain.ChainID"))
		} else {
			return nil, fmt.Errorf("Chain.ChainID has not been set")
		}
		if viper.IsSet("Chain.SMCrypto") {
			config.IsSMCrypto = viper.GetBool("Chain.SMCrypto")
		} else {
			return nil, fmt.Errorf("SMCrypto has not been set")
		}
	} else {
		return nil, fmt.Errorf("chain has not been set")
	}
	if viper.IsSet("Account") {
		accountKeyFile := viper.GetString("Account.KeyFile")
		keyHex, curve, _, err := LoadECPrivateKeyFromPEM(accountKeyFile)
		if err != nil {
			return nil, fmt.Errorf("parse private key failed, err: %v", err)
		}
		if config.IsSMCrypto && curve != sm2p256v1 {
			return nil, fmt.Errorf("smcrypto must use sm2p256v1 private key, but found %s", curve)
		}
		if !config.IsSMCrypto && curve != secp256k1 {
			return nil, fmt.Errorf("must use secp256k1 private key, but found %s", curve)
		}
		config.PrivateKey = keyHex
	} else {
		return nil, fmt.Errorf("network has not been set")
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
				return nil, fmt.Errorf("parse Network.Connection failed. err: %v", err)
			}
			for i := range connections {
				configs = append(configs, *config)
				configs[i].GroupID = connections[i].GroupID
				configs[i].NodeURL = connections[i].NodeURL
			}
		} else {
			return nil, fmt.Errorf("Network.Connection has not been set")
		}
	} else {
		return nil, fmt.Errorf("network has not been set")
	}
	return configs, nil
}

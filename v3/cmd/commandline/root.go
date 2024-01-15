package commandline

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/smcrypto"
	"github.com/spf13/cobra"
)

var cfgFile string
var privateKeyFilePath string
var smCrypto bool
var disableSsl bool
var nodeEndpoint string
var groupID string
var certPath string

// RPC is the client connected to the blockchain
var RPC *client.Client

// GetClient is used for test, it will be init by a config file later.
func getClient(config *client.Config) (*client.Client, error) {
	// RPC API
	c, err := client.DialContext(context.Background(), config) // change to your RPC and groupID
	if err != nil {
		fmt.Println("can not dial to FISCO node, please check config. error message: ", err)
		return nil, err
	}
	return c, nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "console",
	Short:   "console is a command line tool for FISCO BCOS 3.0.0",
	Version: "3.0.0",
	Long: `console is a Golang client for FISCO BCOS 3.0.0 and it supports the JSON-RPC
service and the contract operations(e.g. deploying && writing contracts).

Also, console can be used as a Go package for FISCO BCOS that just simply adding
the import statement:

    import "github.com/FISCO-BCOS/go-sdk"
or
    import "github.com/FISCO-BCOS/go-sdk/v3/client" # using the client package

Please access the github site for more details:
	https://github.com/FISCO-BCOS/go-sdk.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("console flags:", cmd.Flags())
		err := initConfig()
		if err != nil {
			fmt.Println("init config failed, err: ", err)
			os.Exit(1)
		}
	},
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

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	fmt.Printf("private key file: %s, groupID: %s, disableSsl: %v, isSmCrypto: %v, endpoint: %s, certPath: %s\n", privateKeyFilePath, groupID, disableSsl, smCrypto, nodeEndpoint, certPath)
	var privateKey []byte
	if len(privateKeyFilePath) != 0 {
		_, err := os.Stat(privateKeyFilePath)
		if err != nil && os.IsNotExist(err) {
			fmt.Println("private key file set but not exist, use default private key")
		} else if err != nil {
			return fmt.Errorf("check private key file failed, err: %v\n", err)
		} else {
			key, curve, err := client.LoadECPrivateKeyFromPEM(privateKeyFilePath)
			if err != nil {
				return fmt.Errorf("parse private key failed, err: %v\n", err)
			}
			if smCrypto && curve != client.Sm2p256v1 {
				return fmt.Errorf("smCrypto should use sm2p256v1 private key, but found %s\n", curve)
			}
			if !smCrypto && curve != client.Secp256k1 {
				return fmt.Errorf("should use secp256k1 private key, but found %s\n", curve)
			}
			privateKey = key
		}
	}
	if len(privateKey) == 0 {
		address := "0xFbb18d54e9Ee57529cda8c7c52242EFE879f064F"
		privateKey, _ = hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
		if smCrypto {
			address = smcrypto.SM2KeyToAddress(privateKey).Hex()
		}
		fmt.Println("use default private key, address: ", address)
	}
	ret := strings.Split(nodeEndpoint, ":")
	host := ret[0]
	port, _ := strconv.Atoi(ret[1])
	var config *client.Config
	if !smCrypto {
		config = &client.Config{IsSMCrypto: smCrypto, GroupID: groupID, DisableSsl: disableSsl,
			PrivateKey: privateKey, Host: host, Port: port, TLSCaFile: certPath + "/ca.crt", TLSKeyFile: certPath + "/sdk.key", TLSCertFile: certPath + "/sdk.crt"}
	} else {
		config = &client.Config{IsSMCrypto: smCrypto, GroupID: groupID, DisableSsl: disableSsl,
			PrivateKey: privateKey, Host: host, Port: port, TLSCaFile: certPath + "/sm_ca.crt", TLSKeyFile: certPath + "/sm_sdk.key", TLSCertFile: certPath + "/sm_sdk.crt", TLSSmEnKeyFile: certPath + "/sm_ensdk.key", TLSSmEnCertFile: certPath + "/sm_ensdk.crt"}
	}
	var err error
	RPC, err = getClient(config)
	if err != nil {
		return err
	}
	return nil
}

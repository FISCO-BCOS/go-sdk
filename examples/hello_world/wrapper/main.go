package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
)

func main() {
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	// disable ssl of node rpc
	config := &client.Config{IsSMCrypto: false, GroupID: "group0", DisableSsl: false,
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./conf/ca.crt", TLSKeyFile: "./conf/sdk.key", TLSCertFile: "./conf/sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	input := "HelloWorld deployment 1.0"
	fmt.Println("=================DeployHelloWorld===============")
	address, receipt, instance, err := DeployHelloWorld(client.GetTransactOpts(), client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	fmt.Println("transaction hash: ", receipt.TransactionHash)

	// load the contract
	// contractAddress := common.HexToAddress("contract address in hex String")
	// instance, err := NewStore(contractAddress, client)
	// if err != nil {
	//     log.Fatal(err)
	// }

	fmt.Println("================================")
	helloSession := &HelloWorldSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	version, err := helloSession.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("version :", version) // "HelloWorld deployment 1.0"

	ret, err := helloSession.Get()
	if err != nil {
		fmt.Printf("Get() failed: %v", err)
		return
	}
	done := make(chan bool)
	_, err = helloSession.WatchAllSetValue(nil, func(ret int, logs []types.Log) {
		fmt.Printf("receive event SetValue: %d, logs: %v\n", ret, logs)
		setValue, err := helloSession.ParseSetValue(logs[0])
		if err != nil {
			fmt.Printf("WatchAllSetValue() failed: %v", err)
			panic("WatchAllSetValue failed")
		}
		fmt.Printf("receive setValue: %+v\n", *setValue)
		done <- true
	})
	if err != nil {
		fmt.Printf("WatchAllSetValue() failed: %v", err)
		return
	}
	fmt.Printf("Get: %s\n", ret)
	fmt.Println("================================")

	oldValue, _, receipt, err := helloSession.Set("hello world")
	fmt.Println("old value is: ", oldValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	ret, err = helloSession.Get()
	if err != nil {
		fmt.Printf("Get() failed: %v", err)
		return
	}
	fmt.Printf("Get: %s\n", ret)
	<-done
}

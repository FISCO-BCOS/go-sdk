package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	kvtable "github.com/FISCO-BCOS/go-sdk/examples" // import kvtabletest
)

var (
	channel         = make(chan int, 0)
	contractAddress common.Address
)

func deployContractHandler(receipt *types.Receipt, err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("contract address: ", receipt.ContractAddress) // the address should be saved
	fmt.Printf("tx sent: %s\n", receipt.TransactionHash)
	contractAddress = common.HexToAddress(receipt.ContractAddress)
	channel <- 0
}

func invokeSetHandler(receipt *types.Receipt, err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	//setedLines, err := parseOutput(kvtable.KVTableTestABI, "set", receipt)
	//if err != nil {
	//	logrus.Fatalf("error when transfer string to int: %v\n", err)
	//}
	//fmt.Printf("seted lines: %v\n", setedLines.Int64())
	channel <- 0
}

func main() {
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &conf.Config{IsHTTP: false, ChainID: "chain0",
		IsSMCrypto: false, GroupID: "group0", PrivateKey: privateKey, NodeURL: "127.0.0.1:20200"}

	// deploy Asynccontract
	fmt.Println("-------------------starting deploy contract-----------------------")
	client, err := client.Dial(config)
	if err != nil {
		logrus.Fatal(err)
	}
	_, err = kvtable.AsyncDeployKVTableTest(client.GetTransactOpts(), deployContractHandler, client)
	if err != nil {
		logrus.Fatal(err)
	}
	<-channel

	// invoke AsyncSet to insert info
	fmt.Println("\n-------------------starting invoke Set to insert info-----------------------")
	instance, err := kvtable.NewKVTableTest(contractAddress, client)
	if err != nil {
		logrus.Fatal(err)
	}
	kvtabletestSession := &kvtable.KVTableTestSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	id := "100010001001"
	item_name := "Laptop"
	item_age := "29"
	_, err = kvtabletestSession.AsyncInsert(invokeSetHandler, id, item_name, item_age) // call set API
	if err != nil {
		logrus.Fatal(err)
	}
	<-channel

	// invoke Get to query info
	fmt.Println("\n-------------------starting invoke Get to query info-----------------------")
	item_name, item_age, err = kvtabletestSession.Select(id) // call get API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("id: %v, item_age: %v, item_name: %v \n", id, item_name, item_age)
}

func parseOutput(abiStr, name string, receipt *types.Receipt) (*big.Int, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Printf("parse ABI failed, err: %v", err)
	}
	var ret *big.Int
	b, err := hex.DecodeString(receipt.Output[2:])
	if err != nil {
		return nil, fmt.Errorf("decode receipt.Output[2:] failed, err: %v", err)
	}
	err = parsed.Unpack(&ret, name, b)
	if err != nil {
		return nil, fmt.Errorf("unpack %v failed, err: %v", name, err)
	}
	return ret, nil
}

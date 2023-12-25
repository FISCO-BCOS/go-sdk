package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	kvtable "github.com/FISCO-BCOS/go-sdk/v3/examples" // import kvtabletest
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/sirupsen/logrus"
)

func main() {
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &client.Config{IsSMCrypto: false, GroupID: "group0", PrivateKey: privateKey,
		Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		logrus.Fatal(err)
	}

	// deploy contract
	fmt.Println("-------------------starting deploy contract-----------------------")
	address, tx, instance, err := kvtable.DeployKVTableTest(client.GetTransactOpts(), client)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.TransactionHash)
	_ = instance

	// invoke Set to insert info
	fmt.Println("\n-------------------starting invoke Set to insert info-----------------------")
	kvtabletestSession := &kvtable.KVTableTestSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	id := "100010001001"
	item_name := "Laptop"
	item_age := "29"
	_, receipt, err := kvtabletestSession.Insert(id, item_name, item_age) // call set API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", receipt.TransactionHash)
	//setedLines, err := parseOutput(kvtable.KVTableTestABI, "insert", receipt)
	//if err != nil {
	//	logrus.Fatalf("error when transfer string to int: %v\n", err)
	//}
	//fmt.Printf("seted lines: %v\n", setedLines.String())

	// invoke Get to query info
	fmt.Println("\n-------------------starting invoke Get to query info-----------------------")
	item_name, item_age, err = kvtabletestSession.Select(id) // call get API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("id: %v, item_name: %v, item_age: %v \n", id, item_name, item_age)
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

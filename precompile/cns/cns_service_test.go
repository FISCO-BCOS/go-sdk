package cns

import (
	"testing"
	"context"
	"crypto/ecdsa"

	"github.com/KasperLiu/gobcos/client"
	"github.com/KasperLiu/gobcos/crypto"
	"github.com/KasperLiu/gobcos/accounts/abi/bind"
)

func GetClient(t *testing.T) *client.Client {
	groupID := uint(1)
	rpc, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		t.Fatalf("init rpc client failed: %+v", err)
	}
	return rpc
}

func GenerateKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
    if err != nil {
        t.Fatalf("init privateKey failed: %+v", err)
	}
	return privateKey
}

func GetService(t *testing.T) *CnsService {
	rpc := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewCnsService(rpc, privateKey)
	if err != nil {
		t.Fatalf("init CnsService failed: %+v", err)
	}
	return service
}

func TestAll(t *testing.T) {
	name := "store"
	version := "5.0"
	address := "0x0626918C51A1F36c7ad4354BB1197460A533a2B9"
	abi := `[
		{
			"constant": true,
			"inputs": [
				{
					"name": "",
					"type": "bytes32"
				}
			],
			"name": "items",
			"outputs": [
				{
					"name": "",
					"type": "bytes32"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "version",
			"outputs": [
				{
					"name": "",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"name": "key",
					"type": "bytes32"
				},
				{
					"name": "value",
					"type": "bytes32"
				}
			],
			"name": "setItem",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"name": "_version",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "constructor"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"name": "key",
					"type": "bytes32"
				},
				{
					"indexed": false,
					"name": "value",
					"type": "bytes32"
				}
			],
			"name": "ItemSet",
			"type": "event"
		}
	]`
	rpc := GetClient(t)
	service := GetService(t)

	// test RegisterCns
	tx, err := service.RegisterCns(name, version, address, abi)
	if err != nil {
		t.Fatalf("CnsService RegisterCns failed: %+v\n", err)
	}
	// wait for the mining
    receipt, err := bind.WaitMined(context.Background(), rpc, tx)
    if err != nil {
        t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s\n", receipt.GetTransactionHash())
	
	// test GetAddressByContractNameAndVersion
	addr, err := service.GetAddressByContractNameAndVersion(name + ":" +version)
	if err != nil {
		t.Fatalf("GetAddressByContractNameAndVersion failed: %v", err)
	}
	t.Logf("address: %s", addr)

	// test QueryCnsByNameAndVersion
	cnsInfo, err := service.QueryCnsByNameAndVersion(name, version)
	if err != nil {
		t.Fatalf("QueryCnsByNameAndVersion failed: %v\n", err)
	}
	t.Logf("QueryCnsByNameAndVersion: %s", cnsInfo[0].String())

    // test QueryCnsByNameAndVersion
	cnsInfoByName, err := service.QueryCnsByName(name)
	if err != nil {
		t.Fatalf("QueryCnsByName failed: %v\n", err)
	}
	t.Logf("QueryCnsByName: %s", cnsInfoByName[0].String())
}
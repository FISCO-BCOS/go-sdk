package cns

import (
	"context"
	"crypto/ecdsa"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetClient(t *testing.T) *client.Client {
	// config := &conf.ParseConfig("config.toml")[0]
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58",
		NodeURL:    "http://localhost:8545"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("Dial to %s failed of %v", config.NodeURL, err)
	}
	return c
}

func GenerateKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	if err != nil {
		t.Fatalf("init privateKey failed: %+v", err)
	}
	return privateKey
}

func GetService(t *testing.T) *Service {
	c := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewCnsService(c, privateKey)
	if err != nil {
		t.Fatalf("init Service failed: %+v", err)
	}
	return service
}

const (
	name = "store"
	version = "5.0"
	address = "0x0626918C51A1F36c7ad4354BB1197460A533a2B9"
	testABI = `[
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
)

func TestRegisterCns(t *testing.T)  {
	c := GetClient(t)
	service := GetService(t)
	// test RegisterCns
	tx, err := service.RegisterCns(name, version, address, testABI)
	if err != nil {
		t.Fatalf("Service RegisterCns failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), c, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s\n", receipt.GetTransactionHash())
}

func TestGetAddressByContractNameAndVersion(t *testing.T)  {
	service := GetService(t)

	// test GetAddressByContractNameAndVersion
	addr, err := service.GetAddressByContractNameAndVersion(name + ":" + version)
	if err != nil {
		t.Fatalf("GetAddressByContractNameAndVersion failed: %v", err)
	}
	t.Logf("address: %s", addr)
}

func TestQueryCnsByNameAndVersion(t *testing.T)  {
	service := GetService(t)

	// test QueryCnsByNameAndVersion
	cnsInfo, err := service.QueryCnsByNameAndVersion(name, version)
	if err != nil {
		t.Fatalf("QueryCnsByNameAndVersion failed: %v\n", err)
	}
	t.Logf("QueryCnsByNameAndVersion: %s", cnsInfo[0].String())
}

func TestQueryCnsByName(t *testing.T) {
	service := GetService(t)

	// test QueryCnsByNameAndVersion
	cnsInfoByName, err := service.QueryCnsByName(name)
	if err != nil {
		t.Fatalf("QueryCnsByName failed: %v\n", err)
	}
	t.Logf("QueryCnsByName: %s", cnsInfoByName[0].String())
}

package contractlifecycle

import (
	"encoding/hex"
	"testing"

	helloworld "github.com/FISCO-BCOS/go-sdk/.ci/hello"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

const (
	managerContractAccount = "0xae66fbe9ee2b5007e245d98bf7cf9904cc61e394"
	standardOutput         = 1
)

var (
	contractAddress = ""
)

func GetClient(t *testing.T) *client.Client {
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1,
		PrivateKey: "8c47f550380591adab955cf050c439c0ffabb236bf05a64849ee0ba8aed42a41", NodeURL: "http://localhost:8545"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("Dial to %s failed of %v", config.NodeURL, err)
	}
	return c
}

func GetService(t *testing.T) *Service {
	c := GetClient(t)
	service, err := NewService(c)
	if err != nil {
		t.Fatalf("init contractLifeCycleService failed: %+v", err)
	}
	return service
}

func deployHelloWorldContract(t *testing.T) {
	c := GetClient(t)
	address, tx, instance, err := helloworld.DeployHelloWorld(c.GetTransactOpts(), c) // deploy contract
	if err != nil {
		t.Fatalf("deploy HelloWorld contract failed：%v", err)
	}
	t.Logf("the address of contract: %v", address.Hex())
	t.Logf("the hash of transaction: %v", tx.Hash().Hex())
	_ = instance
	contractAddress = address.Hex()
}

func TestFreeze(t *testing.T) {
	deployHelloWorldContract(t)

	service := GetService(t)
	result, err := service.Freeze(common.HexToAddress(contractAddress))
	if err != nil {
		t.Fatalf("TestFreeze failed: %v", err)
	}
	if result != standardOutput {
		t.Fatalf("TestFreeze failed, the result %v is inconsistent with \"1\"", result)
	}
	t.Logf("TestFreeze: %v", result)
}

func TestUnfreeze(t *testing.T) {
	service := GetService(t)
	result, err := service.Unfreeze(common.HexToAddress(contractAddress))
	if err != nil {
		t.Fatalf("TestUnfreeze failed: %v", err)
	}
	if result != standardOutput {
		t.Fatalf("TestUnfreeze failed, the result %v is inconsistent with \"1\"", result)
	}
	t.Logf("TestUnfreeze: %v", result)
}

func TestGrantManager(t *testing.T) {
	service := GetService(t)
	result, err := service.GrantManager(common.HexToAddress(contractAddress), common.HexToAddress(managerContractAccount))
	if err != nil {
		t.Fatalf("TestGrantManager failed: %v", err)
	}
	if result != standardOutput {
		t.Fatalf("TestGrantManager failed, the result %v is inconsistent with \"1\"", result)
	}
	t.Logf("TestGrantManager: %v", result)
}

func TestGetStatus(t *testing.T) {
	service := GetService(t)
	num, message, err := service.GetStatus(common.HexToAddress(contractAddress))
	if err != nil {
		t.Fatalf("TestGetStatus failed: %v", err)
	}
	if num != 0 && message != "The contract is available." {
		t.Fatalf("TestGetStatus failed, the result \"%v\" and \"%v\" is inconsistent with \"0\" and \"The contract is available.\"", num, message)
	}
	t.Logf("TestGetStatus, num: %v, message: %v", num, message)
}

func TestListManager(t *testing.T) {
	service := GetService(t)
	num, managerAddressList, err := service.ListManager(common.HexToAddress(contractAddress))
	if err != nil {
		t.Fatalf("TestListManager failed: %v", err)
	}
	if num != 0 {
		t.Fatalf("TestListManager failed, the result %v is inconsistent with \"0\"", num)
	}
	for i := 0; i < len(managerAddressList); i++ {
		t.Logf("TestListManager, num: %v, managerAddressList: %v", num, hex.EncodeToString(managerAddressList[i][:]))
	}
}

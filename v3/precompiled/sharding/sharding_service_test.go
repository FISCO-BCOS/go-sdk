package sharding

import (
	"context"
	"encoding/hex"
	"os"
	"testing"

	// "fmt"
	"math/big"
	"time"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/precompiled"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	// "github.com/ethereum/go-ethereum/common"
)

const (
	standardOutput = 0
	timeout        = 1 * time.Second
	name           = "hello_v11"
	version        = "11.0"
	address        = "0xc92ad282ba7868b032341a3921b3635b0c45de74"
	addressAsync   = "0x272d69cfdb321d147f63fdfa9126c4ad1969265a"
	shardName      = "shardName"
	shardNameAsync = "shardName_async"
)

func getClient(t *testing.T) *client.Client {
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &client.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
	c, err := client.DialContext(context.Background(), config)
	if err != nil {
		t.Fatalf("Dial to %s:%d failed of %v", config.Host, config.Port, err)
	}
	return c
}

func getService(t *testing.T) {
	c := getClient(t)
	newService, err := NewShardingService(c)
	if err != nil {
		t.Fatalf("init CnsService failed: %+v", err)
	}
	service = newService
}

var (
	service *Service
	channel = make(chan int)
)

func TestMain(m *testing.M) {
	getService(&testing.T{})
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestMakeShard(t *testing.T) {
	ret0, _, _, err := service.MakeShard(shardName)
	if err != nil {
		t.Fatalf("Service MakeShard failed: %+v\n", err)
	}

	if ret0 != standardOutput {
		t.Fatalf("TestMakeShard failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
	}

	t.Logf("TestMakeShard ret0:%v", ret0)
}

func TestAsyncMakeShard(t *testing.T) {
	handler := func(receipt *types.Receipt, err error) {
		if err != nil {
			t.Fatalf("receive receipt failed, %v\n", err)
		}
		var bigNum *big.Int
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
		if err != nil {
			t.Fatalf("parseReturnValue failed, err: %v\n", err)
		}
		result, err := precompiled.BigIntToInt64(bigNum)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if result != 0 {
			t.Fatalf("TestAsyncMakeShard failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncMakeShard(handler, shardNameAsync)
	if err != nil {
		t.Fatalf("remove data failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestLinkShard(t *testing.T) {
	ret0, _, _, err := service.LinkShard(shardName, address)
	if err != nil {
		t.Fatalf("Service LinkShard failed: %+v\n", err)
	}

	if ret0 != standardOutput {
		t.Fatalf("TestLinkShard failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
	}

	t.Logf("TestMakeShard ret0:%v", ret0)
}

func TestAsyncLinkShard(t *testing.T) {
	handler := func(receipt *types.Receipt, err error) {
		if err != nil {
			t.Fatalf("receive receipt failed, %v\n", err)
		}
		var bigNum *big.Int
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
		if err != nil {
			t.Fatalf("parseReturnValue failed, err: %v\n", err)
		}
		result, err := precompiled.BigIntToInt64(bigNum)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if result != 0 {
			t.Fatalf("TestAsyncLinkShard failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncLinkShard(handler, shardNameAsync, addressAsync)
	if err != nil {
		t.Fatalf("remove data failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestGetContractShard(t *testing.T) {
	ret0, ret1, err := service.GetContractShard(address)
	if err != nil {
		t.Fatalf("Service GetContractShard failed: %+v\n", err)
	}

	if ret0 != standardOutput {
		t.Fatalf("TestGetContractShard failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
	}

	t.Logf("TestGetContractShard ret0:%v, ret0:%v", ret0, ret1)
}

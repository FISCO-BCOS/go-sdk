package bfs

import (
	"context"
	"encoding/hex"
	"os"
	"testing"
	"fmt"
	"math/big"
	"reflect"
	"time"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompiled"
	// "github.com/ethereum/go-ethereum/common"
)

const (
	standardOutput = 0
	timeout        = 1 * time.Second
	name           = "hello_v11"
	version        = "11.0"
	address        = "0xc92ad282ba7868b032341a3921b3635b0c45de74"
	testABI        = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"get","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"n","type":"string"}],"name":"set","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	absolutePath_link = "/testlink"
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
	newService, err := NewBfsService(c)
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

func TestList0(t *testing.T) {
	absolutePath := "/sys/bfs"
	testInfo := BfsInfo{"bfs", "link", []string{"000000000000000000000000000000000000100e"} }

	ret0, ret1, err := service.List0(absolutePath)
	if err != nil {
		t.Fatalf("Service List0 failed: %+v\n", err)
	}

	if ret0!=standardOutput || ret1[0].FileName != testInfo.FileName ||  ret1[0].FileType != testInfo.FileType || reflect.DeepEqual(ret1[0].Ext[:len(testInfo.Ext)], testInfo.Ext)==false{
		t.Fatalf("TestList0 failed, the ret1 %v is inconsistent with \"%v\"", ret1, testInfo)
	}

	t.Logf("TestList0 ret0:%v, ret1:%v", ret0, ret1)
}

func TestList(t *testing.T) {
	absolutePath := "/sys"
	offset := big.NewInt(2)
	limit := big.NewInt(1)
	testInfo := BfsInfo{"bfs", "link", []string{"0","0"} }

	ret0, ret1, err := service.List(absolutePath,offset,limit)
	if err != nil {
		t.Fatalf("Service List failed: %+v\n", err)
	}

	if ret1[0].FileName != testInfo.FileName ||  ret1[0].FileType != testInfo.FileType || reflect.DeepEqual(ret1[0].Ext[:len(testInfo.Ext)], testInfo.Ext)==false{
		t.Fatalf("TestList failed, the ret1 %v is inconsistent with \"%v\"", ret1, testInfo)
	}

	t.Logf("TestList ret0:%v, ret1:%v", ret0, ret1)
}

func TestMkdir(t *testing.T) {
	absolutePath := "/apps/testBfsv0"

	ret0, err := service.Mkdir(absolutePath)
	if err != nil {
		t.Fatalf("Service Mkdir failed: %+v\n", err)
	}

	if ret0 != standardOutput {
		t.Fatalf("TestList0 failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
	}

	t.Logf("TestList0 ret0:%v", ret0)
}

func TestAsyncMkdir(t *testing.T) {
	absolutePath := "/apps/testBfsv1"

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
			t.Fatalf("TestAsyncMkdir failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncMkdir(handler, absolutePath)
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

func TestLink(t *testing.T) {
	ret0, err := service.Link(absolutePath_link,address, testABI)
	if err != nil {
		t.Fatalf("Service Link failed: %+v\n", err)
	}

	if ret0 != standardOutput {
		t.Fatalf("TestLink failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
	}

	t.Logf("TestLink ret0:%v", ret0)
}

func TestAsyncLink(t *testing.T) {
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
			t.Fatalf("TestAsyncLink failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncLink(handler,absolutePath_link,address, testABI)
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

func TestLink0(t *testing.T) {
	ret0, err := service.Link0(name, version, address, testABI)
	if err != nil {
		t.Fatalf("Service Link0 failed: %+v\n", err)
	}

	if ret0 != standardOutput {
		t.Fatalf("TestLink0 failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
	}

	t.Logf("TestLink0 ret0:%v", ret0)
}

func TestAsyncLink0(t *testing.T) {
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
			t.Fatalf("TestAsyncLink0 failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncLink0(handler, name, version, address, testABI)
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

func TestReadlink(t *testing.T) {
	ret0, err := service.Readlink("/apps"+absolutePath_link)
	if err != nil {
		t.Fatalf("Service Link0 failed: %+v\n", err)
	}
	fmt.Println("ret0",ret0)

	if strings.ToLower(ret0.Hex()) != address {
		t.Fatalf("TestLink0 failed, the ret0 %v is inconsistent with \"%v\"", ret0, address)
	}

	t.Logf("TestLink0 ret0:%v", ret0)
}

// func TestTouch(t *testing.T) {
// 	absolutePath:="/apps/sbw"
// 	fileType:="txt"
// 	ret0, err := service.Touch(absolutePath, fileType)
// 	if err != nil {
// 		t.Fatalf("Service Link0 failed: %+v\n", err)
// 	}
// 	fmt.Println("ret0",ret0)

// 	if ret0 != standardOutput {
// 		t.Fatalf("TestTouch failed, the ret0 %v is inconsistent with \"%v\"", ret0, standardOutput)
// 	}

// 	t.Logf("TestTouch ret0:%v", ret0)
// }
// Package main go 实现的压测工具
package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/kvTableTest"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/parallelOk"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/server"
	"github.com/sirupsen/logrus"
)

// array 自定义数组参数
type array []string

// String string
func (a *array) String() string {
	return fmt.Sprint(*a)
}

// Set set
func (a *array) Set(s string) error {
	*a = append(*a, s)

	return nil
}

var (
	concurrency    uint64 = 1       // 并发数
	totalNumber    uint64 = 1       // 请求数(单个并发/协程)
	debugStr              = "false" // 是否是debug
	requestURL            = ""      // 压测的url 目前支持，http/https ws/wss
	path                  = ""      // curl文件路径 http接口压测，自定义参数设置
	contractType          = ""      // curl文件路径 http接口压测，自定义参数设置
	contractMothod        = ""      // curl文件路径 http接口压测，自定义参数设置
	cpuNumber             = 1       // CUP 核数，默认为一核，一般场景下单核已经够用了
	timeout        int64  = 0       // 超时时间，默认不设置
)

func init() {
	flag.Uint64Var(&concurrency, "c", concurrency, "并发数")
	flag.Uint64Var(&totalNumber, "n", totalNumber, "请求数(单个并发/协程)")
	flag.StringVar(&debugStr, "d", debugStr, "调试模式")
	flag.StringVar(&contractType, "t", requestURL, "合约类型")
	flag.StringVar(&contractMothod, "m", requestURL, "合约方法")
	flag.IntVar(&cpuNumber, "cpuNumber", cpuNumber, "CUP 核数，默认为一核")
	flag.Int64Var(&timeout, "timeout", timeout, "超时时间 单位 秒,默认不设置")
	flag.Parse()
}

//go:generate go build main.go
func main() {
	runtime.GOMAXPROCS(cpuNumber)
	if concurrency == 0 || totalNumber == 0 || (contractMothod == "" && contractType == "") {
		fmt.Printf("示例: go run main.go -c 1 -n 1 -u https://www.baidu.com/ \n")
		fmt.Printf("压测地址或curl路径必填 \n")
		fmt.Printf("当前请求参数: -c %d -n %d -d %v -u %s \n", concurrency, totalNumber, debugStr, requestURL)
		flag.Usage()
		return
	}
	request, err := model.NewRequestByContractType(contractType, contractMothod)
	if err != nil {
		fmt.Printf("参数不合法 %v \n", err)
		return
	}
	fmt.Printf("\n 开始启动  并发数:%d 请求数:%d 请求参数: \n", concurrency, totalNumber)

	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()
		deadline, ok := ctx.Deadline()
		if ok {
			fmt.Printf(" deadline %s", deadline)
		}
	}

	//
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &conf.Config{ChainID: 1, CAFile: "ca.crt", Key: "sdk.key", Cert: "sdk.crt",
		IsSMCrypto: false, GroupID: "group0", PrivateKey: privateKey, NodeURL: "127.0.0.1:20200"}
	client, err := client.Dial(config)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("-------------------starting deploy contract-----------------------")
	switch request.Form {
	case model.FormTypeKvTable:
		_, _, instance, err := kvTableTest.DeployKVTableTest(client.GetTransactOpts(), client)
		if err != nil {
			logrus.Fatal(err)
		}
		kvtabletestSession := &kvTableTest.KVTableTestSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
		server.Dispose(ctx, concurrency, totalNumber, request, kvtabletestSession)
	case model.FormParallelOk:
		address, tx, instance, err := parallelOk.DeployParallelOk(client.GetTransactOpts(), client)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println("contract address: ", address.Hex()) // the address should be saved
		fmt.Println("transaction hash: ", tx.TransactionHash)
		_ = instance
		parallelOkSession := &parallelOk.ParallelOkSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
		server.Dispose(ctx, concurrency, totalNumber, request, parallelOkSession)
	}

	return
}

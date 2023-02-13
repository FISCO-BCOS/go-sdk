package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/kvTableTest"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/parallelOk"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/server"
	"github.com/sirupsen/logrus"
)

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
	concurrency    uint64 = 1       // concurrency
	totalNumber    uint64 = 1       // Number of Requests
	debugStr              = "false" // debug or not
	requestURL            = ""      //
	path                  = ""
	contractType          = ""
	contractMothod        = ""
	cpuNumber             = 1 // CPU number
	timeout        int64  = 0 // Timeout unit s,Not set by default
)

func init() {
	flag.Uint64Var(&concurrency, "c", concurrency, "concurrency")
	flag.Uint64Var(&totalNumber, "n", totalNumber, "Number of Requests")
	flag.StringVar(&debugStr, "d", debugStr, "debug mode")
	flag.StringVar(&contractType, "t", requestURL, "contract type")
	flag.StringVar(&contractMothod, "m", requestURL, "contract method")
	flag.IntVar(&cpuNumber, "cpuNumber", cpuNumber, "CUP number")
	flag.Int64Var(&timeout, "timeout", timeout, "Timeout unit s,Not set by default")
	flag.Parse()
}

//go:generate go build main.go
func main() {
	runtime.GOMAXPROCS(cpuNumber)
	if concurrency == 0 || totalNumber == 0 || (contractMothod == "" && contractType == "") {
		fmt.Printf("example: go run -ldflags=\"-r ./libs/linux\" main.go -c 1000 -n 100 -t kvTableTest -m select -cpuNumber 8 \n")
		fmt.Printf("Current request parameters: -c %d -n %d -d %v -u %s \n", concurrency, totalNumber, debugStr, requestURL)
		flag.Usage()
		return
	}
	request, err := model.NewRequestByContractType(contractType, contractMothod)
	if err != nil {
		fmt.Printf("parameter invalid %v \n", err)
		return
	}
	fmt.Printf("\n start...  concurrency:%d Number of Requests:%d Requests param: \n", concurrency, totalNumber)

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
	config := &client.Config{IsSMCrypto: false, GroupID: "group0", PrivateKey: privateKey,
		Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
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

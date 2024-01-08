package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/schollz/progressbar/v3"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("Usage: ./%s groupID userCount total qps", os.Args[0])
		return
	}
	groupID := os.Args[1]
	userCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("parse userCount error", err)
		return
	}
	total, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("parse total error", err)
		return
	}
	qps, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("parse qps error", err)
		return
	}
	fmt.Println("start perf groupID:", groupID, "userCount:", userCount, "total:", total, "qps:", qps)

	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &client.Config{IsSMCrypto: false, GroupID: groupID, DisableSsl: false,
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./conf/ca.crt", TLSKeyFile: "./conf/sdk.key", TLSCertFile: "./conf/sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	// client, err := client.Dial("./config.ini", groupID, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=================DeployTransfer===============")
	address, receipt, instance, err := DeployTransfer(client.GetTransactOpts(), client)
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

	transfer := &TransferSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	fmt.Println("add user")
	var wg sync.WaitGroup
	// balance := make(map[int]int64, userCount)
	balance := sync.Map{}
	initValue := int64(1000000000)
	failedCount := 0
	for i := 0; i < userCount; i++ {
		_, err = transfer.AsyncSet(func(receipt *types.Receipt, err error) {
			if err != nil {
				fmt.Println("add user error", err)
			}
			wg.Done()
		}, strconv.Itoa(i), big.NewInt(initValue))
		if err != nil {
			fmt.Println("add user error", err)
			failedCount++
		}
		balance.Store(i, initValue)
		wg.Add(1)
	}
	if failedCount > 0 {
		fmt.Println("add user failed", failedCount)
		return
	}
	wg.Wait()
	fmt.Println("start transfer")
	var wg2 sync.WaitGroup
	sendBar := progressbar.Default(int64(total), "send")
	receiveBar := progressbar.Default(int64(total), "receive")
	// routineCount := (qps + 4000) / 4000
	// sended := int64(0)
	// for i := 0; i < routineCount; i++ {
	// 	fmt.Printf("start routine :%d\n", i)
	// 	go func() {
	// 		for {
	// 			from := rand.Intn(userCount)
	// 			to := rand.Intn(userCount)
	// 			if from == to {
	// 				continue
	// 			}
	// 			if atomic.LoadInt64(&sended) == int64(total) {
	// 				break
	// 			}
	// 			amount := int64(rand.Intn(10))
	// 			_, err = transfer.AsyncTransfer(func(receipt *types.Receipt, err error) {
	// 				receiveBar.Add(1)
	// 				wg2.Done()
	// 				if err != nil {
	// 					fmt.Println("transfer error", err)
	// 					return
	// 				}
	// 				currentFrom, _ := balance.Load(from)
	// 				currentTo, _ := balance.Load(to)
	// 				if !balance.CompareAndSwap(from, currentFrom.(int64), currentFrom.(int64)-amount) {
	// 					for {
	// 						currentFrom, _ := balance.Load(from)
	// 						if balance.CompareAndSwap(from, currentFrom.(int64), currentFrom.(int64)-amount) {
	// 							break
	// 						}
	// 					}
	// 				}
	// 				if !balance.CompareAndSwap(to, currentTo.(int64), currentTo.(int64)+amount) {
	// 					for {
	// 						currentTo, _ := balance.Load(to)
	// 						if balance.CompareAndSwap(to, currentTo.(int64), currentTo.(int64)+amount) {
	// 							break
	// 						}
	// 					}
	// 				}
	// 			}, strconv.Itoa(from), strconv.Itoa(to), big.NewInt(int64(amount)))
	// 			if err != nil {
	// 				fmt.Println("transfer error", err)
	// 				continue
	// 			}
	// 			atomic.AddInt64(&sended, 1)
	// 			sendBar.Add(1)
	// 			wg2.Add(1)
	// 		}
	// 	}()
	// }
	// time.Sleep(time.Second * 5)
	// wg2.Wait()

	for i := 0; i < total; i++ {
		from := i % userCount
		to := (i + userCount/2) % userCount
		amount := int64(1)
		_, err = transfer.AsyncTransfer(func(receipt *types.Receipt, err error) {
			receiveBar.Add(1)
			wg2.Done()
			if err != nil {
				fmt.Println("transfer error", err)
				return
			}
			currentFrom, _ := balance.Load(from)
			currentTo, _ := balance.Load(to)
			if !balance.CompareAndSwap(from, currentFrom.(int64), currentFrom.(int64)-amount) {
				for {
					currentFrom, _ := balance.Load(from)
					if balance.CompareAndSwap(from, currentFrom.(int64), currentFrom.(int64)-amount) {
						break
					}
				}
			}
			if !balance.CompareAndSwap(to, currentTo.(int64), currentTo.(int64)+amount) {
				for {
					currentTo, _ := balance.Load(to)
					if balance.CompareAndSwap(to, currentTo.(int64), currentTo.(int64)+amount) {
						break
					}
				}
			}
		}, strconv.Itoa(from), strconv.Itoa(to), big.NewInt(int64(amount)))
		if err != nil {
			fmt.Println("transfer error", err)
			continue
		}
		sendBar.Add(1)
		wg2.Add(1)
	}
	wg2.Wait()

	// check balance
	fmt.Println("check balance...")
	var wg3 sync.WaitGroup
	for i := 0; i < userCount; i++ {
		wg3.Add(1)
		go func(i int) {
			b, err := transfer.BalanceOf(strconv.Itoa(i))
			if err != nil {
				fmt.Println("check balance error", err)
				return
			}
			current, _ := balance.Load(i)
			if b.Cmp(big.NewInt(int64(current.(int64)))) != 0 {
				fmt.Println("check balance error", i, b, current.(int64))
			}
			wg3.Done()
		}(i)
	}
	wg3.Wait()
	fmt.Println("check balance done")
}

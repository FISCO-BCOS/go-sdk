package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/smcrypto"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/schollz/progressbar/v3"
	flag "github.com/spf13/pflag"
)

func main() {
	pemFileName := flag.StringP("pem", "p", "", "pem file path")
	groupID := flag.StringP("group", "g", "group0", "groupID")
	disableSsl := flag.BoolP("disableSsl", "d", false, "disable ssl")
	isSmCrypto := flag.BoolP("smCrypto", "s", false, "use sm crypto")
	endpoint := flag.StringP("endpoint", "e", "127.0.0.1:20200", "node endpoint")
	certPath := flag.StringP("cert", "c", "./conf/", "cert path")
	userCount := flag.IntP("userCount", "u", 1000, "user count")
	totalTx := flag.IntP("totalTxTx", "t", 10000, "totalTx tx")
	qps := flag.IntP("qps", "q", 1000, "qps")
	flag.Parse()
	fmt.Printf("pem: %s, groupID: %s, disableSsl: %v, isSmCrypto: %v, endpoint: %s, certPath: %s, userCount: %d, totalTx: %d, qps: %d\n", *pemFileName, *groupID, *disableSsl, *isSmCrypto, *endpoint, *certPath, *userCount, *totalTx, *qps)

	var privateKey []byte
	if len(*pemFileName) != 0 {
		_, err := os.Stat(*pemFileName)
		if err != nil && os.IsNotExist(err) {
			fmt.Println("private key file set but not exist, use default private key")
		} else if err != nil {
			fmt.Printf("check private key file failed, err: %v\n", err)
			return
		} else {
			key, curve, err := client.LoadECPrivateKeyFromPEM(*pemFileName)
			if err != nil {
				fmt.Printf("parse private key failed, err: %v\n", err)
				return
			}
			if *isSmCrypto && curve != client.Sm2p256v1 {
				fmt.Printf("smCrypto should use sm2p256v1 private key, but found %s\n", curve)
				return
			}
			if !*isSmCrypto && curve != client.Secp256k1 {
				fmt.Printf("should use secp256k1 private key, but found %s\n", curve)
				return
			}
			privateKey = key
		}
	}
	if len(privateKey) == 0 {
		address := "0xFbb18d54e9Ee57529cda8c7c52242EFE879f064F"
		privateKey, _ = hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
		if *isSmCrypto {
			address = smcrypto.SM2KeyToAddress(privateKey).Hex()
		}
		fmt.Println("use default private key, address: ", address)
	}
	ret := strings.Split(*endpoint, ":")
	host := ret[0]
	port, _ := strconv.Atoi(ret[1])
	var config *client.Config
	if !*isSmCrypto {
		config = &client.Config{IsSMCrypto: *isSmCrypto, GroupID: *groupID, DisableSsl: *disableSsl,
			PrivateKey: privateKey, Host: host, Port: port, TLSCaFile: *certPath + "/ca.crt", TLSKeyFile: *certPath + "/sdk.key", TLSCertFile: *certPath + "/sdk.crt"}
	} else {
		config = &client.Config{IsSMCrypto: *isSmCrypto, GroupID: *groupID, DisableSsl: *disableSsl,
			PrivateKey: privateKey, Host: host, Port: port, TLSCaFile: *certPath + "/sm_ca.crt", TLSKeyFile: *certPath + "/sm_sdk.key", TLSCertFile: *certPath + "/sm_sdk.crt", TLSSmEnKeyFile: *certPath + "/sm_ensdk.key", TLSSmEnCertFile: *certPath + "/sm_ensdk.crt"}
	}
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
	for i := 0; i < *userCount; i++ {
		_, err = transfer.AsyncSet(func(receipt *types.Receipt, err error) {
			if err != nil {
				fmt.Println("add user error", err)
			}
			wg.Done()
		}, strconv.Itoa(i), big.NewInt(initValue))
		if err != nil {
			fmt.Println("add user error", err)
			failedCount++
			continue
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
	sendBar := progressbar.Default(int64(*totalTx), "send")
	receiveBar := progressbar.Default(int64(*totalTx), "receive")
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

	for i := 0; i < *totalTx; i++ {
		from := i % *userCount
		to := (i + *userCount/2) % *userCount
		amount := int64(1)
		_, err = transfer.AsyncTransfer(func(receipt *types.Receipt, err error) {
			receiveBar.Add(1)
			if err != nil {
				fmt.Println("AsyncTransfer error", err)
				return
			}
			wg2.Done()
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
	for i := 0; i < *userCount; i++ {
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

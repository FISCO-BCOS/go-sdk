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
	"time"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/smcrypto"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/schollz/progressbar/v3"
	flag "github.com/spf13/pflag"
	"golang.org/x/time/rate"
)

// TransferABI is the input ABI used to generate the binding from.
const TransferABI = "[{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"selector\":[904814471,3449012829],\"stateMutability\":\"view\",\"type\":\"function\"},{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"selector\":[2319641577,4076138093],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]},{\"kind\":3,\"slot\":0,\"value\":[1]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"selector\":[2608902224,1630350335],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]},{\"kind\":3,\"slot\":0,\"value\":[1]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transferWithRevert\",\"outputs\":[],\"selector\":[4208209799,2876358409],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransferBin is the compiled bytecode used for deploying new contracts.
var TransferBin = "0x608060405234801561001057600080fd5b50610679806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806335ee5f87146100515780638a42ebe9146100815780639b80b0501461009d578063fad42f87146100b9575b600080fd5b61006b60048036038101906100669190610369565b6100d5565b60405161007891906103cb565b60405180910390f35b61009b60048036038101906100969190610412565b6100fc565b005b6100b760048036038101906100b2919061046e565b610123565b005b6100d360048036038101906100ce919061046e565b610192565b005b600080826040516100e69190610573565b9081526020016040518091039020549050919050565b8060008360405161010d9190610573565b9081526020016040518091039020819055505050565b806000846040516101349190610573565b9081526020016040518091039020600082825461015191906105b9565b92505081905550806000836040516101699190610573565b9081526020016040518091039020600082825461018691906105ed565b92505081905550505050565b806000846040516101a39190610573565b908152602001604051809103902060008282546101c091906105b9565b92505081905550806000836040516101d89190610573565b908152602001604051809103902060008282546101f591906105ed565b92505081905550606481111561020a57600080fd5b505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102768261022d565b810181811067ffffffffffffffff821117156102955761029461023e565b5b80604052505050565b60006102a861020f565b90506102b4828261026d565b919050565b600067ffffffffffffffff8211156102d4576102d361023e565b5b6102dd8261022d565b9050602081019050919050565b82818337600083830152505050565b600061030c610307846102b9565b61029e565b90508281526020810184848401111561032857610327610228565b5b6103338482856102ea565b509392505050565b600082601f8301126103505761034f610223565b5b81356103608482602086016102f9565b91505092915050565b60006020828403121561037f5761037e610219565b5b600082013567ffffffffffffffff81111561039d5761039c61021e565b5b6103a98482850161033b565b91505092915050565b6000819050919050565b6103c5816103b2565b82525050565b60006020820190506103e060008301846103bc565b92915050565b6103ef816103b2565b81146103fa57600080fd5b50565b60008135905061040c816103e6565b92915050565b6000806040838503121561042957610428610219565b5b600083013567ffffffffffffffff8111156104475761044661021e565b5b6104538582860161033b565b9250506020610464858286016103fd565b9150509250929050565b60008060006060848603121561048757610486610219565b5b600084013567ffffffffffffffff8111156104a5576104a461021e565b5b6104b18682870161033b565b935050602084013567ffffffffffffffff8111156104d2576104d161021e565b5b6104de8682870161033b565b92505060406104ef868287016103fd565b9150509250925092565b600081519050919050565b600081905092915050565b60005b8381101561052d578082015181840152602081019050610512565b8381111561053c576000848401525b50505050565b600061054d826104f9565b6105578185610504565b935061056781856020860161050f565b80840191505092915050565b600061057f8284610542565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105c4826103b2565b91506105cf836103b2565b9250828210156105e2576105e161058a565b5b828203905092915050565b60006105f8826103b2565b9150610603836103b2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156106385761063761058a565b5b82820190509291505056fea26469706673582212209c3a15397c5fc2d0668cca2ef1bce80e3b8124697ac44a54a13f74f9e30961d164736f6c634300080b0033"
var TransferSMBin = "0x608060405234801561001057600080fd5b50610679806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063612d2bff14610051578063ab71bf091461006d578063cd93c25d14610089578063f2f4ee6d146100b9575b600080fd5b61006b6004803603810190610066919061039f565b6100d5565b005b6100876004803603810190610082919061039f565b610144565b005b6100a3600480360381019061009e919061042a565b6101c1565b6040516100b09190610482565b60405180910390f35b6100d360048036038101906100ce919061049d565b6101e8565b005b806000846040516100e69190610573565b9081526020016040518091039020600082825461010391906105b9565b925050819055508060008360405161011b9190610573565b9081526020016040518091039020600082825461013891906105ed565b92505081905550505050565b806000846040516101559190610573565b9081526020016040518091039020600082825461017291906105b9565b925050819055508060008360405161018a9190610573565b908152602001604051809103902060008282546101a791906105ed565b9250508190555060648111156101bc57600080fd5b505050565b600080826040516101d29190610573565b9081526020016040518091039020549050919050565b806000836040516101f99190610573565b9081526020016040518091039020819055505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7fb95aa35500000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102768261022d565b810181811067ffffffffffffffff821117156102955761029461023e565b5b80604052505050565b60006102a861020f565b90506102b4828261026d565b919050565b600067ffffffffffffffff8211156102d4576102d361023e565b5b6102dd8261022d565b9050602081019050919050565b82818337600083830152505050565b600061030c610307846102b9565b61029e565b90508281526020810184848401111561032857610327610228565b5b6103338482856102ea565b509392505050565b600082601f8301126103505761034f610223565b5b81356103608482602086016102f9565b91505092915050565b6000819050919050565b61037c81610369565b811461038757600080fd5b50565b60008135905061039981610373565b92915050565b6000806000606084860312156103b8576103b7610219565b5b600084013567ffffffffffffffff8111156103d6576103d561021e565b5b6103e28682870161033b565b935050602084013567ffffffffffffffff8111156104035761040261021e565b5b61040f8682870161033b565b92505060406104208682870161038a565b9150509250925092565b6000602082840312156104405761043f610219565b5b600082013567ffffffffffffffff81111561045e5761045d61021e565b5b61046a8482850161033b565b91505092915050565b61047c81610369565b82525050565b60006020820190506104976000830184610473565b92915050565b600080604083850312156104b4576104b3610219565b5b600083013567ffffffffffffffff8111156104d2576104d161021e565b5b6104de8582860161033b565b92505060206104ef8582860161038a565b9150509250929050565b600081519050919050565b600081905092915050565b60005b8381101561052d578082015181840152602081019050610512565b8381111561053c576000848401525b50505050565b600061054d826104f9565b6105578185610504565b935061056781856020860161050f565b80840191505092915050565b600061057f8284610542565b915081905092915050565b7fb95aa35500000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105c482610369565b91506105cf83610369565b9250828210156105e2576105e161058a565b5b828203905092915050565b60006105f882610369565b915061060383610369565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156106385761063761058a565b5b82820190509291505056fea26469706673582212204f8ebd1bd8a95fec67c4a519537e11316b2dd90a6d972b04ef627aac78dc982f64736f6c634300080b0033"

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
	if err != nil {
		log.Fatal(err)
	}
	// deploy parallelok contract
	currentNumber, err := client.GetBlockNumber(context.Background())
	if err != nil {
		log.Fatalf("GetBlockNumber error: %v", err)
	}
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		log.Fatalf("abi.JSON error: %v", err)
	}
	blockLimit := currentNumber + 900
	var input []byte
	// 1. create txData
	if client.SMCrypto() {
		input = common.FromHex(TransferSMBin)
	} else {
		input = common.FromHex(TransferBin)
	}
	tx, txHash, err := encodeAndSignTx(client, nil, input, blockLimit)
	if err != nil {
		log.Fatalf("encodeAndSignTx error: %v", err)
	}
	receipt, err := client.SendEncodedTransaction(context.Background(), tx, true)
	if err != nil {
		log.Fatalf("SendEncodedTransaction error: %v", err)
	}
	address := common.HexToAddress(receipt.ContractAddress)
	fmt.Printf("txHash: %x\n", txHash)
	fmt.Printf("contract address: %s\n", address.Hex())
	// call parallelok set
	fmt.Println("add user")
	var wg sync.WaitGroup
	// balance := make(map[int]int64, userCount)
	balance := sync.Map{}
	initValue := int64(1000000000)
	failedCount := 0
	for i := 0; i < *userCount; i++ {
		input, err = parsed.Pack("set", strconv.Itoa(i), big.NewInt(initValue))
		if err != nil {
			log.Fatalf("parsed.Pack error: %v", err)
		}
		tx, _, err = encodeAndSignTx(client, &address, input, blockLimit)
		if err != nil {
			log.Fatalf("CreateEncodedTransactionDataV1 error: %v", err)
		}
		err = client.AsyncSendEncodedTransaction(context.Background(), tx, false, func(receipt *types.Receipt, err error) {
			if err != nil {
				log.Fatalf("AsyncSendEncodedTransaction error: %v", err)
			}
			if receipt.Status != 0 {
				log.Fatalf("receipt status error: %s", receipt.GetErrorMessage())
			}
			wg.Done()
		})
		if err != nil {
			fmt.Printf("add user error: %v\n", err)
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
	// update blockLimit
	stopChan := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("stop update blockLimit goroutine")
				return
			default:
				currentNumber, err := client.GetBlockNumber(context.Background())
				if err != nil {
					log.Fatalf("GetBlockNumber error: %v", err)
				}
				blockLimit = currentNumber + 900
				time.Sleep(5 * time.Second)
			}
		}
	}()
	// call parallel transfer async

	var wg2 sync.WaitGroup
	sendBar := progressbar.Default(int64(*totalTx), "send")
	receiveBar := progressbar.Default(int64(*totalTx), "receive")
	limiter := rate.NewLimiter(rate.Limit(*totalTx), *qps)
	for i := 0; i < *totalTx; i++ {
		from := i % *userCount
		to := (i + *userCount/2) % *userCount
		amount := int64(1)
		input, err = parsed.Pack("transfer", strconv.Itoa(from), strconv.Itoa(to), big.NewInt(int64(amount)))
		if err != nil {
			log.Fatalf("parsed.Pack error: %v", err)
		}
		// create tx
		tx, _, err = encodeAndSignTx(client, &address, input, blockLimit)
		if err != nil {
			log.Fatalf("encodeAndSignTx error: %v", err)
		}
		err = limiter.Wait(context.Background())
		if err != nil {
			log.Fatalf("limiter Wait error: %v", err)
		}
		err = client.AsyncSendEncodedTransaction(context.Background(), tx, false, func(receipt *types.Receipt, err error) {
			receiveBar.Add(1)
			if err != nil {
				fmt.Println("transfer error", err)
				return
			}
			wg2.Done()
			if receipt.Status != 0 {
				fmt.Println("transfer error", receipt.GetErrorMessage())
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
		})
		if err != nil {
			fmt.Println("transfer error", err)
			continue
		}
		sendBar.Add(1)
		wg2.Add(1)
	}
	wg2.Wait()
	stopChan <- struct{}{}
}

func encodeAndSignTx(client *client.Client, address *common.Address, input []byte, blockLimit int64) ([]byte, []byte, error) {
	txData, txHash, err := client.CreateEncodedTransactionDataV1(address, input, blockLimit, TransferABI)
	if err != nil {
		log.Fatalf("CreateEncodedTransactionDataV1 error: %v", err)
	}
	// 2. sign txData
	signature, err := client.CreateEncodedSignature(txHash)
	if err != nil {
		log.Fatalf("CreateEncodedSignature error: %v", err)
	}
	// 3. create tx, tx include txData, txHash, signature, arrtibute, extraData
	tx, err := client.CreateEncodedTransaction(txData, txHash, signature, 0, "")
	if err != nil {
		log.Fatalf("CreateEncodedTransaction error: %v", err)
	}
	return tx, txHash, nil
	// 4. send tx
	// receipt, err := client.SendEncodedTransaction(context.Background(), tx, true)
	// if err != nil {
	// 	log.Fatalf("SendEncodedTransaction error: %v", err)
	// }
}

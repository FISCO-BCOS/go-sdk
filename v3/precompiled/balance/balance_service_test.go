package balanceprecompiled

import (
	"context"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"math/big"
	"os"
	"testing"
)

func getClient(t *testing.T) *client.Client {
	keyBytes, _, err := client.LoadECPrivateKeyFromPEM("./0x0d5535bc7f3707053e653c8d525983182fa9b4e6.pem")
	if err != nil {
		fmt.Printf("LoadECPrivateKeyFromPEM failed: %+v", err)
	}
	config := &client.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: keyBytes, Host: "127.0.0.1", Port: 20222, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
	c, err := client.DialContext(context.Background(), config)
	if err != nil {
		t.Fatalf("Dial to %s:%d failed of %v", config.Host, config.Port, err)
	}
	return c
}

func getService(t *testing.T) {
	c := getClient(t)
	newService, err := NewBalanceService(c)
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

func TestGetBalance(t *testing.T) {
	address := "0xc92ad282ba7868b032341a3921b3635b0c45de75"
	balance, err := service.GetBalance(address)
	if err != nil {
		t.Fatalf("get balance failed: %+v", err)
	}
	t.Logf("TestGetBalance success %+v", balance)
	if balance.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("getBalance failed, balance is not equal to 0")
	}
	t.Logf("TestGetBalance success %+v", balance)
}

func TestBalance(t *testing.T) {
	account := "0xc92ad282ba7868b032341a3921b3635b0c45de77"
	value := big.NewInt(100)
	_, err := service.addBalance(account, value)
	if err != nil {
		t.Fatalf("add balance failed: %+v", err)
	}
	balance, err := service.GetBalance(account)
	fmt.Println(balance)
	if err != nil {
		t.Fatalf("get balance failed: %+v", err)
	}
	// check balance
	if balance.Cmp(value) != 0 {
		t.Fatalf("add balance failed, balance is not equal to %+v", value)
	}

	// sub balance
	value1 := big.NewInt(50)
	_, err = service.subBalance(account, value1)
	if err != nil {
		t.Fatalf("sub balance failed: %+v", err)
	}

	// check balance
	balance1, err := service.GetBalance(account)
	fmt.Println(balance1)
	if err != nil {
		t.Fatalf("get balance failed: %+v", err)
	}
	if balance1.Cmp(value1) != 0 {
		t.Fatalf("sub balance failed, balance is not equal to %+v", value1)
	}

	t.Logf("TestBalance success")
}

func TestTransferBalance(t *testing.T) {
	from := "0xc92ad282ba7868b032341a3921b3635b0c45de84"
	to := "0x2c7536e3605d9c16a7a3d7b1898e529396a65c45"
	value := big.NewInt(10)
	// addBalance to from
	_, err := service.addBalance(from, big.NewInt(50))
	if err != nil {
		t.Fatalf("add balance failed: %+v", err)
	}
	// check from balance
	balance, err := service.GetBalance(from)
	if err != nil {
		t.Fatalf("get balance failed: %+v", err)
	}
	if balance.Cmp(big.NewInt(50)) != 0 {
		t.Fatalf("add balance failed, from balance is not equal to 50")
	}

	// transfer
	_, err = service.Transfer(from, to, value)
	if err != nil {
		t.Fatalf("transfer balance failed: %+v", err)
	}

	// check from balance
	fromBalance, err := service.GetBalance(from)
	if err != nil {
		t.Fatalf("get balance failed: %+v", err)
	}
	if fromBalance.Cmp(big.NewInt(40)) != 0 {
		t.Fatalf("transfer balance failed, from balance is not equal to 40")
	}

	// check to balance
	toBalance, err := service.GetBalance(to)
	if err != nil {
		t.Fatalf("get balance failed: %+v", err)
	}
	if toBalance.Cmp(big.NewInt(10)) != 0 {
		t.Fatalf("transfer balance failed, to balance is not equal to 10")
	}

	t.Logf("TestTransferBalance success")
}

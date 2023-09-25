package config

import (
	"context"
	"encoding/hex"
	"testing"
	"os"

	"github.com/FISCO-BCOS/go-sdk/client"
)

const (
	standardOutput = 0
	key            = "tx_count_limit"
	value          = "30000000"
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
	newService, err := NewSystemConfigService(c)
	if err != nil {
		t.Fatalf("init CnsService failed: %+v", err)
	}
	service = newService
}

var (
	service *SystemConfigService
)

func TestMain(m *testing.M) {
	getService(&testing.T{})
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSetValueByKey(t *testing.T) {
	result, err := service.SetValueByKey(key, value)
	if err != nil {
		t.Fatalf("Service RegisterCns failed: %+v\n", err)
	}
	if result != standardOutput {
		t.Fatalf("TestRegisterCns failed, the result %v is inconsistent with \"%v\"", result, standardOutput)
	}
	t.Logf("TestRegisterCns result: %v", result)
}

func TestGetValueByKey(t *testing.T) {
	ret0, _, err := service.GetValueByKey(key)
	if err != nil {
		t.Fatalf("Service GetValueByKey failed: %+v\n", err)
	}
	if ret0 != value {
		t.Fatalf("TestGetValueByKey failed, the ret0 %v is inconsistent with \"%v\"", ret0, value)
	}
	t.Logf("TestGetValueByKey ret0: %v", ret0)
}
package config

import (
	"testing"
	"context"

	"github.com/KasperLiu/gobcos/client"
	"github.com/KasperLiu/gobcos/crypto"
	"github.com/KasperLiu/gobcos/accounts/abi/bind"
)

func TestSetValueByKey(t *testing.T) {
	groupID := uint(1)
	rpc, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		t.Fatalf("init rpc client failed: %+v", err)
	}

	privateKey, err := crypto.HexToECDSA("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
    if err != nil {
        t.Fatalf("init privateKey failed: %+v", err)
	}

	service, err := NewSystemConfigService(rpc, privateKey)
	if err != nil {
		t.Fatalf("init SystemConfigService failed: %+v", err)
	}
	
	key := "tx_count_limit"
	value := "30000000"
	tx, err := service.SetValueByKey(key, value)
	if err != nil {
		t.Fatalf("SystemConfigService SetValueByKey failed: %+v", err)
	}
	// wait for the mining
    _, err = bind.WaitMined(context.Background(), rpc, tx)
    if err != nil {
        t.Fatalf("tx mining error:%v\n", err)
	}
	
	result, err := rpc.GetSystemConfigByKey(context.Background(), key)
	if err != nil {
		t.Fatalf("GetSystemConfigByKey failed: %v", err)
	}
	t.Logf("set value: %s, GetSystemConfigByKey: %s\n", value, result[1:len(result)-1])
	if value != string(result[1:len(result)-1]) {
		t.Fatalf("SetValueByKey failed!")
	}
	t.Logf("transaction hash: %s", tx.Hash().Hex())
}
package config

import (
	"context"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func testSetValueByKey(t *testing.T, key string, value string) {
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58", NodeURL: "http://localhost:8545"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("init client failed: %+v", err)
	}

	service, err := NewSystemConfigService(c)
	if err != nil {
		t.Fatalf("init SystemConfigService failed: %+v", err)
	}
	tx, err := service.SetValueByKey(key, value)
	if err != nil {
		t.Fatalf("SystemConfigService SetValueByKey failed: %+v", err)
	}
	// wait for the mining
	_, err = bind.WaitMined(context.Background(), c, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}

	result, err := c.GetSystemConfigByKey(context.Background(), key)
	if err != nil {
		t.Fatalf("GetSystemConfigByKey failed: %v", err)
	}
	t.Logf("set %s value: %s, GetSystemConfigByKey: %s", key, value, result[1:len(result)-1])
	if value != string(result[1:len(result)-1]) {
		t.Fatalf("SetValueByKey failed!")
	}
	t.Logf("transaction hash: %s", tx.Hash().Hex())
}

func TestSetValueByKey(t *testing.T) {
	// test tx_count_limit
	testSetValueByKey(t, "tx_count_limit", "30000000")

	// test tx_gas_limit
	testSetValueByKey(t, "tx_gas_limit", "3000000000")

	// test rpbft_epoch_sealer_num
	testSetValueByKey(t, "rpbft_epoch_sealer_num", "20")

	// test rpbft_epoch_block_num
	testSetValueByKey(t, "rpbft_epoch_block_num", "100")
}

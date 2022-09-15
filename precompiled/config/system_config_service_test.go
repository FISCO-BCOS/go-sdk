package config

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func testSetValueByKey(t *testing.T, key string, value string) {
	privateKey, _ := hex.DecodeString("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62")
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, NodeURL: "127.0.0.1:20200"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("init client failed: %+v", err)
	}
	service, err := NewSystemConfigService(c)
	if err != nil {
		t.Fatalf("init SystemConfigService failed: %+v", err)
	}

	num, err := service.SetValueByKey(key, value)
	if err != nil {
		t.Fatalf("SystemConfigService SetValueByKey failed: %+v", err)
	}
	if num != 0 {
		t.Fatalf("testSetValueByKey failed, the result %v is inconsistent with \"1\"", num)
	}

	result, err := c.GetSystemConfigByKey(context.Background(), key)
	if err != nil {
		t.Fatalf("GetSystemConfigByKey failed: %v", err)
	}
	t.Logf("set %s value: %s, GetSystemConfigByKey: %s", key, value, result.GetValue())
	//t.Logf("set %s value: %s, GetSystemConfigByKey: %s", key, value, result[1:len(result)-1])
	if value != result.GetValue() {
		t.Fatalf("SetValueByKey failed!")
	}
}

func TestSetValueByKey(t *testing.T) {
	// test tx_count_limit
	testSetValueByKey(t, "tx_count_limit", "30000000")

	// test tx_gas_limit
	testSetValueByKey(t, "tx_gas_limit", "3000000000")

	// test rpbft_epoch_sealer_num
	//testSetValueByKey(t, "rpbft_epoch_sealer_num", "20")

	// test rpbft_epoch_block_num
	//testSetValueByKey(t, "rpbft_epoch_block_num", "100")
}

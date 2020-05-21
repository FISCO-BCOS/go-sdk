package config

import (
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
)

// SystemConfigService is a precompile contract service.
type SystemConfigService struct {
	systemConfig *Config
	client       *client.Client
}

const methodSetValueByKey = "setValueByKey"

// contract address
var systemConfigPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001000")

// NewSystemConfigService returns ptr of SystemConfigService
func NewSystemConfigService(client *client.Client) (*SystemConfigService, error) {
	instance, err := NewConfig(systemConfigPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct SystemConfigService failed: %+v", err)
	}
	return &SystemConfigService{systemConfig: instance, client: client}, nil
}

// SetValueByKey returns nil if there is no error occurred.
func (s *SystemConfigService) SetValueByKey(key string, value string) (*types.Transaction, error) {
	tx, err := s.systemConfig.SetValueByKey(s.client.GetTransactOpts(), key, value)
	if err != nil {
		return tx, fmt.Errorf("SystemConfigService setValueByKey failed: %+v", err)
	}
	// receipt, _ := s.client.WaitMined(tx)
	// output := common.FromHex(receipt.Output)
	// ret := new(*big.Int)
	// err = s.abi.Unpack(ret, methodSetValueByKey, output)
	// if err != nil {
	// 	return tx, fmt.Errorf("SystemConfigService setValueByKey failed,ret:%d, err:%+v", ret, err)
	// }
	// fmt.Printf("return value:%v", ret)
	return tx, nil
}

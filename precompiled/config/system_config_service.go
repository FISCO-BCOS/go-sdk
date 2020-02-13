package config

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
)

// SystemConfigService is a precompile contract service.
type SystemConfigService struct {
	systemConfig *Config
	client       *client.Client
	abi          abi.ABI
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
	abi, err := abi.JSON(strings.NewReader(ConfigABI))
	if err != nil {
		return nil, err
	}
	abi.SMCrypto = client.SMCrypto()

	return &SystemConfigService{systemConfig: instance, client: client, abi: abi}, nil
}

// SetValueByKey returns nil if there is no error occurred.
func (s *SystemConfigService) SetValueByKey(key string, value string) (*types.Transaction, error) {
	tx, err := s.systemConfig.SetValueByKey(s.client.GetTransactOpts(), key, value)
	if err != nil {
		return tx, fmt.Errorf("SystemConfigService setValueByKey failed: %+v", err)
	}
	receipt, _ := s.client.WaitMined(tx)
	output := common.FromHex(receipt.Output)
	// var ret int
	ret := new(*big.Int)
	// ret := new(int64)
	err = s.abi.Unpack(ret, methodSetValueByKey, output)
	if err != nil {
		return tx, fmt.Errorf("SystemConfigService setValueByKey failed,ret:%d, err:%+v", ret, err)
	}

	return tx, nil
}

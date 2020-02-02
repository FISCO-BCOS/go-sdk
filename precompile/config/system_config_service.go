package config

import (
	"fmt"
	"crypto/ecdsa"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
)

// SystemConfigService is a precompile contract service.
type SystemConfigService struct {
	systemConfig *Config
	systemConfigAuth *bind.TransactOpts
}

// contract address
var systemConfigPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001000")

// NewSystemConfigService returns ptr of SystemConfigService
func NewSystemConfigService(client *client.Client, privateKey *ecdsa.PrivateKey) (*SystemConfigService, error) {
	instance, err := NewConfig(systemConfigPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct SystemConfigService failed: %+v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = big.NewInt(30000000)
    return &SystemConfigService{systemConfig:instance, systemConfigAuth:auth}, nil
}

// SetValueByKey returns a raw transaction if there is no error occured.
func (service *SystemConfigService) SetValueByKey(key string ,value string) (*types.Transaction, error) {
	tx, err := service.systemConfig.SetValueByKey(service.systemConfigAuth, key, value)
    if err != nil {
        return nil, fmt.Errorf("SystemConfigService setValueByKey failed: %+v", err)
	}
	return tx, nil
}
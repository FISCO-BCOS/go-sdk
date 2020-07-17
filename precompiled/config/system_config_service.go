package config

import (
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompiled"
	"github.com/ethereum/go-ethereum/common"
)

// SystemConfigService is a precompile contract service.
type SystemConfigService struct {
	systemConfig *Config
	client       *client.Client
}

// systemConfig precompiled contract error code
const (
	invalidConfigurationValues int64 = -51300
)

// getErrorCodeMessage returns the message of error code
func getErrorCodeMessage(errorCode int64) string {
	var message string
	switch errorCode {
	case invalidConfigurationValues:
		message = "invalid configuration values"
	default:
		message = ""
	}
	return message
}

// errorCodeToError judges whether the error code represents an error
func errorCodeToError(errorCode int64) error {
	var errorCodeMessage string
	errorCodeMessage = precompiled.GetCommonErrorCodeMessage(errorCode)
	if errorCodeMessage != "" {
		return fmt.Errorf("error code: %v, error code message: %v", errorCode, errorCodeMessage)
	}
	errorCodeMessage = getErrorCodeMessage(errorCode)
	if errorCodeMessage != "" {
		return fmt.Errorf("error code: %v, error code message: %v", errorCode, errorCodeMessage)
	}
	return nil
}

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
func (s *SystemConfigService) SetValueByKey(key string, value string) (int64, error) {
	tx, err := s.systemConfig.SetValueByKey(s.client.GetTransactOpts(), key, value)
	if err != nil {
		return types.PrecompiledError, fmt.Errorf("SystemConfigService setValueByKey failed: %+v", err)
	}
	receipt, err := s.client.WaitMined(tx)
	if err != nil {
		return types.PrecompiledError, fmt.Errorf("client.WaitMined failed, err: %v", err)
	}
	bigNum, err := precompiled.ParseBigIntFromOutput(ConfigABI, "setValueByKey", receipt)
	if err != nil {
		return types.PrecompiledError, fmt.Errorf("systemConfigService setValueByKey failed, err: %+v", err)
	}
	errorCode, err := precompiled.BigIntToInt64(bigNum)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("handleReceipt failed, err: %v", err)
	}
	return errorCode, errorCodeToError(errorCode)
}

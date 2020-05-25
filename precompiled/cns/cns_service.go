package cns

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
)

// Service is a precompile contract service.
type Service struct {
	cns     *Cns
	cnsAuth *bind.TransactOpts
}

const maxVersionLength = 40

// contract address
var cnsPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001004")

// NewCnsService returns ptr of Service
func NewCnsService(client *client.Client, privateKey *ecdsa.PrivateKey) (*Service, error) {
	instance, err := NewCns(cnsPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct Service failed: %+v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = big.NewInt(30000000)
	return &Service{cns: instance, cnsAuth: auth}, nil
}

// SelectByName returns the cns information according to the name string.
func (service *Service) SelectByName(name string) (string, error) {
	opts := &bind.CallOpts{From: service.cnsAuth.From}
	cnsName, err := service.cns.SelectByName(opts, name)
	if err != nil {
		return "", fmt.Errorf("Service SelectByName failed: %+v", err)
	}
	return cnsName, nil
}

// SelectByNameAndVersion returns the cns information according to the name string and version string.
func (service *Service) SelectByNameAndVersion(name string, version string) (string, error) {
	opts := &bind.CallOpts{From: service.cnsAuth.From}
	cnsName, err := service.cns.SelectByNameAndVersion(opts, name, version)
	if err != nil {
		return "", fmt.Errorf("Service SelectByNameAndVersion failed: %+v", err)
	}
	return cnsName, nil
}

// GetAddressByContractNameAndVersion returns the contract address.
func (service *Service) GetAddressByContractNameAndVersion(contractNameAndVersion string) (string, error) {
	if !isValidCnsName(contractNameAndVersion) {
		return contractNameAndVersion, fmt.Errorf("contractNameAndVersion is not valid")
	}
	var address string
	if strings.Contains(contractNameAndVersion, ":") {
		splited := strings.Split(contractNameAndVersion, ":")
		name := splited[0]
		version := splited[1]
		addressInfo, err := service.SelectByNameAndVersion(name, version)
		if err != nil {
			return "", err
		} else if addressInfo == "[\n]" {
			return "", fmt.Errorf("The contract version does not exist")
		}
		// json unmarshal
		var dat []Info
		if err := json.Unmarshal([]byte(addressInfo), &dat); err != nil {
			return "", fmt.Errorf("Unmarshal the addressInfo failed: %s", addressInfo)
		}
		cnsInfo := dat[len(dat)-1]
		address = cnsInfo.GetAddress()
	} else { // onlu contract name
		addressInfo, err := service.SelectByName(contractNameAndVersion)
		if err != nil {
			return "", err
		} else if addressInfo == "[\n]" {
			return "", fmt.Errorf("The contract version does not exist")
		}
		// json unmarshal
		var dat []Info
		if err := json.Unmarshal([]byte(addressInfo), &dat); err != nil {
			return "", fmt.Errorf("Unmarshal the addressInfo failed  %v", err)
		}
		cnsInfo := dat[len(dat)-1]
		address = cnsInfo.GetAddress()
	}
	if !common.IsHexAddress(address) {
		return "", fmt.Errorf("Unable to resolve address for name: %s", contractNameAndVersion)
	}
	return address, nil
}

// RegisterCns registers a contract for its CNS.
func (service *Service) RegisterCns(name string, version string, addr string, abi string) (*types.Transaction, error) {
	if len(version) > maxVersionLength {
		return nil, fmt.Errorf("version string length exceeds the maximum limit")
	}
	tx, err := service.cns.Insert(service.cnsAuth, name, version, addr, abi)
	if err != nil {
		return nil, fmt.Errorf("Service RegisterCns failed: %+v", err)
	}
	return tx, nil
}

// QueryCnsByName returns the CNS info according to the CNS name
func (service *Service) QueryCnsByName(name string) ([]Info, error) {
	cnsInfo, err := service.SelectByName(name)
	if err != nil {
		return nil, err
	}
	// json unmarshal
	var dat []Info
	if err := json.Unmarshal([]byte(cnsInfo), &dat); err != nil {
		return nil, fmt.Errorf("Unmarshal the Info failed")
	}
	return dat, nil
}

// QueryCnsByNameAndVersion returns the CNS info according to the name and version
func (service *Service) QueryCnsByNameAndVersion(name string, version string) ([]Info, error) {
	cnsInfo, err := service.SelectByNameAndVersion(name, version)
	if err != nil {
		return nil, err
	}
	// json unmarshal
	var dat []Info
	if err := json.Unmarshal([]byte(cnsInfo), &dat); err != nil {
		return nil, fmt.Errorf("Unmarshal the Info failed")
	}
	return dat, nil
}

func isValidCnsName(input string) bool {
	return input != "" && (strings.Contains(input, ":") || !common.IsHexAddress(input))
}

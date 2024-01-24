// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package config

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// ConfigABI is the input ABI used to generate the binding from.
const ConfigABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValueByKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setValueByKey\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConfigBin is the compiled bytecode used for deploying new contracts.
var ConfigBin = "0x608060405234801561001057600080fd5b50610406806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631258a93a1461003b578063bd291aef1461006c575b600080fd5b61005560048036038101906100509190610207565b61009c565b6040516100639291906102f1565b60405180910390f35b61008660048036038101906100819190610321565b6100a5565b60405161009391906103b5565b60405180910390f35b60606000915091565b600092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610114826100cb565b810181811067ffffffffffffffff82111715610133576101326100dc565b5b80604052505050565b60006101466100ad565b9050610152828261010b565b919050565b600067ffffffffffffffff821115610172576101716100dc565b5b61017b826100cb565b9050602081019050919050565b82818337600083830152505050565b60006101aa6101a584610157565b61013c565b9050828152602081018484840111156101c6576101c56100c6565b5b6101d1848285610188565b509392505050565b600082601f8301126101ee576101ed6100c1565b5b81356101fe848260208601610197565b91505092915050565b60006020828403121561021d5761021c6100b7565b5b600082013567ffffffffffffffff81111561023b5761023a6100bc565b5b610247848285016101d9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561028a57808201518184015260208101905061026f565b83811115610299576000848401525b50505050565b60006102aa82610250565b6102b4818561025b565b93506102c481856020860161026c565b6102cd816100cb565b840191505092915050565b6000819050919050565b6102eb816102d8565b82525050565b6000604082019050818103600083015261030b818561029f565b905061031a60208301846102e2565b9392505050565b60008060408385031215610338576103376100b7565b5b600083013567ffffffffffffffff811115610356576103556100bc565b5b610362858286016101d9565b925050602083013567ffffffffffffffff811115610383576103826100bc565b5b61038f858286016101d9565b9150509250929050565b60008160030b9050919050565b6103af81610399565b82525050565b60006020820190506103ca60008301846103a6565b9291505056fea26469706673582212207bb572c2d45fbfc66c4f7051acb78da418cfff1f4aaf059601a3c6d758a1c22c64736f6c634300080b0033"

// Config is an auto generated Go binding around a Solidity contract.
type Config struct {
	ConfigCaller     // Read-only binding to the contract
	ConfigTransactor // Write-only binding to the contract
	ConfigFilterer   // Log filterer for contract events
}

// ConfigCaller is an auto generated read-only Go binding around a Solidity contract.
type ConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigTransactor is an auto generated write-only Go binding around a Solidity contract.
type ConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ConfigSession struct {
	Contract     *Config           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ConfigCallerSession struct {
	Contract *ConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConfigTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ConfigTransactorSession struct {
	Contract     *ConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigRaw is an auto generated low-level Go binding around a Solidity contract.
type ConfigRaw struct {
	Contract *Config // Generic contract binding to access the raw methods on
}

// ConfigCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ConfigCallerRaw struct {
	Contract *ConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ConfigTransactorRaw struct {
	Contract *ConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfig creates a new instance of Config, bound to a specific deployed contract.
func NewConfig(address common.Address, backend bind.ContractBackend) (*Config, error) {
	contract, err := bindConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// NewConfigCaller creates a new read-only instance of Config, bound to a specific deployed contract.
func NewConfigCaller(address common.Address, caller bind.ContractCaller) (*ConfigCaller, error) {
	contract, err := bindConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigCaller{contract: contract}, nil
}

// NewConfigTransactor creates a new write-only instance of Config, bound to a specific deployed contract.
func NewConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigTransactor, error) {
	contract, err := bindConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigTransactor{contract: contract}, nil
}

// NewConfigFilterer creates a new log filterer instance of Config, bound to a specific deployed contract.
func NewConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigFilterer, error) {
	contract, err := bindConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigFilterer{contract: contract}, nil
}

// bindConfig binds a generic wrapper to an already deployed contract.
func bindConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Config.Contract.ConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Config.Contract.ConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Config.Contract.ConfigTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Config.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Config.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Config.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetValueByKey is a free data retrieval call binding the contract method 0x1258a93a.
//
// Solidity: function getValueByKey(string key) constant returns(string, int256)
func (_Config *ConfigCaller) GetValueByKey(opts *bind.CallOpts, key string) (string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Config.contract.Call(opts, out, "getValueByKey", key)
	return *ret0, *ret1, err
}

// GetValueByKey is a free data retrieval call binding the contract method 0x1258a93a.
//
// Solidity: function getValueByKey(string key) constant returns(string, int256)
func (_Config *ConfigSession) GetValueByKey(key string) (string, *big.Int, error) {
	return _Config.Contract.GetValueByKey(&_Config.CallOpts, key)
}

// GetValueByKey is a free data retrieval call binding the contract method 0x1258a93a.
//
// Solidity: function getValueByKey(string key) constant returns(string, int256)
func (_Config *ConfigCallerSession) GetValueByKey(key string) (string, *big.Int, error) {
	return _Config.Contract.GetValueByKey(&_Config.CallOpts, key)
}

// SetValueByKey is a paid mutator transaction binding the contract method 0xbd291aef.
//
// Solidity: function setValueByKey(string key, string value) returns(int32)
func (_Config *ConfigTransactor) SetValueByKey(opts *bind.TransactOpts, key string, value string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Config.contract.TransactWithResult(opts, out, "setValueByKey", key, value)
	return *ret0, transaction, receipt, err
}

func (_Config *ConfigTransactor) AsyncSetValueByKey(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Config.contract.AsyncTransact(opts, handler, "setValueByKey", key, value)
}

// SetValueByKey is a paid mutator transaction binding the contract method 0xbd291aef.
//
// Solidity: function setValueByKey(string key, string value) returns(int32)
func (_Config *ConfigSession) SetValueByKey(key string, value string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Config.Contract.SetValueByKey(&_Config.TransactOpts, key, value)
}

func (_Config *ConfigSession) AsyncSetValueByKey(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _Config.Contract.AsyncSetValueByKey(handler, &_Config.TransactOpts, key, value)
}

// SetValueByKey is a paid mutator transaction binding the contract method 0xbd291aef.
//
// Solidity: function setValueByKey(string key, string value) returns(int32)
func (_Config *ConfigTransactorSession) SetValueByKey(key string, value string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Config.Contract.SetValueByKey(&_Config.TransactOpts, key, value)
}

func (_Config *ConfigTransactorSession) AsyncSetValueByKey(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _Config.Contract.AsyncSetValueByKey(handler, &_Config.TransactOpts, key, value)
}

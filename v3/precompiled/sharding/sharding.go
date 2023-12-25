// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sharding

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// ShardingABI is the input ABI used to generate the binding from.
const ShardingABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"}],\"name\":\"getContractShard\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shardName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"linkShard\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shardName\",\"type\":\"string\"}],\"name\":\"makeShard\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ShardingBin is the compiled bytecode used for deploying new contracts.
var ShardingBin = "0x608060405234801561001057600080fd5b5061042c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631d82d998146100465780639c1284bc14610076578063b7ede6cb146100a7575b600080fd5b610060600480360381019061005b9190610249565b6100d7565b60405161006d91906102ab565b60405180910390f35b610090600480360381019061008b9190610249565b6100de565b60405161009e92919061034e565b60405180910390f35b6100c160048036038101906100bc919061037e565b6100e7565b6040516100ce91906102ab565b60405180910390f35b6000919050565b60006060915091565b600092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101568261010d565b810181811067ffffffffffffffff821117156101755761017461011e565b5b80604052505050565b60006101886100ef565b9050610194828261014d565b919050565b600067ffffffffffffffff8211156101b4576101b361011e565b5b6101bd8261010d565b9050602081019050919050565b82818337600083830152505050565b60006101ec6101e784610199565b61017e565b90508281526020810184848401111561020857610207610108565b5b6102138482856101ca565b509392505050565b600082601f8301126102305761022f610103565b5b81356102408482602086016101d9565b91505092915050565b60006020828403121561025f5761025e6100f9565b5b600082013567ffffffffffffffff81111561027d5761027c6100fe565b5b6102898482850161021b565b91505092915050565b6000819050919050565b6102a581610292565b82525050565b60006020820190506102c0600083018461029c565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156103005780820151818401526020810190506102e5565b8381111561030f576000848401525b50505050565b6000610320826102c6565b61032a81856102d1565b935061033a8185602086016102e2565b6103438161010d565b840191505092915050565b6000604082019050610363600083018561029c565b81810360208301526103758184610315565b90509392505050565b60008060408385031215610395576103946100f9565b5b600083013567ffffffffffffffff8111156103b3576103b26100fe565b5b6103bf8582860161021b565b925050602083013567ffffffffffffffff8111156103e0576103df6100fe565b5b6103ec8582860161021b565b915050925092905056fea2646970667358221220baad85981073d50ee009743acebf2bece8d1deb0d1966340563e3341bedf4da464736f6c634300080b0033"

// DeploySharding deploys a new contract, binding an instance of Sharding to it.
func DeploySharding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Sharding, error) {
	parsed, err := abi.JSON(strings.NewReader(ShardingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShardingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Sharding{ShardingCaller: ShardingCaller{contract: contract}, ShardingTransactor: ShardingTransactor{contract: contract}, ShardingFilterer: ShardingFilterer{contract: contract}}, nil
}

func AsyncDeploySharding(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ShardingABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ShardingBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Sharding is an auto generated Go binding around a Solidity contract.
type Sharding struct {
	ShardingCaller     // Read-only binding to the contract
	ShardingTransactor // Write-only binding to the contract
	ShardingFilterer   // Log filterer for contract events
}

// ShardingCaller is an auto generated read-only Go binding around a Solidity contract.
type ShardingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShardingTransactor is an auto generated write-only Go binding around a Solidity contract.
type ShardingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShardingFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ShardingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShardingSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ShardingSession struct {
	Contract     *Sharding         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShardingCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ShardingCallerSession struct {
	Contract *ShardingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ShardingTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ShardingTransactorSession struct {
	Contract     *ShardingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ShardingRaw is an auto generated low-level Go binding around a Solidity contract.
type ShardingRaw struct {
	Contract *Sharding // Generic contract binding to access the raw methods on
}

// ShardingCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ShardingCallerRaw struct {
	Contract *ShardingCaller // Generic read-only contract binding to access the raw methods on
}

// ShardingTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ShardingTransactorRaw struct {
	Contract *ShardingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSharding creates a new instance of Sharding, bound to a specific deployed contract.
func NewSharding(address common.Address, backend bind.ContractBackend) (*Sharding, error) {
	contract, err := bindSharding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sharding{ShardingCaller: ShardingCaller{contract: contract}, ShardingTransactor: ShardingTransactor{contract: contract}, ShardingFilterer: ShardingFilterer{contract: contract}}, nil
}

// NewShardingCaller creates a new read-only instance of Sharding, bound to a specific deployed contract.
func NewShardingCaller(address common.Address, caller bind.ContractCaller) (*ShardingCaller, error) {
	contract, err := bindSharding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShardingCaller{contract: contract}, nil
}

// NewShardingTransactor creates a new write-only instance of Sharding, bound to a specific deployed contract.
func NewShardingTransactor(address common.Address, transactor bind.ContractTransactor) (*ShardingTransactor, error) {
	contract, err := bindSharding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShardingTransactor{contract: contract}, nil
}

// NewShardingFilterer creates a new log filterer instance of Sharding, bound to a specific deployed contract.
func NewShardingFilterer(address common.Address, filterer bind.ContractFilterer) (*ShardingFilterer, error) {
	contract, err := bindSharding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShardingFilterer{contract: contract}, nil
}

// bindSharding binds a generic wrapper to an already deployed contract.
func bindSharding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShardingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sharding *ShardingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sharding.Contract.ShardingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sharding *ShardingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.ShardingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sharding *ShardingRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.ShardingTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sharding *ShardingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sharding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sharding *ShardingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sharding *ShardingTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetContractShard is a free data retrieval call binding the contract method 0x9c1284bc.
//
// Solidity: function getContractShard(string absolutePath) constant returns(int256, string)
func (_Sharding *ShardingCaller) GetContractShard(opts *bind.CallOpts, absolutePath string) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Sharding.contract.Call(opts, out, "getContractShard", absolutePath)
	return *ret0, *ret1, err
}

// GetContractShard is a free data retrieval call binding the contract method 0x9c1284bc.
//
// Solidity: function getContractShard(string absolutePath) constant returns(int256, string)
func (_Sharding *ShardingSession) GetContractShard(absolutePath string) (*big.Int, string, error) {
	return _Sharding.Contract.GetContractShard(&_Sharding.CallOpts, absolutePath)
}

// GetContractShard is a free data retrieval call binding the contract method 0x9c1284bc.
//
// Solidity: function getContractShard(string absolutePath) constant returns(int256, string)
func (_Sharding *ShardingCallerSession) GetContractShard(absolutePath string) (*big.Int, string, error) {
	return _Sharding.Contract.GetContractShard(&_Sharding.CallOpts, absolutePath)
}

// LinkShard is a paid mutator transaction binding the contract method 0xb7ede6cb.
//
// Solidity: function linkShard(string shardName, string _address) returns(int256)
func (_Sharding *ShardingTransactor) LinkShard(opts *bind.TransactOpts, shardName string, _address string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Sharding.contract.TransactWithResult(opts, out, "linkShard", shardName, _address)
	return *ret0, transaction, receipt, err
}

func (_Sharding *ShardingTransactor) AsyncLinkShard(handler func(*types.Receipt, error), opts *bind.TransactOpts, shardName string, _address string) (*types.Transaction, error) {
	return _Sharding.contract.AsyncTransact(opts, handler, "linkShard", shardName, _address)
}

// LinkShard is a paid mutator transaction binding the contract method 0xb7ede6cb.
//
// Solidity: function linkShard(string shardName, string _address) returns(int256)
func (_Sharding *ShardingSession) LinkShard(shardName string, _address string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.LinkShard(&_Sharding.TransactOpts, shardName, _address)
}

func (_Sharding *ShardingSession) AsyncLinkShard(handler func(*types.Receipt, error), shardName string, _address string) (*types.Transaction, error) {
	return _Sharding.Contract.AsyncLinkShard(handler, &_Sharding.TransactOpts, shardName, _address)
}

// LinkShard is a paid mutator transaction binding the contract method 0xb7ede6cb.
//
// Solidity: function linkShard(string shardName, string _address) returns(int256)
func (_Sharding *ShardingTransactorSession) LinkShard(shardName string, _address string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.LinkShard(&_Sharding.TransactOpts, shardName, _address)
}

func (_Sharding *ShardingTransactorSession) AsyncLinkShard(handler func(*types.Receipt, error), shardName string, _address string) (*types.Transaction, error) {
	return _Sharding.Contract.AsyncLinkShard(handler, &_Sharding.TransactOpts, shardName, _address)
}

// MakeShard is a paid mutator transaction binding the contract method 0x1d82d998.
//
// Solidity: function makeShard(string shardName) returns(int256)
func (_Sharding *ShardingTransactor) MakeShard(opts *bind.TransactOpts, shardName string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Sharding.contract.TransactWithResult(opts, out, "makeShard", shardName)
	return *ret0, transaction, receipt, err
}

func (_Sharding *ShardingTransactor) AsyncMakeShard(handler func(*types.Receipt, error), opts *bind.TransactOpts, shardName string) (*types.Transaction, error) {
	return _Sharding.contract.AsyncTransact(opts, handler, "makeShard", shardName)
}

// MakeShard is a paid mutator transaction binding the contract method 0x1d82d998.
//
// Solidity: function makeShard(string shardName) returns(int256)
func (_Sharding *ShardingSession) MakeShard(shardName string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.MakeShard(&_Sharding.TransactOpts, shardName)
}

func (_Sharding *ShardingSession) AsyncMakeShard(handler func(*types.Receipt, error), shardName string) (*types.Transaction, error) {
	return _Sharding.Contract.AsyncMakeShard(handler, &_Sharding.TransactOpts, shardName)
}

// MakeShard is a paid mutator transaction binding the contract method 0x1d82d998.
//
// Solidity: function makeShard(string shardName) returns(int256)
func (_Sharding *ShardingTransactorSession) MakeShard(shardName string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Sharding.Contract.MakeShard(&_Sharding.TransactOpts, shardName)
}

func (_Sharding *ShardingTransactorSession) AsyncMakeShard(handler func(*types.Receipt, error), shardName string) (*types.Transaction, error) {
	return _Sharding.Contract.AsyncMakeShard(handler, &_Sharding.TransactOpts, shardName)
}

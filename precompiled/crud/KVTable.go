// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crud

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
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

// KVTableABI is the input ABI used to generate the binding from.
const KVTableABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KVTable is an auto generated Go binding around a Solidity contract.
type KVTable struct {
	KVTableCaller     // Read-only binding to the contract
	KVTableTransactor // Write-only binding to the contract
	KVTableFilterer   // Log filterer for contract events
}

// KVTableCaller is an auto generated read-only Go binding around a Solidity contract.
type KVTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableTransactor is an auto generated write-only Go binding around a Solidity contract.
type KVTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type KVTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type KVTableSession struct {
	Contract     *KVTable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KVTableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type KVTableCallerSession struct {
	Contract *KVTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KVTableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type KVTableTransactorSession struct {
	Contract     *KVTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KVTableRaw is an auto generated low-level Go binding around a Solidity contract.
type KVTableRaw struct {
	Contract *KVTable // Generic contract binding to access the raw methods on
}

// KVTableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type KVTableCallerRaw struct {
	Contract *KVTableCaller // Generic read-only contract binding to access the raw methods on
}

// KVTableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type KVTableTransactorRaw struct {
	Contract *KVTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKVTable creates a new instance of KVTable, bound to a specific deployed contract.
func NewKVTable(address common.Address, backend bind.ContractBackend) (*KVTable, error) {
	contract, err := bindKVTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KVTable{KVTableCaller: KVTableCaller{contract: contract}, KVTableTransactor: KVTableTransactor{contract: contract}, KVTableFilterer: KVTableFilterer{contract: contract}}, nil
}

// NewKVTableCaller creates a new read-only instance of KVTable, bound to a specific deployed contract.
func NewKVTableCaller(address common.Address, caller bind.ContractCaller) (*KVTableCaller, error) {
	contract, err := bindKVTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableCaller{contract: contract}, nil
}

// NewKVTableTransactor creates a new write-only instance of KVTable, bound to a specific deployed contract.
func NewKVTableTransactor(address common.Address, transactor bind.ContractTransactor) (*KVTableTransactor, error) {
	contract, err := bindKVTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableTransactor{contract: contract}, nil
}

// NewKVTableFilterer creates a new log filterer instance of KVTable, bound to a specific deployed contract.
func NewKVTableFilterer(address common.Address, filterer bind.ContractFilterer) (*KVTableFilterer, error) {
	contract, err := bindKVTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KVTableFilterer{contract: contract}, nil
}

// bindKVTable binds a generic wrapper to an already deployed contract.
func bindKVTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KVTableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTable *KVTableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTable.Contract.KVTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTable *KVTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _KVTable.Contract.KVTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTable *KVTableRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _KVTable.Contract.KVTableTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTable *KVTableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTable *KVTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _KVTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTable *KVTableTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _KVTable.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(bool, string)
func (_KVTable *KVTableCaller) Get(opts *bind.CallOpts, key string) (bool, string, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _KVTable.contract.Call(opts, out, "get", key)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(bool, string)
func (_KVTable *KVTableSession) Get(key string) (bool, string, error) {
	return _KVTable.Contract.Get(&_KVTable.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(bool, string)
func (_KVTable *KVTableCallerSession) Get(key string) (bool, string, error) {
	return _KVTable.Contract.Get(&_KVTable.CallOpts, key)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns(int32)
func (_KVTable *KVTableTransactor) Set(opts *bind.TransactOpts, key string, value string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _KVTable.contract.TransactWithResult(opts, out, "set", key, value)
	return *ret0, transaction, receipt, err
}

func (_KVTable *KVTableTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _KVTable.contract.AsyncTransact(opts, handler, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns(int32)
func (_KVTable *KVTableSession) Set(key string, value string) (int32, *types.Transaction, *types.Receipt, error) {
	return _KVTable.Contract.Set(&_KVTable.TransactOpts, key, value)
}

func (_KVTable *KVTableSession) AsyncSet(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _KVTable.Contract.AsyncSet(handler, &_KVTable.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns(int32)
func (_KVTable *KVTableTransactorSession) Set(key string, value string) (int32, *types.Transaction, *types.Receipt, error) {
	return _KVTable.Contract.Set(&_KVTable.TransactOpts, key, value)
}

func (_KVTable *KVTableTransactorSession) AsyncSet(handler func(*types.Receipt, error), key string, value string) (*types.Transaction, error) {
	return _KVTable.Contract.AsyncSet(handler, &_KVTable.TransactOpts, key, value)
}

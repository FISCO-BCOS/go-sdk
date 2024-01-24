// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auth

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

// AccountManagerABI is the input ABI used to generate the binding from.
const AccountManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAccountStatus\",\"outputs\":[{\"internalType\":\"enumAccountStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"enumAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setAccountStatus\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccountManager is an auto generated Go binding around a Solidity contract.
type AccountManager struct {
	AccountManagerCaller     // Read-only binding to the contract
	AccountManagerTransactor // Write-only binding to the contract
	AccountManagerFilterer   // Log filterer for contract events
}

// AccountManagerCaller is an auto generated read-only Go binding around a Solidity contract.
type AccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerTransactor is an auto generated write-only Go binding around a Solidity contract.
type AccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type AccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type AccountManagerSession struct {
	Contract     *AccountManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountManagerCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type AccountManagerCallerSession struct {
	Contract *AccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountManagerTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type AccountManagerTransactorSession struct {
	Contract     *AccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountManagerRaw is an auto generated low-level Go binding around a Solidity contract.
type AccountManagerRaw struct {
	Contract *AccountManager // Generic contract binding to access the raw methods on
}

// AccountManagerCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type AccountManagerCallerRaw struct {
	Contract *AccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type AccountManagerTransactorRaw struct {
	Contract *AccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManager creates a new instance of AccountManager, bound to a specific deployed contract.
func NewAccountManager(address common.Address, backend bind.ContractBackend) (*AccountManager, error) {
	contract, err := bindAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// NewAccountManagerCaller creates a new read-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerCaller, error) {
	contract, err := bindAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCaller{contract: contract}, nil
}

// NewAccountManagerTransactor creates a new write-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerTransactor, error) {
	contract, err := bindAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerTransactor{contract: contract}, nil
}

// NewAccountManagerFilterer creates a new log filterer instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFilterer, error) {
	contract, err := bindAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFilterer{contract: contract}, nil
}

// bindAccountManager binds a generic wrapper to an already deployed contract.
func bindAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.AccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _AccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _AccountManager.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xfd4fa05a.
//
// Solidity: function getAccountStatus(address addr) constant returns(uint8)
func (_AccountManager *AccountManagerCaller) GetAccountStatus(opts *bind.CallOpts, addr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getAccountStatus", addr)
	return *ret0, err
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xfd4fa05a.
//
// Solidity: function getAccountStatus(address addr) constant returns(uint8)
func (_AccountManager *AccountManagerSession) GetAccountStatus(addr common.Address) (uint8, error) {
	return _AccountManager.Contract.GetAccountStatus(&_AccountManager.CallOpts, addr)
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xfd4fa05a.
//
// Solidity: function getAccountStatus(address addr) constant returns(uint8)
func (_AccountManager *AccountManagerCallerSession) GetAccountStatus(addr common.Address) (uint8, error) {
	return _AccountManager.Contract.GetAccountStatus(&_AccountManager.CallOpts, addr)
}

// SetAccountStatus is a paid mutator transaction binding the contract method 0x0ad2b0a1.
//
// Solidity: function setAccountStatus(address addr, uint8 status) returns(int32)
func (_AccountManager *AccountManagerTransactor) SetAccountStatus(opts *bind.TransactOpts, addr common.Address, status uint8) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _AccountManager.contract.TransactWithResult(opts, out, "setAccountStatus", addr, status)
	return *ret0, transaction, receipt, err
}

func (_AccountManager *AccountManagerTransactor) AsyncSetAccountStatus(handler func(*types.Receipt, error), opts *bind.TransactOpts, addr common.Address, status uint8) (*types.Transaction, error) {
	return _AccountManager.contract.AsyncTransact(opts, handler, "setAccountStatus", addr, status)
}

// SetAccountStatus is a paid mutator transaction binding the contract method 0x0ad2b0a1.
//
// Solidity: function setAccountStatus(address addr, uint8 status) returns(int32)
func (_AccountManager *AccountManagerSession) SetAccountStatus(addr common.Address, status uint8) (int32, *types.Transaction, *types.Receipt, error) {
	return _AccountManager.Contract.SetAccountStatus(&_AccountManager.TransactOpts, addr, status)
}

func (_AccountManager *AccountManagerSession) AsyncSetAccountStatus(handler func(*types.Receipt, error), addr common.Address, status uint8) (*types.Transaction, error) {
	return _AccountManager.Contract.AsyncSetAccountStatus(handler, &_AccountManager.TransactOpts, addr, status)
}

// SetAccountStatus is a paid mutator transaction binding the contract method 0x0ad2b0a1.
//
// Solidity: function setAccountStatus(address addr, uint8 status) returns(int32)
func (_AccountManager *AccountManagerTransactorSession) SetAccountStatus(addr common.Address, status uint8) (int32, *types.Transaction, *types.Receipt, error) {
	return _AccountManager.Contract.SetAccountStatus(&_AccountManager.TransactOpts, addr, status)
}

func (_AccountManager *AccountManagerTransactorSession) AsyncSetAccountStatus(handler func(*types.Receipt, error), addr common.Address, status uint8) (*types.Transaction, error) {
	return _AccountManager.Contract.AsyncSetAccountStatus(handler, &_AccountManager.TransactOpts, addr, status)
}

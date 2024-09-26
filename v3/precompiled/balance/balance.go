// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balanceprecompiled

import (
	"fmt"
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

// BalancePrecompiledABI is the input ABI used to generate the binding from.
const BalancePrecompiledABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listCaller\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"registerCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unregisterCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalancePrecompiledBin is the compiled bytecode used for deploying new contracts.
var BalancePrecompiledBin = "0x608060405234801561001057600080fd5b50610414806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063afb4bfbd1161005b578063afb4bfbd146100d8578063beabacc8146100f4578063cf8eeb7e14610110578063f8b2cb4f1461012c5761007d565b806321e5383a146100825780633b9a32de1461009e578063511ab10c146100ba575b600080fd5b61009c60048036038101906100979190610214565b61015c565b005b6100b860048036038101906100b39190610254565b610160565b005b6100c2610163565b6040516100cf919061033f565b60405180910390f35b6100f260048036038101906100ed9190610254565b610168565b005b61010e60048036038101906101099190610361565b61016b565b005b61012a60048036038101906101259190610214565b610170565b005b61014660048036038101906101419190610254565b610174565b60405161015391906103c3565b60405180910390f35b5050565b50565b606090565b50565b505050565b5050565b6000919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101ab82610180565b9050919050565b6101bb816101a0565b81146101c657600080fd5b50565b6000813590506101d8816101b2565b92915050565b6000819050919050565b6101f1816101de565b81146101fc57600080fd5b50565b60008135905061020e816101e8565b92915050565b6000806040838503121561022b5761022a61017b565b5b6000610239858286016101c9565b925050602061024a858286016101ff565b9150509250929050565b60006020828403121561026a5761026961017b565b5b6000610278848285016101c9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6102b6816101a0565b82525050565b60006102c883836102ad565b60208301905092915050565b6000602082019050919050565b60006102ec82610281565b6102f6818561028c565b93506103018361029d565b8060005b8381101561033257815161031988826102bc565b9750610324836102d4565b925050600181019050610305565b5085935050505092915050565b6000602082019050818103600083015261035981846102e1565b905092915050565b60008060006060848603121561037a5761037961017b565b5b6000610388868287016101c9565b9350506020610399868287016101c9565b92505060406103aa868287016101ff565b9150509250925092565b6103bd816101de565b82525050565b60006020820190506103d860008301846103b4565b9291505056fea2646970667358221220f4f1f0e50be54ee5394905a711259b7039cdf627b4521c2148afb1e185ed7c2a64736f6c634300080b0033"
var BalancePrecompiledSMBin = "0x"

// DeployBalancePrecompiled deploys a new contract, binding an instance of BalancePrecompiled to it.
func DeployBalancePrecompiled(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *BalancePrecompiled, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancePrecompiledABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(BalancePrecompiledSMBin)
	} else {
		bytecode = common.FromHex(BalancePrecompiledBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, BalancePrecompiledABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &BalancePrecompiled{BalancePrecompiledCaller: BalancePrecompiledCaller{contract: contract}, BalancePrecompiledTransactor: BalancePrecompiledTransactor{contract: contract}, BalancePrecompiledFilterer: BalancePrecompiledFilterer{contract: contract}}, nil
}

func AsyncDeployBalancePrecompiled(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancePrecompiledABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(BalancePrecompiledSMBin)
	} else {
		bytecode = common.FromHex(BalancePrecompiledBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, BalancePrecompiledABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// BalancePrecompiled is an auto generated Go binding around a Solidity contract.
type BalancePrecompiled struct {
	BalancePrecompiledCaller     // Read-only binding to the contract
	BalancePrecompiledTransactor // Write-only binding to the contract
	BalancePrecompiledFilterer   // Log filterer for contract events
}

// BalancePrecompiledCaller is an auto generated read-only Go binding around a Solidity contract.
type BalancePrecompiledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancePrecompiledTransactor is an auto generated write-only Go binding around a Solidity contract.
type BalancePrecompiledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancePrecompiledFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type BalancePrecompiledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancePrecompiledSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type BalancePrecompiledSession struct {
	Contract     *BalancePrecompiled // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BalancePrecompiledCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type BalancePrecompiledCallerSession struct {
	Contract *BalancePrecompiledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BalancePrecompiledTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type BalancePrecompiledTransactorSession struct {
	Contract     *BalancePrecompiledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BalancePrecompiledRaw is an auto generated low-level Go binding around a Solidity contract.
type BalancePrecompiledRaw struct {
	Contract *BalancePrecompiled // Generic contract binding to access the raw methods on
}

// BalancePrecompiledCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type BalancePrecompiledCallerRaw struct {
	Contract *BalancePrecompiledCaller // Generic read-only contract binding to access the raw methods on
}

// BalancePrecompiledTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type BalancePrecompiledTransactorRaw struct {
	Contract *BalancePrecompiledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancePrecompiled creates a new instance of BalancePrecompiled, bound to a specific deployed contract.
func NewBalancePrecompiled(address common.Address, backend bind.ContractBackend) (*BalancePrecompiled, error) {
	contract, err := bindBalancePrecompiled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancePrecompiled{BalancePrecompiledCaller: BalancePrecompiledCaller{contract: contract}, BalancePrecompiledTransactor: BalancePrecompiledTransactor{contract: contract}, BalancePrecompiledFilterer: BalancePrecompiledFilterer{contract: contract}}, nil
}

// NewBalancePrecompiledCaller creates a new read-only instance of BalancePrecompiled, bound to a specific deployed contract.
func NewBalancePrecompiledCaller(address common.Address, caller bind.ContractCaller) (*BalancePrecompiledCaller, error) {
	contract, err := bindBalancePrecompiled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancePrecompiledCaller{contract: contract}, nil
}

// NewBalancePrecompiledTransactor creates a new write-only instance of BalancePrecompiled, bound to a specific deployed contract.
func NewBalancePrecompiledTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancePrecompiledTransactor, error) {
	contract, err := bindBalancePrecompiled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancePrecompiledTransactor{contract: contract}, nil
}

// NewBalancePrecompiledFilterer creates a new log filterer instance of BalancePrecompiled, bound to a specific deployed contract.
func NewBalancePrecompiledFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancePrecompiledFilterer, error) {
	contract, err := bindBalancePrecompiled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancePrecompiledFilterer{contract: contract}, nil
}

// bindBalancePrecompiled binds a generic wrapper to an already deployed contract.
func bindBalancePrecompiled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancePrecompiledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancePrecompiled *BalancePrecompiledRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancePrecompiled.Contract.BalancePrecompiledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancePrecompiled *BalancePrecompiledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.BalancePrecompiledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancePrecompiled *BalancePrecompiledRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.BalancePrecompiledTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancePrecompiled *BalancePrecompiledCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancePrecompiled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancePrecompiled *BalancePrecompiledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancePrecompiled *BalancePrecompiledTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) constant returns(uint256)
func (_BalancePrecompiled *BalancePrecompiledCaller) GetBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancePrecompiled.contract.Call(opts, out, "getBalance", account)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) constant returns(uint256)
func (_BalancePrecompiled *BalancePrecompiledSession) GetBalance(account common.Address) (*big.Int, error) {
	return _BalancePrecompiled.Contract.GetBalance(&_BalancePrecompiled.CallOpts, account)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) constant returns(uint256)
func (_BalancePrecompiled *BalancePrecompiledCallerSession) GetBalance(account common.Address) (*big.Int, error) {
	return _BalancePrecompiled.Contract.GetBalance(&_BalancePrecompiled.CallOpts, account)
}

// ListCaller is a free data retrieval call binding the contract method 0x511ab10c.
//
// Solidity: function listCaller() constant returns(address[])
func (_BalancePrecompiled *BalancePrecompiledCaller) ListCaller(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancePrecompiled.contract.Call(opts, out, "listCaller")
	return *ret0, err
}

// ListCaller is a free data retrieval call binding the contract method 0x511ab10c.
//
// Solidity: function listCaller() constant returns(address[])
func (_BalancePrecompiled *BalancePrecompiledSession) ListCaller() ([]common.Address, error) {
	return _BalancePrecompiled.Contract.ListCaller(&_BalancePrecompiled.CallOpts)
}

// ListCaller is a free data retrieval call binding the contract method 0x511ab10c.
//
// Solidity: function listCaller() constant returns(address[])
func (_BalancePrecompiled *BalancePrecompiledCallerSession) ListCaller() ([]common.Address, error) {
	return _BalancePrecompiled.Contract.ListCaller(&_BalancePrecompiled.CallOpts)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(address account, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactor) AddBalance(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _BalancePrecompiled.contract.TransactWithResult(opts, out, "addBalance", account, amount)
	return transaction, receipt, err
}

func (_BalancePrecompiled *BalancePrecompiledTransactor) AsyncAddBalance(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.contract.AsyncTransact(opts, handler, "addBalance", account, amount)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(address account, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledSession) AddBalance(account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.AddBalance(&_BalancePrecompiled.TransactOpts, account, amount)
}

func (_BalancePrecompiled *BalancePrecompiledSession) AsyncAddBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncAddBalance(handler, &_BalancePrecompiled.TransactOpts, account, amount)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(address account, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactorSession) AddBalance(account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.AddBalance(&_BalancePrecompiled.TransactOpts, account, amount)
}

func (_BalancePrecompiled *BalancePrecompiledTransactorSession) AsyncAddBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncAddBalance(handler, &_BalancePrecompiled.TransactOpts, account, amount)
}

// RegisterCaller is a paid mutator transaction binding the contract method 0x3b9a32de.
//
// Solidity: function registerCaller(address account) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactor) RegisterCaller(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _BalancePrecompiled.contract.TransactWithResult(opts, out, "registerCaller", account)
	return transaction, receipt, err
}

func (_BalancePrecompiled *BalancePrecompiledTransactor) AsyncRegisterCaller(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BalancePrecompiled.contract.AsyncTransact(opts, handler, "registerCaller", account)
}

// RegisterCaller is a paid mutator transaction binding the contract method 0x3b9a32de.
//
// Solidity: function registerCaller(address account) returns()
func (_BalancePrecompiled *BalancePrecompiledSession) RegisterCaller(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.RegisterCaller(&_BalancePrecompiled.TransactOpts, account)
}

func (_BalancePrecompiled *BalancePrecompiledSession) AsyncRegisterCaller(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncRegisterCaller(handler, &_BalancePrecompiled.TransactOpts, account)
}

// RegisterCaller is a paid mutator transaction binding the contract method 0x3b9a32de.
//
// Solidity: function registerCaller(address account) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactorSession) RegisterCaller(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.RegisterCaller(&_BalancePrecompiled.TransactOpts, account)
}

func (_BalancePrecompiled *BalancePrecompiledTransactorSession) AsyncRegisterCaller(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncRegisterCaller(handler, &_BalancePrecompiled.TransactOpts, account)
}

// SubBalance is a paid mutator transaction binding the contract method 0xcf8eeb7e.
//
// Solidity: function subBalance(address account, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactor) SubBalance(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _BalancePrecompiled.contract.TransactWithResult(opts, out, "subBalance", account, amount)
	return transaction, receipt, err
}

func (_BalancePrecompiled *BalancePrecompiledTransactor) AsyncSubBalance(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.contract.AsyncTransact(opts, handler, "subBalance", account, amount)
}

// SubBalance is a paid mutator transaction binding the contract method 0xcf8eeb7e.
//
// Solidity: function subBalance(address account, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledSession) SubBalance(account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.SubBalance(&_BalancePrecompiled.TransactOpts, account, amount)
}

func (_BalancePrecompiled *BalancePrecompiledSession) AsyncSubBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncSubBalance(handler, &_BalancePrecompiled.TransactOpts, account, amount)
}

// SubBalance is a paid mutator transaction binding the contract method 0xcf8eeb7e.
//
// Solidity: function subBalance(address account, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactorSession) SubBalance(account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.SubBalance(&_BalancePrecompiled.TransactOpts, account, amount)
}

func (_BalancePrecompiled *BalancePrecompiledTransactorSession) AsyncSubBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncSubBalance(handler, &_BalancePrecompiled.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _BalancePrecompiled.contract.TransactWithResult(opts, out, "transfer", from, to, amount)
	return transaction, receipt, err
}

func (_BalancePrecompiled *BalancePrecompiledTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.contract.AsyncTransact(opts, handler, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.Transfer(&_BalancePrecompiled.TransactOpts, from, to, amount)
}

func (_BalancePrecompiled *BalancePrecompiledSession) AsyncTransfer(handler func(*types.Receipt, error), from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncTransfer(handler, &_BalancePrecompiled.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactorSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.Transfer(&_BalancePrecompiled.TransactOpts, from, to, amount)
}

func (_BalancePrecompiled *BalancePrecompiledTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncTransfer(handler, &_BalancePrecompiled.TransactOpts, from, to, amount)
}

// UnregisterCaller is a paid mutator transaction binding the contract method 0xafb4bfbd.
//
// Solidity: function unregisterCaller(address account) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactor) UnregisterCaller(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _BalancePrecompiled.contract.TransactWithResult(opts, out, "unregisterCaller", account)
	return transaction, receipt, err
}

func (_BalancePrecompiled *BalancePrecompiledTransactor) AsyncUnregisterCaller(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BalancePrecompiled.contract.AsyncTransact(opts, handler, "unregisterCaller", account)
}

// UnregisterCaller is a paid mutator transaction binding the contract method 0xafb4bfbd.
//
// Solidity: function unregisterCaller(address account) returns()
func (_BalancePrecompiled *BalancePrecompiledSession) UnregisterCaller(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.UnregisterCaller(&_BalancePrecompiled.TransactOpts, account)
}

func (_BalancePrecompiled *BalancePrecompiledSession) AsyncUnregisterCaller(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncUnregisterCaller(handler, &_BalancePrecompiled.TransactOpts, account)
}

// UnregisterCaller is a paid mutator transaction binding the contract method 0xafb4bfbd.
//
// Solidity: function unregisterCaller(address account) returns()
func (_BalancePrecompiled *BalancePrecompiledTransactorSession) UnregisterCaller(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _BalancePrecompiled.Contract.UnregisterCaller(&_BalancePrecompiled.TransactOpts, account)
}

func (_BalancePrecompiled *BalancePrecompiledTransactorSession) AsyncUnregisterCaller(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _BalancePrecompiled.Contract.AsyncUnregisterCaller(handler, &_BalancePrecompiled.TransactOpts, account)
}

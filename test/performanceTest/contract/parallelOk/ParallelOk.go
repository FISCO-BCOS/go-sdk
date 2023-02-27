// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parallelOk

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

// ParallelOkABI is the input ABI used to generate the binding from.
const ParallelOkABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transferWithRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ParallelOkBin is the compiled bytecode used for deploying new contracts.
var ParallelOkBin = "0x608060405234801561001057600080fd5b506107ab806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806335ee5f87146100515780638a42ebe9146101205780639b80b050146101e5578063fad42f8714610341575b600080fd5b61010a6004803603602081101561006757600080fd5b810190808035906020019064010000000081111561008457600080fd5b82018360208201111561009657600080fd5b803590602001918460018302840111640100000000831117156100b857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061049d565b6040518082815260200191505060405180910390f35b6101e36004803603604081101561013657600080fd5b810190808035906020019064010000000081111561015357600080fd5b82018360208201111561016557600080fd5b8035906020019184600183028401116401000000008311171561018757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061050f565b005b61033f600480360360608110156101fb57600080fd5b810190808035906020019064010000000081111561021857600080fd5b82018360208201111561022a57600080fd5b8035906020019184600183028401116401000000008311171561024c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102af57600080fd5b8201836020820111156102c157600080fd5b803590602001918460018302840111640100000000831117156102e357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610581565b005b61049b6004803603606081101561035757600080fd5b810190808035906020019064010000000081111561037457600080fd5b82018360208201111561038657600080fd5b803590602001918460018302840111640100000000831117156103a857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561040b57600080fd5b82018360208201111561041d57600080fd5b8035906020019184600183028401116401000000008311171561043f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610674565b005b600080826040518082805190602001908083835b602083106104d457805182526020820191506020810190506020830392506104b1565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b806000836040518082805190602001908083835b602083106105465780518252602082019150602081019050602083039250610523565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055505050565b806000846040518082805190602001908083835b602083106105b85780518252602082019150602081019050602083039250610595565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008282540392505081905550806000836040518082805190602001908083835b6020831061062f578051825260208201915060208101905060208303925061060c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008282540192505081905550505050565b806000846040518082805190602001908083835b602083106106ab5780518252602082019150602081019050602083039250610688565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008282540392505081905550806000836040518082805190602001908083835b6020831061072257805182526020820191506020810190506020830392506106ff565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008282540192505081905550606481111561077057600080fd5b50505056fea2646970667358221220f6831a3b7f43e628f3e42e57faf4a17112d05e8f6bd05de2928b25d9e738e72264736f6c634300060a0033"

// DeployParallelOk deploys a new contract, binding an instance of ParallelOk to it.
func DeployParallelOk(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *ParallelOk, error) {
	parsed, err := abi.JSON(strings.NewReader(ParallelOkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParallelOkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &ParallelOk{ParallelOkCaller: ParallelOkCaller{contract: contract}, ParallelOkTransactor: ParallelOkTransactor{contract: contract}, ParallelOkFilterer: ParallelOkFilterer{contract: contract}}, nil
}

func AsyncDeployParallelOk(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ParallelOkABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ParallelOkBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ParallelOk is an auto generated Go binding around a Solidity contract.
type ParallelOk struct {
	ParallelOkCaller     // Read-only binding to the contract
	ParallelOkTransactor // Write-only binding to the contract
	ParallelOkFilterer   // Log filterer for contract events
}

// ParallelOkCaller is an auto generated read-only Go binding around a Solidity contract.
type ParallelOkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParallelOkTransactor is an auto generated write-only Go binding around a Solidity contract.
type ParallelOkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParallelOkFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ParallelOkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParallelOkSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ParallelOkSession struct {
	Contract     *ParallelOk       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParallelOkCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ParallelOkCallerSession struct {
	Contract *ParallelOkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ParallelOkTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ParallelOkTransactorSession struct {
	Contract     *ParallelOkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ParallelOkRaw is an auto generated low-level Go binding around a Solidity contract.
type ParallelOkRaw struct {
	Contract *ParallelOk // Generic contract binding to access the raw methods on
}

// ParallelOkCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ParallelOkCallerRaw struct {
	Contract *ParallelOkCaller // Generic read-only contract binding to access the raw methods on
}

// ParallelOkTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ParallelOkTransactorRaw struct {
	Contract *ParallelOkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParallelOk creates a new instance of ParallelOk, bound to a specific deployed contract.
func NewParallelOk(address common.Address, backend bind.ContractBackend) (*ParallelOk, error) {
	contract, err := bindParallelOk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParallelOk{ParallelOkCaller: ParallelOkCaller{contract: contract}, ParallelOkTransactor: ParallelOkTransactor{contract: contract}, ParallelOkFilterer: ParallelOkFilterer{contract: contract}}, nil
}

// NewParallelOkCaller creates a new read-only instance of ParallelOk, bound to a specific deployed contract.
func NewParallelOkCaller(address common.Address, caller bind.ContractCaller) (*ParallelOkCaller, error) {
	contract, err := bindParallelOk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParallelOkCaller{contract: contract}, nil
}

// NewParallelOkTransactor creates a new write-only instance of ParallelOk, bound to a specific deployed contract.
func NewParallelOkTransactor(address common.Address, transactor bind.ContractTransactor) (*ParallelOkTransactor, error) {
	contract, err := bindParallelOk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParallelOkTransactor{contract: contract}, nil
}

// NewParallelOkFilterer creates a new log filterer instance of ParallelOk, bound to a specific deployed contract.
func NewParallelOkFilterer(address common.Address, filterer bind.ContractFilterer) (*ParallelOkFilterer, error) {
	contract, err := bindParallelOk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParallelOkFilterer{contract: contract}, nil
}

// bindParallelOk binds a generic wrapper to an already deployed contract.
func bindParallelOk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParallelOkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParallelOk *ParallelOkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParallelOk.Contract.ParallelOkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParallelOk *ParallelOkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.ParallelOkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParallelOk *ParallelOkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.ParallelOkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParallelOk *ParallelOkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParallelOk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParallelOk *ParallelOkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParallelOk *ParallelOkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string name) constant returns(uint256)
func (_ParallelOk *ParallelOkCaller) BalanceOf(opts *bind.CallOpts, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParallelOk.contract.Call(opts, out, "balanceOf", name)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string name) constant returns(uint256)
func (_ParallelOk *ParallelOkSession) BalanceOf(name string) (*big.Int, error) {
	return _ParallelOk.Contract.BalanceOf(&_ParallelOk.CallOpts, name)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string name) constant returns(uint256)
func (_ParallelOk *ParallelOkCallerSession) BalanceOf(name string) (*big.Int, error) {
	return _ParallelOk.Contract.BalanceOf(&_ParallelOk.CallOpts, name)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string name, uint256 num) returns()
func (_ParallelOk *ParallelOkTransactor) Set(opts *bind.TransactOpts, name string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.contract.Transact(opts, "set", name, num)
}

func (_ParallelOk *ParallelOkTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.contract.AsyncTransact(opts, handler, "set", name, num)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string name, uint256 num) returns()
func (_ParallelOk *ParallelOkSession) Set(name string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.Set(&_ParallelOk.TransactOpts, name, num)
}

func (_ParallelOk *ParallelOkSession) AsyncSet(handler func(*types.Receipt, error), name string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.Contract.AsyncSet(handler, &_ParallelOk.TransactOpts, name, num)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string name, uint256 num) returns()
func (_ParallelOk *ParallelOkTransactorSession) Set(name string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.Set(&_ParallelOk.TransactOpts, name, num)
}

func (_ParallelOk *ParallelOkTransactorSession) AsyncSet(handler func(*types.Receipt, error), name string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.Contract.AsyncSet(handler, &_ParallelOk.TransactOpts, name, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 num) returns()
func (_ParallelOk *ParallelOkTransactor) Transfer(opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.contract.Transact(opts, "transfer", from, to, num)
}

func (_ParallelOk *ParallelOkTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.contract.AsyncTransact(opts, handler, "transfer", from, to, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 num) returns()
func (_ParallelOk *ParallelOkSession) Transfer(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.Transfer(&_ParallelOk.TransactOpts, from, to, num)
}

func (_ParallelOk *ParallelOkSession) AsyncTransfer(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.Contract.AsyncTransfer(handler, &_ParallelOk.TransactOpts, from, to, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 num) returns()
func (_ParallelOk *ParallelOkTransactorSession) Transfer(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.Transfer(&_ParallelOk.TransactOpts, from, to, num)
}

func (_ParallelOk *ParallelOkTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.Contract.AsyncTransfer(handler, &_ParallelOk.TransactOpts, from, to, num)
}

// TransferWithRevert is a paid mutator transaction binding the contract method 0xfad42f87.
//
// Solidity: function transferWithRevert(string from, string to, uint256 num) returns()
func (_ParallelOk *ParallelOkTransactor) TransferWithRevert(opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.contract.Transact(opts, "transferWithRevert", from, to, num)
}

func (_ParallelOk *ParallelOkTransactor) AsyncTransferWithRevert(handler func(*types.Receipt, error), opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.contract.AsyncTransact(opts, handler, "transferWithRevert", from, to, num)
}

// TransferWithRevert is a paid mutator transaction binding the contract method 0xfad42f87.
//
// Solidity: function transferWithRevert(string from, string to, uint256 num) returns()
func (_ParallelOk *ParallelOkSession) TransferWithRevert(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.TransferWithRevert(&_ParallelOk.TransactOpts, from, to, num)
}

func (_ParallelOk *ParallelOkSession) AsyncTransferWithRevert(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.Contract.AsyncTransferWithRevert(handler, &_ParallelOk.TransactOpts, from, to, num)
}

// TransferWithRevert is a paid mutator transaction binding the contract method 0xfad42f87.
//
// Solidity: function transferWithRevert(string from, string to, uint256 num) returns()
func (_ParallelOk *ParallelOkTransactorSession) TransferWithRevert(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ParallelOk.Contract.TransferWithRevert(&_ParallelOk.TransactOpts, from, to, num)
}

func (_ParallelOk *ParallelOkTransactorSession) AsyncTransferWithRevert(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _ParallelOk.Contract.AsyncTransferWithRevert(handler, &_ParallelOk.TransactOpts, from, to, num)
}

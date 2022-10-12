// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworld

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
	_ = types.BloomLookup
)

// HelloWorldABI is the input ABI used to generate the binding from.
const HelloWorldABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"v\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"v\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HelloWorldBin is the compiled bytecode used for deploying new contracts.
var HelloWorldBin = "0x60806040526040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506001908051906020019061004f9291906100ae565b5034801561005c57600080fd5b506040518060400160405280600d81526020017f48656c6c6f2c20576f726c642100000000000000000000000000000000000000815250600090805190602001906100a89291906100ae565b50610153565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ef57805160ff191683800117855561011d565b8280016001018555821561011d579182015b8281111561011c578251825591602001919060010190610101565b5b50905061012a919061012e565b5090565b61015091905b8082111561014c576000816000905550600101610134565b5090565b90565b610490806101626000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634ed3885e1461004657806354fd4d50146100bf5780636d4ce63c14610142575b600080fd5b6100bd6004803603602081101561005c57600080fd5b810190808035906020019064010000000081111561007957600080fd5b82018360208201111561008b57600080fd5b803590602001918460018302840111640100000000831117156100ad57600080fd5b90919293919293905050506101c5565b005b6100c7610275565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101075780820151818401526020810190506100ec565b50505050905090810190601f1680156101345780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61014a610313565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561018a57808201518184015260208101905061016f565b50505050905090810190601f1680156101b75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b8181600091906101d69291906103b5565b503373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff167f1cede41e194608a414a2e1d67987cf390338e67d0ff22be86dee2f3737c23d538484600160405180806020018381526020018281038252858582818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a35050565b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561030b5780601f106102e05761010080835404028352916020019161030b565b820191906000526020600020905b8154815290600101906020018083116102ee57829003601f168201915b505050505081565b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103ab5780601f10610380576101008083540402835291602001916103ab565b820191906000526020600020905b81548152906001019060200180831161038e57829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103f657803560ff1916838001178555610424565b82800160010185558215610424579182015b82811115610423578235825591602001919060010190610408565b5b5090506104319190610435565b5090565b61045791905b8082111561045357600081600090555060010161043b565b5090565b9056fea2646970667358221220dcd5e07266c2bd45adf7583c743a9c32073feecbd74ecd34e4a314f35d2a186164736f6c634300060a0033"

// DeployHelloWorld deploys a new contract, binding an instance of HelloWorld to it.
func DeployHelloWorld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *HelloWorld, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloWorldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &HelloWorld{HelloWorldCaller: HelloWorldCaller{contract: contract}, HelloWorldTransactor: HelloWorldTransactor{contract: contract}, HelloWorldFilterer: HelloWorldFilterer{contract: contract}}, nil
}

func AsyncDeployHelloWorld(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(HelloWorldBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// HelloWorld is an auto generated Go binding around a Solidity contract.
type HelloWorld struct {
	HelloWorldCaller     // Read-only binding to the contract
	HelloWorldTransactor // Write-only binding to the contract
	HelloWorldFilterer   // Log filterer for contract events
}

// HelloWorldCaller is an auto generated read-only Go binding around a Solidity contract.
type HelloWorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldTransactor is an auto generated write-only Go binding around a Solidity contract.
type HelloWorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type HelloWorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type HelloWorldSession struct {
	Contract     *HelloWorld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type HelloWorldCallerSession struct {
	Contract *HelloWorldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloWorldTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type HelloWorldTransactorSession struct {
	Contract     *HelloWorldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloWorldRaw is an auto generated low-level Go binding around a Solidity contract.
type HelloWorldRaw struct {
	Contract *HelloWorld // Generic contract binding to access the raw methods on
}

// HelloWorldCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type HelloWorldCallerRaw struct {
	Contract *HelloWorldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloWorldTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type HelloWorldTransactorRaw struct {
	Contract *HelloWorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloWorld creates a new instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorld(address common.Address, backend bind.ContractBackend) (*HelloWorld, error) {
	contract, err := bindHelloWorld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloWorld{HelloWorldCaller: HelloWorldCaller{contract: contract}, HelloWorldTransactor: HelloWorldTransactor{contract: contract}, HelloWorldFilterer: HelloWorldFilterer{contract: contract}}, nil
}

// NewHelloWorldCaller creates a new read-only instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldCaller(address common.Address, caller bind.ContractCaller) (*HelloWorldCaller, error) {
	contract, err := bindHelloWorld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldCaller{contract: contract}, nil
}

// NewHelloWorldTransactor creates a new write-only instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloWorldTransactor, error) {
	contract, err := bindHelloWorld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldTransactor{contract: contract}, nil
}

// NewHelloWorldFilterer creates a new log filterer instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloWorldFilterer, error) {
	contract, err := bindHelloWorld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloWorldFilterer{contract: contract}, nil
}

// bindHelloWorld binds a generic wrapper to an already deployed contract.
func bindHelloWorld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorld *HelloWorldRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HelloWorld.Contract.HelloWorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorld *HelloWorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorld *HelloWorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorld *HelloWorldCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HelloWorld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorld *HelloWorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorld *HelloWorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_HelloWorld *HelloWorldCaller) Get(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HelloWorld.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_HelloWorld *HelloWorldSession) Get() (string, error) {
	return _HelloWorld.Contract.Get(&_HelloWorld.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_HelloWorld *HelloWorldCallerSession) Get() (string, error) {
	return _HelloWorld.Contract.Get(&_HelloWorld.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_HelloWorld *HelloWorldCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HelloWorld.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_HelloWorld *HelloWorldSession) Version() (string, error) {
	return _HelloWorld.Contract.Version(&_HelloWorld.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_HelloWorld *HelloWorldCallerSession) Version() (string, error) {
	return _HelloWorld.Contract.Version(&_HelloWorld.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string v) returns()
func (_HelloWorld *HelloWorldTransactor) Set(opts *bind.TransactOpts, v string) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.contract.Transact(opts, "set", v)
}

func (_HelloWorld *HelloWorldTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, v string) (*types.Transaction, error) {
	return _HelloWorld.contract.AsyncTransact(opts, handler, "set", v)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string v) returns()
func (_HelloWorld *HelloWorldSession) Set(v string) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.Set(&_HelloWorld.TransactOpts, v)
}

func (_HelloWorld *HelloWorldSession) AsyncSet(handler func(*types.Receipt, error), v string) (*types.Transaction, error) {
	return _HelloWorld.Contract.AsyncSet(handler, &_HelloWorld.TransactOpts, v)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string v) returns()
func (_HelloWorld *HelloWorldTransactorSession) Set(v string) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.Set(&_HelloWorld.TransactOpts, v)
}

func (_HelloWorld *HelloWorldTransactorSession) AsyncSet(handler func(*types.Receipt, error), v string) (*types.Transaction, error) {
	return _HelloWorld.Contract.AsyncSet(handler, &_HelloWorld.TransactOpts, v)
}

// HelloWorldSetValue represents a SetValue event raised by the HelloWorld contract.
type HelloWorldSetValue struct {
	V     string
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchSetValue is a free log subscription operation binding the contract event 0x1cede41e194608a414a2e1d67987cf390338e67d0ff22be86dee2f3737c23d53.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, uint256 value)
func (_HelloWorld *HelloWorldFilterer) WatchSetValue(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address) error {
	return _HelloWorld.contract.WatchLogs(fromBlock, handler, "set", from, to)
}

func (_HelloWorld *HelloWorldFilterer) WatchAllSetValue(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _HelloWorld.contract.WatchLogs(fromBlock, handler, "set")
}

// ParseSetValue is a log parse operation binding the contract event 0x1cede41e194608a414a2e1d67987cf390338e67d0ff22be86dee2f3737c23d53.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, uint256 value)
func (_HelloWorld *HelloWorldFilterer) ParseSetValue(log types.Log) (*HelloWorldSetValue, error) {
	event := new(HelloWorldSetValue)
	if err := _HelloWorld.contract.UnpackLog(event, "set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchSetValue is a free log subscription operation binding the contract event 0x1cede41e194608a414a2e1d67987cf390338e67d0ff22be86dee2f3737c23d53.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, uint256 value)
func (_HelloWorld *HelloWorldSession) WatchSetValue(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address) error {
	return _HelloWorld.Contract.WatchSetValue(fromBlock, handler, from, to)
}

func (_HelloWorld *HelloWorldSession) WatchAllSetValue(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _HelloWorld.Contract.WatchAllSetValue(fromBlock, handler)
}

// ParseSetValue is a log parse operation binding the contract event 0x1cede41e194608a414a2e1d67987cf390338e67d0ff22be86dee2f3737c23d53.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, uint256 value)
func (_HelloWorld *HelloWorldSession) ParseSetValue(log types.Log) (*HelloWorldSetValue, error) {
	return _HelloWorld.Contract.ParseSetValue(log)
}

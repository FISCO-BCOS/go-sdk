// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// HelloWorldABI is the input ABI used to generate the binding from.
const HelloWorldABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"initValue\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"v\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setValue\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"v\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HelloWorldBin is the compiled bytecode used for deploying new contracts.
var HelloWorldBin = "0x60806040523480156200001157600080fd5b50604051620009fd380380620009fd8339818101604052810190620000379190620002ac565b80600090805190602001906200004f9291906200005f565b5060006001819055505062000362565b8280546200006d906200032c565b90600052602060002090601f016020900481019282620000915760008555620000dd565b82601f10620000ac57805160ff1916838001178555620000dd565b82800160010185558215620000dd579182015b82811115620000dc578251825591602001919060010190620000bf565b5b509050620000ec9190620000f0565b5090565b5b808211156200010b576000816000905550600101620000f1565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000178826200012d565b810181811067ffffffffffffffff821117156200019a57620001996200013e565b5b80604052505050565b6000620001af6200010f565b9050620001bd82826200016d565b919050565b600067ffffffffffffffff821115620001e057620001df6200013e565b5b620001eb826200012d565b9050602081019050919050565b60005b8381101562000218578082015181840152602081019050620001fb565b8381111562000228576000848401525b50505050565b6000620002456200023f84620001c2565b620001a3565b90508281526020810184848401111562000264576200026362000128565b5b62000271848285620001f8565b509392505050565b600082601f83011262000291576200029062000123565b5b8151620002a38482602086016200022e565b91505092915050565b600060208284031215620002c557620002c462000119565b5b600082015167ffffffffffffffff811115620002e657620002e56200011e565b5b620002f48482850162000279565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200034557607f821691505b602082108114156200035c576200035b620002fd565b5b50919050565b61068b80620003726000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634ed3885e1461004657806354fd4d50146100765780636d4ce63c14610094575b600080fd5b610060600480360381019061005b9190610387565b6100b2565b60405161006d919061046d565b60405180910390f35b61007e6101dd565b60405161008b91906104a8565b60405180910390f35b61009c6101e3565b6040516100a9919061046d565b60405180910390f35b606060008080546100c2906104f2565b80601f01602080910402602001604051908101604052809291908181526020018280546100ee906104f2565b801561013b5780601f106101105761010080835404028352916020019161013b565b820191906000526020600020905b81548152906001019060200180831161011e57829003601f168201915b50505050509050838360009190610153929190610275565b50600180546101629190610553565b6001819055503373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff167fc3bf5911f8e0476e774566ef3fa1259f04156ba5c61ea5ff35c0201390381f9686866001546040516101cb93929190610623565b60405180910390a38091505092915050565b60015481565b6060600080546101f2906104f2565b80601f016020809104026020016040519081016040528092919081815260200182805461021e906104f2565b801561026b5780601f106102405761010080835404028352916020019161026b565b820191906000526020600020905b81548152906001019060200180831161024e57829003601f168201915b5050505050905090565b828054610281906104f2565b90600052602060002090601f0160209004810192826102a357600085556102ea565b82601f106102bc57803560ff19168380011785556102ea565b828001600101855582156102ea579182015b828111156102e95782358255916020019190600101906102ce565b5b5090506102f791906102fb565b5090565b5b808211156103145760008160009055506001016102fc565b5090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261034757610346610322565b5b8235905067ffffffffffffffff81111561036457610363610327565b5b6020830191508360018202830111156103805761037f61032c565b5b9250929050565b6000806020838503121561039e5761039d610318565b5b600083013567ffffffffffffffff8111156103bc576103bb61031d565b5b6103c885828601610331565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561040e5780820151818401526020810190506103f3565b8381111561041d576000848401525b50505050565b6000601f19601f8301169050919050565b600061043f826103d4565b61044981856103df565b93506104598185602086016103f0565b61046281610423565b840191505092915050565b600060208201905081810360008301526104878184610434565b905092915050565b6000819050919050565b6104a28161048f565b82525050565b60006020820190506104bd6000830184610499565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061050a57607f821691505b6020821081141561051e5761051d6104c3565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061055e8261048f565b91506105698361048f565b9250817f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038313600083121516156105a4576105a3610524565b5b817f80000000000000000000000000000000000000000000000000000000000000000383126000831216156105dc576105db610524565b5b828201905092915050565b82818337600083830152505050565b600061060283856103df565b935061060f8385846105e7565b61061883610423565b840190509392505050565b6000604082019050818103600083015261063e8185876105f6565b905061064d6020830184610499565b94935050505056fea2646970667358221220f474bd1d28e84751caca4356bb3cca5453b846289fe3aed4ecbc8cd022fb484464736f6c634300080b0033"

// DeployHelloWorld deploys a new contract, binding an instance of HelloWorld to it.
func DeployHelloWorld(auth *bind.TransactOpts, backend bind.ContractBackend, initValue string) (common.Address, *types.Receipt, *HelloWorld, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloWorldBin), backend, initValue)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &HelloWorld{HelloWorldCaller: HelloWorldCaller{contract: contract}, HelloWorldTransactor: HelloWorldTransactor{contract: contract}, HelloWorldFilterer: HelloWorldFilterer{contract: contract}}, nil
}

func AsyncDeployHelloWorld(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, initValue string) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(HelloWorldBin), backend, initValue)
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
	Contract     *HelloWorld             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type HelloWorldCallerSession struct {
	Contract *HelloWorldCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloWorldTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type HelloWorldTransactorSession struct {
	Contract     *HelloWorldTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
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
func (_HelloWorld *HelloWorldRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.TransactWithResult(opts, result, method, params...)
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
func (_HelloWorld *HelloWorldTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.contract.TransactWithResult(opts, result, method, params...)
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
// Solidity: function version() constant returns(int256)
func (_HelloWorld *HelloWorldCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HelloWorld.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(int256)
func (_HelloWorld *HelloWorldSession) Version() (*big.Int, error) {
	return _HelloWorld.Contract.Version(&_HelloWorld.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(int256)
func (_HelloWorld *HelloWorldCallerSession) Version() (*big.Int, error) {
	return _HelloWorld.Contract.Version(&_HelloWorld.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string v) returns(string)
func (_HelloWorld *HelloWorldTransactor) Set(opts *bind.TransactOpts, v string) (string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	transaction, receipt, err := _HelloWorld.contract.TransactWithResult(opts, out, "set", v)
	return *ret0, transaction, receipt, err
}

func (_HelloWorld *HelloWorldTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, v string) (*types.Transaction, error) {
	return _HelloWorld.contract.AsyncTransact(opts, handler, "set", v)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string v) returns(string)
func (_HelloWorld *HelloWorldSession) Set(v string) (string, *types.Transaction, *types.Receipt, error) {
	return _HelloWorld.Contract.Set(&_HelloWorld.TransactOpts, v)
}

func (_HelloWorld *HelloWorldSession) AsyncSet(handler func(*types.Receipt, error), v string) (*types.Transaction, error) {
	return _HelloWorld.Contract.AsyncSet(handler, &_HelloWorld.TransactOpts, v)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string v) returns(string)
func (_HelloWorld *HelloWorldTransactorSession) Set(v string) (string, *types.Transaction, *types.Receipt, error) {
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

// WatchSetValue is a free log subscription operation binding the contract event 0xc3bf5911f8e0476e774566ef3fa1259f04156ba5c61ea5ff35c0201390381f96.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, int256 value)
func (_HelloWorld *HelloWorldFilterer) WatchSetValue(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address) (string, error) {
	return _HelloWorld.contract.WatchLogs(fromBlock, handler, "setValue", from, to)
}

func (_HelloWorld *HelloWorldFilterer) WatchAllSetValue(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _HelloWorld.contract.WatchLogs(fromBlock, handler, "setValue")
}

// ParseSetValue is a log parse operation binding the contract event 0xc3bf5911f8e0476e774566ef3fa1259f04156ba5c61ea5ff35c0201390381f96.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, int256 value)
func (_HelloWorld *HelloWorldFilterer) ParseSetValue(log types.Log) (*HelloWorldSetValue, error) {
	event := new(HelloWorldSetValue)
	if err := _HelloWorld.contract.UnpackLog(event, "setValue", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchSetValue is a free log subscription operation binding the contract event 0xc3bf5911f8e0476e774566ef3fa1259f04156ba5c61ea5ff35c0201390381f96.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, int256 value)
func (_HelloWorld *HelloWorldSession) WatchSetValue(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address) (string, error) {
	return _HelloWorld.Contract.WatchSetValue(fromBlock, handler, from, to)
}

func (_HelloWorld *HelloWorldSession) WatchAllSetValue(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _HelloWorld.Contract.WatchAllSetValue(fromBlock, handler)
}

// ParseSetValue is a log parse operation binding the contract event 0xc3bf5911f8e0476e774566ef3fa1259f04156ba5c61ea5ff35c0201390381f96.
//
// Solidity: event setValue(string v, address indexed from, address indexed to, int256 value)
func (_HelloWorld *HelloWorldSession) ParseSetValue(log types.Log) (*HelloWorldSetValue, error) {
	return _HelloWorld.Contract.ParseSetValue(log)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus

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

// ConsensusABI is the input ABI used to generate the binding from.
const ConsensusABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"addObserver\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addSealer\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"setWeight\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConsensusBin is the compiled bytecode used for deploying new contracts.
var ConsensusBin = "0x608060405234801561001057600080fd5b506103d1806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632800efc014610051578063359168561461008157806380599e4b146100b1578063ce6fa5c5146100e1575b600080fd5b61006b60048036038101906100669190610289565b610111565b60405161007891906102ee565b60405180910390f35b61009b6004803603810190610096919061033f565b610118565b6040516100a891906102ee565b60405180910390f35b6100cb60048036038101906100c69190610289565b610120565b6040516100d891906102ee565b60405180910390f35b6100fb60048036038101906100f6919061033f565b610127565b60405161010891906102ee565b60405180910390f35b6000919050565b600092915050565b6000919050565b600092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101968261014d565b810181811067ffffffffffffffff821117156101b5576101b461015e565b5b80604052505050565b60006101c861012f565b90506101d4828261018d565b919050565b600067ffffffffffffffff8211156101f4576101f361015e565b5b6101fd8261014d565b9050602081019050919050565b82818337600083830152505050565b600061022c610227846101d9565b6101be565b90508281526020810184848401111561024857610247610148565b5b61025384828561020a565b509392505050565b600082601f8301126102705761026f610143565b5b8135610280848260208601610219565b91505092915050565b60006020828403121561029f5761029e610139565b5b600082013567ffffffffffffffff8111156102bd576102bc61013e565b5b6102c98482850161025b565b91505092915050565b60008160030b9050919050565b6102e8816102d2565b82525050565b600060208201905061030360008301846102df565b92915050565b6000819050919050565b61031c81610309565b811461032757600080fd5b50565b60008135905061033981610313565b92915050565b6000806040838503121561035657610355610139565b5b600083013567ffffffffffffffff8111156103745761037361013e565b5b6103808582860161025b565b92505060206103918582860161032a565b915050925092905056fea2646970667358221220f9bf1c92bf30c11b640bf87ec39377551d33a2375d32f144e87f2bc541dc28ab64736f6c634300080b0033"

// Consensus is an auto generated Go binding around a Solidity contract.
type Consensus struct {
	ConsensusCaller     // Read-only binding to the contract
	ConsensusTransactor // Write-only binding to the contract
	ConsensusFilterer   // Log filterer for contract events
}

// ConsensusCaller is an auto generated read-only Go binding around a Solidity contract.
type ConsensusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusTransactor is an auto generated write-only Go binding around a Solidity contract.
type ConsensusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ConsensusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ConsensusSession struct {
	Contract     *Consensus        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsensusCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ConsensusCallerSession struct {
	Contract *ConsensusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConsensusTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ConsensusTransactorSession struct {
	Contract     *ConsensusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConsensusRaw is an auto generated low-level Go binding around a Solidity contract.
type ConsensusRaw struct {
	Contract *Consensus // Generic contract binding to access the raw methods on
}

// ConsensusCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ConsensusCallerRaw struct {
	Contract *ConsensusCaller // Generic read-only contract binding to access the raw methods on
}

// ConsensusTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ConsensusTransactorRaw struct {
	Contract *ConsensusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsensus creates a new instance of Consensus, bound to a specific deployed contract.
func NewConsensus(address common.Address, backend bind.ContractBackend) (*Consensus, error) {
	contract, err := bindConsensus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Consensus{ConsensusCaller: ConsensusCaller{contract: contract}, ConsensusTransactor: ConsensusTransactor{contract: contract}, ConsensusFilterer: ConsensusFilterer{contract: contract}}, nil
}

// NewConsensusCaller creates a new read-only instance of Consensus, bound to a specific deployed contract.
func NewConsensusCaller(address common.Address, caller bind.ContractCaller) (*ConsensusCaller, error) {
	contract, err := bindConsensus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusCaller{contract: contract}, nil
}

// NewConsensusTransactor creates a new write-only instance of Consensus, bound to a specific deployed contract.
func NewConsensusTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsensusTransactor, error) {
	contract, err := bindConsensus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusTransactor{contract: contract}, nil
}

// NewConsensusFilterer creates a new log filterer instance of Consensus, bound to a specific deployed contract.
func NewConsensusFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsensusFilterer, error) {
	contract, err := bindConsensus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsensusFilterer{contract: contract}, nil
}

// bindConsensus binds a generic wrapper to an already deployed contract.
func bindConsensus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsensusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consensus *ConsensusRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Consensus.Contract.ConsensusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consensus *ConsensusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.ConsensusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consensus *ConsensusRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.ConsensusTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consensus *ConsensusCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Consensus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consensus *ConsensusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consensus *ConsensusTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// AddObserver is a paid mutator transaction binding the contract method 0x2800efc0.
//
// Solidity: function addObserver(string ) returns(int32)
func (_Consensus *ConsensusTransactor) AddObserver(opts *bind.TransactOpts, arg0 string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Consensus.contract.TransactWithResult(opts, out, "addObserver", arg0)
	return *ret0, transaction, receipt, err
}

func (_Consensus *ConsensusTransactor) AsyncAddObserver(handler func(*types.Receipt, error), opts *bind.TransactOpts, arg0 string) (*types.Transaction, error) {
	return _Consensus.contract.AsyncTransact(opts, handler, "addObserver", arg0)
}

// AddObserver is a paid mutator transaction binding the contract method 0x2800efc0.
//
// Solidity: function addObserver(string ) returns(int32)
func (_Consensus *ConsensusSession) AddObserver(arg0 string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.AddObserver(&_Consensus.TransactOpts, arg0)
}

func (_Consensus *ConsensusSession) AsyncAddObserver(handler func(*types.Receipt, error), arg0 string) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncAddObserver(handler, &_Consensus.TransactOpts, arg0)
}

// AddObserver is a paid mutator transaction binding the contract method 0x2800efc0.
//
// Solidity: function addObserver(string ) returns(int32)
func (_Consensus *ConsensusTransactorSession) AddObserver(arg0 string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.AddObserver(&_Consensus.TransactOpts, arg0)
}

func (_Consensus *ConsensusTransactorSession) AsyncAddObserver(handler func(*types.Receipt, error), arg0 string) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncAddObserver(handler, &_Consensus.TransactOpts, arg0)
}

// AddSealer is a paid mutator transaction binding the contract method 0x35916856.
//
// Solidity: function addSealer(string , uint256 ) returns(int32)
func (_Consensus *ConsensusTransactor) AddSealer(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Consensus.contract.TransactWithResult(opts, out, "addSealer", arg0, arg1)
	return *ret0, transaction, receipt, err
}

func (_Consensus *ConsensusTransactor) AsyncAddSealer(handler func(*types.Receipt, error), opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Consensus.contract.AsyncTransact(opts, handler, "addSealer", arg0, arg1)
}

// AddSealer is a paid mutator transaction binding the contract method 0x35916856.
//
// Solidity: function addSealer(string , uint256 ) returns(int32)
func (_Consensus *ConsensusSession) AddSealer(arg0 string, arg1 *big.Int) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.AddSealer(&_Consensus.TransactOpts, arg0, arg1)
}

func (_Consensus *ConsensusSession) AsyncAddSealer(handler func(*types.Receipt, error), arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncAddSealer(handler, &_Consensus.TransactOpts, arg0, arg1)
}

// AddSealer is a paid mutator transaction binding the contract method 0x35916856.
//
// Solidity: function addSealer(string , uint256 ) returns(int32)
func (_Consensus *ConsensusTransactorSession) AddSealer(arg0 string, arg1 *big.Int) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.AddSealer(&_Consensus.TransactOpts, arg0, arg1)
}

func (_Consensus *ConsensusTransactorSession) AsyncAddSealer(handler func(*types.Receipt, error), arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncAddSealer(handler, &_Consensus.TransactOpts, arg0, arg1)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(string ) returns(int32)
func (_Consensus *ConsensusTransactor) Remove(opts *bind.TransactOpts, arg0 string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Consensus.contract.TransactWithResult(opts, out, "remove", arg0)
	return *ret0, transaction, receipt, err
}

func (_Consensus *ConsensusTransactor) AsyncRemove(handler func(*types.Receipt, error), opts *bind.TransactOpts, arg0 string) (*types.Transaction, error) {
	return _Consensus.contract.AsyncTransact(opts, handler, "remove", arg0)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(string ) returns(int32)
func (_Consensus *ConsensusSession) Remove(arg0 string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.Remove(&_Consensus.TransactOpts, arg0)
}

func (_Consensus *ConsensusSession) AsyncRemove(handler func(*types.Receipt, error), arg0 string) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncRemove(handler, &_Consensus.TransactOpts, arg0)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(string ) returns(int32)
func (_Consensus *ConsensusTransactorSession) Remove(arg0 string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.Remove(&_Consensus.TransactOpts, arg0)
}

func (_Consensus *ConsensusTransactorSession) AsyncRemove(handler func(*types.Receipt, error), arg0 string) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncRemove(handler, &_Consensus.TransactOpts, arg0)
}

// SetWeight is a paid mutator transaction binding the contract method 0xce6fa5c5.
//
// Solidity: function setWeight(string , uint256 ) returns(int32)
func (_Consensus *ConsensusTransactor) SetWeight(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Consensus.contract.TransactWithResult(opts, out, "setWeight", arg0, arg1)
	return *ret0, transaction, receipt, err
}

func (_Consensus *ConsensusTransactor) AsyncSetWeight(handler func(*types.Receipt, error), opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Consensus.contract.AsyncTransact(opts, handler, "setWeight", arg0, arg1)
}

// SetWeight is a paid mutator transaction binding the contract method 0xce6fa5c5.
//
// Solidity: function setWeight(string , uint256 ) returns(int32)
func (_Consensus *ConsensusSession) SetWeight(arg0 string, arg1 *big.Int) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.SetWeight(&_Consensus.TransactOpts, arg0, arg1)
}

func (_Consensus *ConsensusSession) AsyncSetWeight(handler func(*types.Receipt, error), arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncSetWeight(handler, &_Consensus.TransactOpts, arg0, arg1)
}

// SetWeight is a paid mutator transaction binding the contract method 0xce6fa5c5.
//
// Solidity: function setWeight(string , uint256 ) returns(int32)
func (_Consensus *ConsensusTransactorSession) SetWeight(arg0 string, arg1 *big.Int) (int32, *types.Transaction, *types.Receipt, error) {
	return _Consensus.Contract.SetWeight(&_Consensus.TransactOpts, arg0, arg1)
}

func (_Consensus *ConsensusTransactorSession) AsyncSetWeight(handler func(*types.Receipt, error), arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Consensus.Contract.AsyncSetWeight(handler, &_Consensus.TransactOpts, arg0, arg1)
}
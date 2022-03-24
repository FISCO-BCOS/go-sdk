// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
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

// BN256ABI is the input ABI used to generate the binding from.
const BN256ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"rounds\",\"type\":\"uint32\"},{\"name\":\"h\",\"type\":\"bytes32[2]\"},{\"name\":\"m\",\"type\":\"bytes32[4]\"},{\"name\":\"t\",\"type\":\"bytes8[2]\"},{\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"F\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"base\",\"type\":\"bytes32\"},{\"name\":\"exponent\",\"type\":\"bytes32\"},{\"name\":\"modulus\",\"type\":\"bytes32\"}],\"name\":\"BigModExp\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"},{\"name\":\"scalar\",\"type\":\"bytes32\"}],\"name\":\"Bn256ScalarMul\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"Bn256Pairing\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ax\",\"type\":\"bytes32\"},{\"name\":\"ay\",\"type\":\"bytes32\"},{\"name\":\"bx\",\"type\":\"bytes32\"},{\"name\":\"by\",\"type\":\"bytes32\"}],\"name\":\"Bn256Add\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BN256Bin is the compiled bytecode used for deploying new contracts.
var BN256Bin = "0x608060405234801561001057600080fd5b50610bce806100206000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806349f049c0146100725780635a218c3c146100af5780637c8a4a45146100ec578063aee5bc0a14610129578063ba85222314610166575b600080fd5b34801561007e57600080fd5b506100996004803603610094919081019061093f565b6101a3565b6040516100a69190610a12565b60405180910390f35b3480156100bb57600080fd5b506100d660048036036100d1919081019061084c565b6103cc565b6040516100e39190610a2d565b60405180910390f35b3480156100f857600080fd5b50610113600480360361010e919081019061084c565b610421565b6040516101209190610a12565b60405180910390f35b34801561013557600080fd5b50610150600480360361014b91908101906108fe565b6104c3565b60405161015d9190610a2d565b60405180910390f35b34801561017257600080fd5b5061018d6004803603610188919081019061089b565b610516565b60405161019a9190610a12565b60405180910390f35b6101ab6105dc565b6101b36105dc565b6060878760006002811015156101c557fe5b60200201518860016002811015156101d957fe5b60200201518860006004811015156101ed57fe5b602002015189600160048110151561020157fe5b60200201518a600260048110151561021557fe5b60200201518b600360048110151561022957fe5b60200201518b600060028110151561023d57fe5b60200201518c600160028110151561025157fe5b60200201518c604051602001808b63ffffffff1663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018a60001916600019168152602001896000191660001916815260200188600019166000191681526020018760001916600019168152602001866000191660001916815260200185600019166000191681526020018477ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff191681526008018377ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260080182151515157f01000000000000000000000000000000000000000000000000000000000000000281526001019a5050505050505050505050604051602081830303815290604052905060408260d5602084016009600019fa15156103be57600080fd5b819250505095945050505050565b600060405160208152602080820152602060408201528460608201528360808201528260a082015260208160c083600060055af180600081146104125782519350610417565b600080fd5b5050509392505050565b6104296105dc565b6104316105fe565b8481600060038110151561044157fe5b602002019060001916908160001916815250508381600160038110151561046457fe5b602002019060001916908160001916815250508281600260038110151561048757fe5b60200201906000191690816000191681525050604082606083600060075af180600081146104b4576104b9565b600080fd5b5050509392505050565b60008082519050600060c0828115156104d857fe5b061415156104e557600080fd5b6040516020818360208701600060085af18060008114610508578251945061050d565b600080fd5b50505050919050565b61051e6105dc565b610526610621565b8581600060048110151561053657fe5b602002019060001916908160001916815250508481600160048110151561055957fe5b602002019060001916908160001916815250508381600260048110151561057c57fe5b602002019060001916908160001916815250508281600360048110151561059f57fe5b60200201906000191690816000191681525050604082608083600060065af180600081146105cc576105d1565b600080fd5b505050949350505050565b6040805190810160405280600290602082028038833980820191505090505090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b600082601f830112151561065757600080fd5b600261066a61066582610a75565b610a48565b9150818385602084028201111561068057600080fd5b60005b838110156106b0578161069688826107ba565b845260208401935060208301925050600181019050610683565b5050505092915050565b600082601f83011215156106cd57600080fd5b60046106e06106db82610a97565b610a48565b915081838560208402820111156106f657600080fd5b60005b83811015610726578161070c88826107ba565b8452602084019350602083019250506001810190506106f9565b5050505092915050565b600082601f830112151561074357600080fd5b600261075661075182610ab9565b610a48565b9150818385602084028201111561076c57600080fd5b60005b8381101561079c578161078288826107ce565b84526020840193506020830192505060018101905061076f565b5050505092915050565b60006107b28235610b33565b905092915050565b60006107c68235610b3f565b905092915050565b60006107da8235610b49565b905092915050565b600082601f83011215156107f557600080fd5b813561080861080382610adb565b610a48565b9150808252602083016020830185838301111561082457600080fd5b61082f838284610b85565b50505092915050565b60006108448235610b75565b905092915050565b60008060006060848603121561086157600080fd5b600061086f868287016107ba565b9350506020610880868287016107ba565b9250506040610891868287016107ba565b9150509250925092565b600080600080608085870312156108b157600080fd5b60006108bf878288016107ba565b94505060206108d0878288016107ba565b93505060406108e1878288016107ba565b92505060606108f2878288016107ba565b91505092959194509250565b60006020828403121561091057600080fd5b600082013567ffffffffffffffff81111561092a57600080fd5b610936848285016107e2565b91505092915050565b6000806000806000610140868803121561095857600080fd5b600061096688828901610838565b955050602061097788828901610644565b9450506060610988888289016106ba565b93505060e061099988828901610730565b9250506101206109ab888289016107a6565b9150509295509295909350565b6109c181610b11565b6109ca82610b07565b60005b828110156109fc576109e0858351610a03565b6109e982610b1c565b91506020850194506001810190506109cd565b5050505050565b610a0c81610b29565b82525050565b6000604082019050610a2760008301846109b8565b92915050565b6000602082019050610a426000830184610a03565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610a6b57600080fd5b8060405250919050565b600067ffffffffffffffff821115610a8c57600080fd5b602082029050919050565b600067ffffffffffffffff821115610aae57600080fd5b602082029050919050565b600067ffffffffffffffff821115610ad057600080fd5b602082029050919050565b600067ffffffffffffffff821115610af257600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b600060029050919050565b6000602082019050919050565b6000819050919050565b60008115159050919050565b6000819050919050565b60007fffffffffffffffff00000000000000000000000000000000000000000000000082169050919050565b600063ffffffff82169050919050565b828183376000838301525050505600a265627a7a72305820a73cfba29580aec571ccc9b027ee05391712e0e9bb0e22b314955b550406139b6c6578706572696d656e74616cf50037"

// DeployBN256 deploys a new contract, binding an instance of BN256 to it.
func DeployBN256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BN256, error) {
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BN256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN256{BN256Caller: BN256Caller{contract: contract}, BN256Transactor: BN256Transactor{contract: contract}, BN256Filterer: BN256Filterer{contract: contract}}, nil
}

// BN256 is an auto generated Go binding around a Solidity contract.
type BN256 struct {
	BN256Caller     // Read-only binding to the contract
	BN256Transactor // Write-only binding to the contract
	BN256Filterer   // Log filterer for contract events
}

// BN256Caller is an auto generated read-only Go binding around a Solidity contract.
type BN256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Transactor is an auto generated write-only Go binding around a Solidity contract.
type BN256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Filterer is an auto generated log filtering Go binding around a Solidity contract events.
type BN256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Session is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type BN256Session struct {
	Contract     *BN256            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256CallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type BN256CallerSession struct {
	Contract *BN256Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BN256TransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type BN256TransactorSession struct {
	Contract     *BN256Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256Raw is an auto generated low-level Go binding around a Solidity contract.
type BN256Raw struct {
	Contract *BN256 // Generic contract binding to access the raw methods on
}

// BN256CallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type BN256CallerRaw struct {
	Contract *BN256Caller // Generic read-only contract binding to access the raw methods on
}

// BN256TransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type BN256TransactorRaw struct {
	Contract *BN256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBN256 creates a new instance of BN256, bound to a specific deployed contract.
func NewBN256(address common.Address, backend bind.ContractBackend) (*BN256, error) {
	contract, err := bindBN256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN256{BN256Caller: BN256Caller{contract: contract}, BN256Transactor: BN256Transactor{contract: contract}, BN256Filterer: BN256Filterer{contract: contract}}, nil
}

// NewBN256Caller creates a new read-only instance of BN256, bound to a specific deployed contract.
func NewBN256Caller(address common.Address, caller bind.ContractCaller) (*BN256Caller, error) {
	contract, err := bindBN256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN256Caller{contract: contract}, nil
}

// NewBN256Transactor creates a new write-only instance of BN256, bound to a specific deployed contract.
func NewBN256Transactor(address common.Address, transactor bind.ContractTransactor) (*BN256Transactor, error) {
	contract, err := bindBN256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN256Transactor{contract: contract}, nil
}

// NewBN256Filterer creates a new log filterer instance of BN256, bound to a specific deployed contract.
func NewBN256Filterer(address common.Address, filterer bind.ContractFilterer) (*BN256Filterer, error) {
	contract, err := bindBN256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN256Filterer{contract: contract}, nil
}

// bindBN256 binds a generic wrapper to an already deployed contract.
func bindBN256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN256 *BN256Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BN256.Contract.BN256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN256 *BN256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.BN256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.BN256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN256 *BN256CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BN256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN256 *BN256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.contract.Transact(opts, method, params...)
}

// F is a free data retrieval call binding the contract method 0x49f049c0.
//
// Solidity: function F(uint32 rounds, bytes32[2] h, bytes32[4] m, bytes8[2] t, bool f) constant returns(bytes32[2])
func (_BN256 *BN256Caller) F(opts *bind.CallOpts, rounds uint32, h [2][32]byte, m [4][32]byte, t [2][8]byte, f bool) ([2][32]byte, error) {
	var (
		ret0 = new([2][32]byte)
	)
	out := ret0
	err := _BN256.contract.Call(opts, out, "F", rounds, h, m, t, f)
	return *ret0, err
}

// F is a free data retrieval call binding the contract method 0x49f049c0.
//
// Solidity: function F(uint32 rounds, bytes32[2] h, bytes32[4] m, bytes8[2] t, bool f) constant returns(bytes32[2])
func (_BN256 *BN256Session) F(rounds uint32, h [2][32]byte, m [4][32]byte, t [2][8]byte, f bool) ([2][32]byte, error) {
	return _BN256.Contract.F(&_BN256.CallOpts, rounds, h, m, t, f)
}

// F is a free data retrieval call binding the contract method 0x49f049c0.
//
// Solidity: function F(uint32 rounds, bytes32[2] h, bytes32[4] m, bytes8[2] t, bool f) constant returns(bytes32[2])
func (_BN256 *BN256CallerSession) F(rounds uint32, h [2][32]byte, m [4][32]byte, t [2][8]byte, f bool) ([2][32]byte, error) {
	return _BN256.Contract.F(&_BN256.CallOpts, rounds, h, m, t, f)
}

// BigModExp is a paid mutator transaction binding the contract method 0x5a218c3c.
//
// Solidity: function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus) returns(bytes32 result)
func (_BN256 *BN256Transactor) BigModExp(opts *bind.TransactOpts, base [32]byte, exponent [32]byte, modulus [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.contract.Transact(opts, "BigModExp", base, exponent, modulus)
}

// BigModExp is a paid mutator transaction binding the contract method 0x5a218c3c.
//
// Solidity: function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus) returns(bytes32 result)
func (_BN256 *BN256Session) BigModExp(base [32]byte, exponent [32]byte, modulus [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.BigModExp(&_BN256.TransactOpts, base, exponent, modulus)
}

// BigModExp is a paid mutator transaction binding the contract method 0x5a218c3c.
//
// Solidity: function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus) returns(bytes32 result)
func (_BN256 *BN256TransactorSession) BigModExp(base [32]byte, exponent [32]byte, modulus [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.BigModExp(&_BN256.TransactOpts, base, exponent, modulus)
}

// Bn256Add is a paid mutator transaction binding the contract method 0xba852223.
//
// Solidity: function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by) returns(bytes32[2] result)
func (_BN256 *BN256Transactor) Bn256Add(opts *bind.TransactOpts, ax [32]byte, ay [32]byte, bx [32]byte, by [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.contract.Transact(opts, "Bn256Add", ax, ay, bx, by)
}

// Bn256Add is a paid mutator transaction binding the contract method 0xba852223.
//
// Solidity: function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by) returns(bytes32[2] result)
func (_BN256 *BN256Session) Bn256Add(ax [32]byte, ay [32]byte, bx [32]byte, by [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.Bn256Add(&_BN256.TransactOpts, ax, ay, bx, by)
}

// Bn256Add is a paid mutator transaction binding the contract method 0xba852223.
//
// Solidity: function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by) returns(bytes32[2] result)
func (_BN256 *BN256TransactorSession) Bn256Add(ax [32]byte, ay [32]byte, bx [32]byte, by [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.Bn256Add(&_BN256.TransactOpts, ax, ay, bx, by)
}

// Bn256Pairing is a paid mutator transaction binding the contract method 0xaee5bc0a.
//
// Solidity: function Bn256Pairing(bytes input) returns(bytes32 result)
func (_BN256 *BN256Transactor) Bn256Pairing(opts *bind.TransactOpts, input []byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.contract.Transact(opts, "Bn256Pairing", input)
}

// Bn256Pairing is a paid mutator transaction binding the contract method 0xaee5bc0a.
//
// Solidity: function Bn256Pairing(bytes input) returns(bytes32 result)
func (_BN256 *BN256Session) Bn256Pairing(input []byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.Bn256Pairing(&_BN256.TransactOpts, input)
}

// Bn256Pairing is a paid mutator transaction binding the contract method 0xaee5bc0a.
//
// Solidity: function Bn256Pairing(bytes input) returns(bytes32 result)
func (_BN256 *BN256TransactorSession) Bn256Pairing(input []byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.Bn256Pairing(&_BN256.TransactOpts, input)
}

// Bn256ScalarMul is a paid mutator transaction binding the contract method 0x7c8a4a45.
//
// Solidity: function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar) returns(bytes32[2] result)
func (_BN256 *BN256Transactor) Bn256ScalarMul(opts *bind.TransactOpts, x [32]byte, y [32]byte, scalar [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.contract.Transact(opts, "Bn256ScalarMul", x, y, scalar)
}

// Bn256ScalarMul is a paid mutator transaction binding the contract method 0x7c8a4a45.
//
// Solidity: function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar) returns(bytes32[2] result)
func (_BN256 *BN256Session) Bn256ScalarMul(x [32]byte, y [32]byte, scalar [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.Bn256ScalarMul(&_BN256.TransactOpts, x, y, scalar)
}

// Bn256ScalarMul is a paid mutator transaction binding the contract method 0x7c8a4a45.
//
// Solidity: function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar) returns(bytes32[2] result)
func (_BN256 *BN256TransactorSession) Bn256ScalarMul(x [32]byte, y [32]byte, scalar [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _BN256.Contract.Bn256ScalarMul(&_BN256.TransactOpts, x, y, scalar)
}

func main() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("parseConfig failed, err: %v", err)
	}
	client, err := client.Dial(&configs[0])
	if err != nil {
		fmt.Printf("Dial Client failed, err:%v", err)
		return
	}
	address, tx, instance, err := DeployBN256(client.GetTransactOpts(), client)
	if err != nil {
		fmt.Printf("Deploy failed, err:%v", err)
		return
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())

	bn256Precompiled := &BN256Session{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// g1x := new(bn256.G1)
	_, g1x, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		fmt.Printf("bn256.RandomG1 failed of :%v", err)
		return
	}
	k, g1y, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		fmt.Printf("bn256.RandomG1 failed of :%v", err)
		return
	}
	g1xBytes := g1x.Marshal()
	g1yBytes := g1y.Marshal()
	var ax, ay, bx, by [32]byte
	copy(ax[:], g1xBytes[:32])
	copy(ay[:], g1xBytes[32:])
	copy(bx[:], g1yBytes[:32])
	copy(by[:], g1yBytes[32:])
	tx, receipt, err := bn256Precompiled.Bn256Add(ax, ay, bx, by)
	if err != nil {
		fmt.Printf("client.WaitMined failed of :%v", err)
		return
	}
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		fmt.Printf("abi.JSON(strings.NewReader(BN256ABI)) failed of :%v", err)
		return
	}
	var ret [2][32]byte
	err = parsed.Unpack(&ret, "Bn256Add", common.FromHex(receipt.Output))
	if err != nil {
		fmt.Printf("Unpack Bn256Add failed of :%v", err)
		return
	}
	p := new(bn256.G1)
	p.Add(g1x, g1y)
	precompiledResult := make([]byte, 64)
	copy(precompiledResult[:], ret[0][:])
	copy(precompiledResult[32:], ret[1][:])
	retBytes := p.Marshal()
	if bytes.Compare(retBytes, precompiledResult) != 0 {
		fmt.Printf("precompiled Bn256Add not equal local add\n")
		fmt.Printf("local=%v\nBn256Add=%v\n", retBytes, precompiledResult)
		fmt.Printf("a=%x\nb=%x", g1xBytes, g1yBytes)
		return
	}
	fmt.Printf("Bn256Add success\n")

	var k32 [32]byte
	copy(k32[32-len(k.Bytes()):], k.Bytes())
	_, g1z, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		fmt.Printf("bn256.RandomG1 failed of :%v", err)
		return
	}
	g1zBytes := g1z.Marshal()
	var zx, zy [32]byte
	copy(zx[:], g1zBytes[:32])
	copy(zy[:], g1zBytes[32:])
	tx, receipt, err = bn256Precompiled.Bn256ScalarMul(zx, zy, k32)
	var ret2 [2][32]byte

	err = parsed.Unpack(&ret2, "Bn256ScalarMul", common.FromHex(receipt.Output))
	if err != nil {
		fmt.Printf("Unpack Bn256ScalarMul failed of :%v", err)
		return
	}
	precompiledResult2 := make([]byte, 64)

	copy(precompiledResult2[:], ret2[0][:])
	copy(precompiledResult2[32:], ret2[1][:])
	p = new(bn256.G1)
	p = p.ScalarMult(g1z, k)
	retBytes = p.Marshal()
	if bytes.Compare(retBytes, precompiledResult2) != 0 {
		fmt.Printf("precompiled Bn256ScalarMul not equal\n")
		fmt.Printf("local=%v\nprecompiled=%v\n k=%x\n p=%x\n", retBytes, precompiledResult, k.Bytes(), g1zBytes)
		return
	}
	fmt.Printf("Bn256ScalarMul success\n")

	var hexInput = "0x1c76476f4def4bb94541d57ebba1193381ffa7aa76ada664dd31c16024c43f593034dd2920f673e204fee2811c678745fc819b55d3e9d294e45c9b03a76aef41209dd15ebff5d46c4bd888e51a93cf99a7329636c63514396b4a452003a35bf704bf11ca01483bfa8b34b43561848d28905960114c8ac04049af4b6315a416782bb8324af6cfc93537a2ad1a445cfd0ca2a71acd7ac41fadbf933c2a51be344d120a2a4cf30c1bf9845f20c6fe39e07ea2cce61f0c9bb048165fe5e4de877550111e129f1cf1097710d41c4ac70fcdfa5ba2023c6ff1cbeac322de49d1b6df7c2032c61a830e3c17286de9462bf242fca2883585b93870a73853face6a6bf411198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c21800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa"
	tx, receipt, err = bn256Precompiled.Bn256Pairing(common.FromHex(hexInput))
	var ret3 [32]byte
	err = parsed.Unpack(&ret3, "Bn256Pairing", common.FromHex(receipt.Output))
	if err != nil {
		fmt.Printf("Unpack Bn256Pairing failed of :%v", err)
		return
	}
	if ret3[31] != 1 {
		fmt.Printf("precompiled Bn256Pairing not equal 1\n")
		fmt.Printf("result=%v\n", ret3)
		return
	}
	fmt.Printf("Bn256Pairing success\n")

	base := generateBigInt()
	exponent := generateBigInt()
	modulus := generateBigInt()
	var b, e, m [32]byte
	copy(b[32-len(base.Bytes()):], base.Bytes())
	copy(e[32-len(exponent.Bytes()):], exponent.Bytes())
	copy(m[32-len(modulus.Bytes()):], modulus.Bytes())
	tx, receipt, err = bn256Precompiled.BigModExp(b, e, m)
	var ret4, local4 [32]byte
	err = parsed.Unpack(&ret4, "BigModExp", common.FromHex(receipt.Output))
	if err != nil {
		fmt.Printf("Unpack BigModExp failed of :%v", err)
		return
	}
	r := new(big.Int).Exp(base, exponent, modulus)
	copy(local4[32-len(r.Bytes()):], r.Bytes())
	if bytes.Compare(ret4[:], local4[:]) != 0 {
		fmt.Printf("precompiled BigModExp not equal\n")
		fmt.Printf("local=%x\nBigModExp=%x\n", ret4, r.Bytes())
		fmt.Printf("b=%x\ne=%x\nm=%x\n", base, exponent, modulus)
		os.Exit(1)
	}
	fmt.Printf("BigModExp success\n")
}

func generateBigInt() *big.Int {
	randBytes := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, randBytes)
	if err != nil {
		fmt.Printf("random bigint failed of :%v", err)
		return nil

	}
	return new(big.Int).SetBytes(randBytes)
}

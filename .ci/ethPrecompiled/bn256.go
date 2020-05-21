// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	"github.com/bitherhq/go-bither/crypto/bn256"
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
	_ = event.NewSubscription
)

// BN256ABI is the input ABI used to generate the binding from.
const BN256ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"base\",\"type\":\"bytes32\"},{\"name\":\"exponent\",\"type\":\"bytes32\"},{\"name\":\"modulus\",\"type\":\"bytes32\"}],\"name\":\"BigModExp\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"},{\"name\":\"scalar\",\"type\":\"bytes32\"}],\"name\":\"Bn256ScalarMul\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"Bn256Pairing\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ax\",\"type\":\"bytes32\"},{\"name\":\"ay\",\"type\":\"bytes32\"},{\"name\":\"bx\",\"type\":\"bytes32\"},{\"name\":\"by\",\"type\":\"bytes32\"}],\"name\":\"Bn256Add\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BN256Bin is the compiled bytecode used for deploying new contracts.
var BN256Bin = "0x608060405234801561001057600080fd5b50610698806100206000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631aa5d42f146100675780633adf65f5146100a457806390ea88a1146100e1578063fda5d6261461011e575b600080fd5b34801561007357600080fd5b5061008e6004803603610089919081019061043d565b61015b565b60405161009b91906105a5565b60405180910390f35b3480156100b057600080fd5b506100cb60048036036100c6919081019061043d565b6101b0565b6040516100d8919061058a565b60405180910390f35b3480156100ed57600080fd5b50610108600480360361010391908101906104ef565b610252565b60405161011591906105a5565b60405180910390f35b34801561012a57600080fd5b506101456004803603610140919081019061048c565b6102a5565b604051610152919061058a565b60405180910390f35b600060405160208152602080820152602060408201528460608201528360808201528260a082015260208160c083600060055af180600081146101a157825193506101a6565b600080fd5b5050509392505050565b6101b861036b565b6101c061038d565b848160006003811015156101d057fe5b60200201906000191690816000191681525050838160016003811015156101f357fe5b602002019060001916908160001916815250508281600260038110151561021657fe5b60200201906000191690816000191681525050604082606083600060075af1806000811461024357610248565b600080fd5b5050509392505050565b60008082519050600060c08281151561026757fe5b0614151561027457600080fd5b6040516020818360208701600060085af18060008114610297578251945061029c565b600080fd5b50505050919050565b6102ad61036b565b6102b56103b0565b858160006004811015156102c557fe5b60200201906000191690816000191681525050848160016004811015156102e857fe5b602002019060001916908160001916815250508381600260048110151561030b57fe5b602002019060001916908160001916815250508281600360048110151561032e57fe5b60200201906000191690816000191681525050604082608083600060065af1806000811461035b57610360565b600080fd5b505050949350505050565b6040805190810160405280600290602082028038833980820191505090505090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b60006103df8235610645565b905092915050565b600082601f83011215156103fa57600080fd5b813561040d610408826105ed565b6105c0565b9150808252602083016020830185838301111561042957600080fd5b61043483828461064f565b50505092915050565b60008060006060848603121561045257600080fd5b6000610460868287016103d3565b9350506020610471868287016103d3565b9250506040610482868287016103d3565b9150509250925092565b600080600080608085870312156104a257600080fd5b60006104b0878288016103d3565b94505060206104c1878288016103d3565b93505060406104d2878288016103d3565b92505060606104e3878288016103d3565b91505092959194509250565b60006020828403121561050157600080fd5b600082013567ffffffffffffffff81111561051b57600080fd5b610527848285016103e7565b91505092915050565b61053981610623565b61054282610619565b60005b828110156105745761055885835161057b565b6105618261062e565b9150602085019450600181019050610545565b5050505050565b6105848161063b565b82525050565b600060408201905061059f6000830184610530565b92915050565b60006020820190506105ba600083018461057b565b92915050565b6000604051905081810181811067ffffffffffffffff821117156105e357600080fd5b8060405250919050565b600067ffffffffffffffff82111561060457600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b600060029050919050565b6000602082019050919050565b6000819050919050565b6000819050919050565b828183376000838301525050505600a265627a7a723058206a87baa154a9cbf59cd93f7c27e8ae81374b6b0500f04320e287dd65fc0ee2746c6578706572696d656e74616cf50037"

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

// BN256 is an auto generated Go binding around an contract.
type BN256 struct {
	BN256Caller     // Read-only binding to the contract
	BN256Transactor // Write-only binding to the contract
	BN256Filterer   // Log filterer for contract events
}

// BN256Caller is an auto generated read-only Go binding around an contract.
type BN256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Transactor is an auto generated write-only Go binding around an contract.
type BN256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Filterer is an auto generated log filtering Go binding around an contract events.
type BN256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Session is an auto generated Go binding around an contract,
// with pre-set call and transact options.
type BN256Session struct {
	Contract     *BN256            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256CallerSession is an auto generated read-only Go binding around an contract,
// with pre-set call options.
type BN256CallerSession struct {
	Contract *BN256Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BN256TransactorSession is an auto generated write-only Go binding around an contract,
// with pre-set transact options.
type BN256TransactorSession struct {
	Contract     *BN256Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256Raw is an auto generated low-level Go binding around an contract.
type BN256Raw struct {
	Contract *BN256 // Generic contract binding to access the raw methods on
}

// BN256CallerRaw is an auto generated low-level read-only Go binding around an contract.
type BN256CallerRaw struct {
	Contract *BN256Caller // Generic read-only contract binding to access the raw methods on
}

// BN256TransactorRaw is an auto generated low-level write-only Go binding around an contract.
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
func (_BN256 *BN256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN256.Contract.BN256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
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
func (_BN256 *BN256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN256.Contract.contract.Transact(opts, method, params...)
}

// BigModExp is a paid mutator transaction binding the contract method 0x1aa5d42f.
//
// Solidity: function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus) returns(bytes32 result)
func (_BN256 *BN256Transactor) BigModExp(opts *bind.TransactOpts, base [32]byte, exponent [32]byte, modulus [32]byte) (*types.Transaction, error) {
	return _BN256.contract.Transact(opts, "BigModExp", base, exponent, modulus)
}

// BigModExp is a paid mutator transaction binding the contract method 0x1aa5d42f.
//
// Solidity: function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus) returns(bytes32 result)
func (_BN256 *BN256Session) BigModExp(base [32]byte, exponent [32]byte, modulus [32]byte) (*types.Transaction, error) {
	return _BN256.Contract.BigModExp(&_BN256.TransactOpts, base, exponent, modulus)
}

// BigModExp is a paid mutator transaction binding the contract method 0x1aa5d42f.
//
// Solidity: function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus) returns(bytes32 result)
func (_BN256 *BN256TransactorSession) BigModExp(base [32]byte, exponent [32]byte, modulus [32]byte) (*types.Transaction, error) {
	return _BN256.Contract.BigModExp(&_BN256.TransactOpts, base, exponent, modulus)
}

// Bn256Add is a paid mutator transaction binding the contract method 0xfda5d626.
//
// Solidity: function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by) returns(bytes32[2] result)
func (_BN256 *BN256Transactor) Bn256Add(opts *bind.TransactOpts, ax [32]byte, ay [32]byte, bx [32]byte, by [32]byte) (*types.Transaction, error) {
	return _BN256.contract.Transact(opts, "Bn256Add", ax, ay, bx, by)
}

// Bn256Add is a paid mutator transaction binding the contract method 0xfda5d626.
//
// Solidity: function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by) returns(bytes32[2] result)
func (_BN256 *BN256Session) Bn256Add(ax [32]byte, ay [32]byte, bx [32]byte, by [32]byte) (*types.Transaction, error) {
	return _BN256.Contract.Bn256Add(&_BN256.TransactOpts, ax, ay, bx, by)
}

// Bn256Add is a paid mutator transaction binding the contract method 0xfda5d626.
//
// Solidity: function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by) returns(bytes32[2] result)
func (_BN256 *BN256TransactorSession) Bn256Add(ax [32]byte, ay [32]byte, bx [32]byte, by [32]byte) (*types.Transaction, error) {
	return _BN256.Contract.Bn256Add(&_BN256.TransactOpts, ax, ay, bx, by)
}

// Bn256Pairing is a paid mutator transaction binding the contract method 0x90ea88a1.
//
// Solidity: function Bn256Pairing(bytes input) returns(bytes32 result)
func (_BN256 *BN256Transactor) Bn256Pairing(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _BN256.contract.Transact(opts, "Bn256Pairing", input)
}

// Bn256Pairing is a paid mutator transaction binding the contract method 0x90ea88a1.
//
// Solidity: function Bn256Pairing(bytes input) returns(bytes32 result)
func (_BN256 *BN256Session) Bn256Pairing(input []byte) (*types.Transaction, error) {
	return _BN256.Contract.Bn256Pairing(&_BN256.TransactOpts, input)
}

// Bn256Pairing is a paid mutator transaction binding the contract method 0x90ea88a1.
//
// Solidity: function Bn256Pairing(bytes input) returns(bytes32 result)
func (_BN256 *BN256TransactorSession) Bn256Pairing(input []byte) (*types.Transaction, error) {
	return _BN256.Contract.Bn256Pairing(&_BN256.TransactOpts, input)
}

// Bn256ScalarMul is a paid mutator transaction binding the contract method 0x3adf65f5.
//
// Solidity: function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar) returns(bytes32[2] result)
func (_BN256 *BN256Transactor) Bn256ScalarMul(opts *bind.TransactOpts, x [32]byte, y [32]byte, scalar [32]byte) (*types.Transaction, error) {
	return _BN256.contract.Transact(opts, "Bn256ScalarMul", x, y, scalar)
}

// Bn256ScalarMul is a paid mutator transaction binding the contract method 0x3adf65f5.
//
// Solidity: function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar) returns(bytes32[2] result)
func (_BN256 *BN256Session) Bn256ScalarMul(x [32]byte, y [32]byte, scalar [32]byte) (*types.Transaction, error) {
	return _BN256.Contract.Bn256ScalarMul(&_BN256.TransactOpts, x, y, scalar)
}

// Bn256ScalarMul is a paid mutator transaction binding the contract method 0x3adf65f5.
//
// Solidity: function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar) returns(bytes32[2] result)
func (_BN256 *BN256TransactorSession) Bn256ScalarMul(x [32]byte, y [32]byte, scalar [32]byte) (*types.Transaction, error) {
	return _BN256.Contract.Bn256ScalarMul(&_BN256.TransactOpts, x, y, scalar)
}

func main() {
	configs := conf.ParseConfig("config.toml")
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
	tx, err = bn256Precompiled.Bn256Add(ax, ay, bx, by)
	if err != nil {
		fmt.Printf("bn256Precompiled.Bn256Add failed of :%v", err)
		return
	}
	receipt, err := client.WaitMined(tx)
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
	tx, err = bn256Precompiled.Bn256ScalarMul(zx, zy, k32)
	receipt, err = client.WaitMined(tx)
	if err != nil {
		fmt.Printf("Bn256ScalarMul WaitMined failed of :%v", err)
		return
	}
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
		fmt.Printf("local=%v\nprecompiled=%v\n k=%x\n p=%x", retBytes, precompiledResult, k.Bytes(), g1zBytes)
		return
	}
	fmt.Printf("Bn256ScalarMul success\n")

	var hexInput = "0x1c76476f4def4bb94541d57ebba1193381ffa7aa76ada664dd31c16024c43f593034dd2920f673e204fee2811c678745fc819b55d3e9d294e45c9b03a76aef41209dd15ebff5d46c4bd888e51a93cf99a7329636c63514396b4a452003a35bf704bf11ca01483bfa8b34b43561848d28905960114c8ac04049af4b6315a416782bb8324af6cfc93537a2ad1a445cfd0ca2a71acd7ac41fadbf933c2a51be344d120a2a4cf30c1bf9845f20c6fe39e07ea2cce61f0c9bb048165fe5e4de877550111e129f1cf1097710d41c4ac70fcdfa5ba2023c6ff1cbeac322de49d1b6df7c2032c61a830e3c17286de9462bf242fca2883585b93870a73853face6a6bf411198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c21800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa"
	tx, err = bn256Precompiled.Bn256Pairing(common.FromHex(hexInput))
	receipt, err = client.WaitMined(tx)
	if err != nil {
		fmt.Printf("Bn256Pairing WaitMined failed of :%v", err)
		return
	}
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
	tx, err = bn256Precompiled.BigModExp(b, e, m)
	receipt, err = client.WaitMined(tx)
	if err != nil {
		fmt.Printf("BigModExp WaitMined failed of :%v", err)
		return
	}
	var ret4 [32]byte
	err = parsed.Unpack(&ret4, "BigModExp", common.FromHex(receipt.Output))
	if err != nil {
		fmt.Printf("Unpack BigModExp failed of :%v", err)
		return
	}
	r := new(big.Int).Exp(base, exponent, modulus)
	if bytes.Compare(ret4[:], r.Bytes()) != 0 {
		fmt.Printf("precompiled BigModExp not equal\n")
		fmt.Printf("local=%v\nBigModExp=%v\n", retBytes, precompiledResult)
		return
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

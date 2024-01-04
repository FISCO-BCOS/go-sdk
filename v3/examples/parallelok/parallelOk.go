// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"fmt"
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

// TransferABI is the input ABI used to generate the binding from.
const TransferABI = "[{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"selector\":[904814471,3449012829],\"stateMutability\":\"view\",\"type\":\"function\"},{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"selector\":[2319641577,4076138093],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]},{\"kind\":3,\"slot\":0,\"value\":[1]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"selector\":[2608902224,1630350335],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"conflictFields\":[{\"kind\":3,\"slot\":0,\"value\":[0]},{\"kind\":3,\"slot\":0,\"value\":[1]}],\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transferWithRevert\",\"outputs\":[],\"selector\":[4208209799,2876358409],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransferBin is the compiled bytecode used for deploying new contracts.
var TransferBin = "0x608060405234801561001057600080fd5b50610679806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806335ee5f87146100515780638a42ebe9146100815780639b80b0501461009d578063fad42f87146100b9575b600080fd5b61006b60048036038101906100669190610369565b6100d5565b60405161007891906103cb565b60405180910390f35b61009b60048036038101906100969190610412565b6100fc565b005b6100b760048036038101906100b2919061046e565b610123565b005b6100d360048036038101906100ce919061046e565b610192565b005b600080826040516100e69190610573565b9081526020016040518091039020549050919050565b8060008360405161010d9190610573565b9081526020016040518091039020819055505050565b806000846040516101349190610573565b9081526020016040518091039020600082825461015191906105b9565b92505081905550806000836040516101699190610573565b9081526020016040518091039020600082825461018691906105ed565b92505081905550505050565b806000846040516101a39190610573565b908152602001604051809103902060008282546101c091906105b9565b92505081905550806000836040516101d89190610573565b908152602001604051809103902060008282546101f591906105ed565b92505081905550606481111561020a57600080fd5b505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102768261022d565b810181811067ffffffffffffffff821117156102955761029461023e565b5b80604052505050565b60006102a861020f565b90506102b4828261026d565b919050565b600067ffffffffffffffff8211156102d4576102d361023e565b5b6102dd8261022d565b9050602081019050919050565b82818337600083830152505050565b600061030c610307846102b9565b61029e565b90508281526020810184848401111561032857610327610228565b5b6103338482856102ea565b509392505050565b600082601f8301126103505761034f610223565b5b81356103608482602086016102f9565b91505092915050565b60006020828403121561037f5761037e610219565b5b600082013567ffffffffffffffff81111561039d5761039c61021e565b5b6103a98482850161033b565b91505092915050565b6000819050919050565b6103c5816103b2565b82525050565b60006020820190506103e060008301846103bc565b92915050565b6103ef816103b2565b81146103fa57600080fd5b50565b60008135905061040c816103e6565b92915050565b6000806040838503121561042957610428610219565b5b600083013567ffffffffffffffff8111156104475761044661021e565b5b6104538582860161033b565b9250506020610464858286016103fd565b9150509250929050565b60008060006060848603121561048757610486610219565b5b600084013567ffffffffffffffff8111156104a5576104a461021e565b5b6104b18682870161033b565b935050602084013567ffffffffffffffff8111156104d2576104d161021e565b5b6104de8682870161033b565b92505060406104ef868287016103fd565b9150509250925092565b600081519050919050565b600081905092915050565b60005b8381101561052d578082015181840152602081019050610512565b8381111561053c576000848401525b50505050565b600061054d826104f9565b6105578185610504565b935061056781856020860161050f565b80840191505092915050565b600061057f8284610542565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105c4826103b2565b91506105cf836103b2565b9250828210156105e2576105e161058a565b5b828203905092915050565b60006105f8826103b2565b9150610603836103b2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156106385761063761058a565b5b82820190509291505056fea26469706673582212209c3a15397c5fc2d0668cca2ef1bce80e3b8124697ac44a54a13f74f9e30961d164736f6c634300080b0033"
var TransferSMBin = "0x608060405234801561001057600080fd5b50610679806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063612d2bff14610051578063ab71bf091461006d578063cd93c25d14610089578063f2f4ee6d146100b9575b600080fd5b61006b6004803603810190610066919061039f565b6100d5565b005b6100876004803603810190610082919061039f565b610144565b005b6100a3600480360381019061009e919061042a565b6101c1565b6040516100b09190610482565b60405180910390f35b6100d360048036038101906100ce919061049d565b6101e8565b005b806000846040516100e69190610573565b9081526020016040518091039020600082825461010391906105b9565b925050819055508060008360405161011b9190610573565b9081526020016040518091039020600082825461013891906105ed565b92505081905550505050565b806000846040516101559190610573565b9081526020016040518091039020600082825461017291906105b9565b925050819055508060008360405161018a9190610573565b908152602001604051809103902060008282546101a791906105ed565b9250508190555060648111156101bc57600080fd5b505050565b600080826040516101d29190610573565b9081526020016040518091039020549050919050565b806000836040516101f99190610573565b9081526020016040518091039020819055505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7fb95aa35500000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102768261022d565b810181811067ffffffffffffffff821117156102955761029461023e565b5b80604052505050565b60006102a861020f565b90506102b4828261026d565b919050565b600067ffffffffffffffff8211156102d4576102d361023e565b5b6102dd8261022d565b9050602081019050919050565b82818337600083830152505050565b600061030c610307846102b9565b61029e565b90508281526020810184848401111561032857610327610228565b5b6103338482856102ea565b509392505050565b600082601f8301126103505761034f610223565b5b81356103608482602086016102f9565b91505092915050565b6000819050919050565b61037c81610369565b811461038757600080fd5b50565b60008135905061039981610373565b92915050565b6000806000606084860312156103b8576103b7610219565b5b600084013567ffffffffffffffff8111156103d6576103d561021e565b5b6103e28682870161033b565b935050602084013567ffffffffffffffff8111156104035761040261021e565b5b61040f8682870161033b565b92505060406104208682870161038a565b9150509250925092565b6000602082840312156104405761043f610219565b5b600082013567ffffffffffffffff81111561045e5761045d61021e565b5b61046a8482850161033b565b91505092915050565b61047c81610369565b82525050565b60006020820190506104976000830184610473565b92915050565b600080604083850312156104b4576104b3610219565b5b600083013567ffffffffffffffff8111156104d2576104d161021e565b5b6104de8582860161033b565b92505060206104ef8582860161038a565b9150509250929050565b600081519050919050565b600081905092915050565b60005b8381101561052d578082015181840152602081019050610512565b8381111561053c576000848401525b50505050565b600061054d826104f9565b6105578185610504565b935061056781856020860161050f565b80840191505092915050565b600061057f8284610542565b915081905092915050565b7fb95aa35500000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105c482610369565b91506105cf83610369565b9250828210156105e2576105e161058a565b5b828203905092915050565b60006105f882610369565b915061060383610369565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156106385761063761058a565b5b82820190509291505056fea26469706673582212204f8ebd1bd8a95fec67c4a519537e11316b2dd90a6d972b04ef627aac78dc982f64736f6c634300080b0033"

// DeployTransfer deploys a new contract, binding an instance of Transfer to it.
func DeployTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *Transfer, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransferSMBin)
	} else {
		bytecode = common.FromHex(TransferBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

func AsyncDeployTransfer(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(TransferSMBin)
	} else {
		bytecode = common.FromHex(TransferBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Transfer is an auto generated Go binding around a Solidity contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
	TransferFilterer   // Log filterer for contract events
}

// TransferCaller is an auto generated read-only Go binding around a Solidity contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around a Solidity contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around a Solidity contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// NewTransferFilterer creates a new log filterer instance of Transfer, bound to a specific deployed contract.
func NewTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferFilterer, error) {
	contract, err := bindTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferFilterer{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.TransferTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string name) constant returns(uint256)
func (_Transfer *TransferCaller) BalanceOf(opts *bind.CallOpts, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "balanceOf", name)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string name) constant returns(uint256)
func (_Transfer *TransferSession) BalanceOf(name string) (*big.Int, error) {
	return _Transfer.Contract.BalanceOf(&_Transfer.CallOpts, name)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string name) constant returns(uint256)
func (_Transfer *TransferCallerSession) BalanceOf(name string) (*big.Int, error) {
	return _Transfer.Contract.BalanceOf(&_Transfer.CallOpts, name)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string name, uint256 num) returns()
func (_Transfer *TransferTransactor) Set(opts *bind.TransactOpts, name string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Transfer.contract.TransactWithResult(opts, out, "set", name, num)
	return transaction, receipt, err
}

func (_Transfer *TransferTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.AsyncTransact(opts, handler, "set", name, num)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string name, uint256 num) returns()
func (_Transfer *TransferSession) Set(name string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.Set(&_Transfer.TransactOpts, name, num)
}

func (_Transfer *TransferSession) AsyncSet(handler func(*types.Receipt, error), name string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.AsyncSet(handler, &_Transfer.TransactOpts, name, num)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string name, uint256 num) returns()
func (_Transfer *TransferTransactorSession) Set(name string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.Set(&_Transfer.TransactOpts, name, num)
}

func (_Transfer *TransferTransactorSession) AsyncSet(handler func(*types.Receipt, error), name string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.AsyncSet(handler, &_Transfer.TransactOpts, name, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 num) returns()
func (_Transfer *TransferTransactor) Transfer(opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Transfer.contract.TransactWithResult(opts, out, "transfer", from, to, num)
	return transaction, receipt, err
}

func (_Transfer *TransferTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.AsyncTransact(opts, handler, "transfer", from, to, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 num) returns()
func (_Transfer *TransferSession) Transfer(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, from, to, num)
}

func (_Transfer *TransferSession) AsyncTransfer(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.AsyncTransfer(handler, &_Transfer.TransactOpts, from, to, num)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string from, string to, uint256 num) returns()
func (_Transfer *TransferTransactorSession) Transfer(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.Transfer(&_Transfer.TransactOpts, from, to, num)
}

func (_Transfer *TransferTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.AsyncTransfer(handler, &_Transfer.TransactOpts, from, to, num)
}

// TransferWithRevert is a paid mutator transaction binding the contract method 0xfad42f87.
//
// Solidity: function transferWithRevert(string from, string to, uint256 num) returns()
func (_Transfer *TransferTransactor) TransferWithRevert(opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Transfer.contract.TransactWithResult(opts, out, "transferWithRevert", from, to, num)
	return transaction, receipt, err
}

func (_Transfer *TransferTransactor) AsyncTransferWithRevert(handler func(*types.Receipt, error), opts *bind.TransactOpts, from string, to string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.contract.AsyncTransact(opts, handler, "transferWithRevert", from, to, num)
}

// TransferWithRevert is a paid mutator transaction binding the contract method 0xfad42f87.
//
// Solidity: function transferWithRevert(string from, string to, uint256 num) returns()
func (_Transfer *TransferSession) TransferWithRevert(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.TransferWithRevert(&_Transfer.TransactOpts, from, to, num)
}

func (_Transfer *TransferSession) AsyncTransferWithRevert(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.AsyncTransferWithRevert(handler, &_Transfer.TransactOpts, from, to, num)
}

// TransferWithRevert is a paid mutator transaction binding the contract method 0xfad42f87.
//
// Solidity: function transferWithRevert(string from, string to, uint256 num) returns()
func (_Transfer *TransferTransactorSession) TransferWithRevert(from string, to string, num *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Transfer.Contract.TransferWithRevert(&_Transfer.TransactOpts, from, to, num)
}

func (_Transfer *TransferTransactorSession) AsyncTransferWithRevert(handler func(*types.Receipt, error), from string, to string, num *big.Int) (*types.Transaction, error) {
	return _Transfer.Contract.AsyncTransferWithRevert(handler, &_Transfer.TransactOpts, from, to, num)
}

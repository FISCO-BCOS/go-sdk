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

// ContractAuthABI is the input ABI used to generate the binding from.
const ContractAuthABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_func\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkMethodAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeDeployAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_func\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeMethodAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"contractAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_func\",\"type\":\"bytes4\"}],\"name\":\"getMethodAuth\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasDeployAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"}],\"name\":\"initAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"openDeployAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_func\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"openMethodAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"resetAdmin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isFreeze\",\"type\":\"bool\"}],\"name\":\"setContractStatus\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"enumStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setContractStatus\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"setDeployAuthType\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_func\",\"type\":\"bytes4\"},{\"internalType\":\"uint8\",\"name\":\"authType\",\"type\":\"uint8\"}],\"name\":\"setMethodAuthType\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractAuth is an auto generated Go binding around a Solidity contract.
type ContractAuth struct {
	ContractAuthCaller     // Read-only binding to the contract
	ContractAuthTransactor // Write-only binding to the contract
	ContractAuthFilterer   // Log filterer for contract events
}

// ContractAuthCaller is an auto generated read-only Go binding around a Solidity contract.
type ContractAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAuthTransactor is an auto generated write-only Go binding around a Solidity contract.
type ContractAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAuthFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ContractAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAuthSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ContractAuthSession struct {
	Contract     *ContractAuth     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractAuthCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ContractAuthCallerSession struct {
	Contract *ContractAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractAuthTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ContractAuthTransactorSession struct {
	Contract     *ContractAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractAuthRaw is an auto generated low-level Go binding around a Solidity contract.
type ContractAuthRaw struct {
	Contract *ContractAuth // Generic contract binding to access the raw methods on
}

// ContractAuthCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ContractAuthCallerRaw struct {
	Contract *ContractAuthCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAuthTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ContractAuthTransactorRaw struct {
	Contract *ContractAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAuth creates a new instance of ContractAuth, bound to a specific deployed contract.
func NewContractAuth(address common.Address, backend bind.ContractBackend) (*ContractAuth, error) {
	contract, err := bindContractAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAuth{ContractAuthCaller: ContractAuthCaller{contract: contract}, ContractAuthTransactor: ContractAuthTransactor{contract: contract}, ContractAuthFilterer: ContractAuthFilterer{contract: contract}}, nil
}

// NewContractAuthCaller creates a new read-only instance of ContractAuth, bound to a specific deployed contract.
func NewContractAuthCaller(address common.Address, caller bind.ContractCaller) (*ContractAuthCaller, error) {
	contract, err := bindContractAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAuthCaller{contract: contract}, nil
}

// NewContractAuthTransactor creates a new write-only instance of ContractAuth, bound to a specific deployed contract.
func NewContractAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAuthTransactor, error) {
	contract, err := bindContractAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAuthTransactor{contract: contract}, nil
}

// NewContractAuthFilterer creates a new log filterer instance of ContractAuth, bound to a specific deployed contract.
func NewContractAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAuthFilterer, error) {
	contract, err := bindContractAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAuthFilterer{contract: contract}, nil
}

// bindContractAuth binds a generic wrapper to an already deployed contract.
func bindContractAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAuth *ContractAuthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAuth.Contract.ContractAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAuth *ContractAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.ContractAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAuth *ContractAuthRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.ContractAuthTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAuth *ContractAuthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAuth *ContractAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAuth *ContractAuthTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// CheckMethodAuth is a free data retrieval call binding the contract method 0xd8662aa4.
//
// Solidity: function checkMethodAuth(address contractAddr, bytes4 _func, address account) constant returns(bool)
func (_ContractAuth *ContractAuthCaller) CheckMethodAuth(opts *bind.CallOpts, contractAddr common.Address, _func [4]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAuth.contract.Call(opts, out, "checkMethodAuth", contractAddr, _func, account)
	return *ret0, err
}

// CheckMethodAuth is a free data retrieval call binding the contract method 0xd8662aa4.
//
// Solidity: function checkMethodAuth(address contractAddr, bytes4 _func, address account) constant returns(bool)
func (_ContractAuth *ContractAuthSession) CheckMethodAuth(contractAddr common.Address, _func [4]byte, account common.Address) (bool, error) {
	return _ContractAuth.Contract.CheckMethodAuth(&_ContractAuth.CallOpts, contractAddr, _func, account)
}

// CheckMethodAuth is a free data retrieval call binding the contract method 0xd8662aa4.
//
// Solidity: function checkMethodAuth(address contractAddr, bytes4 _func, address account) constant returns(bool)
func (_ContractAuth *ContractAuthCallerSession) CheckMethodAuth(contractAddr common.Address, _func [4]byte, account common.Address) (bool, error) {
	return _ContractAuth.Contract.CheckMethodAuth(&_ContractAuth.CallOpts, contractAddr, _func, account)
}

// ContractAvailable is a free data retrieval call binding the contract method 0x2c8c4a4f.
//
// Solidity: function contractAvailable(address _address) constant returns(bool)
func (_ContractAuth *ContractAuthCaller) ContractAvailable(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAuth.contract.Call(opts, out, "contractAvailable", _address)
	return *ret0, err
}

// ContractAvailable is a free data retrieval call binding the contract method 0x2c8c4a4f.
//
// Solidity: function contractAvailable(address _address) constant returns(bool)
func (_ContractAuth *ContractAuthSession) ContractAvailable(_address common.Address) (bool, error) {
	return _ContractAuth.Contract.ContractAvailable(&_ContractAuth.CallOpts, _address)
}

// ContractAvailable is a free data retrieval call binding the contract method 0x2c8c4a4f.
//
// Solidity: function contractAvailable(address _address) constant returns(bool)
func (_ContractAuth *ContractAuthCallerSession) ContractAvailable(_address common.Address) (bool, error) {
	return _ContractAuth.Contract.ContractAvailable(&_ContractAuth.CallOpts, _address)
}

// DeployType is a free data retrieval call binding the contract method 0x1749bea9.
//
// Solidity: function deployType() constant returns(uint256)
func (_ContractAuth *ContractAuthCaller) DeployType(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractAuth.contract.Call(opts, out, "deployType")
	return *ret0, err
}

// DeployType is a free data retrieval call binding the contract method 0x1749bea9.
//
// Solidity: function deployType() constant returns(uint256)
func (_ContractAuth *ContractAuthSession) DeployType() (*big.Int, error) {
	return _ContractAuth.Contract.DeployType(&_ContractAuth.CallOpts)
}

// DeployType is a free data retrieval call binding the contract method 0x1749bea9.
//
// Solidity: function deployType() constant returns(uint256)
func (_ContractAuth *ContractAuthCallerSession) DeployType() (*big.Int, error) {
	return _ContractAuth.Contract.DeployType(&_ContractAuth.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address contractAddr) constant returns(address)
func (_ContractAuth *ContractAuthCaller) GetAdmin(opts *bind.CallOpts, contractAddr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractAuth.contract.Call(opts, out, "getAdmin", contractAddr)
	return *ret0, err
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address contractAddr) constant returns(address)
func (_ContractAuth *ContractAuthSession) GetAdmin(contractAddr common.Address) (common.Address, error) {
	return _ContractAuth.Contract.GetAdmin(&_ContractAuth.CallOpts, contractAddr)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address contractAddr) constant returns(address)
func (_ContractAuth *ContractAuthCallerSession) GetAdmin(contractAddr common.Address) (common.Address, error) {
	return _ContractAuth.Contract.GetAdmin(&_ContractAuth.CallOpts, contractAddr)
}

// GetMethodAuth is a free data retrieval call binding the contract method 0x0578519a.
//
// Solidity: function getMethodAuth(address contractAddr, bytes4 _func) constant returns(uint8, string[], string[])
func (_ContractAuth *ContractAuthCaller) GetMethodAuth(opts *bind.CallOpts, contractAddr common.Address, _func [4]byte) (uint8, []string, []string, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([]string)
		ret2 = new([]string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ContractAuth.contract.Call(opts, out, "getMethodAuth", contractAddr, _func)
	return *ret0, *ret1, *ret2, err
}

// GetMethodAuth is a free data retrieval call binding the contract method 0x0578519a.
//
// Solidity: function getMethodAuth(address contractAddr, bytes4 _func) constant returns(uint8, string[], string[])
func (_ContractAuth *ContractAuthSession) GetMethodAuth(contractAddr common.Address, _func [4]byte) (uint8, []string, []string, error) {
	return _ContractAuth.Contract.GetMethodAuth(&_ContractAuth.CallOpts, contractAddr, _func)
}

// GetMethodAuth is a free data retrieval call binding the contract method 0x0578519a.
//
// Solidity: function getMethodAuth(address contractAddr, bytes4 _func) constant returns(uint8, string[], string[])
func (_ContractAuth *ContractAuthCallerSession) GetMethodAuth(contractAddr common.Address, _func [4]byte) (uint8, []string, []string, error) {
	return _ContractAuth.Contract.GetMethodAuth(&_ContractAuth.CallOpts, contractAddr, _func)
}

// HasDeployAuth is a free data retrieval call binding the contract method 0x630577e5.
//
// Solidity: function hasDeployAuth(address account) constant returns(bool)
func (_ContractAuth *ContractAuthCaller) HasDeployAuth(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAuth.contract.Call(opts, out, "hasDeployAuth", account)
	return *ret0, err
}

// HasDeployAuth is a free data retrieval call binding the contract method 0x630577e5.
//
// Solidity: function hasDeployAuth(address account) constant returns(bool)
func (_ContractAuth *ContractAuthSession) HasDeployAuth(account common.Address) (bool, error) {
	return _ContractAuth.Contract.HasDeployAuth(&_ContractAuth.CallOpts, account)
}

// HasDeployAuth is a free data retrieval call binding the contract method 0x630577e5.
//
// Solidity: function hasDeployAuth(address account) constant returns(bool)
func (_ContractAuth *ContractAuthCallerSession) HasDeployAuth(account common.Address) (bool, error) {
	return _ContractAuth.Contract.HasDeployAuth(&_ContractAuth.CallOpts, account)
}

// CloseDeployAuth is a paid mutator transaction binding the contract method 0x56bd7084.
//
// Solidity: function closeDeployAuth(address account) returns(int256)
func (_ContractAuth *ContractAuthTransactor) CloseDeployAuth(opts *bind.TransactOpts, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "closeDeployAuth", account)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncCloseDeployAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "closeDeployAuth", account)
}

// CloseDeployAuth is a paid mutator transaction binding the contract method 0x56bd7084.
//
// Solidity: function closeDeployAuth(address account) returns(int256)
func (_ContractAuth *ContractAuthSession) CloseDeployAuth(account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.CloseDeployAuth(&_ContractAuth.TransactOpts, account)
}

func (_ContractAuth *ContractAuthSession) AsyncCloseDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncCloseDeployAuth(handler, &_ContractAuth.TransactOpts, account)
}

// CloseDeployAuth is a paid mutator transaction binding the contract method 0x56bd7084.
//
// Solidity: function closeDeployAuth(address account) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) CloseDeployAuth(account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.CloseDeployAuth(&_ContractAuth.TransactOpts, account)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncCloseDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncCloseDeployAuth(handler, &_ContractAuth.TransactOpts, account)
}

// CloseMethodAuth is a paid mutator transaction binding the contract method 0xcb7c5c11.
//
// Solidity: function closeMethodAuth(address contractAddr, bytes4 _func, address account) returns(int256)
func (_ContractAuth *ContractAuthTransactor) CloseMethodAuth(opts *bind.TransactOpts, contractAddr common.Address, _func [4]byte, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "closeMethodAuth", contractAddr, _func, account)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncCloseMethodAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, _func [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "closeMethodAuth", contractAddr, _func, account)
}

// CloseMethodAuth is a paid mutator transaction binding the contract method 0xcb7c5c11.
//
// Solidity: function closeMethodAuth(address contractAddr, bytes4 _func, address account) returns(int256)
func (_ContractAuth *ContractAuthSession) CloseMethodAuth(contractAddr common.Address, _func [4]byte, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.CloseMethodAuth(&_ContractAuth.TransactOpts, contractAddr, _func, account)
}

func (_ContractAuth *ContractAuthSession) AsyncCloseMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, _func [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncCloseMethodAuth(handler, &_ContractAuth.TransactOpts, contractAddr, _func, account)
}

// CloseMethodAuth is a paid mutator transaction binding the contract method 0xcb7c5c11.
//
// Solidity: function closeMethodAuth(address contractAddr, bytes4 _func, address account) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) CloseMethodAuth(contractAddr common.Address, _func [4]byte, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.CloseMethodAuth(&_ContractAuth.TransactOpts, contractAddr, _func, account)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncCloseMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, _func [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncCloseMethodAuth(handler, &_ContractAuth.TransactOpts, contractAddr, _func, account)
}

// InitAuth is a paid mutator transaction binding the contract method 0x8805dd3c.
//
// Solidity: function initAuth(string account) returns(int256)
func (_ContractAuth *ContractAuthTransactor) InitAuth(opts *bind.TransactOpts, account string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "initAuth", account)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncInitAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, account string) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "initAuth", account)
}

// InitAuth is a paid mutator transaction binding the contract method 0x8805dd3c.
//
// Solidity: function initAuth(string account) returns(int256)
func (_ContractAuth *ContractAuthSession) InitAuth(account string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.InitAuth(&_ContractAuth.TransactOpts, account)
}

func (_ContractAuth *ContractAuthSession) AsyncInitAuth(handler func(*types.Receipt, error), account string) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncInitAuth(handler, &_ContractAuth.TransactOpts, account)
}

// InitAuth is a paid mutator transaction binding the contract method 0x8805dd3c.
//
// Solidity: function initAuth(string account) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) InitAuth(account string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.InitAuth(&_ContractAuth.TransactOpts, account)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncInitAuth(handler func(*types.Receipt, error), account string) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncInitAuth(handler, &_ContractAuth.TransactOpts, account)
}

// OpenDeployAuth is a paid mutator transaction binding the contract method 0x61548099.
//
// Solidity: function openDeployAuth(address account) returns(int256)
func (_ContractAuth *ContractAuthTransactor) OpenDeployAuth(opts *bind.TransactOpts, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "openDeployAuth", account)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncOpenDeployAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "openDeployAuth", account)
}

// OpenDeployAuth is a paid mutator transaction binding the contract method 0x61548099.
//
// Solidity: function openDeployAuth(address account) returns(int256)
func (_ContractAuth *ContractAuthSession) OpenDeployAuth(account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.OpenDeployAuth(&_ContractAuth.TransactOpts, account)
}

func (_ContractAuth *ContractAuthSession) AsyncOpenDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncOpenDeployAuth(handler, &_ContractAuth.TransactOpts, account)
}

// OpenDeployAuth is a paid mutator transaction binding the contract method 0x61548099.
//
// Solidity: function openDeployAuth(address account) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) OpenDeployAuth(account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.OpenDeployAuth(&_ContractAuth.TransactOpts, account)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncOpenDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncOpenDeployAuth(handler, &_ContractAuth.TransactOpts, account)
}

// OpenMethodAuth is a paid mutator transaction binding the contract method 0x0c82b73d.
//
// Solidity: function openMethodAuth(address contractAddr, bytes4 _func, address account) returns(int256)
func (_ContractAuth *ContractAuthTransactor) OpenMethodAuth(opts *bind.TransactOpts, contractAddr common.Address, _func [4]byte, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "openMethodAuth", contractAddr, _func, account)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncOpenMethodAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, _func [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "openMethodAuth", contractAddr, _func, account)
}

// OpenMethodAuth is a paid mutator transaction binding the contract method 0x0c82b73d.
//
// Solidity: function openMethodAuth(address contractAddr, bytes4 _func, address account) returns(int256)
func (_ContractAuth *ContractAuthSession) OpenMethodAuth(contractAddr common.Address, _func [4]byte, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.OpenMethodAuth(&_ContractAuth.TransactOpts, contractAddr, _func, account)
}

func (_ContractAuth *ContractAuthSession) AsyncOpenMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, _func [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncOpenMethodAuth(handler, &_ContractAuth.TransactOpts, contractAddr, _func, account)
}

// OpenMethodAuth is a paid mutator transaction binding the contract method 0x0c82b73d.
//
// Solidity: function openMethodAuth(address contractAddr, bytes4 _func, address account) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) OpenMethodAuth(contractAddr common.Address, _func [4]byte, account common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.OpenMethodAuth(&_ContractAuth.TransactOpts, contractAddr, _func, account)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncOpenMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, _func [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncOpenMethodAuth(handler, &_ContractAuth.TransactOpts, contractAddr, _func, account)
}

// ResetAdmin is a paid mutator transaction binding the contract method 0xc53057b4.
//
// Solidity: function resetAdmin(address contractAddr, address admin) returns(int256)
func (_ContractAuth *ContractAuthTransactor) ResetAdmin(opts *bind.TransactOpts, contractAddr common.Address, admin common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "resetAdmin", contractAddr, admin)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncResetAdmin(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "resetAdmin", contractAddr, admin)
}

// ResetAdmin is a paid mutator transaction binding the contract method 0xc53057b4.
//
// Solidity: function resetAdmin(address contractAddr, address admin) returns(int256)
func (_ContractAuth *ContractAuthSession) ResetAdmin(contractAddr common.Address, admin common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.ResetAdmin(&_ContractAuth.TransactOpts, contractAddr, admin)
}

func (_ContractAuth *ContractAuthSession) AsyncResetAdmin(handler func(*types.Receipt, error), contractAddr common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncResetAdmin(handler, &_ContractAuth.TransactOpts, contractAddr, admin)
}

// ResetAdmin is a paid mutator transaction binding the contract method 0xc53057b4.
//
// Solidity: function resetAdmin(address contractAddr, address admin) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) ResetAdmin(contractAddr common.Address, admin common.Address) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.ResetAdmin(&_ContractAuth.TransactOpts, contractAddr, admin)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncResetAdmin(handler func(*types.Receipt, error), contractAddr common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncResetAdmin(handler, &_ContractAuth.TransactOpts, contractAddr, admin)
}

// SetContractStatus is a paid mutator transaction binding the contract method 0x81c81cdc.
//
// Solidity: function setContractStatus(address _address, bool isFreeze) returns(int256)
func (_ContractAuth *ContractAuthTransactor) SetContractStatus(opts *bind.TransactOpts, _address common.Address, isFreeze bool) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "setContractStatus", _address, isFreeze)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncSetContractStatus(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, isFreeze bool) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "setContractStatus", _address, isFreeze)
}

// SetContractStatus is a paid mutator transaction binding the contract method 0x81c81cdc.
//
// Solidity: function setContractStatus(address _address, bool isFreeze) returns(int256)
func (_ContractAuth *ContractAuthSession) SetContractStatus(_address common.Address, isFreeze bool) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetContractStatus(&_ContractAuth.TransactOpts, _address, isFreeze)
}

func (_ContractAuth *ContractAuthSession) AsyncSetContractStatus(handler func(*types.Receipt, error), _address common.Address, isFreeze bool) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetContractStatus(handler, &_ContractAuth.TransactOpts, _address, isFreeze)
}

// SetContractStatus is a paid mutator transaction binding the contract method 0x81c81cdc.
//
// Solidity: function setContractStatus(address _address, bool isFreeze) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) SetContractStatus(_address common.Address, isFreeze bool) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetContractStatus(&_ContractAuth.TransactOpts, _address, isFreeze)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncSetContractStatus(handler func(*types.Receipt, error), _address common.Address, isFreeze bool) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetContractStatus(handler, &_ContractAuth.TransactOpts, _address, isFreeze)
}

// SetContractStatus0 is a paid mutator transaction binding the contract method 0xf67f8dcc.
//
// Solidity: function setContractStatus(address _address, uint8 _status) returns(int256)
func (_ContractAuth *ContractAuthTransactor) SetContractStatus0(opts *bind.TransactOpts, _address common.Address, _status uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "setContractStatus0", _address, _status)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncSetContractStatus0(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, _status uint8) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "setContractStatus0", _address, _status)
}

// SetContractStatus0 is a paid mutator transaction binding the contract method 0xf67f8dcc.
//
// Solidity: function setContractStatus(address _address, uint8 _status) returns(int256)
func (_ContractAuth *ContractAuthSession) SetContractStatus0(_address common.Address, _status uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetContractStatus0(&_ContractAuth.TransactOpts, _address, _status)
}

func (_ContractAuth *ContractAuthSession) AsyncSetContractStatus0(handler func(*types.Receipt, error), _address common.Address, _status uint8) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetContractStatus0(handler, &_ContractAuth.TransactOpts, _address, _status)
}

// SetContractStatus0 is a paid mutator transaction binding the contract method 0xf67f8dcc.
//
// Solidity: function setContractStatus(address _address, uint8 _status) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) SetContractStatus0(_address common.Address, _status uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetContractStatus0(&_ContractAuth.TransactOpts, _address, _status)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncSetContractStatus0(handler func(*types.Receipt, error), _address common.Address, _status uint8) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetContractStatus0(handler, &_ContractAuth.TransactOpts, _address, _status)
}

// SetDeployAuthType is a paid mutator transaction binding the contract method 0xbb0aa40c.
//
// Solidity: function setDeployAuthType(uint8 _type) returns(int256)
func (_ContractAuth *ContractAuthTransactor) SetDeployAuthType(opts *bind.TransactOpts, _type uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "setDeployAuthType", _type)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncSetDeployAuthType(handler func(*types.Receipt, error), opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "setDeployAuthType", _type)
}

// SetDeployAuthType is a paid mutator transaction binding the contract method 0xbb0aa40c.
//
// Solidity: function setDeployAuthType(uint8 _type) returns(int256)
func (_ContractAuth *ContractAuthSession) SetDeployAuthType(_type uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetDeployAuthType(&_ContractAuth.TransactOpts, _type)
}

func (_ContractAuth *ContractAuthSession) AsyncSetDeployAuthType(handler func(*types.Receipt, error), _type uint8) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetDeployAuthType(handler, &_ContractAuth.TransactOpts, _type)
}

// SetDeployAuthType is a paid mutator transaction binding the contract method 0xbb0aa40c.
//
// Solidity: function setDeployAuthType(uint8 _type) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) SetDeployAuthType(_type uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetDeployAuthType(&_ContractAuth.TransactOpts, _type)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncSetDeployAuthType(handler func(*types.Receipt, error), _type uint8) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetDeployAuthType(handler, &_ContractAuth.TransactOpts, _type)
}

// SetMethodAuthType is a paid mutator transaction binding the contract method 0x9cc3ca0f.
//
// Solidity: function setMethodAuthType(address contractAddr, bytes4 _func, uint8 authType) returns(int256)
func (_ContractAuth *ContractAuthTransactor) SetMethodAuthType(opts *bind.TransactOpts, contractAddr common.Address, _func [4]byte, authType uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _ContractAuth.contract.TransactWithResult(opts, out, "setMethodAuthType", contractAddr, _func, authType)
	return *ret0, transaction, receipt, err
}

func (_ContractAuth *ContractAuthTransactor) AsyncSetMethodAuthType(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, _func [4]byte, authType uint8) (*types.Transaction, error) {
	return _ContractAuth.contract.AsyncTransact(opts, handler, "setMethodAuthType", contractAddr, _func, authType)
}

// SetMethodAuthType is a paid mutator transaction binding the contract method 0x9cc3ca0f.
//
// Solidity: function setMethodAuthType(address contractAddr, bytes4 _func, uint8 authType) returns(int256)
func (_ContractAuth *ContractAuthSession) SetMethodAuthType(contractAddr common.Address, _func [4]byte, authType uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetMethodAuthType(&_ContractAuth.TransactOpts, contractAddr, _func, authType)
}

func (_ContractAuth *ContractAuthSession) AsyncSetMethodAuthType(handler func(*types.Receipt, error), contractAddr common.Address, _func [4]byte, authType uint8) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetMethodAuthType(handler, &_ContractAuth.TransactOpts, contractAddr, _func, authType)
}

// SetMethodAuthType is a paid mutator transaction binding the contract method 0x9cc3ca0f.
//
// Solidity: function setMethodAuthType(address contractAddr, bytes4 _func, uint8 authType) returns(int256)
func (_ContractAuth *ContractAuthTransactorSession) SetMethodAuthType(contractAddr common.Address, _func [4]byte, authType uint8) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _ContractAuth.Contract.SetMethodAuthType(&_ContractAuth.TransactOpts, contractAddr, _func, authType)
}

func (_ContractAuth *ContractAuthTransactorSession) AsyncSetMethodAuthType(handler func(*types.Receipt, error), contractAddr common.Address, _func [4]byte, authType uint8) (*types.Transaction, error) {
	return _ContractAuth.Contract.AsyncSetMethodAuthType(handler, &_ContractAuth.TransactOpts, contractAddr, _func, authType)
}

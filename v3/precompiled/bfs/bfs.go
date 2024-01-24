// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bfs

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

// BfsInfo is an auto generated low-level Go binding around an user-defined struct.
type BfsInfo struct {
	FileName string
	FileType string
	Ext      []string
}

// BfsABI is the input ABI used to generate the binding from.
const BfsABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_abi\",\"type\":\"string\"}],\"name\":\"link\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_abi\",\"type\":\"string\"}],\"name\":\"link\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"file_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"file_type\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ext\",\"type\":\"string[]\"}],\"internalType\":\"structBfsInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"file_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"file_type\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ext\",\"type\":\"string[]\"}],\"internalType\":\"structBfsInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"}],\"name\":\"mkdir\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"}],\"name\":\"readlink\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebuildBfs\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"absolutePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"}],\"name\":\"touch\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Bfs is an auto generated Go binding around a Solidity contract.
type Bfs struct {
	BfsCaller     // Read-only binding to the contract
	BfsTransactor // Write-only binding to the contract
	BfsFilterer   // Log filterer for contract events
}

// BfsCaller is an auto generated read-only Go binding around a Solidity contract.
type BfsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BfsTransactor is an auto generated write-only Go binding around a Solidity contract.
type BfsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BfsFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type BfsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BfsSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type BfsSession struct {
	Contract     *Bfs              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BfsCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type BfsCallerSession struct {
	Contract *BfsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BfsTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type BfsTransactorSession struct {
	Contract     *BfsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BfsRaw is an auto generated low-level Go binding around a Solidity contract.
type BfsRaw struct {
	Contract *Bfs // Generic contract binding to access the raw methods on
}

// BfsCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type BfsCallerRaw struct {
	Contract *BfsCaller // Generic read-only contract binding to access the raw methods on
}

// BfsTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type BfsTransactorRaw struct {
	Contract *BfsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBfs creates a new instance of Bfs, bound to a specific deployed contract.
func NewBfs(address common.Address, backend bind.ContractBackend) (*Bfs, error) {
	contract, err := bindBfs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bfs{BfsCaller: BfsCaller{contract: contract}, BfsTransactor: BfsTransactor{contract: contract}, BfsFilterer: BfsFilterer{contract: contract}}, nil
}

// NewBfsCaller creates a new read-only instance of Bfs, bound to a specific deployed contract.
func NewBfsCaller(address common.Address, caller bind.ContractCaller) (*BfsCaller, error) {
	contract, err := bindBfs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BfsCaller{contract: contract}, nil
}

// NewBfsTransactor creates a new write-only instance of Bfs, bound to a specific deployed contract.
func NewBfsTransactor(address common.Address, transactor bind.ContractTransactor) (*BfsTransactor, error) {
	contract, err := bindBfs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BfsTransactor{contract: contract}, nil
}

// NewBfsFilterer creates a new log filterer instance of Bfs, bound to a specific deployed contract.
func NewBfsFilterer(address common.Address, filterer bind.ContractFilterer) (*BfsFilterer, error) {
	contract, err := bindBfs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BfsFilterer{contract: contract}, nil
}

// bindBfs binds a generic wrapper to an already deployed contract.
func bindBfs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BfsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bfs *BfsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bfs.Contract.BfsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bfs *BfsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.BfsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bfs *BfsRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.BfsTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bfs *BfsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bfs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bfs *BfsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bfs *BfsTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// List is a free data retrieval call binding the contract method 0x912f3095.
//
// Solidity: function list(string absolutePath, uint256 offset, uint256 limit) constant returns(int256, []BfsInfo)
func (_Bfs *BfsCaller) List(opts *bind.CallOpts, absolutePath string, offset *big.Int, limit *big.Int) (*big.Int, []BfsInfo, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]BfsInfo)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Bfs.contract.Call(opts, out, "list", absolutePath, offset, limit)
	return *ret0, *ret1, err
}

// List is a free data retrieval call binding the contract method 0x912f3095.
//
// Solidity: function list(string absolutePath, uint256 offset, uint256 limit) constant returns(int256, []BfsInfo)
func (_Bfs *BfsSession) List(absolutePath string, offset *big.Int, limit *big.Int) (*big.Int, []BfsInfo, error) {
	return _Bfs.Contract.List(&_Bfs.CallOpts, absolutePath, offset, limit)
}

// List is a free data retrieval call binding the contract method 0x912f3095.
//
// Solidity: function list(string absolutePath, uint256 offset, uint256 limit) constant returns(int256, []BfsInfo)
func (_Bfs *BfsCallerSession) List(absolutePath string, offset *big.Int, limit *big.Int) (*big.Int, []BfsInfo, error) {
	return _Bfs.Contract.List(&_Bfs.CallOpts, absolutePath, offset, limit)
}

// List0 is a free data retrieval call binding the contract method 0xfe42bf1a.
//
// Solidity: function list(string absolutePath) constant returns(int32, []BfsInfo)
func (_Bfs *BfsCaller) List0(opts *bind.CallOpts, absolutePath string) (int32, []BfsInfo, error) {
	var (
		ret0 = new(int32)
		ret1 = new([]BfsInfo)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Bfs.contract.Call(opts, out, "list0", absolutePath)
	return *ret0, *ret1, err
}

// List0 is a free data retrieval call binding the contract method 0xfe42bf1a.
//
// Solidity: function list(string absolutePath) constant returns(int32, []BfsInfo)
func (_Bfs *BfsSession) List0(absolutePath string) (int32, []BfsInfo, error) {
	return _Bfs.Contract.List0(&_Bfs.CallOpts, absolutePath)
}

// List0 is a free data retrieval call binding the contract method 0xfe42bf1a.
//
// Solidity: function list(string absolutePath) constant returns(int32, []BfsInfo)
func (_Bfs *BfsCallerSession) List0(absolutePath string) (int32, []BfsInfo, error) {
	return _Bfs.Contract.List0(&_Bfs.CallOpts, absolutePath)
}

// Readlink is a free data retrieval call binding the contract method 0x1d05a836.
//
// Solidity: function readlink(string absolutePath) constant returns(address)
func (_Bfs *BfsCaller) Readlink(opts *bind.CallOpts, absolutePath string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bfs.contract.Call(opts, out, "readlink", absolutePath)
	return *ret0, err
}

// Readlink is a free data retrieval call binding the contract method 0x1d05a836.
//
// Solidity: function readlink(string absolutePath) constant returns(address)
func (_Bfs *BfsSession) Readlink(absolutePath string) (common.Address, error) {
	return _Bfs.Contract.Readlink(&_Bfs.CallOpts, absolutePath)
}

// Readlink is a free data retrieval call binding the contract method 0x1d05a836.
//
// Solidity: function readlink(string absolutePath) constant returns(address)
func (_Bfs *BfsCallerSession) Readlink(absolutePath string) (common.Address, error) {
	return _Bfs.Contract.Readlink(&_Bfs.CallOpts, absolutePath)
}

// Link is a paid mutator transaction binding the contract method 0x8df118af.
//
// Solidity: function link(string absolutePath, string _address, string _abi) returns(int256)
func (_Bfs *BfsTransactor) Link(opts *bind.TransactOpts, absolutePath string, _address string, _abi string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Bfs.contract.TransactWithResult(opts, out, "link", absolutePath, _address, _abi)
	return *ret0, transaction, receipt, err
}

func (_Bfs *BfsTransactor) AsyncLink(handler func(*types.Receipt, error), opts *bind.TransactOpts, absolutePath string, _address string, _abi string) (*types.Transaction, error) {
	return _Bfs.contract.AsyncTransact(opts, handler, "link", absolutePath, _address, _abi)
}

// Link is a paid mutator transaction binding the contract method 0x8df118af.
//
// Solidity: function link(string absolutePath, string _address, string _abi) returns(int256)
func (_Bfs *BfsSession) Link(absolutePath string, _address string, _abi string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Link(&_Bfs.TransactOpts, absolutePath, _address, _abi)
}

func (_Bfs *BfsSession) AsyncLink(handler func(*types.Receipt, error), absolutePath string, _address string, _abi string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncLink(handler, &_Bfs.TransactOpts, absolutePath, _address, _abi)
}

// Link is a paid mutator transaction binding the contract method 0x8df118af.
//
// Solidity: function link(string absolutePath, string _address, string _abi) returns(int256)
func (_Bfs *BfsTransactorSession) Link(absolutePath string, _address string, _abi string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Link(&_Bfs.TransactOpts, absolutePath, _address, _abi)
}

func (_Bfs *BfsTransactorSession) AsyncLink(handler func(*types.Receipt, error), absolutePath string, _address string, _abi string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncLink(handler, &_Bfs.TransactOpts, absolutePath, _address, _abi)
}

// Link0 is a paid mutator transaction binding the contract method 0xe19c2fcf.
//
// Solidity: function link(string name, string version, string _address, string _abi) returns(int32)
func (_Bfs *BfsTransactor) Link0(opts *bind.TransactOpts, name string, version string, _address string, _abi string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Bfs.contract.TransactWithResult(opts, out, "link0", name, version, _address, _abi)
	return *ret0, transaction, receipt, err
}

func (_Bfs *BfsTransactor) AsyncLink0(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, version string, _address string, _abi string) (*types.Transaction, error) {
	return _Bfs.contract.AsyncTransact(opts, handler, "link0", name, version, _address, _abi)
}

// Link0 is a paid mutator transaction binding the contract method 0xe19c2fcf.
//
// Solidity: function link(string name, string version, string _address, string _abi) returns(int32)
func (_Bfs *BfsSession) Link0(name string, version string, _address string, _abi string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Link0(&_Bfs.TransactOpts, name, version, _address, _abi)
}

func (_Bfs *BfsSession) AsyncLink0(handler func(*types.Receipt, error), name string, version string, _address string, _abi string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncLink0(handler, &_Bfs.TransactOpts, name, version, _address, _abi)
}

// Link0 is a paid mutator transaction binding the contract method 0xe19c2fcf.
//
// Solidity: function link(string name, string version, string _address, string _abi) returns(int32)
func (_Bfs *BfsTransactorSession) Link0(name string, version string, _address string, _abi string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Link0(&_Bfs.TransactOpts, name, version, _address, _abi)
}

func (_Bfs *BfsTransactorSession) AsyncLink0(handler func(*types.Receipt, error), name string, version string, _address string, _abi string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncLink0(handler, &_Bfs.TransactOpts, name, version, _address, _abi)
}

// Mkdir is a paid mutator transaction binding the contract method 0x876b0eb2.
//
// Solidity: function mkdir(string absolutePath) returns(int32)
func (_Bfs *BfsTransactor) Mkdir(opts *bind.TransactOpts, absolutePath string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Bfs.contract.TransactWithResult(opts, out, "mkdir", absolutePath)
	return *ret0, transaction, receipt, err
}

func (_Bfs *BfsTransactor) AsyncMkdir(handler func(*types.Receipt, error), opts *bind.TransactOpts, absolutePath string) (*types.Transaction, error) {
	return _Bfs.contract.AsyncTransact(opts, handler, "mkdir", absolutePath)
}

// Mkdir is a paid mutator transaction binding the contract method 0x876b0eb2.
//
// Solidity: function mkdir(string absolutePath) returns(int32)
func (_Bfs *BfsSession) Mkdir(absolutePath string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Mkdir(&_Bfs.TransactOpts, absolutePath)
}

func (_Bfs *BfsSession) AsyncMkdir(handler func(*types.Receipt, error), absolutePath string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncMkdir(handler, &_Bfs.TransactOpts, absolutePath)
}

// Mkdir is a paid mutator transaction binding the contract method 0x876b0eb2.
//
// Solidity: function mkdir(string absolutePath) returns(int32)
func (_Bfs *BfsTransactorSession) Mkdir(absolutePath string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Mkdir(&_Bfs.TransactOpts, absolutePath)
}

func (_Bfs *BfsTransactorSession) AsyncMkdir(handler func(*types.Receipt, error), absolutePath string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncMkdir(handler, &_Bfs.TransactOpts, absolutePath)
}

// RebuildBfs is a paid mutator transaction binding the contract method 0xa007e274.
//
// Solidity: function rebuildBfs() returns(int256)
func (_Bfs *BfsTransactor) RebuildBfs(opts *bind.TransactOpts) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Bfs.contract.TransactWithResult(opts, out, "rebuildBfs")
	return *ret0, transaction, receipt, err
}

func (_Bfs *BfsTransactor) AsyncRebuildBfs(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bfs.contract.AsyncTransact(opts, handler, "rebuildBfs")
}

// RebuildBfs is a paid mutator transaction binding the contract method 0xa007e274.
//
// Solidity: function rebuildBfs() returns(int256)
func (_Bfs *BfsSession) RebuildBfs() (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.RebuildBfs(&_Bfs.TransactOpts)
}

func (_Bfs *BfsSession) AsyncRebuildBfs(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncRebuildBfs(handler, &_Bfs.TransactOpts)
}

// RebuildBfs is a paid mutator transaction binding the contract method 0xa007e274.
//
// Solidity: function rebuildBfs() returns(int256)
func (_Bfs *BfsTransactorSession) RebuildBfs() (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.RebuildBfs(&_Bfs.TransactOpts)
}

func (_Bfs *BfsTransactorSession) AsyncRebuildBfs(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncRebuildBfs(handler, &_Bfs.TransactOpts)
}

// Touch is a paid mutator transaction binding the contract method 0x131ffcdd.
//
// Solidity: function touch(string absolutePath, string fileType) returns(int32)
func (_Bfs *BfsTransactor) Touch(opts *bind.TransactOpts, absolutePath string, fileType string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Bfs.contract.TransactWithResult(opts, out, "touch", absolutePath, fileType)
	return *ret0, transaction, receipt, err
}

func (_Bfs *BfsTransactor) AsyncTouch(handler func(*types.Receipt, error), opts *bind.TransactOpts, absolutePath string, fileType string) (*types.Transaction, error) {
	return _Bfs.contract.AsyncTransact(opts, handler, "touch", absolutePath, fileType)
}

// Touch is a paid mutator transaction binding the contract method 0x131ffcdd.
//
// Solidity: function touch(string absolutePath, string fileType) returns(int32)
func (_Bfs *BfsSession) Touch(absolutePath string, fileType string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Touch(&_Bfs.TransactOpts, absolutePath, fileType)
}

func (_Bfs *BfsSession) AsyncTouch(handler func(*types.Receipt, error), absolutePath string, fileType string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncTouch(handler, &_Bfs.TransactOpts, absolutePath, fileType)
}

// Touch is a paid mutator transaction binding the contract method 0x131ffcdd.
//
// Solidity: function touch(string absolutePath, string fileType) returns(int32)
func (_Bfs *BfsTransactorSession) Touch(absolutePath string, fileType string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Bfs.Contract.Touch(&_Bfs.TransactOpts, absolutePath, fileType)
}

func (_Bfs *BfsTransactorSession) AsyncTouch(handler func(*types.Receipt, error), absolutePath string, fileType string) (*types.Transaction, error) {
	return _Bfs.Contract.AsyncTouch(handler, &_Bfs.TransactOpts, absolutePath, fileType)
}

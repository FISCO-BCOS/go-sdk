// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tableManager

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

// TableInfo is an auto generated low-level Go binding around an user-defined struct.
type TableInfo struct {
	KeyOrder     uint8
	KeyColumn    string
	ValueColumns []string
}

// TableManagerABI is the input ABI used to generate the binding from.
const TableManagerABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"newColumns\",\"type\":\"string[]\"}],\"name\":\"appendColumns\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"keyField\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"valueField\",\"type\":\"string\"}],\"name\":\"createKVTable\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumKeyOrder\",\"name\":\"keyOrder\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"keyColumn\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"valueColumns\",\"type\":\"string[]\"}],\"internalType\":\"structTableInfo\",\"name\":\"tableInfo\",\"type\":\"tuple\"}],\"name\":\"createTable\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"}],\"name\":\"descWithKeyOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"enumKeyOrder\",\"name\":\"keyOrder\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"keyColumn\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"valueColumns\",\"type\":\"string[]\"}],\"internalType\":\"structTableInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"openTable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TableManager is an auto generated Go binding around a Solidity contract.
type TableManager struct {
	TableManagerCaller     // Read-only binding to the contract
	TableManagerTransactor // Write-only binding to the contract
	TableManagerFilterer   // Log filterer for contract events
}

// TableManagerCaller is an auto generated read-only Go binding around a Solidity contract.
type TableManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableManagerTransactor is an auto generated write-only Go binding around a Solidity contract.
type TableManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableManagerFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TableManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableManagerSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TableManagerSession struct {
	Contract     *TableManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableManagerCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TableManagerCallerSession struct {
	Contract *TableManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TableManagerTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TableManagerTransactorSession struct {
	Contract     *TableManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TableManagerRaw is an auto generated low-level Go binding around a Solidity contract.
type TableManagerRaw struct {
	Contract *TableManager // Generic contract binding to access the raw methods on
}

// TableManagerCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TableManagerCallerRaw struct {
	Contract *TableManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TableManagerTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TableManagerTransactorRaw struct {
	Contract *TableManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTableManager creates a new instance of TableManager, bound to a specific deployed contract.
func NewTableManager(address common.Address, backend bind.ContractBackend) (*TableManager, error) {
	contract, err := bindTableManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TableManager{TableManagerCaller: TableManagerCaller{contract: contract}, TableManagerTransactor: TableManagerTransactor{contract: contract}, TableManagerFilterer: TableManagerFilterer{contract: contract}}, nil
}

// NewTableManagerCaller creates a new read-only instance of TableManager, bound to a specific deployed contract.
func NewTableManagerCaller(address common.Address, caller bind.ContractCaller) (*TableManagerCaller, error) {
	contract, err := bindTableManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TableManagerCaller{contract: contract}, nil
}

// NewTableManagerTransactor creates a new write-only instance of TableManager, bound to a specific deployed contract.
func NewTableManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TableManagerTransactor, error) {
	contract, err := bindTableManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TableManagerTransactor{contract: contract}, nil
}

// NewTableManagerFilterer creates a new log filterer instance of TableManager, bound to a specific deployed contract.
func NewTableManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TableManagerFilterer, error) {
	contract, err := bindTableManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TableManagerFilterer{contract: contract}, nil
}

// bindTableManager binds a generic wrapper to an already deployed contract.
func bindTableManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TableManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableManager *TableManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableManager.Contract.TableManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableManager *TableManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.TableManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableManager *TableManagerRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.TableManagerTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableManager *TableManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableManager *TableManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableManager *TableManagerTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// DescWithKeyOrder is a free data retrieval call binding the contract method 0xb8764d3f.
//
// Solidity: function descWithKeyOrder(string tableName) constant returns(TableInfo)
func (_TableManager *TableManagerCaller) DescWithKeyOrder(opts *bind.CallOpts, tableName string) (TableInfo, error) {
	var (
		ret0 = new(TableInfo)
	)
	out := ret0
	err := _TableManager.contract.Call(opts, out, "descWithKeyOrder", tableName)
	return *ret0, err
}

// DescWithKeyOrder is a free data retrieval call binding the contract method 0xb8764d3f.
//
// Solidity: function descWithKeyOrder(string tableName) constant returns(TableInfo)
func (_TableManager *TableManagerSession) DescWithKeyOrder(tableName string) (TableInfo, error) {
	return _TableManager.Contract.DescWithKeyOrder(&_TableManager.CallOpts, tableName)
}

// DescWithKeyOrder is a free data retrieval call binding the contract method 0xb8764d3f.
//
// Solidity: function descWithKeyOrder(string tableName) constant returns(TableInfo)
func (_TableManager *TableManagerCallerSession) DescWithKeyOrder(tableName string) (TableInfo, error) {
	return _TableManager.Contract.DescWithKeyOrder(&_TableManager.CallOpts, tableName)
}

// OpenTable is a free data retrieval call binding the contract method 0xf23f63c9.
//
// Solidity: function openTable(string path) constant returns(address)
func (_TableManager *TableManagerCaller) OpenTable(opts *bind.CallOpts, path string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TableManager.contract.Call(opts, out, "openTable", path)
	return *ret0, err
}

// OpenTable is a free data retrieval call binding the contract method 0xf23f63c9.
//
// Solidity: function openTable(string path) constant returns(address)
func (_TableManager *TableManagerSession) OpenTable(path string) (common.Address, error) {
	return _TableManager.Contract.OpenTable(&_TableManager.CallOpts, path)
}

// OpenTable is a free data retrieval call binding the contract method 0xf23f63c9.
//
// Solidity: function openTable(string path) constant returns(address)
func (_TableManager *TableManagerCallerSession) OpenTable(path string) (common.Address, error) {
	return _TableManager.Contract.OpenTable(&_TableManager.CallOpts, path)
}

// AppendColumns is a paid mutator transaction binding the contract method 0x302baee0.
//
// Solidity: function appendColumns(string path, string[] newColumns) returns(int32)
func (_TableManager *TableManagerTransactor) AppendColumns(opts *bind.TransactOpts, path string, newColumns []string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _TableManager.contract.TransactWithResult(opts, out, "appendColumns", path, newColumns)
	return *ret0, transaction, receipt, err
}

func (_TableManager *TableManagerTransactor) AsyncAppendColumns(handler func(*types.Receipt, error), opts *bind.TransactOpts, path string, newColumns []string) (*types.Transaction, error) {
	return _TableManager.contract.AsyncTransact(opts, handler, "appendColumns", path, newColumns)
}

// AppendColumns is a paid mutator transaction binding the contract method 0x302baee0.
//
// Solidity: function appendColumns(string path, string[] newColumns) returns(int32)
func (_TableManager *TableManagerSession) AppendColumns(path string, newColumns []string) (int32, *types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.AppendColumns(&_TableManager.TransactOpts, path, newColumns)
}

func (_TableManager *TableManagerSession) AsyncAppendColumns(handler func(*types.Receipt, error), path string, newColumns []string) (*types.Transaction, error) {
	return _TableManager.Contract.AsyncAppendColumns(handler, &_TableManager.TransactOpts, path, newColumns)
}

// AppendColumns is a paid mutator transaction binding the contract method 0x302baee0.
//
// Solidity: function appendColumns(string path, string[] newColumns) returns(int32)
func (_TableManager *TableManagerTransactorSession) AppendColumns(path string, newColumns []string) (int32, *types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.AppendColumns(&_TableManager.TransactOpts, path, newColumns)
}

func (_TableManager *TableManagerTransactorSession) AsyncAppendColumns(handler func(*types.Receipt, error), path string, newColumns []string) (*types.Transaction, error) {
	return _TableManager.Contract.AsyncAppendColumns(handler, &_TableManager.TransactOpts, path, newColumns)
}

// CreateKVTable is a paid mutator transaction binding the contract method 0xb0e89adb.
//
// Solidity: function createKVTable(string tableName, string keyField, string valueField) returns(int32)
func (_TableManager *TableManagerTransactor) CreateKVTable(opts *bind.TransactOpts, tableName string, keyField string, valueField string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _TableManager.contract.TransactWithResult(opts, out, "createKVTable", tableName, keyField, valueField)
	return *ret0, transaction, receipt, err
}

func (_TableManager *TableManagerTransactor) AsyncCreateKVTable(handler func(*types.Receipt, error), opts *bind.TransactOpts, tableName string, keyField string, valueField string) (*types.Transaction, error) {
	return _TableManager.contract.AsyncTransact(opts, handler, "createKVTable", tableName, keyField, valueField)
}

// CreateKVTable is a paid mutator transaction binding the contract method 0xb0e89adb.
//
// Solidity: function createKVTable(string tableName, string keyField, string valueField) returns(int32)
func (_TableManager *TableManagerSession) CreateKVTable(tableName string, keyField string, valueField string) (int32, *types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.CreateKVTable(&_TableManager.TransactOpts, tableName, keyField, valueField)
}

func (_TableManager *TableManagerSession) AsyncCreateKVTable(handler func(*types.Receipt, error), tableName string, keyField string, valueField string) (*types.Transaction, error) {
	return _TableManager.Contract.AsyncCreateKVTable(handler, &_TableManager.TransactOpts, tableName, keyField, valueField)
}

// CreateKVTable is a paid mutator transaction binding the contract method 0xb0e89adb.
//
// Solidity: function createKVTable(string tableName, string keyField, string valueField) returns(int32)
func (_TableManager *TableManagerTransactorSession) CreateKVTable(tableName string, keyField string, valueField string) (int32, *types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.CreateKVTable(&_TableManager.TransactOpts, tableName, keyField, valueField)
}

func (_TableManager *TableManagerTransactorSession) AsyncCreateKVTable(handler func(*types.Receipt, error), tableName string, keyField string, valueField string) (*types.Transaction, error) {
	return _TableManager.Contract.AsyncCreateKVTable(handler, &_TableManager.TransactOpts, tableName, keyField, valueField)
}

// CreateTable is a paid mutator transaction binding the contract method 0x75b14eea.
//
// Solidity: function createTable(string path, TableInfo tableInfo) returns(int32)
func (_TableManager *TableManagerTransactor) CreateTable(opts *bind.TransactOpts, path string, tableInfo TableInfo) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _TableManager.contract.TransactWithResult(opts, out, "createTable", path, tableInfo)
	return *ret0, transaction, receipt, err
}

func (_TableManager *TableManagerTransactor) AsyncCreateTable(handler func(*types.Receipt, error), opts *bind.TransactOpts, path string, tableInfo TableInfo) (*types.Transaction, error) {
	return _TableManager.contract.AsyncTransact(opts, handler, "createTable", path, tableInfo)
}

// CreateTable is a paid mutator transaction binding the contract method 0x75b14eea.
//
// Solidity: function createTable(string path, TableInfo tableInfo) returns(int32)
func (_TableManager *TableManagerSession) CreateTable(path string, tableInfo TableInfo) (int32, *types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.CreateTable(&_TableManager.TransactOpts, path, tableInfo)
}

func (_TableManager *TableManagerSession) AsyncCreateTable(handler func(*types.Receipt, error), path string, tableInfo TableInfo) (*types.Transaction, error) {
	return _TableManager.Contract.AsyncCreateTable(handler, &_TableManager.TransactOpts, path, tableInfo)
}

// CreateTable is a paid mutator transaction binding the contract method 0x75b14eea.
//
// Solidity: function createTable(string path, TableInfo tableInfo) returns(int32)
func (_TableManager *TableManagerTransactorSession) CreateTable(path string, tableInfo TableInfo) (int32, *types.Transaction, *types.Receipt, error) {
	return _TableManager.Contract.CreateTable(&_TableManager.TransactOpts, path, tableInfo)
}

func (_TableManager *TableManagerTransactorSession) AsyncCreateTable(handler func(*types.Receipt, error), path string, tableInfo TableInfo) (*types.Transaction, error) {
	return _TableManager.Contract.AsyncCreateTable(handler, &_TableManager.TransactOpts, path, tableInfo)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crud

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

// Condition is an auto generated low-level Go binding around an user-defined struct.
type Condition struct {
	Op    uint8
	Field string
	Value string
}

// Entry is an auto generated low-level Go binding around an user-defined struct.
type Entry struct {
	Key    string
	Fields []string
}

// Limit is an auto generated low-level Go binding around an user-defined struct.
type Limit struct {
	Offset uint32
	Count  uint32
}

// UpdateField is an auto generated low-level Go binding around an user-defined struct.
type UpdateField struct {
	ColumnName string
	Value      string
}

// TableABI is the input ABI used to generate the binding from.
const TableABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumConditionOP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"}],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"fields\",\"type\":\"string[]\"}],\"internalType\":\"structEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"insert\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConditionOP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"offset\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"internalType\":\"structLimit\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConditionOP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"offset\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"internalType\":\"structLimit\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"select\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"fields\",\"type\":\"string[]\"}],\"internalType\":\"structEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"select\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"fields\",\"type\":\"string[]\"}],\"internalType\":\"structEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"columnName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structUpdateField[]\",\"name\":\"updateFields\",\"type\":\"tuple[]\"}],\"name\":\"update\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumConditionOP\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"offset\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"internalType\":\"structLimit\",\"name\":\"limit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"columnName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structUpdateField[]\",\"name\":\"updateFields\",\"type\":\"tuple[]\"}],\"name\":\"update\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Table is an auto generated Go binding around a Solidity contract.
type Table struct {
	TableCaller     // Read-only binding to the contract
	TableTransactor // Write-only binding to the contract
	TableFilterer   // Log filterer for contract events
}

// TableCaller is an auto generated read-only Go binding around a Solidity contract.
type TableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableTransactor is an auto generated write-only Go binding around a Solidity contract.
type TableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TableSession struct {
	Contract     *Table            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TableCallerSession struct {
	Contract *TableCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TableTransactorSession struct {
	Contract     *TableTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableRaw is an auto generated low-level Go binding around a Solidity contract.
type TableRaw struct {
	Contract *Table // Generic contract binding to access the raw methods on
}

// TableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TableCallerRaw struct {
	Contract *TableCaller // Generic read-only contract binding to access the raw methods on
}

// TableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TableTransactorRaw struct {
	Contract *TableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTable creates a new instance of Table, bound to a specific deployed contract.
func NewTable(address common.Address, backend bind.ContractBackend) (*Table, error) {
	contract, err := bindTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Table{TableCaller: TableCaller{contract: contract}, TableTransactor: TableTransactor{contract: contract}, TableFilterer: TableFilterer{contract: contract}}, nil
}

// NewTableCaller creates a new read-only instance of Table, bound to a specific deployed contract.
func NewTableCaller(address common.Address, caller bind.ContractCaller) (*TableCaller, error) {
	contract, err := bindTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TableCaller{contract: contract}, nil
}

// NewTableTransactor creates a new write-only instance of Table, bound to a specific deployed contract.
func NewTableTransactor(address common.Address, transactor bind.ContractTransactor) (*TableTransactor, error) {
	contract, err := bindTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TableTransactor{contract: contract}, nil
}

// NewTableFilterer creates a new log filterer instance of Table, bound to a specific deployed contract.
func NewTableFilterer(address common.Address, filterer bind.ContractFilterer) (*TableFilterer, error) {
	contract, err := bindTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TableFilterer{contract: contract}, nil
}

// bindTable binds a generic wrapper to an already deployed contract.
func bindTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Table *TableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Table.Contract.TableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Table *TableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Table.Contract.TableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Table *TableRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Table.Contract.TableTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Table *TableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Table.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Table *TableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Table.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Table *TableTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Table.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Count is a free data retrieval call binding the contract method 0xda46335f.
//
// Solidity: function count([]Condition conditions) constant returns(uint32)
func (_Table *TableCaller) Count(opts *bind.CallOpts, conditions []Condition) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Table.contract.Call(opts, out, "count", conditions)
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0xda46335f.
//
// Solidity: function count([]Condition conditions) constant returns(uint32)
func (_Table *TableSession) Count(conditions []Condition) (uint32, error) {
	return _Table.Contract.Count(&_Table.CallOpts, conditions)
}

// Count is a free data retrieval call binding the contract method 0xda46335f.
//
// Solidity: function count([]Condition conditions) constant returns(uint32)
func (_Table *TableCallerSession) Count(conditions []Condition) (uint32, error) {
	return _Table.Contract.Count(&_Table.CallOpts, conditions)
}

// Select is a free data retrieval call binding the contract method 0x2a922441.
//
// Solidity: function select([]Condition conditions, Limit limit) constant returns([]Entry)
func (_Table *TableCaller) Select(opts *bind.CallOpts, conditions []Condition, limit Limit) ([]Entry, error) {
	var (
		ret0 = new([]Entry)
	)
	out := ret0
	err := _Table.contract.Call(opts, out, "select", conditions, limit)
	return *ret0, err
}

// Select is a free data retrieval call binding the contract method 0x2a922441.
//
// Solidity: function select([]Condition conditions, Limit limit) constant returns([]Entry)
func (_Table *TableSession) Select(conditions []Condition, limit Limit) ([]Entry, error) {
	return _Table.Contract.Select(&_Table.CallOpts, conditions, limit)
}

// Select is a free data retrieval call binding the contract method 0x2a922441.
//
// Solidity: function select([]Condition conditions, Limit limit) constant returns([]Entry)
func (_Table *TableCallerSession) Select(conditions []Condition, limit Limit) ([]Entry, error) {
	return _Table.Contract.Select(&_Table.CallOpts, conditions, limit)
}

// Select0 is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string key) constant returns(Entry)
func (_Table *TableCaller) Select0(opts *bind.CallOpts, key string) (Entry, error) {
	var (
		ret0 = new(Entry)
	)
	out := ret0
	err := _Table.contract.Call(opts, out, "select0", key)
	return *ret0, err
}

// Select0 is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string key) constant returns(Entry)
func (_Table *TableSession) Select0(key string) (Entry, error) {
	return _Table.Contract.Select0(&_Table.CallOpts, key)
}

// Select0 is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string key) constant returns(Entry)
func (_Table *TableCallerSession) Select0(key string) (Entry, error) {
	return _Table.Contract.Select0(&_Table.CallOpts, key)
}

// Insert is a paid mutator transaction binding the contract method 0x5c6e105f.
//
// Solidity: function insert(Entry entry) returns(int32)
func (_Table *TableTransactor) Insert(opts *bind.TransactOpts, entry Entry) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Table.contract.TransactWithResult(opts, out, "insert", entry)
	return *ret0, transaction, receipt, err
}

func (_Table *TableTransactor) AsyncInsert(handler func(*types.Receipt, error), opts *bind.TransactOpts, entry Entry) (*types.Transaction, error) {
	return _Table.contract.AsyncTransact(opts, handler, "insert", entry)
}

// Insert is a paid mutator transaction binding the contract method 0x5c6e105f.
//
// Solidity: function insert(Entry entry) returns(int32)
func (_Table *TableSession) Insert(entry Entry) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Insert(&_Table.TransactOpts, entry)
}

func (_Table *TableSession) AsyncInsert(handler func(*types.Receipt, error), entry Entry) (*types.Transaction, error) {
	return _Table.Contract.AsyncInsert(handler, &_Table.TransactOpts, entry)
}

// Insert is a paid mutator transaction binding the contract method 0x5c6e105f.
//
// Solidity: function insert(Entry entry) returns(int32)
func (_Table *TableTransactorSession) Insert(entry Entry) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Insert(&_Table.TransactOpts, entry)
}

func (_Table *TableTransactorSession) AsyncInsert(handler func(*types.Receipt, error), entry Entry) (*types.Transaction, error) {
	return _Table.Contract.AsyncInsert(handler, &_Table.TransactOpts, entry)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(string key) returns(int32)
func (_Table *TableTransactor) Remove(opts *bind.TransactOpts, key string) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Table.contract.TransactWithResult(opts, out, "remove", key)
	return *ret0, transaction, receipt, err
}

func (_Table *TableTransactor) AsyncRemove(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Table.contract.AsyncTransact(opts, handler, "remove", key)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(string key) returns(int32)
func (_Table *TableSession) Remove(key string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Remove(&_Table.TransactOpts, key)
}

func (_Table *TableSession) AsyncRemove(handler func(*types.Receipt, error), key string) (*types.Transaction, error) {
	return _Table.Contract.AsyncRemove(handler, &_Table.TransactOpts, key)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(string key) returns(int32)
func (_Table *TableTransactorSession) Remove(key string) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Remove(&_Table.TransactOpts, key)
}

func (_Table *TableTransactorSession) AsyncRemove(handler func(*types.Receipt, error), key string) (*types.Transaction, error) {
	return _Table.Contract.AsyncRemove(handler, &_Table.TransactOpts, key)
}

// Remove0 is a paid mutator transaction binding the contract method 0xa7260717.
//
// Solidity: function remove([]Condition conditions, Limit limit) returns(int32)
func (_Table *TableTransactor) Remove0(opts *bind.TransactOpts, conditions []Condition, limit Limit) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Table.contract.TransactWithResult(opts, out, "remove0", conditions, limit)
	return *ret0, transaction, receipt, err
}

func (_Table *TableTransactor) AsyncRemove0(handler func(*types.Receipt, error), opts *bind.TransactOpts, conditions []Condition, limit Limit) (*types.Transaction, error) {
	return _Table.contract.AsyncTransact(opts, handler, "remove0", conditions, limit)
}

// Remove0 is a paid mutator transaction binding the contract method 0xa7260717.
//
// Solidity: function remove([]Condition conditions, Limit limit) returns(int32)
func (_Table *TableSession) Remove0(conditions []Condition, limit Limit) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Remove0(&_Table.TransactOpts, conditions, limit)
}

func (_Table *TableSession) AsyncRemove0(handler func(*types.Receipt, error), conditions []Condition, limit Limit) (*types.Transaction, error) {
	return _Table.Contract.AsyncRemove0(handler, &_Table.TransactOpts, conditions, limit)
}

// Remove0 is a paid mutator transaction binding the contract method 0xa7260717.
//
// Solidity: function remove([]Condition conditions, Limit limit) returns(int32)
func (_Table *TableTransactorSession) Remove0(conditions []Condition, limit Limit) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Remove0(&_Table.TransactOpts, conditions, limit)
}

func (_Table *TableTransactorSession) AsyncRemove0(handler func(*types.Receipt, error), conditions []Condition, limit Limit) (*types.Transaction, error) {
	return _Table.Contract.AsyncRemove0(handler, &_Table.TransactOpts, conditions, limit)
}

// Update is a paid mutator transaction binding the contract method 0x41ffd75f.
//
// Solidity: function update(string key, []UpdateField updateFields) returns(int32)
func (_Table *TableTransactor) Update(opts *bind.TransactOpts, key string, updateFields []UpdateField) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Table.contract.TransactWithResult(opts, out, "update", key, updateFields)
	return *ret0, transaction, receipt, err
}

func (_Table *TableTransactor) AsyncUpdate(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, updateFields []UpdateField) (*types.Transaction, error) {
	return _Table.contract.AsyncTransact(opts, handler, "update", key, updateFields)
}

// Update is a paid mutator transaction binding the contract method 0x41ffd75f.
//
// Solidity: function update(string key, []UpdateField updateFields) returns(int32)
func (_Table *TableSession) Update(key string, updateFields []UpdateField) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Update(&_Table.TransactOpts, key, updateFields)
}

func (_Table *TableSession) AsyncUpdate(handler func(*types.Receipt, error), key string, updateFields []UpdateField) (*types.Transaction, error) {
	return _Table.Contract.AsyncUpdate(handler, &_Table.TransactOpts, key, updateFields)
}

// Update is a paid mutator transaction binding the contract method 0x41ffd75f.
//
// Solidity: function update(string key, []UpdateField updateFields) returns(int32)
func (_Table *TableTransactorSession) Update(key string, updateFields []UpdateField) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Update(&_Table.TransactOpts, key, updateFields)
}

func (_Table *TableTransactorSession) AsyncUpdate(handler func(*types.Receipt, error), key string, updateFields []UpdateField) (*types.Transaction, error) {
	return _Table.Contract.AsyncUpdate(handler, &_Table.TransactOpts, key, updateFields)
}

// Update0 is a paid mutator transaction binding the contract method 0x9924c17e.
//
// Solidity: function update([]Condition conditions, Limit limit, []UpdateField updateFields) returns(int32)
func (_Table *TableTransactor) Update0(opts *bind.TransactOpts, conditions []Condition, limit Limit, updateFields []UpdateField) (int32, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	transaction, receipt, err := _Table.contract.TransactWithResult(opts, out, "update0", conditions, limit, updateFields)
	return *ret0, transaction, receipt, err
}

func (_Table *TableTransactor) AsyncUpdate0(handler func(*types.Receipt, error), opts *bind.TransactOpts, conditions []Condition, limit Limit, updateFields []UpdateField) (*types.Transaction, error) {
	return _Table.contract.AsyncTransact(opts, handler, "update0", conditions, limit, updateFields)
}

// Update0 is a paid mutator transaction binding the contract method 0x9924c17e.
//
// Solidity: function update([]Condition conditions, Limit limit, []UpdateField updateFields) returns(int32)
func (_Table *TableSession) Update0(conditions []Condition, limit Limit, updateFields []UpdateField) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Update0(&_Table.TransactOpts, conditions, limit, updateFields)
}

func (_Table *TableSession) AsyncUpdate0(handler func(*types.Receipt, error), conditions []Condition, limit Limit, updateFields []UpdateField) (*types.Transaction, error) {
	return _Table.Contract.AsyncUpdate0(handler, &_Table.TransactOpts, conditions, limit, updateFields)
}

// Update0 is a paid mutator transaction binding the contract method 0x9924c17e.
//
// Solidity: function update([]Condition conditions, Limit limit, []UpdateField updateFields) returns(int32)
func (_Table *TableTransactorSession) Update0(conditions []Condition, limit Limit, updateFields []UpdateField) (int32, *types.Transaction, *types.Receipt, error) {
	return _Table.Contract.Update0(&_Table.TransactOpts, conditions, limit, updateFields)
}

func (_Table *TableTransactorSession) AsyncUpdate0(handler func(*types.Receipt, error), conditions []Condition, limit Limit, updateFields []UpdateField) (*types.Transaction, error) {
	return _Table.Contract.AsyncUpdate0(handler, &_Table.TransactOpts, conditions, limit, updateFields)
}

package crud

import (
	//"encoding/json"
	"fmt"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompiled"
	"github.com/ethereum/go-ethereum/common"
)

const (
	TableKeyMaxLength int = 255

	//crud precompiled contract error code
	conditionOperationUndefined int64 = -51502
	parseConditionError         int64 = -51501
	parseEntryError             int64 = -51500
)

// getErrorMessage returns the message of error code
func getErrorMessage(errorCode int64) string {
	var message string
	switch errorCode {
	case conditionOperationUndefined:
		message = "condition operation undefined"
	case parseConditionError:
		message = "parse condition error"
	case parseEntryError:
		message = "parse entry error"
	default:
		message = ""
	}
	return message
}

// errorCodeToError judges whether the error code represents an error
func errorCodeToError(errorCode int64) error {
	var errorCodeMessage string
	errorCodeMessage = precompiled.GetCommonErrorCodeMessage(errorCode)
	if errorCodeMessage != "" {
		return fmt.Errorf("error code: %v, error code message: %v", errorCode, errorCodeMessage)
	}
	errorCodeMessage = getErrorMessage(errorCode)
	if errorCodeMessage != "" {
		return fmt.Errorf("error code: %v, error code message: %v", errorCode, errorCodeMessage)
	}
	return nil
}

// CRUDService is a precompile contract service.
type Service struct {
	crud         *Table
	tableFactory *TableManager
	KVTable		 *KVTable
	crudAuth     *bind.TransactOpts
	CallOpts     *bind.CallOpts
	client       *client.Client
}

// TableFactoryPrecompileAddress is the contract address of TableFactory
var TableFactoryPrecompileAddress = common.HexToAddress("0x1002")

// CRUDPrecompileAddress is the contract address of CRUD
var CRUDPrecompileAddress = common.HexToAddress("0x1002")

var KVTablePrecompileAddress = common.HexToAddress("0x1002")

// NewCRUDService returns ptr of CRUDService
func NewCRUDService(client *client.Client) (*Service, error) {
	crudInstance, err := NewTable(CRUDPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct CRUD failed: %+v", err)
	}
	tableInstance, err := NewTableManager(TableFactoryPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	KVTableInstance, err := NewKVTable(KVTablePrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct KVTable failed: %+v", err)
	}
	auth := client.GetTransactOpts()
	callOpts := client.GetCallOpts()
	return &Service{crud: crudInstance, tableFactory: tableInstance, KVTable:KVTableInstance, crudAuth: auth, CallOpts: callOpts, client: client}, nil
}

// TableManager
func (service *Service) CreateTable(tableName string, key string, valueFields []string) (int64, error) {
	tableInfo := TableInfo{
		KeyColumn:    key,
		ValueColumns: valueFields,
	}
	_, _, receipt, err := service.tableFactory.CreateTable(service.crudAuth, tableName, tableInfo)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService CreateTable failed: %v", err)
	}

	address, err := service.tableFactory.OpenTable(service.CallOpts, tableName)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService OpenTable failed: %v", err)
	}
	fmt.Println("CreateTable address:", address)
	crudInstance, err := NewTable(address, service.client)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	service.crud = crudInstance
	return parseReturnValue(receipt, "createTable")
}

func (service *Service) AsyncCreateTable(handler func(*types.Receipt, error), tableName string, key string, valueFields []string) (*types.Transaction, error) {
	tableInfo := TableInfo{
		KeyColumn:    key,
		ValueColumns: valueFields,
	}
	return service.tableFactory.AsyncCreateTable(handler, service.crudAuth, tableName, tableInfo)
}

func (service *Service) CreateKVTable(tableName string, keyField string, valueField string) (int64, error) {
	_, _, receipt, err := service.tableFactory.CreateKVTable(service.crudAuth, tableName, keyField, valueField)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService CreateKVTable failed: %v", err)
	}

	address, err := service.tableFactory.OpenTable(service.CallOpts, tableName)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService OpenTable failed: %v", err)
	}
	fmt.Println("CreateKVTable address:", address)
	crudInstance, err := NewTable(address, service.client)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	service.crud = crudInstance
	return parseReturnValue(receipt, "createKVTable")
}

func (service *Service) AsyncCreateKVTable(handler func(*types.Receipt, error), tableName string, keyField string, valueField string) (*types.Transaction, error) {
	return service.tableFactory.AsyncCreateKVTable(handler, service.crudAuth, tableName, keyField, valueField)
}

func (service *Service) OpenTable(tableName string) (int64, error) {
	address, err := service.tableFactory.OpenTable(service.CallOpts, tableName)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService OpenTable failed: %v", err)
	}
	fmt.Println("OpenTable address:", address)
	crudInstance, err := NewTable(address, service.client)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	service.crud = crudInstance
	return 0, nil
}

func (service *Service) OpenKVTable(tableName string) (int64, error) {
	address, err := service.tableFactory.OpenTable(service.CallOpts, tableName)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService OpenTable failed: %v", err)
	}
	fmt.Println("OpenTable address:", address)
	KVTableInstance, err := NewKVTable(address, service.client)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	service.KVTable = KVTableInstance
	return 0, nil
}

func (service *Service) AppendColumns(path string, newColumns []string) (int64, error) {
	_, _, receipt, err := service.tableFactory.AppendColumns(service.crudAuth, path, newColumns)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService AppendColumns failed: %v", err)
	}
	return parseReturnValue(receipt, "createTable")
}

func (service *Service) AsyncAppendColumns(handler func(*types.Receipt, error), path string, newColumns []string) (*types.Transaction, error) {
	return service.tableFactory.AsyncAppendColumns(handler, service.crudAuth, path, newColumns)
}

func (service *Service) DescWithKeyOrder(tableName string) (TableInfo, error) {
	opts := &bind.CallOpts{From: service.crudAuth.From}
	ret0, err := service.tableFactory.DescWithKeyOrder(opts, tableName)
	if err != nil {
		return *new(TableInfo), fmt.Errorf("CRUDService DescWithKeyOrder failed: %v", err)
	}
	return ret0, err
}

// Table
func (service *Service) Insert(tableName string, entry Entry) (int64, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return -1, fmt.Errorf("the table does not exist")
	}
	if len(entry.Key) > TableKeyMaxLength {
		return -1, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	_, _, receipt, err := service.crud.Insert(service.crudAuth, entry)
	if err != nil {
		return -1, fmt.Errorf("CRUDService Insert failed: %v", err)
	}
	return parseReturnValue(receipt, "insert")
}

func (service *Service) AsyncInsert(handler func(*types.Receipt, error), tableName string, entry Entry) (*types.Transaction, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return nil, fmt.Errorf("the table does not exist")
	}
	if len(entry.Key) > TableKeyMaxLength {
		return nil, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	return service.crud.AsyncInsert(handler, service.crudAuth, entry)
}

func (service *Service) Select0(tableName string, key string) (Entry, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return Entry{}, fmt.Errorf("the table does not exist")
	}
	if len(key) > TableKeyMaxLength {
		return Entry{}, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	entry, err := service.crud.Select0(service.CallOpts, key)
	if err != nil {
		return Entry{}, fmt.Errorf("CRUDService Select failed: %v", err)
	}
	return entry, nil
}

func (service *Service) Select(tableName string, conditions []Condition, limit Limit) ([]Entry, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return nil, fmt.Errorf("the table does not exist")
	}
	entries, err := service.crud.Select(service.CallOpts, conditions, limit)
	if err != nil {
		return nil, fmt.Errorf("CRUDService Select failed: %v", err)
	}
	return entries, nil
}

func (service *Service) Update(tableName string, key string, updateFields []UpdateField) (int64, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return -1, fmt.Errorf("the table does not exist")
	}
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	_, _, receipt, err := service.crud.Update(service.crudAuth, key, updateFields)
	if err != nil {
		return -1, fmt.Errorf("CRUDService Update failed: %v", err)
	}
	return parseReturnValue(receipt, "update")
}

func (service *Service) Update0(tableName string, conditions []Condition, limit Limit, updateFields []UpdateField) (int64, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return -1, fmt.Errorf("the table does not exist")
	}
	_, _, receipt, err := service.crud.Update0(service.crudAuth, conditions, limit, updateFields)
	if err != nil {
		return -1, fmt.Errorf("CRUDService Update failed: %v", err)
	}
	return parseReturnValue(receipt, "update0")
}

func (service *Service) AsyncUpdate(handler func(*types.Receipt, error), tableName string, key string, updateFields []UpdateField) (*types.Transaction, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return nil, fmt.Errorf("the table does not exist")
	}
	if len(key) > TableKeyMaxLength {
		return nil, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	return service.crud.AsyncUpdate(handler, service.crudAuth, key, updateFields)
}

func (service *Service) AsyncUpdate0(handler func(*types.Receipt, error), tableName string, conditions []Condition, limit Limit, updateFields []UpdateField) (*types.Transaction, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return nil, fmt.Errorf("the table does not exist")
	}
	return service.crud.AsyncUpdate0(handler, service.crudAuth, conditions, limit, updateFields)
}

func (service *Service) Remove(tableName string, key string) (int64, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return -1, fmt.Errorf("the table does not exist")
	}
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	_, _, receipt, err := service.crud.Remove(service.crudAuth, key)
	if err != nil {
		return -1, fmt.Errorf("CRUDService Remove failed: %v", err)
	}
	return parseReturnValue(receipt, "remove")
}

func (service *Service) Remove0(tableName string, conditions []Condition, limit Limit) (int64, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return -1, fmt.Errorf("the table does not exist")
	}
	_, _, receipt, err := service.crud.Remove0(service.crudAuth, conditions, limit)
	if err != nil {
		return -1, fmt.Errorf("CRUDService Remove failed: %v", err)
	}
	return parseReturnValue(receipt, "remove0")
}

func (service *Service) AsyncRemove(handler func(*types.Receipt, error), tableName string, key string) (*types.Transaction, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return nil, fmt.Errorf("the table does not exist")
	}
	if len(key) > TableKeyMaxLength {
		return nil, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	return service.crud.AsyncRemove(handler, service.crudAuth, key)
}

func (service *Service) AsyncRemove0(handler func(*types.Receipt, error), tableName string, conditions []Condition, limit Limit) (*types.Transaction, error) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		return nil, fmt.Errorf("the table does not exist")
	}
	return service.crud.AsyncRemove0(handler, service.crudAuth, conditions, limit)
}

// KVTable
func (service *Service) Set(tableName string, key string, value string) (int64, error) {
	_, err := service.OpenKVTable(tableName)
	if err !=nil{
		return -1, fmt.Errorf("the table does not exist")
	}
	_, _, receipt, err := service.KVTable.Set(service.crudAuth, key, value)
	if err != nil {
		return -1, fmt.Errorf("CRUDService Set failed: %v", err)
	}
	return parseReturnValue(receipt, "set")
}

func (service *Service) Get(tableName string, key string) (bool, string, error) {
	_, err := service.OpenKVTable(tableName)
	if err !=nil{
		return false, "", fmt.Errorf("the table does not exist")
	}
	ret0, ret1, err := service.KVTable.Get(service.CallOpts, key)
	if err != nil {
		return false, "", fmt.Errorf("KVTable Get failed: %v", err)
	}
	return ret0, ret1, err
}

func parseReturnValue(receipt *types.Receipt, name string) (int64, error) {
	errorMessage := receipt.GetErrorMessage()
	if errorMessage != "" {
		return int64(receipt.GetStatus()), fmt.Errorf("receipt.Status err: %v", errorMessage)
	}
	var bigNum *big.Int
	var err error
	if name == "createTable" {
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
	} else {
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
	}
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("parseReturnValue failed, err: %v", err)
	}
	errorCode, err := precompiled.BigIntToInt64(bigNum)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("parseReturnValue failed, err: %v", err)
	}
	return errorCode, errorCodeToError(errorCode)
}
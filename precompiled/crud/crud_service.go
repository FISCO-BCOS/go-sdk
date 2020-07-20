package crud

import (
	"encoding/json"
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
	crud         *Crud
	tableFactory *TableFactory
	crudAuth     *bind.TransactOpts
	client       *client.Client
}

// TableFactoryPrecompileAddress is the contract address of TableFactory
var TableFactoryPrecompileAddress = common.HexToAddress("0x0000000000000000000000000000000000001001")

// CRUDPrecompileAddress is the contract address of CRUD
var CRUDPrecompileAddress = common.HexToAddress("0x0000000000000000000000000000000000001002")

// NewCRUDService returns ptr of CRUDService
func NewCRUDService(client *client.Client) (*Service, error) {
	crudInstance, err := NewCrud(CRUDPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct CRUD failed: %+v", err)
	}
	tableInstance, err := NewTableFactory(TableFactoryPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	auth := client.GetTransactOpts()
	return &Service{crud: crudInstance, tableFactory: tableInstance, crudAuth: auth, client: client}, nil
}

func (service *Service) CreateTable(tableName string, key string, valueFields string) (int64, error) {
	tx, err := service.tableFactory.CreateTable(service.crudAuth, tableName, key, valueFields)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CRUDService CreateTable failed: %v", err)
	}
	return handleReceipt(service.client, tx, "createTable")

}

// Insert entry
func (service *Service) Insert(tableName string, key string, entry *Entry) (int64, error) {
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	// change to string
	entryJSON, err := json.MarshalIndent(entry.GetFields(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change entry to json struct failed: %v", err)
	}
	tx, err := service.crud.Insert(service.crudAuth, tableName, key, string(entryJSON[:]), "")
	if err != nil {
		return -1, fmt.Errorf("CRUDService Insert failed: %v", err)
	}
	return handleReceipt(service.client, tx, "insert")
}

// Update entry
func (service *Service) Update(tableName string, key string, entry *Entry, condition *Condition) (int64, error) {
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	// change to string
	entryJSON, err := json.MarshalIndent(entry.GetFields(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change entry to json struct failed: %v", err)
	}
	conditionJSON, err := json.MarshalIndent(condition.GetConditions(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change condition to json struct failed: %v", err)
	}

	tx, err := service.crud.Update(service.crudAuth, tableName, key, string(entryJSON[:]), string(conditionJSON[:]), "")
	if err != nil {
		return -1, fmt.Errorf("CRUDService Update failed: %v", err)
	}
	return handleReceipt(service.client, tx, "update")
}

func (service *Service) Remove(tableName string, key string, condition *Condition) (int64, error) {
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	conditionJSON, err := json.MarshalIndent(condition.GetConditions(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change condition to json struct failed: %v", err)
	}

	tx, err := service.crud.Remove(service.crudAuth, tableName, key, string(conditionJSON[:]), "")
	if err != nil {
		return -1, fmt.Errorf("CRUDService Remove failed: %v", err)
	}
	return handleReceipt(service.client, tx, "remove")
}

// Select entry
func (service *Service) Select(tableName string, key string, condition *Condition) ([]map[string]string, error) {
	if len(key) > TableKeyMaxLength {
		return nil, fmt.Errorf("the value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	conditionJSON, err := json.MarshalIndent(condition.GetConditions(), "", "\t")
	if err != nil {
		return nil, fmt.Errorf("change condition to json struct failed: %v", err)
	}
	opts := &bind.CallOpts{From: service.crudAuth.From}
	result, err := service.crud.Select(opts, tableName, key, string(conditionJSON[:]), "")
	if err != nil {
		return nil, fmt.Errorf("CRUDService Select failed: %v", err)
	}
	// unmarshal result
	var results []map[string]string
	if err := json.Unmarshal([]byte(result), &results); err != nil {
		return nil, fmt.Errorf("CRUDService: Unmarshal the Select result failed: %v", err)
	}
	return results, nil
}

// Desc is used for Table
func (service *Service) Desc(userTableName string) (string, string, error) {
	opts := &bind.CallOpts{From: service.crudAuth.From}
	keyField, valueField, err := service.crud.Desc(opts, userTableName)
	if err != nil {
		return "", "", fmt.Errorf("desc failed, select table error: %v", err)
	}
	return keyField, valueField, nil
}

func handleReceipt(c *client.Client, tx *types.Transaction, name string) (int64, error) {
	// wait for the mining
	receipt, err := c.WaitMined(tx)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("CrudService wait for the transaction receipt failed, err: %v", err)
	}
	status := receipt.GetStatus()
	if types.Success != status {
		return int64(status), fmt.Errorf(types.GetStatusMessage(status))
	}
	var bigNum *big.Int
	if name == "createTable" {
		bigNum, err = precompiled.ParseBigIntFromOutput(TableFactoryABI, name, receipt)
	} else {
		bigNum, err = precompiled.ParseBigIntFromOutput(CrudABI, name, receipt)
	}
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("handleReceipt failed, err: %v", err)
	}
	errorCode, err := precompiled.BigIntToInt64(bigNum)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("handleReceipt failed, err: %v", err)
	}
	return errorCode, errorCodeToError(errorCode)
}

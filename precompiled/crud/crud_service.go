package crud

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	TableKeyMaxLength int    = 255
	SysTable          string = "_sys_tables_"
	UserTablePrefix   string = "u_"
	// UserTablePrefixOld   string = "_user_"
)

// CRUDService is a precompile contract service.
type CRUDService struct {
	crud         *Crud
	tableFactory *TableFactory
	crudAuth     *bind.TransactOpts
	client       *client.Client
}

// TableFactoryPrecompileAddress is the contract address of TableFactory
var TableFactoryPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001001")

// CRUDPrecompileAddress is the contract address of CRUD
var CRUDPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001002")

// NewCRUDService returns ptr of CRUDService
func NewCRUDService(client *client.Client, privateKey *ecdsa.PrivateKey) (*CRUDService, error) {
	crudInstance, err := NewCrud(CRUDPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct CRUD failed: %+v", err)
	}
	tableInstance, err := NewTableFactory(TableFactoryPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct TableFactor failed: %+v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = big.NewInt(30000000)
	return &CRUDService{crud: crudInstance, tableFactory: tableInstance, crudAuth: auth, client: client}, nil
}

func (service *CRUDService) CreateTable(tableName string, key string, valueFields string) (int, error) {
	tx, err := service.tableFactory.CreateTable(service.crudAuth, tableName, key, valueFields)
	if err != nil {
		return -1, fmt.Errorf("CRUDService CreateTable failed: %v", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), service.client, tx)
	if err != nil {
		return -1, fmt.Errorf("CRUDService wait for the transaction receipt failed: %v", err)
	}
	// handle receipt
	return handleReceipt(receipt)
}

// Insert entry
func (service *CRUDService) Insert(tableName string, key string, entry *Entry) (int, error) {
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
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
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), service.client, tx)
	if err != nil {
		return -1, fmt.Errorf("CRUDService wait for the transaction receipt failed: %v", err)
	}
	// handle receipt
	return handleReceipt(receipt)
}

// Update entry
func (service *CRUDService) Update(tableName string, key string, entry *Entry, condition *Condition) (int, error) {
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
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
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), service.client, tx)
	if err != nil {
		return -1, fmt.Errorf("CRUDService wait for the transaction receipt failed: %v", err)
	}
	// handle receipt
	return handleReceipt(receipt)
}

func (service *CRUDService) Remove(tableName string, key string, condition *Condition) (int, error) {
	if len(key) > TableKeyMaxLength {
		return -1, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	conditionJSON, err := json.MarshalIndent(condition.GetConditions(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change condition to json struct failed: %v", err)
	}

	tx, err := service.crud.Remove(service.crudAuth, tableName, key, string(conditionJSON[:]), "")
	if err != nil {
		return -1, fmt.Errorf("CRUDService Remove failed: %v", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), service.client, tx)
	if err != nil {
		return -1, fmt.Errorf("CRUDService wait for the transaction receipt failed: %v", err)
	}
	// handle receipt
	return handleReceipt(receipt)
}

// Select entry
func (service *CRUDService) Select(tableName string, key string, condition *Condition) ([]map[string]string, error) {
	if len(key) > TableKeyMaxLength {
		return nil, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
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
func (service *CRUDService) Desc(userTableName string) (string, string, error) {
	tableName := SysTable
	key := UserTablePrefix + userTableName
	condition := NewCondition()
	userTable, err := service.Select(tableName, key, condition)
	if err != nil {
		return "", "", fmt.Errorf("select table failed: %v", err)
	}
	if len(userTable) == 0 {
		return "", "", fmt.Errorf("the table %s does not exist", tableName)
	}
	return userTable[0]["key_field"], userTable[0]["value_field"], nil
}

func handleReceipt(receipt *types.Receipt) (int, error) {
	status := receipt.GetStatus()
	if types.Success != status {
		return -1, fmt.Errorf(types.GetStatusMessage(status))
	}
	output := receipt.GetOutput()
	if output != "0x" {
		i := new(big.Int)
		var flag bool
		i, flag = i.SetString(output[2:], 16)
		if !flag {
			return -1, fmt.Errorf("handleReceipt: convert receipt output to int failed")
		}
		return int(i.Uint64()), nil
	}
	return -1, fmt.Errorf("Transaction is handled failure")
}

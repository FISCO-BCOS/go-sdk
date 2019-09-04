package crud

import (
	"fmt"
	"crypto/ecdsa"
	"math/big"
	"context"
	"encoding/json"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
)

const (
	TableKeyMaxLength int = 255
	SysTable string = "_sys_tables_"
	UserTablePrefix string = "_user_";
)

// CRUDService is a precompile contract service.
type CRUDService struct {
	crud *Crud
	tableFactory *TableFactory
	crudAuth *bind.TransactOpts
	client *client.Client
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
    return &CRUDService{crud:crudInstance, tableFactory:tableInstance, crudAuth:auth, client: client}, nil
}

// CreateTable returns the status of the creating that 0 represents succeed 
func (service *CRUDService) CreateTable(table *Table) (int, error) {
	tx, err := service.tableFactory.CreateTable(service.crudAuth, table.GetTableName(), table.GetKey(), table.GetValueFields())
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
func (service *CRUDService) Insert(table *Table, entry *Entry) (int, error) {
    if (len(table.GetKey()) > TableKeyMaxLength) {
		return -1, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
    // change to string
	entryJSON, err := json.MarshalIndent(entry.GetFields(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change entry to json struct failed: %v", err)
	}

	tx, err := service.crud.Insert(service.crudAuth, table.GetTableName(), table.GetKey(), string(entryJSON[:]), table.GetOptional())
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
func (service *CRUDService) Update(table *Table, entry *Entry, condition *Condition) (int, error) {
    if (len(table.GetKey()) > TableKeyMaxLength) {
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

	tx, err := service.crud.Update(service.crudAuth, table.GetTableName(), table.GetKey(), string(entryJSON[:]), string(conditionJSON[:]), table.GetOptional())
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

// Remove entry
func (service *CRUDService) Remove(table *Table, condition *Condition) (int, error) {
    if (len(table.GetKey()) > TableKeyMaxLength) {
		return -1, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	conditionJSON, err := json.MarshalIndent(condition.GetConditions(), "", "\t")
	if err != nil {
		return -1, fmt.Errorf("change condition to json struct failed: %v", err)
	}

	tx, err := service.crud.Remove(service.crudAuth, table.GetTableName(), table.GetKey(), string(conditionJSON[:]), table.GetOptional())
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
func (service *CRUDService) Select(table *Table, condition *Condition) ([]map[string]string, error) {
    if (len(table.GetKey()) > TableKeyMaxLength) {
		return nil, fmt.Errorf("The value of the table key exceeds the maximum limit( %d )", TableKeyMaxLength)
	}
	conditionJSON, err := json.MarshalIndent(condition.GetConditions(), "", "\t")
	if err != nil {
		return nil, fmt.Errorf("change condition to json struct failed: %v", err)
	}
	opts := &bind.CallOpts{From: service.crudAuth.From}
	result, err := service.crud.Select(opts, table.GetTableName(), table.GetKey(), string(conditionJSON[:]), table.GetOptional())
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
func (service *CRUDService) Desc(tableName string) (*Table, error) {
	table := &Table{}
	table.SetTableName(SysTable)
	table.SetKey(UserTablePrefix + tableName)
	condition := table.GetCondition()
	userTable, err := service.Select(table, condition)
	if err != nil {
        return nil, fmt.Errorf("select table failed: %v", err)
	}
	tableInfo := &Table{}
	if len(userTable) == 0 {
        return nil, fmt.Errorf("The table %s does not exist", tableName)
	}
	tableInfo.SetTableName(tableName)
	tableInfo.SetKey(userTable[0]["key_field"])
	tableInfo.SetValueFields(userTable[0]["value_field"])
	return table, nil
}

func handleReceipt(receipt *types.Receipt) (int, error) {
	status := receipt.GetStatus()
	if "0x0" != status {
		return -1, fmt.Errorf(common.GetStatusMessage(status))
	} 
	output := receipt.GetOutput()
	if output != "0x" {
		i := new(big.Int)
		var flag bool
		i,flag = i.SetString(output[2:len(output)], 16)
		if flag == false {
			return -1, fmt.Errorf("handleReceipt: convert receipt output to int failed")
		}
		return int(i.Uint64()), nil
	}
	return -1, fmt.Errorf("Transaction is handled failure")
}
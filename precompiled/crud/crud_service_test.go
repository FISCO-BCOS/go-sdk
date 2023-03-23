package crud

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompiled"
)

const (
	tableName         = "t_test9"
	tableNameForAsync = "t_test_async9"
	key               = "name"
	keyAsync          = "name_async"
	timeout           = 1 * time.Second
)

var valueFields = []string{"item_name"}
var valueFields_update = []string{"item_name_update"}

var (
	service *Service
	channel = make(chan int)
)

func getClient(t *testing.T) *client.Client {
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &client.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
	c, err := client.DialContext(context.Background(), config)
	if err != nil {
		t.Fatalf("Dial to %s:%d failed of %v", config.Host, config.Port, err)
	}
	return c
}

func getService(t *testing.T) {
	c := getClient(t)
	newService, err := NewCRUDService(c)
	if err != nil {
		t.Fatalf("init CrudService failed: %+v", err)
	}
	service = newService
}

func TestMain(m *testing.M) {
	getService(&testing.T{})
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateTable(t *testing.T) {
	result, err := service.CreateTable(tableName, key, valueFields)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	if result != 0 {
		t.Fatalf("TestCreateTable failed, the result \"%v\" is inconsistent with \"0\"", result)
	}
	t.Logf("result: %d\n", result)
}

func TestAsyncCreateTable(t *testing.T) {
	handler := func(receipt *types.Receipt, err error) {
		if err != nil {
			t.Fatalf("receive receipt failed, %v\n", err)
		}
		var bigNum *big.Int
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
		if err != nil {
			t.Fatalf("parseReturnValue failed, err: %v\n", err)
		}
		result, err := precompiled.BigIntToInt64(bigNum)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if result != 0 {
			t.Fatalf("TestAsyncCreateTable failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncCreateTable(handler, tableNameForAsync, key, valueFields)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestInsert(t *testing.T) {
	var insertResults int
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	for i := 1; i <= 5; i++ {
		_, receipt, err := service.crud.Insert(service.crudAuth, entry)
		if err != nil {
			t.Fatalf("insert table failed: %v", err)
		}
		insertResults += receipt.Status
	}
	if insertResults != 0 {
		t.Fatalf("TestInsert failed, the insertResults \"%v\" is inconsistent with \"5\"", insertResults)
	}
	t.Logf("insertResults: %d\n", insertResults)
}

func TestAsyncInsert(t *testing.T) {
	handler := func(receipt *types.Receipt, err error) {
		if err != nil {
			t.Fatalf("receive receipt failed, %v\n", err)
		}
		var bigNum *big.Int
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
		if err != nil {
			t.Fatalf("parseReturnValue failed, err: %v\n", err)
		}
		result, err := precompiled.BigIntToInt64(bigNum)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if result != 1 {
			t.Fatalf("TestAsyncInsert failed, the result \"%v\" is inconsistent with \"1\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	entry := Entry{
		Key:    keyAsync,
		Fields: valueFields,
	}
	_, err := service.crud.AsyncInsert(handler, service.crudAuth, entry)
	if err != nil {
		t.Fatalf("insert table failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestSelect(t *testing.T) {

	resultSelect, err := service.crud.Select0(service.CallOpts, key)
	if err != nil {
		t.Fatalf("select table failed: %v", err)
	}
	if resultSelect.Fields[0] != valueFields[0] {
		t.Fatalf("TestSelect failed, the result of resultSelect \"%v\" is not inconsistent", resultSelect.Fields[0])
	}
	t.Logf("resultSelect :\n")

	for i := 0; i < len(resultSelect.Fields); i++ {
		t.Logf("resultSelect[%d]'s item_name is：%s\n", i, resultSelect.Fields[i])
	}
}

func TestUpdate(t *testing.T) {
	updateField := UpdateField{
		ColumnName: valueFields[0],
		Value:      valueFields_update[0],
	}
	var updateFieldList []UpdateField
	updateFieldList = append(updateFieldList, updateField)
	_, receipt, err := service.crud.Update(service.crudAuth, key, updateFieldList)
	if err != nil {
		t.Fatalf("update table failed: %v", err)
	}
	if receipt.Status != 0 {
		t.Fatalf("TestUpdate failed, the updateResult ")
	}
	resultSelect, err := service.crud.Select0(service.CallOpts, key)
	if err != nil {
		t.Fatalf("select table failed: %v", err)
	}
	if resultSelect.Fields[0] != valueFields_update[0] {
		t.Fatalf("TestUpdate failed, the result of resultUpdate \"%v\" is not inconsistent", resultSelect.Fields[0])
	}
	t.Logf("updateResult: %s", resultSelect.Fields[0])
}

func TestAsyncUpdate(t *testing.T) {
	handler := func(receipt *types.Receipt, err error) {
		if err != nil {
			t.Fatalf("receive receipt failed, %v\n", err)
		}
		var bigNum *big.Int
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
		if err != nil {
			t.Fatalf("parseReturnValue failed, err: %v\n", err)
		}
		result, err := precompiled.BigIntToInt64(bigNum)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if result != 1 {
			t.Fatalf("TestAsyncUpdate failed, the result \"%v\" is inconsistent with \"1\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	updateField := UpdateField{
		ColumnName: valueFields[0],
		Value:      valueFields_update[0],
	}
	var updateFieldList []UpdateField
	updateFieldList = append(updateFieldList, updateField)
	_, err := service.crud.AsyncUpdate(handler, service.crudAuth, keyAsync, updateFieldList)
	if err != nil {
		t.Fatalf("update table failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestRemove(t *testing.T) {
	_, receipt, err := service.crud.Remove0(service.crudAuth, key)
	if err != nil {
		t.Fatalf("remove table failed: %v", err)
	}
	if receipt.Status != 0 {
		t.Fatalf("TestRemove failed, the removeResult \"%v\" is not inconsistent with \"0\"", receipt.Status)
	}
	t.Logf("removeResult: %d\n", receipt.Status)
}

func TestAsyncRemove(t *testing.T) {
	handler := func(receipt *types.Receipt, err error) {
		if err != nil {
			t.Fatalf("receive receipt failed, %v\n", err)
		}
		var bigNum *big.Int
		bigNum, err = precompiled.ParseBigIntFromOutput(receipt)
		if err != nil {
			t.Fatalf("parseReturnValue failed, err: %v\n", err)
		}
		result, err := precompiled.BigIntToInt64(bigNum)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if result != 1 {
			t.Fatalf("TestAsyncRemove failed, the result \"%v\" is inconsistent with \"1\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.crud.AsyncRemove0(handler, service.crudAuth, keyAsync)
	if err != nil {
		t.Fatalf("remove data failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestDesc(t *testing.T) {
	tableInfo, err := service.tableFactory.Desc(service.CallOpts, tableName)
	if err != nil {
		t.Fatalf("query table info by tableName failed: %v", err)
	}
	if tableInfo.KeyColumn != key {
		t.Fatalf("TestDesc failed, the keyField \"%v\" is not inconsistent with \"name\"", tableInfo.KeyColumn)
	}
	if tableInfo.ValueColumns[0] != valueFields[0] {
		t.Fatalf("TestDesc failed, the valueField \"%v\" is not inconsistent with \"item_id,item_name\"", tableInfo.ValueColumns[0])
	}
	t.Logf("keyFiled is：%s\n", tableInfo.KeyColumn)
	t.Logf("valueField is：%s\n", tableInfo.ValueColumns[0])
}

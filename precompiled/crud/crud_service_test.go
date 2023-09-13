package crud

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"testing"
	"time"
	// "fmt"
    "math/rand"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompiled"
)

const (
	tableName         = "t_test0"
	tableNameForAsync = "t_test_async0"
	KVTableName       = "t_test1"
	KVTableNameForAsync = "t_test_async1"
	tablePath		  = "/tables/t_test0"
	key               = "name"
	timeout           = 1 * time.Second
)
var valueField  = "item_name"
var valueFields = []string{"item_name"}
var condition = Condition{
	Op:		uint8(4),  //EQ
	Field:	"item_name",
	Value:	"item",
}
var limit = Limit{
	Offset:	uint32(0),
	Count:	uint32(4),
}
var valueFields_update = []string{"item_name_update"}
var conditions = []Condition{condition}

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

func TestCreateKVTable(t *testing.T) {
	result, err := service.CreateKVTable(KVTableName, key, valueField)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	if result != 0 {
		t.Fatalf("TestCreateKVTable failed, the result \"%v\" is inconsistent with \"0\"", result)
	}
	t.Logf("result: %d\n", result)
}

func TestAsyncCreateKVTable(t *testing.T) {
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
			t.Fatalf("TestAsyncCreateKVTable failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncCreateKVTable(handler, KVTableNameForAsync, key, valueField)
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

func TestAppendColumns(t *testing.T) {
	newColumns:=[]string{"test0","test1"} 

	result, err := service.AppendColumns(tablePath, newColumns)
	if err != nil {
		t.Fatalf("create AppendColumns failed: %v", err)
	}
	if result != 0 {
		t.Fatalf("TestAppendColumns failed, the result \"%v\" is inconsistent with \"0\"", result)
	}
	t.Logf("result: %d\n", result)
}

func TestAsyncAppendColumns(t *testing.T) {
	newColumns:=[]string{"test2","test3"} 

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
			t.Fatalf("TestAsyncAppendColumns failed, the result \"%v\" is inconsistent with \"0\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	_, err := service.AsyncAppendColumns(handler, tablePath, newColumns)
	if err != nil {
		t.Fatalf("AsyncAppendColumns failed: %v", err)
	}
	select {
	case <-channel:
		return
	case <-time.After(timeout):
		t.Fatal("timeout")
	}
}

func TestDescWithKeyOrder(t *testing.T) {
	tableInfo, err := service.DescWithKeyOrder(tableName)
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

func TestInsert(t *testing.T) {
	_, err := service.OpenTable(tableName)
	if err !=nil{
		service.CreateTable(tableName, key, valueFields)
	}
	// valueFields's length needs to be the same as the number of columns
	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}

	ret0, err := service.Insert(tableName,entry)
	t.Logf("ret0: %v\n", ret0)
	if err != nil {
		t.Fatalf("insert table failed: %v", err)
	}

	if ret0 != 1 {
		t.Fatalf("TestInsert failed, the ret0 \"%v\" is inconsistent with \"1\"", ret0)
	}
	t.Logf("ret0: %d\n", ret0)
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

	_, err := service.OpenTable(tableName)
	if err !=nil{
		service.CreateTable(tableName, key, valueFields)
	}
	// valueFields's length needs to be the same as the number of columns
	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}

	_, err = service.AsyncInsert(handler, tableName, entry)
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

func TestSelect0(t *testing.T) {
	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	service.Insert(tableName,entry)

	resultSelect, err := service.Select0(tableName, key)
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

func TestSelect(t *testing.T) {
	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	service.Insert(tableName,entry)

	resultSelects, err := service.Select(tableName, conditions, limit)
	if err != nil {
		t.Fatalf("select table failed: %v", err)
	}
	for i := 0; i < len(resultSelects); i++ {
		if resultSelects[i].Fields[0] != valueFields[0] {
			t.Fatalf("TestSelect failed, the result of resultSelect \"%v\" is not inconsistent", resultSelects[i].Fields[0])
		}
	}
	t.Logf("resultSelects %v:",resultSelects)
}

func TestUpdate(t *testing.T) {
	newValue := randStringBytes(10)
	updateField := UpdateField{
		ColumnName: valueFields[0],
		Value:      newValue,
	}
	var updateFields []UpdateField
	updateFields = append(updateFields, updateField)
	_, err := service.Update(tableName, key, updateFields)
	if err != nil {
		t.Fatalf("update table failed: %v", err)
	}
	// check update results
	resultSelect, err := service.Select0(tableName, key)
	if err != nil {
		t.Fatalf("select table failed: %v", err)
	}
	if resultSelect.Fields[0] != newValue {
		t.Fatalf("TestUpdate failed, the result of resultUpdate \"%v\" is not inconsistent", resultSelect.Fields[0])
	}
	t.Logf("updateResult: %s", resultSelect.Fields[0])
}

func TestUpdate0(t *testing.T) {
	newValue := randStringBytes(10)
	updateField := UpdateField{
		ColumnName: valueFields[0],
		Value:      newValue,
	}
	var updateFields []UpdateField
	updateFields = append(updateFields, updateField)

	condition := Condition{
		Op:		uint8(4),  //EQ
		Field:	valueField,
		Value:	"item",
	}
	conditions := []Condition{}
	conditions=append(conditions,condition)
	limit := Limit{
		Offset:	uint32(0),
		Count:	uint32(4),
	}

	newCondition := Condition{
		Op:		uint8(4),  //EQ
		Field:	valueField,
		Value:	newValue,
	}
	newConditions := []Condition{}
	newConditions=append(newConditions,newCondition)

	// key origin results
	originResultSelects, _ := service.Select(tableName, conditions, limit)
	t.Logf("originResultSelects %v:",originResultSelects)

	// perform update
	_, err := service.Update0(tableName, conditions, limit, updateFields)
	if err != nil {
		t.Fatalf("update table failed: %v", err)
	}
	
	// check update results
	afterResultSelects := []Entry{}
	for i := 0; i < len(originResultSelects); i++ {
		tempKey := originResultSelects[i].Key
		tempResultSelect, _ := service.Select0(tableName, tempKey)
		if tempResultSelect.Fields[0] != newValue {
			t.Fatalf("TestSelect failed, the result of resultSelect \"%v\" is not inconsistent", tempResultSelect.Fields[0])
		}
		afterResultSelects=append(afterResultSelects,tempResultSelect)
	}
	t.Logf("afterResultSelects %v:",afterResultSelects)
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

	newValue := randStringBytes(10)
	t.Logf("newValue: %v\n", newValue)
	updateField := UpdateField{
		ColumnName: valueFields[0],
		Value:      newValue,
	}
	var updateFields []UpdateField
	updateFields = append(updateFields, updateField)
	_, err := service.AsyncUpdate(handler, tableName, key, updateFields)
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

func TestAsyncUpdate0(t *testing.T) {
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
		if result < 1 {
			t.Fatalf("TestAsyncUpdate failed, the result \"%v\" is less then \"1\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	newValue := randStringBytes(10)
	t.Logf("newValue: %v\n", newValue)
	updateField := UpdateField{
		ColumnName: valueFields[0],
		Value:      newValue,
	}
	var updateFields []UpdateField
	updateFields = append(updateFields, updateField)
	_, err := service.AsyncUpdate0(handler, tableName, conditions, limit, updateFields)
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
	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	service.Insert(tableName,entry)

	ret0, err := service.Remove(tableName, key)
	if err != nil {
		t.Fatalf("remove table failed: %v", err)
	}
	if ret0 != 1 {
		t.Fatalf("TestRemove failed, the ret0 \"%v\" is inconsistent with \"1\"", ret0)
	}
	t.Logf("ret0: %d\n", ret0)
}

func TestRemove0(t *testing.T) {
	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	service.Insert(tableName,entry)

	ret0, err := service.Remove0(tableName, conditions, limit)
	if err != nil {
		t.Fatalf("remove table failed: %v", err)
	}
	if ret0 < 1 {
		t.Fatalf("TestRemove failed, the ret0 \"%v\" is less then \"1\"", ret0)
	}
	t.Logf("ret0: %d\n", ret0)
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

	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	service.Insert(tableName,entry)

	_, err = service.AsyncRemove(handler, tableName, key)
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

func TestAsyncRemove0(t *testing.T) {
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
		if result < 1 {
			t.Fatalf("TestAsyncRemove failed, the result \"%v\" is less then \"1\"", result)
		}
		t.Logf("result: %d\n", result)
		channel <- 0
	}

	valueFields := []string{}
	tableInfo, err := service.DescWithKeyOrder(tableName)
	for i := 0; i < len(tableInfo.ValueColumns); i++ {
        valueFields=append(valueFields,"item")
    }
	key := randStringBytes(10)
	t.Logf("key: %v\n", key)
	entry := Entry{
		Key:    key,
		Fields: valueFields,
	}
	service.Insert(tableName,entry)

	_, err = service.AsyncRemove0(handler, tableName, conditions, limit)
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

func TestSet(t *testing.T) {
	_, err := service.OpenKVTable(KVTableName)
	if err !=nil{
		service.CreateKVTable(KVTableName, key, valueField)
	}

	key := randStringBytes(10)
	value := randStringBytes(10)

	ret0, err := service.Set(KVTableName,key,value)
	if err != nil {
		t.Fatalf("KVTable set failed: %v", err)
	}
	if ret0 != 1 {
		t.Fatalf("TestSet failed, the ret0 \"%v\" is inconsistent with \"1\"", ret0)
	}

	_, tempValue, _ := service.Get(KVTableName,key)
	if value != tempValue {
		t.Fatalf("TestSet failed, the value \"%v\" is inconsistent with the tempValue \"%v\"", value, tempValue)
	}

	t.Logf("key: %v\n", key)
	t.Logf("value: %v\n", value)
}

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
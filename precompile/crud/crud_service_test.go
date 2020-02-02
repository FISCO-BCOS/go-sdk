package crud

import (
	"crypto/ecdsa"
	"math/rand"
	"strconv"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetClient(t *testing.T) *client.Client {
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58", NodeURL: "http://localhost:8545"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("can not dial to the RPC API: %v", err)
	}
	return c
}

func GenerateKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	if err != nil {
		t.Fatalf("init privateKey failed: %+v", err)
	}
	return privateKey
}

func GetService(t *testing.T) *CRUDService {
	rpc := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewCRUDService(rpc, privateKey)
	if err != nil {
		t.Fatalf("init CRUDService failed: %+v", err)
	}
	return service
}

func TestCRUD(t *testing.T) {
	tableName := "t_test" + strconv.Itoa(rand.Intn(100000))
	key := "name"
	valueFields := "item_id, item_name"
	table := &Table{TableName: tableName, Key: key, ValueFields: valueFields}

	service := GetService(t)

	// create table
	resultCreate, err := service.CreateTable(table)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	t.Logf("resultCreate: %d\n", resultCreate)

	// insert records
	var insertResults int
	for i := 1; i <= 5; i++ {
		insertEnrty := table.GetEntry()
		insertEnrty.Put("item_id", "1")
		insertEnrty.Put("item_name", "apple"+strconv.Itoa(i))
		table.SetKey("fruit")
		insertResult, err := service.Insert(table, insertEnrty)
		if err != nil {
			t.Fatalf("insert table faied: %v", err)
		}
		insertResults += insertResult
	}
	t.Logf("insertResults: %d\n", insertResults)

	// select records
	condition1 := table.GetCondition()
	condition1.EQ("item_id", "1")
	condition1.Limit(1)

	resultSelect1, err := service.Select(table, condition1)
	if err != nil {
		t.Fatalf("select table faied: %v", err)
	}
	t.Logf("resultSelect1 :\n")
	t.Logf("%s\n", resultSelect1[0]["name"])
	t.Logf("%s\n", resultSelect1[0]["item_id"])
	t.Logf("%s\n", resultSelect1[0]["item_name"])

	// update records
	updateEntry := table.GetEntry()
	updateEntry.Put("item_id", "1")
	updateEntry.Put("item_name", "orange")
	updateCondition := table.GetCondition()
	updateCondition.EQ("item_id", "1")
	updateResult, err := service.Update(table, updateEntry, updateCondition)
	if err != nil {
		t.Fatalf("update table failed: %v", err)
	}
	t.Logf("updateResult: %d", updateResult)

	// select records
	condition2 := table.GetCondition()
	condition2.EQ("item_id", "1")
	condition2.Limit(1)

	resultSelect2, err := service.Select(table, condition2)
	if err != nil {
		t.Fatalf("select table faied: %v", err)
	}
	t.Logf("resultSelect2 :\n")
	t.Logf("%s\n", resultSelect2[0]["name"])
	t.Logf("%s\n", resultSelect2[0]["item_id"])
	t.Logf("%s\n", resultSelect2[0]["item_name"])

	// remove records
	removeCondition := table.GetCondition()
	removeCondition.EQ("item_id", "1")
	removeResult, err := service.Remove(table, removeCondition)
	if err != nil {
		t.Fatalf("remove table faied: %v", err)
	}
	t.Logf("removeResult: %d\n", removeResult)
}

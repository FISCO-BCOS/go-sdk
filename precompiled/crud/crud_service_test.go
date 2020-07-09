package crud

import (
	"crypto/ecdsa"
	"strconv"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	tableName   = "t_test"
	key         = "name"
	valueFields = "item_id, item_name"
)

func GetClient(t *testing.T) *client.Client {
	// config := &conf.ParseConfig("config.toml")[0]
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58",
		NodeURL:    "http://localhost:8545"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("Dial to %s failed of %v", config.NodeURL, err)
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
	c := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewCRUDService(c, privateKey)
	if err != nil {
		t.Fatalf("init CRUDService failed: %+v", err)
	}
	return service
}

func TestCreateTable(t *testing.T) {
	service := GetService(t)

	resultCreate, err := service.CreateTable(tableName, key, valueFields)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	t.Logf("resultCreate: %d\n", resultCreate)
}

func TestInsert(t *testing.T) {
	service := GetService(t)

	var insertResults int
	insertEntry := NewEntry()
	for i := 1; i <= 5; i++ {
		insertEntry.Put("item_id", "1")
		insertEntry.Put("item_name", "apple"+strconv.Itoa(i))
		insertResult, err := service.Insert(tableName, "fruit", insertEntry)
		if err != nil {
			t.Fatalf("insert table failed: %v", err)
		}
		insertResults += insertResult
	}
	t.Logf("insertResults: %d\n", insertResults)
}

func TestSelect(t *testing.T) {
	service := GetService(t)

	condition := NewCondition()
	condition.EQ("item_id", "1")
	condition.Limit(5)

	resultSelect, err := service.Select(tableName, "fruit", condition)
	if err != nil {
		t.Fatalf("select table failed: %v", err)
	}
	t.Logf("resultSelect :\n")
	t.Logf("%d", len(resultSelect))
	for i := 0; i < len(resultSelect); i++ {
		t.Logf("resultSelect[%d]'s name is：%s\n", i, resultSelect[i]["name"])
		t.Logf("resultSelect[%d]'s item_id is：%s\n", i, resultSelect[i]["item_id"])
		t.Logf("resultSelect[%d]'s item_name is：%s\n", i, resultSelect[i]["item_name"])
	}
}

func TestUpdate(t *testing.T) {
	service := GetService(t)

	updateEntry := NewEntry()
	updateEntry.Put("item_id", "1")
	updateEntry.Put("item_name", "orange")
	updateCondition := NewCondition()
	updateCondition.EQ("item_id", "1")
	updateResult, err := service.Update(tableName, "fruit", updateEntry, updateCondition)
	if err != nil {
		t.Fatalf("update table failed: %v", err)
	}
	t.Logf("updateResult: %d", updateResult)
}

func TestRemove(t *testing.T) {
	service := GetService(t)

	removeCondition := NewCondition()
	removeCondition.EQ("item_id", "1")
	removeResult, err := service.Remove(tableName, "fruit", removeCondition)
	if err != nil {
		t.Fatalf("remove table failed: %v", err)
	}
	t.Logf("removeResult: %d\n", removeResult)
}

func TestDesc(t *testing.T) {
	service := GetService(t)

	keyField, valueField, err := service.Desc(tableName)
	if err != nil {
		t.Fatalf("query table info by tableName failed: %v", err)
	}
	t.Logf("keyFiled is：%s\n", keyField)
	t.Logf("valueField is：%s\n", valueField)
}

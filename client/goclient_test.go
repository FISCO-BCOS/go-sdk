package client

import (
	"context"
	"fmt"
	helloworld "github.com/FISCO-BCOS/go-sdk/.ci/hello"
	"strings"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/conf"
)

var (
	contractAddress = ""
	blockHash = ""       // get blockHash by TestBlockHashByNumber test case
	transactionHash = ""
)

func GetClient(t *testing.T) *Client {
	config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58", NodeURL: "http://localhost:8545"}
	c, err := Dial(config)
	if err != nil {
		t.Fatalf("Dial to %s failed of %v", config.NodeURL, err)
	}
	return c
}

func deployHelloWorldContract(t *testing.T) error {
	client := GetClient(t)
	address, tx, instance, err := helloworld.DeployHelloWorld(client.GetTransactOpts(), client) // deploy contract
	if err != nil {
		return fmt.Errorf("deploy HelloWorld contract failed：%v", err)
	}
	contractAddress = address.Hex()
	transactionHash = tx.Hash().Hex()
	_ = instance
	return nil
}

// Get contractAddress、transactionHash、blockHash by this test
func TestBlockHashByNumber(t *testing.T) {
	err := deployHelloWorldContract(t)
	if err != nil {
		t.Logf("excute TestBlockHashByNumber test case failed：%v", err)
	}

	c := GetClient(t)

	bnum := "0x1"
	raw, err := c.GetBlockHashByNumber(context.Background(), bnum)
	if err != nil {
		t.Fatalf("block hash not found: %v", err)
	}

	t.Logf("block hash by number:\n%s", raw)
	blockHash = strings.Trim(string(raw),"\"")
}

func TestClientVersion(t *testing.T) {
	c := GetClient(t)

	cv, err := c.GetClientVersion(context.Background())
	if err != nil {
		t.Fatalf("client version not found: %v", err)
	}

	t.Logf("client version:\n%s", cv)
}

func TestBlockNumber(t *testing.T) {
	c := GetClient(t)
	// cannot use big.NewInt to construct json request
	// TODO: analysis the ethereum's big.NewInt
	bn, err := c.GetBlockNumber(context.Background())
	if err != nil {
		t.Fatalf("block number not found: %v", err)
	}

	t.Logf("latest block number: \n%s", bn)
}

func TestPBFTView(t *testing.T) {
	c := GetClient(t)
	pv, err := c.GetPBFTView(context.Background())
	if err != nil {
		t.Fatalf("PBFT view not found: %v", err)
	}

	t.Logf("PBFT view: \n%s", pv)
}

func TestBlockLimit(t *testing.T) {
    c := GetClient(t)
	// cannot use big.NewInt to construct json request
	// TODO: analysis the ethereum's big.NewInt
	bl, err := c.GetBlockLimit(context.Background())
	if err != nil {
		t.Fatalf("blockLimit not found: %v", err)
	}

	t.Logf("latest blockLimit: \n%s", bl)
}

func TestGroupID(t *testing.T) {
    c := GetClient(t)
	// cannot use big.NewInt to construct json request
	// TODO: analysis the ethereum's big.NewInt
	groupid := c.GetGroupID()
	t.Logf("current groupID: \n%s", groupid)
}

func TestChainID(t *testing.T) {
    c := GetClient(t)
	// cannot use big.NewInt to construct json request
	// TODO: analysis the ethereum's big.NewInt
	chainid, err := c.GetChainID(context.Background())
	if err != nil {
		t.Fatalf("Chain ID not found: %v", err)
	}
	t.Logf("Chain ID: \n%s", chainid)
}

func TestSealerList(t *testing.T) {
	c := GetClient(t)
	sl, err := c.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("sealer list not found: %v", err)
	}

	t.Logf("sealer list:\n%s", sl)
}

func TestObserverList(t *testing.T) {
	c := GetClient(t)
	ol, err := c.GetObserverList(context.Background())
	if err != nil {
		t.Fatalf("observer list not found: %v", err)
	}

	t.Logf("observer list:\n%s", ol)
}

func TestConsensusStatus(t *testing.T) {
	c := GetClient(t)
	status, err := c.GetConsensusStatus(context.Background())
	if err != nil {
		t.Fatalf("consensus status not found: %v", err)
	}

	t.Logf("consensus status:\n%s", status)
}

func TestSyncStatus(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetSyncStatus(context.Background())
	if err != nil {
		t.Fatalf("synchronization status not found: %v", err)
	}

	t.Logf("synchronization Status:\n%s", raw)
}

func TestPeers(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetPeers(context.Background())
	if err != nil {
		t.Fatalf("peers not found: %v", err)
	}

	t.Logf("peers:\n%s", raw)
}

func TestGroupPeers(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetGroupPeers(context.Background())
	if err != nil {
		t.Fatalf("group peers not found: %v", err)
	}

	t.Logf("group peers:\n%s", raw)
}

func TestNodeIDList(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetNodeIDList(context.Background())
	if err != nil {
		t.Fatalf("nodeID list not found: %v", err)
	}

	t.Logf("nodeID list:\n %s", raw)
}

func TestGroupList(t *testing.T) {
	c := GetClient(t)
	raw, err := c.GetGroupList(context.Background())
	if err != nil {
		t.Fatalf("group list not found: %v", err)
	}

	t.Logf("group list:\n%s", raw)
}

func TestBlockByHash(t *testing.T) {
	c := GetClient(t)

	includeTx := false
	raw, err := c.GetBlockByHash(context.Background(), blockHash, includeTx)
	if err != nil {
		t.Fatalf("block not found: %v", err)
	}

	t.Logf("block by hash:\n%s", raw)
}

func TestBlockByNumber(t *testing.T) {
	c := GetClient(t)

	bnum := "0x1"
	includeTx := true
	raw, err := c.GetBlockByNumber(context.Background(), bnum, includeTx)
	if err != nil {
		t.Fatalf("block not found: %v", err)
	}

	t.Logf("block by number:\n%s", raw)
}

func TestTransactionByHash(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetTransactionByHash(context.Background(), transactionHash)
	if err != nil {
		t.Fatalf("transaction not found: %v", err)
	}

	t.Logf("transaction by hash:\n%s", raw)
}

func TestTransactionByBlockHashAndIndex(t *testing.T) {
	c := GetClient(t)

	txindex := "0x0"
	raw, err := c.GetTransactionByBlockHashAndIndex(context.Background(), blockHash, txindex)
	if err != nil {
		t.Fatalf("transaction not found: %v", err)
	}

	t.Logf("transaction by block hash and transaction index:\n%s", raw)
}

func TestTransactionByBlockNumberAndIndex(t *testing.T) {
	c := GetClient(t)

	bnum := "0x1"
	txindex := "0x0"
	raw, err := c.GetTransactionByBlockNumberAndIndex(context.Background(), bnum, txindex)
	if err != nil {
		t.Fatalf("transaction not found: %v", err)
	}

	t.Logf("transaction by block number and transaction index:\n%s", raw)
}

func TestTransactionReceipt(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetTransactionReceipt(context.Background(), transactionHash)
	if err != nil {
		t.Fatalf("transaction receipt not found: %v", err)
	}
	t.Logf("transaction receipt by transaction hash:\n%s", raw)
}

func TestContractAddress(t *testing.T) {
	c := GetClient(t)

	ca, err := c.GetContractAddress(context.Background(), transactionHash)
	if err != nil {
		t.Fatalf("ContractAddress not found: %v", err)
	}

	t.Logf("ContractAddress: \n%s", ca.String())
}

func TestPendingTransactions(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetPendingTransactions(context.Background())
	if err != nil {
		t.Fatalf("pending transactions not found: %v", err)
	}

	t.Logf("pending transactions:\n%s", raw)
}

func TestPendingTxSize(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetPendingTxSize(context.Background())
	if err != nil {
		t.Fatalf("pending transactions not found: %v", err)
	}

	t.Logf("the amount of the pending transactions:\n%s", raw)
}

func TestGetCode(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetCode(context.Background(), contractAddress)
	if err != nil {
		t.Fatalf("contract not found: %v", err)
	}

	t.Logf("the contract code:\n%s", raw)
}

func TestTotalTransactionCount(t *testing.T) {
	c := GetClient(t)

	raw, err := c.GetTotalTransactionCount(context.Background())
	if err != nil {
		t.Fatalf("transactions not found: %v", err)
	}

	t.Logf("the totoal transactions and present block height:\n%s", raw)
}

func TestSystemConfigByKey(t *testing.T) {
	c := GetClient(t)

	findkey := "tx_count_limit"
	raw, err := c.GetSystemConfigByKey(context.Background(), findkey)
	if err != nil {
		t.Fatalf("the value not found: %v", err)
	}

	t.Logf("the value got by the key:\n%s", raw)
}

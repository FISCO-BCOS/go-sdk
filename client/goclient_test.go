package client

import (
	"testing"
	"context"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func GetClient(t *testing.T) (*Client) {
	config := &conf.Config{IsHTTP:true,ChainID:1,IsSMCrypto:false,ConnInfo:conf.Connection{GroupID:1,NodeURL:"http://localhost:8545"}}
	// RPC API
	c, err := Dial(config)  // change to your RPC and groupID
	if err != nil {
		t.Fatalf("can not dial to the RPC API: %v", err)
	}
	return c
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

// func TestBlockLimit(t *testing.T) {
//     c := GetClient(t)
// 	// cannot use big.NewInt to construct json request
// 	// TODO: analysis the ethereum's big.NewInt
// 	bl, err := c.GetBlockLimit(context.Background())
// 	if err != nil {
// 		t.Fatalf("blockLimit not found: %v", err)
// 	}

// 	t.Logf("latest blockLimit: \n%s", bl)
// }

// func TestGroupID(t *testing.T) {
//     c := GetClient(t)
// 	// cannot use big.NewInt to construct json request
// 	// TODO: analysis the ethereum's big.NewInt
// 	groupid := c.GetGroupID()
// 	t.Logf("current groupID: \n%s", groupid)
// }

// func TestChainID(t *testing.T) {
//     c := GetClient(t)
// 	// cannot use big.NewInt to construct json request
// 	// TODO: analysis the ethereum's big.NewInt
// 	chainid, err := c.GetChainID(context.Background())
// 	if err != nil {
// 		t.Fatalf("Chain ID not found: %v", err)
// 	}
// 	t.Logf("Chain ID: \n%s", chainid)
// }

// func TestSealerList(t *testing.T) {
// 	c := GetClient(t)
// 	sl, err := c.GetSealerList(context.Background())
// 	if err != nil {
// 		t.Fatalf("sealer list not found: %v", err)
// 	}

// 	t.Logf("sealer list:\n%s", sl)
// }

// func TestObserverList(t *testing.T) {
// 	c := GetClient(t)
// 	ol, err := c.GetObserverList(context.Background())
// 	if err != nil {
// 		t.Fatalf("observer list not found: %v", err)
// 	}

// 	t.Logf("observer list:\n%s", ol)
// }

// func TestConsensusStatus(t *testing.T) {
// 	c := GetClient(t)
// 	status, err := c.GetConsensusStatus(context.Background())
// 	if err != nil {
// 		t.Fatalf("consensus status not found: %v", err)
// 	}

// 	t.Logf("consensus status:\n%s", status)
// }

// func TestSyncStatus(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetSyncStatus(context.Background())
// 	if err != nil {
// 		t.Fatalf("synchronization status not found: %v", err)
// 	}

// 	t.Logf("synchronization Status:\n%s", raw)
// }

// func TestPeers(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetPeers(context.Background())
// 	if err != nil {
// 		t.Fatalf("peers not found: %v", err)
// 	}

// 	t.Logf("peers:\n%s", raw)
// }

// func TestGroupPeers(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetGroupPeers(context.Background())
// 	if err != nil {
// 		t.Fatalf("group peers not found: %v", err)
// 	}

// 	t.Logf("group peers:\n%s", raw)
// }

// func TestNodeIDList(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetNodeIDList(context.Background())
// 	if err != nil {
// 		t.Fatalf("nodeID list not found: %v", err)
// 	}
    
// 	t.Logf("nodeID list:\n %s", raw)
// }

// func TestGroupList(t *testing.T) {
// 	c := GetClient(t)
// 	raw, err := c.GetGroupList(context.Background())
// 	if err != nil {
// 		t.Fatalf("group list not found: %v", err)
// 	}

// 	t.Logf("group list:\n%s", raw)
// }

// func TestBlockByHash(t *testing.T) {
// 	c := GetClient(t)
	
// 	bhash := "0xc0b21d064b97bafda716e07785fe8bb20cc23506bb980f12c7f7a4f4ef50ce30"
// 	includeTx := false
// 	raw, err := c.GetBlockByHash(context.Background(), bhash, includeTx)
// 	if err != nil {
// 		t.Fatalf("block not found: %v", err)
// 	}

// 	t.Logf("block by hash:\n%s", raw)
// }

// func TestBlockByNumber(t *testing.T) {
// 	c := GetClient(t)
	
// 	bnum := "0x1"
// 	includeTx := true
// 	raw, err := c.GetBlockByNumber(context.Background(), bnum, includeTx)
// 	if err != nil {
// 		t.Fatalf("block not found: %v", err)
// 	}

// 	t.Logf("block by number:\n%s", raw)
// }

// func TestBlockHashByNumber(t *testing.T) {
// 	c := GetClient(t)
	
// 	bnum := "0x1"
// 	raw, err := c.GetBlockHashByNumber(context.Background(), bnum)
// 	if err != nil {
// 		t.Fatalf("block hash not found: %v", err)
// 	}

// 	t.Logf("block hash by number:\n%s", raw)
// }

// func TestTransactionByHash(t *testing.T) {
// 	c := GetClient(t)
	
// 	txhash := "0xed51827558939e8d103cbf8f6ff37f8a99582f09afa29e5636d0e54a073d0893"
// 	raw, err := c.GetTransactionByHash(context.Background(), txhash)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by hash:\n%s", raw)
// }

// func TestTransactionByBlockHashAndIndex(t *testing.T) {
// 	c := GetClient(t)
	
// 	bhash := "0xc0b21d064b97bafda716e07785fe8bb20cc23506bb980f12c7f7a4f4ef50ce30"
// 	txindex := "0x0"
// 	raw, err := c.GetTransactionByBlockHashAndIndex(context.Background(), bhash, txindex)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by block hash and transaction index:\n%s", raw)
// }

// func TestTransactionByBlockNumberAndIndex(t *testing.T) {
// 	c := GetClient(t)
	
// 	bnum := "0x1"
// 	txindex := "0x0"
// 	raw, err := c.GetTransactionByBlockNumberAndIndex(context.Background(), bnum, txindex)
// 	if err != nil {
// 		t.Fatalf("transaction not found: %v", err)
// 	}

// 	t.Logf("transaction by block number and transaction index:\n%s", raw)
// }

// func TestTransactionReceipt(t *testing.T) {
// 	c := GetClient(t)
	
// 	txhash := "0x6613da911621248ffc2983a4553d44fb299887ae2d803eb81d9890465fbf29e5"
// 	raw, err := c.GetTransactionReceipt(context.Background(), txhash)
// 	if err != nil {
// 		t.Fatalf("transaction receipt not found: %v", err)
// 	}
// 	t.Logf("transaction receipt by transaction hash:\n%s", raw)
// }

// func TestContractAddress(t *testing.T) {
// 	c := GetClient(t)
// 	txhash := "0x4a2a4d878318a83491383d29d6550c088bdcf692e3055b060342dcd85177c621"
// 	ca, err := c.GetContractAddress(context.Background(), txhash)
// 	if err != nil {
// 		t.Fatalf("ContractAddress not found: %v", err)
// 	}

// 	t.Logf("ContractAddress: \n%s", ca.String())
// }

// func TestPendingTransactions(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetPendingTransactions(context.Background())
// 	if err != nil {
// 		t.Fatalf("pending transactions not found: %v", err)
// 	}

// 	t.Logf("pending transactions:\n%s", raw)
// }

// func TestPendingTxSize(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetPendingTxSize(context.Background())
// 	if err != nil {
// 		t.Fatalf("pending transactions not found: %v", err)
// 	}

// 	t.Logf("the amount of the pending transactions:\n%s", raw)
// }

// func TestGetCode(t *testing.T) {
// 	c := GetClient(t)
	
// 	addr := "0x27c1b5d9fe3ab035c2e9db7199d4beb139e12292"
// 	raw, err := c.GetCode(context.Background(), addr)
// 	if err != nil {
// 		t.Fatalf("contract not found: %v", err)
// 	}

// 	t.Logf("the contract code:\n%s", raw)
// }

// func TestTotalTransactionCount(t *testing.T) {
// 	c := GetClient(t)
	
// 	raw, err := c.GetTotalTransactionCount(context.Background())
// 	if err != nil {
// 		t.Fatalf("transactions not found: %v", err)
// 	}

// 	t.Logf("the totoal transactions and present block height:\n%s", raw)
// }

// func TestSystemConfigByKey(t *testing.T) {
// 	c := GetClient(t)
	
// 	findkey := "tx_count_limit"
// 	raw, err := c.GetSystemConfigByKey(context.Background(), findkey)
// 	if err != nil {
// 		t.Fatalf("the value not found: %v", err)
// 	}

// 	t.Logf("the value got by the key:\n%s", raw)
// }

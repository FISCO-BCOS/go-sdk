package consensus

import (
	"context"
	"crypto/ecdsa"
	"regexp"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	nodeID = ""
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

func GetService(t *testing.T) *ConsensusService {
	c := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewConsensusService(c, privateKey)
	if err != nil {
		t.Fatalf("init ConsensusService failed: %+v", err)
	}
	return service
}

// Get nodeID
// TODO: try to use TestMain function to init before excute test case
func TestGetNodeID(t *testing.T) {
	c := GetClient(t)
	sealerList, err := c.GetNodeIDList(context.Background())
	if err != nil {
		t.Fatalf("sealer list not found: %v", err)
	}
	reg := regexp.MustCompile(`[\w]+`)
	nodeList := reg.FindAllString(string(sealerList), -1)
	if len(nodeList) < 4 { // pbft consensus needs 2f+1
		t.Fatalf("the number of nodes does not exceed 4")
	}
	nodeID = nodeList[1]
}

func TestAddObserver(t *testing.T) {
	c := GetClient(t)
	service := GetService(t)

	observer, err := c.GetObserverList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService GetObserverList failed: %+v\n", err)
	}
	t.Logf("Observer list: %s\n", observer)

	tx, err := service.AddObserver(nodeID)
	if err != nil {
		t.Fatalf("ConsensusService AddObserver failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), c, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s", receipt.GetTransactionHash())

	observer, err = c.GetObserverList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService invoke GetObserverList second time failed: %+v\n", err)
	}
	t.Logf("Observer list: %s\n", observer)
}

func TestAddSealer(t *testing.T) {
	c := GetClient(t)
	service := GetService(t)

	observer, err := c.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService GetSealerList failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)

	tx, err := service.AddSealer(nodeID)
	if err != nil {
		t.Fatalf("ConsensusService AddSealer failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), c, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s", receipt.GetTransactionHash())

	observer, err = c.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService invoke GetSealerList second time failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)
}

func TestRemove(t *testing.T) {
	c := GetClient(t)
	service := GetService(t)

	observer, err := c.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService GetSealerList failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)

	tx, err := service.RemoveNode(nodeID)

	if err != nil {
		t.Fatalf("ConsensusService Remove failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), c, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s", receipt.GetTransactionHash())

	observer, err = c.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService invoke GetSealerList second time failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)
}

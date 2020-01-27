package consensus

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/crypto"
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

func GetService(t *testing.T) *ConsensusService {
	rpc := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewConsensusService(rpc, privateKey)
	if err != nil {
		t.Fatalf("init ConsensusService failed: %+v", err)
	}
	return service
}

func TestAddObserver(t *testing.T) {
	rpc := GetClient(t)
	service := GetService(t)

	observer, err := rpc.GetObserverList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService GetObserverList failed: %+v\n", err)
	}
	t.Logf("Observer list: %s\n", observer)

	nodeID := "da72d42af7228b7fcbd0c2ca1128a9cf5b1a3a648c64878ebba4177a751507a0e1d686c2a6ccdfdadcfc60c1d6ec6d5d07797880f2f6a1f176d480b98ed5a13c"
	tx, err := service.AddObserver(nodeID)
	if err != nil {
		t.Fatalf("ConsensusService AddObserver failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), rpc, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s", receipt.GetTransactionHash())

	observer, err = rpc.GetObserverList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService invoke GetObserverList second time failed: %+v\n", err)
	}
	t.Logf("Observer list: %s\n", observer)
}

func TestAddSealer(t *testing.T) {
	rpc := GetClient(t)
	service := GetService(t)

	observer, err := rpc.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService GetSealerList failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)

	nodeID := "da72d42af7228b7fcbd0c2ca1128a9cf5b1a3a648c64878ebba4177a751507a0e1d686c2a6ccdfdadcfc60c1d6ec6d5d07797880f2f6a1f176d480b98ed5a13c"
	tx, err := service.AddSealer(nodeID)
	if err != nil {
		t.Fatalf("ConsensusService AddSealer failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), rpc, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s", receipt.GetTransactionHash())

	observer, err = rpc.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService invoke GetSealerList second time failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)
}

func TestRemove(t *testing.T) {
	rpc := GetClient(t)
	service := GetService(t)

	observer, err := rpc.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService GetSealerList failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)

	nodeID := "da72d42af7228b7fcbd0c2ca1128a9cf5b1a3a648c64878ebba4177a751507a0e1d686c2a6ccdfdadcfc60c1d6ec6d5d07797880f2f6a1f176d480b98ed5a13c"
	tx, err := service.RemoveNode(nodeID)
	if err != nil {
		t.Fatalf("ConsensusService Remove failed: %+v\n", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), rpc, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}
	t.Logf("transaction hash: %s", receipt.GetTransactionHash())

	observer, err = rpc.GetSealerList(context.Background())
	if err != nil {
		t.Fatalf("ConsensusService invoke GetSealerList second time failed: %+v\n", err)
	}
	t.Logf("Sealer list: %s\n", observer)
}

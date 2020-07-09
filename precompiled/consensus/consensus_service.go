package consensus

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

// ConsensusService is a precompile contract service.
type ConsensusService struct {
	consensus     *Consensus
	consensusAuth *bind.TransactOpts
	client        *client.Client
}

// contract address
var consensusPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001003")

// NewConsensusService returns ptr of ConsensusService
func NewConsensusService(client *client.Client, privateKey *ecdsa.PrivateKey) (*ConsensusService, error) {
	instance, err := NewConsensus(consensusPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct ConsensusService failed: %+v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = big.NewInt(30000000)
	return &ConsensusService{consensus: instance, consensusAuth: auth, client: client}, nil
}

// AddObserver add a new observe node according to the node ID
func (service *ConsensusService) AddObserver(nodeID string) (*types.Transaction, error) {
	flag, err := service.isValidNodeID(nodeID)
	if err != nil {
		return nil, err
	} else if !flag {
		return nil, fmt.Errorf("the node is not reachable")
	}

	observerRaw, err := service.client.GetObserverList(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get the observer list failed: %v", err)
	}

	var nodeIDs []string
	err = json.Unmarshal(observerRaw, &nodeIDs)
	if err != nil {
		return nil, fmt.Errorf("unmarshal the observer list failed: %v", err)
	}

	for _, nID := range nodeIDs {
		if nID == nodeID {
			return nil, fmt.Errorf("the node is already in the observer lisn")
		}
	}
	tx, err := service.consensus.AddObserver(service.consensusAuth, nodeID)
	if err != nil {
		return nil, fmt.Errorf("ConsensusService addObserver failed: %+v", err)
	}
	return tx, nil
}

// AddSealer add a new sealer node according to the node ID
func (service *ConsensusService) AddSealer(nodeID string) (*types.Transaction, error) {
	flag, err := service.isValidNodeID(nodeID)
	if err != nil {
		return nil, err
	} else if !flag {
		return nil, fmt.Errorf("the node is not reachable")
	}

	sealerRaw, err := service.client.GetSealerList(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get the sealer list failed: %v", err)
	}

	var nodeIDs []string
	err = json.Unmarshal(sealerRaw, &nodeIDs)
	if err != nil {
		return nil, fmt.Errorf("unmarshal the sealer list failed: %v", err)
	}

	for _, nID := range nodeIDs {
		if nID == nodeID {
			return nil, fmt.Errorf("the node is already in the sealer list")
		}
	}

	tx, err := service.consensus.AddSealer(service.consensusAuth, nodeID)
	if err != nil {
		return nil, fmt.Errorf("ConsensusService addSealer failed: %+v", err)
	}

	return tx, nil
}

// RemoveNode remove a sealer node according to the node ID
func (service *ConsensusService) RemoveNode(nodeID string) (*types.Transaction, error) {
	peersRaw, err := service.client.GetGroupPeers(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get the group peers failed: %v", err)
	}

	var nodeIDs []string
	err = json.Unmarshal(peersRaw, &nodeIDs)
	if err != nil {
		return nil, fmt.Errorf("unmarshal the group peers failed: %v", err)
	}

	var flag = true
	for _, nID := range nodeIDs {
		if nID == nodeID {
			flag = false
			break
		}
	}
	if flag {
		return nil, fmt.Errorf("the node is not a group peer")
	}

	tx, err := service.consensus.Remove(service.consensusAuth, nodeID)
	// maybe will occur something wrong
	// when request the receipt from the SDK since the connected node of SDK is removed
	if err != nil {
		return nil, fmt.Errorf("ConsensusService Remove failed: %+v", err)
	}
	return tx, nil
}

// isValidNodeID returns true if the nodeID exits in NodeIDlist.
func (service *ConsensusService) isValidNodeID(nodeID string) (bool, error) {
	var flag = false
	nodeIDRaw, err := service.client.GetNodeIDList(context.Background())
	if err != nil {
		return flag, fmt.Errorf("get the valid Node IDs failed: %v", err)
	}
	var nodeIDs []string
	err = json.Unmarshal(nodeIDRaw, &nodeIDs)
	if err != nil {
		return flag, fmt.Errorf("unmarshal the valid Node IDs failed: %v", err)
	}
	for _, nID := range nodeIDs {
		if nID == nodeID {
			flag = true
		}
	}
	return flag, nil
}

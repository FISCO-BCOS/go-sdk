package sharding

import (
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompiled"
	"github.com/ethereum/go-ethereum/common"
)

type Service struct {
	sharding     *Sharding
	shardingAuth *bind.TransactOpts
	client  *client.Client
}

// contract address
var shardingPrecompileAddress = common.HexToAddress("0000000000000000000000000000000000001010")

// NewCnsService returns ptr of Service
func NewShardingService(client *client.Client) (*Service, error) {
	instance, err := NewSharding(shardingPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct Service failed: %+v", err)
	}
	auth := client.GetTransactOpts()
	return &Service{sharding: instance, shardingAuth: auth, client: client}, nil
}

func (service *Service) GetContractShard(absolutePath string) (int64, string, error) {
	opts := &bind.CallOpts{From: service.shardingAuth.From}
	ret0, ret1, err := service.sharding.GetContractShard(opts, absolutePath)
	if err != nil {
		return precompiled.DefaultErrorCode, ret1, fmt.Errorf("service GetContractShard failed: %+v", err)
	}
	return ret0.Int64(), ret1, err
}

func (service *Service) MakeShard(shardName string) (int64, *types.Transaction, *types.Receipt, error) {
	ret0, ret1, ret2, err := service.sharding.MakeShard(service.shardingAuth, shardName)
	if err != nil {
		return precompiled.DefaultErrorCode, nil, nil, fmt.Errorf("service MakeShard failed: %+v", err)
	}
	return ret0.Int64(), ret1, ret2, err
}

func (service *Service) AsyncMakeShard(handler func(*types.Receipt, error), shardName string) (*types.Transaction, error) {
	return service.sharding.AsyncMakeShard(handler, service.shardingAuth, shardName)
}

func (service *Service) LinkShard(shardName string, _address string) (int64, *types.Transaction, *types.Receipt, error) {
	ret0, ret1, ret2, err := service.sharding.LinkShard(service.shardingAuth, shardName, _address)
	if err != nil {
		return precompiled.DefaultErrorCode, nil, nil, fmt.Errorf("service LinkShard failed: %+v", err)
	}
	return ret0.Int64(), ret1, ret2, err
}

func (service *Service) AsyncLinkShard(handler func(*types.Receipt, error), shardName string, _address string) (*types.Transaction, error) {
	return service.sharding.AsyncLinkShard(handler, service.shardingAuth, shardName, _address)
}
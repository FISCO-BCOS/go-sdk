package bfs

import (
	"fmt"

	"math/big"

	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/precompiled"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

type Service struct {
	bfs     *Bfs
	bfsAuth *bind.TransactOpts
	client  *client.Client
}

// contract address
var bfsPrecompileAddress = common.HexToAddress("000000000000000000000000000000000000100e")

// NewCnsService returns ptr of Service
func NewBfsService(client *client.Client) (*Service, error) {
	instance, err := NewBfs(bfsPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct Service failed: %+v", err)
	}
	auth := client.GetTransactOpts()
	return &Service{bfs: instance, bfsAuth: auth, client: client}, nil
}

func (service *Service) List0(absolutePath string) (int64, []BfsInfo, error) {
	opts := &bind.CallOpts{From: service.bfsAuth.From}
	ret0, ret1, err := service.bfs.List0(opts, absolutePath)
	if err != nil {
		return precompiled.DefaultErrorCode, nil, fmt.Errorf("service List0 failed: %+v", err)
	}
	return int64(ret0), ret1, err
}

func (service *Service) List(absolutePath string, offset *big.Int, limit *big.Int) (int64, []BfsInfo, error) {
	opts := &bind.CallOpts{From: service.bfsAuth.From}
	ret0, ret1, err := service.bfs.List(opts, absolutePath, offset, limit)
	if err != nil {
		return precompiled.DefaultErrorCode, nil, fmt.Errorf("service List failed: %+v", err)
	}
	return ret0.Int64(), ret1, err
}

func (service *Service) AsyncMkdir(handler func(*types.Receipt, error), absolutePath string) (*types.Transaction, error) {
	return service.bfs.AsyncMkdir(handler, service.bfsAuth, absolutePath)
}

func (service *Service) Mkdir(absolutePath string) (int64, error) {
	ret0, _, _, err := service.bfs.Mkdir(service.bfsAuth, absolutePath)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("service Mkdir failed: %+v", err)
	}
	return int64(ret0), err
}

func (service *Service) Link(absolutePath string, _address string, _abi string) (int64, error) {
	ret0, _, _, err := service.bfs.Link(service.bfsAuth, absolutePath, _address, _abi)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("service Link failed: %+v", err)
	}
	return ret0.Int64(), err
}

func (service *Service) AsyncLink(handler func(*types.Receipt, error), absolutePath string, _address string, _abi string) (*types.Transaction, error) {
	return service.bfs.AsyncLink(handler, service.bfsAuth, absolutePath, _address, _abi)
}

func (service *Service) Link0(name string, version string, _address string, _abi string) (int64, error) {
	ret0, _, _, err := service.bfs.Link0(service.bfsAuth, name, version, _address, _abi)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("service Link0 failed: %+v", err)
	}
	return int64(ret0), err
}

func (service *Service) AsyncLink0(handler func(*types.Receipt, error), name string, version string, _address string, _abi string) (*types.Transaction, error) {
	return service.bfs.AsyncLink0(handler, service.bfsAuth, name, version, _address, _abi)
}

func (service *Service) Readlink(absolutePath string) (common.Address, error) {
	opts := &bind.CallOpts{From: service.bfsAuth.From}
	ret0, err := service.bfs.Readlink(opts, absolutePath)
	if err != nil {
		return common.Address{}, fmt.Errorf("service Readlink failed: %+v", err)
	}
	return ret0, err
}

// func (service *Service) Touch(absolutePath string, fileType string) (int64, error) {
// 	ret0, _, _, err := service.bfs.Touch(service.bfsAuth, absolutePath, fileType)
// 	if err != nil {
// 		return precompiled.DefaultErrorCode, fmt.Errorf("service Touch failed: %+v", err)
// 	}
// 	return int64(ret0), err
// }

// func (service *Service) AsyncTouch(handler func(*types.Receipt, error), absolutePath string, fileType string) (*types.Transaction, error) {
// 	opts := &bind.CallOpts{From: service.bfsAuth.From}
// 	return service.bfs.AsyncTouch(handler, opts, absolutePath, fileType)
// }

// func (service *Service) RebuildBfs() (*big.Int, error) {
// 	opts := &bind.CallOpts{From: service.bfsAuth.From}
// 	ret0, _, _, err := service.bfs.RebuildBfs(opts)
// 	if err != nil {
// 		return common.Address{}, fmt.Errorf("service RebuildBfs failed: %+v", err)
// 	}
// 	return ret0, err
// }

// func (service *Service) AsyncRebuildBfs(handler func(*types.Receipt, error)) (*types.Transaction, error) {
// 	opts := &bind.CallOpts{From: service.bfsAuth.From}
// 	return service.bfs.AsyncRebuildBfs(handler, opts)
// }

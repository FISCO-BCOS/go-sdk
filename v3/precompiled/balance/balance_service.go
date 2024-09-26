package balanceprecompiled

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Service struct {
	balance                *BalancePrecompiled
	BalancePrecompiledAuth *bind.TransactOpts
	client                 *client.Client
}

var balancePrecompiledAddress = common.HexToAddress("0000000000000000000000000000000000001011")

func NewBalanceService(client *client.Client) (*Service, error) {
	instance, err := NewBalancePrecompiled(balancePrecompiledAddress, client)
	if err != nil {
		return nil, err
	}
	auth := client.GetTransactOpts()
	return &Service{balance: instance, BalancePrecompiledAuth: auth, client: client}, nil
}

func (service *Service) GetBalance(account string) (*big.Int, error) {
	opts := &bind.CallOpts{From: service.BalancePrecompiledAuth.From}
	ret0, err := service.balance.GetBalance(opts, common.HexToAddress(account))
	if err != nil {
		return big.NewInt(0), err
	}
	return ret0, nil
}

func (service *Service) Transfer(from string, to string, value *big.Int) (receipt *types.Receipt, err error) {
	_, receipt, err = service.balance.Transfer(service.BalancePrecompiledAuth, common.HexToAddress(from), common.HexToAddress(to), value)
	if err != nil {
		return nil, fmt.Errorf("service Transfer failed: %+v", err)
	}
	return receipt, nil
}

func (service *Service) addBalance(account string, value *big.Int) (receipt *types.Receipt, err error) {
	_, receipt, err = service.balance.AddBalance(service.BalancePrecompiledAuth, common.HexToAddress(account), value)
	if err != nil {
		return nil, fmt.Errorf("service addBalance failed: %+v", err)
	}
	return receipt, nil
}

func (service *Service) subBalance(account string, value *big.Int) (receipt *types.Receipt, err error) {
	_, receipt, err = service.balance.SubBalance(service.BalancePrecompiledAuth, common.HexToAddress(account), value)
	if err != nil {
		return nil, fmt.Errorf("service subBalance failed: %+v", err)
	}
	return receipt, nil
}

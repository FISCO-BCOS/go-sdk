package permission

import (
	"fmt"
	"crypto/ecdsa"
	"math/big"
	"context"
	"encoding/json"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/precompile/crud"
)

const (
	SysConsensus = "_sys_consensus_"
	SysCNS = "_sys_cns_"
	SysTableAccess = "_sys_table_access_"
	SysConfig = "_sys_config_"
)

// PermissionService is a precompile contract service.
type PermissionService struct {
	permission *Permission
	permissionAuth *bind.TransactOpts
	client *client.Client
	privateKey *ecdsa.PrivateKey
}

// PermissionPrecompileAddress is the contract address of Permission
var PermissionPrecompileAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000001005")

// NewPermissionService returns ptr of PermissionService
func NewPermissionService(client *client.Client, privateKey *ecdsa.PrivateKey) (*PermissionService, error) {
	instance, err := NewPermission(PermissionPrecompileAddress, client)
	if err != nil {
		return nil, fmt.Errorf("construct PermissionService failed: %+v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = big.NewInt(30000000)
    return &PermissionService{permission:instance, permissionAuth:auth, client: client, privateKey:privateKey}, nil
}

// GrantUserTableManager grants the info by the table name and user address
func (service *PermissionService) GrantUserTableManager(tableName string, grantress string) (string, error) {
	crudService,err := crud.NewCRUDService(service.client, service.privateKey)
	if err != nil {
		return "", fmt.Errorf("PermissionService create CRUDService failed: %v", err)
	}
	_, err = crudService.Desc(tableName)
    if err != nil {
		return "",fmt.Errorf("GrantUserTableManager failed: %v", err)
	}
    return service.grant(tableName, grantress)
}

// RevokeUserTableManager revokes a grantress' right of the table name
func (service *PermissionService) RevokeUserTableManager(tableName string, grantress string) (string, error) {
    return service.revoke(tableName, grantress)
}

// ListUserTableManager returns the list of permission info
func (service *PermissionService) ListUserTableManager(tableName string) ([]PermissionInfo, error) {
	return service.list(tableName)
}

// GrantDeployAndCreateManager grants the deploy and create option to an address
func (service *PermissionService) GrantDeployAndCreateManager(grantress string) (string, error) {
	return service.grant(crud.SysTable, grantress)
}

// RevokeDeployAndCreateManager revokes a grantress's right of the deploy and create option
func (service *PermissionService) RevokeDeployAndCreateManager(grantress string) (string, error) {
	return service.revoke(crud.SysTable, grantress)
}


// ListDeployAndCreateManager returns the list of permission info
func (service *PermissionService) ListDeployAndCreateManager() ([]PermissionInfo, error) {
	return service.list(crud.SysTable)
}

// GrantPermissionManager grants the permission
func (service *PermissionService) GrantPermissionManager(grantress string) (string, error) {
	return service.grant(SysTableAccess, grantress)
}

// RevokePermissionManager revokes the permission
func (service *PermissionService) RevokePermissionManager(grantress string) (string, error) {
	return service.revoke(SysTableAccess, grantress)
}

// ListPermissionManager returns the list of permission
func (service *PermissionService) ListPermissionManager() ([]PermissionInfo, error) {
	return service.list(SysTableAccess)
}

// GrantNodeManager grants the Node
func (service *PermissionService) GrantNodeManager(grantress string) (string, error) {
	return service.grant(SysConsensus, grantress)
}

// RevokeNodeManager revokes the Node
func (service *PermissionService) RevokeNodeManager(grantress string ) (string, error) {
	return service.revoke(SysConsensus, grantress)
}

// ListNodeManager returns the list of Node manager
func (service *PermissionService) ListNodeManager() ([]PermissionInfo, error) {
	return service.list(SysConsensus)
}

// GrantCNSManager grants the CNS
func (service *PermissionService) GrantCNSManager(grantress string ) (string, error) {
	return service.grant(SysCNS, grantress)
}

// RevokeCNSManager revokes the CNS
func (service *PermissionService) RevokeCNSManager(grantress string ) (string, error) {
	return service.revoke(SysCNS, grantress)
}

// ListCNSManager returns the list of CNS manager
func (service *PermissionService) ListCNSManager() ([]PermissionInfo, error) {
	return service.list(SysCNS)
}

// GrantSysConfigManager grants the System configuration manager
func (service *PermissionService) GrantSysConfigManager(grantress string ) (string, error) {
	return service.grant(SysConfig, grantress)
}

// RevokeSysConfigManager revokes the System configuration manager
func (service *PermissionService) RevokeSysConfigManager(grantress string ) (string, error) {
	return service.revoke(SysConfig, grantress)
}

// ListSysConfigManager returns the list of System configuration manager
func (service *PermissionService) ListSysConfigManager() ([]PermissionInfo, error) {
	return service.list(SysConfig)
}

func (service *PermissionService) grant(tableName string, grantress string) (string, error) {
	tx, err := service.permission.Insert(service.permissionAuth, tableName, grantress)
	if err != nil {
		return "", fmt.Errorf("PermissionService grant failed: %v", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), service.client, tx)
	if err != nil {
        return "", fmt.Errorf("PermissionService wait for the transaction receipt failed: %v", err)
	}
	return handleReceipt(receipt)
}

func (service *PermissionService) revoke(tableName string, address string) (string, error) {
	tx, err := service.permission.Remove(service.permissionAuth, tableName, address)
	if err != nil {
		return "", fmt.Errorf("PermissionService revoke failed: %v", err)
	}
	// wait for the mining
	receipt, err := bind.WaitMined(context.Background(), service.client, tx)
	if err != nil {
        return "", fmt.Errorf("PermissionService wait for the transaction receipt failed: %v", err)
	}
	return handleReceipt(receipt)
}

func (service *PermissionService) list(tableName string) ([]PermissionInfo, error) {
	opts := &bind.CallOpts{From: service.permissionAuth.From}
	permissionyInfo, err := service.permission.QueryByName(opts, tableName)
	if err != nil {
		return nil, fmt.Errorf("PermissionService List failed: %v", err)
	}
	// unmarshal result
	var results []PermissionInfo
	if err := json.Unmarshal([]byte(permissionyInfo), &results); err != nil {
		return nil, fmt.Errorf("PermissionService: Unmarshal the List result failed: %v", err)
	}
	return results, nil
}

func handleReceipt(receipt *types.Receipt) (string, error) {
	status := receipt.GetStatus()
	if "0x0" != status {
		return "", fmt.Errorf(common.GetStatusMessage(status))
	} 
	output := receipt.GetOutput()
	if output != "" {
		return common.GetJsonStr(output)
	}
	return "", fmt.Errorf("Transaction is handled failure")
}
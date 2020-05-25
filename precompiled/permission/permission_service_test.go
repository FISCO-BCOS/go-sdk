package permission

import (
	"crypto/ecdsa"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	success        = "{\"code\":0,\"msg\":\"success\"}"
	tableName      = "t_test"
	permisstionAdd = "0xFbb18d54e9Ee57529cda8c7c52242EFE879f064F"
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

func GetService(t *testing.T) *PermissionService {
	c := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewPermissionService(c, privateKey)
	if err != nil {
		t.Fatalf("init PermissionService failed: %+v", err)
	}
	return service
}

func TestGrant(t *testing.T) {
	service := GetService(t)
	// grant permission
	result, err := service.GrantPermissionManager(permisstionAdd)
	if err != nil {
		t.Fatalf("TestPermissionManager failed: %v", err)
	}
	t.Logf("TestPermissionManager: %v", result)

	listResult, err := service.ListPermissionManager()
	if err != nil {
		t.Fatalf("ListPermissionManager failed: %v", err)
	}
	t.Logf("ListPermissionManager: %+v", listResult)

	result, err = service.RevokePermissionManager(permisstionAdd)
	if err != nil {
		t.Fatalf("RevokePermissionManager failed: %v", err)
	}
	t.Logf("RevokePermissionManager: %v", result)

	listResult, err = service.ListPermissionManager()
	if err != nil {
		t.Fatalf("ListPermissionManager failed: %v", err)
	}
	t.Logf("ListPermissionManager: %v", listResult)
}

// func TestUserTableManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantUserTableManager(tableName, permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestUserTableManager failed: %v", err)
// 	}
// 	t.Logf("TestUserTableManager: %v", result)
// 	revokeResult, err := service.RevokeUserTableManager(tableName, permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestUserTableManager failed: %v", err)
// 	}
// 	t.Logf("TestUserTableManager revoke result: %v", revokeResult)
// }

// func TestDeployAndCreateManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantDeployAndCreateManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestDeployAndCreateManager failed: %v", err)
// 	}
// 	t.Logf("TestDeployAndCreateManager: %v", result)

// 	revokeResult, err := service.RevokeDeployAndCreateManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestDeployAndCreateManager failed: %v", err)
// 	}
// 	t.Logf("TestDeployAndCreateManager revoke result: %v", revokeResult)
// }

// func TestNodeManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantNodeManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestNodeManager failed: %v", err)
// 	}
// 	t.Logf("TestNodeManager: %v", result)

// 	revokeResult, err := service.RevokeNodeManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestNodeManager failed: %v", err)
// 	}
// 	t.Logf("TestNodeManager revoke result: %v", revokeResult)
// }

// func TestCNSManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantCNSManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestCNSManager failed: %v", err)
// 	}
// 	t.Logf("TestCNSManager: %v", result)

// 	revokeResult, err := service.RevokeCNSManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestCNSManager failed: %v", err)
// 	}
// 	t.Logf("TestCNSManager revoke result: %v", revokeResult)
// }

// func TestSysConfigManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantSysConfigManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestSysConfigManager failed: %v", err)
// 	}
// 	t.Logf("TestSysConfigManager: %v", result)

// 	revokeResult, err := service.RevokeSysConfigManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestSysConfigManager failed: %v", err)
// 	}
// 	t.Logf("TestSysConfigManager revoke result: %v", revokeResult)
// 	t.Logf("Success result: %s", success)
// }

// func TestListUser(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.ListUserTableManager(tableName)
// 	if err != nil {
// 		t.Fatalf("ListUserTableManager failed: %v", err)
// 	}
// 	t.Logf("ListUserTableManager: %v", result)
// }

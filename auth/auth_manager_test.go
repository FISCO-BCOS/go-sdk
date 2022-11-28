package auth

import (
	"encoding/hex"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"testing"
)

var (
	service *AuthManagerService
)

func getClient(t *testing.T) *client.Client {
	privateKey, _ := hex.DecodeString("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62")
	config := &conf.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, NodeURL: "http://localhost:8545"}
	c, err := client.Dial(config)
	if err != nil {
		t.Fatalf("Dial to %s failed of %v", config.NodeURL, err)
	}
	return c
}

func getService(t *testing.T) {
	c := getClient(t)
	newService, err := NewAuthManagerService(c)
	if err != nil {
		t.Fatalf("init AuthManagerService failed: %+v", err)
	}
	service = newService
}

func TestAuthManagerService_GetDeployAuthType(t *testing.T) {
	t.Logf("starting test ....................")

	result, err := service.GetDeployAuthType()
	if err != nil {
		t.Fatalf("GetDeployAuthType failed: %v", err)
	}

	t.Logf("GetDeployAuthType: %v", result)
}

package main

import (
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/auth"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"log"
)

var (
	service *auth.AuthManagerService
)

func main() {
	fmt.Println("starting test..........................")
	getService()

	TestAuthManagerService_GetCommitteeInfo()

}

func getClient() *client.Client {
	privateKey, _ := hex.DecodeString("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62")
	config := &conf.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, NodeURL: "127.0.0.1:20200"}
	c, err := client.Dial(config)
	if err != nil {
		log.Fatalf("Dial to %s failed of %v", config.NodeURL, err)
	}
	return c
}

func getService() {
	c := getClient()
	newService, err := auth.NewAuthManagerService(c)
	if err != nil {
		log.Fatalf("init AuthManagerService failed: %+v", err)
	}
	service = newService
}

func TestAuthManagerService_GetDeployAuthType() {
	log.Println("starting test TestAuthManagerService_GetDeployAuthType ....................")

	result, err := service.GetDeployAuthType()
	if err != nil {
		log.Fatalf("GetDeployAuthType failed: %v", err)
	}

	log.Printf("GetDeployAuthType: %v", result)
}

func TestAuthManagerService_GetCommitteeInfo() {
	log.Println("starting test TestAuthManagerService_GetCommitteeInfo ....................")

	result, err := service.GetCommitteeInfo()
	if err != nil {
		log.Fatalf("GetCommitteeInfo failed: %v", err)
	}

	log.Printf("GetCommitteeInfo: %v", result)
}

package main

import (
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/auth"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

var (
	service *auth.AuthManagerService
)

func main() {
	fmt.Println("starting test..........................")
	getService()

	TestAuthManagerService_GetAdmin()

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

func TestAuthManagerService_CheckDeployAuth() {
	log.Println("starting test TestAuthManagerService_CheckDeployAuth ....................")

	result, err := service.CheckDeployAuth(common.HexToAddress("0x83309d045a19c44Dc3722D15A6AbD472f95866bC"))
	if err != nil {
		log.Fatalf("CheckDeployAuth failed: %v", err)
	}

	log.Printf("CheckDeployAuth: %v", *result)
}

func TestAuthManagerService_GetAdmin() {
	log.Println("starting test TestAuthManagerService_GetAdmin ....................")

	//result, err := service.GetAdmin(common.HexToAddress("0000000000000000000000000000000000001005"))
	result, err := service.GetAdmin(common.HexToAddress("0000000000000000000000000000000000010001"))
	if err != nil {
		log.Fatalf("GetAdmin failed: %v", err)
	}

	log.Printf("GetAdmin: %v", result)
}

func TestAuthManagerService_GetCommitteeInfo() {
	log.Println("starting test TestAuthManagerService_GetCommitteeInfo ....................")

	result, err := service.GetCommitteeInfo()
	if err != nil {
		log.Fatalf("GetCommitteeInfo failed: %v", err)
	}

	log.Printf("GetCommitteeInfo: %v", result)
}

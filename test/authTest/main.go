package main

import (
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/auth"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

var (
	service *auth.AuthManagerService
)

func main() {
	fmt.Println("starting test..........................")
	getService()

	//TestAuthManagerService_GetDeployAuthType()
	//TestAuthManagerService_SetDeployAuthType()

	//TestAuthManagerService_ModifyDeployAuth()

	//TestAuthManagerService_RevokeProposal()
	//TestAuthManagerService_VoteProposal()

	TestAuthManagerService_SetMethodAuthType()

}

func getClient() *client.Client {
	//privateKey, _ := hex.DecodeString("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62")
	privateKey, _ := hex.DecodeString("8f21f97898615eb58c24f8310e7ee3fae148ebc03d14938c2ebc87587129e44d")
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

// todo
func TestAuthManagerService_GetCommitteeInfo() {
	log.Println("starting test TestAuthManagerService_GetCommitteeInfo ....................")

	result, err := service.GetCommitteeInfo()
	if err != nil {
		log.Fatalf("GetCommitteeInfo failed: %v", err)
	}

	log.Printf("GetCommitteeInfo: %v", result)
}

// todo
func TestAuthManagerService_GetAdmin() {
	log.Println("starting test TestAuthManagerService_GetAdmin ....................")

	result, err := service.GetAdmin(common.HexToAddress("0x1005"))
	//result, err := service.GetAdmin(common.HexToAddress("0000000000000000000000000000000000010001"))
	if err != nil {
		log.Fatalf("GetAdmin failed: %v", err)
	}

	log.Printf("GetAdmin: %v", result)
}

// todo
func TestAuthManagerService_ResetAdmin() {
	log.Println("starting test TestAuthManagerService_GetAdmin ....................")

	result, err := service.ResetAdmin(common.HexToAddress("0xe2b91bb57b43239788740295db49301382d05021"), common.HexToAddress("0000000000000000000000000000000000010001"))
	if err != nil {
		log.Fatalf("ResetAdmin failed: %v", err)
	}

	log.Printf("ResetAdmin: %v", result)
}

func TestAuthManagerService_UpdateGovernor() {
	log.Println("starting test TestAuthManagerService_UpdateGovernor ....................")

	result, err := service.UpdateGovernor(common.HexToAddress("0xe2b91bb57b43239788740295db49301382d05021"), 2)
	if err != nil {
		log.Fatalf("UpdateGovernor failed: %v", err)
	}

	log.Printf("UpdateGovernor: %v", result)
}

func TestAuthManagerService_SetRate() {
	log.Println("starting test TestAuthManagerService_SetRate ....................")

	result, err := service.SetRate(1, 2)
	if err != nil {
		log.Fatalf("SetRate failed: %v", err)
	}

	log.Printf("SetRate: %v", result)
}

func TestAuthManagerService_SetDeployAuthType() {
	log.Println("starting test TestAuthManagerService_SetDeployAuthType ....................")

	TestAuthManagerService_GetDeployAuthType()

	result, err := service.SetDeployAuthType(1)
	if err != nil {
		log.Fatalf("SetDeployAuthType failed: %v", err)
	}

	log.Printf("SetDeployAuthType: %v", result)
}

func TestAuthManagerService_ModifyDeployAuth() {
	log.Println("starting test TestAuthManagerService_ModifyDeployAuth ....................")

	TestAuthManagerService_GetDeployAuthType()

	result, err := service.ModifyDeployAuth(common.HexToAddress("0xe2b91bb57b43239788740295db49301382d05021"), true)
	if err != nil {
		log.Fatalf("ModifyDeployAuth failed: %v", err)
	}

	log.Printf("ModifyDeployAuth: %v", result)
}

func TestAuthManagerService_RevokeProposal() {
	log.Println("starting test TestAuthManagerService_RevokeProposal ....................")

	result, err := service.RevokeProposal(*big.NewInt(5))
	if err != nil {
		log.Fatalf("RevokeProposal failed: %v", err)
	}

	log.Printf("RevokeProposal: %v", result)
}

func TestAuthManagerService_VoteProposal() {
	log.Println("starting test TestAuthManagerService_VoteProposal ....................")

	result, err := service.VoteProposal(*big.NewInt(5), false)
	if err != nil {
		log.Fatalf("VoteProposal failed: %v", err)
	}

	log.Printf("VoteProposal: %v", result)
}

func TestAuthManagerService_SetMethodAuthType() {
	log.Println("starting test TestAuthManagerService_SetMethodAuthType ....................")

	result, err := service.SetMethodAuthType(common.HexToAddress("0000000000000000000000000000000000001005"), [4]byte{}, 1)
	if err != nil {
		log.Fatalf("SetMethodAuthType failed: %v", err)
	}

	log.Printf("SetMethodAuthType: %v", result)
}

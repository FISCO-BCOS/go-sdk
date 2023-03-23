package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/precompiled/auth"
	"github.com/ethereum/go-ethereum/common"
)

var (
	service *auth.AuthManagerService
)

func main() {
	fmt.Println("starting test..........................")
	getService()

	//GetDeployAuthType()
	//SetDeployAuthType()

	//ModifyDeployAuth()

	//RevokeProposal()
	//VoteProposal()

	//SetMethodAuthType()

	//GetAdmin()
	//ResetAdmin()

	CheckMethodAuth()
}

func getClient() *client.Client {
	//privateKey, _ := hex.DecodeString("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62")
	privateKey, _ := hex.DecodeString("8f21f97898615eb58c24f8310e7ee3fae148ebc03d14938c2ebc87587129e44d")
	config := &client.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}
	c, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatalf("Dial to %s:%d failed of %v", config.Host, config.Port, err)
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

func GetDeployAuthType() {
	log.Println("starting test GetDeployAuthType ....................")

	result, err := service.GetDeployAuthType()
	if err != nil {
		log.Fatalf("GetDeployAuthType failed: %v", err)
	}

	log.Printf("GetDeployAuthType: %v", result)
}

func CheckDeployAuth() {
	log.Println("starting test CheckDeployAuth ....................")

	result, err := service.CheckDeployAuth(common.HexToAddress("0x83309d045a19c44Dc3722D15A6AbD472f95866bC"))
	if err != nil {
		log.Fatalf("CheckDeployAuth failed: %v", err)
	}

	log.Printf("CheckDeployAuth: %v", *result)
}

// todo
func GetCommitteeInfo() {
	log.Println("starting test GetCommitteeInfo ....................")

	result, err := service.GetCommitteeInfo()
	if err != nil {
		log.Fatalf("GetCommitteeInfo failed: %v", err)
	}

	log.Printf("GetCommitteeInfo: %v", result)
}

func CheckMethodAuth() {
	log.Println("starting test CheckMethodAuth ....................")

	// console
	//	[group0]: /apps> deploy ContractAuthPrecompiled
	//	transaction hash: 0x1fa19bbff520ec13c67b5cd8ed202c6a0bf656b4f2790000e94b7de1c3ba065e
	//	contract address: 0x41a1281dba209614f2ada8ecc75fd957ad179d7b
	//	currentAccount: 0x357d2f663c8868b777eccc69a7bc8a9d7e4862ce
	//	[group0]: /apps>

	var funcs [4]byte
	//0xc53057b4
	funcs[0] = 0xc5
	funcs[1] = 0x30
	funcs[2] = 0x57
	funcs[3] = 0xb4

	result, err := service.CheckMethodAuth(common.HexToAddress("0x41a1281dba209614f2ada8ecc75fd957ad179d7b"), funcs, common.HexToAddress("0x357d2f663c8868b777eccc69a7bc8a9d7e4862ce"))
	if err != nil {
		log.Fatalf("CheckMethodAuth failed: %v", err)
	}

	log.Printf("CheckMethodAuth: %v", *result)
}

func GetAdmin() {
	log.Println("starting test GetAdmin ....................")

	// console
	//	[group0]: /apps> deploy ContractAuthPrecompiled
	//	transaction hash: 0x1fa19bbff520ec13c67b5cd8ed202c6a0bf656b4f2790000e94b7de1c3ba065e
	//	contract address: 0x41a1281dba209614f2ada8ecc75fd957ad179d7b
	//	currentAccount: 0x357d2f663c8868b777eccc69a7bc8a9d7e4862ce
	//	[group0]: /apps>
	result, err := service.GetAdmin(common.HexToAddress("0x41a1281dba209614f2ada8ecc75fd957ad179d7b"))
	if err != nil {
		log.Fatalf("GetAdmin failed: %v", err)
	}

	log.Printf("GetAdmin: %v", result)
}

func ResetAdmin() {
	log.Println("starting test GetAdmin ....................")

	result, err := service.ResetAdmin(common.HexToAddress("0xe2b91bb57b43239788740295db49301382d05021"), common.HexToAddress("0x41a1281dba209614f2ada8ecc75fd957ad179d7b"))
	if err != nil {
		log.Fatalf("ResetAdmin failed: %v", err)
	}

	log.Printf("ResetAdmin: %v", result)
}

func UpdateGovernor() {
	log.Println("starting test UpdateGovernor ....................")

	result, err := service.UpdateGovernor(common.HexToAddress("0xe2b91bb57b43239788740295db49301382d05021"), 2)
	if err != nil {
		log.Fatalf("UpdateGovernor failed: %v", err)
	}

	log.Printf("UpdateGovernor: %v", result)
}

func SetRate() {
	log.Println("starting test SetRate ....................")

	result, err := service.SetRate(1, 2)
	if err != nil {
		log.Fatalf("SetRate failed: %v", err)
	}

	log.Printf("SetRate: %v", result)
}

func SetDeployAuthType() {
	log.Println("starting test SetDeployAuthType ....................")

	GetDeployAuthType()

	result, err := service.SetDeployAuthType(1)
	if err != nil {
		log.Fatalf("SetDeployAuthType failed: %v", err)
	}

	log.Printf("SetDeployAuthType: %v", result)
}

func ModifyDeployAuth() {
	log.Println("starting test ModifyDeployAuth ....................")

	GetDeployAuthType()

	result, err := service.ModifyDeployAuth(common.HexToAddress("0xe2b91bb57b43239788740295db49301382d05021"), true)
	if err != nil {
		log.Fatalf("ModifyDeployAuth failed: %v", err)
	}

	log.Printf("ModifyDeployAuth: %v", result)
}

func RevokeProposal() {
	log.Println("starting test RevokeProposal ....................")

	result, err := service.RevokeProposal(*big.NewInt(5))
	if err != nil {
		log.Fatalf("RevokeProposal failed: %v", err)
	}

	log.Printf("RevokeProposal: %v", result)
}

func VoteProposal() {
	log.Println("starting test VoteProposal ....................")

	result, err := service.VoteProposal(*big.NewInt(5), false)
	if err != nil {
		log.Fatalf("VoteProposal failed: %v", err)
	}

	log.Printf("VoteProposal: %v", result)
}

func SetMethodAuthType() {
	log.Println("starting test SetMethodAuthType ....................")

	result, err := service.SetMethodAuthType(common.HexToAddress("0000000000000000000000000000000000001005"), [4]byte{}, 1)
	if err != nil {
		log.Fatalf("SetMethodAuthType failed: %v", err)
	}

	log.Printf("SetMethodAuthType: %v", result)
}

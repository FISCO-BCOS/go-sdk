package auth

import (
	"context"
	"encoding/hex"
	"strconv"
	"fmt"
	"os"
	"testing"
	"math/big"
	"github.com/FISCO-BCOS/go-sdk/client"
	// "github.com/FISCO-BCOS/go-sdk/core/types"
	// "github.com/FISCO-BCOS/go-sdk/precompiled"
	"github.com/ethereum/go-ethereum/common"
)

const (
	accountAddress_string = "0xfe5625acd8b8effbf87ef65f9ed9ddc3390114f5"	//治理委员，contractAddress的管理员
	normalAccountAddress_string = "0x4fdc7f1e05e48b4b252df0e815dbe107935c8618"
	funcSelector_string = "4ed3885e" //set(string)
	status = uint8(1)
	authType = uint8(2)	// 1:white_list. 2:black_list
	isOpen = false
	weight = 5
	participatesRate = uint8(50)
	winRate	= uint8(60)
	openFlag = false
)

var (
	service *AuthManagerService
	channel = make(chan int)
	proposalManager_string string
)

func getClient(t *testing.T) *client.Client {
	// privatekey of 0x83309d045a19c44dc3722d15a6abd472f95866ac
	privateKey, _ := hex.DecodeString("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62")
	config := &client.Config{IsSMCrypto: false, GroupID: "group0",
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./ca.crt", TLSKeyFile: "./sdk.key", TLSCertFile: "./sdk.crt"}

	c, err := client.DialContext(context.Background(), config)
	if err != nil {
		t.Fatalf("Dial to %s:%d failed of %v", config.Host, config.Port, err)
	}

	// deploy helloWorld contract
	address_common := common.HexToAddress(accountAddress_string)
	address, _, _, _ := DeployProposalManager(c.GetTransactOpts(), c, address_common, address_common)
	proposalManager_string = address.String()
	fmt.Println("proposalManager_string",proposalManager_string)

	return c
}

func getService(t *testing.T) {
	c := getClient(t)
	newService, err := NewAuthManagerService(c)
	if err != nil {
		t.Fatalf("init CrudService failed: %+v", err)
	}
	service = newService
}

func TestMain(m *testing.M) {
	getService(&testing.T{})
	exitCode := m.Run()
	os.Exit(exitCode)
}

/**
 * **************************************************************************************************************
 * AccountManager
 * **************************************************************************************************************
**/

/**
 * *************************************************
 * 无需权限的查询接口
 * *************************************************
**/



func TestGetAccountStatus(t *testing.T) {
	address_common := common.HexToAddress(accountAddress_string)

	ret0, err := service.GetAccountStatus(address_common)
	if err != nil {
		t.Fatalf("TestGetAccountStatus failed: %v", err)
	}

	t.Logf("ret0: %v\n", ret0)
}

/**
 * *************************************************
 * 治理委员账号专用接口
 * *************************************************
**/

func TestSetAccountStatus(t *testing.T) {
	address_common := common.HexToAddress(normalAccountAddress_string)
	ret0, err := service.SetAccountStatus(address_common, status)
	
	if err != nil {
		t.Fatalf("TestSetAccountStatus failed: %v", err)
	}
	if ret0 != 0 {
		t.Fatalf("TestSetAccountStatus failed, the ret0 \"%v\" is inconsistent with \"0\"", ret0)
	}

	tempStatus, err := service.GetAccountStatus(address_common)
	if tempStatus != status {
		t.Fatalf("TestSetAccountStatus failed, the tempStatus \"%v\" is inconsistent with the status \"%v\"", tempStatus, status)
	}

	t.Logf("status: %v\n", status)
	t.Logf("tempStatus: %v\n", tempStatus)
}

/**
 * **************************************************************************************************************
 * Committee
 * **************************************************************************************************************
**/

/**
 * *************************************************
 * 无需权限的查询接口
 * *************************************************
**/

func TestGetCommitteeInfo(t *testing.T) {
	ret0, err := service.GetCommitteeInfo()
	if err != nil {
		t.Fatalf("TestGetCommitteeInfo failed: %v", err)
	}

	t.Logf("ret0: %v\n", ret0)
}

/**
 * **************************************************************************************************************
 * ProposalManager
 * **************************************************************************************************************
**/

/**
 * *************************************************
 * 无需权限的查询接口
 * *************************************************
**/

func TestGetProposalInfo(t *testing.T) {
	proposalId := big.NewInt(1)
	ret0, err := service.GetProposalInfo(proposalId)
	if err != nil {
		t.Fatalf("TestGetProposalInfo failed: %v", err)
	}

	t.Logf("ret0: %v\n", ret0)
}

/**
 * **************************************************************************************************************
 * ContractAuth
 * **************************************************************************************************************
**/

/**
 * *************************************************
 * 无需权限的查询接口
 * *************************************************
**/

func TestGetDeployAuthType(t *testing.T) {
	ret0, err := service.GetDeployAuthType()
	if err != nil {
		t.Fatalf("TestGetDeployAuthType failed: %v", err)
	}

	t.Logf("ret0: %v\n", ret0)
}

func TestCheckDeployAuth(t *testing.T) {
	address_common := common.HexToAddress(accountAddress_string)
	ret0, err := service.CheckDeployAuth(address_common)
	if err != nil {
		t.Fatalf("TestCheckDeployAuth failed: %v", err)
	}

	t.Logf("ret0: %v\n", *ret0)
}

func TestCheckMethodAuth(t *testing.T) {
	accountAddress_common := common.HexToAddress(accountAddress_string)
	contractAddress_common := common.HexToAddress(proposalManager_string)
	funcSelector := StringToByteList_FuncSelector(funcSelector_string)

	ret0, err := service.CheckMethodAuth(contractAddress_common,funcSelector,accountAddress_common)
	if err != nil {
		t.Fatalf("TestCheckMethodAuth failed: %v", err)
	}

	t.Logf("ret0: %v\n", *ret0)
}

func TestGetAdmin(t *testing.T) {
	contractAddress_common := common.HexToAddress(proposalManager_string)
	ret0, err := service.GetAdmin(contractAddress_common)
	if err != nil {
		t.Fatalf("TestGetAdmin failed: %v", err)
	}

	t.Logf("ret0: %v\n", *ret0)
}


/**
 * *************************************************
 * 合约管理员账号专用接口
 * *************************************************
**/

func TestSetMethodAuthType(t *testing.T) {
	// funcSelector_string: "4ed3885e"
	accountAddress_common := common.HexToAddress(accountAddress_string)
	contractAddress_common := common.HexToAddress(proposalManager_string)
	fmt.Println("proposalManager_string",proposalManager_string)

	tempByte, err := hex.DecodeString(funcSelector_string)
	var funcSelector [4]byte
	copy(funcSelector[:],tempByte)
	fmt.Println("funcSelector",funcSelector)

	ret0, err := service.SetMethodAuthType(contractAddress_common, funcSelector, authType)
	
	if err != nil {
		t.Fatalf("TestSetMethodAuthType failed: %v", err)
	}
	if ret0.Int64() != 0 {
		t.Fatalf("TestSetMethodAuthType failed, the ret0 \"%v\" is inconsistent with \"0\"", ret0)
	}

	tempBool:=true
	if authType==1 {
		tempBool=false
	}

	ret1, err := service.CheckMethodAuth(contractAddress_common,funcSelector,accountAddress_common)
	if *ret1 != tempBool {
		t.Fatalf("TestSetMethodAuthType failed, the ret1 \"%v\" is inconsistent with the tempBool \"%v\"", *ret1, tempBool)
	}

	t.Logf("authType: %v\n", authType)
	t.Logf("ret1: %v\n", *ret1)
}

func TestSetMethodAuth(t *testing.T) {
	accountAddress_common := common.HexToAddress(accountAddress_string)
	contractAddress_common := common.HexToAddress(proposalManager_string)

	tempByte, err := hex.DecodeString(funcSelector_string)
	var funcSelector [4]byte
	copy(funcSelector[:],tempByte)
	fmt.Println("funcSelector",funcSelector)

	ret0, err := service.SetMethodAuth(contractAddress_common, funcSelector, accountAddress_common, isOpen)
	
	if err != nil {
		t.Fatalf("TestSetMethodAuth failed: %v", err)
	}
	if ret0.Int64() != 0 {
		t.Fatalf("TestSetMethodAuth failed, the ret0 \"%v\" is inconsistent with \"0\"", ret0)
	}

	ret1, err := service.CheckMethodAuth(contractAddress_common,funcSelector,accountAddress_common)
	if *ret1 != isOpen {
		t.Fatalf("TestSetMethodAuthType failed, the ret1 \"%v\" is inconsistent with the isOpen \"%v\"", *ret1, isOpen)
	}

	t.Logf("isOpen: %v\n", isOpen)
	t.Logf("ret1: %v\n", *ret1)
}

func StringToByteList_FuncSelector(s string) [4]byte {
	byteList := [4]byte{}
	index:=0
	for i := 0; i < len(s); i+=2 {
        tempStr:=s[i:i+2]
		byteList[index]=Hex2Dec(tempStr)
		index++
	}
	return byteList

}

func Hex2Dec(val string) uint8 {
	n, err := strconv.ParseUint(val, 16, 8)
	if err != nil {
		fmt.Println(err)
	}
	return uint8(n)
}

/**
 * **************************************************************************************************************
 * CommitteeManager
 * **************************************************************************************************************
**/

/**
 * *************************************************
 * 治理委员账号专用接口
 * *************************************************
**/

func TestUpdateGovernor(t *testing.T) {
	accountAddress_common := common.HexToAddress(accountAddress_string)
	ret0, err := service.UpdateGovernor(accountAddress_common, weight)
	
	if err != nil {
		t.Fatalf("TestUpdateGovernor failed: %v", err)
	}
	if ret0.Int64() < 0 {
		t.Fatalf("TestUpdateGovernor failed, the ret0 \"%v\" is less then \"0\"", ret0.Int64())
	}

	t.Logf("ret0: %v\n", ret0.Int64())
}

func TestSetRate(t *testing.T) {
	ret0, err := service.SetRate(participatesRate, winRate)
	
	if err != nil {
		t.Fatalf("TestSetRate failed: %v", err)
	}
	if ret0.Int64() < 0 {
		t.Fatalf("TestSetRate failed, the ret0 \"%v\" is less then \"0\"", ret0.Int64())
	}

	ret1, err := service.GetCommitteeInfo()
	if ret1.ParticipatesRate != participatesRate {
		t.Fatalf("TestSetRate failed, the participatesRate \"%v\" is inconsistent with ret1.ParticipatesRate \"%v\"", participatesRate, ret1.ParticipatesRate)
	}
	if ret1.WinRate != winRate {
		t.Fatalf("TestSetRate failed, the winRate \"%v\" is inconsistent with ret1.WinRate \"%v\"", winRate, ret1.WinRate)
	}

	t.Logf("ret0: %v\n", ret0.Int64())
}

func TestSetDeployAuthType(t *testing.T) {
	ret0, err := service.GetDeployAuthType()

	// You need to set the deployAuthType to be different from before
	var deployAuthType uint8
	if uint8(ret0.Uint64()) == uint8(2){
		deployAuthType=uint8(1)
	}else{
		deployAuthType=uint8(2)
	}

	ret1, err := service.SetDeployAuthType(deployAuthType)
	
	if err != nil {
		t.Fatalf("TestSetDeployAuthType failed: %v", err)
	}
	if ret1.Int64() < 0 {
		t.Fatalf("TestSetDeployAuthType failed, the ret1 \"%v\" is less then \"0\"", ret1.Int64())
	}

	t.Logf("ret1: %v\n", ret1.Int64())
}

func TestModifyDeployAuth(t *testing.T) {
	address_common := common.HexToAddress(normalAccountAddress_string)
	ret0, err := service.ModifyDeployAuth(address_common, openFlag)
	
	if err != nil {
		t.Fatalf("TestModifyDeployAuth failed: %v", err)
	}
	if ret0.Int64() < 0 {
		t.Fatalf("TestModifyDeployAuth failed, the ret0 \"%v\" is less then \"0\"", ret0.Int64())
	}

	t.Logf("ret0: %v\n", ret0.Int64())
}

func TestResetAdmin(t *testing.T) {
	newAdmin_common := common.HexToAddress(normalAccountAddress_string)
	contractAddress_common := common.HexToAddress(proposalManager_string)
	ret0, err := service.ResetAdmin(newAdmin_common, contractAddress_common)
	
	if err != nil {
		t.Fatalf("TestResetAdmin failed: %v", err)
	}
	if ret0.Int64() < 0 {
		t.Fatalf("TestResetAdmin failed, the ret0 \"%v\" is less then \"0\"", ret0.Int64())
	}

	t.Logf("ret0: %v\n", ret0.Int64())
}

func TestRevokeProposal(t *testing.T) {
	// accountAddress_common := common.HexToAddress(accountAddress_string)
	// service.UpdateGovernor(accountAddress_common, weight)
	// service.SetRate(participatesRate, winRate)

	lastProposalNum, _ := service.ProposalCount()
	lastProposalStatus, _:= service.GetProposalStatus(lastProposalNum)

	_, err := service.RevokeProposal(*lastProposalNum)
	if err != nil {
		t.Fatalf("TestRevokeProposal failed: %v", err)
	}

	lastProposalNewStatus, _:= service.GetProposalStatus(lastProposalNum)

	//0-not exist 1-created 2-passed 3-denied 4-revoked 5-outdated
	if lastProposalStatus==uint8(1) && lastProposalNewStatus!=uint8(4){
		t.Fatalf("TestRevokeProposal failed")
	}

	t.Logf("lastProposalNum: %v\n", lastProposalNum)
	t.Logf("lastProposalStatus: %v\n", lastProposalStatus)
	t.Logf("lastProposalNewStatus: %v\n", lastProposalNewStatus)
}

func TestVoteProposal(t *testing.T) {
	accountAddress_common := common.HexToAddress(accountAddress_string)
	service.UpdateGovernor(accountAddress_common, weight)

	lastProposalNum, _ := service.ProposalCount()
	lastProposalStatus, _:= service.GetProposalStatus(lastProposalNum)

	agree := false
	_, err := service.VoteProposal(*lastProposalNum, agree)
	if err != nil {
		t.Fatalf("TestVoteProposal failed: %v", err)
	}

	proposalInfo, err := service.GetProposalInfo(lastProposalNum)
	againstVoters := proposalInfo.AgainstVoters

	lastProposalNewStatus, _:= service.GetProposalStatus(lastProposalNum)

	t.Logf("lastProposalNum: %v\n", lastProposalNum)
	t.Logf("lastProposalStatus: %v\n", lastProposalStatus)
	t.Logf("againstVoters: %v\n", againstVoters)
	t.Logf("lastProposalNewStatus: %v\n", lastProposalNewStatus)
}
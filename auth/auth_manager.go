package auth

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
)

type ProposalInfo struct {
	ResourceId          common.Address
	Proposer            common.Address
	ProposalType        uint8
	BlockNumberInterval *big.Int
	Status              uint8
	AgreeVoters         []common.Address
	AgainstVoters       []common.Address
}

type CommitteeInfo struct {
	ParticipatesRate uint8
	WinRate          uint8
	Governors        []common.Address
	Weights          []uint32
}

type AuthManagerService struct {
	client          *client.Client
	authManagerAuth *bind.TransactOpts

	contractAuthPrecompiled *ContractAuthPrecompiled
	committeeManager        *CommitteeManager
	committee               *Committee
	proposalManager         *ProposalManager
}

//var committeeAddress = common.HexToAddress("0000000000000000000000000000000000010001")

//var proposalManagerAddress = common.HexToAddress("0000000000000000000000000000000000010001")

var committeeManagerAddress = common.HexToAddress("0000000000000000000000000000000000010001")

var contractAuthPrecompiledAddress = common.HexToAddress("0000000000000000000000000000000000001005")

var DEFAULT_BLOCK_NUMBER_INTERVAL = big.NewInt(3600 * 24 * 7)

func NewAuthManagerService(client *client.Client) (services *AuthManagerService, err error) {
	authManagerAuth := client.GetTransactOpts()

	mgr, err := NewCommitteeManager(committeeManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewCommitteeManager construct Service failed, err: %+v", err)
	}

	pre, err := NewContractAuthPrecompiled(contractAuthPrecompiledAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewContractAuthPrecompiled construct Service failed, err: %+v", err)
	}

	//opts := &bind.CallOpts{From: authManagerAuth.From}
	//fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	//
	//fmt.Printf("services.committeeManager : %s", services.committeeManager)
	//
	//committeeCon, err := services.committeeManager.CommitteeManagerCaller.Committee(opts)
	//
	//fmt.Printf("committeeCon : %s", committeeCon)
	//fmt.Printf("committeeCon err : %w", err)
	//
	//committee, err := NewCommittee(committeeCon, client)
	//if err != nil {
	//	return nil, fmt.Errorf("NewCommittee construct Service failed, err: %+v", err)
	//}
	//
	//proposalMgrCon, err := services.committeeManager.ProposalMgr(opts)
	//proposalManager, err := NewProposalManager(proposalMgrCon, client)
	//if err != nil {
	//	return nil, fmt.Errorf("NewProposalManager construct Service failed, err: %+v", err)
	//}

	s := &AuthManagerService{client: client,
		authManagerAuth:         authManagerAuth,
		committeeManager:        mgr,
		contractAuthPrecompiled: pre,
		//committee:               committee,
		//proposalManager:         proposalManager,
	}

	return s, nil
}

//6.1 无需权限的查询接口
//get Committee info
func (service *AuthManagerService) GetCommitteeInfo() (c *CommitteeInfo, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.committee.CommitteeCaller.GetCommitteeInfo(opts)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetCommitteeInfo failed, err: %v", err)
	}

	if reflect.DeepEqual(result, CommitteeInfo{}) {
		return nil, fmt.Errorf("AuthManagerService GetCommitteeInfo is empty, err: %v", err)
	}

	var info CommitteeInfo

	info.Governors = result.Governors
	info.WinRate = result.WinRate
	info.ParticipatesRate = result.ParticipatesRate
	info.Weights = result.Weights

	return &info, nil
}

//get proposal info
func (service *AuthManagerService) GetProposalInfo(proposalId big.Int) (proposalInfo *ProposalInfo, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.proposalManager.GetProposalInfo(opts, &proposalId)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetProposalInfo failed, err: %v", err)
	}

	if reflect.DeepEqual(result, ProposalInfo{}) {
		return nil, fmt.Errorf("AuthManagerService GetProposalInfo is empty, err: %v", err)
	}

	var info ProposalInfo
	info.ResourceId = result.ResourceId
	info.Proposer = result.Proposer
	info.ProposalType = result.ProposalType
	info.BlockNumberInterval = result.BlockNumberInterval
	info.Status = result.Status
	info.AgreeVoters = result.AgreeVoters
	info.AgainstVoters = result.AgainstVoters

	return &info, nil
}

//get global deploy auth type
func (service *AuthManagerService) GetDeployAuthType() (*big.Int, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuthPrecompiled.DeployType(opts)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetDeployAuthType failed, err: %v", err)
	}

	return result, nil
}

//check the account whether this account can deploy contract
func (service *AuthManagerService) CheckDeployAuth(account common.Address) (*bool, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuthPrecompiled.HasDeployAuth(opts, account)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService CheckDeployAuth failed, err: %v", err)
	}
	return &result, nil
}

//check the contract interface func whether this account can call
func (service *AuthManagerService) CheckMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (*bool, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuthPrecompiled.CheckMethodAuth(opts, contractAddr, funcSelector, account)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService CheckMethodAuth failed, err: %v", err)
	}
	return &result, nil
}

//get a specific contract admin
func (service *AuthManagerService) GetAdmin(contractAddr common.Address) (account *common.Address, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuthPrecompiled.GetAdmin(opts, contractAddr)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetAdmin failed, err: %v", err)
	}
	return &result, nil
}

//6.2 治理委员账号专用接口
//apply for update governor, only governor can call it
func (service *AuthManagerService) UpdateGovernor(account common.Address, weight uint32) (proposalId *big.Int, err error) {

	_, receipt, err := service.committeeManager.
		CreateUpdateGovernorProposal(service.client.GetTransactOpts(), account, weight, DEFAULT_BLOCK_NUMBER_INTERVAL)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService UpdateGovernor failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService UpdateGovernor failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	//service.committeeManager.CreateUpdateGovernorProposal(receipt.GetOutput())

	return nil, nil
}

//apply set participate rate and win rate. only governor can call it
func (service *AuthManagerService) SetRate(participatesRate big.Int, winRate big.Int) (proposalId *big.Int, err error) {
	return nil, nil
}

//submit a proposal of setting deploy contract auth type, only governor can call it
func (service *AuthManagerService) SetDeployAuthType(deployAuthType big.Int) (proposalId *big.Int, err error) {
	return nil, nil
}

//submit a proposal of adding deploy contract auth for account, only governor can call it
func (service *AuthManagerService) ModifyDeployAuth(account common.Address, openFlag bool) (proposalId *big.Int, err error) {
	return nil, nil
}

//submit a proposal of resetting contract admin, only governor can call it
func (service *AuthManagerService) ResetAdmin(newAdmin common.Address, contractAddr common.Address) (proposalId *big.Int, err error) {
	return nil, nil
}

//revoke proposal, only governor can call it
func (service *AuthManagerService) RevokeProposal(proposalId big.Int) (receipt *types.Receipt, err error) {
	return nil, nil
}

//unified vote, only governor can call it
func (service *AuthManagerService) VoteProposal(proposalId big.Int, agree bool) (receipt *types.Receipt, err error) {
	return nil, nil
}

//6.3 合约管理员账号专用接口
//set a specific contract's method auth type, only contract admin can call it
func (service *AuthManagerService) SetMethodAuthType(contractAddr common.Address, funcSelector [4]byte, authType big.Int) (rtCdoe *big.Int, err error) {
	return nil, nil
}

//set a specific contract's method ACL, only contract admin can call it
func (service *AuthManagerService) SetMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address, isOpen bool) (rtCdoe *big.Int, err error) {
	return nil, nil
}

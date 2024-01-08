package auth

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/precompiled"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
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
	client           *client.Client
	authManagerAuth  *bind.TransactOpts
	contractAuth     *ContractAuth
	accountManager   *AccountManager
	committee        *Committee
	committeeManager *CommitteeManager
	proposalManager  *ProposalManager
}

const (
	ACCOUNT_STATUS_NORMAL = iota
	ACCOUNT_STATUS_FREEZE
	ACCOUNT_STATUS_ABOLISH
)

var committeeManagerAddress = common.HexToAddress("0000000000000000000000000000000000010001")

var accountManagerAddress = common.HexToAddress("0000000000000000000000000000000000010003")

var contractAuthAddress = common.HexToAddress("0000000000000000000000000000000000001005")

var DEFAULT_BLOCK_NUMBER_INTERVAL = big.NewInt(3600 * 24 * 7)

func NewAuthManagerService(client *client.Client) (services *AuthManagerService, err error) {
	authManagerAuth := client.GetTransactOpts()

	accountManagerInstance, err := NewAccountManager(accountManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewAccountManager construct Service failed: %+v", err)
	}

	committeeManagerInstance, err := NewCommitteeManager(committeeManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewCommitteeManager construct Service failed, err: %+v", err)
	}

	contractAuthInstance, err := NewContractAuth(contractAuthAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewContractAuth construct Service failed, err: %+v", err)
	}

	opts := &bind.CallOpts{From: authManagerAuth.From}
	committeeCon, err := committeeManagerInstance.Committee(opts)
	committeeInstance, err := NewCommittee(committeeCon, client)
	if err != nil {
		return nil, fmt.Errorf("NewCommittee construct Service failed, err: %+v", err)
	}

	proposalMgrCon, err := committeeManagerInstance.ProposalMgr(opts)
	proposalManagerInstance, err := NewProposalManager(proposalMgrCon, client)
	if err != nil {
		return nil, fmt.Errorf("NewProposalManager construct Service failed, err: %+v", err)
	}

	s := &AuthManagerService{client: client,
		authManagerAuth:  authManagerAuth,
		accountManager:   accountManagerInstance,
		committeeManager: committeeManagerInstance,
		committee:        committeeInstance,
		contractAuth:     contractAuthInstance,
		proposalManager:  proposalManagerInstance,
	}

	return s, nil
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

func (service *AuthManagerService) GetAccountStatus(addr common.Address) (uint8, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	ret0, err := service.accountManager.GetAccountStatus(opts, addr)
	return ret0, err
}

/**
 * *************************************************
 * 治理委员账号专用接口
 * *************************************************
**/

func (service *AuthManagerService) SetAccountStatus(addr common.Address, status uint8) (int64, error) {
	ret0, _, _, err := service.accountManager.SetAccountStatus(service.authManagerAuth, addr, status)
	if err != nil {
		return precompiled.DefaultErrorCode, fmt.Errorf("AccountManagerService SetAccountStatus failed: %v", err)
	}
	return int64(ret0), err
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

func (service *AuthManagerService) GetCommitteeInfo() (c *CommitteeInfo, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.committee.GetCommitteeInfo(opts)

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

func (service *AuthManagerService) GetProposalInfo(proposalId *big.Int) (proposalInfo *ProposalInfo, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.proposalManager.GetProposalInfo(opts, proposalId)

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

func (service *AuthManagerService) ProposalCount() (*big.Int, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.proposalManager.ProposalCount(opts)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService ProposalCount failed, err: %v", err)
	}

	return result, err
}

func (service *AuthManagerService) GetProposalStatus(proposalId *big.Int) (uint8, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.proposalManager.GetProposalStatus(opts, proposalId)

	return result, err
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

// get global deploy auth type
func (service *AuthManagerService) GetDeployAuthType() (*big.Int, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuth.DeployType(opts)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetDeployAuthType failed, err: %v", err)
	}

	return result, nil
}

// check the account whether this account can deploy contract
func (service *AuthManagerService) CheckDeployAuth(account common.Address) (*bool, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuth.HasDeployAuth(opts, account)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService CheckDeployAuth failed, err: %v", err)
	}
	return &result, nil
}

// check the contract interface func whether this account can call
func (service *AuthManagerService) CheckMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (*bool, error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuth.CheckMethodAuth(opts, contractAddr, funcSelector, account)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService CheckMethodAuth failed, err: %v", err)
	}
	return &result, nil
}

// get a specific contract admin
func (service *AuthManagerService) GetAdmin(contractAddr common.Address) (account *common.Address, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.contractAuth.GetAdmin(opts, contractAddr)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetAdmin failed, err: %v", err)
	}
	return &result, nil
}

/**
 * *************************************************
 * 合约管理员账号专用接口
 * *************************************************
**/

// set a specific contract's method auth type, only contract admin can call it
// authType white_list or black_list
func (service *AuthManagerService) SetMethodAuthType(contractAddr common.Address, funcSelector [4]byte, authType uint8) (rtCode *big.Int, err error) {
	_, _, receipt, err := service.contractAuth.SetMethodAuthType(service.client.GetTransactOpts(), contractAddr, funcSelector, authType)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService SetMethodAuthType failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService SetMethodAuthType failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	return parseReturnValue(receipt, "setMethodAuthType")
}

// set a specific contract's method ACL, only contract admin can call it
// isOpen if open, then white_list type is true, black_list is false; if close, then white_list type is false, black_list is true
func (service *AuthManagerService) SetMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address, isOpen bool) (rtCode *big.Int, err error) {

	var receipt *types.Receipt
	if isOpen {
		_, _, receipt, err = service.contractAuth.OpenMethodAuth(service.client.GetTransactOpts(), contractAddr, funcSelector, account)

		if err != nil {
			return nil, fmt.Errorf("AuthManagerService OpenMethodAuth failed, err: %v", err)
		}
	} else {
		_, _, receipt, err = service.contractAuth.CloseMethodAuth(service.client.GetTransactOpts(), contractAddr, funcSelector, account)

		if err != nil {
			return nil, fmt.Errorf("AuthManagerService CloseMethodAuth failed, err: %v", err)
		}
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService SetMethodAuth failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}
	if isOpen {
		return parseReturnValue(receipt, "openMethodAuth")
	} else {
		return parseReturnValue(receipt, "closeMethodAuth")
	}

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

// apply for update governor, only governor can call it
// account new governor address
// weight 0 == delete, bigger than 0 == update or insert
func (service *AuthManagerService) UpdateGovernor(account common.Address, weight uint32) (proposalId *big.Int, err error) {
	_, _, receipt, err := service.committeeManager.
		CreateUpdateGovernorProposal(service.client.GetTransactOpts(), account, weight, DEFAULT_BLOCK_NUMBER_INTERVAL)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService UpdateGovernor failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService UpdateGovernor failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	return parseReturnValue(receipt, "createUpdateGovernorProposal")
}

// apply set participate rate and win rate. only governor can call it
// participatesRate [0,100]. if 0, always succeed.
// winRate [0,100].
func (service *AuthManagerService) SetRate(participatesRate uint8, winRate uint8) (proposalId *big.Int, err error) {
	_, _, receipt, err := service.committeeManager.
		CreateSetRateProposal(service.client.GetTransactOpts(), participatesRate, winRate, DEFAULT_BLOCK_NUMBER_INTERVAL)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService SetRate failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService SetRate failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	return parseReturnValue(receipt, "createSetRateProposal")
}

// submit a proposal of setting deploy contract auth type, only governor can call it
// deployAuthType 1-whitelist; 2-blacklist
func (service *AuthManagerService) SetDeployAuthType(deployAuthType uint8) (proposalId *big.Int, err error) {
	_, _, receipt, err := service.committeeManager.
		CreateSetDeployAuthTypeProposal(service.client.GetTransactOpts(), deployAuthType, DEFAULT_BLOCK_NUMBER_INTERVAL)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService SetDeployAuthType failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService SetDeployAuthType failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	return parseReturnValue(receipt, "createSetDeployAuthTypeProposal")
}

// submit a proposal of adding deploy contract auth for account, only governor can call it
// openFlag true-open; false-close
func (service *AuthManagerService) ModifyDeployAuth(account common.Address, openFlag bool) (proposalId *big.Int, err error) {
	_, _, receipt, err := service.committeeManager.
		CreateModifyDeployAuthProposal(service.client.GetTransactOpts(), account, openFlag, DEFAULT_BLOCK_NUMBER_INTERVAL)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService ModifyDeployAuth failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService ModifyDeployAuth failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	return parseReturnValue(receipt, "createModifyDeployAuthProposal")
}

// submit a proposal of resetting contract admin, only governor can call it
func (service *AuthManagerService) ResetAdmin(newAdmin common.Address, contractAddr common.Address) (proposalId *big.Int, err error) {
	_, _, receipt, err := service.committeeManager.
		CreateResetAdminProposal(service.client.GetTransactOpts(), newAdmin, contractAddr, DEFAULT_BLOCK_NUMBER_INTERVAL)
	if err != nil {
		return nil, fmt.Errorf("AuthManagerService ResetAdmin failed, err: %v", err)
	}

	if receipt.Status != 0 {
		return nil, fmt.Errorf("AuthManagerService ResetAdmin failed, ErrorMessage: %v", receipt.GetErrorMessage())
	}

	return parseReturnValue(receipt, "createResetAdminProposal")
}

// revoke proposal, only governor can call it
func (service *AuthManagerService) RevokeProposal(proposalId big.Int) (receipt *types.Receipt, err error) {

	_, receipt, err = service.committeeManager.RevokeProposal(service.client.GetTransactOpts(), &proposalId)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService RevokeProposal failed, err: %v", err)
	}

	return receipt, nil
}

// unified vote, only governor can call it
func (service *AuthManagerService) VoteProposal(proposalId big.Int, agree bool) (receipt *types.Receipt, err error) {

	_, receipt, err = service.committeeManager.VoteProposal(service.client.GetTransactOpts(), &proposalId, agree)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService VoteProposal failed, err: %v", err)
	}

	return receipt, nil
}

func parseReturnValue(receipt *types.Receipt, name string) (*big.Int, error) {
	// todo
	// fmt.Println(receipt)

	errorMessage := receipt.GetErrorMessage()
	if errorMessage != "" {
		return nil, fmt.Errorf("receipt.Status err: %v", errorMessage)
	}
	bigNum, err := precompiled.ParseBigIntFromOutput(receipt)
	if err != nil {
		return nil, fmt.Errorf("parseReturnValue failed, err: %v", err)
	}
	return bigNum, nil
}

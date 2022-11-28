package auth

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

type AuthManagerService struct {
	client          *client.Client
	authManagerAuth *bind.TransactOpts

	contractAuthPrecompiled *ContractAuthPrecompiled
	committeeManager        *CommitteeManager
	committee               *Committee
	proposalManager         *ProposalManager
}

var committeeAddress = common.HexToAddress("0000000000000000000000000000000000010001")

var proposalManagerAddress = common.HexToAddress("0000000000000000000000000000000000010001")

var committeeManagerAddress = common.HexToAddress("0000000000000000000000000000000000010001")

var contractAuthPrecompiledAddress = common.HexToAddress("0000000000000000000000000000000000001005")

func NewAuthManagerService(client *client.Client) (services *AuthManagerService, err error) {
	mgr, err := NewCommitteeManager(committeeManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewCommitteeManager construct Service failed, err: %+v", err)
	}

	pre, err := NewContractAuthPrecompiled(contractAuthPrecompiledAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewContractAuthPrecompiled construct Service failed, err: %+v", err)
	}

	committee, err := NewCommittee(committeeAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewCommittee construct Service failed, err: %+v", err)
	}

	proposalManager, err := NewProposalManager(proposalManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("NewProposalManager construct Service failed, err: %+v", err)
	}

	authManagerAuth := client.GetTransactOpts()

	s := &AuthManagerService{client: client,
		authManagerAuth:         authManagerAuth,
		committeeManager:        mgr,
		contractAuthPrecompiled: pre,
		committee:               committee,
		proposalManager:         proposalManager,
	}

	return s, nil
}

//6.1 无需权限的查询接口
//get Committee info
func (service *AuthManagerService) GetCommitteeInfo() {

	//service.committee.

}

//get proposal info
func (service *AuthManagerService) GetProposalInfo(proposalId big.Int) (info *ProposalInfo, err error) {
	opts := &bind.CallOpts{From: service.authManagerAuth.From}
	result, err := service.proposalManager.GetProposalInfo(opts, &proposalId)

	if err != nil {
		return nil, fmt.Errorf("AuthManagerService GetProposalInfo failed, err: %v", err)
	}

	info.ResourceId = result.ResourceId
	info.Proposer = result.Proposer
	info.ProposalType = result.ProposalType
	info.BlockNumberInterval = result.BlockNumberInterval
	info.Status = result.Status
	info.AgreeVoters = result.AgreeVoters
	info.AgainstVoters = result.AgainstVoters

	return info, nil
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

	return nil, nil
}

//check the contract interface func whether this account can call
func (service *AuthManagerService) CheckMethodAuth(contractAddr common.Address, funcSelector string, account common.Address) (*bool, error) {
	return nil, nil
}

//get a specific contract admin
func (service *AuthManagerService) GetAdmin(contractAddr common.Address) (account *common.Address, err error) {
	return nil, nil
}

//6.2 治理委员账号专用接口
//apply for update governor, only governor can call it
func (service *AuthManagerService) UpdateGovernor(account common.Address, weight big.Int) (proposalId *big.Int, err error) {
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
func (service *AuthManagerService) SetMethodAuthType(contractAddr common.Address, funcSelector string, authType big.Int) (rtCdoe *big.Int, err error) {
	return nil, nil
}

//set a specific contract's method ACL, only contract admin can call it
func (service *AuthManagerService) SetMethodAuth(contractAddr common.Address, funcSelector string, account common.Address, isOpen bool) (rtCdoe *big.Int, err error) {
	return nil, nil
}

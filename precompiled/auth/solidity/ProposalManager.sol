// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;
pragma experimental ABIEncoderV2;

import "./BasicAuth.sol";
import "./VoteComputer.sol";

contract ProposalManager is BasicAuth {
    struct ProposalInfo {
        // unique address
        address resourceId;
        // creator external address
        address proposer;
        // proposal type: 1X-committee；2X-deploy contract auth；3X-admin auth
        uint8 proposalType;
        // block number interval
        uint256 blockNumberInterval;
        //0-not exist 1-created 2-passed 3-denied 4-revoked 5-outdated
        uint8 status;
        // approve voters list
        address[] agreeVoters;
        // against voters List
        address[] againstVoters;
    }
    // Committee handler
    VoteComputerTemplate public _voteComputer;
    // auto generated proposal id
    uint256 public _proposalCount;
    // (id, proposal)
    mapping(uint256 => ProposalInfo) public _proposals;
    // (type, (resource id, proposal id))
    mapping(uint8 => mapping(address => uint256)) public _proposalIndex;

    modifier proposalExist(uint256 proposalId) {
        require(_proposals[proposalId].status != 0, "Proposal not exist");
        _;
    }

    modifier proposalVotable(uint256 proposalId) {
        require(_proposals[proposalId].status == 1, "Proposal is not votable");
        _;
    }

    constructor(address committeeMgrAddress, address committeeAddress) public {
        _voteComputer = new VoteComputer(committeeMgrAddress, committeeAddress);
    }

    function setVoteComputer(address addr) public onlyOwner {
        _voteComputer = VoteComputerTemplate(addr);
    }

    /*
     * predicate proposal outdated
     * @param proposal id
     */
    function refreshProposalStatus(uint256 proposalId)
        public
        proposalExist(proposalId)
        returns (uint8)
    {
        ProposalInfo storage proposal = _proposals[proposalId];
        if (proposal.status == 1) {
            if (block.number > proposal.blockNumberInterval) {
                proposal.status = 5;
                return 5;
            }
        }
        return proposal.status;
    }

    /*
     * create proposal
     * @param create address
     * @param  proposal type : 1X-committee；2X-deploy contract auth；3X-admin auth
     * @param resource id
     * @param  after the block number interval, the proposal would be outdated.
     */
    function create(
        address proposer,
        uint8 proposalType,
        address resourceId,
        uint256 blockNumberInterval
    ) public onlyOwner returns (uint256) {
        uint256 alreadExistProposalId = _proposalIndex[proposalType][
            resourceId
        ];
        if (_proposals[alreadExistProposalId].status == 1) {
            refreshProposalStatus(alreadExistProposalId);
        }
        require(
            _proposals[alreadExistProposalId].status != 1,
            "Current proposal not end"
        );
        _proposalCount++;
        uint256 proposalId = _proposalCount;
        address[] memory agreeVoters;
        address[] memory againstVoters;
        ProposalInfo memory proposal = ProposalInfo(
            resourceId,
            proposer,
            proposalType,
            block.number + blockNumberInterval,
            1,
            agreeVoters,
            againstVoters
        );

        _proposals[proposalId] = proposal;
        _proposalIndex[proposalType][resourceId] = proposalId;
        return proposalId;
    }

    /*
     * unified vote
     * @param proposal id
     * @param true or false
     * @param voter address
     */
    function vote(
        uint256 proposalId,
        bool agree,
        address voterAddress
    )
        public
        onlyOwner
        proposalExist(proposalId)
        proposalVotable(proposalId)
        returns (uint8)
    {
        ProposalInfo storage proposal = _proposals[proposalId];
        require(!hasVoted(proposal, voterAddress), "Already voted");
        if (agree) {
            proposal.agreeVoters.push(voterAddress);
        } else {
            proposal.againstVoters.push(voterAddress);
        }
        uint8 status = _voteComputer.determineVoteResult(
            proposal.agreeVoters,
            proposal.againstVoters
        );
        proposal.status = status;
        return status;
    }

    /*
     * revoke proposal
     * @param proposal id
     * @param voter address
     */
    function revoke(uint256 proposalId, address voterAddress) public onlyOwner {
        ProposalInfo storage proposal = _proposals[proposalId];
        require(
            refreshProposalStatus(proposalId) == 1,
            "Only newly created proposal can be revoked"
        );
        require(proposal.proposer == voterAddress, "Only proposer can revoke");
        proposal.status = 4;
    }

    /*
     * get proposal info
     * @param proposal id
     */
    function getProposalInfo(uint256 proposalId)
        public
        view
        returns (
            address resourceId,
            address proposer,
            uint8 proposalType,
            uint256 blockNumberInterval,
            uint8 status,
            address[] memory agreeVoters,
            address[] memory againstVoters
        )
    {
        ProposalInfo storage proposal = _proposals[proposalId];
        resourceId = proposal.resourceId;
        proposer = proposal.proposer;
        proposalType = proposal.proposalType;
        blockNumberInterval = proposal.blockNumberInterval;
        status = proposal.status;
        agreeVoters = proposal.agreeVoters;
        againstVoters = proposal.againstVoters;
    }

    /*
     * get proposalInfo list, range in [from, to]
     */
    function getProposalInfoList(uint256 from, uint256 to)
        public
        view
        returns (ProposalInfo[] memory)
    {
        require(
            from <= _proposalCount,
            "'from' is greater than 'proposalCount'"
        );
        require(from <= to, "'from' is greater than 'to'");
        if (to > _proposalCount) {
            to = _proposalCount;
        }
        if (from < 1) {
            from = 1;
        }
        ProposalInfo[] memory _infoList = new ProposalInfo[](to - from + 1);
        uint256 _infoListIndex = 0;
        for (uint256 i = from; i <= to; i++) {
            ProposalInfo storage proposal = _proposals[i];
            _infoList[_infoListIndex++] = proposal;
        }
        return _infoList;
    }

    /*
     * get proposal status by proposal id
     * @param proposal id
     */
    function getProposalStatus(uint256 proposalId) public view returns (uint8) {
        return _proposals[proposalId].status;
    }

    /*
     * get info by proposal type and resource id
     * @param create address
     * @param  proposal type : 1X-committee；2X-deploy contract auth；3X-admin auth
     * @param resource id
     * @param  after the block number interval, the proposal would be outdated.
     */
    function getIdByTypeAndResourceId(uint8 proposalType, address resourceId)
        public
        view
        returns (uint256)
    {
        return _proposalIndex[proposalType][resourceId];
    }

    /**
     * judge account if voted for the proposal
     */
    function hasVoted(ProposalInfo memory proposal, address account)
        internal
        pure
        returns (bool)
    {
        if (
            contains(proposal.agreeVoters, account) ||
            contains(proposal.againstVoters, account)
        ) {
            return true;
        }
        return false;
    }

    function contains(address[] memory array, address value)
        internal
        pure
        returns (bool)
    {
        for (uint256 i = 0; i < array.length; i++) {
            if (value == array[i]) {
                return true;
            }
        }
        return false;
    }
}
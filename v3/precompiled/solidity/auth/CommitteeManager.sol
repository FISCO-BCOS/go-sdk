// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;
// pragma experimental ABIEncoderV2;

import "./Committee.sol";
import "./ProposalManager.sol";
import "./ContractAuthPrecompiled.sol";
import "./ConsensusPrecompiled.sol";
import "./SystemConfigPrecompiled.sol";
import "./VoteComputerTemplate.sol";

contract CommitteeManager {
    // Governors and threshold
    Committee public _committee;
    // proposal manager
    ProposalManager public _proposalMgr;
    SystemConfigPrecompiled constant _systemConfigPrecompiled =
        SystemConfigPrecompiled(address(0x1000));
    ConsensusPrecompiled constant _consensusPrecompiled =
        ConsensusPrecompiled(address(0x1003));
    ContractAuthPrecompiled constant _contractPrecompiled =
        ContractAuthPrecompiled(address(0x1005));

    // exec proposal when vote pass through
    // 0 == success, others exec error
    event execResult(int256);

    struct ProposalInfo {
        // Committee management: 11-set governor weight; 12-set rate; 13-upgrade VoteComputer contract;
        // access control: 21-set deploy auth type; 22-modify deploy auth;
        // contract admin: 31-reset admin
        // system config management: 41- set config
        // consensus node management: 51- set weight (weigh > 0, sealer; weight = 0, observer), 52- remove
        uint8 proposalType;
        // unique address
        address resourceId;
        // uint8 array
        uint8[] uint8Array;
        // string array
        string[] strArray;
        uint32 weight;
        // address array
        address[] addressArray;
        bool flag;
    }
    // [id, Proposal]
    mapping(uint256 => ProposalInfo) private _proposalInfoMap;

    modifier onlyGovernor() {
        require(isGovernor(msg.sender), "you must be governor");
        _;
    }

    constructor(
        address[] memory initGovernors,
        uint32[] memory weights,
        uint8 participatesRate,
        uint8 winRate
    ) public {
        //_contractPrecompiled = ContractAuthPrecompiled(0x1005);
        _committee = new Committee(
            initGovernors,
            weights,
            participatesRate,
            winRate
        );
        _proposalMgr = new ProposalManager(address(this), address(_committee));
    }

    /*
     * apply for update governor
     * @param external account
     * @param weight, 0-delete, >0-update or insert
     * @param blockNumberInterval, after current block number, it will be outdated.
     */
    function createUpdateGovernorProposal(
        address account,
        uint32 weight,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        address[] memory addressArray = new address[](1);
        uint8[] memory uint8Array;
        string[] memory strArray;
        addressArray[0] = account;
        ProposalInfo memory proposalInfo = ProposalInfo(
            11,
            account,
            uint8Array,
            strArray,
            weight,
            addressArray,
            true
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * apply set participate rate and win rate.
     * @param participate rate, [0,100]. if 0, always succeed.
     * @param win rate, [0,100].
     * @param blockNumberInterval, after current block number, it will be outdated.
     */
    function createSetRateProposal(
        uint8 participatesRate,
        uint8 winRate,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        require(
            participatesRate >= 0 && participatesRate <= 100,
            "invalid range of participatesRate"
        );
        require(winRate >= 0 && winRate <= 100, "invalid range of winRate");
        address[] memory addressArray;
        uint8[] memory uint8Array = new uint8[](2);
        string[] memory strArray;
        uint8Array[0] = participatesRate;
        uint8Array[1] = winRate;
        ProposalInfo memory proposalInfo = ProposalInfo(
            12,
            address(this),
            uint8Array,
            strArray,
            0,
            addressArray,
            false
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit an propsal of upgrade VoteCompter.sol
     * @param new address of VoteComputer
     * @param contractAddr the address of contract which will propose to reset admin
     */
    function createUpgradeVoteComputerProposal(
        address newAddr,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        address[] memory allGovernors;
        address[] memory emptyAddress;
        (, , allGovernors, ) = _committee.getCommitteeInfo();
        VoteComputerTemplate newVoteComputer = VoteComputerTemplate(newAddr);
        require(
            newVoteComputer._committee() == _committee,
            "new vote computer committee address mismatch"
        );
        require(
            newVoteComputer.determineVoteResult(allGovernors, emptyAddress) ==
                2,
            "new vote computer imperfection"
        );
        address[] memory addressArray = new address[](1);
        uint8[] memory uint8Array;
        string[] memory strArray;
        addressArray[0] = newAddr;
        ProposalInfo memory proposalInfo = ProposalInfo(
            13,
            address(_proposalMgr),
            uint8Array,
            strArray,
            0,
            addressArray,
            false
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit an proposal of setting deploy contract auth type
     * @param deployAuthType: 1- whitelist; 2-blacklist
     */
    function createSetDeployAuthTypeProposal(
        uint8 deployAuthType,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        require(
            _contractPrecompiled.deployType() != deployAuthType,
            "the current deploy auth type is the same as you want to set"
        );

        address[] memory addressArray;
        uint8[] memory uint8Array = new uint8[](1);
        string[] memory strArray;
        uint8Array[0] = deployAuthType;
        ProposalInfo memory proposalInfo = ProposalInfo(
            21,
            address(_contractPrecompiled),
            uint8Array,
            strArray,
            0,
            addressArray,
            false
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit an proposal of adding deploy contract auth for account
     * @param account
     * @param openFlag: true-open; false-close
     */
    function createModifyDeployAuthProposal(
        address account,
        bool openFlag,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        address[] memory addressArray = new address[](1);
        addressArray[0] = account;
        uint8[] memory uint8Array;
        string[] memory strArray;
        ProposalInfo memory proposalInfo = ProposalInfo(
            22,
            account,
            uint8Array,
            strArray,
            0,
            addressArray,
            openFlag
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit an propsal of resetting contract admin
     * @param newAdmin
     * @param contractAddr the address of contract which will propose to reset admin
     */
    function createResetAdminProposal(
        address newAdmin,
        address contractAddr,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        require(contractAddr != address(0), "contract address not exists.");
        // require(methodAuthMgr._owner() == address(this), "caller is not owner");
        require(
            newAdmin != _contractPrecompiled.getAdmin(contractAddr),
            "the account has been the admin of concurrt contract."
        );
        address[] memory addressArray = new address[](1);
        uint8[] memory uint8Array;
        string[] memory strArray;
        addressArray[0] = newAdmin;
        ProposalInfo memory proposalInfo = ProposalInfo(
            31,
            contractAddr,
            uint8Array,
            strArray,
            0,
            addressArray,
            false
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit a propsal of set system config
     * @param key （tx_count_limit,consensus_leader_period,tx_gas_limit)
     * @param value
     */
    function createSetSysConfigProposal(
        string memory key,
        string memory value,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        require(bytes(key).length > 1, "invalid key length.");
        address[] memory addressArray;
        uint8[] memory uint8Array;
        string[] memory strArray = new string[](2);
        strArray[0] = key;
        strArray[1] = value;
        ProposalInfo memory proposalInfo = ProposalInfo(
            41,
            address(_systemConfigPrecompiled),
            uint8Array,
            strArray,
            0,
            addressArray,
            false
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit a proposal of set consensus weight
     * @param node
     * @param weight: weigh > 0, sealer; weight = 0, observer
     * @param addFlag true-> add, false-> set
     * @param blockNumberInterval, after current block number, it will be outdated.
     */
    function createSetConsensusWeightProposal(
        string memory node,
        uint32 weight,
        bool addFlag,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        require(bytes(node).length > 1, "invalid node.");
        address[] memory addressArray;
        uint8[] memory uint8Array;
        string[] memory strArray = new string[](1);
        strArray[0] = node;
        ProposalInfo memory proposalInfo = ProposalInfo(
            51,
            address(_consensusPrecompiled),
            uint8Array,
            strArray,
            weight,
            addressArray,
            addFlag
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * submit a proposal of remove node
     * @param node
     * @param blockNumberInterval, after current block number, it will be outdated.
     */
    function createRmNodeProposal(
        string memory node,
        uint256 blockNumberInterval
    ) public onlyGovernor returns (uint256 currentproposalId) {
        require(bytes(node).length > 1, "invalid node.");
        address[] memory addressArray;
        uint8[] memory uint8Array;
        string[] memory strArray = new string[](1);
        strArray[0] = node;
        ProposalInfo memory proposalInfo = ProposalInfo(
            52,
            address(_consensusPrecompiled),
            uint8Array,
            strArray,
            0,
            addressArray,
            true
        );
        currentproposalId = _createProposal(proposalInfo, blockNumberInterval);
    }

    /*
     * create proposal
     * @param create address
     * @param  proposal type : 1X-committee；2X-deploy contract auth；3X-admin auth
     * @param resource id
     * @param  after the block number interval, the proposal would be outdated.
     */
    function _createProposal(
        ProposalInfo memory proposalInfo,
        uint256 blockNumberInterval
    ) internal returns (uint256) {
        uint256 proposalId = _proposalMgr.create(
            msg.sender,
            proposalInfo.proposalType,
            proposalInfo.resourceId,
            blockNumberInterval
        );
        _proposalInfoMap[proposalId] = proposalInfo;
        //detault vote agree for the proposal.
        voteProposal(proposalId, true);
        return proposalId;
    }

    /*
     * revoke proposal
     * @param proposal id
     */
    function revokeProposal(uint256 proposalId) public onlyGovernor {
        _proposalMgr.revoke(proposalId, msg.sender);
    }

    /*
     * unified vote
     * @param proposal id
     * @param true or false
     * @event error code, 
        0 == exec success, others exec error
     */
    function voteProposal(uint256 proposalId, bool agree) public onlyGovernor {
        uint8 voteStatus = _proposalMgr.vote(proposalId, agree, msg.sender);
        ProposalInfo memory proposalInfo;
        if (voteStatus == 2) {
            int256 retCode = 0;
            uint8 proposalType = getProposalType(proposalId);
            proposalInfo = _proposalInfoMap[proposalId];
            if (proposalType == 11) {
                if (proposalInfo.weight == 0) {
                    require(
                        proposalInfo.addressArray[0] != msg.sender,
                        "You can not remove yourself!"
                    );
                }
                _committee.setWeight(
                    proposalInfo.addressArray[0],
                    proposalInfo.weight
                );
            } else if (proposalType == 12) {
                _committee.setRate(
                    proposalInfo.uint8Array[0],
                    proposalInfo.uint8Array[1]
                );
            } else if (proposalType == 13) {
                _proposalMgr.setVoteComputer(proposalInfo.addressArray[0]);
            } else if (proposalType == 21) {
                retCode = _contractPrecompiled.setDeployAuthType(
                    proposalInfo.uint8Array[0]
                );
            } else if (proposalType == 22) {
                if (proposalInfo.flag) {
                    retCode = _contractPrecompiled.openDeployAuth(
                        proposalInfo.addressArray[0]
                    );
                } else {
                    retCode = _contractPrecompiled.closeDeployAuth(
                        proposalInfo.addressArray[0]
                    );
                }
            } else if (proposalType == 31) {
                // (contractAddress, adminAddress)
                retCode = _contractPrecompiled.resetAdmin(
                    proposalInfo.resourceId,
                    proposalInfo.addressArray[0]
                );
            } else if (proposalType == 41) {
                retCode = _systemConfigPrecompiled.setValueByKey(
                    proposalInfo.strArray[0],
                    proposalInfo.strArray[1]
                );
            } else if (proposalType == 51) {
                if (proposalInfo.flag) {
                    if (proposalInfo.weight == 0) {
                        retCode = _consensusPrecompiled.addObserver(
                            proposalInfo.strArray[0]
                        );
                    } else {
                        retCode = _consensusPrecompiled.addSealer(
                            proposalInfo.strArray[0],
                            uint256(proposalInfo.weight)
                        );
                    }
                } else {
                    retCode = _consensusPrecompiled.setWeight(
                        proposalInfo.strArray[0],
                        uint256(proposalInfo.weight)
                    );
                }
            } else if (proposalType == 52) {
                retCode = _consensusPrecompiled.remove(
                    proposalInfo.strArray[0]
                );
            } else {
                revert("vote type error.");
            }
            emit execResult(retCode);
        }
    }

    /*
     * predicate governor
     * @param external account
     */
    function isGovernor(address account) public view returns (bool) {
        return _committee.isGovernor(account);
    }

    /*
     * get proposal type
     * @param proposal id
     */
    function getProposalType(uint256 proposalId) public view returns (uint8) {
        return _proposalInfoMap[proposalId].proposalType;
    }
}
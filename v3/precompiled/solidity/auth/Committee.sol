// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

import "./LibAddressSet.sol";
import "./BasicAuth.sol";

contract Committee is BasicAuth {
    using LibAddressSet for LibAddressSet.AddressSet;

    LibAddressSet.AddressSet private _governorSet;
    mapping(address => uint32) private _weightMapping;
    uint8 public _participatesRate;
    uint8 public _winRate;

    constructor(
        address[] memory governorList,
        uint32[] memory weightList,
        uint8 participatesRate,
        uint8 winRate
    ) public {
        for (uint32 i = 0; i < governorList.length; i++) {
            setWeight(governorList[i], weightList[i]);
        }
        _winRate = winRate;
        _participatesRate = participatesRate;
        setOwner(msg.sender);
    }

    /*
     * set rate by owner
     * @param participate rate
     * @param win rate
     */
    function setRate(uint8 participatesRate, uint8 winRate) public onlyOwner {
        _winRate = winRate;
        _participatesRate = participatesRate;
    }

    /*
     * set weight only by owner
     * @param governor external address
     * @param weight
     */
    function setWeight(address governor, uint32 weight) public onlyOwner {
        if (weight == 0) {
            require(governor != tx.origin, "You can not remove yourself!");
            delete _weightMapping[governor];
            _governorSet.remove(governor);
        } else if (_governorSet.contains(governor)) {
            _weightMapping[governor] = weight;
        } else {
            _weightMapping[governor] = weight;
            _governorSet.add(governor);
        }
    }

    /*
     * get committee info
     */
    function getCommitteeInfo()
        public
        view
        returns (
            uint8 participatesRate,
            uint8 winRate,
            address[] memory governors,
            uint32[] memory weights
        )
    {
        governors = _governorSet.getAll();
        weights = new uint32[](governors.length);
        for (uint256 i = 0; i < governors.length; i++) {
            weights[i] = _weightMapping[governors[i]];
        }
        winRate = _winRate;
        participatesRate = _participatesRate;
    }

    /*
     * predicate governor
     * @param governor address
     */
    function isGovernor(address governor) public view returns (bool) {
        return _governorSet.contains(governor);
    }

    /*
     * get weight
     * @param governor address
     */
    function getWeight(address governor) public view returns (uint32) {
        return _weightMapping[governor];
    }

    /*
     * compute weights with given votes list
+     * @param computed voters list
     */
    function getWeights(address[] memory votes) public view returns (uint32) {
        uint32 totalVotes = 0;
        for (uint32 i = 0; i < votes.length; i++) {
            totalVotes += _weightMapping[votes[i]];
        }
        return totalVotes;
    }

    /*
     * compute weights with governor set
     */
    function getWeights() public view returns (uint32) {
        return getWeights(_governorSet.getAll());
    }
}
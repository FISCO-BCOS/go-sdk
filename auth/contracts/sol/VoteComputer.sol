// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

import "./Committee.sol";
import "./VoteComputerTemplate.sol";

contract VoteComputer is VoteComputerTemplate {
    // Governors and threshold
    constructor(address committeeMgrAddress, address committeeAddress)
        public
        VoteComputerTemplate(committeeMgrAddress, committeeAddress)
    {}

    /*
     * predicate vote result and return the status
     * @param for voters list
     * @param against voters list
     */
    function determineVoteResult(
        address[] memory agreeVoters,
        address[] memory againstVoters
    ) public view override returns (uint8) {
        uint32 agreeVotes = _committee.getWeights(agreeVoters);
        uint32 doneVotes = agreeVotes + _committee.getWeights(againstVoters);
        uint32 allVotes = _committee.getWeights();
        return
            voteResultCalc(
                agreeVotes,
                doneVotes,
                allVotes,
                _committee._participatesRate(),
                _committee._winRate()
            );
    }

    /*
     * calculate vote result and return the status,
     * for convenience, this method for committee check calculate logic
     * @param agree voter total weight
     * @param voted voter total weight
     * @param all voter total weight
     * @param participate threshold, percentage
     * @param win threshold, percentage
     */
    function voteResultCalc(
        uint32 agreeVotes,
        uint32 doneVotes,
        uint32 allVotes,
        uint8 participatesRate,
        uint8 winRate
    ) public pure override returns (uint8) {
        //1. Checks enough voters: totalVotes/totalVotesPower >= p_rate/100
        if (doneVotes * 100 < allVotes * participatesRate) {
            //not enough voters, need more votes
            return 1;
        }
        //2. Checks whether for votes wins: agreeVotes/totalVotes >= win_rate/100
        if (agreeVotes * 100 >= winRate * doneVotes) {
            return 2;
        } else {
            return 3;
        }
    }
}

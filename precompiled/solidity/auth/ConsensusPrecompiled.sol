// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

contract ConsensusPrecompiled {
    function addSealer(string memory, uint256) public returns (int256) {}

    function addObserver(string memory) public returns (int256) {}

    function remove(string memory) public returns (int256) {}

    function setWeight(string memory, uint256) public returns (int256) {}
}
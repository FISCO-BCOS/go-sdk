// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

contract ShardingPrecompiled{
    function getContractShard(string memory absolutePath) public view returns (int256,string memory) {}
    function makeShard(string memory shardName) public returns (int256) {}
    function linkShard(string memory shardName, string memory _address) public returns (int256) {}
}
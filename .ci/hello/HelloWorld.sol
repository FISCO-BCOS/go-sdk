// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

contract HelloWorld {
    string value;
    event setValue(string v, address indexed from, address indexed to, uint256 value);
    string public version = "1";

    constructor() {
        value = "Hello, World!";
    }

    function get() public view returns (string memory) {
        return value;
    }

    function set(string calldata v) public returns (string memory) {
        string memory old = value;
        value = v;
        emit setValue(v, tx.origin, msg.sender, 1);
        return old;
    }
}

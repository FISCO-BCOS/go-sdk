// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

contract HelloWorld {
    string value;
    event setValue(string v, address indexed from, address indexed to, int256 value);
    int public version;

    constructor(string memory initValue) {
        value = initValue;
        version = 0;
    }

    function get() public view returns (string memory) {
        return value;
    }

    function set(string calldata v) public returns (string memory) {
        string memory old = value;
        value = v;
        version = version + 1;
        emit setValue(v, tx.origin, msg.sender, version);
        return old;
    }
}

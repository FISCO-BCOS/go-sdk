pragma solidity>=0.4.24 <0.6.11;

contract HelloWorld {
    string value;
    event setValue(string v, address indexed from, address indexed to, uint256 value);
    string public version = "1";

    constructor() public {
        value = "Hello, World!";
    }

    function get() public view returns (string memory) {
        return value;
    }

    function set(string calldata v) public {
        value = v;
        emit setValue(v, tx.origin, msg.sender, 1);
    }
}
pragma solidity>=0.4.24 <0.6.11;

contract HelloWorld {
    string value;
    event setValue(string);
    string public version = "1";

    constructor() public {
        value = "Hello, World!";
    }

    function get() public view returns (string memory) {
        return value;
    }

    function set(string v) public {
        value = v;
        emit setValue(v);
    }
}
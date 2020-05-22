pragma solidity ^0.4.24;


contract HelloWorld {
    string value;
    event setValue(string);
    string public version = "1";

    constructor() public {
        value = "Hello, World!";
    }

    function get() public view returns (string) {
        return value;
    }

    function set(string v) public {
        value = v;
        emit setValue(v);
    }
}

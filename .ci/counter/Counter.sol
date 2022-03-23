pragma solidity>=0.4.24 <0.6.11;

contract Counter {
    int256 public version;
    int256 count;

    constructor() public {
        version = 0;
        count = 0;
    }

    function get() public view returns (int256) {
        return count;
    }

    function add() public returns (int256) {
        count++;
        return count;
    }

    function set(int256 n) public {
        count = n;
    }
}

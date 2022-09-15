pragma solidity ^0.6.0;

// A parallel contract example
contract ParallelOk
{
    mapping (string => uint256) _balance;

     // Just an example, overflow is ok, use 'SafeMath' if needed
    function transfer(string memory from, string memory to, uint256 num) public
    {
        _balance[from] -= num;
        _balance[to] += num;
    }

    // Just for testing whether the parallel revert function is working well, no practical use
    function transferWithRevert(string memory from, string memory to, uint256 num) public
    {
        _balance[from] -= num;
        _balance[to] += num;
        require(num <= 100);
    }

    function set(string memory name, uint256 num) public
    {
        _balance[name] = num;
    }

    function balanceOf(string memory name) public view returns (uint256)
    {
        return _balance[name];
    }
}
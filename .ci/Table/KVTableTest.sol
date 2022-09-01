pragma solidity>=0.4.24 <0.6.11;

import "./Table.sol";

contract KVTableTest {
    event InsertResult(int256 count);

    // 创建TableManager对象，其在区块链上的固定地址是0x1002
    TableManager constant tm =  TableManager(address(0x1002));
    Table table;
    string constant TABLE_NAME = "t_test";
    constructor () public{
        // 创建t_test表，表的主键名为id，其他字段名为name和age
        string[] memory columnNames = new string[](2);
        columnNames[0] = "name";
        columnNames[1] = "age";
        TableInfo memory tf = TableInfo("id", columnNames);
        tm.createTable(TABLE_NAME, tf);

        // 获取真实的地址，存在合约中
        address t_address = tm.openTable(TABLE_NAME);
        require(t_address!=address(0x0),"");
        table = Table(t_address);
    }

    function insert(string memory id,string memory name,string memory age) public returns (int32){
        string[] memory columns = new string[](2);
        columns[0] = name;
        columns[1] = age;
        Entry memory entry = Entry(id, columns);
        int32 result = table.insert(entry);
        emit InsertResult(result);
        return result;
    }

    function select(string memory id) public view returns (string memory,string memory)
    {
        Entry memory entry = table.select(id);

        string memory name;
        string memory age;
        if(entry.fields.length==2){
            name = entry.fields[0];
            age = entry.fields[1];
        }
        return (name,age);
    }

}
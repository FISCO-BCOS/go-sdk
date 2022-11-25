// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

contract SystemConfigPrecompiled {
    function setValueByKey(string memory key, string memory value)
        public
        returns (int256)
    {}

    function getValueByKey(string memory key)
        public
        view
        returns (string memory, int256)
    {}
}

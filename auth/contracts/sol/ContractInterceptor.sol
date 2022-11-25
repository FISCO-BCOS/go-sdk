// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

import "./ContractAuthPrecompiled.sol";

contract ContractInterceptor {
    //function login(address account) public view {}

    //function logout(address account) public view {}

    /*
     * when deploy contract, create function will check the deploy contract auth firstly.
     * @param account
     */
    function create(address account) public view returns (bool) {
        ContractAuthPrecompiled auth = ContractAuthPrecompiled(address(0x1005));
        return auth.hasDeployAuth(account);
    }

    /*
     * when call method, call function will check the method access auth firstly.
     * @param methodAuthMgrAddr
     * @param methodId
     * @param account
     */
    function call(
        address contractAddr,
        bytes4 methodId,
        address account
    ) public view returns (bool) {
        return checkAccessMethodAuth(contractAddr, methodId, account);
    }

    /*
     * when send transaction by the method, sendTransaction function will check the method access auth firstly.
     * @param methodAuthMgrAddr
     * @param methodId
     * @param account
     */
    function sendTransaction(
        address contractAddr,
        bytes4 methodId,
        address account
    ) public view returns (bool) {
        return checkAccessMethodAuth(contractAddr, methodId, account);
    }

    /*
     * when access the method, checkAccessMethodAuth function will check the method access auth firstly.
     * @param methodAuthMgrAddr
     * @param methodId
     * @param account
     */
    function checkAccessMethodAuth(
        address contractAddr,
        bytes4 methodId,
        address account
    ) internal view returns (bool) {
        ContractAuthPrecompiled auth = ContractAuthPrecompiled(address(0x1005));
        return auth.checkMethodAuth(contractAddr, methodId, account);
    }
}

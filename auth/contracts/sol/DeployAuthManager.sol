// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

import "./BasicAuth.sol";

contract DeployAuthManager is BasicAuth {
    //1-white list; 2-black list
    uint8 public _deployAuthType;
    //store white account list
    mapping(address => bool) private _deployAuthWhiteMap;
    //store black account list
    mapping(address => bool) private _deployAuthBlackMap;

    constructor(
        address owner //committee manager address
    ) public {
        setOwner(owner);
        _deployAuthType = 0;
    }

    function setDeployAuthType(uint8 deployAuthType) public onlyOwner {
        require(
            deployAuthType == 1 || deployAuthType == 2,
            "deploy auth type must be 1 or 2."
        );
        _deployAuthType = deployAuthType;
    }

    /*
     * open account the auth of deploying contract
     */
    function openDeployAuth(address account) public onlyOwner {
        if (_deployAuthType == 1) {
            _deployAuthWhiteMap[account] = true;
        } else if (_deployAuthType == 2) {
            _deployAuthBlackMap[account] = false;
        }
    }

    /*
     * close account the auth of deploying contract
     */
    function closeDeployAuth(address account) public onlyOwner {
        if (_deployAuthType == 1) {
            _deployAuthWhiteMap[account] = false;
        } else if (_deployAuthType == 2) {
            _deployAuthBlackMap[account] = true;
        }
    }

    /*
     * check account if has auth of deploying contract
     */
    function hasDeployAuth(address account) public view returns (bool) {
        if (_deployAuthType == 0) return true;
        if (_deployAuthType == 1 && _deployAuthWhiteMap[account]) return true;
        if (_deployAuthType == 2 && !_deployAuthBlackMap[account]) return true;
        return false;
    }
}

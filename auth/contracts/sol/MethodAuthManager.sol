// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

import "./BasicAuth.sol";

contract MethodAuthManager is BasicAuth {
    //admin: contract admin, manage contract access auth
    address public _admin;
    address public _contractAddress;

    //method auth type: mapping[methodId][aclType]
    //aclType : 1-wirte list;2-black list
    mapping(bytes4 => uint8) private _methodAuthTypeMap;

    //method auth whitelist management: mapping[methodId][accAddress][bool]
    mapping(bytes4 => mapping(address => bool)) private _methodAuthWhiteMap;
    //method auth blacklist management: mapping[methodId][accAddress][bool]
    mapping(bytes4 => mapping(address => bool)) private _methodAuthBlackMap;

    constructor(
        address contractAddress,
        address admin,
        address owner
    ) public {
        _contractAddress = contractAddress;
        _admin = admin;
        setOwner(owner);
    }

    modifier onlyAdmin() {
        require(_admin == msg.sender, "you must be admin");
        _;
    }

    /*
     * set contract method access auth type
     * @param methodId
     * @param authType: 1-witer list; 2-black list
     */
    function setMethodAccessAuthType(bytes4 methodId, uint8 authType)
        public
        onlyAdmin
    {
        require(authType == 1 || authType == 2, "auth type must be 1 or 2.");
        _methodAuthTypeMap[methodId] = authType;
    }

    /*
     * open contract method access auth for account
     */
    function openMethodAccessAuth(bytes4 methodId, address account)
        public
        onlyAdmin
    {
        require(
            _methodAuthTypeMap[methodId] != 0,
            "you should set the method access auth type firstly."
        );
        if (_methodAuthTypeMap[methodId] == 1) {
            _methodAuthWhiteMap[methodId][account] = true;
        } else if (_methodAuthTypeMap[methodId] == 2) {
            _methodAuthBlackMap[methodId][account] = false;
        }
    }

    /*
     * close contract method access auth for account
     */
    function closeMehtodAccessAuth(bytes4 methodId, address account)
        public
        onlyAdmin
    {
        require(
            _methodAuthTypeMap[methodId] != 0,
            "you should set the method access auth type firstly."
        );
        if (_methodAuthTypeMap[methodId] == 1) {
            _methodAuthWhiteMap[methodId][account] = false;
        } else if (_methodAuthTypeMap[methodId] == 2) {
            _methodAuthBlackMap[methodId][account] = true;
        }
    }

    /*
     * reset method auth contract admin
     */
    function resetAdmin(address admin) public onlyOwner {
        _admin = admin;
    }

    /*
     * judge account if has contract method access auth
     */
    function hasMethodAccessAuth(bytes4 methodId, address account)
        public
        view
        returns (bool)
    {
        if (_methodAuthTypeMap[methodId] == 0) return true;
        if (
            _methodAuthTypeMap[methodId] == 1 &&
            _methodAuthWhiteMap[methodId][account]
        ) return true;
        if (
            _methodAuthTypeMap[methodId] == 2 &&
            !_methodAuthBlackMap[methodId][account]
        ) return true;
        return false;
    }
}

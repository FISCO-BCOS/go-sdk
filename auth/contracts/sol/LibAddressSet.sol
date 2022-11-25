// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.10 <0.8.20;

library LibAddressSet {
    struct AddressSet {
        mapping(address => uint256) indexMapping;
        address[] values;
    }

    function add(AddressSet storage self, address value) internal {
        require(value != address(0x0), "LibAddressSet: value can't be 0x0");
        require(
            !contains(self, value),
            "LibAddressSet: value already exists in the set."
        );
        self.values.push(value);
        self.indexMapping[value] = self.values.length;
    }

    function contains(AddressSet storage self, address value)
        internal
        view
        returns (bool)
    {
        return self.indexMapping[value] != 0;
    }

    function remove(AddressSet storage self, address value) internal {
        require(contains(self, value), "LibAddressSet: value doesn't exist.");
        uint256 toDeleteindexMapping = self.indexMapping[value] - 1;
        uint256 lastindexMapping = self.values.length - 1;
        address lastValue = self.values[lastindexMapping];
        self.values[toDeleteindexMapping] = lastValue;
        self.indexMapping[lastValue] = toDeleteindexMapping + 1;
        delete self.indexMapping[value];
        self.values.pop();
    }

    function getSize(AddressSet storage self) internal view returns (uint256) {
        return self.values.length;
    }

    function get(AddressSet storage self, uint256 index)
        internal
        view
        returns (address)
    {
        return self.values[index];
    }

    function getAll(AddressSet storage self)
        internal
        view
        returns (address[] memory)
    {
        address[] memory output = new address[](self.values.length);
        for (uint256 i; i < self.values.length; i++) {
            output[i] = self.values[i];
        }
        return output;
    }
}

pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;


contract Precompiledbn256 {
    constructor() public {}

    function Bn256Add(bytes32 ax, bytes32 ay, bytes32 bx, bytes32 by)
        public
        returns (bytes32[2] memory result)
    {
        bytes32[4] memory input;
        input[0] = ax;
        input[1] = ay;
        input[2] = bx;
        input[3] = by;
        assembly {
            let success := call(gas, 0x06, 0, input, 0x80, result, 0x40)
            switch success
                case 0 {
                    revert(0, 0)
                }
        }
    }

    function Bn256ScalarMul(bytes32 x, bytes32 y, bytes32 scalar)
        public
        returns (bytes32[2] memory result)
    {
        bytes32[3] memory input;
        input[0] = x;
        input[1] = y;
        input[2] = scalar;
        assembly {
            let success := call(gas, 0x07, 0, input, 0x60, result, 0x40)
            switch success
                case 0 {
                    revert(0, 0)
                }
        }
    }

    function Bn256Pairing(bytes memory input) public returns (bytes32 result) {
        // input is a serialized bytes stream of (a1, b1, a2, b2, ..., ak, bk) from (G_1 x G_2)^k
        uint256 len = input.length;
        require(len % 192 == 0);
        assembly {
            let memPtr := mload(0x40)
            let success := call(
                gas,
                0x08,
                0,
                add(input, 0x20),
                len,
                memPtr,
                0x20
            )
            switch success
                case 0 {
                    revert(0, 0)
                }
                default {
                    result := mload(memPtr)
                }
        }
    }

    function BigModExp(bytes32 base, bytes32 exponent, bytes32 modulus)
        public
        returns (bytes32 result)
    {
        assembly {
            // free memory pointer
            let memPtr := mload(0x40)
            // length of base, exponent, modulus
            mstore(memPtr, 0x20)
            mstore(add(memPtr, 0x20), 0x20)
            mstore(add(memPtr, 0x40), 0x20)
            // assign base, exponent, modulus
            mstore(add(memPtr, 0x60), base)
            mstore(add(memPtr, 0x80), exponent)
            mstore(add(memPtr, 0xa0), modulus)
            // call the precompiled contract BigModExp (0x05)
            let success := call(gas, 0x05, 0x0, memPtr, 0xc0, memPtr, 0x20)
            switch success
                case 0 {
                    revert(0x0, 0x0)
                }
                default {
                    result := mload(memPtr)
                }
        }
    }

    // https://eips.ethereum.org/EIPS/eip-152
    function F(
        uint32 rounds,
        bytes32[2] memory h,
        bytes32[4] memory m,
        bytes8[2] memory t,
        bool f
    ) public view returns (bytes32[2] memory) {
        bytes32[2] memory output;

        bytes memory args = abi.encodePacked(
            rounds,
            h[0],
            h[1],
            m[0],
            m[1],
            m[2],
            m[3],
            t[0],
            t[1],
            f
        );

        assembly {
            if iszero(
                staticcall(not(0), 0x09, add(args, 32), 0xd5, output, 0x40)
            ) {
                revert(0, 0)
            }
        }

        return output;
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/MintGem.sol";
import "../src/KZGVerifier.sol";
import "forge-std/console.sol";

contract MintTest is Test, Verifier {
    Mint public mint;

    function setUp() public {
        mint = new Mint();
    }

    receive() external payable {}
    
    function testWeights() external view {
        // x = uint256[1,2,3,4,5,6,7,8,9,10]
        uint256[] memory x = new uint256[](10);
        for (uint256 i = 0; i < 10; i++) {
            x[i] = i + 1;
        }
        uint256[] memory weights = getBarycentricWeights(x);
        console.logUint(weights[0]);
        console.logUint(weights[1]);
        console.logUint(weights[2]);
    }
}

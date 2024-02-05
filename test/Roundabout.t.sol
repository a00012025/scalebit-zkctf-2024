// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console2} from "forge-std/Test.sol";
import {Roundabout} from "../src/Roundabout/Contract/verifier.sol";

contract RoundaboutTest is Test {
    Roundabout public roundabout;

    function setUp() public {
        roundabout = new Roundabout();
    }

    function testVerify() public {
        roundabout.verify(
            [
                0x175e46d103fabd516c95b041906c94f215674d38f007edcfd53745a503eac325,
                0x2e09b5fcb8799d77191fa7dff80ab86f40616a91f4c2fb40131ecddabeaa09df
            ],
            [
                [
                    0x24069e2981ecef69692727dc59ff430f22931b3bd61b51eb9384d0b18a9a3b03,
                    0x2b14cd37989af54ced367444600cf79d7906bd30b2fda4c5fe05bf91a9600261
                ],
                [
                    0x03c5d8b18e7778336f1c759bc26983140b845fa6c372adbc261db7eb4a37f39d,
                    0x167aa8436d8b18fc3b26695cd49490524e26a99a5fb468731d3e1073184361d7
                ]
            ],
            [
                0x2fb6722183bcb4562133dd2532a0efa7c6d420f72267f23750d9a5f6db1a07ee,
                0x1ee608bba33e593c219085478d6c05f88323d4c65c32fcca5b32cf76eec0ea80
            ],
            [uint(0x25d2f72)]
        );
        assertTrue(roundabout.isSolved());
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console2} from "forge-std/Test.sol";
import {Mixer} from "../src/Mixer/Mixer.sol";
import {Challenge} from "../src/Mixer/Challenge.sol";

contract CheckinTest is Test {
    Challenge public ch;

    function setUp() public {
        Mixer mixer = new Mixer();
        ch = new Challenge(mixer);
    }

    function testVerify() public {
        assertTrue(ch.isSolved());
    }
}

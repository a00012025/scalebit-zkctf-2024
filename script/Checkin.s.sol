// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console2} from "forge-std/Script.sol";
import {Checkin} from "../src/CheckIn/verifier.sol";

contract SolveScript is Script {
    function setUp() public {}

    function run() public {
        Checkin checkin = Checkin(0x4cD12eE863345551f23055e32758b3EDC76e2fE1);
        vm.startBroadcast();
        checkin.verify(
            [
                0x0de868e308ad5c7f0f8701b1fbf2c78a65f6abc4490f4a7fe3aa39a8c28d571b,
                0x1b747bef08f7d12d8935e2e625efb7e424c12e2b7c0df4c31d355a9b21789736
            ],
            [
                [
                    0x225b3ed4742b98188ed99860f39a9ff45c996be697c1bb070545f1b1934498cc,
                    0x151ca4e3a630224134bbcea93e8529132fe7c56ffe2e51fb773433b37124d4b7
                ],
                [
                    0x09db935dfe20716f461a33af25741ede2e651d5a6b4f8a7c334033a103e018c0,
                    0x28b4e30552d8de83b412060d1600261c2b1b7554709781ae1a68754c014a3f5e
                ]
            ],
            [
                0x255569acd9c32bc51824d19368988866b594cd0cb62b16391b97c234e5409883,
                0x1c84b5c94cf6d0d12e3ddb94db061f94adc4ccac58961afdff9a7fafa9a662ef
            ],
            [uint(0x21)]
        );
        vm.stopBroadcast();
    }
}

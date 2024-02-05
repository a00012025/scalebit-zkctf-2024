# ZK  CFT - Halo2 - Is Zero

## Overview
1. This is a quiz about Halo2 usage. The standard answer would not be given to participants
2. Instead, a file would be given, with all necessary structs and missing logics. Participants are required to fill in the blanks, compile the program and run tests
3. Multiple test cases would be given, the prover program shall be tested against them

## Problem description

The is_zero circuit is an important component within the Halo2 zero-knowledge proof (ZKP) system, a common requirement is to prove that a certain value is zero without revealing the value itself. The is_zero circuit is designed to facilitate this need. It's a gadget or a sub-circuit that can be used within a larger circuit to prove that a given value is zero, which is crucial in various computations and conditional checks within a zero-knowledge proof system.

Now you are required to simulate an is_zero checking in Halo2. When the input is zero, it should output 1, otherwise, it should output 0.

## Comand
cargo run

## Example
input == 0 and output == 1
input == 2 and output == 0

## Solve
Sumbit is_zero.rs to solve the puzzle.

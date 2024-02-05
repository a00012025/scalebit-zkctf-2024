// Modified from https://github.com/weijiekoh/libkzg
// Modified from https://github.com/appliedzkp/semaphore/blob/master/contracts/sol/verifier.sol
pragma experimental ABIEncoderV2;
pragma solidity ^0.8.15;

import "./Pairing.sol";
import { Constants } from "./Constants.sol";

contract Verifier is Constants {

    using Pairing for *;

    // The G1 generator
    Pairing.G1Point SRS_G1_0 = Pairing.G1Point({
        X: Constants.SRS_G1_X[0],
        Y: Constants.SRS_G1_Y[0]
    });

    // The G2 generator
    Pairing.G2Point g2Generator = Pairing.G2Point({
        X: [ Constants.SRS_G2_X_0[0], Constants.SRS_G2_X_1[0] ],
        Y: [ Constants.SRS_G2_Y_0[0], Constants.SRS_G2_Y_1[0] ]

    });

    Pairing.G2Point SRS_G2_1 = Pairing.G2Point({
        X: [ Constants.SRS_G2_X_0[1], Constants.SRS_G2_X_1[1] ],
        Y: [ Constants.SRS_G2_Y_0[1], Constants.SRS_G2_Y_1[1] ]
    });

    function g1Check(Pairing.G1Point memory _point) internal pure
    {
        require(_point.X != 0 || _point.Y != 0, "G1 point at infinity");
        // check encoding
        require(_point.X < PRIME_Q);
        require(_point.Y < PRIME_Q);
        // check on curve
        uint256 lhs = mulmod(_point.Y, _point.Y, PRIME_Q); // y^2
        uint256 rhs = mulmod(_point.X, _point.X, PRIME_Q); // x^2
        rhs = mulmod(rhs, _point.X, PRIME_Q); // x^3
        rhs = addmod(rhs, bn254_b_coeff, PRIME_Q); // x^3 + b
        require(lhs == rhs);
    }

    function verify(
        Pairing.G1Point memory _commitment,
        Pairing.G1Point memory _proof,
        uint256 _index,
        uint256 _value
    ) public view returns (bool) {
        // Make sure each parameter is less than the prime q
        g1Check(_commitment);
        g1Check(_proof);
        require(_index < PRIME_Q, "Verifier.verifyKZG: _index is out of range");
        require(_value < PRIME_Q, "Verifier.verifyKZG: _value is out of range");
        
        // Compute commitment - aCommitment
        Pairing.G1Point memory commitmentMinusA = Pairing.plus(
            _commitment,
            Pairing.negate(
                Pairing.mulScalar(SRS_G1_0, _value)
            )
        );

        // Negate the proof
        Pairing.G1Point memory negProof = Pairing.negate(_proof);

        // Compute index * proof
        Pairing.G1Point memory indexMulProof = Pairing.mulScalar(_proof, _index);

        // Returns true if and only if
        // e((index * proof) + (commitment - aCommitment), G2.g) * e(-proof, xCommit) == 1
        return Pairing.pairing(
            Pairing.plus(indexMulProof, commitmentMinusA),
            g2Generator,
            negProof,
            SRS_G2_1
        );
    }

    function verifySoulBox(
        Pairing.G1Point memory _commitment,
        Pairing.G1Point memory _proof,
        Pairing.G1Point memory _soulBox
    ) public view returns (bool) {
        // Make sure each parameter is less than the prime q
        g1Check(_commitment);
        g1Check(_proof);
        uint256 _index = 0;

        // Compute commitment - aCommitment
        Pairing.G1Point memory commitmentMinusA = Pairing.plus(
            _commitment,
            Pairing.negate(
                _soulBox
            )
        );

        // _commitment - _soulBox = _proof

        // Negate the proof
        Pairing.G1Point memory negProof = Pairing.negate(_proof);

        // Compute index * proof
        Pairing.G1Point memory indexMulProof = Pairing.mulScalar(_proof, _index);

        // Returns true if and only if
        // e((index * proof) + (commitment - aCommitment), G2.g) * e(-proof, xCommit) == 1
        return Pairing.pairing(
            Pairing.plus(indexMulProof, commitmentMinusA),
            g2Generator,
            negProof,
            SRS_G2_1
        );
    }

    function fr_inverse(uint256 a) internal view returns (uint256) {
        require(a != 0, "Division by zero");
        return fr_pow(a, Constants.PRIME_R - 2);
    }

    function fr_pow(uint256 a, uint256 power) internal view returns (uint256) {
        uint256[6] memory input;
        uint256[1] memory result;
        bool ret;

        input[0] = 32;
        input[1] = 32;
        input[2] = 32;
        input[3] = a;
        input[4] = power;
        input[5] = Constants.PRIME_R;

        assembly {
            ret := staticcall(gas(), 0x05, input, 0xc0, result, 0x20)
        }
        require(ret);

        return result[0];
    }

    function fr_div(uint256 a, uint256 b) internal view returns (uint256) {
        require(b != 0);
        return mulmod(a, fr_inverse(b), Constants.PRIME_R);
    }

    function fr_mul_add(
        uint256 a,
        uint256 b,
        uint256 c
    ) internal pure returns (uint256) {
        return addmod(mulmod(a, b, Constants.PRIME_R), c, Constants.PRIME_R);
    }

    // Function to calculate barycentric weights
    function getBarycentricWeights(uint256[] memory xValues) public view returns (uint256[] memory) {
        uint256 m = Constants.PRIME_R;
        uint256 n = xValues.length;
        uint256[] memory weights = new uint256[](n);

        for (uint256 i = 0; i < n; i++) {
            weights[i] = 1;
            for (uint256 j = 0; j < n; j++) {
                if (i != j) {
                    uint256 difference = addmod(xValues[i], m - xValues[j], m);
                    weights[i] = mulmod(weights[i], difference, m);
                }
            }
            weights[i] = fr_div(1, weights[i]);
        }

        return weights;
    }

    function evaluateBarycentricPolynomial(uint256 x, uint256[] memory yValues, uint256[] memory weights) public view returns (uint256) {
        uint256 m = Constants.PRIME_R;

        require(yValues.length == weights.length, "Array lengths must match");

        uint256 n = weights.length;
        uint256 numerator = 0;
        uint256 denominator = 0;

        for (uint256 i = 0; i < n; i++) {
            if (x == (i+1)) {
                return yValues[i];
            }
            uint256 temp = 0;
            uint256 temp2 = 0;
            temp = addmod(x, m - (i+1), m);               // x - xValues[i]
            temp = fr_div(weights[i], temp);              // weights[i] / temp
            temp2 = mulmod(temp, yValues[i], m);          // temp * yValues[i]
            numerator = addmod(numerator, temp2, m);      // numerator + temp2
            denominator = addmod(denominator, temp, m);   // denominator + temp
        }

        return fr_div(numerator, denominator);
    }

}

pragma circom 2.0.0;

template Multiplier() {
    signal input a;
    signal input b;
    signal input c;
    signal output d;

    d <== (-3 * a) * (2 * a + 4) + 12 + 6 * b;
}

component main = Multiplier();

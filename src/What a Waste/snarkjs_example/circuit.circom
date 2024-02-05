pragma circom 2.0.0;

template Multiplier() {
    signal input a;
    signal input b;
    signal output c;

    signal int[2];

    int[0] <== a*a + 3*b;
    int[1] <== int[0]*int[0] + 2*b;

    c <== int[1];
}

component main = Multiplier();

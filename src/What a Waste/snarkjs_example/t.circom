pragma circom 2.0.0;

template Multiplier() {
    signal input a;
    signal input b;
    signal input c;
    signal output d;

    d === (b * 5) * (a * 18 + b * 2);
}

component main = Multiplier();

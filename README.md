# Scalebit CTF 2024 Writeup

Author: cbd1913

Problems:

- [https://github.com/scalebit/zkCTF-day1](https://github.com/scalebit/zkCTF-day1)
- [https://github.com/scalebit/zkCTF-day2](https://github.com/scalebit/zkCTF-day2)

# CheckIn (100 pt)

This is a simple circuit with one constraint `c = a * b` , so we can compile the circuit and provide an input, then use snarkjs to generate call data to verify it on chain.

```bash
circom CheckIn.circom --r1cs --wasm --sym --c -o ./output
cat <<EOT > input.json
{"a": 1, "b": 2}
EOT
node output/CheckIn_js/generate_witness.js output/CheckIn_js/CheckIn.wasm input.json witness.wtns
snarkjs groth16 prove CheckIn_groth16.zkey witness.wtns proof.json public.json
snarkjs generatecall
```

# **Kid's Math (200 pt)**

We need to implement a Fibonacci Chip using halo2. So a constraint `a + b = c`  should be added for each row, and make sure a new row’s a and b value equals previous row’s b and c value. More info can be found in the implementation of Fibonacci chip using halo2 here: [https://youtu.be/60lkR8DZKUA?si=6A658HfYwynzAAUz&t=1388](https://youtu.be/60lkR8DZKUA?si=6A658HfYwynzAAUz&t=1388)

# **Int Division (200 pt)**

We need to enforce integer division constraints in this problem, including `a = b * c + r`  and `r < b` . Also, a range check should be performed to make sure all of the elements are within 0~255. I didn’t figure out how to write lookup table in the contest, so I used a brute force way to write it like:

```rust
let range_check = |range: usize, a: Expression<F>| {
    let mut value = F::ZERO;
    (1..range).fold(Expression::Constant(F::ONE), |expr, _| {
        let result = expr * (Expression::Constant(value) - a.clone());
        value = value + F::ONE;
        result
    })
};
constraints.push(s.clone() * range_check(255, a.clone()));
constraints.push(s.clone() * range_check(255, b.clone()));
constraints.push(s.clone() * range_check(255, c.clone()));
constraints.push(s.clone() * range_check(255, r.clone()));
constraints.push(s.clone() * range_check(255, k.clone()));
```

# **Roundabout (300 pt)**

Trace the code of `MiMCSponge`  line by line and we can find that its output equals to `(a+1)^5` , therefore we can factorize `3066844496547985532785966973086993824`  and find `a = 19830713` 

For second part, just replace `c` value in the equation then we can solve `b = 2` .

# **Ethereal Quest (400 pt)**

There are two steps to solve this problem. First is minting 10 gems, and second is minting 20. If we run the contract and mint program locally then send several mint transactions, we will notice that when minting the 10th gem, it has following error: `Soul dispersed` 

The main reason of it is: When minter provides too many points on the polynomial (which should firstly pass the KZG commitment verification), the contract can recover the original polynomial by Lagrange interpolation and get the private key of the hero (i.e. the evaluation at 0 of the polynomial)

So the relation between the number of points (`N`) which can be posted to contract (then mint a gem), and the degree of original polynomial (`d`), is that `N = d`  (i.e. when posting the `d+1` th points it can recover the polynomial and revert)

Checking the polynomial generation code in `ForgeSword()` , we see that it’s generating a polynomial with degree 9, that’s why we can only mint 9 gems. So the solution of first step is changing 10 to 12 so that we can mint more gems.

```go
f := make([]fr.Element, 12)
```

But it doesn’t work for the second step, because we only have 16 power of tau values in the contract, which can only support a polynomial up to degree 16! We need another way to bypass the check.

The condition to pass step 2 is that, we need to post at least 20 points to the smart contract which will pass KZG verification, but they don’t lie on a polynomial with degree 16. So we must forge part of the values we provide to the contract.

When calling `mint`  function, the contract will ensure that `P(i) = value` with the given KZG commitment `comm(P)` , and it’s a standard implementation so we must act honestly when providing these values. Which means anyway we need to generate a polynomial (with degree ≤ 16) and evaluate it on `i = 1~20`  honestly to pass the verification.

The vulnerability is in `verifySoulBox()` , we can see that it performs following check on `comm(P)` , `proof` , `soulBox` **:

```solidity
e(comm(P) - soulBox, G2.g) * e(-proof, τ * G2.g) == 1
```

Where `τ`  is secret from KZG commitment (`_index`  = 0 so `index * proof` ** is omitted). It doesn’t verify that whether caller knows the discrete log of `soulBox` , so we can provide a value that we don’t know the private key.

We want the forged values satisfy: `τ * proof = comm(P) - soulBox` , so just setting `soulBox = comm(P) - τ * G1.g`  and `proof = G1.g`, which can pass the verification.

```go
tauG1 := &SRS.Pk.G1[1]
soulBoxPoint := new(bn254.G1Affine)
soulBoxPoint.Add(soulBoxPoint, commitmentPoint)
soulBoxPoint.Sub(soulBoxPoint, tauG1)
bladeProofPoint := G1AffineToG1Point(g1)
```

Because the `soulBox` value we forge doesn’t correspond to the actual `P(0)` value, it will always pass the contract’s check on polynomial recovery, so we can mint as many gems as we want now.

# **Is Zero (200 pt)**

We can introduce advice `c`  which equals to `1/a`  when `a`  is not zero, or equals to 0 when `a`  is zero. So the constraint should be: `(b - 1) * (a * c - 1) = 0`  

We still need to prevent `(a, b, c) = (1, 2, 1)`  to pass the verification, so another constraint `a * b = 0`  should be added

And prevent `(a, b, c) = (0, 2, 0)` to pass the verification, so `b * (1-b) = 0`  should also be added. The rest of the code is to assign correct value in `synthesize` .

# **Familiar Strangers (300 pt)**

For level 1, the provided 2 values can be written as:

`A = 6026017665971213533282357846279359759458261226685473132380160 = 2^201 + 2^200 + 2^199 + 2^198` 

`B = -401734511064747568885490523085290650630550748445698208825344 = -2^201 + 2^200 + 2^199 + 2^198`

We need to find `x`  to pass the check of `x < A`  and `B < x` . The circom `LessThan(201)` is used to check these constraints (`GreaterThan`  is also `LessThan` with reversed input order).

Following is the implementation of `LessThan` :

```go
template LessThan(n) {
    assert(n <= 252);
    signal input in[2];
    signal output out;

    component n2b = Num2Bits(n+1);

    n2b.in <== in[0]+ (1<<n) - in[1];

    out <== 1-n2b.out[n];
}
```

So we only need to make sure `in[0] + (2^n) - in[1]`  has value 0 at the `2^n` bit. And because `A`  and `B`  are overflow to the range of `[0, 2^201 - 1]` , the assumption of this circuit isn’t true which can be exploited.

If we provide `x = 2^200 + 2^199 + 2^198` , it can pass both constraints because:

`x + 2^201 - (2^201 + 2^200 + 2^199 + 2^198)` and `(-2^201 + 2^200 + 2^199 + 2^198) + 2^201 - x` both equal to 0, which means their `2^n`  bit is also zero, then we can solve level 1.

For level 2, the values can be written as:

`A = 3533700027045102098369050084895387317199177651876580346993442643999981568 = 2^241 + 2^221 + 2^220 + 2^219` 

`B = -3618502782768642773547390826438087570129142810943142283802299270005870559232 = -2^251 + 2^221 + 2^220 + 2^219` 

And the constraints we need to pass is also `x < A` (under 241 bit) and `B < x`  (under 251 bit)

We can use the same trick to pass the verification, i.e. using `x = 2^221 + 2^220 + 2^219` , but it can’t pass the length verification in `stranger_judge.js` , which requires `inputs.l2.length <= 70` 

But the judge code doesn’t check the range of provided input, so we can use `x + Fr`  to pass the verification, where `Fr`  is the scalar field size of BN254 curve. This is known is input aliasing issue.

# Mixer (300 pt)

After generating the challenge contract, we can notice that some transactions are sent to the contract, including 5 deposits and 2 withdraw.

In the `withdraw`  function, the contract only checks whether the current `nullifier` is used, and will mod the value using `BASE_FIELD_SIZE` . So we can replay other transactions’ proof data with  `nullifier + BASE_FIELD_SIZE` and preserve all of the other data.
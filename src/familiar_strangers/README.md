# Familiar Strangers

In the realm of the familiar, where logic gates are the silent sentinels of truth, lies a challenge shrouded in enigma. "Familiar Strangers" beckons you to a world where the simplest of circuits hide secrets just beyond the veil of obviousness. These circuits, reminiscent of the ones you've met countless times, now hold a mystery that only a true Circom savant can unravel. Two levels, each a step deeper into the cryptic dance of numbers and logic, await your prowess. Are you ready to discover the true inputs and reveal the concealed answers within? Find the key, communicate with our judge service, and claim your place among the elite who see beyond the familiar to the truth that lies beneath.

## How to play?

1. Install the latest circom, check <https://docs.circom.io/getting-started/installation/>.
2. Install the dependencies with `npm install`.
3. Read the [circuit](circuits/challenge.circom) and [stranger_judge.js](stranger_judge.js) carefully to understand the challenge.
4. Find the correct inputs and send to the judge service to win the challenge.

## How to run judge service?

For local testing, you can run the judge service with:

```shell
node stranger_judge.js
```

## How to send request to judge service?

Example request with curl:

```shell
curl -X POST http://localhost:3000/judge -H "Content-Type: application/json" -d '{"l1": "123", "l2": "123"}'
```

The `l1` and `l2` are the two levels of this challenge.

The service will return `{"success":true}` when the input is correct, otherwise it will return `{"success":false}`.

## Reference

We include the below files in README for your convenience. You should always read the latest code in the correct location.

The [challenge.circom](circuits/challenge.circom) is the circuits for this challenge.

```js
pragma circom 2.1.7;

include "../node_modules/circomlib/circuits/comparators.circom";

template Level1() {
    signal input in;
    signal output out;
    out <== 1;
    component lt = LessThan(201);
    lt.in[0] <== in;
    lt.in[1] <== 6026017665971213533282357846279359759458261226685473132380160;
    lt.out === out;
    component gt = GreaterThan(201);
    gt.in[0] <== in;
    gt.in[1] <== -401734511064747568885490523085290650630550748445698208825344;
    gt.out === out;
}

template Level2() {
    signal input in;
    signal output out;
    out <== 1;
    component lt = GreaterThan(241);
    lt.in[0] <== 3533700027045102098369050084895387317199177651876580346993442643999981568;
    lt.in[1] <== in;
    lt.out === out;
    component gt = LessThan(251);
    gt.in[0] <== -3618502782768642773547390826438087570129142810943142283802299270005870559232;
    gt.in[1] <== in;
    gt.out === out;
}
```

The [stranger_judge.js](stranger_judge.js) is the code for the judge service.

```js
const express = require('express');
const bodyParser = require('body-parser');
const { wasm: wasm_tester } = require('circom_tester');
const path = require('path');

const app = express();
app.use(bodyParser.json());

const PASS = "pass";
const FAIL = "fail";

async function runCircuit(circuitPath, input) {
    const circuit = await wasm_tester(circuitPath, {
        output: path.join(__dirname, "./test/generated"),
        // recompile: true,
    });
    try {
        const witness = await circuit.calculateWitness({ in: input });
        await circuit.checkConstraints(witness);
        return PASS;
    } catch (error) {
        return error.toString();
    }
}

// Example inputs {"l1": "123", "l2": "123"}
// Test with curl
// curl -X POST http://localhost:3000/judge -H "Content-Type: application/json" -d '{"l1": "123", "l2": "123"}'

app.post('/judge', async (req, res) => {
    // random number for logging
    const identifier = Math.floor(Math.random() * 100000000000);
    const inputs = req.body;
    console.log("identifier", identifier, "inputs", inputs);

    let result = FAIL;

    // Level 1
    result = await runCircuit(path.join(__dirname, "test/level1_test.circom"), inputs.l1);
    if (result != PASS) {
        res.json({ success: false });
        return;
    }

    // Level 2
    if (inputs.l2.length <= 70) {
        res.json({ success: false });
        return;
    }
    result = await runCircuit(path.join(__dirname, "test/level2_test.circom"), inputs.l2);
    if (result != PASS) {
        res.json({ success: false });
        return;
    }

    res.json({ success: true });
    console.log("identifier", identifier, "success");
    return;
});

const port = 3000;
app.listen(port, () => {
    console.log(`Stranger Judge Service running on port ${port}`);
});
```

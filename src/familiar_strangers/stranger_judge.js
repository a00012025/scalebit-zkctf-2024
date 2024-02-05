const express = require("express");
const bodyParser = require("body-parser");
const { wasm: wasm_tester } = require("circom_tester");
const path = require("path");

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

app.post("/judge", async (req, res) => {
  // random number for logging
  const identifier = Math.floor(Math.random() * 100000000000);
  const inputs = req.body;
  console.log("identifier", identifier, "inputs", inputs);

  let result = FAIL;

  // Level 1
  result = await runCircuit(
    path.join(__dirname, "test/level1_test.circom"),
    inputs.l1
  );
  if (result != PASS) {
    console.log("Level 1 no");
    res.json({ success: false });
    return;
  }

  // Level 2
  if (inputs.l2.length <= 70) {
    res.json({ success: false });
    return;
  }
  result = await runCircuit(
    path.join(__dirname, "test/level2_test.circom"),
    inputs.l2
  );
  if (result != PASS) {
    console.log("Level 2 no");
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

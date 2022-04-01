const { readFileSync, writeFileSync } = require("fs");
const path = require("path");

require("wabt")().then(wabt => {
  const inputWat = "main.wat";
  const outputWasm = "main.wasm";

  const wasmModule = wabt.parseWat(inputWat, readFileSync(inputWat, "utf8"));
  const { buffer } = wasmModule.toBinary({});

  writeFileSync(outputWasm, new Buffer(buffer));
});

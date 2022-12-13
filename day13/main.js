const fs = require("fs");

const input = fs
  .readFileSync("./input.txt")
  .toString()
  .split("\n")
  .filter((value) => value != "")
  .map((value) => eval(value));

const compare = (a, b) => {
  if (!Array.isArray(a) && !Array.isArray(b)) {
    if (a < b) return 1;
    if (a > b) return -1;
    return 0;
  }

  if (!Array.isArray(a)) a = [a];
  if (!Array.isArray(b)) b = [b];

  const min = Math.min(a.length, b.length);
  for (let i = 0; i < min; i++) {
    const res = compare(a[i], b[i]);
    if (res != 0) return res;
  }

  if (a.length == b.length) return 0;
  if (a.length > b.length) return -1;
  return 1;
};

// part a
let count = 1;
let sum = 0;
for (let i = 0; i < input.length; i += 2) {
  const a = input[i];
  const b = input[i + 1];
  if (compare(a, b) == 1) {
    sum += count;
  }
  count += 1;
}
console.log("a:", sum);

// part b
input.push([[2]]);
input.push([[6]]);
const l = input.length;
for (let i = 0; i < l - 1; i++) {
  for (let j = i + 1; j < l; j++) {
    const res = compare(input[i], input[j]);
    if (res == -1) {
      [input[i], input[j]] = [input[j], input[i]];
    }
  }
}

const pack1 = input.findIndex(
  (v) => JSON.stringify(v) == JSON.stringify([[2]])
);
const pack2 = input.findIndex(
  (v) => JSON.stringify(v) == JSON.stringify([[6]])
);
console.log("b:", (pack1 + 1) * (pack2 + 1));

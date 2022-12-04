const fs = require('fs');

const input = fs
  .readFileSync('./input.txt')
  .toString()
  .split('\n')
  .filter((row) => row != '')
  .map((row) =>
    row.split(',').map((pair) => pair.split('-').map((v) => Number.parseInt(v)))
  );

const a = input.filter((pair) => {
  if (pair[0][0] <= pair[1][0] && pair[0][1] >= pair[1][1]) {
    return true;
  }
  if (pair[1][0] <= pair[0][0] && pair[1][1] >= pair[0][1]) {
    return true;
  }
  return false;
});

const b = input.filter((pair) => {
  if (pair[0][0] >= pair[1][0] && pair[0][0] <= pair[1][1]) {
    return true;
  }
  if (pair[0][1] >= pair[1][0] && pair[0][1] <= pair[1][1]) {
    return true;
  }
  if (pair[1][0] >= pair[0][0] && pair[1][0] <= pair[0][1]) {
    return true;
  }
  if (pair[1][1] >= pair[0][0] && pair[1][0] <= pair[0][1]) {
    return true;
  }
  return false;
});

console.log(`a: ${a.length}`);
console.log(`b: ${b.length}`);

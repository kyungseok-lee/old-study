"use strict";
debugger;

let result = [1, 2, 3].find((value, index, allData) => value === 2);
console.log(result); // 2

result = [1, 2, 1].find(function (value, index, allData) {
  return value === 1 && value === this.key;
}, { key: 1 });
console.log(result); // 1

result = [1, 2, 1].find(function (value, index, allData) {
  return value === 1 && value === this.key;
}, { key: 3 });
console.log(result); // 1

let dummy;

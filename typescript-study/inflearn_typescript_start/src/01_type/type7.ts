export {};

// union 교집합 &

let v1: (1 | 3 | 5) & (3 | 5 | 7); // (3 | 5)
// v1 = 1;
v1 = 3;
v1 = 5;
// v1 = 6;
// v1 = 7;
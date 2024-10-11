// Monkey programming language clojures 
let add = fn(a, b) { a + b; };
let addTwo = fn(a) { add(a, 2); };
let addThree = fn(a) { add(a, 3); };
let applyFunc = fn(a, b, func) { func(a, b); };
applyFunc(3, 4, add); // 7
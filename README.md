# Monkey

Monkey is a toy programming language that is dynamically typed and has a C-like syntax. This project is an interpreter for the Monkey programming language.

It is based on the book [Writing An Interpreter In Go](https://interpreterbook.com/).

## Instructions to run 
```bash
$ go run main.go
# or if you want to run a file
$ go run main.go -file <filename>.monkey
```

To learn more about available commands:
```bash
$ go run main.go -help
```

## Features
- [x] Mathematical expressions
- [x] Variable bindings
- [x] functions
- [x] conditionals
- [x] return statements
- [x] higher-order functions
- [x] closures
- [x] integers
- [x] booleans
- [x] strings
- [x] arrays
- [x] hashes
  

## Examples:
### Church Encoding
```rust
let to_integer = fn(proc) { 
    return proc(fn(x) { return x + 1 })(0) 
};

let ZERO = fn(f) { fn(x) { x } }; 
let ONE = fn(f) { fn(x) { f(x) } };
let TWO = fn(f) { fn(x) { f(f(x)) } };
let THREE = fn(f) { fn(x) { f(f(f(x))) } };

let EXP = fn(m) { fn(n) { m(n) } };
let SUCC = fn(n) { fn(f) { fn(x) { f(n(f)(x)) } } };


puts(to_integer(TWO));
puts("succ one: ", to_integer(SUCC(ONE)));
puts("exp two three: ", to_integer(EXP(TWO)(THREE)));
puts("number 10: ", to_integer(fn(f) { fn(x) { f(f(f(f(f(f(f(f(f(f(x)))))))))) } }));
```

### Fibonacci
```rust
let fibonacci = fn(x) {
    if (x == 0) {
        return 0;
    } else {
        if (x == 1) {
            return 1;
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

let result = fibonacci(10);
puts(result); // 55
```

### higher-order functions
```rust
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            accumulated
        } else {
            iter(rest(arr), push(accumulated, f(first(arr))))
        }
    };
    iter(arr, []);
};

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)))
        }
    };
    iter(arr, initial)
};

let doubled = map([1, 2, 3, 4, 5], fn(x) {
    return x * 2
});
puts((doubled)); // [2, 4, 6, 8, 10]

let sum = reduce([1, 2, 3, 4, 5], 0, fn(acc, value) {
    return acc + value
});
puts(sum); // 15
```

### Closures
```rust
let add = fn(a, b) { a + b; };
let addTwo = fn(a) { add(a, 2); };
let addThree = fn(a) { add(a, 3); };
let applyFunc = fn(a, b, func) { func(a, b); };
applyFunc(3, 4, add); // 7
```
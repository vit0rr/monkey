let to_integer = fn(proc) { proc(fn(x) { x + 1 })(0) };


let ZERO = fn(f) { fn(x) { x } }; 
let ONE = fn(f) { fn(x) { f(x) } };
let TWO = fn(f) { fn(x) { f(f(x)) } };
let THREE = fn(f) { fn(x) { f(f(f(x))) } };

let TRUE = fn(x) { fn(y) { x } };
let FALSE = fn(x) { fn(y) { y } };

let EXP = fn(m) { fn(n) { m(n) } };
let SUCC = fn(n) { fn(f) { fn(x) { f(n(f)(x)) } } };


// puts(to_integer(TWO));
// puts("succ one: ", to_integer(SUCC(ONE)));
// puts("exp two three: ", to_integer(EXP(TWO)(THREE)));
// puts("number 10: ", to_integer(fn(f) { fn(x) { f(f(f(f(f(f(f(f(f(f(x)))))))))) } }));


let user = {
    "name": "John",
    "age": 30,
    "is_admin": true,
}

let get_age = fn(user) { user["age"] };
let get_name = fn(user) { user["name"] };
let get_is_admin = fn(user) { user["is_admin"] };

puts(get_age(user));
puts(get_name(user));
puts(get_is_admin(user));

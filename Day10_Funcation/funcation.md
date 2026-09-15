# 🧩 Go Functions

## 1. What is a Function?

A function is a reusable block of code that performs a specific task.

Instead of writing the same code multiple times, we can put it inside a function and call it whenever needed.

Basic syntax:

```go
func functionName(parameters) returnType {
    // code
}
```

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

Calling the function:

```go
sum := add(1, 2)

fmt.Println(sum)
```

Output:

```text
3
```

---

# 2. Function Parameters

Parameters are values that a function receives.

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

Here:

```text
a int
b int
```

are parameters.

When calling:

```go
add(10, 20)
```

`10` goes into `a` and `20` goes into `b`.

---

# 3. Same Type Parameters

If multiple parameters have the same type, Go allows a shorter syntax.

Instead of:

```go
func sub(a int, b int) int {
    return a - b
}
```

We can write:

```go
func sub(a, b int) int {
    return a - b
}
```

Both are equivalent.

Example:

```go
func multiply(a, b int) int {
    return a * b
}
```

---

# 4. Return Value

A function can return a value.

Example:

```go
func add(a, b int) int {
    return a + b
}
```

The final `int` means the function returns an integer.

```go
result := add(5, 3)

fmt.Println(result)
```

Output:

```text
8
```

The `return` statement sends the result back to the caller.

---

# 5. Multiple Return Values

One of Go's useful features is that a function can return multiple values.

Example:

```go
func getLanguages() (string, string, string) {
    return "golang", "javascript", "python"
}
```

This function returns **three strings**.

We can receive them separately:

```go
lang1, lang2, lang3 := getLanguages()

fmt.Println(lang1)
fmt.Println(lang2)
fmt.Println(lang3)
```

Output:

```text
golang
javascript
python
```

---

# 6. Ignore a Return Value Using `_`

Sometimes we don't need every returned value.

Example:

```go
lang1, lang2, _ := getLanguages()

fmt.Println(lang1, lang2)
```

Output:

```text
golang javascript
```

`_` is called the **blank identifier**.

It means:

> "I don't need this value."

This is very common in Go.

---

# 7. Function as an Argument

Go allows us to pass a function to another function.

Example:

```go
func processIt(fn func(a int) int) {
    fn(1)
}
```

Here:

```go
fn func(a int) int
```

means:

> `fn` must be a function that takes an `int` and returns an `int`.

For example:

```go
fn := func(a int) int {
    return a * 2
}

processIt(fn)
```

A function can therefore be treated like a value.

This is useful for:

* Callbacks
* Custom behavior
* Middleware
* Sorting
* Event handling
* Functional programming patterns

---

# 8. Anonymous Function

An anonymous function is a function without a name.

Example:

```go
fn := func(a int) int {
    return 2
}
```

There is no function name between `func` and `(`.

Normally:

```go
func add(a, b int) int {
    return a + b
}
```

Anonymous:

```go
func(a int) int {
    return 2
}
```

We can store an anonymous function in a variable:

```go
fn := func(a int) int {
    return 2
}

result := fn(6)

fmt.Println(result)
```

Output:

```text
2
```

---

# 9. Returning a Function

A function can also return another function.

Example:

```go
func processIt() func(a int) int {
    return func(a int) int {
        return 2
    }
}
```

The return type is:

```go
func(a int) int
```

This means:

> The function returns another function that takes an `int` and returns an `int`.

We can use it like this:

```go
fn := processIt()

result := fn(6)

fmt.Println(result)
```

Output:

```text
2
```

Flow:

```text
processIt()
     ↓
returns a function
     ↓
fn
     ↓
fn(6)
     ↓
2
```

---

# 10. Function as Argument + Function as Return Value

This is an important concept.

Example:

```go
func take(fn func(a int) int) func(a int) int {
    return func(a int) int {
        return fn(a)
    }
}
```

This function:

1. Takes a function as an argument.
2. Returns another function.

The parameter:

```go
fn func(a int) int
```

means:

```text
input  → int
output → int
```

The return type:

```go
func(a int) int
```

also means:

```text
input  → int
output → int
```

So the function signature is:

```text
Function
   │
   ├── takes → function(int) int
   │
   └── returns → function(int) int
```

---

# 11. Understanding `take()` Step by Step

Suppose:

```go
func double(a int) int {
    return a * 2
}
```

Now:

```go
fn := take(double)
```

Inside `take()`:

```go
return func(a int) int {
    return fn(a)
}
```

The returned function calls the original `fn`.

So:

```go
result := fn(5)
```

becomes:

```text
5
 ↓
double(5)
 ↓
10
```

Therefore:

```go
fmt.Println(result)
```

prints:

```text
10
```

---

# 12. Higher-Order Functions

A function is called a **higher-order function** when it does at least one of these:

* Takes another function as an argument.
* Returns a function.

Example:

```go
func take(fn func(int) int) func(int) int {
    return func(a int) int {
        return fn(a)
    }
}
```

`take()` is a higher-order function because it both:

* accepts a function
* returns a function

---

# 13. Function Type

This:

```go
func(a int) int
```

is a function type.

It describes a function with:

```text
Input:  int
Output: int
```

Example functions matching this type:

```go
func double(a int) int {
    return a * 2
}
```

```go
func square(a int) int {
    return a * a
}
```

Both can be used where:

```go
func(a int) int
```

is expected.

---

# 14. Anonymous Function Immediately Called

An anonymous function can also be called immediately.

```go
result := func(a int) int {
    return a * 2
}(5)

fmt.Println(result)
```

Output:

```text
10
```

The important part is:

```go
}(5)
```

The `(5)` immediately calls the anonymous function.

---

# 15. Functions Are First-Class Values

In Go, functions can be treated as values.

We can:

### Store a function in a variable

```go
fn := add
```

### Pass a function to another function

```go
processIt(add)
```

### Return a function

```go
fn := processIt()
```

### Call the stored function

```go
fn(5)
```

This is why Go functions are often described as **first-class values**.

---

# 16. Important Difference: Function vs Function Call

This is very important.

### Passing a function

```go
take(double)
```

Here we are passing the function itself.

### Calling a function

```go
double(5)
```

Here we are executing the function.

Think:

```text
double     → function
double(5)  → result of calling function
```

---

# 17. Your `processIt()` Example

Your code:

```go
func processIt() func(a int) int {
    return func(a int) int {
        return 2
    }
}
```

Then:

```go
fn := processIt()
```

At this point `fn` contains a function.

Calling:

```go
fn(6)
```

executes that returned function.

However, in your current code:

```go
fn(6)
```

the returned value is `2`, but you are not storing or printing it.

If you want to see it:

```go
result := fn(6)

fmt.Println(result)
```

Output:

```text
2
```

---

# 18. Function Without Parameters

Functions don't always need parameters.

```go
func greet() {
    fmt.Println("Hello")
}
```

Call:

```go
greet()
```

---

# 19. Function Without Return Value

A function can perform an action without returning anything.

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

There is no return type after `)`.

---

# 20. Function With Multiple Parameters and Return

Example:

```go
func calculate(a, b int) int {
    return a + b
}
```

Call:

```go
result := calculate(10, 20)
```

---

# ⭐ Pro Tips

1. Use functions to keep code reusable and readable.

2. If parameters have the same type:

```go
func add(a, b int) int
```

is cleaner than:

```go
func add(a int, b int) int
```

3. Go supports multiple return values:

```go
func getData() (string, int) {
    return "Go", 10
}
```

4. Use `_` when you don't need a returned value:

```go
name, _, _ := getData()
```

5. `func(a int) int` is a **function type**.

6. Anonymous functions don't have a name:

```go
func(a int) int {
    return a * 2
}
```

7. Functions can be stored in variables:

```go
fn := add
```

8. Functions can be passed as arguments.

9. Functions can return other functions.

10. A function that takes or returns another function is called a **higher-order function**.

11. Don't confuse:

```go
add
```

with:

```go
add(1, 2)
```

The first is the function itself; the second calls the function.

---

# 🧠 What I Learned Today

* A function is a reusable block of code.
* Functions can accept parameters.
* Functions can return values.
* Go supports multiple return values.
* `_` can ignore unwanted return values.
* Functions can be stored in variables.
* Anonymous functions don't have names.
* A function can accept another function as an argument.
* A function can return another function.
* `func(a int) int` represents a function type.
* Functions can behave like first-class values.
* Higher-order functions can accept or return other functions.

---

# 📌 Quick Reference

```go
// Normal function
func add(a, b int) int {
    return a + b
}

// Call
result := add(1, 2)

// Multiple return values
func getData() (string, int) {
    return "Go", 10
}

// Receive multiple values
name, age := getData()

// Ignore value
name, _ := getData()

// Anonymous function
fn := func(a int) int {
    return a * 2
}

// Function as argument
func process(fn func(int) int) {
    fmt.Println(fn(10))
}

// Function as return value
func create() func(int) int {
    return func(a int) int {
        return a * 2
    }
}

// Higher-order function
func take(fn func(int) int) func(int) int {
    return func(a int) int {
        return fn(a)
    }
}
```

> **Key idea:** In Go, functions are values. You can store them, pass them to other functions, and return them from functions.

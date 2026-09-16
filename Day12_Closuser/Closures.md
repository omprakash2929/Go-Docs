# 🔐 Go Closures

## 1. What is a Closure?

A **closure** is a function that can remember and access variables from the surrounding scope where it was created.

In simple words:

> A closure is a function + the variables it remembers from its surrounding environment.

Example:

```go
func counter() func() int {
    var count int = 0

    return func() int {
        count += 1
        return count
    }
}
```

The anonymous function remembers the `count` variable.

---

# 2. Understanding the `counter()` Example

Your complete code:

```go
package main

import "fmt"

func counter() func() int {
    var count int = 0

    return func() int {
        count += 1
        return count
    }
}

func main() {

    increment := counter()

    fmt.Println(increment())
    fmt.Println(increment())
}
```

Output:

```text
1
2
```

Why does the second call return `2` instead of `1`?

Because the returned function **remembers `count`**.

---

# 3. Step-by-Step Execution

When this runs:

```go
increment := counter()
```

`counter()` starts executing.

Inside it:

```go
var count int = 0
```

So:

```text
count = 0
```

Then `counter()` returns this function:

```go
func() int {
    count += 1
    return count
}
```

That function is stored in:

```go
increment
```

So conceptually:

```text
increment
    │
    ▼
┌───────────────────┐
│ function          │
│                   │
│ count += 1        │
│ return count      │
└─────────┬─────────┘
          │
          ▼
     remembers
     count = 0
```

---

# 4. First Call

Now:

```go
increment()
```

The closure executes:

```go
count += 1
```

Current value:

```text
count = 0
```

After increment:

```text
count = 1
```

Then:

```go
return count
```

returns:

```text
1
```

---

# 5. Second Call

Now we call:

```go
increment()
```

again.

The closure still remembers the same `count`.

Current value:

```text
count = 1
```

Then:

```go
count += 1
```

becomes:

```text
count = 2
```

So it returns:

```text
2
```

That's why the output is:

```text
1
2
```

---

# 6. Why Doesn't `count` Disappear?

Normally, you might think:

```go
func counter() {
    count := 0
}
```

After `counter()` finishes, `count` is no longer accessible directly.

But here the returned function still uses `count`:

```go
return func() int {
    count += 1
    return count
}
```

Go keeps the captured variable available because the returned function still needs it.

So:

```text
counter()
   │
   ├── creates count
   │
   └── returns closure
            │
            └── remembers count
```

The closure keeps access to the variable.

---

# 7. Closure = Function + Captured Variables

A useful mental model:

```text
Closure
   =
Function
   +
Variables from surrounding scope
```

For your example:

```text
Closure
   │
   ├── function: func() int
   │
   └── captured variable: count
```

---

# 8. Closure Can Modify Captured Variables

Closures don't just read variables.

They can also modify them.

Example:

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

The closure changes:

```go
count
```

every time it is called.

---

# 9. Multiple Closures Have Separate State

This is a very important concept.

Consider:

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

Now:

```go
counter1 := counter()
counter2 := counter()
```

These are **two different closure instances**.

Calling:

```go
fmt.Println(counter1())
fmt.Println(counter1())
```

gives:

```text
1
2
```

Calling:

```go
fmt.Println(counter2())
```

gives:

```text
1
```

Why?

Because each call to `counter()` creates its own `count`.

Conceptually:

```text
counter1
   │
   └── count = 2


counter2
   │
   └── count = 1
```

They don't share the same `count`.

---

# 10. Closure as Private State

Closures can be used to create **private state**.

Example:

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

From `main()` we cannot directly access:

```go
count
```

We can only interact with it through:

```go
increment()
```

So:

```text
main()
  │
  │ calls
  ▼
increment()
  │
  ▼
private count
```

This is similar to keeping some state hidden inside an object.

---

# 11. Closure vs Normal Function

### Normal function

```go
func add(a, b int) int {
    return a + b
}
```

It doesn't remember previous calls.

```go
add(1, 2)
add(1, 2)
```

Both calls are independent.

### Closure

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

The closure remembers state:

```text
call 1 → 1
call 2 → 2
call 3 → 3
```

---

# 12. Closure With Different Variables

A closure can capture any variable from an outer scope.

Example:

```go
func multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}
```

Now:

```go
double := multiplier(2)

fmt.Println(double(5))
fmt.Println(double(10))
```

Output:

```text
10
20
```

The returned function remembers:

```text
x = 2
```

So:

```text
double(5)
   ↓
2 × 5
   ↓
10
```

---

# 13. Another Example: Greeting

```go
func greeting(name string) func() string {
    return func() string {
        return "Hello " + name
    }
}
```

Then:

```go
helloOm := greeting("Omprakash")

fmt.Println(helloOm())
```

Output:

```text
Hello Omprakash
```

The closure remembers `name`.

---

# 14. Common Uses of Closures

Closures are useful for:

* Maintaining state
* Counters
* Configuration
* Callbacks
* Middleware
* Function factories
* Encapsulation
* HTTP handlers
* Goroutines and concurrency patterns

For example, an HTTP handler can capture configuration without making that configuration globally accessible.

---

# 15. Closure and Function Factory

A function that creates and returns another function is often called a **function factory**.

Example:

```go
func multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}
```

We can create different functions:

```go
double := multiplier(2)
triple := multiplier(3)
```

Now:

```go
fmt.Println(double(10))
fmt.Println(triple(10))
```

Output:

```text
20
30
```

Conceptually:

```text
multiplier(2)
     ↓
double function
     ↓
remembers x = 2


multiplier(3)
     ↓
triple function
     ↓
remembers x = 3
```

---

# 16. Important Difference: Closure vs Anonymous Function

These two concepts are related but not exactly the same.

### Anonymous function

A function without a name:

```go
func(a int) int {
    return a * 2
}
```

### Closure

A function that **captures variables from its surrounding scope**:

```go
x := 10

func() {
    fmt.Println(x)
}
```

The function captures `x`, so it is a closure.

Therefore:

> An anonymous function is not automatically a closure. It becomes a closure when it captures variables from an outer scope.

---

# ⭐ Pro Tips

1. A closure remembers variables from its surrounding scope.

2. Your `counter()` is a classic closure example.

3. Each call to `counter()` creates a separate captured state.

4. The captured variable can be modified by the closure.

5. Closures are useful for maintaining state without using global variables.

6. Think of a closure as:

```text
Function + captured environment
```

7. Anonymous function ≠ always closure.

8. A closure can be returned from another function.

9. Closures are heavily used with callbacks, handlers, middleware, and function factories.

10. Don't make everything a closure. Use a closure when keeping local state or capturing configuration makes the design simpler.

---

# 🧠 What I Learned Today

* A closure is a function that remembers variables from its surrounding scope.
* A closure can access captured variables even after the outer function has returned.
* A closure can modify captured variables.
* `counter()` is a classic example of a closure.
* Each call to `counter()` creates separate state.
* Closures can provide private/local state.
* A function factory can create different closures with different captured values.
* Anonymous functions can become closures when they capture outer variables.
* Closures are useful for counters, callbacks, middleware, handlers, and state management.

---

# 📌 Quick Reference

### Basic Closure

```go
func main() {
    count := 0

    increment := func() {
        count++
    }

    increment()
    increment()

    fmt.Println(count)
}
```

Output:

```text
2
```

### Closure Returned From Function

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

### Use

```go
increment := counter()

fmt.Println(increment()) // 1
fmt.Println(increment()) // 2
fmt.Println(increment()) // 3
```

### Separate State

```go
counter1 := counter()
counter2 := counter()

fmt.Println(counter1()) // 1
fmt.Println(counter1()) // 2

fmt.Println(counter2()) // 1
```

> **Key idea:** A closure allows a function to remember and use variables from the scope where the function was created.

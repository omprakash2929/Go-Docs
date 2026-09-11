# 📦 Go Values & Variables

## 1. Simple Values

Go allows us to directly print different types of values.

```go
package main

import "fmt"

func main() {
    // Integer
    fmt.Println(1)

    // String
    fmt.Println("Hello")

    // Boolean
    fmt.Println(true)
}
```

### Common Values

```text
1          → Integer
"Hello"    → String
true       → Boolean
```

---

# 📌 Variables

A variable is a named location used to store a value.

Go provides **multiple ways to define variables**, which gives developers flexibility depending on the situation.

---

## 2. Explicit Variable Declaration

We can define the **type and value** explicitly.

```go
var name string = "Omprakash"
```

Here:

```text
var    → declares a variable
name   → variable name
string → data type
"Omprakash" → value
```

---

## 3. Type Inference

Go can automatically determine the type from the value.

```go
var name = "Omprakash"
var num = 23
var isAdult = true
```

Go understands:

```text
name    → string
num     → int
isAdult → bool
```

So we don't always need to write the type.

---

## 4. Short Variable Declaration

Inside a function, we can use the `:=` operator.

```go
name := "Golang"
```

Go automatically determines the type.

```go
num := 23
isAdult := true
```

This is one of the most commonly used ways to create local variables.

> ⚠️ `:=` can only be used inside functions.

---

## 5. Declare First, Assign Later

We can declare a variable without assigning a value immediately.

```go
var name string

name = "Chauhan"
```

The variable gets its **zero value** initially.

For example:

```go
var name string  // ""
var age int      // 0
var active bool  // false
```

---

# 🌍 Package-Level Variable

Variables can also be declared outside a function.

```go
package main

import "fmt"

var language string = "JavaScript"

func main() {
    fmt.Println(language)
}
```

A variable declared outside a function belongs to the package scope.

---

# 🔒 Constants

A constant is a value that **cannot be changed after it is declared**.

```go
const age = 22
```

For example:

```go
const language = "Golang"
```

You cannot assign a new value to it later.

---

# 🧠 What I Learned Today

* Go supports different types of values such as `int`, `string`, and `bool`.
* Variables can be declared in multiple ways.
* We can explicitly specify the type.
* Go can automatically infer the type.
* `:=` provides a short way to declare and initialize variables.
* Variables can be declared first and assigned later.
* Variables declared outside functions have package-level scope.
* `const` is used for values that should not change.
* Go provides flexibility while keeping the syntax simple.

---

# ⚡ Quick Reference

| Method              | Example                           | Where?          |
| ------------------- | --------------------------------- | --------------- |
| Explicit            | `var name string = "Go"`          | Anywhere        |
| Type inference      | `var name = "Go"`                 | Anywhere        |
| Short declaration   | `name := "Go"`                    | Inside function |
| Declare then assign | `var name string` → `name = "Go"` | Anywhere        |
| Constant            | `const age = 22`                  | Anywhere        |

> **Key takeaway:** Go gives you multiple ways to declare values and variables, so you can choose between explicit, readable declarations and concise syntax depending on the situation.

# 🔀 Go Conditional Statements

Conditional statements allow a program to make decisions based on a condition.

Go mainly provides:

* `if`
* `else if`
* `else`

---

# 1. Basic `if`

The `if` statement runs code when a condition is `true`.

```go
age := 21

if age >= 18 {
    fmt.Println("You are an adult")
}
```

If `age >= 18` is `true`, the message is printed.

---

# 2. `if` + `else if` + `else`

We can check multiple conditions.

```go
age := 15

if age >= 12 && age <= 19 {
    fmt.Println("It's a teenager")
} else if age >= 20 {
    fmt.Println("It's an adult")
} else {
    fmt.Println("It's a kid")
}
```

### Flow

```text
        age?
          │
    ┌─────┴─────┐
  12-19        >=20
    │             │
Teenager        Adult
    │
   else
    │
   Kid
```

---

# 3. Logical Operators

Go provides logical operators to combine conditions.

## `&&` — AND

Both conditions must be `true`.

```go
if age >= 12 && age <= 19 {
    fmt.Println("Teenager")
}
```

Meaning:

```text
age >= 12  AND  age <= 19
```

---

## `||` — OR

At least one condition must be `true`.

```go
if role == "admin" || hasPermission == true {
    fmt.Println("You have access")
}
```

Meaning:

```text
role is admin
OR
hasPermission is true
```

---

## `!` — NOT

`!` reverses a boolean value.

```go
isAdmin := false

if !isAdmin {
    fmt.Println("User is not an admin")
}
```

```text
!true  → false
!false → true
```

---

# 4. Combining Conditions

We can combine multiple logical operators.

```go
role := "admin"
hasPermission := true

if role == "admin" && hasPermission {
    fmt.Println("You can access the admin panel")
}
```

### Tip

You don't need to write:

```go
hasPermission == true
```

This is cleaner:

```go
if hasPermission {
    // code
}
```

---

# 5. Short Variable Declaration in `if`

Go allows us to declare a variable directly inside an `if` statement.

```go
if age := 21; age >= 19 {
    fmt.Println("You are an adult")
} else if age <= 19 {
    fmt.Println("You are a teenager")
}
```

Here:

```go
age := 21
```

is declared as part of the `if` statement.

The variable is available only inside the `if`, `else if`, and `else` blocks.

For example:

```go
if age := 21; age >= 18 {
    fmt.Println(age)
}

// fmt.Println(age) ❌
```

`age` cannot be accessed outside the `if` statement.

This is called **block scope**.

---

# 6. Important Go Syntax Rule

In Go, the opening `{` must be on the same line as the condition.

Correct:

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

Not:

```go
if age >= 18
{
    fmt.Println("Adult")
}
```

---

# 7. Comparison Operators

Conditions commonly use comparison operators.

| Operator | Meaning               |
| -------- | --------------------- |
| `==`     | Equal                 |
| `!=`     | Not equal             |
| `>`      | Greater than          |
| `<`      | Less than             |
| `>=`     | Greater than or equal |
| `<=`     | Less than or equal    |

Example:

```go
age := 22

if age >= 18 {
    fmt.Println("Adult")
}
```

---

# 🧠 What I Learned Today

* `if` statements
* `else if`
* `else`
* Comparison operators
* `&&` (AND)
* `||` (OR)
* `!` (NOT)
* Combining multiple conditions
* Short variable declaration inside `if`
* Variable scope inside `if`
* Go conditional syntax

---

# ⚡ Quick Reference

```go
// Basic if
if condition {
    // code
}

// if + else
if condition {
    // code
} else {
    // code
}

// Multiple conditions
if condition1 {
    // code
} else if condition2 {
    // code
} else {
    // code
}

// Short variable declaration
if value := 10; value > 5 {
    // code
}
```

> **Key takeaway:** Go keeps conditional logic simple. We use `if`, `else if`, and `else`, together with comparison and logical operators, to make decisions in our programs.

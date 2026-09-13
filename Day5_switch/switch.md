# 🔀 Go Switch Statement

A `switch` statement is used when we want to compare a value with multiple possible cases.

It can make code cleaner than writing many `if else if` statements.

---

# 1. Simple Switch

```go
package main

import "fmt"

func main() {
    num := 3

    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Not supported")
    }
}
```

Output:

```text
Three
```

### Important

In Go, we **do not need to write `break`** after every case.

Go automatically stops after the matching case.

For example:

```go
case 3:
    fmt.Println("Three")
```

Once this case matches, the switch ends automatically.

---

# 2. `default`

`default` runs when none of the cases match.

```go
num := 10

switch num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Not supported")
}
```

Output:

```text
Not supported
```

`default` is similar to the final `else` in an `if-else` statement.

---

# 3. Multiple Values in One Case

We can put multiple values in the same `case`.

```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("It's weekend")
default:
    fmt.Println("It's a work day")
}
```

Both `"Saturday"` and `"Sunday"` execute the same code.

This is useful when multiple values should produce the same result.

---

# 4. Switch Without an Expression

Go also allows a switch without a value.

```go
age := 22

switch {
case age < 13:
    fmt.Println("Kid")
case age < 20:
    fmt.Println("Teenager")
case age >= 20:
    fmt.Println("Adult")
}
```

This works similar to an `if-else if` chain.

It can make multiple conditions easier to read.

---

# 5. Type Switch

A **type switch** is used to check the type of a value.

Example:

```go
whoAmI := func(i interface{}) {
    switch k := i.(type) {
    case int:
        fmt.Println("It's integer", k)
    case string:
        fmt.Println("It's string", k)
    default:
        fmt.Println("Not supported", k)
    }
}

whoAmI(true)
```

Output:

```text
Not supported true
```

Here:

```go
i interface{}
```

allows the function to receive different types of values.

Then:

```go
switch k := i.(type)
```

checks the actual type of `i`.

---

# 6. Type Switch Examples

```go
whoAmI(10)
```

Output:

```text
It's integer 10
```

```go
whoAmI("Golang")
```

Output:

```text
It's string Golang
```

```go
whoAmI(true)
```

Output:

```text
Not supported true
```

---

# 7. Why Use Type Switch?

Type switches are useful when a function can receive values of different types and we need different behavior for each type.

Example:

```go
switch k := i.(type) {
case int:
    // Handle integer
case string:
    // Handle string
case bool:
    // Handle boolean
default:
    // Handle unknown type
}
```

---

# 🧠 Important Points

* `switch` is used to handle multiple possible cases.
* Go automatically stops after a matching case.
* We normally don't need `break`.
* `default` handles unmatched values.
* Multiple values can be used in one case.
* A switch can be used without an expression.
* A switch without an expression can replace some `if-else` chains.
* A **type switch** checks the type of a value.
* `i.(type)` is used specifically inside a type switch.
* `interface{}` can hold values of different types.

---

# ⚡ Quick Reference

### Normal Switch

```go
switch value {
case 1:
    // code
case 2:
    // code
default:
    // code
}
```

### Multiple Values

```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Work day")
}
```

### Condition Switch

```go
switch {
case age < 18:
    fmt.Println("Minor")
case age >= 18:
    fmt.Println("Adult")
}
```

### Type Switch

```go
switch value := data.(type) {
case int:
    fmt.Println("Integer", value)
case string:
    fmt.Println("String", value)
default:
    fmt.Println("Unknown type")
}
```

---

# 📝 What I Learned Today

* Simple `switch`
* `case`
* `default`
* Multiple conditions in one case
* Switch without an expression
* Automatic `break` behavior
* Type switch
* `interface{}`
* `i.(type)`

> **Key takeaway:** Go's `switch` is simple and powerful. It removes the need for manual `break` statements and can also be used to check conditions or the type of a value.

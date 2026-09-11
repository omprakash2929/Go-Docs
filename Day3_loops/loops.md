# 🔄 Go Loops

## Introduction

In Go, there is **only one looping keyword: `for`**.

Go does not have separate `while` or `do-while` keywords.

The `for` loop can be used in different ways to create:

* While-style loops
* Infinite loops
* Classic for loops
* Range loops

---

# 1. While-Style Loop

Go does not have a `while` keyword.

Instead, we can use `for` with a condition.

```go
i := 1

for i <= 9 {
    fmt.Println(i)
    i = i + 1
}
```

Output:

```text
1
2
3
4
5
6
7
8
9
```

The loop continues while the condition is `true`.

### Important

We must update `i` inside the loop.

```go
i = i + 1
```

Otherwise, the loop can become infinite.

---

# 2. Infinite Loop

A `for` loop without a condition runs forever.

```go
for {
    fmt.Println("Hello")
}
```

This is called an **infinite loop**.

It is useful when a program needs to keep running, such as:

* Servers
* Background workers
* Event loops
* Continuous monitoring

Usually, we use `break` to stop an infinite loop.

```go
for {
    if condition {
        break
    }
}
```

---

# 3. Classic For Loop

Go also supports the traditional three-part `for` loop.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

It has three parts:

```text
initialization ; condition ; update
```

Example:

```go
for i := 0; i < 5; i++ {
    // code
}
```

### Flow

```text
i := 0
   ↓
check condition
   ↓
run code
   ↓
i++
   ↓
check condition again
```

Output:

```text
0
1
2
3
4
```

---

# 4. `break`

`break` immediately stops the loop.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }

    fmt.Println(i)
}
```

Output:

```text
0
1
2
3
4
```

When `i` becomes `5`, the loop stops.

---

# 5. `continue`

`continue` skips the current iteration and moves to the next iteration.

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        continue
    }

    fmt.Println(i)
}
```

Output:

```text
0
1
2
4
```

`3` is skipped.

### Difference

```text
break     → Stop the entire loop
continue  → Skip current iteration
```

---

# 6. Range

Go provides `range` to iterate over values.

For example:

```go
for i := range 9 {
    fmt.Println(i)
}
```

Output:

```text
0
1
2
3
4
5
6
7
8
```

`range 9` generates values from `0` to `8`.

---

## Range with Slice

`range` is commonly used with slices.

```go
names := []string{"Om", "Prakash", "Chauhan"}

for index, name := range names {
    fmt.Println(index, name)
}
```

Output:

```text
0 Om
1 Prakash
2 Chauhan
```

Here:

```text
index → position
name  → value
```

---

## Ignore the Index

If we only need the value, use `_`.

```go
names := []string{"Om", "Prakash", "Chauhan"}

for _, name := range names {
    fmt.Println(name)
}
```

Output:

```text
Om
Prakash
Chauhan
```

`_` means:

> I don't need this value.

---

# 7. Nested Loops

A loop can be placed inside another loop.

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Println(i, j)
    }
}
```

This is called a **nested loop**.

Nested loops are useful for:

* Matrices
* Tables
* 2D data
* Comparing multiple values

---

# 8. Looping Over a Map

`range` can also be used with maps.

```go
users := map[string]int{
    "Om":     22,
    "Rahul":  25,
}

for name, age := range users {
    fmt.Println(name, age)
}
```

Here:

```text
name → key
age  → value
```

---

# 🧠 Important Points

* Go has only one looping keyword: `for`.
* Go does not have a `while` keyword.
* `for condition {}` can work like a while loop.
* `for {}` creates an infinite loop.
* `for init; condition; update {}` is the classic form.
* `break` stops the loop.
* `continue` skips the current iteration.
* `range` is useful for iterating over collections.
* `range` can be used with slices, arrays, maps, strings, and other iterable values.
* `_` can be used when a returned value is not needed.
* Loops can be nested inside other loops.

---

# 📝 What I Learned Today

Today I learned:

* `for` loop
* While-style loop
* Infinite loop
* Classic `for` loop
* `break`
* `continue`
* `range`
* Looping through slices
* Looping through maps
* Nested loops

---

# ⚡ Quick Reference

| Syntax                      | Use                              |
| --------------------------- | -------------------------------- |
| `for condition {}`          | While-style loop                 |
| `for {}`                    | Infinite loop                    |
| `for i := 0; i < 5; i++ {}` | Classic loop                     |
| `break`                     | Stop loop                        |
| `continue`                  | Skip iteration                   |
| `for i := range 9 {}`       | Range over integers              |
| `for i, v := range data {}` | Iterate with index/key and value |
| `for _, v := range data {}` | Iterate only values              |

> **Key takeaway:** Go keeps looping simple. Instead of having `while`, `do-while`, and `for`, Go mainly uses one powerful `for` keyword.

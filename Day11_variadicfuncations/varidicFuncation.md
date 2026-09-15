# 🔢 Go Variadic Functions

## 1. What is a Variadic Function?

A **variadic function** is a function that can accept **zero or more arguments** of the same type.

Example:

```go
func sum(nums ...int) int {
    total := 0

    for _, num := range nums {
        total = total + num
    }

    return total
}
```

Here:

```go
nums ...int
```

means:

> The function can receive any number of `int` arguments.

For example:

```go
sum()
sum(1)
sum(1, 2)
sum(1, 2, 3, 4, 5)
```

All of these are valid.

---

# 2. Calling a Variadic Function

Example:

```go
func sum(nums ...int) int {
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
```

Call it:

```go
res := sum(1, 2, 3, 4, 5)

fmt.Println(res)
```

Output:

```text
15
```

The number of arguments is flexible.

---

# 3. How `...int` Works

When we write:

```go
func sum(nums ...int) int
```

inside the function, `nums` behaves like a **slice of `int`**.

Conceptually:

```text
sum(1, 2, 3, 4, 5)

        ↓

nums = []int{1, 2, 3, 4, 5}
```

That's why we can use:

```go
for _, num := range nums {
    ...
}
```

We can also use:

```go
len(nums)
```

and:

```go
nums[0]
```

because `nums` is a slice inside the function.

---

# 4. Using `range` With Variadic Parameters

Your code:

```go
func sum(nums ...int) int {

    total := 0

    for _, num := range nums {
        total = total + num
    }

    return total
}
```

The loop:

```go
for _, num := range nums
```

means:

```text
_   → ignore index
num → current value
```

Example:

```text
nums = [1, 2, 3, 4]

range:

index   value
  0       1
  1       2
  2       3
  3       4
```

Then:

```text
total = 0
total = 0 + 1
total = 1 + 2
total = 3 + 3
total = 6 + 4
total = 10
```

---

# 5. Passing a Slice to a Variadic Function

This is an important part of your example.

You have:

```go
n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
```

You cannot pass the slice directly like this:

```go
sum(n) // ❌
```

because `sum()` expects individual `int` arguments.

Instead, use:

```go
sum(n...)
```

The `...` **expands the slice elements into individual arguments**.

Conceptually:

```go
sum(n...)
```

becomes:

```go
sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
```

So:

```go
fmt.Println(sum(n...))
```

outputs:

```text
45
```

---

# 6. Slice Expansion

The `...` operator can be used when passing a slice to a variadic parameter.

Example:

```go
nums := []int{10, 20, 30}

result := sum(nums...)
```

Conceptually:

```text
nums...
  ↓
10, 20, 30
  ↓
sum(10, 20, 30)
```

This is called **slice expansion** or **unpacking**.

---

# 7. Zero Arguments Are Allowed

A variadic function can be called without arguments.

```go
fmt.Println(sum())
```

Output:

```text
0
```

Why?

```go
total := 0
```

and there are no numbers to add.

---

# 8. One Argument

```go
fmt.Println(sum(10))
```

Output:

```text
10
```

---

# 9. Multiple Arguments

```go
fmt.Println(sum(10, 20, 30))
```

Output:

```text
60
```

---

# 10. Mixing Normal and Variadic Parameters

A variadic parameter does not have to be the only parameter.

Example:

```go
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Println(greeting, name)
    }
}
```

Call:

```go
greet("Hello", "Om", "Rahul", "Amit")
```

Output:

```text
Hello Om
Hello Rahul
Hello Amit
```

Here:

```text
greeting → normal parameter
names    → variadic parameter
```

### Important Rule

The variadic parameter must be the **last parameter**.

Correct:

```go
func example(a int, b ...int)
```

Incorrect:

```go
func example(a ...int, b int) // ❌
```

---

# 11. Variadic Function With Different Types

A variadic parameter has one fixed type.

For example:

```go
func sum(nums ...int)
```

accepts only `int` values.

```go
sum(1, 2, 3) // ✅
```

But:

```go
sum(1, "hello", 3) // ❌
```

is not valid.

If you need different types, you can use `any`, but that should be used only when there is a real reason.

Example:

```go
func printAll(values ...any) {
    for _, value := range values {
        fmt.Println(value)
    }
}
```

Now:

```go
printAll(10, "Go", true)
```

is valid.

---

# 12. Variadic Functions and Slices

Inside the function:

```go
func sum(nums ...int)
```

`nums` is treated as a slice:

```go
[]int
```

Therefore you can use:

```go
len(nums)
```

```go
nums[0]
```

```go
for _, num := range nums
```

Example:

```go
func count(nums ...int) int {
    return len(nums)
}
```

Call:

```go
fmt.Println(count(10, 20, 30))
```

Output:

```text
3
```

---

# 13. `...` Has Two Related Uses

You will see `...` in two places.

### 1. Function declaration

```go
func sum(nums ...int) int
```

This means:

> Accept zero or more `int` arguments.

### 2. Function call

```go
sum(nums...)
```

This means:

> Expand this slice into individual arguments.

So remember:

```text
Declaration:
...int → collect multiple arguments

Call:
slice... → expand slice into arguments
```

---

# 14. Variadic Functions Are Useful For

Variadic functions are useful when the number of arguments is not fixed.

Common examples:

* `fmt.Println()`
* Logging functions
* Mathematical operations
* Utility functions
* Configuration helpers
* APIs where multiple values can be supplied

For example:

```go
fmt.Println("Go", "Python", "JavaScript")
```

`fmt.Println` itself accepts a variable number of arguments.

---

# 15. Example: Find Maximum Number

A variadic function can be used to find the maximum number:

```go
func max(nums ...int) int {
    if len(nums) == 0 {
        return 0
    }

    maximum := nums[0]

    for _, num := range nums {
        if num > maximum {
            maximum = num
        }
    }

    return maximum
}
```

Call:

```go
fmt.Println(max(10, 50, 20, 30))
```

Output:

```text
50
```

---

# 16. Example: Average

```go
func average(nums ...int) float64 {
    if len(nums) == 0 {
        return 0
    }

    total := 0

    for _, num := range nums {
        total += num
    }

    return float64(total) / float64(len(nums))
}
```

Call:

```go
fmt.Println(average(10, 20, 30))
```

Output:

```text
20
```

---

# ⭐ Pro Tips

1. `...int` means **zero or more `int` arguments**.

2. Inside the function, a variadic parameter behaves like a slice:

```go
nums ...int
```

→ `nums` behaves like `[]int`.

3. You can use `range`, `len`, and indexing with it.

4. To pass a slice:

```go
sum(nums...)
```

not:

```go
sum(nums)
```

5. The `...` in a function declaration and function call have related but different meanings.

6. A variadic parameter must be the **last parameter**.

7. Variadic functions can accept zero arguments:

```go
sum()
```

8. Variadic parameters have one fixed type:

```go
func sum(nums ...int)
```

accepts only `int` values.

9. Don't use variadic parameters just because you can. Use them when the function naturally supports a variable number of arguments.

---

# 🧠 What I Learned Today

* A variadic function accepts a variable number of arguments.
* `...int` allows zero or more `int` arguments.
* Inside the function, the variadic parameter behaves like a slice.
* `range` can be used to iterate over the values.
* A slice can be expanded using `...`.
* `sum(nums...)` passes each slice element as an individual argument.
* A variadic parameter must be the last parameter.
* Variadic functions can accept zero, one, or many arguments.
* Variadic functions are useful when the number of inputs is not fixed.

---

# 📌 Quick Reference

```go
// Variadic function
func sum(nums ...int) int {
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
```

### Multiple arguments

```go
sum(1, 2, 3, 4)
```

### No arguments

```go
sum()
```

### Slice

```go
nums := []int{1, 2, 3, 4}
```

### Expand slice

```go
sum(nums...)
```

### Normal + variadic parameter

```go
func greet(message string, names ...string) {
    for _, name := range names {
        fmt.Println(message, name)
    }
}
```

> **Key idea:** `...` in a function parameter collects multiple arguments, while `...` when calling the function expands a slice into individual arguments.

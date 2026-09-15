# 🔄 Go Range — Iterating Over Data Structures

## 1. What is `range`?

`range` is used with `for` loops to iterate over data structures.

It makes it easy to access elements one by one.

Commonly used with:

* Slices
* Arrays
* Maps
* Strings
* Channels

Basic syntax:

```go
for index, value := range collection {
    // code
}
```

The values returned by `range` depend on the data structure.

---

# 2. Range Over Slice

Example:

```go
nums := []int{6, 7, 8, 9}

for i, num := range nums {
    fmt.Println(i, num)
}
```

Output:

```text
0 6
1 7
2 8
3 9
```

Here:

* `i` → index
* `num` → value

So:

```text
i     num
0      6
1      7
2      8
3      9
```

### Using Normal `for` Loop

We can also iterate using a traditional loop:

```go
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}
```

### Using `range`

```go
for i, num := range nums {
    fmt.Println(i, num)
}
```

`range` is usually cleaner when we need both index and value.

---

# 3. Calculate Sum Using Range

Example:

```go
nums := []int{6, 7, 8, 9}

sum := 0

for _, num := range nums {
    sum = sum + num
}

fmt.Println(sum)
```

Output:

```text
30
```

Here `_` is used because we don't need the index.

```go
for _, num := range nums
```

Means:

```text
ignore index → use value
```

### Important

If you don't need a returned value, use `_`.

```go
for _, value := range nums {
    fmt.Println(value)
}
```

---

# 4. Range Over Maps

Example:

```go
m := map[string]string{
    "fname": "omprakash",
    "lname": "chauhan",
    "age":   "21",
}

for k, v := range m {
    fmt.Println(k, v)
}
```

Here:

* `k` → key
* `v` → value

Example output could be:

```text
fname omprakash
lname chauhan
age 21
```

But **map order is not guaranteed**.

You might get:

```text
age 21
fname omprakash
lname chauhan
```

or another order.

### Important Rule

Never depend on map iteration order.

---

# 5. Get Only Keys

If we only need keys:

```go
for k := range m {
    fmt.Println(k)
}
```

Example output:

```text
fname
lname
age
```

We don't need to write `_` for the value because:

```go
for k := range m
```

directly gives the key.

---

# 6. Get Only Values

For values only:

```go
for _, v := range m {
    fmt.Println(v)
}
```

Here:

```text
_ → ignore key
v → value
```

---

# 7. Range Over String

Strings are slightly different.

Example:

```go
for i, c := range "Omprakash" {
    fmt.Println(i, string(c))
}
```

Output:

```text
0 O
1 m
2 p
3 r
4 a
5 k
6 a
7 s
8 h
```

Here:

* `i` → starting byte position
* `c` → Unicode code point (`rune`)

`c` is a `rune`, which is an alias for `int32`.

We use:

```go
string(c)
```

to convert the rune into a string.

---

# 8. Byte Index vs Character Index

This is an important concept.

In Go, strings are stored as **UTF-8 encoded bytes**.

For ASCII characters:

```go
for i, c := range "Omprakash" {
    fmt.Println(i, string(c))
}
```

The byte position and character position appear the same because ASCII characters use one byte.

But Unicode characters can use multiple bytes.

Example:

```go
for i, c := range "Go 🚀" {
    fmt.Println(i, string(c))
}
```

The `i` value is the **byte position**, not simply the character number.

So for strings:

> `range` gives the starting byte index of each decoded UTF-8 rune.

This is why `range` is preferred when working with Unicode text.

---

# 9. Range Over Array

`range` also works with arrays.

```go
nums := [4]int{10, 20, 30, 40}

for i, num := range nums {
    fmt.Println(i, num)
}
```

Output:

```text
0 10
1 20
2 30
3 40
```

The behavior is similar to slices.

---

# 10. Range With Only Value

If index is not required:

```go
nums := []int{10, 20, 30}

for _, num := range nums {
    fmt.Println(num)
}
```

Output:

```text
10
20
30
```

This is one of the most common patterns in Go.

---

# 11. Range and Modifying Slice Values

Be careful when modifying values inside a `range` loop.

Example:

```go
nums := []int{1, 2, 3}

for _, num := range nums {
    num = num * 2
}

fmt.Println(nums)
```

Output:

```text
[1 2 3]
```

Why?

`num` is a copy of each slice element.

If you want to modify the original slice:

```go
for i := range nums {
    nums[i] = nums[i] * 2
}
```

Now:

```text
[2 4 6]
```

### Remember

```go
for _, value := range nums
```

→ `value` is a copy.

```go
for i := range nums
```

→ use `nums[i]` to modify the original element.

---

# 12. Quick Reference

| Data Structure | First Value    | Second Value |
| -------------- | -------------- | ------------ |
| Array          | Index          | Value        |
| Slice          | Index          | Value        |
| Map            | Key            | Value        |
| String         | Byte index     | Rune         |
| Channel        | Received value | —            |

Examples:

```go
// Slice
for i, v := range nums {}

// Map
for k, v := range m {}

// String
for i, r := range str {}

// Only values
for _, v := range nums {}
```

---

# 13. `range` vs Traditional `for`

### Traditional loop

```go
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}
```

### Range

```go
for _, num := range nums {
    fmt.Println(num)
}
```

### Use `range` when:

* You want to visit every element.
* You need index + value.
* You need map keys + values.
* You need to process string runes.

### Use traditional `for` when:

* You need custom index control.
* You don't want to iterate over every element.
* You need more control over the loop condition.

---

# ⭐ Pro Tips

1. `range` is used with `for`.

2. For slices and arrays:

```go
for i, value := range nums
```

`i` = index, `value` = element.

3. For maps:

```go
for key, value := range m
```

`key` = map key, `value` = map value.

4. Map iteration order is **not guaranteed**.

5. Use `_` when you don't need a value:

```go
for _, value := range nums
```

6. String `range` works with UTF-8 and gives **runes**, not raw bytes.

7. The string index returned by `range` is a **byte index**.

8. A `range` value is generally a copy of the element. To modify a slice element, use its index:

```go
for i := range nums {
    nums[i] = nums[i] * 2
}
```

---

# 🧠 What I Learned Today

* `range` is used to iterate over data structures.
* `range` works with arrays, slices, maps, strings, and channels.
* Slice/array → `index, value`
* Map → `key, value`
* String → `byte index, rune`
* `_` can ignore unwanted values.
* Map iteration order is not guaranteed.
* String `range` handles Unicode correctly.
* `range` values are copies, so modifying the loop variable doesn't modify the original slice element.

---

# 📌 Quick Syntax

```go
// Slice / Array
for i, value := range nums {
}

// Only value
for _, value := range nums {
}

// Only index
for i := range nums {
}

// Map
for key, value := range m {
}

// Only key
for key := range m {
}

// String
for i, r := range str {
}
```

> **Key idea:** `range` makes iterating over Go data structures simple and readable.

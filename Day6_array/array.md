# 📦 Go Arrays

An **array** is a collection of elements of the **same type** with a **fixed size**.

Example:

```go
var nums [4]int
```

This creates an array that can store exactly **4 integers**.

---

# 1. Creating an Array

We can declare an array first and assign values later.

```go
var num [4]int

num[0] = 1
num[1] = 2

fmt.Println(num)
```

Output:

```text
[1 2 0 0]
```

The values that are not assigned get their **zero value**.

For `int`, the zero value is:

```text
0
```

---

# 2. Array Index

Array indexing starts from **0**.

```go
nums := [4]int{10, 20, 30, 40}

fmt.Println(nums[0]) // 10
fmt.Println(nums[1]) // 20
fmt.Println(nums[3]) // 40
```

Index:

```text
Value:  10   20   30   40
Index:   0    1    2    3
```

Trying to access index `4` would cause a runtime error because the valid indexes are `0` to `3`.

---

# 3. String Array

Arrays can store strings too.

```go
var names [4]string

names[0] = "Golang"
names[1] = "Python"

fmt.Println(names)
```

Output:

```text
[Golang Python  ]
```

The remaining elements contain the zero value of `string`, which is an empty string.

---

# 4. Initialize an Array in One Line

We can create and initialize an array at the same time.

```go
nums := [3]int{1, 2, 3}

fmt.Println(nums)
```

Output:

```text
[1 2 3]
```

Here:

```text
[3]int → array of 3 integers
{1, 2, 3} → initial values
```

---

# 5. Array Length

Use `len()` to get the number of elements.

```go
nums := [3]int{1, 2, 3}

fmt.Println(len(nums))
```

Output:

```text
3
```

---

# 6. Array Capacity

You can use `cap()` with an array.

```go
nums := [3]int{1, 2, 3}

fmt.Println(cap(nums))
```

Output:

```text
3
```

For an array, the capacity is fixed and is the same as its length.

> `cap()` becomes much more important when we learn **slices**.

---

# <font color="#76923c"> 7. Let Go Count the Size</font>

We can use `...` when we don't want to manually specify the array length.

```go
nums := [...]int{1, 2, 3, 4, 5}

fmt.Println(nums)
fmt.Println(len(nums))
```

Go automatically calculates the size.

```text
[5]int
```

This is useful when we know the values but don't want to count them manually.

---

# 8. Partial Initialization

We don't have to provide a value for every index.

```go
nums := [5]int{1, 2}
```

Result:

```text
[1 2 0 0 0]
```

The remaining elements get their zero value.

---

# 9. Initialize Specific Indexes

We can directly assign values to specific indexes while creating the array.

```go
nums := [5]int{
    0: 10,
    3: 40,
}

fmt.Println(nums)
```

Output:

```text
[10 0 0 40 0]
```

This can be useful when only certain positions need initial values.

---

# 10. Two-Dimensional Array

Go supports multidimensional arrays.

```go
nums := [2][2]int{
    {1, 2},
    {4, 5},
}

fmt.Println(nums)
```

Output:

```text
[[1 2] [4 5]]
```

Think of it like a table:

```text
1  2
4  5
```

Access an element:

```go
fmt.Println(nums[0][1])
```

Output:

```text
2
```

---

# 11. Loop Through an Array

We can use a normal `for` loop.

```go
nums := [4]int{10, 20, 30, 40}

for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}
```

Or use `range`:

```go
for index, value := range nums {
    fmt.Println(index, value)
}
```

If we only need the values:

```go
for _, value := range nums {
    fmt.Println(value)
}
```

---

# 12. Arrays Are Fixed Size

The most important property of an array is that its size is fixed.

```go
nums := [3]int{1, 2, 3}
```

You cannot add another element to this array.

There is no:

```go
nums.append(4) // ❌
```

If you need a dynamic collection, Go provides **[[slices]]**.

```go
nums := []int{1, 2, 3}
```

Slices can grow using `append()`:

```go
nums = append(nums, 4)
```

> **Array = fixed size**  
> **Slice = dynamic size**

We will use slices much more often in real Go applications.

---

# 13. Arrays Are Value Types

Arrays are copied when assigned to another variable.

```go
a := [3]int{1, 2, 3}

b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[1 2 3]
[100 2 3]
```

Changing `b` does not change `a`.

This is an important difference compared with slices.

---

# 🚀 Why Use Arrays?

Arrays are useful when:

- The size is known in advance.
    
- The size should never change.
    
- You need predictable memory usage.
    
- You need fast index-based access.
    
- You are working with fixed-size data.
    

Array element access is **O(1)** because Go can directly access an element using its index.

```go
nums[3]
```

does not require searching through the previous elements.

---

# 💡 Pro Tips

### 1. Prefer slices for most application code

In real-world Go applications, you will usually work with:

```go
[]int
[]string
[]User
```

rather than fixed arrays.

---

### 2. Remember the difference

```go
[5]int    // Array → fixed size
[]int     // Slice → dynamic size
```

The number inside `[ ]` makes a big difference.

---

### 3. Array size is part of its type

These are different types:

```go
[3]int
[4]int
```

You cannot directly assign one to the other.

```go
var a [3]int
var b [4]int

// a = b ❌
```

---

### 4. Use `len()` instead of hardcoding the size

Instead of:

```go
for i := 0; i < 5; i++ {
```

Prefer:

```go
for i := 0; i < len(nums); i++ {
```

This makes your code safer if the array size changes.

---

# 🧠 What I Learned Today

- What an array is
    
- How to declare an array
    
- How to initialize an array
    
- Array indexing
    
- Zero values
    
- `len()`
    
- `cap()`
    
- `[...]` automatic size
    
- Partial initialization
    
- Specific index initialization
    
- Two-dimensional arrays
    
- Looping through arrays
    
- Fixed-size nature of arrays
    
- Difference between arrays and slices
    
- Arrays are value types
    
- Constant-time index access
    

---

# ⚡ Quick Reference

```go
// Declare
var nums [4]int

// Initialize
nums := [3]int{1, 2, 3}

// Let Go calculate size
nums := [...]int{1, 2, 3}

// Length
len(nums)

// Capacity
cap(nums)

// Access
nums[0]

// Update
nums[0] = 100

// Loop
for i, value := range nums {
    fmt.Println(i, value)
}

// 2D array
matrix := [2][2]int{
    {1, 2},
    {3, 4},
}
```

> **Key takeaway:** An array in Go is a fixed-size collection of values of the same type. It provides predictable memory usage and fast index access, but when you need a collection that can grow or shrink, use a **slice** instead.
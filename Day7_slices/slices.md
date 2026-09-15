# 🧩 Go Slices

## What is a Slice?

A **slice** is a dynamic, flexible view of an underlying array.

Unlike arrays, slices can **grow and shrink** as needed.

```go
nums := []int{1, 2, 3, 4, 5}
```

Slices are one of the **most commonly used data structures in Go**.

> **Array → Fixed size**
> **Slice → Dynamic size**

---

# 1. Nil Slice

A slice can be declared without assigning any value.

```go
var nums []int

fmt.Println(nums == nil)
fmt.Println(len(nums))
```

Output:

```text
true
0
```

A slice declared like this is a **nil slice**.

```go
var nums []int
```

It has:

```text
len = 0
cap = 0
nil = true
```

A nil slice is valid and can be used with `append()`.

```go
nums = append(nums, 10)

fmt.Println(nums)
```

Output:

```text
[10]
```

---

# 2. Creating a Slice with `make()`

We can create a slice using `make()`.

```go
nums := make([]int, 3)

fmt.Println(nums)
```

Output:

```text
[0 0 0]
```

Syntax:

```go
make([]type, length)
```

Example:

```go
make([]int, 3)
```

This creates:

```text
length = 3
capacity = 3
```

The elements initially contain their zero value.

---

# 3. `make()` with Length and Capacity

We can specify both length and capacity.

```go
nums := make([]int, 0, 5)

fmt.Println(len(nums))
fmt.Println(cap(nums))
```

Output:

```text
0
5
```

Here:

```text
length  = 0
capacity = 5
```

The slice has space for up to 5 elements before it needs to grow.

---

# 4. `len()` vs `cap()`

### `len()`

Returns the number of elements currently in the slice.

```go
nums := []int{1, 2, 3}

fmt.Println(len(nums))
```

Output:

```text
3
```

### `cap()`

Returns the capacity of the slice's underlying storage.

```go
fmt.Println(cap(nums))
```

Capacity can be greater than length.

Example:

```go
nums := make([]int, 0, 5)

len(nums) // 0
cap(nums) // 5
```

Think:

```text
Capacity: [ _  _  _  _  _ ]
Length:    0
```

After:

```go
nums = append(nums, 10)
```

```text
Capacity: [10  _  _  _  _ ]
Length:    1
```

---

# 5. `append()`

`append()` adds elements to a slice.

```go
nums := []int{1, 2, 3}

nums = append(nums, 4)

fmt.Println(nums)
```

Output:

```text
[1 2 3 4]
```

### Important

`append()` can return a **new slice**, so we normally assign it back:

```go
nums = append(nums, 4)
```

Not:

```go
append(nums, 4) // ❌
```

---

# 6. Slice Can Grow Automatically

```go
nums := make([]int, 0, 3)

nums = append(nums, 1)
nums = append(nums, 2)
nums = append(nums, 3)
nums = append(nums, 4)
```

When the capacity becomes full, Go can allocate a larger underlying array and move the elements.

You don't need to manually resize the slice.

---

# 7. Another Way to Create a Slice

We can directly initialize a slice.

```go
nums := []int{1, 2, 3, 4, 5}
```

Unlike an array, there is **no number inside `[]`**.

```go
[5]int{}   // Array
[]int{}    // Slice
```

---

# 8. Updating Slice Elements

We can access elements using indexes.

```go
nums := []int{1, 2, 3}

nums[0] = 100

fmt.Println(nums)
```

Output:

```text
[100 2 3]
```

Indexing starts from `0`.

---

# 9. Copying a Slice

Use the built-in `copy()` function.

```go
nums := []int{1, 2, 3}

nums2 := make([]int, len(nums))

copy(nums2, nums)

fmt.Println("Nums1:", nums)
fmt.Println("Nums2:", nums2)
```

Output:

```text
Nums1: [1 2 3]
Nums2: [1 2 3]
```

`copy()` copies elements from one slice to another.

### Important

The destination slice must have enough length.

```go
nums2 := make([]int, len(nums))
copy(nums2, nums)
```

---

# 10. Slice Operator

We can create a smaller slice from another slice.

```go
nums := []int{1, 2, 3, 4, 5}
```

### `nums[0:3]`

```go
fmt.Println(nums[0:3])
```

Output:

```text
[1 2 3]
```

The starting index is included, but the ending index is excluded.

```text
0   1   2   3   4
1   2   3   4   5
└───────┘
```

---

### `nums[:4]`

```go
fmt.Println(nums[:4])
```

Output:

```text
[1 2 3 4]
```

Means:

```go
nums[0:4]
```

---

### `nums[3:]`

```go
fmt.Println(nums[3:])
```

Output:

```text
[4 5]
```

Means:

```go
nums[3:len(nums)]
```

---

# 11. Important Slice Behavior

A slice is **not the actual array**.

A slice contains information that points to an underlying array.

Conceptually:

```text
Slice
 ├── Pointer → underlying array
 ├── Length
 └── Capacity
```

Example:

```go
nums := []int{1, 2, 3, 4}
```

Think:

```text
Slice
  │
  ▼
[1] [2] [3] [4]
```

This is why slices are lightweight and flexible.

---

# 12. Slices Can Share the Same Underlying Array

Example:

```go
nums := []int{1, 2, 3, 4, 5}

part := nums[1:4]

part[0] = 100

fmt.Println(nums)
fmt.Println(part)
```

Output:

```text
[1 100 3 4 5]
[100 3 4]
```

Changing `part` also changed `nums`.

Why?

Because both slices can refer to the **same underlying array**.

> ⚠️ This is very important when working with slices.

---

# 13. Append One Slice to Another

Use `...` to expand a slice when passing it to `append()`.

```go
num1 := []int{1, 2, 3}
num2 := []int{4, 5, 6}

num3 := append(num1, num2...)

fmt.Println(num3)
```

Output:

```text
[1 2 3 4 5 6]
```

Without `...`:

```go
append(num1, num2) // ❌
```

Because `num2` itself is a slice.

With:

```go
num2...
```

its elements are passed individually.

---

# 📦 `slices` Package

Go provides the standard library `slices` package with useful functions for working with slices.

```go
import "slices"
```

---

# 14. `slices.Max()`

Returns the largest value.

```go
nums := []int{1, 5, 3, 6, 2, 4}

fmt.Println(slices.Max(nums))
```

Output:

```text
6
```

---

# 15. `slices.Min()`

Returns the smallest value.

```go
fmt.Println(slices.Min(nums))
```

Output:

```text
1
```

---

# 16. `slices.Contains()`

Checks whether a value exists.

```go
nums := []int{1, 2, 3}

fmt.Println(slices.Contains(nums, 2))
```

Output:

```text
true
```

---

# 17. `slices.Index()`

Returns the index of a value.

```go
nums := []int{1, 5, 3, 6}

fmt.Println(slices.Index(nums, 5))
```

Output:

```text
1
```

If the value does not exist, it returns:

```text
-1
```

---

# 18. `slices.Equal()`

Checks whether two slices contain the same elements in the same order.

```go
num1 := []int{1, 2}
num2 := []int{1, 2}

fmt.Println(slices.Equal(num1, num2))
```

Output:

```text
true
```

Example:

```go
num1 := []int{1, 2}
num2 := []int{2, 1}

fmt.Println(slices.Equal(num1, num2))
```

Output:

```text
false
```

Order matters.

---

# 19. `slices.Sort()`

Sorts a slice.

```go
nums := []int{5, 2, 8, 1, 3}

slices.Sort(nums)

fmt.Println(nums)
```

Output:

```text
[1 2 3 5 8]
```

The slice is sorted **in place**.

---

# 20. Custom Sorting with `slices.SortFunc()`

We can define our own sorting logic.

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Om", 24},
    {"Raj", 21},
}

slices.SortFunc(people, func(a, b Person) int {
    return a.Age - b.Age
})
```

Now people are sorted by age.

```text
Raj → 21
Om  → 24
```

This is useful when sorting structs or custom data.

---

# 21. `slices.Reverse()`

Reverses the slice.

```go
nums := []int{1, 2, 3, 4}

slices.Reverse(nums)

fmt.Println(nums)
```

Output:

```text
[4 3 2 1]
```

It modifies the original slice **in place**.

---

# 22. `slices.Insert()`

Inserts values at a specific index.

```go
nums := []int{1, 2, 3}

nums = slices.Insert(nums, 1, 100)

fmt.Println(nums)
```

Output:

```text
[1 100 2 3]
```

Here:

```text
index = 1
value = 100
```

---

# 23. `slices.Delete()`

Deletes elements from a slice.

```go
nums := []int{1, 2, 3, 4, 5}

nums = slices.Delete(nums, 1, 3)

fmt.Println(nums)
```

Output:

```text
[1 4 5]
```

The range is:

```text
start = 1
end   = 3
```

The end index is excluded.

So indexes `1` and `2` are removed.

---

# 24. `slices.Clone()`

Creates a copy of a slice.

```go
nums := []int{1, 2, 3}

copyNums := slices.Clone(nums)
```

Now `copyNums` has its own underlying storage.

```go
copyNums[0] = 100

fmt.Println(nums)
fmt.Println(copyNums)
```

Output:

```text
[1 2 3]
[100 2 3]
```

This is useful when you want an independent copy.

---

# 25. `slices.Compact()`

Removes consecutive duplicate values.

```go
nums := []int{1, 1, 2, 2, 2, 3, 3}

nums = slices.Compact(nums)

fmt.Println(nums)
```

Output:

```text
[1 2 3]
```

### Important

`Compact()` removes **consecutive duplicates**.

It does not automatically remove every duplicate from any position.

Example:

```go
[1, 2, 1, 2]
```

will not become:

```text
[1, 2]
```

because the duplicates are not next to each other.

---

# 🚀 Pro Tips

## 1. Prefer slices over arrays for dynamic data

Most application code uses:

```go
[]string
[]int
[]User
```

instead of:

```go
[10]string
```

when the size can change.

---

## 2. Use `make()` when you know the expected size

If you know you will store many elements:

```go
nums := make([]int, 0, 100)
```

This gives the slice an initial capacity of `100` and can reduce allocations as it grows.

---

## 3. Don't confuse length and capacity

```go
nums := make([]int, 0, 5)
```

means:

```text
len = 0
cap = 5
```

You cannot directly do:

```go
nums[0] = 10 // ❌
```

because the length is `0`.

Instead:

```go
nums = append(nums, 10) // ✅
```

---

## 4. Nil slices are usually safe

This is valid:

```go
var nums []int

nums = append(nums, 10)
```

You don't always need to initialize a slice with `make()`.

---

## 5. Be careful when slicing

These can share the same underlying array:

```go
a := []int{1, 2, 3, 4}
b := a[1:3]
```

Changing `b` may affect `a`.

Use `slices.Clone()` when you need an independent copy.

---

## 6. Remember `append()` can change the underlying array

When a slice runs out of capacity, Go may allocate a new underlying array.

So don't rely on the underlying storage staying the same after `append()`.

---

# 🧠 Array vs Slice

| Feature           | Array       | Slice       |
| ----------------- | ----------- | ----------- |
| Size              | Fixed       | Dynamic     |
| Syntax            | `[3]int`    | `[]int`     |
| Can grow          | ❌           | ✅           |
| `append()`        | ❌           | ✅           |
| Length            | Fixed       | Can change  |
| Capacity          | Fixed       | Can change  |
| Common in Go apps | Less common | Very common |

---

# 📝 What I Learned Today

* What slices are
* Nil slices
* `make()`
* Length and capacity
* `append()`
* Creating slices
* Updating slice elements
* `copy()`
* Slice operators
* Slices and underlying arrays
* Shared underlying arrays
* Appending one slice to another
* `slices.Max()`
* `slices.Min()`
* `slices.Contains()`
* `slices.Index()`
* `slices.Equal()`
* `slices.Sort()`
* `slices.SortFunc()`
* `slices.Reverse()`
* `slices.Insert()`
* `slices.Delete()`
* `slices.Clone()`
* `slices.Compact()`

---

# ⚡ Quick Reference

```go
// Nil slice
var nums []int

// Create with make
nums := make([]int, 0, 5)

// Create with values
nums := []int{1, 2, 3}

// Add values
nums = append(nums, 4)

// Length
len(nums)

// Capacity
cap(nums)

// Slice
nums[1:3]

// Copy
copy(destination, source)

// Append another slice
nums = append(nums, other...)

// Contains
slices.Contains(nums, 2)

// Index
slices.Index(nums, 2)

// Sort
slices.Sort(nums)

// Reverse
slices.Reverse(nums)

// Clone
newNums := slices.Clone(nums)
```

> **Key takeaway:** A slice is Go's flexible and commonly used collection type. It provides dynamic size while using an underlying array for storage. Understanding **length, capacity, `append()`, and shared underlying arrays** is especially important before moving to more advanced Go programming.

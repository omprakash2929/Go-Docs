# 🗺️ Go Maps

## What is a Map?

A **map** is a collection of **key-value pairs**.

It is similar to:

* JavaScript → Object
* Python → Dictionary
* Other languages → Hash Map

Example:

```go
map[string]string
```

Here:

```text
string → key type
string → value type
```

Example:

```go
"name" → "Golang"
"area" → "backend"
"ext"  → ".go"
```

---

# 1. Creating a Map with `make()`

We can create a map using `make()`.

```go
m := make(map[string]string)
```

Then we can add elements:

```go
m["name"] = "Golang"
m["area"] = "backend"
m["ext"] = ".go"
```

Print the map:

```go
fmt.Println(m)
```

Output:

```text
map[area:backend ext:.go name:Golang]
```

> Map order is not guaranteed.

---

# 2. Getting a Value

We can access a value using its key.

```go
fmt.Println(m["name"])
```

Output:

```text
Golang
```

Example:

```go
m["area"]
```

returns:

```text
backend
```

---

# 3. What Happens If a Key Does Not Exist?

If the key does not exist, Go returns the **zero value** for the map's value type.

```go
m := make(map[string]int)

fmt.Println(m["age"])
```

Output:

```text
0
```

For a string:

```go
m := make(map[string]string)

fmt.Println(m["name"])
```

Output:

```text
""
```

For a boolean:

```go
map[string]bool
```

the zero value is:

```text
false
```

> ⚠️ This creates an important problem: `0`, `""`, or `false` could mean either "the key doesn't exist" or "the key exists with that value."

Use the **comma-ok pattern** to solve this.

---

# 4. Check Whether a Key Exists

Go provides a special way to check if a key exists.

```go
m := map[string]int{
    "id":    1,
    "price": 400,
}

value, ok := m["price"]

if ok {
    fmt.Println("Key exists:", value)
} else {
    fmt.Println("Key does not exist")
}
```

Here:

```text
value → value stored in the map
ok    → true if key exists
```

If the key exists:

```text
ok = true
```

If it doesn't:

```text
ok = false
```

---

# 5. Using `_` When You Don't Need the Value

Sometimes we only want to check whether a key exists.

We can ignore the value using `_`.

```go
_, ok := m["price"]

if ok {
    fmt.Println("Price exists")
}
```

`_` is called the **blank identifier**.

---

# 6. Add or Update an Element

The same syntax is used for both adding and updating.

### Add

```go
m["name"] = "Golang"
```

### Update

```go
m["name"] = "Python"
```

If `"name"` already exists, its value is replaced.

---

# 7. Delete an Element

Use the built-in `delete()` function.

```go
delete(m, "price")
```

Syntax:

```go
delete(map, key)
```

Example:

```go
m := map[string]int{
    "id":    1,
    "price": 400,
}

delete(m, "price")

fmt.Println(m)
```

Output:

```text
map[id:1]
```

If the key doesn't exist, `delete()` does nothing.

---

# 8. Clear a Map

Go provides `clear()` to remove all elements.

```go
clear(m)
```

After:

```go
clear(m)

fmt.Println(len(m))
```

Output:

```text
0
```

The map still exists, but it contains no elements.

---

# 9. Map Length

Use `len()` to get the number of key-value pairs.

```go
m := map[string]int{
    "id":    1,
    "price": 400,
    "phone": 3,
}

fmt.Println(len(m))
```

Output:

```text
3
```

---

# 10. Create a Map Without `make()`

We can directly initialize a map.

```go
m := map[string]int{
    "id":    1,
    "price": 400,
    "phone": 3,
}
```

This is very common when we already know the initial data.

---

# 11. Empty Map

We can create an empty, but initialized, map.

```go
m := map[string]string{}
```

Then we can add values:

```go
m["name"] = "Om"

fmt.Println(m)
```

This works because the map has been initialized.

---

# 12. Nil Map

This is different:

```go
var m map[string]string
```

This is a **nil map**.

You can safely read from it:

```go
fmt.Println(m["name"])
```

But you cannot add values to it.

```go
m["name"] = "Golang" // ❌ panic
```

To make it usable for writing:

```go
m = make(map[string]string)
```

Then:

```go
m["name"] = "Golang"
```

works.

### Remember

```go
var m map[string]int       // nil map
m := make(map[string]int)  // initialized map
m := map[string]int{}      // initialized empty map
```

---

# 13. Loop Through a Map

Use `range` to iterate over a map.

```go
m := map[string]int{
    "id":    1,
    "price": 400,
    "phone": 3,
}

for key, value := range m {
    fmt.Println(key, value)
}
```

Output order can be different each time.

> ⚠️ Never depend on map iteration order.

---

## Only Keys

If we only need the keys:

```go
for key := range m {
    fmt.Println(key)
}
```

---

## Only Values

If we only need the values:

```go
for _, value := range m {
    fmt.Println(value)
}
```

---

# 14. Maps Don't Guarantee Order

For example:

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}
```

You might get:

```text
a 1
b 2
c 3
```

But another run could produce:

```text
c 3
a 1
b 2
```

This is normal.

If you need ordered data, use a slice or sort the keys.

---

# 15. Map With Different Value Types

The key and value types can be different.

```go
m := map[string]int{
    "age":   22,
    "price": 500,
}
```

Here:

```text
key   → string
value → int
```

Another example:

```go
users := map[int]string{
    1: "Om",
    2: "Raj",
}
```

Here:

```text
key   → int
value → string
```

---

# 16. Nested Maps

A map can contain another map.

```go
users := map[string]map[string]string{
    "user1": {
        "name": "Om",
        "role": "admin",
    },
}
```

Access:

```go
fmt.Println(users["user1"]["name"])
```

Output:

```text
Om
```

Nested maps can be useful for structured data, but for complex data, structs are often easier to maintain.

---

# 📦 `maps` Package

Go also provides the standard library `maps` package.

```go
import "maps"
```

It provides useful functions for working with maps.

---

# 17. `maps.Equal()`

Checks whether two maps contain the same key-value pairs.

```go
m1 := map[string]int{
    "id":    1,
    "price": 400,
}

m2 := map[string]int{
    "id":    1,
    "price": 400,
}

fmt.Println(maps.Equal(m1, m2))
```

Output:

```text
true
```

Map order does not matter for equality.

These are equal:

```text
{id:1, price:400}
```

and:

```text
{price:400, id:1}
```

---

# 18. `maps.Clone()`

Creates a copy of a map.

```go
m2 := map[string]int{
    "id":    1,
    "price": 400,
}

m3 := maps.Clone(m2)
```

Now `m3` is a separate map.

```go
m3["price"] = 500

fmt.Println(m2["price"])
fmt.Println(m3["price"])
```

Output:

```text
400
500
```

For a map with simple values like `int` or `string`, this behaves like an independent copy.

> For maps containing reference-like values such as slices, maps, or pointers, `Clone()` is **shallow** with respect to those nested values.

---

# 19. `maps.Copy()`

`maps.Copy()` copies key-value pairs from one map into another.

```go
src := map[string]int{
    "a": 1,
    "b": 2,
}

dst := make(map[string]int)

maps.Copy(dst, src)

fmt.Println(dst)
```

Output:

```text
map[a:1 b:2]
```

Syntax:

```go
maps.Copy(destination, source)
```

If a key already exists in the destination, its value is overwritten.

---

# 🧠 Important Points

* A map stores **key-value pairs**.
* Keys must be of a comparable type.
* Values can be any type.
* Map size is dynamic.
* Use `make()` to initialize an empty map.
* A nil map can be read but cannot be written to.
* `len()` returns the number of entries.
* `delete()` removes an entry.
* `clear()` removes all entries.
* Missing keys return the value type's zero value.
* Use `value, ok := m[key]` to check key existence.
* Use `range` to iterate over a map.
* Map iteration order is not guaranteed.
* Maps are reference-like data structures; assigning a map variable to another map variable makes them refer to the same underlying map.
* Use `maps.Clone()` when you need a separate map.

---

# ⚡ Quick Reference

```go
// Create
m := make(map[string]int)

// Create with values
m := map[string]int{
    "age": 22,
}

// Add / Update
m["price"] = 500

// Get
fmt.Println(m["price"])

// Check key
value, ok := m["price"]

// Delete
delete(m, "price")

// Clear
clear(m)

// Length
len(m)

// Loop
for key, value := range m {
    fmt.Println(key, value)
}

// Only keys
for key := range m {
    fmt.Println(key)
}

// Only values
for _, value := range m {
    fmt.Println(value)
}

// Compare
maps.Equal(m1, m2)

// Clone
m2 := maps.Clone(m1)

// Copy
maps.Copy(destination, source)
```

---

# 🆚 Map vs Slice vs Array

| Feature   | Array      | Slice             | Map                   |
| --------- | ---------- | ----------------- | --------------------- |
| Structure | Indexed    | Indexed           | Key-value             |
| Size      | Fixed      | Dynamic           | Dynamic               |
| Access    | `arr[0]`   | `slice[0]`        | `map["key"]`          |
| Add       | ❌          | `append()`        | `map[key] = value`    |
| Main use  | Fixed data | Lists/collections | Lookup/key-value data |

---

# 🚀 Pro Tips

### 1. Use maps when you need fast lookup by a key

Instead of searching through a list:

```go
users["om"]
```

is a natural way to retrieve data by a unique key.

---

### 2. Always remember the nil map problem

This is safe:

```go
var m map[string]int

fmt.Println(m["age"])
```

But this causes a panic:

```go
m["age"] = 22 // ❌
```

Initialize it first:

```go
m = make(map[string]int)
```

---

### 3. Use comma-ok when the zero value matters

Don't rely only on:

```go
value := m["age"]
```

If `0` is a valid stored value, you cannot know whether the key exists.

Use:

```go
value, ok := m["age"]

if ok {
    fmt.Println(value)
}
```

---

### 4. Don't expect map order

Never write application logic that depends on:

```go
for key, value := range m {
```

being in a particular order.

---

### 5. Use structs for complex objects

Instead of making a map like:

```go
map[string]interface{}
```

for structured data, a `struct` is often better:

```go
type User struct {
    Name string
    Age  int
}
```

Structs give you better type safety and clearer code.

---

# 📝 What I Learned Today

* What maps are
* Key-value pairs
* Creating maps with `make()`
* Creating maps directly
* Adding and updating values
* Getting values
* Missing-key behavior
* `len()`
* `delete()`
* `clear()`
* Nil maps
* Checking key existence with comma-ok
* Blank identifier `_`
* Looping through maps with `range`
* Map iteration order
* Nested maps
* `maps.Equal()`
* `maps.Clone()`
* `maps.Copy()`

> **Key takeaway:** A map is Go's key-value data structure. Use it when you need to store and quickly access data using a key. The most important things to remember are **nil maps, comma-ok, and the fact that map order is not guaranteed**.

# 🧱 Go Structs, Constructors & Methods

## 1. What is a Struct?

A **struct** is a custom data type that groups multiple related values together.

For example, an order can have:

* ID
* Amount
* Status
* Created time

Instead of keeping separate variables:

```go
id := "001"
amount := 300.0
status := "shipped"
```

we can group them:

```go
type order struct {
    id        string
    amount    float32
    status    string
    createdAt time.Time
}
```

Now `order` represents one complete order.

---

# 2. Creating a Struct

We can create a struct using a struct literal:

```go
myOrder := order{
    id:     "001",
    amount: 200,
    status: "Done",
}
```

Now:

```go
fmt.Println(myOrder)
```

prints all the fields.

We can access individual fields using `.`:

```go
fmt.Println(myOrder.id)
fmt.Println(myOrder.amount)
fmt.Println(myOrder.status)
```

---

# 3. Struct Fields

Your struct:

```go
type order struct {
    id        string
    amount    float32
    status    string
    createdAt time.Time
}
```

has four fields:

| Field       | Type        |
| ----------- | ----------- |
| `id`        | `string`    |
| `amount`    | `float32`   |
| `status`    | `string`    |
| `createdAt` | `time.Time` |

A struct can contain different types.

---

# 4. Zero Values in Struct

If you don't provide a value for a field, Go gives it the **zero value** of its type.

Example:

```go
myOrder := order{
    id: "001",
}
```

The other fields get their zero values.

```text
id        → "001"
amount    → 0
status    → ""
createdAt → zero time.Time value
```

Common zero values:

| Type      | Zero Value |
| --------- | ---------- |
| `int`     | `0`        |
| `float32` | `0`        |
| `string`  | `""`       |
| `bool`    | `false`    |
| pointer   | `nil`      |
| slice     | `nil`      |
| map       | `nil`      |

---

# 5. Updating Struct Fields

Struct fields can be modified after creation.

```go
myOrder.status = "Confirmed"
```

And:

```go
myOrder.createdAt = time.Now()
```

Now the fields contain the new values.

---

# 6. `time.Time`

Your struct contains:

```go
createdAt time.Time
```

`time.Time` comes from the standard `time` package.

Import:

```go
import "time"
```

We can get the current time using:

```go
time.Now()
```

Example:

```go
myOrder.createdAt = time.Now()
```

`time.Time` stores date and time information with high precision.

---

# 7. Constructor Pattern in Go

Go does **not** have a special `constructor` keyword like some languages.

Instead, developers commonly create a normal function that initializes and returns a struct.

Your example:

```go
func newOrder(id string, amount float32, status string) *order {
    myOrder := order{
        id:     id,
        amount: amount,
        status: status,
    }

    return &myOrder
}
```

This function acts as a **constructor-like function**.

---

# 8. Why Name It `newOrder`?

Go has a common naming convention:

```go
newOrder()
newUser()
newServer()
newClient()
```

These functions are usually used to create and initialize values.

There is no special constructor syntax in Go.

---

# 9. Understanding `newOrder()`

When we call:

```go
myOrder := newOrder("001", 300, "shipped")
```

the arguments are:

```text
id     → "001"
amount → 300
status → "shipped"
```

Inside the function:

```go
myOrder := order{
    id:     id,
    amount: amount,
    status: status,
}
```

creates an `order`.

Then:

```go
return &myOrder
```

returns its address.

So the return type:

```go
*order
```

means:

> pointer to an `order`.

---

# 10. Why Return `*order`?

Your function:

```go
func newOrder(...) *order
```

returns a pointer to the struct.

So:

```go
myOrder := newOrder(...)
```

makes `myOrder` a pointer to an `order`.

Conceptually:

```text
myOrder
   │
   ▼
┌──────────────────┐
│ order            │
│ id: "001"        │
│ amount: 300      │
│ status: shipped  │
└──────────────────┘
```

Go automatically allows:

```go
myOrder.amount
```

even though `myOrder` is a pointer.

You don't normally need to write:

```go
(*myOrder).amount
```

Go handles the field access for you.

---

# 11. Methods

A **method** is a function associated with a specific type.

Your method:

```go
func (o *order) changeStatus(status string) {
    o.status = status
}
```

This is a method of the `order` type.

The part:

```go
(o *order)
```

is called the **receiver**.

---

# 12. Receiver

A receiver connects a method to a type.

Example:

```go
func (o *order) changeStatus(status string) {
    o.status = status
}
```

Here:

```text
o       → receiver variable
*order  → receiver type
```

So this method belongs to `order`.

We can call it:

```go
myOrder.changeStatus("Confirmed")
```

---

# 13. Pointer Receiver

Your method uses:

```go
(o *order)
```

This is called a **pointer receiver**.

Use a pointer receiver when the method needs to modify the original struct.

Example:

```go
func (o *order) changeStatus(status string) {
    o.status = status
}
```

Calling:

```go
myOrder.changeStatus("Confirmed")
```

changes the original order.

---

# 14. Why Pointer Receiver?

Suppose:

```go
myOrder.status = "Pending"
```

Then:

```go
myOrder.changeStatus("Confirmed")
```

The method changes:

```text
Pending
   ↓
Confirmed
```

Because `o` points to the original struct.

---

# 15. Value Receiver

Your second method:

```go
func (o order) getAmount() float32 {
    return o.amount
}
```

uses:

```go
(o order)
```

instead of:

```go
(o *order)
```

This is called a **value receiver**.

The method receives a copy of the struct value.

Since we're only reading:

```go
return o.amount
```

a value receiver is fine.

Call it:

```go
amount := myOrder.getAmount()

fmt.Println(amount)
```

---

# 16. Pointer Receiver vs Value Receiver

This is an important concept.

### Pointer Receiver

```go
func (o *order) changeStatus(status string) {
    o.status = status
}
```

Use when the method needs to modify the struct.

```text
Original struct
      ↑
      │
   pointer
      │
   method
      │
   modifies
```

### Value Receiver

```go
func (o order) getAmount() float32 {
    return o.amount
}
```

Use when the method only needs to read the value.

```text
Original struct
      │
      │ copy
      ▼
   method
      │
    reads
```

---

# 17. Method Call

Instead of calling a normal function like:

```go
changeStatus(myOrder, "Confirmed")
```

we can call the method:

```go
myOrder.changeStatus("Confirmed")
```

This makes the relationship clearer:

```text
myOrder
   │
   └── changeStatus()
```

---

# 18. Struct + Methods

Structs and methods work together very naturally.

Example:

```go
type order struct {
    id     string
    amount float32
    status string
}

func (o *order) changeStatus(status string) {
    o.status = status
}

func (o order) getAmount() float32 {
    return o.amount
}
```

Now `order` has behavior:

```text
order
 ├── id
 ├── amount
 ├── status
 │
 ├── changeStatus()
 └── getAmount()
```

This is one of the ways Go organizes code around data and behavior.

---

# 19. Anonymous Struct

Your code also contains:

```go
language := struct {
    name   string
    isGood bool
}{
    "golang",
    true,
}
```

This is an **anonymous struct**.

It means we create a struct without giving it a named type.

---

# 20. Named Struct vs Anonymous Struct

### Named Struct

```go
type language struct {
    name   string
    isGood bool
}
```

Now we can reuse it:

```go
lang1 := language{
    name:   "golang",
    isGood: true,
}
```

### Anonymous Struct

```go
language := struct {
    name   string
    isGood bool
}{
    "golang",
    true,
}
```

This struct type is created directly and is useful when we only need it in one place.

---

# 21. Named Fields vs Positional Values

You can initialize a struct using field names:

```go
order{
    id:     "001",
    amount: 300,
    status: "shipped",
}
```

This is generally clearer.

You can also use values in field order:

```go
order{
    "001",
    300,
    "shipped",
    time.Now(),
}
```

But field names are usually preferred because they make the code easier to understand and safer when the struct changes.

---

# 22. Complete Example

```go
package main

import (
    "fmt"
    "time"
)

type order struct {
    id        string
    amount    float32
    status    string
    createdAt time.Time
}

func newOrder(id string, amount float32, status string) *order {
    return &order{
        id:        id,
        amount:    amount,
        status:    status,
        createdAt: time.Now(),
    }
}

func (o *order) changeStatus(status string) {
    o.status = status
}

func (o order) getAmount() float32 {
    return o.amount
}

func main() {
    myOrder := newOrder("001", 300, "shipped")

    fmt.Println(myOrder.amount)

    myOrder.changeStatus("Delivered")

    fmt.Println(myOrder.status)
    fmt.Println(myOrder.getAmount())
    fmt.Println(myOrder.createdAt)
}
```

Possible output:

```text
300
Delivered
300
2026-09-17 ...
```

---

# 23. Struct + Pointer + Method Connection

Your example combines several concepts:

```text
                    order struct
                        │
             ┌──────────┼──────────┐
             ▼          ▼          ▼
            data      methods    constructor
             │          │          │
       id/amount/   receiver     newOrder()
       status       │
                    ▼
              *order / order
```

### Constructor

Creates the object/value:

```go
newOrder(...)
```

### Pointer receiver

Changes the original struct:

```go
func (o *order) changeStatus(...)
```

### Value receiver

Reads the struct:

```go
func (o order) getAmount()
```

---

# ⭐ Pro Tips

1. Go does not have a special `constructor` keyword.

2. Functions like `newOrder()` are commonly used as constructor-like functions.

3. `*order` means a pointer to an `order`.

4. Use a **pointer receiver** when a method needs to modify the struct.

5. A **value receiver** receives a copy of the struct.

6. Go automatically handles pointer field access:

```go
myOrder.amount
```

even when `myOrder` is a pointer.

7. Struct fields with no assigned value get their type's zero value.

8. Prefer named field initialization:

```go
order{
    id:     "001",
    amount: 300,
    status: "shipped",
}
```

9. `time.Time` comes from the standard `time` package.

10. Anonymous structs are useful for small, one-time data structures.

11. Keep pointer vs value receiver choices consistent for a type when possible. If a type has methods with pointer receivers, using pointer receivers consistently is often clearer.

12. Returning a pointer from a constructor-like function is common when the caller should work with the same struct instance and its methods may modify it.

---

# 🧠 What I Learned Today

* A struct groups related data together.
* Structs can contain different data types.
* Every struct field has a zero value if not initialized.
* Struct fields can be accessed using `.`.
* Struct fields can be modified after creation.
* Go does not have a special constructor syntax.
* `newOrder()` is a constructor-like function.
* `*order` means pointer to an `order`.
* A method is a function associated with a type.
* The variable before the method name is called the receiver.
* `*order` receiver is a pointer receiver.
* `order` receiver is a value receiver.
* Pointer receivers can modify the original struct.
* Value receivers work with a copy of the struct.
* Go automatically handles pointer field access.
* Anonymous structs allow creating a struct without defining a named type.
* `time.Time` is used for date and time values.

---

# 📌 Quick Reference

### Struct

```go
type order struct {
    id     string
    amount float32
    status string
}
```

### Create Struct

```go
myOrder := order{
    id:     "001",
    amount: 300,
    status: "shipped",
}
```

### Access Field

```go
fmt.Println(myOrder.amount)
```

### Modify Field

```go
myOrder.status = "Delivered"
```

### Constructor-like Function

```go
func newOrder(id string, amount float32, status string) *order {
    return &order{
        id:     id,
        amount: amount,
        status: status,
    }
}
```

### Pointer Receiver

```go
func (o *order) changeStatus(status string) {
    o.status = status
}
```

### Value Receiver

```go
func (o order) getAmount() float32 {
    return o.amount
}
```

### Anonymous Struct

```go
language := struct {
    name   string
    isGood bool
}{
    "golang",
    true,
}
```

> **Key idea:** In Go, structs hold data, methods define behavior, and receivers connect those methods to the struct type.

# 🧩 Go Struct Embedding & Composition

## 1. What is Struct Embedding?

**Struct embedding** means putting one struct inside another struct without giving it a separate field name.

Example:

```go
type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}
```

Here, `customer` is **embedded** inside `order`.

So an `order` contains:

* `id`
* `amount`
* `status`
* `createdAt`
* `customer` data

---

# 2. Composition in Go

Go commonly uses **composition** instead of traditional class inheritance.

Composition means:

> Build a larger type by combining smaller types.

For example:

```text
order
 ├── id
 ├── amount
 ├── status
 ├── createdAt
 └── customer
      ├── name
      └── phone
```

This creates a **has-a relationship**:

```text
Order has a Customer
```

It is not an `Order is a Customer` relationship.

---

# 3. Embedded Struct

In this example:

```go
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}
```

The line:

```go
customer
```

is an **embedded field**.

Normally, you could write:

```go
type order struct {
	id       string
	amount   float32
	status   string
	customer customer
}
```

But embedding:

```go
customer
```

gives additional behavior such as **field promotion**.

---

# 4. Creating an Embedded Struct

There are multiple ways to initialize the embedded `customer`.

## Method 1 — Create Customer Separately

```go
newCustomer := customer{
	name:  "omprakash",
	phone: "9510276296",
}

newOrder := order{
	id:       "001",
	amount:   300,
	status:   "Received",
	customer: newCustomer,
}
```

Here:

```go
customer: newCustomer
```

assigns the existing `customer` value to the embedded field.

---

# 5. Method 2 — Inline Initialization

You can create the embedded struct directly inside the `order`.

```go
newOrder := order{
	id:     "001",
	amount: 300,
	status: "Received",

	customer: customer{
		name:  "prakash",
		phone: "123456789",
	},
}
```

This is called **inline struct initialization**.

It is useful when you don't need the customer as a separate variable.

---

# 6. Accessing Embedded Fields

You can access the embedded struct normally:

```go
fmt.Println(newOrder.customer.name)
```

Output:

```text
prakash
```

You can also modify it:

```go
newOrder.customer.name = "Rhaul"
```

Now:

```go
fmt.Println(newOrder.customer.name)
```

Output:

```text
Rhaul
```

---

# 7. Field Promotion ⭐

One of the most useful features of embedding is **field promotion**.

Because `customer` is embedded, you can access its fields directly:

```go
fmt.Println(newOrder.name)
```

instead of:

```go
fmt.Println(newOrder.customer.name)
```

Similarly:

```go
fmt.Println(newOrder.phone)
```

works because `name` and `phone` are promoted from the embedded `customer`.

You can also modify them:

```go
newOrder.name = "Rahul"
```

This is effectively accessing:

```go
newOrder.customer.name
```

---

# 8. Embedded Field vs Normal Nested Field

### Normal nested struct

```go
type order struct {
	customer customer
}
```

Access:

```go
newOrder.customer.name
```

### Embedded struct

```go
type order struct {
	customer
}
```

You can access:

```go
newOrder.customer.name
```

and also:

```go
newOrder.name
```

The second form works because of **field promotion**.

---

# 9. Complete Example

```go
package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {

	newOrder := order{
		id:     "001",
		amount: 300,
		status: "Received",
		customer: customer{
			name:  "Prakash",
			phone: "123456789",
		},
	}

	// Access through embedded struct
	fmt.Println(newOrder.customer.name)

	// Access through field promotion
	fmt.Println(newOrder.name)

	// Modify promoted field
	newOrder.name = "Rahul"

	fmt.Println(newOrder)
}
```

---

# 10. Struct Embedding with Methods

Embedding also promotes methods.

Example:

```go
type customer struct {
	name string
}

func (c customer) getName() string {
	return c.name
}

type order struct {
	id string
	customer
}
```

Now:

```go
newOrder := order{
	id: "001",
	customer: customer{
		name: "Rahul",
	},
}

fmt.Println(newOrder.getName())
```

The method from `customer` can be accessed through `order`.

This is another important reason why embedding is powerful.

---

# 11. Struct Embedding Is NOT Inheritance

Go does not use traditional class inheritance like Java or C++.

Instead of:

```text
Order extends Customer
```

Go encourages:

```text
Order
 └── Customer
```

This is composition.

The idea is:

> Prefer combining types instead of creating deep inheritance hierarchies.

---

# 12. Why Composition Is Useful

Composition helps you create reusable structures.

For example:

```go
type address struct {
	city    string
	country string
}

type customer struct {
	name string
	address
}

type company struct {
	name string
	address
}
```

Both `customer` and `company` can reuse `address`.

This avoids duplicating:

```go
city
country
```

in multiple structs.

---

# 13. Embedding `time.Time`

Go's standard library also uses embedding patterns.

For example:

```go
type MyTime struct {
	time.Time
}
```

Now methods of `time.Time` can be promoted to `MyTime`.

This demonstrates that embedding isn't just for your own structs; it can also be used with types from packages.

---

# 14. Your `createdAt` Field

You have:

```go
createdAt time.Time
```

`time.Time` represents a point in time.

You can initialize it with:

```go
createdAt: time.Now(),
```

Example:

```go
newOrder := order{
	id:        "001",
	amount:    300,
	status:    "Received",
	createdAt: time.Now(),
	customer: customer{
		name:  "Prakash",
		phone: "123456789",
	},
}
```

Then:

```go
fmt.Println(newOrder.createdAt)
```

will print the timestamp.

`time.Time` supports very high-resolution timestamps, but the actual precision you observe can depend on the platform/runtime clock.

---

# 15. Important: Embedded Field Name

When you write:

```go
type order struct {
	customer
}
```

the embedded field's name is the type name:

```text
customer
```

Therefore this works:

```go
newOrder.customer
```

and:

```go
newOrder.customer.name
```

---

# 16. Naming Convention

Your structs are currently:

```go
type customer struct {}
type order struct {}
```

These names are **unexported** because they start with lowercase letters.

If you need to use them from another package:

```go
type Customer struct {}

type Order struct {}
```

Capital letters make identifiers **exported**.

Example:

```go
type Customer struct {
	Name  string
	Phone string
}
```

Now another package can access:

```go
customer.Name
```

Lowercase fields such as:

```go
name
phone
```

are only accessible within the same package.

---

# 17. Key Concept: Promoted Fields

Remember this:

```go
type order struct {
	customer
}
```

Given:

```go
newOrder.name
```

Go looks through the embedded `customer` and finds:

```go
customer.name
```

Conceptually:

```text
newOrder.name
     ↓
newOrder.customer.name
```

This is called **field promotion**.

---

# 18. If Names Conflict

Suppose both structs contain the same field:

```go
type customer struct {
	name string
}

type order struct {
	name string
	customer
}
```

Now:

```go
newOrder.name
```

refers to the `order` field because the direct field takes precedence.

You can explicitly access the embedded field:

```go
newOrder.customer.name
```

This removes ambiguity.

---

# 19. Multiple Embedded Structs

You can embed multiple types:

```go
type address struct {
	city string
}

type contact struct {
	phone string
}

type customer struct {
	name string
	address
	contact
}
```

Now:

```go
customer.name
customer.city
customer.phone
```

can all be accessed directly through promotion.

---

# 20. Mental Model

Think of embedding like this:

```text
Customer
 ├── name
 └── phone

        ↓ embedded into

Order
 ├── id
 ├── amount
 ├── status
 ├── createdAt
 └── Customer
      ├── name
      └── phone
```

Because `Customer` is embedded, Go promotes its fields and methods.

So:

```go
order.name
```

can reach:

```go
order.customer.name
```

---

# 21. Struct Embedding vs Composition

These terms are related but not exactly identical.

### Composition

General design idea:

```text
Order contains Customer
```

### Struct Embedding

A specific Go language feature:

```go
type order struct {
	customer
}
```

So:

> Embedding is one way to implement composition in Go.

You can also use composition without embedding:

```go
type order struct {
	customer customer
}
```

---

# 22. Common Mistakes

### ❌ Mistake 1 — Thinking embedding means inheritance

```go
type order struct {
	customer
}
```

does not mean:

```text
Order extends Customer
```

It means the `customer` type is embedded in `order`.

---

### ❌ Mistake 2 — Forgetting the embedded type name

This:

```go
type order struct {
	customer
}
```

does not mean there is no `customer` field.

There is an embedded field named:

```go
customer
```

---

### ❌ Mistake 3 — Confusing promoted fields with duplicated fields

When you write:

```go
newOrder.name
```

Go isn't copying `name` into `order`.

The field still belongs to:

```go
customer
```

It is simply **promoted for convenient access**.

---

# 💡 Pro Tips

* Go prefers **composition over inheritance**.
* Embedding is a powerful way to reuse fields and methods.
* Learn **field promotion** carefully; it appears frequently in real Go code.
* You can embed structs as well as other named types.
* Embedded methods can also be promoted.
* Use explicit access like `obj.customer.name` when it makes the structure clearer.
* Be careful with naming conflicts between the outer struct and embedded types.
* Use exported names (`Customer`, `Name`) when the type/field must be accessible from another package.
* Don't use embedding just to make code shorter; use it when the relationship/design makes sense.

---

# 🧠 What I Learned Today

Today I learned:

* What **struct embedding** means.
* How Go uses **composition**.
* How to embed one struct inside another.
* How to initialize an embedded struct separately.
* How to initialize an embedded struct inline.
* How to access an embedded struct.
* What **field promotion** means.
* How promoted fields can be accessed directly.
* How methods can also be promoted.
* Why embedding is not traditional inheritance.
* How multiple structs can be embedded.
* How name conflicts work.
* How `time.Time` can be used as a struct field.

---

# ⚡ Quick Reference

| Concept                 | Example                   |
| ----------------------- | ------------------------- |
| Normal nested struct    | `customer customer`       |
| Embedded struct         | `customer`                |
| Access embedded struct  | `order.customer`          |
| Access nested field     | `order.customer.name`     |
| Promoted field          | `order.name`              |
| Modify promoted field   | `order.name = "Rahul"`    |
| Inline initialization   | `customer: customer{...}` |
| Composition             | `Order has a Customer`    |
| Traditional inheritance | ❌ Not used in Go          |
| Exported type           | `Customer`                |
| Unexported type         | `customer`                |

## ⭐ Most Important

Remember these three ideas:

```text
Composition
    ↓
Struct Embedding
    ↓
Field / Method Promotion
```

Example:

```go
type order struct {
	customer
}
```

Then:

```go
order.customer.name
```

and, because of promotion:

```go
order.name
```

Both can access the customer's `name` field.

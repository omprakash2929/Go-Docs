# 🔢 Go Enumerated Types

## 1. What Are Enumerated Types?

An **enumerated type (enum)** represents a value that should come from a **fixed set of possible values**.

For example, an order can have only certain statuses:

```text
Received
Confirmed
Prepared
Delivered
```

Go does not have a traditional `enum` keyword like some other languages.

Instead, Go commonly uses:

> **Custom types + constants**

---

# 2. Creating a Custom Type

Your code:

```go
type OrderStatus string
```

creates a new type called:

```text
OrderStatus
```

whose underlying type is:

```text
string
```

So:

```go
var status OrderStatus
```

is not exactly the same as:

```go
var status string
```

Even though both are based on strings.

### Why create a custom type?

It gives your code more meaning.

Compare:

```go
func changeOrderStatus(status string)
```

with:

```go
func changeOrderStatus(status OrderStatus)
```

The second one clearly communicates:

> This function expects an order status.

---

# 3. Creating Constants

You created:

```go
const (
	Received  OrderStatus = "received"
	Confirmed            = "confirmed"
	Prepared             = "prepared"
	Delivered             = "delivered"
)
```

The first constant explicitly specifies the type:

```go
Received OrderStatus = "received"
```

The following constants can use the same type:

```go
Confirmed = "confirmed"
Prepared  = "prepared"
Delivered = "delivered"
```

So effectively:

```text
Received  → OrderStatus
Confirmed → OrderStatus
Prepared  → OrderStatus
Delivered → OrderStatus
```

---

# 4. Why Not Just Use Strings?

You could write:

```go
func changeOrderStatus(status string) {
	fmt.Println(status)
}
```

Then someone could accidentally pass:

```go
changeOrderStatus("hello")
```

That doesn't make sense as an order status.

With:

```go
type OrderStatus string
```

your API communicates that the value represents an order status.

For example:

```go
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order status to", status)
}
```

Now the function is specifically designed around `OrderStatus`.

---

# 5. Using the Constants

You can call:

```go
changeOrderStatus(Delivered)
```

Output:

```text
Changing Order status to delivered
```

Other valid predefined values:

```go
changeOrderStatus(Received)
changeOrderStatus(Confirmed)
changeOrderStatus(Prepared)
changeOrderStatus(Delivered)
```

---

# 6. Custom Type Is Still Based on String

Because:

```go
type OrderStatus string
```

you can create a value using a string conversion:

```go
status := OrderStatus("delivered")
```

Now:

```go
fmt.Println(status)
```

prints:

```text
delivered
```

But remember:

```go
OrderStatus
```

is a distinct named type.

---

# 7. Custom Type vs Type Alias ⭐

This is important.

### Custom type

```go
type OrderStatus string
```

Creates a **new named type**.

### Type alias

```go
type OrderStatus = string
```

This does **not** create a new type.

It is simply another name for `string`.

So:

```go
type OrderStatus string
```

is what you want for enum-like values.

---

# 8. Why Constants Are Useful

Constants represent values that should not change.

For example:

```go
const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)
```

You can't do:

```go
Delivered = "something else"
```

because constants cannot be reassigned.

---

# 9. Your `OrderStatus` Design

Your code creates a clean domain model:

```text
OrderStatus
    │
    ├── Received
    ├── Confirmed
    ├── Prepared
    └── Delivered
```

Then:

```go
func changeOrderStatus(status OrderStatus)
```

works with that domain concept.

This is much clearer than passing random strings everywhere.

---

# 10. Using `switch` with Enum-Like Values

This becomes very useful when combined with the `switch` statement you learned earlier.

```go
func processOrder(status OrderStatus) {

	switch status {
	case Received:
		fmt.Println("Order received")

	case Confirmed:
		fmt.Println("Order confirmed")

	case Prepared:
		fmt.Println("Order prepared")

	case Delivered:
		fmt.Println("Order delivered")

	default:
		fmt.Println("Unknown order status")
	}
}
```

Now:

```go
processOrder(Delivered)
```

Output:

```text
Order delivered
```

---

# 11. Another Real-World Example — Payment Status

You can create another custom type:

```go
type PaymentStatus string

const (
	Pending PaymentStatus = "pending"
	Paid                   = "paid"
	Failed                 = "failed"
	Refunded               = "refunded"
)
```

Then:

```go
func updatePaymentStatus(status PaymentStatus) {
	fmt.Println("Payment status:", status)
}
```

Usage:

```go
updatePaymentStatus(Paid)
```

This is the same pattern as your `OrderStatus`.

---

# 12. Another Example — User Role

```go
type UserRole string

const (
	Admin  UserRole = "admin"
	Manager         = "manager"
	User            = "user"
	Guest           = "guest"
)
```

Then:

```go
func checkPermission(role UserRole) {
	switch role {
	case Admin:
		fmt.Println("Full access")

	case Manager:
		fmt.Println("Management access")

	case User:
		fmt.Println("Normal access")

	case Guest:
		fmt.Println("Read-only access")
	}
}
```

---

# 13. `iota` — Numeric Enum Values

You also tried this:

```go
const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)
```

This is where `iota` becomes important.

`iota` is a special constant generator that starts at:

```text
0
```

and increases by `1` for each constant declaration in the same `const` block.

So:

```go
const (
	Received = iota
	Confirmed
	Prepared
	Delivered
)
```

produces:

```text
Received  = 0
Confirmed = 1
Prepared  = 2
Delivered = 3
```

---

# 14. But Your `iota` Example Has a Type Problem

You wrote:

```go
const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)
```

But:

```go
type OrderStatus string
```

means `OrderStatus` is based on `string`.

`iota` produces an integer constant.

So this design doesn't match:

```text
OrderStatus → string
iota        → integer
```

If you want string statuses like:

```text
received
confirmed
prepared
delivered
```

your original version is better:

```go
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)
```

---

# 15. When Should You Use `iota`?

Use `iota` when numeric values make sense.

Example:

```go
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)
```

Values become:

```text
Debug   = 0
Info    = 1
Warning = 2
Error   = 3
```

This is a good use case for `iota`.

---

# 16. `iota` Doesn't Have to Start at 0

You can skip zero:

```go
type LogLevel int

const (
	Debug LogLevel = iota + 1
	Info
	Warning
	Error
)
```

Now:

```text
Debug   = 1
Info    = 2
Warning = 3
Error   = 4
```

This can be useful when `0` should represent an invalid/unset value.

---

# 17. Bit Flags with `iota`

`iota` is also useful for bit flags.

Example:

```go
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Delete
	Execute
)
```

This generates powers of two:

```text
Read    = 1
Write   = 2
Delete  = 4
Execute = 8
```

This is a more advanced use of `iota`.

You don't need to focus on it yet, but remember that `iota` is useful beyond simple numbering.

---

# 18. Important Difference: Constants Don't Automatically Restrict Values

This is a subtle but important point.

Even though you have:

```go
const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)
```

you can still create:

```go
status := OrderStatus("something")
```

Go does not automatically say:

> "Only `Received`, `Confirmed`, `Prepared`, and `Delivered` are allowed."

The type provides semantic meaning, but it does not enforce that the value is one of those constants.

If you need validation, you can write:

```go
func isValidOrderStatus(status OrderStatus) bool {

	switch status {
	case Received, Confirmed, Prepared, Delivered:
		return true

	default:
		return false
	}
}
```

---

# 19. Validation Example

```go
func changeOrderStatus(status OrderStatus) {

	if !isValidOrderStatus(status) {
		fmt.Println("Invalid order status")
		return
	}

	fmt.Println("Changing order status to", status)
}
```

Now:

```go
changeOrderStatus(Delivered)
```

works.

But:

```go
changeOrderStatus(OrderStatus("cancelled"))
```

can be rejected by your validation.

---

# 20. A Better Order Example

You can combine today's topics with structs and methods:

```go
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

type Order struct {
	ID     string
	Status OrderStatus
}

func (o *Order) changeStatus(status OrderStatus) {

	switch status {
	case Received, Confirmed, Prepared, Delivered:
		o.Status = status

	default:
		fmt.Println("Invalid status")
	}
}
```

Usage:

```go
order := Order{
	ID:     "001",
	Status: Received,
}

order.changeStatus(Confirmed)

fmt.Println(order.Status)
```

Output:

```text
confirmed
```

This is getting closer to how you might model real application/domain code.

---

# 21. Why This Pattern Is Useful in Backend Development

You'll frequently have fixed categories such as:

```text
Order Status
Payment Status
User Role
Log Level
Environment
Request Method
Job Status
Deployment Status
Node Status
```

Instead of scattering strings everywhere:

```go
"pending"
"pending"
"pending"
"pendng" // typo 😬
```

you can use:

```go
Pending
```

This makes code easier to read and reduces spelling mistakes.

---

# 22. Go Does Not Have Traditional Enums

Remember:

```text
Go
 └── no enum keyword
```

Instead, common patterns are:

```go
type Status string

const (
	Pending Status = "pending"
	Success        = "success"
	Failed         = "failed"
)
```

or numeric:

```go
type Status int

const (
	Pending Status = iota
	Success
	Failed
)
```

---

# 23. String Enum vs Numeric Enum

| Approach   | Example              | Good For                |
| ---------- | -------------------- | ----------------------- |
| String     | `Status string`      | API/domain values       |
| Integer    | `Status int`         | Compact internal states |
| `iota`     | `Pending = iota`     | Sequential constants    |
| `iota + 1` | `Pending = iota + 1` | Avoiding zero           |
| Bit flags  | `1 << iota`          | Permissions/options     |

For your order-status example, **string-based status is a natural choice**, especially when the value may appear in JSON or an API.

---

# 💡 Pro Tips

### 1. Use meaningful custom types

Instead of:

```go
func update(status string)
```

prefer:

```go
func update(status OrderStatus)
```

when the domain concept is important.

### 2. Use constants for known values

```go
Received
Confirmed
Prepared
Delivered
```

is better than repeatedly typing raw strings.

### 3. Be careful with typos in constant values

Your original code has:

```go
"recived"
"perpared"
```

The English spellings should be:

```go
"received"
"prepared"
```

### 4. Use `iota` for numeric sequences

Don't use `iota` when your actual values need to be meaningful strings.

### 5. Validate when necessary

Custom types + constants don't automatically prevent invalid values.

### 6. Remember zero values

For:

```go
type Status int
```

the zero value is:

```text
0
```

So design numeric enums carefully if `0` has special meaning.

---

# 🧠 What I Learned Today

Today I learned:

* Go doesn't have a traditional `enum` keyword.
* Go commonly creates enum-like values using **custom types + constants**.
* `type OrderStatus string` creates a new named type.
* Constants represent the predefined values.
* `OrderStatus` gives semantic meaning to a string.
* Custom types are different from type aliases.
* `iota` generates sequential constant values.
* `iota` starts from `0`.
* `iota + 1` can make the first value `1`.
* `iota` is useful for numeric enums and bit flags.
* String-based constants are useful for API/domain values.
* Constants don't automatically prevent invalid values.
* Validation can be added using `switch`.
* Enum-like values work nicely with `switch`.
* This pattern is useful for order status, payment status, user roles, log levels, and job states.

---

# ⚡ Quick Reference

### String-based enum

```go
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)
```

### Numeric enum

```go
type Status int

const (
	Pending Status = iota
	Success
	Failed
)
```

### Start from 1

```go
const (
	One Status = iota + 1
	Two
	Three
)
```

### Function using the custom type

```go
func changeOrderStatus(status OrderStatus) {
	fmt.Println(status)
}
```

### Use the constant

```go
changeOrderStatus(Delivered)
```

### Validate

```go
switch status {
case Received, Confirmed, Prepared, Delivered:
	// valid
default:
	// invalid
}
```

---

# ⭐ Mental Model

Remember:

```text
Custom Type
     ↓
OrderStatus
     ↓
   Constants
     ↓
Received | Confirmed | Prepared | Delivered
```

So instead of passing random strings around your application:

```go
changeOrderStatus("delivered")
```

you can express the domain more clearly:

```go
changeOrderStatus(Delivered)
```

**Custom type = meaning**

**Constant = predefined value**

**`iota` = automatic numeric sequence**

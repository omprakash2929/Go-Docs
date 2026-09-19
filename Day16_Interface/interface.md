# 🔌 Go Interfaces

## 1. What is an Interface?

An **interface** defines a set of methods that a type must implement.

Example:

```go
type paymenter interface {
	pay(amount float32)
}
```

This interface says:

> Any type that has a `pay(float32)` method can be used as a `paymenter`.

The interface does **not** contain the implementation of `pay()`.

It only defines the **behavior** that is required.

---

# 2. Interface vs Function

In an interface:

```go
type paymenter interface {
	pay(amount float32)
}
```

`pay()` is a **method requirement**, not a function implementation.

We don't write:

```go
func pay(amount float32) {
	// ...
}
```

inside the interface.

The actual implementation is provided by another type.

---

# 3. Payment Example

We have this interface:

```go
type paymenter interface {
	pay(amount float32)
}
```

Now we create Razorpay:

```go
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using Razorpay", amount)
}
```

And Stripe:

```go
type strip struct{}

func (s strip) pay(amount float32) {
	fmt.Println("Making payment using Stripe", amount)
}
```

Both types implement:

```go
pay(amount float32)
```

Therefore both can be used as a `paymenter`.

---

# 4. Important Rule ⭐

Go interfaces use **implicit implementation**.

You don't have to write:

```go
implements paymenter
```

There is no `implements` keyword in Go.

If a type has all the methods required by an interface, it automatically satisfies that interface.

For example:

```go
type paymenter interface {
	pay(amount float32)
}
```

Razorpay has:

```go
func (r razorpay) pay(amount float32)
```

Therefore:

```text
razorpay → satisfies paymenter
```

Stripe also has:

```go
func (s strip) pay(amount float32)
```

Therefore:

```text
strip → satisfies paymenter
```

---

# 5. The Main Problem Without Interface

Imagine we write:

```go
func (p payment) makePayment(amount float32) {

	razorpayPayment := razorpay{}
	razorpayPayment.pay(amount)
}
```

Now `payment` is tightly connected to Razorpay.

If tomorrow we want Stripe:

```go
func (p payment) makePayment(amount float32) {

	stripePayment := strip{}
	stripePayment.pay(amount)
}
```

We have to modify the `payment` code.

This creates unnecessary coupling.

---

# 6. Interface Solves This Problem

Instead, we create:

```go
type payment struct {
	gatway paymenter
}
```

Notice:

```go
gatway paymenter
```

We don't say:

```go
gatway razorpay
```

or:

```go
gatway strip
```

We say:

```go
gatway paymenter
```

This means:

> `payment` doesn't care which payment gateway is being used. It only cares that the gateway can `pay()`.

---

# 7. `makePayment()`

Now:

```go
func (p payment) makePayment(amount float32) {
	p.gatway.pay(amount)
}
```

This is the important part.

`makePayment()` doesn't know whether:

```text
Razorpay
Stripe
Fake payment
Future gateway
```

is being used.

It only knows:

```text
I have something that can pay().
```

---

# 8. Using Razorpay

```go
razorpayGateway := razorpay{}

newPayment := payment{
	gatway: razorpayGateway,
}

newPayment.makePayment(100)
```

Output:

```text
Making payment using Razorpay 100
```

---

# 9. Using Stripe

We don't need to change `makePayment()`.

Just change the gateway:

```go
stripeGateway := strip{}

newPayment := payment{
	gatway: stripeGateway,
}

newPayment.makePayment(100)
```

Output:

```text
Making payment using Stripe 100
```

This is the power of interfaces.

---

# 10. Fake Payment for Testing ⭐

Your example has another very important use case:

```go
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making payment using Fake for testing:", amount)
}
```

`fakepayment` also satisfies:

```go
paymenter
```

because it has:

```go
pay(amount float32)
```

Now:

```go
fakeGw := fakepayment{}

newPayment := payment{
	gatway: fakeGw,
}

newPayment.makePayment(100)
```

Output:

```text
Making payment using Fake for testing: 100
```

---

# 11. Why Fake Payment Is Useful?

Imagine your real application uses a real payment provider.

During testing, you don't want to:

```text
Application
    ↓
Real Payment Gateway
    ↓
Real transaction
    ↓
Real money
```

Instead:

```text
Application
    ↓
Fake Payment
    ↓
Test
```

You can test your application without making a real transaction.

This is called using a **test double/fake implementation**.

---

# 12. Open/Closed Principle

Your comment says:

```go
// open close principle
```

This example demonstrates the idea behind the **Open/Closed Principle (OCP)**:

> Software should be open for extension but closed for modification.

Suppose we already have:

```text
Razorpay
Stripe
```

Later we add:

```text
PayPal
```

We can create:

```go
type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using PayPal", amount)
}
```

We don't need to change:

```go
func (p payment) makePayment(amount float32)
```

We simply pass the new gateway:

```go
paypalGateway := paypal{}

newPayment := payment{
	gatway: paypalGateway,
}

newPayment.makePayment(100)
```

That's the important design benefit.

---

# 13. Simple Mental Model

Think of an interface as a **contract**.

```text
             paymenter
            /    |     \
           /     |      \
      Razorpay  Stripe  Fake
```

The contract says:

```text
"You must have pay(amount)."
```

Razorpay says:

```text
"I have pay()."
```

Stripe says:

```text
"I have pay()."
```

Fake says:

```text
"I have pay()."
```

Therefore all three can be used wherever `paymenter` is expected.

---

# 14. Interface Is About Behavior

This is one of the most important concepts.

Don't think:

> Interface = common data

Think:

> Interface = common behavior

For example:

```go
type paymenter interface {
	pay(amount float32)
}
```

The interface doesn't care about:

```text
How Razorpay works
How Stripe works
How Fake works
```

It only cares:

```text
Can this thing pay?
```

---

# 15. Another Easy Example — Notification

Imagine your application needs to send notifications.

We can create:

```go
type notifier interface {
	send(message string)
}
```

Now email:

```go
type email struct{}

func (e email) send(message string) {
	fmt.Println("Sending Email:", message)
}
```

SMS:

```go
type sms struct{}

func (s sms) send(message string) {
	fmt.Println("Sending SMS:", message)
}
```

WhatsApp:

```go
type whatsapp struct{}

func (w whatsapp) send(message string) {
	fmt.Println("Sending WhatsApp:", message)
}
```

All of them satisfy:

```go
notifier
```

So we can create:

```go
func notify(n notifier, message string) {
	n.send(message)
}
```

Then:

```go
notify(email{}, "Your order has been shipped")
notify(sms{}, "Your OTP is 1234")
notify(whatsapp{}, "Your order is ready")
```

The `notify()` function doesn't care about the actual notification provider.

---

# 16. Another Easy Example — Storage

Imagine an application needs to save data.

We can define:

```go
type storage interface {
	save(data string)
}
```

File storage:

```go
type fileStorage struct{}

func (f fileStorage) save(data string) {
	fmt.Println("Saving to file:", data)
}
```

Database:

```go
type database struct{}

func (d database) save(data string) {
	fmt.Println("Saving to database:", data)
}
```

Cloud storage:

```go
type cloudStorage struct{}

func (c cloudStorage) save(data string) {
	fmt.Println("Saving to cloud:", data)
}
```

Now:

```go
func storeData(s storage, data string) {
	s.save(data)
}
```

We can use:

```go
storeData(fileStorage{}, "hello")
storeData(database{}, "hello")
storeData(cloudStorage{}, "hello")
```

Same function.

Different implementations.

---

# 17. Another Example — Logger

Interfaces are very common for logging.

```go
type logger interface {
	log(message string)
}
```

Console logger:

```go
type consoleLogger struct{}

func (c consoleLogger) log(message string) {
	fmt.Println("Console:", message)
}
```

File logger:

```go
type fileLogger struct{}

func (f fileLogger) log(message string) {
	fmt.Println("File:", message)
}
```

Application code:

```go
func process(l logger) {
	l.log("Processing request...")
}
```

Now:

```go
process(consoleLogger{})
process(fileLogger{})
```

The business logic doesn't need to know where the log is going.

---

# 18. Interface as a Function Parameter

You don't always need to store an interface inside a struct.

You can directly accept an interface:

```go
func makePayment(gateway paymenter, amount float32) {
	gateway.pay(amount)
}
```

Now:

```go
makePayment(razorpay{}, 100)
makePayment(strip{}, 200)
makePayment(fakepayment{}, 300)
```

This is a very common pattern.

---

# 19. Interface as a Return Type

A function can also return an interface.

Example:

```go
func getPaymentGateway() paymenter {
	return razorpay{}
}
```

Then:

```go
gateway := getPaymentGateway()

gateway.pay(100)
```

The caller only knows:

```text
paymenter
```

not the exact implementation.

---

# 20. Multiple Methods in an Interface

An interface can have multiple methods:

```go
type paymenter interface {
	pay(amount float32)
	refund(amount float32)
}
```

Now a type must implement **both** methods:

```go
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment:", amount)
}

func (r razorpay) refund(amount float32) {
	fmt.Println("Refund:", amount)
}
```

Now Razorpay satisfies the interface.

---

# 21. Small Interfaces Are Better ⭐

In Go, interfaces are often kept small.

Instead of:

```go
type hugeInterface interface {
	pay()
	refund()
	cancel()
	verify()
	update()
	delete()
	save()
	log()
}
```

prefer focused interfaces:

```go
type payer interface {
	pay(amount float32)
}
```

A small interface is easier to implement, test, and reuse.

---

# 22. Interface Composition

Interfaces can also be combined.

```go
type payer interface {
	pay(amount float32)
}

type refunder interface {
	refund(amount float32)
}

type paymentService interface {
	payer
	refunder
}
```

Now `paymentService` requires both:

```text
pay()
refund()
```

---

# 23. Empty Interface

You may also see:

```go
interface{}
```

Modern Go commonly writes:

```go
any
```

These are equivalent:

```go
interface{}
```

and:

```go
any
```

An empty interface doesn't require any methods.

Therefore almost any value can be stored in it.

Example:

```go
var value any

value = 10
value = "hello"
value = true
```

But don't use `any` everywhere.

Prefer a specific interface when you know the required behavior.

---

# 24. Your Complete Payment Example

```go
package main

import "fmt"

// Interface
type paymenter interface {
	pay(amount float32)
}

// Payment service
type payment struct {
	gatway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gatway.pay(amount)
}

// Razorpay implementation
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using Razorpay:", amount)
}

// Stripe implementation
type strip struct{}

func (s strip) pay(amount float32) {
	fmt.Println("Making payment using Stripe:", amount)
}

// Fake implementation for testing
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making payment using Fake for testing:", amount)
}

func main() {

	fakeGw := fakepayment{}

	newPayment := payment{
		gatway: fakeGw,
	}

	newPayment.makePayment(100)
}
```

The important relationship is:

```text
payment
   |
   | depends on
   ↓
paymenter
   ↑
   |
-----------------------------
|            |             |
Razorpay    Stripe        Fake
```

---

# 🧠 The Main Idea

Without interface:

```text
Payment → Razorpay
```

Payment becomes tightly coupled to Razorpay.

With interface:

```text
              paymenter
             /    |    \
            /     |     \
      Razorpay  Stripe  Fake
```

Now:

```text
Payment → paymenter
```

The implementation can change without changing the payment logic.

---

# 🚀 Real-World Places Where Interfaces Are Used

You'll see interfaces frequently in:

* Payment gateways
* Database repositories
* HTTP clients
* Storage systems
* Logging
* Authentication providers
* Notification systems
* Cloud services
* Message queues
* Cache systems
* Testing/mocking
* File systems
* Different API implementations

For your **DevOps/backend/distributed-systems direction**, interfaces are especially important because large Go systems often have many interchangeable implementations.

---

# 💡 Pro Tips

### 1. Think behavior, not object

Instead of:

> "What object is this?"

Think:

> "What can this type do?"

---

### 2. Interfaces are satisfied implicitly

No:

```go
implements paymenter
```

Just implement the required methods.

---

### 3. Keep interfaces small

Prefer:

```go
type payer interface {
	pay(amount float32)
}
```

over unnecessarily large interfaces.

---

### 4. Interfaces make testing easier

Real implementation:

```text
payment → Razorpay
```

Testing:

```text
payment → FakePayment
```

---

### 5. Don't create interfaces for everything

An interface is useful when you actually need abstraction, substitution, testing, or multiple implementations.

---

# 🧠 What I Learned Today

Today I learned:

* What an **interface** is.
* An interface defines required **behavior** through methods.
* Go interfaces are implemented **implicitly**.
* A type satisfies an interface by implementing its required methods.
* Razorpay, Stripe, and FakePayment can all implement the same interface.
* Interfaces reduce **coupling**.
* Interfaces help follow the **Open/Closed Principle**.
* Interfaces make testing easier through fake implementations.
* Interfaces can be used as struct fields.
* Interfaces can be used as function parameters.
* Interfaces can be returned from functions.
* Interfaces can contain multiple methods.
* Small interfaces are usually easier to work with.
* Interfaces can be composed from other interfaces.
* `any` is the modern name for `interface{}`.
* Interfaces describe **what something can do**, not how it does it.

---

# ⚡ Quick Reference

| Concept                       | Example                            |
| ----------------------------- | ---------------------------------- |
| Define interface              | `type paymenter interface { ... }` |
| Interface method              | `pay(amount float32)`              |
| Implement interface           | Define the required method         |
| Explicit `implements` keyword | ❌ Not required                     |
| Interface field               | `gateway paymenter`                |
| Interface parameter           | `func pay(p paymenter)`            |
| Multiple implementations      | Razorpay / Stripe / Fake           |
| Testing                       | Fake implementation                |
| Main benefit                  | Loose coupling                     |
| Design principle              | Open/Closed Principle              |
| Empty interface               | `any` / `interface{}`              |
| Good interface                | Small and behavior-focused         |

## ⭐ Remember This

```text
Interface = Contract

Type = Implementation

If the type has all required methods
        ↓
It satisfies the interface
```

For your payment example:

```go
type paymenter interface {
	pay(amount float32)
}
```

means:

> **"I don't care whether you are Razorpay, Stripe, or Fake. If you can `pay(float32)`, I can work with you."**

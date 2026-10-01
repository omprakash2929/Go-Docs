# Go Mutex — Protecting Shared Data

## 📌 Overview

When multiple goroutines access and modify the **same shared data**, a **data race** can happen.

Go provides `sync.Mutex` to protect shared data.

A Mutex works like a lock:

```text
Lock → Only one goroutine can enter
       ↓
    Do work
       ↓
Unlock → Other goroutines can enter
```

The main idea:

> **Only one goroutine should modify the protected data at a time.**

---

# 1. The Problem: Shared Data

Suppose we have a post:

```go
type post struct {
    views int
}
```

And 100 goroutines increase the views:

```go
for i := 0; i < 100; i++ {
    go myPost.inc()
}
```

All goroutines are accessing:

```go
p.views
```

at the same time.

This can cause a **data race**.

---

# 2. What Is a Data Race?

A data race happens when multiple goroutines access the same memory concurrently and at least one of them modifies it, without proper synchronization.

For example:

```text
Goroutine 1 → read views
Goroutine 2 → read views
Goroutine 1 → write views
Goroutine 2 → write views
```

Imagine:

```text
views = 10
```

Two goroutines both read `10`.

```text
Goroutine 1: 10 → 11
Goroutine 2: 10 → 11
```

Expected:

```text
12
```

But the result can become:

```text
11
```

because both goroutines operated on the same old value.

This is called a **lost update**.

---

# 3. Mutex

A Mutex stands for:

> **Mutual Exclusion**

Go provides it through:

```go
sync.Mutex
```

Import:

```go
import "sync"
```

Create one:

```go
var mu sync.Mutex
```

A Mutex provides two main methods:

```go
mu.Lock()
mu.Unlock()
```

### `Lock()`

```go
mu.Lock()
```

means:

> "I want exclusive access to the protected data."

If another goroutine already owns the lock, the current goroutine waits.

### `Unlock()`

```go
mu.Unlock()
```

means:

> "I'm finished. Another goroutine can access the protected data."

---

# 4. Your Struct

Your code:

```go
type post struct {
    views int
    mu    sync.Mutex
}
```

This is a good design.

The Mutex is stored together with the data it protects:

```text
post
 ├── views
 └── mu
```

The idea is:

```text
mu protects views
```

This makes it easier to understand which data the lock belongs to.

---

# 5. Your `inc()` Method

Your method:

```go
func (p *post) inc(wg *sync.WaitGroup) {
    defer func() {
        p.mu.Unlock()
        wg.Done()
    }()

    p.mu.Lock()
    p.views += 1
}
```

The important part is:

```go
p.mu.Lock()

p.views += 1

p.mu.Unlock()
```

This creates a **critical section**.

---

# 6. What Is a Critical Section?

A critical section is a part of code where shared data is accessed or modified and therefore needs synchronization.

In your example:

```go
p.mu.Lock()

p.views += 1

p.mu.Unlock()
```

The critical section is:

```go
p.views += 1
```

Only one goroutine should perform this protected operation at a time.

Think:

```text
Goroutine 1
     |
     | Lock
     ↓
┌──────────────┐
│ views += 1   │  ← Critical Section
└──────────────┘
     |
     | Unlock
     ↓

Goroutine 2 can now enter
```

---

# 7. Why `defer Unlock()`?

Your code does:

```go
defer func() {
    p.mu.Unlock()
    wg.Done()
}()
```

This works because the deferred function runs when `inc()` returns.

A cleaner common pattern is:

```go
func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()

    p.mu.Lock()
    defer p.mu.Unlock()

    p.views++
}
```

This is easier to read.

Another common pattern is:

```go
func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()

    p.mu.Lock()
    p.views++
    p.mu.Unlock()
}
```

Both are valid.

The important rule is:

> Every successful `Lock()` must eventually have a corresponding `Unlock()`.

---

# 8. Why Lock Before Modifying?

Your code:

```go
p.mu.Lock()
p.views += 1
```

is correct.

The lock must happen **before** accessing the shared data.

Correct:

```go
p.mu.Lock()

p.views++

p.mu.Unlock()
```

Not:

```go
p.views++
p.mu.Lock()
```

Because by the time you acquire the lock, the unsafe operation has already happened.

---

# 9. Complete Corrected Example

```go
package main

import (
    "fmt"
    "sync"
)

type post struct {
    views int
    mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()

    p.mu.Lock()
    defer p.mu.Unlock()

    p.views++
}

func main() {
    var wg sync.WaitGroup

    myPost := post{views: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg)
    }

    wg.Wait()

    fmt.Println(myPost.views)
}
```

Expected output:

```text
100
```

Every goroutine increments the value exactly once.

---

# 10. Why WaitGroup + Mutex?

These two tools solve **different problems**.

### WaitGroup

```go
wg.Add(1)
wg.Done()
wg.Wait()
```

Answers:

> "Are all goroutines finished?"

### Mutex

```go
mu.Lock()
mu.Unlock()
```

Answers:

> "Can this goroutine safely access this shared data right now?"

So your program uses both:

```text
              100 Goroutines
                    |
          ┌─────────┴─────────┐
          ↓                   ↓
       WaitGroup            Mutex
          |                   |
    Wait for all        Protect shared
      goroutines            data
          |                   |
          └─────────┬─────────┘
                    ↓
              views = 100
```

This distinction is extremely important.

---

# 11. Why `views` Doesn't Need `*int`

You originally had:

```go
type post struct {
    views *int
    mu    sync.Mutex
}
```

But you can simply use:

```go
type post struct {
    views int
    mu    sync.Mutex
}
```

Then:

```go
p.views++
```

is enough.

You don't need:

```go
*int
```

because the `post` itself is already being accessed through a pointer receiver:

```go
func (p *post) inc(...)
```

So `p.views` directly refers to the field in the original `post`.

---

# 12. Pointer Receiver

You have:

```go
func (p *post) inc(wg *sync.WaitGroup)
```

The `*post` is important because the method modifies:

```go
p.views
```

We want to modify the original `post`.

So:

```text
myPost
  |
  ↓
*post
  |
  ├── views
  └── mu
```

All goroutines operate on the same `post`.

That means they also use the same Mutex.

---

# 13. Very Important: Don't Copy a Mutex

This is one of the important rules when working with `sync.Mutex`.

Avoid copying a struct containing a Mutex after the Mutex has been used.

For example, don't casually do:

```go
func process(p post) {
    // p is a copy
}
```

Instead, when appropriate:

```go
func process(p *post) {
    // original post
}
```

Your pointer receiver:

```go
func (p *post) inc(...)
```

is therefore a good choice.

---

# 14. Race Detector

Go has a built-in race detector.

Run your program with:

```bash
go run -race main.go
```

If your code contains a data race, Go can report it.

This is extremely useful when working with concurrent programs.

You can also use:

```bash
go test -race ./...
```

for tests.

### Important

A Mutex prevents a particular race only if **all accesses to that shared data follow the same synchronization strategy**.

Adding a Mutex somewhere does not automatically make every access safe.

---

# 15. Mutex and Channels

You have now learned two different approaches to concurrency.

### Channels

Used mainly for:

> Communication between goroutines.

```go
ch <- data
data := <-ch
```

### Mutex

Used mainly for:

> Protecting shared memory.

```go
mu.Lock()
sharedData++
mu.Unlock()
```

Simple mental model:

```text
Channels → "Talk to each other"

Mutex → "Protect this shared thing"
```

Go's concurrency model often encourages communication through channels, but Mutexes are absolutely useful when shared state is the natural design.

---

# 16. `Mutex` vs `WaitGroup` vs `Channel`

| Tool      | Main Purpose                             |
| --------- | ---------------------------------------- |
| Goroutine | Run work concurrently                    |
| WaitGroup | Wait for goroutines                      |
| Channel   | Communicate between goroutines           |
| Mutex     | Protect shared data                      |
| RWMutex   | Allow multiple readers / protect writers |

This table is worth remembering.

---

# 17. `sync.RWMutex`

Go also provides:

```go
sync.RWMutex
```

It has:

```go
RLock()
RUnlock()

Lock()
Unlock()
```

The idea is:

```text
Multiple readers → Can read together

Writer → Needs exclusive access
```

Example:

```go
type post struct {
    views int
    mu    sync.RWMutex
}
```

Reading:

```go
p.mu.RLock()
views := p.views
p.mu.RUnlock()
```

Writing:

```go
p.mu.Lock()
p.views++
p.mu.Unlock()
```

This is useful when:

* Reads happen frequently.
* Writes happen less frequently.
* You want concurrent readers.

You'll learn this more deeply later.

---

# 18. Common Mistakes

### ❌ Forgetting `Unlock()`

```go
mu.Lock()

sharedData++

// Forgot Unlock
```

Other goroutines may remain blocked forever.

Use:

```go
mu.Lock()
defer mu.Unlock()
```

when appropriate.

---

### ❌ Unlocking without locking

```go
mu.Unlock()
```

without owning the lock can cause a runtime panic.

---

### ❌ Locking too much code

Avoid:

```go
mu.Lock()

doLargeNetworkRequest()
doHeavyCalculation()
doSomethingElse()

p.views++

mu.Unlock()
```

If only `views++` needs protection, keep the critical section small:

```go
mu.Lock()
p.views++
mu.Unlock()
```

Why?

Because while one goroutine holds the lock, other goroutines may have to wait.

---

### ❌ Using a Mutex without understanding what it protects

Ask:

> "Which shared data is this lock protecting?"

In your program:

```text
mu → protects views
```

That makes the design clear.

---

# 19. A Useful Pattern

A common Mutex pattern is:

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()

    return c.value
}
```

Now the struct controls access to its own shared state.

This pattern is common in:

* Counters
* Caches
* In-memory stores
* Connection managers
* Metrics
* Concurrent services

---

# 20. Mental Model

Imagine a bathroom with one key.

```text
             Shared Data
                 |
                 v
          ┌──────────────┐
          │     Mutex    │
          │     🔑       │
          └──────┬───────┘
                 |
       ┌─────────┴─────────┐
       ↓                   ↓
  Goroutine 1          Goroutine 2
       |                   |
    gets lock           waits
       |
    modifies
       |
   unlocks
       |
       └──────────────→ Goroutine 2
                          gets lock
```

Only one goroutine can enter the critical section at a time.

---

# 🎯 What I Learned Today

* A **data race** happens when goroutines access shared data concurrently without proper synchronization.
* `sync.Mutex` provides mutual exclusion.
* `Lock()` acquires the lock.
* `Unlock()` releases the lock.
* A **critical section** is the code that accesses protected shared data.
* `WaitGroup` and `Mutex` solve different problems.
* `WaitGroup` waits for goroutines.
* `Mutex` protects shared data.
* A pointer receiver is useful when modifying the original struct.
* `views` does not need to be `*int`.
* Avoid copying a Mutex after it has been used.
* `go run -race` can detect data races.
* `sync.RWMutex` provides separate read and write locking.
* Keep critical sections small.
* Every successful `Lock()` needs a matching `Unlock()`.

---

# ⚡ Quick Reference

```go
import "sync"
```

### Mutex

```go
var mu sync.Mutex

mu.Lock()

// critical section

mu.Unlock()
```

### Recommended Pattern

```go
mu.Lock()
defer mu.Unlock()

// protected code
```

### WaitGroup + Mutex

```go
var wg sync.WaitGroup

for i := 0; i < 100; i++ {
    wg.Add(1)

    go func() {
        defer wg.Done()

        mu.Lock()
        counter++
        mu.Unlock()
    }()
}

wg.Wait()
```

### Race Detector

```bash
go run -race main.go
```

### Remember

```text
Goroutine → concurrent execution

Channel   → communication

WaitGroup → wait for goroutines

Mutex     → protect shared data

RWMutex   → multiple readers / exclusive writer
```

> **Concurrency is not only about running things at the same time.
> It's also about making shared data safe.**

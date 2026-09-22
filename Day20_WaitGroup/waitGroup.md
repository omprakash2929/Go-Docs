# Go WaitGroup — Waiting for Goroutines

## 📌 Overview

`sync.WaitGroup` is used to **wait until a group of goroutines finishes their work**.

Without `WaitGroup`, the `main()` function may finish before goroutines complete.

Think of it like:

> "Main goroutines start kar raha hoon, but program tab tak finish nahi hoga jab tak sab goroutines complete nahi ho jaate."

---

# 1. Import `sync`

```go
import (
    "fmt"
    "sync"
)
```

`sync` is a standard Go package that provides synchronization tools.

One of its tools is:

```go
sync.WaitGroup
```

---

# 2. Create a WaitGroup

```go
var wg sync.WaitGroup
```

This creates a `WaitGroup`.

We can use it to keep track of how many goroutines are still running.

---

# 3. `wg.Add(1)`

Inside the loop:

```go
wg.Add(1)
```

This tells the WaitGroup:

> "One more goroutine is going to start."

Example:

```go
for i := 0; i <= 10; i++ {
    wg.Add(1)
    go task(i, &wg)
}
```

The loop runs 11 times.

So the counter becomes:

```text
Add(1) → 1
Add(1) → 2
Add(1) → 3
...
Add(1) → 11
```

---

# 4. Passing WaitGroup to the Goroutine

```go
go task(i, &wg)
```

We pass the **address** of `wg`.

Why?

Because the goroutine needs to call:

```go
w.Done()
```

and modify the same WaitGroup counter.

If we passed:

```go
go task(i, wg)
```

we would be passing a copy, which is not what we want.

So:

```go
&wg
```

means:

> Pass the address of the original WaitGroup.

---

# 5. `w.Done()`

Inside `task()`:

```go
func task(id int, w *sync.WaitGroup) {
    defer w.Done()

    fmt.Println("Doing task", id)
}
```

The important line is:

```go
defer w.Done()
```

`Done()` decreases the WaitGroup counter by 1.

Conceptually:

```text
wg.Add(1)
    ↓
Goroutine starts
    ↓
Goroutine finishes
    ↓
wg.Done()
    ↓
Counter - 1
```

If there are 11 goroutines:

```text
Counter = 11

Goroutine 1 finishes → 10
Goroutine 2 finishes → 9
Goroutine 3 finishes → 8
...
Last goroutine finishes → 0
```

---

# 6. Why use `defer`?

You could write:

```go
func task(id int, w *sync.WaitGroup) {
    fmt.Println("Doing task", id)
    w.Done()
}
```

But:

```go
defer w.Done()
```

is safer.

It means:

> "When this function returns, call `Done()`."

So even if later you add more code or an early return:

```go
func task(id int, w *sync.WaitGroup) {
    defer w.Done()

    if id == 5 {
        return
    }

    fmt.Println("Doing task", id)
}
```

`Done()` will still be called.

### Pro Tip

When using `WaitGroup`, a common pattern is:

```go
func task(wg *sync.WaitGroup) {
    defer wg.Done()

    // work
}
```

---

# 7. `wg.Wait()`

After starting all goroutines:

```go
wg.Wait()
```

This is the most important part.

`Wait()` blocks the current goroutine until the WaitGroup counter becomes `0`.

In this example:

```go
for i := 0; i <= 10; i++ {
    wg.Add(1)
    go task(i, &wg)
}

wg.Wait()

fmt.Println("Completed")
```

The flow is:

```text
Create WaitGroup
       ↓
Start goroutine 0
       ↓
Add counter +1
       ↓
Start goroutine 1
       ↓
Add counter +1
       ↓
...
Start goroutine 10
       ↓
Counter = 11
       ↓
wg.Wait()
       ↓
Wait for all goroutines
       ↓
All goroutines call Done()
       ↓
Counter = 0
       ↓
Continue execution
       ↓
Print "Completed"
```

Therefore:

```text
Doing task ...
Doing task ...
Doing task ...
...
Completed
```

`Completed` will be printed **after all goroutines have called `Done()`**.

---

# 8. Complete Code

```go
package main

import (
    "fmt"
    "sync"
)

func task(id int, w *sync.WaitGroup) {
    defer w.Done()

    fmt.Println("Doing task", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 0; i <= 10; i++ {
        wg.Add(1)
        go task(i, &wg)
    }

    wg.Wait()

    fmt.Println("Completed")
}
```

---

# 9. Important: Output Order

Even with `WaitGroup`, the order of tasks is **not guaranteed**.

You might get:

```text
Doing task 10
Doing task 3
Doing task 7
Doing task 1
Doing task 5
...
Completed
```

Or:

```text
Doing task 1
Doing task 2
Doing task 0
Doing task 5
...
Completed
```

That's normal.

`WaitGroup` only guarantees:

> All goroutines are finished before `wg.Wait()` returns.

It does **not** guarantee execution order.

---

# 10. WaitGroup vs `time.Sleep`

Previously you used:

```go
time.Sleep(time.Second * 1)
```

That is not proper synchronization.

### `time.Sleep`

```go
go task()

time.Sleep(time.Second)
```

Problem:

What if the task takes 2 seconds?

```text
Task starts
   ↓
Sleep 1 second
   ↓
main finishes ❌
   ↓
Task may still be running
```

What if the task only needs 100ms?

Then you're unnecessarily waiting.

---

### `WaitGroup`

```go
go task()

wg.Wait()
```

Now Go waits exactly until the goroutines finish.

```text
Task starts
   ↓
Task finishes
   ↓
Done()
   ↓
WaitGroup counter = 0
   ↓
wg.Wait() returns
```

So:

> `WaitGroup` is synchronization.
> `Sleep` is just a delay.

---

# 11. WaitGroup Mental Model

Think of `WaitGroup` as a counter.

```text
              WaitGroup
                  │
        ┌─────────┴─────────┐
        ↓                   ↓
     Add(1)               Done()
        │                   │
        ↓                   ↓
   "One task added"    "One task finished"
        │                   │
        └─────────┬─────────┘
                  ↓
             Counter = 0
                  ↓
              Wait() ends
```

The three important methods are:

```go
wg.Add(1)
wg.Done()
wg.Wait()
```

Remember:

```text
Add()  → increase counter
Done() → decrease counter
Wait() → wait until counter becomes 0
```

---

# 12. Why `*sync.WaitGroup`?

Your function:

```go
func task(id int, w *sync.WaitGroup)
```

Here:

```go
*sync.WaitGroup
```

means:

> `w` is a pointer to a WaitGroup.

And:

```go
go task(i, &wg)
```

passes the address of the original WaitGroup.

This is important because all goroutines need to work with the **same WaitGroup**.

```text
             wg
             │
             ▼
       ┌─────────────┐
       │ WaitGroup   │
       │ counter: 11 │
       └─────────────┘
          ▲   ▲   ▲
          │   │   │
         G1  G2  G3
```

All goroutines update the same synchronization object.

---

# 13. `WaitGroup` Should Not Be Copied

A WaitGroup should be used by reference:

```go
func task(wg *sync.WaitGroup)
```

Not:

```go
func task(wg sync.WaitGroup)
```

And pass:

```go
&wg
```

instead of copying it.

### Important rule

> Don't copy a `sync.WaitGroup` after first use.

Use one WaitGroup and pass its pointer to the goroutines that need it.

---

# 14. Common Pattern

This pattern is worth memorizing:

```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)

    go func(i int) {
        defer wg.Done()

        // work
    }(i)
}

wg.Wait()
```

This is one of the most common basic concurrency patterns in Go.

---

# 15. Real-World Example

Imagine you need to fetch data from 5 APIs.

Instead of:

```text
API 1 → wait
API 2 → wait
API 3 → wait
API 4 → wait
API 5 → wait
```

you can start them concurrently:

```text
             ┌→ API 1
             ├→ API 2
Main ────────┼→ API 3
             ├→ API 4
             └→ API 5
                  ↓
             WaitGroup
                  ↓
             All completed
```

Example:

```go
var wg sync.WaitGroup

for i := 1; i <= 5; i++ {
    wg.Add(1)

    go func(id int) {
        defer wg.Done()

        fmt.Println("Fetching API", id)
    }(i)
}

wg.Wait()

fmt.Println("All APIs completed")
```

This pattern is useful for:

* Parallel API calls
* File processing
* Background jobs
* Worker tasks
* Data processing
* DevOps automation tools
* Infrastructure monitoring

---

# 16. WaitGroup Does NOT Solve Data Races

This is very important.

`WaitGroup` answers:

> "Are all goroutines finished?"

It does **not** answer:

> "Are goroutines safely accessing shared data?"

For example:

```go
counter := 0

var wg sync.WaitGroup

for i := 0; i < 100; i++ {
    wg.Add(1)

    go func() {
        defer wg.Done()
        counter++
    }()
}

wg.Wait()
```

This can have a **data race** because multiple goroutines modify:

```go
counter
```

at the same time.

For shared data, you may need:

* `sync.Mutex`
* `sync.RWMutex`
* `sync/atomic`
* Channels

So remember:

```text
WaitGroup → waiting/synchronization
Mutex     → protecting shared data
Channel   → communication between goroutines
```

---

# 17. Common Mistakes

### ❌ Calling `Done()` without `Add()`

```go
go task()

wg.Wait()
```

without:

```go
wg.Add(1)
```

is incorrect.

---

### ❌ Forgetting `Done()`

```go
func task(wg *sync.WaitGroup) {
    fmt.Println("Doing work")
}
```

Then:

```go
wg.Wait()
```

may wait forever because the counter never reaches zero.

---

### ❌ Calling `Add()` after `Wait()` has started

The usual pattern is:

```go
wg.Add(1)
go task()
```

then:

```go
wg.Wait()
```

Add the work before waiting for completion.

---

### ❌ Copying WaitGroup

Avoid:

```go
func task(wg sync.WaitGroup)
```

Use:

```go
func task(wg *sync.WaitGroup)
```

---

# 18. WaitGroup vs Channel

These are different tools.

### WaitGroup

Used when you mainly need:

> "Wait until these goroutines finish."

```go
wg.Add(1)
go task(&wg)
wg.Wait()
```

### Channel

Used when goroutines need to **communicate or send data**.

```go
ch := make(chan int)

go func() {
    ch <- 100
}()

result := <-ch
```

Later you will learn how these can work together.

---

# 🧠 Mental Model

Remember this simple model:

```text
                 WaitGroup
                    │
             ┌──────┴──────┐
             │             │
          Add(1)         Done()
             │             │
             │             │
       "Task started"  "Task finished"
             │             │
             └──────┬──────┘
                    ↓
              Counter = 0
                    ↓
                 Wait()
                    ↓
              Continue main
```

Or simply:

```text
Add    → "I'm starting work"
Done   → "I'm finished"
Wait   → "Tell me when everyone is finished"
```

---

# 🎯 What I Learned Today

* `sync.WaitGroup` is used to wait for multiple goroutines.
* `wg.Add(1)` increases the number of active tasks.
* `wg.Done()` decreases the counter.
* `wg.Wait()` blocks until the counter becomes `0`.
* `defer wg.Done()` is a common and safe pattern.
* A WaitGroup should normally be passed using a pointer.
* `WaitGroup` doesn't guarantee goroutine execution order.
* `WaitGroup` is better than `time.Sleep` for synchronization.
* `WaitGroup` doesn't protect shared data from race conditions.
* `WaitGroup` is mainly for **waiting**, while channels are mainly for **communication**.

---

# ⚡ Quick Reference

| Code                    | Meaning                         |
| ----------------------- | ------------------------------- |
| `var wg sync.WaitGroup` | Create WaitGroup                |
| `wg.Add(1)`             | Add one task                    |
| `wg.Done()`             | Mark one task complete          |
| `wg.Wait()`             | Wait for all tasks              |
| `&wg`                   | Pass WaitGroup address          |
| `defer wg.Done()`       | Call Done when function returns |

### Standard Pattern

```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)

    go func(i int) {
        defer wg.Done()

        // work
    }(i)
}

wg.Wait()
```

> **Goroutine starts the work.
> WaitGroup tracks the work.
> Done marks the work complete.
> Wait waits for everyone.**

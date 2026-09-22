# ⚡ Go Goroutines — Concurrency Basics

## 1. What Is a Goroutine?

A **goroutine** is a lightweight concurrent execution unit managed by the Go runtime.

You start a goroutine using the `go` keyword:

```go
go task(i)
```

Example:

```go
func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {
	go task(1)

	time.Sleep(time.Second)
}
```

The `go` keyword tells Go:

> Start this function concurrently and let `main()` continue.

---

# 2. Normal Function Call vs Goroutine

Normal function:

```go
task(1)
task(2)
task(3)
```

The calls execute one after another.

Conceptually:

```text
task 1
  ↓
task 2
  ↓
task 3
```

With goroutines:

```go
go task(1)
go task(2)
go task(3)
```

Go can execute them concurrently:

```text
       ┌── task 1
main ──┼── task 2
       └── task 3
```

The exact execution order is not guaranteed.

---

# 3. Your `task()` Function

You created:

```go
func task(id int) {
	fmt.Println("Doing task", id)
}
```

This is just a normal function.

When you write:

```go
go task(i)
```

it becomes a goroutine.

So:

```go
task(i)
```

means:

> Execute this function normally.

While:

```go
go task(i)
```

means:

> Start this function as a goroutine and continue without waiting for it to finish.

---

# 4. Your Main Loop

You have:

```go
for i := 0; i <= 10; i++ {
	go func(i int) {
		fmt.Println(i)
	}(i)
}
```

This creates **11 goroutines**:

```text
i = 0
i = 1
i = 2
...
i = 10
```

Each iteration starts another goroutine.

---

# 5. Anonymous Function

This part:

```go
func(i int) {
	fmt.Println(i)
}
```

is an **anonymous function**.

It has no name.

Normally:

```go
func printNumber(i int) {
	fmt.Println(i)
}
```

has a name:

```text
printNumber
```

But:

```go
func(i int) {
	fmt.Println(i)
}
```

doesn't.

---

# 6. Immediately Calling the Anonymous Function

You wrote:

```go
func(i int) {
	fmt.Println(i)
}(i)
```

The:

```go
(i)
```

at the end calls the anonymous function immediately.

It's similar to:

```go
func printNumber(i int) {
	fmt.Println(i)
}

printNumber(i)
```

The difference is that the anonymous function doesn't have a name.

---

# 7. The Important Part: `go func`

You have:

```go
go func(i int) {
	fmt.Println(i)
}(i)
```

This means:

1. Create an anonymous function.
2. Give it an `int` parameter.
3. Start it as a goroutine.
4. Pass the current `i` to it.

So:

```go
go func(i int) {
	fmt.Println(i)
}(i)
```

is essentially:

```text
Current loop value
       ↓
     pass to
       ↓
goroutine's own i
       ↓
   fmt.Println(i)
```

---

# 8. Why Pass `i` as a Parameter? ⭐

This is a very important concurrency concept.

You could write:

```go
for i := 0; i <= 10; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

But this can create problems around which loop variable the goroutine observes, depending on the Go version and exact code.

A very clear and robust pattern is:

```go
for i := 0; i <= 10; i++ {
	go func(i int) {
		fmt.Println(i)
	}(i)
}
```

Here, each goroutine receives its own function parameter value.

For example:

```text
Loop i = 0 → goroutine receives 0
Loop i = 1 → goroutine receives 1
Loop i = 2 → goroutine receives 2
```

This avoids accidentally depending on shared loop-variable state.

---

# 9. Why Is the Output Order Random?

You might expect:

```text
0
1
2
3
4
5
6
7
8
9
10
```

But goroutines are scheduled independently.

You might see:

```text
3
0
7
1
5
2
9
4
10
6
8
```

Another run might produce:

```text
0
2
1
5
3
4
8
6
7
9
10
```

There is no guarantee that goroutines will finish in the order they were started.

---

# 10. Concurrency vs Parallelism

These terms are related but different.

### Concurrency

Multiple tasks are **in progress** during overlapping periods.

```text
Task A ───────
     Task B ─────
          Task C ─────
```

### Parallelism

Multiple tasks are actually executing at the same time on different CPU cores.

```text
CPU Core 1 → Task A
CPU Core 2 → Task B
```

Go supports concurrency, and the runtime can execute goroutines in parallel when multiple CPU cores are available and `GOMAXPROCS` permits it.

---

# 11. Goroutine vs OS Thread

A goroutine is **not the same thing as an OS thread**.

Think of it roughly like:

```text
Your program
     ↓
Go Runtime
     ↓
Many Goroutines
     ↓
Scheduled onto OS threads
     ↓
CPU
```

Go's runtime manages goroutines and schedules them onto available operating-system threads.

That's one reason Go can handle very large numbers of concurrent tasks efficiently.

---

# 12. `main()` Doesn't Wait Automatically

This is extremely important.

Consider:

```go
func main() {
	go task(1)
}
```

The goroutine starts, but `main()` can finish immediately.

When the `main` goroutine exits, the entire program exits.

It does **not** automatically wait for other goroutines.

---

# 13. Your `time.Sleep()`

You used:

```go
time.Sleep(time.Second * 1)
```

This keeps `main()` alive for one second.

So your program gets time to let the goroutines execute.

Conceptually:

```text
main
 │
 ├── start goroutine
 ├── start goroutine
 ├── start goroutine
 │
 └── sleep 1 second
          ↓
       goroutines run
          ↓
       main exits
```

It works for a simple learning example.

But it is **not the proper synchronization technique** for production code.

---

# 14. Why `time.Sleep()` Is Not Ideal

Imagine:

```go
go verySlowTask()

time.Sleep(time.Second)
```

What if the task takes:

```text
1.5 seconds?
```

Your program exits too early.

What if it takes:

```text
100 milliseconds?
```

You unnecessarily waited almost another second.

So:

```go
time.Sleep(...)
```

doesn't actually tell us:

> "Wait until the goroutine finishes."

It only says:

> "Wait for this amount of time."

---

# 15. `sync.WaitGroup` — Proper Way ⭐

Go provides `sync.WaitGroup` for this kind of synchronization.

Example:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {

		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
```

Now:

```go
wg.Add(1)
```

means:

> We have one more goroutine to wait for.

And:

```go
wg.Done()
```

means:

> This goroutine has finished.

Finally:

```go
wg.Wait()
```

means:

> Wait until all registered goroutines finish.

---

# 16. `defer wg.Done()`

This:

```go
defer wg.Done()
```

is a very common pattern.

It means `Done()` will execute when the goroutine function returns.

So:

```go
go func(i int) {

	defer wg.Done()

	fmt.Println(i)

}(i)
```

is safer than manually remembering to call:

```go
wg.Done()
```

at every possible exit point.

---

# 17. Mental Model of WaitGroup

Think of it like a counter:

```text
wg.Add(1) → 1
wg.Add(1) → 2
wg.Add(1) → 3

goroutine finishes
wg.Done() → 2

goroutine finishes
wg.Done() → 1

goroutine finishes
wg.Done() → 0

wg.Wait() returns
```

So:

```text
Add → increase work
Done → decrease work
Wait → wait until zero
```

---

# 18. Goroutines Can Run Functions

You can start:

```go
go task(10)
```

You can also start an anonymous function:

```go
go func() {
	fmt.Println("Hello")
}()
```

And with parameters:

```go
go func(name string) {
	fmt.Println(name)
}("Om")
```

---

# 19. Goroutines + Closures

This connects directly to your previous **closure** lesson.

A closure can capture variables:

```go
name := "Om"

go func() {
	fmt.Println(name)
}()
```

The anonymous function accesses the surrounding variable.

This is useful, but when variables are modified or shared between goroutines, you need to think carefully about synchronization and data races.

Passing values as parameters is often clearer:

```go
go func(name string) {
	fmt.Println(name)
}(name)
```

---

# 20. Goroutines + Functions

Your earlier function:

```go
func task(id int) {
	fmt.Println("Doing task", id)
}
```

can directly become a goroutine:

```go
go task(1)
```

You don't need an anonymous function unless you need some extra logic around the call.

For example:

```go
for i := 0; i <= 10; i++ {
	go task(i)
}
```

is simpler if all you need is to call `task`.

---

# 21. Why Use Anonymous Functions?

Anonymous goroutines are useful when you need a small piece of logic:

```go
go func() {
	fmt.Println("Processing...")
}()
```

Or when passing the loop value explicitly:

```go
go func(i int) {
	fmt.Println("Processing", i)
}(i)
```

But don't create anonymous functions unnecessarily.

---

# 22. Real-World Example — Processing Tasks

Imagine you have 5 files:

```text
file1
file2
file3
file4
file5
```

Instead of processing them one by one:

```go
process(file1)
process(file2)
process(file3)
process(file4)
process(file5)
```

you could run independent work concurrently:

```go
for _, file := range files {
	go process(file)
}
```

For real applications, you'd normally also use proper synchronization and consider limits on concurrency.

---

# 23. Real-World Example — HTTP Requests

Suppose you need to call multiple independent services:

```text
Service A
Service B
Service C
```

You could potentially run the requests concurrently:

```text
             ┌── Service A
Appl
```

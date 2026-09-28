# Go Channels, Select & Buffered Channels

## 1. Introduction to Channels

A **channel** is a communication mechanism that allows goroutines to send and receive data.

Channels help goroutines communicate safely without directly sharing data.

For example, one goroutine can send a number, and another goroutine can receive it.

```go
ch := make(chan int)
```

Here:

- `chan` is the channel keyword.
    
- `int` is the type of data the channel can carry.
    
- `make()` creates the channel.
    

A channel can carry values of a specific type.

```go
numChan := make(chan int)
messageChan := make(chan string)
```

You cannot send a string to an integer channel.

---

## 2. Sending and Receiving Data

There are two main operations with channels.

**Send data:**

```go
ch <- 10
```

**Receive data:**

```go
num := <-ch
```

The `<-` operator is used for both sending and receiving.

- `ch <- 10`: Send `10` into the channel.
    
- `num := <-ch`: Receive a value from the channel and assign it to `num`.
    

### Example

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello, Go"
    }()

    message := <-ch
    fmt.Println(message)
}
```

Output:

```text
Hello, Go
```

The goroutine sends a message, and the main goroutine receives it.

---

## 3. Channels Are Blocking

Unbuffered channels are blocking.

This means:

- A send blocks until another goroutine receives the value.
    
- A receive blocks until another goroutine sends a value.
    

Example:

```go
ch := make(chan string)

ch <- "Ping"
```

This code blocks because there is no receiver ready to receive the message.

If the main goroutine is the only goroutine and gets stuck like this, Go reports a deadlock.

### Correct example

```go
package main

import "fmt"

func main() {
    messageChan := make(chan string)

    go func() {
        messageChan <- "Ping"
    }()

    msg := <-messageChan
    fmt.Println(msg)
}
```

Output:

```text
Ping
```

The sender and receiver synchronize with each other.

**Mental model:**

```text
Sender                         Receiver
   |                                |
   | ------ Send "Ping" ----------> |
   |                                |
   |                          Receive value
   |                                |
   | <------ Both continue -------- |
```

An unbuffered channel acts like a handoff: the sender and receiver must meet.

---

## 4. Processing Numbers with Channels

Your example:

```go
func proccessNum(numChan *chan int) {
    fmt.Println("Proccessing number", <-numChan)
}
```

The idea is correct, but the channel parameter should usually be `chan int`, not `*chan int`.

Correct version:

```go
func processNum(numChan chan int) {
    fmt.Println("Processing number", <-numChan)
}
```

Complete example:

```go
package main

import "fmt"

func processNum(numChan chan int) {
    fmt.Println("Processing number", <-numChan)
}

func main() {
    numChan := make(chan int)

    go processNum(numChan)

    numChan <- 5
}
```

Output:

```text
Processing number 5
```

The main goroutine sends `5`, and the other goroutine receives and processes it.

### Why Don't We Need `*chan int`?

A channel value already refers to the channel's internal data structure. Passing a channel to a function allows that function to send and receive values.

Usually, use:

```go
func processNum(ch chan int)
```

Instead of:

```go
func processNum(ch *chan int)
```

---

## 5. Returning Results Through a Channel

Channels can also be used to send results from one goroutine to another.

Your example:

```go
func sum(result *chan int, num1 int, num2 int) {
    numResult := num1 + num2
    result <- numResult
}
```

Corrected version:

```go
func sum(result chan int, num1 int, num2 int) {
    numResult := num1 + num2
    result <- numResult
}
```

Complete example:

```go
package main

import "fmt"

func sum(result chan int, num1 int, num2 int) {
    numResult := num1 + num2
    result <- numResult
}

func main() {
    sumNum := make(chan int)

    go sum(sumNum, 4, 5)

    res := <-sumNum
    fmt.Println(res)
}
```

Output:

```text
9
```

### How it works

1. The main goroutine creates a channel.
    
2. It starts the `sum()` function as a goroutine.
    
3. The function calculates `4 + 5`.
    
4. It sends the result to the channel.
    
5. The main goroutine receives the result.
    
6. The result is printed.
    

This is useful when a goroutine performs work and needs to return its result.

---

## 6. Channels as a Goroutine Synchronizer

A channel can also be used to signal that a goroutine has finished.

Your example:

```go
func task(done *chan bool) {
    defer func() {
        done <- true
    }()

    fmt.Println("Processing...")
}
```

Corrected version:

```go
func task(done chan bool) {
    defer func() {
        done <- true
    }()

    fmt.Println("Processing...")
}
```

Complete example:

```go
package main

import "fmt"

func task(done chan bool) {
    defer func() {
        done <- true
    }()

    fmt.Println("Processing...")
}

func main() {
    done := make(chan bool)

    go task(done)

    <-done

    fmt.Println("Task completed")
}
```

Output:

```text
Processing...
Task completed
```

### Why does `<-done` block?

The main goroutine waits for a value from the `done` channel.

The task goroutine sends `true` when it finishes.

When the value is received, the main goroutine continues.

```text
Main goroutine              Task goroutine
      |                           |
      | ------ Start task ------> |
      |                           |
      |                       Processing...
      |                           |
      |                       Send true
      |                           |
      | <------- true ----------- |
      |                           |
  Continue
```

This is a simple way to synchronize goroutines.

For waiting for many goroutines, `sync.WaitGroup` is often more convenient. A channel is especially useful when you also need to communicate data or signals.

---

## 7. Buffered Channels

A buffered channel can hold a limited number of values without requiring a receiver for every send immediately.

Syntax:

```go
ch := make(chan int, 3)
```

Here, `3` is the channel's capacity.

It can hold up to three values before a sender blocks.

Example:

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)

    ch <- 10
    ch <- 20
    ch <- 30

    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

Output:

```text
10
20
30
```

The first three sends can complete without a receiver because the channel has enough capacity.

But sending a fourth value before receiving any values would block:

```go
ch := make(chan int, 3)

ch <- 10
ch <- 20
ch <- 30
ch <- 40 // Blocks
```

The buffer is full.

### Unbuffered vs Buffered Channels

|Feature|Unbuffered|Buffered|
|---|---|---|
|Creation|`make(chan int)`|`make(chan int, 3)`|
|Capacity|0|Specified capacity|
|Send|Waits for a receiver|Can proceed if buffer has space|
|Receive|Waits for a sender|Can proceed if buffer has data|
|Common use|Synchronization and handoff|Queues and asynchronous work|

A buffered channel does not make goroutines automatically faster. It changes when senders and receivers block.

---

## 8. Email Sender Using a Buffered Channel

Your email example demonstrates a practical use of buffered channels.

```go
func emailSender(emailChan chan string, done chan bool) {
    defer func() {
        done <- true
    }()

    for email := range emailChan {
        fmt.Println("Sending email to", email)
        time.Sleep(time.Second)
    }
}
```

The sender receives email addresses from the channel and processes them one by one.

Complete example:

```go
package main

import (
    "fmt"
    "time"
)

func emailSender(emailChan chan string, done chan bool) {
    defer func() {
        done <- true
    }()

    for email := range emailChan {
        fmt.Println("Sending email to", email)
        time.Sleep(time.Second)
    }
}

func main() {
    emailChan := make(chan string, 100)
    done := make(chan bool)

    go emailSender(emailChan, done)

    for i := 0; i < 100; i++ {
        emailChan <- fmt.Sprintf("%d@gmail.com", i)
    }

    fmt.Println("Done sending emails to the channel")

    close(emailChan)

    <-done

    fmt.Println("Email sender finished")
}
```

### How it works

1. Create a buffered channel with capacity `100`.
    
2. Start the email sender goroutine.
    
3. Send 100 email addresses into the channel.
    
4. The sender receives and processes emails.
    
5. Close the email channel to indicate that no more emails will be sent.
    
6. Wait for the sender to finish using `done`.
    

**Important:** `Done sending emails to the channel` means all email addresses have been queued. It does not mean all emails have been processed.

The sender processes them one by one in this example.

---

## 9. Closing a Channel

Use `close()` to indicate that no more values will be sent to a channel.

```go
close(emailChan)
```

After closing a channel:

- Receivers can still receive values that were already buffered.
    
- Once all buffered values are received, a receive returns the zero value immediately.
    
- Sending to a closed channel causes a panic.
    
- Closing an already closed channel causes a panic.
    

Example:

```go
ch := make(chan int, 2)

ch <- 10
ch <- 20

close(ch)

fmt.Println(<-ch)
fmt.Println(<-ch)
```

Output:

```text
10
20
```

### Important rule

Usually, the goroutine that sends values should be responsible for closing the channel.

Do not close a channel just because you have finished receiving from it.

---

## 10. Receiving with `range`

You can use `range` to receive values from a channel until it is closed.

Example:

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)

    ch <- 10
    ch <- 20
    ch <- 30

    close(ch)

    for num := range ch {
        fmt.Println(num)
    }
}
```

Output:

```text
10
20
30
```

The loop ends when the channel is closed and all buffered values have been received.

**Important:** If the channel is never closed, a `range` loop can wait forever after all values have been received.

---

## 11. Channel Direction

Go allows you to specify whether a function can send or receive from a channel.

This is called **channel direction**.

### Send-only channel

```go
func sendData(ch chan<- int) {
    ch <- 10
}
```

`chan<- int` means the function can only send integers.

### Receive-only channel

```go
func receiveData(ch <-chan int) {
    num := <-ch
    fmt.Println(num)
}
```

`<-chan int` means the function can only receive integers.

### Why use channel direction?

It makes functions easier to understand and prevents accidental operations.

For example, a sender should not need to receive from its output channel.

A receiver should not need to send values into its input channel.

---

## 12. The `select` Statement

The `select` statement lets a goroutine wait on multiple channel operations.

It is similar to `switch`, but its cases are channel operations.

Example:

```go
package main

import "fmt"

func main() {
    chan1 := make(chan int)
    chan2 := make(chan string)

    go func() {
        chan1 <- 10
        chan2 <- "Golang"
    }()

    for i := 0; i < 2; i++ {
        select {
        case value := <-chan1:
            fmt.Println("Received from channel 1:", value)

        case value := <-chan2:
            fmt.Println("Received from channel 2:", value)
        }
    }
}
```

Possible output:

```text
Received from channel 2: Golang
Received from channel 1: 10
```

The order may be different.

### How `select` works

- Go checks which channel operations can proceed.
    
- If one or more cases are ready, one ready case is selected.
    
- If no case is ready, the goroutine blocks.
    
- If multiple cases are ready, Go selects one pseudo-randomly.
    
- A `default` case runs immediately if no communication is ready.
    

In your example, `select` receives two values, one from each channel, over two iterations.

However, it does not guarantee which channel will be selected first.

### `select` with `default`

```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("No message available")
}
```

This does not block when the channel is not ready. It executes the `default` case instead.

Be careful: a `default` case in a loop can cause busy waiting if there is no pause or other blocking operation.

---

## 13. `select` with Multiple Goroutines

Imagine you have multiple goroutines sending results to different channels.

```go
select {
case result := <-apiChan:
    fmt.Println("API result:", result)

case result := <-dbChan:
    fmt.Println("Database result:", result)
}
```

The goroutine can receive from whichever channel becomes ready.

This is useful for:

- Handling multiple workers
    
- Listening to different events
    
- Processing results from concurrent operations
    
- Implementing timeouts with `time.After`
    

### Example: Timeout

```go
select {
case result := <-ch:
    fmt.Println("Received:", result)

case <-time.After(2 * time.Second):
    fmt.Println("Operation timed out")
}
```

If no value arrives within two seconds, the timeout case is selected.

---

## 14. Important Corrections to My Code

### Correction 1: Don't use `*chan`

Instead of:

```go
func processNum(numChan *chan int)
```

Use:

```go
func processNum(numChan chan int)
```

A channel can be passed directly to a function.

### Correction 2: `make()` needs the channel type

Incorrect:

```go
chan1 := make(chan int)
```

This is actually correct for an unbuffered channel.

For a buffered channel, specify the capacity:

```go
chan1 := make(chan int, 10)
```

### Correction 3: Don't send without a receiver on an unbuffered channel

This can deadlock:

```go
ch := make(chan int)
ch <- 10
```

Use a receiving goroutine or a buffered channel with available capacity.

### Correction 4: Close the channel after sending is complete

For your email sender, this is correct:

```go
close(emailChan)
<-done
```

The channel is closed after all email addresses have been sent. The `done` signal lets `main()` wait until the sender has finished processing them.

### Correction 5: `select` does not guarantee order

If multiple channel cases are ready, the selected case is not guaranteed to be the first one in your source code.

---

## 15. Common Mistakes

|Mistake|Explanation|
|---|---|
|Sending to an unbuffered channel without a receiver|The sender blocks|
|Receiving when no value will arrive|The receiver blocks|
|Sending to a closed channel|Causes a panic|
|Closing a channel more than once|Causes a panic|
|Closing a channel while another goroutine is sending|Can cause a panic|
|Using `range` on a channel that is never closed|May wait forever|
|Assuming `select` cases run in order|Incorrect; ready cases are selected without a guaranteed order|
|Thinking buffered channels never block|They block when the buffer is full|
|Using `*chan` unnecessarily|Usually, a channel value is sufficient|

---

## 16. Real-World Use Cases

### Buffered channels as queues

A buffered channel can temporarily store tasks while a worker processes them.

```text
Producer
   |
   v
Buffered Channel
   |
   v
Worker
   |
   v
Process task
```

### Worker pools

Multiple workers can receive jobs from the same channel and process them concurrently.

### Notifications

A channel can signal that a task has finished or that an event has occurred.

### API requests

A program can use `select` to wait for API results or a timeout.

### DevOps and infrastructure tools

Channels can help coordinate:

- Concurrent health checks
    
- Background monitoring jobs
    
- Log processing
    
- Parallel infrastructure tasks
    
- Worker pools
    

---

## 17. Channels vs WaitGroup

|Feature|Channel|WaitGroup|
|---|---|---|
|Main purpose|Communication and synchronization|Waiting for goroutines|
|Send data|Yes|No|
|Receive data|Yes|No|
|Wait for tasks|Possible with signals|Yes|
|Carry results|Yes|No|
|Track task count|No built-in counter|Yes|

Use a `WaitGroup` when you only need to wait for goroutines to finish.

Use a channel when goroutines need to exchange data or send signals.

You can also use both together when a program needs both communication and task tracking.

---

## 18. What I Learned Today

- Channels allow goroutines to communicate.
    
- `make(chan int)` creates an unbuffered integer channel.
    
- `ch <- value` sends data.
    
- `<-ch` receives data.
    
- Unbuffered channel sends and receives block until they can communicate.
    
- Channels can return results from goroutines.
    
- Channels can be used to signal task completion.
    
- Buffered channels can hold values up to their capacity.
    
- `close(ch)` indicates that no more values will be sent.
    
- `range` can receive values until a channel is closed and drained.
    
- `select` waits on multiple channel operations.
    
- `default` allows a `select` to continue without blocking.
    
- Channel directions restrict functions to sending or receiving.
    
- Channels and WaitGroups solve related but different problems.
    

---

## 19. Quick Reference

|Syntax|Meaning|
|---|---|
|`make(chan int)`|Create an unbuffered channel|
|`make(chan int, 5)`|Create a buffered channel with capacity 5|
|`ch <- 10`|Send 10|
|`num := <-ch`|Receive a value|
|`close(ch)`|Close a channel|
|`for v := range ch`|Receive until the channel is closed and drained|
|`select`|Wait on multiple channel operations|
|`default`|Run if no `select` case is ready|
|`chan<- int`|Send-only channel|
|`<-chan int`|Receive-only channel|
|`cap(ch)`|Get channel capacity|
|`len(ch)`|Get the number of buffered values currently queued|

---

## 20. Learning Roadmap

-  Goroutines
    
-  WaitGroup
    
-  Channels
    
-  Blocking and synchronization
    
-  Buffered channels
    
-  Channel closing
    
-  `select`
    
-  Channel direction
    
-  Mutex and RWMutex
    
-  Data races and race detector
    
-  Context
    
-  Worker pools
    
-  Fan-in and fan-out
    
-  Concurrency patterns
    

---

## Final Mental Model

```text
Goroutine
    |
    | communicates using
    v
  Channel
    |
    | can send and receive
    v
  Receiver

Unbuffered channel:
Sender and receiver must meet.

Buffered channel:
Values can wait in a limited queue.

select:
Wait for one of several channel operations.

WaitGroup:
Wait for multiple goroutines to finish.
```

**Remember:** Channels are not just containers for data. They are also a way for goroutines to coordinate their work.
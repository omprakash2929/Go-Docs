# 🐹 Go Introduction

## What is Go?

Go (Golang) is an **open-source, statically typed programming language** created by Google.

It is designed for building **fast, reliable, and scalable applications**, especially backend services, APIs, cloud infrastructure, and microservices.

---

## 🚀 Key Features of Go

* ⚡ **Fast** — Compiled directly to machine code.
* 🔒 **Statically Typed** — Types are checked at compile time.
* 🧵 **Concurrency** — Goroutines make concurrent programming easier.
* 📦 **Simple** — Clean and easy-to-read syntax.
* 🛠️ **Built-in Tooling** — Formatting, testing, and dependency management are built into Go.
* 🌐 **Cross-platform** — Go programs can run on different operating systems.
* 📈 **Scalable** — Well suited for backend systems, cloud infrastructure, and microservices.

---

## 💡 Benefits of Go

Go is commonly used for:

* Backend development
* REST APIs
* Microservices
* Cloud applications
* DevOps tools
* Kubernetes-related tools
* Networking applications
* CLI tools

---

# 👋 My First Go Program

Create a file:

```text
main.go
```

Write:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run it:

```bash
go run main.go
```

Output:

```text
Hello, World!
```

---

## 📦 Understanding `fmt.Println()`

`fmt` is a **standard library package** in Go.

`Println()` is a function inside the `fmt` package.

```go
fmt.Println("Hello")
```

Think of it as:

```text
fmt
 └── Println()
```

We use `fmt.Println()` instead of just `Println()` because `Println` belongs to the `fmt` package.

---

## 🧠 What I Learned Today

* What Go is
* Why Go is useful
* Main features and benefits of Go
* How to create a Go program
* `package main`
* `func main()`
* How to import a package
* How `fmt.Println()` works
* How to run a Go program using `go run`

---

## 📝 Quick Note

```go
package main       // Defines the main package

import "fmt"       // Imports the fmt package

func main() {      // Program execution starts here
    fmt.Println("Hello, World!")
}
```

# Go Packages and Scope

## 1. What Are Packages in Go?

A package is a way to organize related Go code into separate files and directories.

Packages help us:

* Organize large projects.
* Reuse code across different files.
* Separate responsibilities.
* Keep code maintainable.
* Control which functions and variables are accessible outside a package.

For example, in a web application, we can create separate packages for authentication, users, payments, and database operations.

## 2. Go Project Structure

Today's project structure:

```text
Day24_Packages/
│
├── auth/
│   ├── credentials.go
│   └── session.go
│
├── user/
│   └── user.go
│
├── go.mod
├── go.sum
└── main.go
```

Each directory can contain a separate Go package.

| File                  | Responsibility                  |
| --------------------- | ------------------------------- |
| `main.go`             | Application entry point         |
| `auth/credentials.go` | Authentication functions        |
| `auth/session.go`     | Session-related functions       |
| `user/user.go`        | User data structure             |
| `go.mod`              | Module name and dependencies    |
| `go.sum`              | Dependency checksum information |

**Important:** Go files in the same directory generally belong to the same package. A package can contain multiple `.go` files.

## 3. The `main` Package

The `main` package is special in Go.

It is the entry point for an executable application.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

The `main()` function runs when you execute the program.

A reusable library package does not need a `main()` function.

## 4. Creating Your Own Packages

For example, inside `auth/session.go`:

```go
package auth

func extractSession() string {
    return "Logged in"
}

func GetSession() string {
    return extractSession()
}
```

Here, `auth` is the package name.

We have two functions:

* `extractSession()` is private to the `auth` package.
* `GetSession()` is exported and can be used by other packages.

### Exported and Unexported Identifiers

Go uses the first letter of an identifier to control whether it is exported.

| Identifier         | Accessibility |
| ------------------ | ------------- |
| `Login()`          | Exported      |
| `GetSession()`     | Exported      |
| `extractSession()` | Unexported    |
| `password`         | Unexported    |
| `User`             | Exported      |
| `Email`            | Exported      |

An identifier beginning with an uppercase letter is exported. An identifier beginning with a lowercase letter is unexported.

This applies to functions, types, variables, constants, and struct fields.

**Important:** Unexported does not mean private to one file. Other files in the same package can access unexported identifiers.

## 5. Understanding Scope

Scope defines where a variable, function, or identifier can be accessed.

Go has different scopes, including:

### Local Scope

A variable declared inside a function is generally accessible only within that function.

```go
func main() {
    message := "Hello Go"
    fmt.Println(message)
}
```

The `message` variable cannot be accessed directly from another function.

### Package Scope

Declarations made at the top level of a package can be accessed by other files in the same package, subject to whether they are exported when accessed from another package.

```go
package auth

var appName = "My App"

func GetAppName() string {
    return appName
}
```

Other files in the `auth` package can use `appName`.

Code in another package cannot access `auth.appName` because it is unexported.

### Block Scope

Variables declared inside an `if`, `for`, or other block are limited to that block.

```go
func main() {
    if true {
        message := "Inside if"
        fmt.Println(message)
    }

    // message is not accessible here
}
```

## 6. Importing Packages

The `import` keyword lets us use code from other packages.

Example:

```go
import (
    "fmt"
    "github.com/fatih/color"
    "github.com/omprakash2929/todoaspp/auth"
    "github.com/omprakash2929/todoaspp/user"
)
```

Here:

* `fmt` is part of Go's standard library.
* `github.com/fatih/color` is an external package.
* `auth` and `user` are packages from your own module.

The import path for your own packages depends on the module name in `go.mod`.

For example:

```go
module github.com/omprakash2929/todoaspp

go 1.25
```

With this module path, the `auth` package can be imported using:

```go
import "github.com/omprakash2929/todoaspp/auth"
```

The Go version in your own `go.mod` should match the version you use.

## 7. Using Your Own Packages

Example from `main.go`:

```go
package main

import (
    "fmt"
    "github.com/omprakash2929/todoaspp/auth"
)

func main() {
    session := auth.GetSession()
    fmt.Println(session)
}
```

The `auth` package is imported, and its exported `GetSession()` function is called using the package name.

```go
auth.GetSession()
```

You cannot call an unexported function from another package:

```go
auth.extractSession() // Error
```

Because `extractSession()` starts with a lowercase letter.

## 8. Structs Across Packages

You can define a struct in one package and use it in another package.

Inside `user/user.go`:

```go
package user

type User struct {
    Email string
    Name  string
}
```

Inside `main.go`:

```go
package main

import (
    "fmt"
    "github.com/omprakash2929/todoaspp/user"
)

func main() {
    userData := user.User{
        Email: "omprakash@gmail.com",
        Name:  "Omprakash",
    }

    fmt.Println(userData)
}
```

Important points:

* `User` is exported, so another package can use it.
* `Email` and `Name` are exported fields.
* The struct can be created using `user.User{}`.
* If the fields were lowercase, another package could not initialize them directly.

## 9. External Packages

Go supports third-party packages that provide functionality beyond the standard library.

Today's example uses:

```go
"github.com/fatih/color"
```

Example:

```go
package main

import (
    "github.com/fatih/color"
)

func main() {
    color.Blue("Hello Go")
    color.Yellow("Hello Python")
    color.HiRed("Hello World")
}
```

The package provides functions to print colored terminal text.

The comment in your original code says `HiRed` prints cyan, but `HiRed` is a red color style.

## 10. `go mod init`, `go get`, and `go mod tidy`

### `go mod init`

Creates a new Go module.

```bash
go mod init github.com/omprakash2929/todoaspp
```

This creates the `go.mod` file.

Run it from your project's root directory.

### `go get`

Adds or updates a dependency.

```bash
go get github.com/fatih/color
```

This downloads the package and records the dependency in your module.

### `go mod tidy`

Cleans up your module dependencies.

```bash
go mod tidy
```

It:

* Adds dependencies needed by your code.
* Removes dependencies that are no longer needed.
* Updates `go.mod` and `go.sum`.

**Pro tip:** Run `go mod tidy` after adding or removing imports, especially before committing your project.

## 11. Private Functions and Encapsulation

Your `auth` example demonstrates a useful design pattern.

```go
package auth

func extractSession() string {
    return "Logged in"
}

func GetSession() string {
    return extractSession()
}
```

The internal function is hidden from other packages, while the exported function provides a public way to access the functionality.

This is a simple form of encapsulation.

For example, you could change how sessions are extracted internally without changing the public function used by other packages.

Note that Go's unexported identifiers are package-level access control, not a security boundary.

## 12. Corrected Version of Your Example

Your original code calls `auth.LoginWithCredentials()`, but that function was not included in the shown `auth` code. It must be implemented and exported for the example to compile.

Also, `extractSession()` should return a `string`, not `*string`, if it returns a string literal directly.

### `auth/session.go`

```go
package auth

func extractSession() string {
    return "Logged in"
}

func GetSession() string {
    return extractSession()
}
```

### `auth/credentials.go`

```go
package auth

import "fmt"

func LoginWithCredentials(username, password string) {
    fmt.Println("Logging in:", username)
}
```

This is only a learning example. A real authentication system should securely verify credentials and avoid printing sensitive information.

### `user/user.go`

```go
package user

type User struct {
    Email string
    Name  string
}
```

### `main.go`

```go
package main

import (
    "fmt"

    "github.com/fatih/color"
    "github.com/omprakash2929/todoaspp/auth"
    "github.com/omprakash2929/todoaspp/user"
)

func main() {
    auth.LoginWithCredentials("omprakash", "123456")

    session := auth.GetSession()
    fmt.Println(session)

    userData := user.User{
        Email: "omprakash@gmail.com",
        Name:  "omprakash",
    }

    color.Blue(userData.Email)
    color.Yellow(userData.Name)
    color.HiRed("Hello from Go")
}
```

Notice that `userData` is used instead of `user` as the variable name. This avoids shadowing the imported `user` package.

## 13. Common Mistakes

| Mistake                                             | Explanation                                                       |
| --------------------------------------------------- | ----------------------------------------------------------------- |
| Calling a lowercase function from another package   | Unexported functions cannot be accessed outside their package.    |
| Using the wrong import path                         | The import path must match the module path and package directory. |
| Forgetting `go mod init`                            | A project using modules needs a module definition.                |
| Returning a string where `*string` is expected      | The return type must match the function signature.                |
| Naming a variable the same as an imported package   | It can shadow the package name in its scope.                      |
| Assuming files in different folders share a package | Each directory normally represents a separate package.            |

## 14. Important Package Design Tips

* Keep packages focused on one responsibility.
* Use clear, meaningful package names.
* Avoid creating too many tiny packages without a reason.
* Export only the identifiers that other packages need.
* Use lowercase names for internal implementation details.
* Prefer small public APIs that are easy to understand.
* Avoid circular imports: packages cannot import each other in a cycle.

## 15. What I Learned Today

* A package organizes related Go code.
* The `main` package is the entry point of an executable.
* Each directory normally represents a separate package.
* `import` lets us use code from other packages.
* Uppercase identifiers are exported.
* Lowercase identifiers are unexported.
* Scope controls where variables and identifiers can be accessed.
* `go.mod` defines the module and its dependencies.
* `go get` adds dependencies.
* `go mod tidy` cleans up dependencies.
* Structs and functions can be shared across packages when exported.
* Package boundaries help organize code and hide implementation details.

## 16. Quick Reference

| Command / Syntax        | Purpose                                  |
| ----------------------- | ---------------------------------------- |
| `package main`          | Defines the executable package           |
| `import "fmt"`          | Imports a package                        |
| `func GetSession()`     | Exported function                        |
| `func extractSession()` | Unexported function                      |
| `type User struct{}`    | Declares a struct                        |
| `user.User{}`           | Creates a struct from the `user` package |
| `go mod init <module>`  | Initializes a module                     |
| `go get <package>`      | Adds or updates a dependency             |
| `go mod tidy`           | Cleans dependency definitions            |
| `go run .`              | Runs the current module's main package   |

## 17. Mental Model

Think of a Go project like a company:

* **Module:** The entire company or project.
* **Package:** A department with a specific responsibility.
* **Exported identifiers:** Public services that other departments can use.
* **Unexported identifiers:** Internal implementation details.
* **Scope:** The area in which a name can be accessed.

For example, the `main` package can call `auth.GetSession()`, but it cannot directly call `auth.extractSession()`.

---

**Next topics to learn:**

* Go error handling
* `errors` package
* Creating custom errors
* Building a small CLI application using multiple packages

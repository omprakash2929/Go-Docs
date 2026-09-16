# 👉 Go Pointers

## 1. What is a Pointer?

A **pointer** is a variable that stores the **memory address of another variable**.

Normally:

```go
num := 1
```

`num` stores the value:

```text
num → 1
```

A pointer stores the address of `num`:

```text
pointer → address of num
                 ↓
                1
```

Pointers are useful when we want a function to **modify the original value**.

---

# 2. Two Important Operators

Go mainly uses two operators with pointers:

### `&` → Address of

`&` gives the memory address of a variable.

```go
num := 1

fmt.Println(&num)
```

This prints an address similar to:

```text
0xc0000120a0
```

The exact address will be different each time/program.

Think:

```text
&num
 ↓
address of num
```

---

### `*` → Dereference

`*` is used to access the value stored at an address.

Example:

```go
num := 1

ptr := &num

fmt.Println(*ptr)
```

Output:

```text
1
```

Think:

```text
ptr
 ↓
address of num
 ↓
*ptr
 ↓
1
```

---

# 3. Your Example

Your code:

```go
package main

import "fmt"

func changeNum(num *int) {

    *num = 5

    fmt.Println("In Change num:", *num)
}

func main() {

    num := 1

    changeNum(&num)

    fmt.Println("in main function", num)
}
```

Output:

```text
In Change num: 5
in main function 5
```

The important question is:

> How did `changeNum()` change `num` inside `main()`?

Because we passed the **address of `num`**.

---

# 4. Step-by-Step

First:

```go
num := 1
```

We have:

```text
num
 │
 ▼
 1
```

Then:

```go
changeNum(&num)
```

`&num` means:

> Give the address of `num`.

So we pass the address to the function.

---

# 5. Function Parameter `*int`

The function is:

```go
func changeNum(num *int)
```

Here:

```go
num *int
```

means:

> `num` is a pointer to an `int`.

It doesn't contain the actual integer directly.

Conceptually:

```text
num pointer
    │
    ▼
 address
    │
    ▼
 original integer
```

---

# 6. Dereferencing the Pointer

Inside the function:

```go
*num = 5
```

The `*` means:

> Go to the address stored in `num` and access the value there.

So:

```text
*num
 ↓
original variable
 ↓
1
```

Then:

```go
*num = 5
```

changes:

```text
1 → 5
```

The original variable in `main()` is changed.

---

# 7. Why Does `main()` See `5`?

After:

```go
changeNum(&num)
```

the original `num` has been modified.

So:

```go
fmt.Println(num)
```

prints:

```text
5
```

Memory concept:

```text
main()

num
 │
 ▼
┌─────────┐
│    5    │
└─────────┘
    ▲
    │
    │ pointer
    │
changeNum()
```

Both are referring to the same underlying variable.

---

# 8. `&` vs `*`

This is the most important thing to remember.

| Operator | Meaning                               |
| -------- | ------------------------------------- |
| `&`      | Get address                           |
| `*`      | Dereference / access value at address |

Example:

```go
num := 10

ptr := &num

fmt.Println(ptr)  // address
fmt.Println(*ptr) // 10
```

Think:

```text
&num
 ↓
address


*ptr
 ↓
value at that address
```

---

# 9. Changing Value Through Pointer

Example:

```go
num := 10

ptr := &num

*ptr = 20

fmt.Println(num)
```

Output:

```text
20
```

Why?

Because `ptr` points to `num`.

```text
num = 10

ptr
 ↓
num


*ptr = 20

num = 20
```

---

# 10. Why Use Pointers?

One major reason is to allow a function to modify the original variable.

Without pointer:

```go
func changeNum(num int) {
    num = 5
}

func main() {
    num := 1

    changeNum(num)

    fmt.Println(num)
}
```

Output:

```text
1
```

Why?

Because the function receives a copy of the value.

Conceptually:

```text
main:
num = 1

      ↓ copy

function:
num = 1

function changes its copy:

num = 5

main's num:
1
```

---

# 11. With Pointer

Now:

```go
func changeNum(num *int) {
    *num = 5
}
```

Call:

```go
changeNum(&num)
```

Now the function receives the address of the original variable.

```text
main:
num = 1
  │
  │ address
  ▼
function:
pointer → num
           │
           ▼
           1

*num = 5

           ▼
           5
```

So the original value changes.

---

# 12. Pointer Declaration

You can declare a pointer:

```go
var ptr *int
```

This means:

> `ptr` can store the address of an `int`.

Initially:

```text
ptr → nil
```

A pointer with no address is called a **nil pointer**.

---

# 13. Assigning an Address

Example:

```go
num := 10

var ptr *int

ptr = &num
```

Now:

```text
ptr
 ↓
address of num
```

You can access the value:

```go
fmt.Println(*ptr)
```

Output:

```text
10
```

---

# 14. Pointer Example

```go
package main

import "fmt"

func main() {
    num := 10

    ptr := &num

    fmt.Println(num)
    fmt.Println(ptr)
    fmt.Println(*ptr)
}
```

Possible output:

```text
10
0xc0000120a0
10
```

Remember that the address is not fixed.

---

# 15. Modify Through Pointer

```go
package main

import "fmt"

func main() {
    num := 10

    ptr := &num

    *ptr = 50

    fmt.Println(num)
}
```

Output:

```text
50
```

Because:

```go
*ptr = 50
```

changes the original variable.

---

# 16. Pointer to Struct

Pointers are also commonly used with structs.

Example:

```go
type User struct {
    name string
    age  int
}

func changeAge(user *User) {
    user.age = 25
}
```

Call:

```go
user := User{
    name: "Omprakash",
    age: 21,
}

changeAge(&user)
```

Now:

```go
fmt.Println(user.age)
```

Output:

```text
25
```

Go automatically allows:

```go
user.age
```

instead of manually writing:

```go
(*user).age
```

---

# 17. Pointer to Pointer

Go can also have a pointer that points to another pointer.

Example:

```go
num := 10

ptr := &num

ptr2 := &ptr
```

Conceptually:

```text
ptr2
 ↓
ptr
 ↓
num
 ↓
10
```

Then:

```go
fmt.Println(**ptr2)
```

prints:

```text
10
```

This is more advanced and usually isn't needed for basic Go programming.

---

# 18. Important: Go Is Pass-by-Value

This is an important Go concept.

Go passes function arguments **by value**.

Even when using pointers, the pointer itself is passed by value.

Example:

```go
func changeNum(ptr *int) {
    *ptr = 5
}
```

When calling:

```go
changeNum(&num)
```

Go copies the pointer value (the address), so both pointer values point to the same original variable.

Therefore the function can modify the original data through that address.

---

# 19. Nil Pointer

A pointer can be `nil`.

```go
var ptr *int

fmt.Println(ptr)
```

Output:

```text
<nil>
```

But this is dangerous:

```go
fmt.Println(*ptr)
```

because there is no valid address to dereference.

So check when necessary:

```go
if ptr != nil {
    fmt.Println(*ptr)
}
```

---

# 20. Common Pointer Mistake

Don't confuse:

```go
num := 10

ptr := &num
```

with:

```go
ptr := num
```

First:

```go
ptr := &num
```

`ptr` contains the address.

Second:

```go
ptr := num
```

`ptr` contains a copy of the value.

```text
ptr := &num
       ↓
    address


ptr := num
       ↓
     value
```

---

# 21. Your Code in One Diagram

Your code:

```go
num := 1

changeNum(&num)
```

creates this relationship:

```text
                main()
                  │
                  │
              num = 1
                  ▲
                  │
             address
                  │
                  │
          changeNum(num *int)
                  │
                  │
               *num
                  │
                  ▼
             same num
                  │
                  ▼
               5
```

So after:

```go
*num = 5
```

the original `num` becomes:

```text
num = 5
```

---

# ⭐ Pro Tips

1. `&variable` → gets the address of a variable.

2. `*pointer` → accesses the value at that address.

3. `*int` in a parameter means "pointer to an int".

4. Pointers are useful when a function needs to modify the original value.

5. Go is still **pass-by-value**. A pointer allows the copied pointer to refer to the same underlying variable.

6. A pointer that doesn't point anywhere is `nil`.

7. Don't dereference a nil pointer.

8. You can modify a variable through its pointer:

```go
*ptr = newValue
```

9. Don't memorize pointers only as symbols. Think in terms of:

```text
& → address
* → value at address
```

10. Pointers become especially important when working with:

    * Structs
    * Methods
    * Large data structures
    * Linked lists / trees
    * APIs
    * Concurrency
    * Performance-sensitive code

---

# 🧠 What I Learned Today

* A pointer stores the address of another variable.
* `&` gets the address of a variable.
* `*` dereferences a pointer.
* `*int` means pointer to an `int`.
* A function can modify the original variable using a pointer.
* `changeNum(&num)` passes the address of `num`.
* `*num = 5` modifies the original value.
* Go uses pass-by-value, including for pointer arguments.
* A pointer can be `nil`.
* Dereferencing a nil pointer causes a runtime panic.
* Pointers are useful for modifying existing data without copying the data itself.

---

# 📌 Quick Reference

```go
// Normal variable
num := 10

// Get address
ptr := &num

// Get value from pointer
fmt.Println(*ptr)

// Modify original value
*ptr = 20

fmt.Println(num) // 20
```

### Function with pointer

```go
func changeNum(num *int) {
    *num = 5
}

func main() {
    num := 1

    changeNum(&num)

    fmt.Println(num) // 5
}
```

### Remember

```text
&num
 ↓
address of num


ptr
 ↓
stores address


*ptr
 ↓
value at that address
```

> **Key idea:** `&` gives you the address, and `*` lets you access or modify the value stored at that address.

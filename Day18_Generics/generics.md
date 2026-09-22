# 🧬 Go Generics

## 1. What Are Generics?

**Generics** allow us to write functions, structs, and other code that can work with **multiple types**.

Without generics, we might need separate functions:

```go
func printInts(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStrings(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
```

The logic is exactly the same.

Generics let us write this once:

```go
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
```

Now the same function can work with:

```go
[]int
[]string
[]bool
```

and many other types.

---

# 2. Basic Generic Function

Your example:

```go
func printSlice[T any](items []T) {

	for _, item := range items {
		fmt.Println(item)
	}
}
```

The important part is:

```go
[T any]
```

Here:

```text
T = type parameter
any = constraint
```

Think of `T` as:

> "Some type that will be decided when this function is used."

---

# 3. Calling a Generic Function

You can do:

```go
num := []int{1, 2, 3, 4, 5}

printSlice(num)
```

Go can infer that:

```text
T = int
```

So conceptually:

```text
printSlice([]int)
       ↓
T becomes int
```

You usually don't need to explicitly write the type.

---

# 4. Multiple Types

The same function can work with strings:

```go
name := []string{
	"golang",
	"javascript",
	"python",
}

printSlice(name)
```

Now:

```text
T = string
```

For booleans:

```go
values := []bool{
	true,
	false,
	true,
}

printSlice(values)
```

Now:

```text
T = bool
```

Same function.

Different type.

---

# 5. What Does `any` Mean?

You wrote:

```go
[T any]
```

`any` is an alias for:

```go
interface{}
```

So these are equivalent:

```go
[T any]
```

and:

```go
[T interface{}]
```

For example:

```go
func printSlice[T any](items []T) {
	// ...
}
```

is equivalent to:

```go
func printSlice[T interface{}](items []T) {
	// ...
}
```

### Important

`any` does **not** mean:

> "This is no longer type-safe."

Generics still preserve the type relationship.

---

# 6. Why Use `any`?

`any` means the type parameter has **no specific restriction**.

For example:

```go
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
```

`T` can be:

```text
int
string
bool
float64
struct
pointer
etc.
```

as long as the value can be used as the element type of the slice.

---

# 7. Generic Type Constraint

You also tried:

```go
func printSlice[T int | string](items []T) {
	// ...
}
```

This is different from:

```go
[T any]
```

Here:

```go
int | string
```

means:

> T can be `int` OR `string`.

So:

```go
printSlice([]int{1, 2, 3})
```

works.

And:

```go
printSlice([]string{"Go", "Python"})
```

works.

But:

```go
printSlice([]bool{true, false})
```

does **not** satisfy the constraint.

---

# 8. `|` Means Type Set Choice

This:

```go
[T int | string]
```

means the allowed types are:

```text
int
string
```

You can have more:

```go
[T int | string | float64]
```

Now `T` can be:

```text
int
string
float64
```

---

# 9. `comparable` Constraint ⭐

Your final version uses:

```go
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
```

`comparable` is a built-in Go constraint.

It means:

> `T` must be a type that supports `==` and `!=`.

Examples of comparable types include:

```text
int
string
bool
float64
pointers
channels
arrays of comparable elements
structs whose fields are comparable
```

Slices, maps, and functions are not comparable for equality.

---

# 10. Why Is `comparable` Useful?

Your current function doesn't actually need `comparable`.

This:

```go
func printSlice[T comparable](items []T)
```

could simply be:

```go
func printSlice[T any](items []T)
```

because you're only printing the values.

You need `comparable` when your generic function needs to compare values.

For example:

```go
func contains[T comparable](items []T, target T) bool {

	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}
```

Now `comparable` is necessary because we use:

```go
item == target
```

---

# 11. Generic `contains()` Example

```go
func contains[T comparable](items []T, target T) bool {

	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}
```

Usage:

```go
nums := []int{1, 2, 3, 4}

fmt.Println(contains(nums, 3))
```

Output:

```text
true
```

String:

```go
names := []string{"Om", "Rahul", "Amit"}

fmt.Println(contains(names, "Om"))
```

Output:

```text
true
```

Same function.

Different types.

---

# 12. `any` vs `comparable`

This distinction is very important.

### `any`

```go
[T any]
```

Means:

```text
Almost any type is allowed.
```

### `comparable`

```go
[T comparable]
```

Means:

```text
Only comparable types are allowed.
```

### Specific type set

```go
[T int | string]
```

Means:

```text
Only int or string.
```

Think:

```text
any
 ↓
very broad

comparable
 ↓
types supporting == and !=

int | string
 ↓
specific allowed types
```

---

# 13. Generic Structs

Generics are not limited to functions.

You can also create **generic structs**.

Your example:

```go
type stack[T any] struct {
	elements []T
}
```

This means:

> `stack` can store elements of any type.

---

# 14. Creating a Generic Stack

You created:

```go
myStack := stack[string]{
	elements: []string{
		"hello",
		"world",
	},
}
```

Here:

```go
stack[string]
```

means:

```text
T = string
```

Therefore:

```go
elements []T
```

becomes conceptually:

```go
elements []string
```

---

# 15. Stack of Integers

The same generic struct can store integers:

```go
myStack := stack[int]{
	elements: []int{
		1,
		2,
		3,
	},
}
```

Now:

```text
T = int
```

So the stack contains:

```go
[]int
```

---

# 16. Stack of Booleans

```go
myStack := stack[bool]{
	elements: []bool{
		true,
		false,
	},
}
```

Now:

```text
T = bool
```

No need to create:

```text
intStack
stringStack
boolStack
```

You have one reusable type:

```go
stack[T]
```

---

# 17. Generic Struct + Methods

This is where generics become more powerful.

We can add a method:

```go
type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(value T) {
	s.elements = append(s.elements, value)
}
```

Notice:

```go
(s *stack[T])
```

The receiver also knows about `T`.

And:

```go
value T
```

means the value must have the same type as the stack.

---

# 18. Using `push()`

```go
myStack := stack[int]{}

myStack.push(10)
myStack.push(20)
myStack.push(30)

fmt.Println(myStack.elements)
```

Output:

```text
[10 20 30]
```

But:

```go
myStack.push("hello")
```

will not work because:

```text
stack[int]
```

expects:

```text
int
```

---

# 19. Generic Stack with Pop

We can create:

```go
func (s *stack[T]) pop() T {

	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]

	s.elements = s.elements[:lastIndex]

	return value
}
```

Usage:

```go
myStack := stack[int]{
	elements: []int{10, 20, 30},
}

value := myStack.pop()

fmt.Println(value)
```

Output:

```text
30
```

This demonstrates how generics can create reusable data structures.

---

# 20. Why Generics Are Useful

Without generics, you might write:

```text
IntStack
StringStack
FloatStack
BoolStack
```

With generics:

```go
stack[T]
```

One implementation can support different types.

This gives us:

```text
Reusable code
+
Type safety
+
Less duplication
```

---

# 21. Generics vs `any`

This distinction is very important.

You might think:

> "Why don't I just use `any`?"

Example:

```go
func printItems(items []any)
```

This can accept different values, but the relationship between the input type and the type parameter is different.

With generics:

```go
func printItems[T any](items []T)
```

Go knows that:

```text
T = the element type of this slice
```

The compiler can preserve that type information throughout the generic code.

Generics are useful when you want **reusable algorithms/data structures while retaining type relationships**.

---

# 22. Generics vs Interface

This is also important after your previous lesson.

### Interface

Interface describes **behavior**.

```go
type paymenter interface {
	pay(amount float32)
}
```

It answers:

> What can this type do?

### Generic

Generics describe **type-parameterized code**.

```go
func printSlice[T any](items []T)
```

It answers:

> How can I write this algorithm once for different types?

So:

```text
Interface → behavior
Generics  → reusable type-based code
```

They solve different problems.

---

# 23. Generic Function with Multiple Type Parameters

You can have more than one type parameter.

Example:

```go
func pair[A any, B any](first A, second B) {
	fmt.Println(first, second)
}
```

Usage:

```go
pair("age", 21)
pair("name", "Om")
```

Here:

```text
A → first type
B → second type
```

---

# 24. Type Inference

Go can often figure out the type parameter automatically.

Instead of:

```go
printSlice[int](num)
```

you can simply write:

```go
printSlice(num)
```

Go sees:

```go
num := []int{1, 2, 3}
```

and infers:

```text
T = int
```

This is called **type inference**.

---

# 25. Explicit Type Argument

You can also explicitly provide the type:

```go
printSlice[int](num)
```

For your example:

```go
myStack := stack[string]{
	elements: []string{"hello", "world"},
}
```

Here the type argument is explicitly specified:

```text
string
```

---

# 26. Type Parameters vs Type Arguments

These names can be confusing.

### Type parameter

Declared here:

```go
func printSlice[T any](items []T)
```

`T` is a **type parameter**.

### Type argument

Provided here:

```go
printSlice[int](num)
```

`int` is a **type argument**.

Think:

```text
Declaration:
T = placeholder

Usage:
int = actual type
```

---

# 27. Your Complete Example

```go
package main

import "fmt"

func printSlice[T comparable](items []T) {

	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(value T) {
	s.elemen
```

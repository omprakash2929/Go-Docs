package main

import "fmt"

//? Simple funcation
func add(a int, b int) int {
	return a + b
}

//? If all parameter same type then you are write like that
func sub(a, b int) int {
	return a - b
}

//* In go Funcation returen miltiple value

func getLanguages() (string, string, string) {
	return "golang", "javascript", "python"
}

//? pass function to funcation as argumenet

// func proccesIt(fn func(a int) int) {
// 	fn(1)
// }

//? return function
func proccesIt() func(a int) int {
	return func(a int) int {
		return 2
	}
}

//? it's take one funcation and also return funcation
func take(fn func(a int) int) func(a int) int {
	return func(a int) int {
		return fn(a)
	}
}
func main() {
	// sum := add(1, 2)

	// fmt.Println(getLanguages())

	len1, len2, _ := getLanguages() //* You get them spreated and if you not any value then use _ "underscroe"

	fmt.Println(len1, len2)
	//* It's anonymous funcation
	// fn := func(a int) int {
	// 	return 2
	// }

	fn := proccesIt()
	fn(6)
}

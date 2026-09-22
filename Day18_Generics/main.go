package main

import (
	"fmt"
)

//* Funcation
// func printSlice[T any](items []T) {   //? Also you write interface{} insted of "any"

//		for _, items := range items {
//			fmt.Println(items)
//		}
//	}
// func printSlice[T int | string](items []T) { //? Also you write interface{} insted of "any"

//		for _, items := range items {
//			fmt.Println(items)
//		}
//	}
func printSlice[T comparable](items []T) {

	for _, items := range items {
		fmt.Println(items)
	}
}

//* Struct

type stack[T any] struct { //? Also you write interface{} insted of "any"
	elements []T
}

func main() {
	//* Funcation
	// name := []string{"golang", "js", "python"}
	// bolll := []bool{true, false, true}
	num := []int{1, 2, 3, 4, 5}
	printSlice(num)

	//* Struct
	myStack := stack[string]{
		elements: []string{"hello", "world"},
	}
	fmt.Println(myStack)
}

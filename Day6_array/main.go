package main

import "fmt"

func main() {

	//! Normal way to create array
	// var num [4]int
	// num[0] = 1
	// num[1] = 2
	// fmt.Println(len(num), num)

	//! string array
	// var name [4]string
	// name[0] = "golang"
	// fmt.Println(name)

	//! declare in single line
	// nums := [3]int{1, 2, 3}

	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//! make 2d arrays
	// nums := [2][2]int{{1, 2}, {4, 5}}
	// fmt.Println(nums)

	//!automatically calculates the size

	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(nums)
	fmt.Println(len(nums))

	//? Fixed size that is predicatable
	//? Memory optimazation
	//? contant time access

}

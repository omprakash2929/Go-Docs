package main

import "fmt"

func main() {

	// age := 1

	// if age >= 12 && age <= 19 {
	// 	fmt.Println("it's a Teeager")
	// } else if age >= 20 {
	// 	fmt.Println("It's Adult")
	// } else {
	// 	fmt.Println("It's kids")
	// }

	//! && and ||

	// var role = "admin"
	// var hasPermission = false

	// if role == "admin" && hasPermission == true {
	// 	fmt.Println("You can access a admin panel")
	// } else if role == "admin" || hasPermission == false {
	// 	fmt.Println("You can access a cutomer panel")
	// }

	if age := 21; age >= 19 {
		fmt.Println("You are adult")
	} else if age <= 19 {
		fmt.Println("You are teenager")
	}

}

package main

import "fmt"

//? -> only construct in go for looping

func main() {

	//! While loop

	// i := 1

	// for i <= 9 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//! infinte while loop
	// for {
	// 	println("")
	// }

	// classic for loop

	// for i := 0; i < 5; i++ {
	// 	// break
	// 	if i == 3 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	//* Range

	for i := range 9 {
		fmt.Println(i)
	}

}

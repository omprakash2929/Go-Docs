package main

import "fmt"

//! itreating over data structures
func main() {

	//* For slices
	// nums := []int{6, 7, 8, 9}

	//? useing Foor loops
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//? useing Range method
	// sum := 0
	// for i, num := range nums { //! index, value
	// 	fmt.Println(i, num)
	// sum = sum + num }

	// fmt.Println(sum)

	//* For Maps

	m := map[string]string{"fname": "omprakash", "lname": "chauhan", "age": "21"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//* For string

	//? i -> return starting byte of rune it not index ,  c -> return Unicode code point rune
	for i, c := range "Omprakash" {
		fmt.Println(i, string(c))
	}
}

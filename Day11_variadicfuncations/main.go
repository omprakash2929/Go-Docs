package main

import "fmt"

//* Variadic Functions

func sum(nums ...int) int {

	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {

	// fmt.Println(1, 2, 3, 4, 5)
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := sum(1, 2, 3, 4, 5)
	fmt.Println(res)
	fmt.Println(sum(n...))
}

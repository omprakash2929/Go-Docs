package main

import "fmt"

func changeNum(num *int) {

	*num = 5

	fmt.Println("In Chanhe num:", *num)

}

func main() {

	num := 1

	changeNum(&num)
	// fmt.Println(&num)
	fmt.Println("in main funcation", num)

}

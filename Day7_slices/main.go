package main

import (
	"fmt"
	"slices"
)

//? slices is a dynaminc arrays
//? most used construct in go
//? + useful methods

func main() {
	//* uninitialized slice is nil mean "null"

	// var nums []int

	// fmt.Println(nums == nil) // true mean is nill
	// fmt.Println(len(nums))

	//var nums = make([]int, 3) // here we asined a 3 number of size and now it not nil
	// fmt.Println(nums == nil) // false

	// var nums = make([]int, 0, 5)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	//! second way top create slices
	// nums := []int{1, 2, 3, 4, 5}
	// nums = append(nums, 12)

	// nums[0] = 11

	// fmt.Println(nums)

	//! copy Funcation in Slices

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)
	// fmt.Println("Nums1", nums)
	// fmt.Println("Nums2", nums2)

	//! Slice Operator

	// var nums = []int{1, 2, 3, 4, 5}

	// fmt.Println(nums[0:3])
	// fmt.Println(nums[:4])
	// fmt.Println(nums[3:])

	//? append another slice

	// num1 := []int{1, 2, 3}
	// num2 := []int{4, 5, 6}

	// num3 := append(num1, num2...)

	// fmt.Println(num3)

	//! Slices Package

	var num1 = []int{1, 2}
	var num2 = []int{1, 5, 3, 6, 2, 4}

	// fmt.Println(slices.Max(num1))
	// fmt.Println(slices.Max(num2))
	// fmt.Println(slices.Min(num2))
	fmt.Println(slices.Equal(num1, num2))

	// fmt.Println(slices.Contains(num1, 2))
	// fmt.Println(slices.Index(num2, 5))
	// slices.Sort(num2)
	// fmt.Println(num2)

	//? custome sort

	// type Person struct {
	// 	Name string
	// 	age  int
	// }
	// people := []Person{{"om", 24}, {"Raj", 21}}
	// slices.SortFunc(people, func(a, b Person) int {
	// 	return a.age - b.age
	// })

	// slices.Reverse(num2) //? In place
	// fmt.Println(num2)

	// num2 = slices.Insert(num2, 1, 100) //? index:1 -> value:100

	// fmt.Println(num2)

	// num2 = slices.Delete(num2, 1, 3)  //? delete to index 1 to 3
	// fmt.Println(num2)

	// s := slices.Clone(num2) //? deep copy
	// fmt.Println(s)

	// s := slices.Compact(num2) //? It rmeove the duplicate data
	// fmt.Println(s)
}

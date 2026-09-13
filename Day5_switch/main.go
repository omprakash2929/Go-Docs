package main

import (
	"fmt"
)

func main() {

	// Simple switch

	// num := 3

	// switch num {
	// case 1:
	// 	fmt.Println("One") //! you do not write a brack. it can automatically handel it.
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Not supported")
	// }

	//? Multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekends")
	// default:
	// 	fmt.Println("it's Work Day")
	// }

	//? Type Switch

	whoAmi := func(i interface{}) {
		switch k := i.(type) {
		case int:
			fmt.Println("it's integer", k)
		case string:
			fmt.Println("It's String", k)
		default:
			fmt.Println("Not supported", k)
		}
	}
	whoAmi(true)
}

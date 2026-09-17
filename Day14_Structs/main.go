package main

import (
	"fmt"
	"time"
)

// Order Struct

// * if you don't set any field, defult value is zero value
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //* nanosecond precision
}

// ! In go constructor
func newOrder(id string, amount float32, status string) *order {
	//! Initial setup goes here...
	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myorder
}

// ? receiver type
func (o *order) changeStatus(status string) {

	o.status = status //* If you want to change and modify a struct value then you use pointer
}

// ? If you only get value then you not use pointer is well
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// var myorder order = order{
	// 	id:     "001",
	// 	amount: 200,
	// 	status: "Done",
	// }

	// myorder := order{
	// 	id:     "001",
	// 	amount: 200,
	// 	status: "Done",
	// }

	// myorder.changeStatus("Confirmed")

	// myorder.createdAt = time.Now()
	// fmt.Println(myorder)
	// fmt.Println(myorder.amount)

	myOrder := newOrder("001", 300, "shipped")

	fmt.Println(myOrder.amount)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}

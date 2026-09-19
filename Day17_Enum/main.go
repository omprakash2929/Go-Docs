package main

import "fmt"

//! enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "recived"
	Confirmed             = "confirmed"
	Prepared              = "perpared"
	Delivered             = "delivered"
)

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

func changeorderStatus(status OrderStatus) {
	fmt.Println("Changing Order status to", status)
}

func main() {

	changeorderStatus(Delivered)
}

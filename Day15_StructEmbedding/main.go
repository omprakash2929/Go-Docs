package main

import (
	"fmt"
	"time"
)

//! Struct Embedding

type customer struct {
	name  string
	phone string
}

//* composition

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //* nanosecond precision
	customer
}

func main() {

	//? method 1
	// newCustome := customer{
	// 	name:  "ompraksh",
	// 	phone: "9510276296",
	// }

	newOrder := order{
		id:     "001",
		amount: 300,
		status: "Recevied",
		// customer: newCustome,
		customer: customer{ //* Inline
			name:  "prakash",
			phone: "123456789",
		},
	}

	newOrder.customer.name = "Rhaul"
	newOrder.phone = "992311272"
	fmt.Println(newOrder)
}

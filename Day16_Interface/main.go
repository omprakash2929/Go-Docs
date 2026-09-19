package main

import "fmt"

//! Interface

type paymenter interface {
	//* Here we are create Method not funcation
	pay(amount float32)
}

type payment struct {
	gatway paymenter
}

// open close principle
func (p payment) makePayment(amount float32) {

	// razorpayPayementgw := razorpay{}
	// razorpayPayementgw.pay(amount)

	// strippayementgw := strip{}
	p.gatway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//? logic to make payment

	fmt.Println("Making paymenet using Razorpay", amount)
}

type strip struct{}

func (s strip) pay(amount float32) {
	fmt.Println("Making payment using strip", amount)
}

//? If we want testing then it casuse issue that reasone we use interface

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making payment using Fake for testing:", amount)
}

func main() {
	// strippayementgw := strip{}
	fakeGw := fakepayment{}

	newPayment := payment{
		gatway: fakeGw,
	}

	newPayment.makePayment(100)
}

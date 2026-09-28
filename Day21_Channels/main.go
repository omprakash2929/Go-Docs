package main

import (
	"fmt"
	"time"
)

// * Sending
func proccessNum(numChan chan int) {
	fmt.Println("Proccessing number", <-numChan)
}

// * recived
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2

	result <- numResult
}

// * gorouting synchronizer && Using channel block insted using waitgroup
func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("Processing...")
}

//* buffer channel golang

func emailSender(emailChan chan string, done chan bool) { //?(emailChan <- chan string, done chan <- bool)
	defer func() {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("sending Email to", email)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
		chan2 <- "Golang"
	}()
	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1val := <-chan1:
			fmt.Println("Recevied data to chan 1", chan1val)
		case chan2val := <-chan2:
			fmt.Println("Recevied data to chan 2", chan2val)

		}
	}

	//!-----------------------------------------
	//! buffer channel golang
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// // emailChan <- "1@example.com"
	// // emailChan <- "om@gmail.com"

	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan)
	// fmt.Println("Done Sending...")
	// close(emailChan) //? !Imporant to close channel after the buffer use
	// <-done
	//! ----------------------
	// done := make(chan bool)

	// go task(done)

	// <-done //? block

	//! ----------------------

	// sumnum := make(chan int)

	// go sum(sumnum, 4, 5)

	// res := <- sumnum  //? Blocking
	// fmt.Println(res)

	//! ----------------------
	// numChan := make(chan int)

	// go proccessNum(numChan)

	// numChan <- 5

	// time.Sleep(time.Second * 2)

	// messageChan := make(chan string)

	// messageChan <- "Ping" //? Channels are blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}

package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {

	for i := 0; i <= 10; i++ {
		// go task(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}

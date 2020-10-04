package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	// receive
	for v := 0; v < 100; v++ {
		fmt.Println(<-c)
	}

	fmt.Println("about to exit")
}

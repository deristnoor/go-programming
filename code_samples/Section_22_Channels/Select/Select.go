package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odds := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odds, quit)

	// receive
	receive(even, odds, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odds channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
}

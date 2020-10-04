package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// doesn't work:

//func main() {
//	cs := make(chan<- int)
//
//	go func() {
//		cs <- 42
//	}()
//	fmt.Println(<-cs)
//
//	fmt.Printf("------\n")
//	fmt.Printf("cs\t%T\n", cs)
//}
// error: invalid operation: <-cs (receive from send-only type chan<- int)

//func main() {
//	cr := make(<-chan int)
//
//	go func() {
//		cr <- 42
//	}()
//	fmt.Println(<-cr)
//
//	fmt.Printf("------\n")
//	fmt.Printf("cr\t%T\n", cr)
//}
// error: invalid operation: cr <- 42 (send to receive-only type <-chan int)

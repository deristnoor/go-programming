package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	// assign function to a variable
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("the year big brother watching:", x)
	}
	f2(1984)
}

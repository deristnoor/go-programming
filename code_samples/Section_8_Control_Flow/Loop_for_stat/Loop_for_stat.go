package main

import (
	"fmt"
)

func main() {
	// for statement with single condition
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}

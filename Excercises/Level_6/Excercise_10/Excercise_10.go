package main

import (
	"fmt"
)

func main() {
	a := any()
	b := any()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func any() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

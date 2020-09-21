package main

import (
	"fmt"
)

func main() {
	x := 0
	foo(x)
	fmt.Println(x)
}

// step 1: no pointers
func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

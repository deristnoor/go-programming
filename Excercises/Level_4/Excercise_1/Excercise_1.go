package main

import (
	"fmt"
)

func main() {
	x := [5]int{4, 2, 6, 3, 7}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

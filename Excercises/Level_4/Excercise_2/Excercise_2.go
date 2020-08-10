package main

import (
	"fmt"
)

func main() {
	x := []int{4, 2, 6, 3, 7, 8, 3, 11, 1, 25}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

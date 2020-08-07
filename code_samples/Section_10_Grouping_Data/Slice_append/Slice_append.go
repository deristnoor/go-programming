package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 101)
	fmt.Println(x)

	y := []int{234, 245, 3567, 345}
	x = append(x, y...)
	fmt.Println(x)

	// y... >> all value of the same type

}

package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...)
	x2 := bar(xi)
	fmt.Println(x)
	fmt.Println(x2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, i := range xi {
		sum += i
	}
	return sum
}

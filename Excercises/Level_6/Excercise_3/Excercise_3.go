package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	defer foo(xi...)
	bar(xi)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("this is foo", sum)
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, i := range xi {
		sum += i
	}
	fmt.Println("this is bar", sum)
	return sum
}

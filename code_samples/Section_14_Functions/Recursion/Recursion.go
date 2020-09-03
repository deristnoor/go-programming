package main

import (
	"fmt"
)

func main() {
	fmt.Println(4 * 3 * 2 * 1)

	n := factorial(4)
	fmt.Println(n)

	n2 := loopFact(4)
	fmt.Println(n2)
}

// recursion start here
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 4 * 3 * 2 * 1 * 1
// end here

// another way for recursion using loop
func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

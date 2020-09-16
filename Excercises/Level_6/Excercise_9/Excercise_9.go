package main

import (
	"fmt"
)

func main() {
	y := func(x []int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}

	z := any(y, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(z)
}

func any(f func(x []int) int, ii []int) int {
	a := f(ii)
	return a

}

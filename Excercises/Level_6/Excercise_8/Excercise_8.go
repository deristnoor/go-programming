package main

import (
	"fmt"
)

func main() {
	x := try()

	y := x()

	fmt.Println(y)
}

func try() func() int {
	return func() int {
		return 12
	}
}

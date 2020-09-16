package main

import (
	"fmt"
)

func main() {

	list := func(x int) {
		for i := 0; i <= x; i++ {
			fmt.Println(i)
		}
	}
	list(50)

	fmt.Println("done")
}

package main

import (
	"fmt"
)

func main() {
	// nesting loop >> loop within a loop
	for i := 0; i <= 100; i++ {
		fmt.Printf("the outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}
}

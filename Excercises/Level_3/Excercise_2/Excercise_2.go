package main

import (
	"fmt"
)

func main() {
	for j := 65; j <= 90; j++ {
		fmt.Println(j)
		for k := 0; k < 3; k++ {
			fmt.Printf("\t%#U\n", j)
		}
	}
}

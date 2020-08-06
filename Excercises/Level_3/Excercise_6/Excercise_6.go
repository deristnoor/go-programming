package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		if i%6 == 0 {
			fmt.Println(i)
		}
	}
}

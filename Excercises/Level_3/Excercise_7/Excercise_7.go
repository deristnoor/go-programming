package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		} else if i%10 == 2 {
			fmt.Println("Two")
		} else {
			fmt.Println("nothing")
		}
	}
}

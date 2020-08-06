package main

import (
	"fmt"
)

// my solution
func main() {
	for i := 0; i <= 2020; i++ {
		if i >= 1993 {
			fmt.Println(i)
		}
	}
}

// todd solution
func main2() {
	bd := 1993
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}

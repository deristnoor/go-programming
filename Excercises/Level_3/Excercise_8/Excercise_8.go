package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("not print")
	case true:
		fmt.Println("print")
	}
}

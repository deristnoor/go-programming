package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	z := [][]string{x, y}
	fmt.Println(z)

	for i, a := range z {
		fmt.Println("record: ", i)
		for j, val := range a {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}

// Using iota, create 4 constants for the NEXT 4 years.
// Print the constant values.

package main

import (
	"fmt"
)

const (
	a = 2020 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

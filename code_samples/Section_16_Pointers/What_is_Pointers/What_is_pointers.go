package main

import (
	"fmt"
)

func main() {
	a := 12
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an adress when you have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)  // decimal
	fmt.Printf("%b\n", n)  // binary
	fmt.Printf("%#X\n", n) // hexadecimal
}

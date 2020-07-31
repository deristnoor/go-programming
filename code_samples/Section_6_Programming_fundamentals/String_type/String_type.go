package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// utf-8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	// hexadecimal
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x", i, v)
	}
}

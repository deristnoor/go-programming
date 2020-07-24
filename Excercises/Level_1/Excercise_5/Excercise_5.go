package main

import "fmt"

type haha int

var x haha
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 76
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

// change name

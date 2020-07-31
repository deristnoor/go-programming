package main

import "fmt"

// int & uint have u/int8, u/int16, u/int32, u/int64
// uint from 0, int from negative number to positive
// int16 signed 16-bit integers (-32768 to 32767)
// uint16 signed 16-bit integers (0 to 65535)
// uintptr unsigned integer

var x int8
var y float64

func main() {
	x = 42
	y = 42.018645
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

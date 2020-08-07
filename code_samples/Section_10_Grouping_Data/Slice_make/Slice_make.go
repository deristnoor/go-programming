package main

import (
	"fmt"
)

func main() {
	// if you know how many capacity of your array should make >> use make
	// make([]type, length, capacity)

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 234)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 345)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// if the capacity is exceeded, go will double the capacity to 24
	x = append(x, 1973)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

package main

import (
	"fmt"
)

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "Lady",
	}
	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(f *person) {
	f.first = "Slope"

	// this also works
	// (*f).first = "Klose"
}

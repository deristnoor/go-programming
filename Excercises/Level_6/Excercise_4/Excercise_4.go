package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "age", p.age)
}

func main() {
	a := person{
		first: "Frank",
		last:  "Martin",
		age:   32,
	}

	a.speak()
}

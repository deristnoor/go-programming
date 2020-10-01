package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"James",
		"Bond",
	}

	fmt.Println(p1)

	// this work
	saySomething(&p1)

	// this doesn't work
	// saySomething(p1)

}

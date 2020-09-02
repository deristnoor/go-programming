package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// INTERFACE allow a VALUE to be of more than one TYPE
// if you have empty interface, every value can be put in there
type human interface {
	speak()
}

type hotdog int

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	// polymorphism: the function can take many different types
	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

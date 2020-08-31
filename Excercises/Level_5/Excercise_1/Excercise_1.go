package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIceCream: []string{
			"Hazelnut",
			"Chocolate",
			"Vanilla",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favIceCream: []string{
			"Strawberry",
			"Green tea",
			"Avocado",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}

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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("---------")
	}
}

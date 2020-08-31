package main

import (
	"fmt"
)

func main() {
	t := struct {
		vehicle []string
		doors   []int
		color   []string
	}{
		vehicle: []string{
			"coupe",
			"sedan",
			"hatchback",
		},
		doors: []int{2, 4, 5},
		color: []string{
			"white",
			"red",
			"black",
			"blue"},
	}

	m := map[string]int{
		t.vehicle[0]: t.doors[0],
		t.vehicle[1]: t.doors[1],
		t.vehicle[2]: t.doors[2],
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v, "doors")
		for i, val := range t.color {
			fmt.Println(i, val)
		}
		fmt.Println("---------")
	}
}

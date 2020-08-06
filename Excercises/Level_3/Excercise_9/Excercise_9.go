package main

import (
	"fmt"
)

func main() {
	favSport := "Soccer"
	switch favSport {
	case "Baseball":
		fmt.Println("Who is your favourite Baseball palyer?")
	case "Soccer":
		fmt.Println("Who is your favourite Soccer palyer?")
	case "Basketball":
		fmt.Println("Who is your favourite Basketball palyer?")
	}
}

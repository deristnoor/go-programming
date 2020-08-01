// Using the following operators, write expressions and assign their values to variables:
// ==
// <=
// >=
// !=
// <
// >

package main

import (
	"fmt"
)

func main() {
	a := (42 == 4)
	b := (42 <= 4)
	c := (42 >= 4)
	d := (42 != 4)
	e := (42 < 4)
	f := (42 > 4)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

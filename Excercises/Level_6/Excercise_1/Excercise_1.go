package main

import (
	"fmt"
)

func main() {
	foo()
	bar()
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 21
}

func bar() (int, string) {
	return 23, "hello"
}

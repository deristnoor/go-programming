package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width

}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

type shape interface {
	area() float64
}

func info(h shape) {
	fmt.Println(h.area())
}

func main() {
	x := square{
		length: 45,
		width:  12,
	}

	y := circle{
		radius: 32,
	}
	info(x)
	info(y)
}

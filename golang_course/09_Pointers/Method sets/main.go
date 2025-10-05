package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s *square) area() float64 {
	return s.length * s.length
}
func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	//NON-POINTER RECEIVER & NON-POINTER VALUE
	c := circle{5}
	info(c)
	// NON-POINTER RECEIVER & NON-POINTER VALUE
	c1 := circle{radius: 10}
	info(&c1)
	// POINTER RECEIVER & POINTER VALUE
	s := square{
		length: 3.5,
	}
	info(&s)
}

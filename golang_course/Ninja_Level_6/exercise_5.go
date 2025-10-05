package main

import "C"
import (
	"fmt"
	"math"
)

/*
create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/

type SQUARE struct {
	Length float64
}

type CIRCLE struct {
	radius float64
}
type SHAPE interface {
	area() float64
}

func (s SQUARE) area() float64 {
	return s.Length * s.Length
}
func (c CIRCLE) area() float64 {
	return math.Pi * c.radius * c.radius
}

func INFO(i SHAPE) {

	fmt.Println(i.area())
}

func main() {

	c1 := CIRCLE{
		radius: 12.056,
	}
	s1 := SQUARE{
		Length: 25.56,
	}
	INFO(c1)
	INFO(s1)
}

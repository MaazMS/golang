## Method sets  
1. when we have a type, we can attach methods to it.   
1. Those methods attach to a type are known as its method set.   
1. IMPORTANT: “The method set of a type determines the INTERFACES that the type implements.....   

|Receivers |Values|  
|---|---|  
|(t T) | T and *T|  
|(t *T) |*T  

## NON-POINTER RECEIVER & NON-POINTER VALUE  
```go
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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	//NON-POINTER RECEIVER & NON-POINTER VALUE
	c := circle{5}
	info(c)
}
/* output
area 78.53981633974483
*/
``` 

## NON-POINTER RECEIVER & POINTER VALUE  
```go
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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	
	// NON-POINTER RECEIVER & NON-POINTER VALUE
	c1 := circle{radius: 10}
	info(&c1)
}
/* output
area 314.1592653589793
*/
``` 
## POINTER RECEIVER & POINTER VALUE
```go
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
	// POINTER RECEIVER & POINTER VALUE
	s := square{
		length: 3.5,
	}
	info(&s)
}
/*
area 12.25
*/
```
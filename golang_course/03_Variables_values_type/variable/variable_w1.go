package main

import (
	"fmt"
	"strconv"
)

var no1 int = 1
var (
	firstname string = "Maaz"
	lastname  string = "Shaikh"
	age       int    = 23
)

func main() {

	fmt.Printf("%v, %T  \n", no1, no1)
	var no1 int = 11
	fmt.Printf("%v, %T  \n", no1, no1)
	var no1integer int
	no1integer = 10

	var no2integer int = 20

	no3integer := 30

	var no1float float64
	no1float = 11.0

	var no2float float64 = 12.

	no3float := 13.

	fmt.Printf("%v, %T  \n", no1integer, no1integer)
	fmt.Printf("%v, %T  \n", no2integer, no2integer)
	fmt.Printf("%v, %T  \n", no3integer, no3integer)

	fmt.Printf("%v, %T  \n", no1float, no1float)
	fmt.Printf("%v, %T  \n", no2float, no2float)
	fmt.Printf("%v, %T  \n", no3float, no3float)

	fmt.Printf("%v, %T  \n", no1, no1)
	fmt.Printf("%v, %T  \n", firstname, firstname)
	fmt.Printf("%v, %T  \n", lastname, lastname)
	fmt.Printf("%v, %T  \n", age, age)

	// int to float conversion
	fmt.Println("int to float conversion")
	var i int = 16
	fmt.Printf("%v, %T  \n", i, i)
	var j float64
	j = float64(i)
	fmt.Printf("%v, %T  \n", j, j)

	// float to int conversion
	fmt.Println("float to int conversion")
	var a float32 = 16.5
	fmt.Printf("%v, %T  \n", a, a)
	var b int
	b = int(a)
	fmt.Printf("%v, %T  \n", b, b)

	// int to String conversion
	fmt.Println("int to String conversion")
	var m int = 16
	fmt.Printf("%v, %T  \n", m, m)
	var n string
	n = strconv.Itoa(m)
	fmt.Printf("%v, %T  \n", n, n)

}

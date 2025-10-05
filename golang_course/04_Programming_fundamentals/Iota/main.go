package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)
const (
	l = iota
	m
	n
)

func main()  {

	fmt.Println(a)  // a = 0
	fmt.Println(b) // b = 1
	fmt.Println(c) // c = 2
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)

	fmt.Println(l)  // l = 0
	fmt.Println(m) // m = 1
	fmt.Println(n) // n = 2
	fmt.Printf("%T\n",l)
	fmt.Printf("%T\n",m)
	fmt.Printf("%T\n",n)



}
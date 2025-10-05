package main

import "fmt"

/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values

*/

const (
	a = 2012 + iota
	b = 2012 + iota
	c = 2012 + iota
	d = 2012 + iota

)

func main()  {
	fmt.Println(a, b, c, d)
}
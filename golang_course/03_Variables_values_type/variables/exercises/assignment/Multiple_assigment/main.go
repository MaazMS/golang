package main

import "fmt"

func main()  {

	var (
		name string
		phono int
		force float64
	)


	name, phono, force= "Maaz", 123, 1
	fmt.Println(name, phono, force)

	name, phono, force = "Shaikh",  987, 100
	fmt.Println(name, phono, force)
}

package main

import "fmt"

func main()  {

	var firstname = "Maaz"
	var lastname = "Shaikh"

	fmt.Println(firstname,lastname)

	firstname, lastname = lastname,  firstname
	fmt.Println(firstname,lastname)
}

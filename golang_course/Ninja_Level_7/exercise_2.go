package main

import "fmt"

/*
create a person struct
● create a func called “changeMe” with a *person as a parameter
	○ change the value stored at the *person address
create a value of type person
	○ print out the value
● in func main
	○ call “changeMe”
● in func main
	○ print out the value
*/

type person struct {
	Name string
}

func main() {
	p1 := person{
		Name: "Maaz",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(a *person) {
	a.Name = "Shaikh"
}

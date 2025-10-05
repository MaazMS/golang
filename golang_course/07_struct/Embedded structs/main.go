package main

import "fmt"

type person struct {
	name string
	age  int
}

type work struct {
	person
	filed string
}

func main() {

	w1 := work{
		person: person{
			name: "Maaz",
			age:  23,
		},
		filed: "IT",
	}
	fmt.Println(w1.filed, w1.name, w1.age, w1.person)
}

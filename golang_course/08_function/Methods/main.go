package main

import "fmt"

type person struct {
	Name string
	age  int
}

func main() {
	p := person{
		Name: "Maaz",
		age:  24,
	}
	fmt.Println(p)
	p.speak()
}

func (s person) speak() {
	fmt.Println("my self is", s)
}

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p1 := person{
		name: "Maaz",
		age:  20,
	}
	fmt.Println(p1)

	p2 := person{
		name: "Shaikh",
		age:  25,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.name, p1.age)
	fmt.Println(p2.name, p2.age)

}

package main

import "fmt"

type person struct {
	Name string
	age  int
}

type work struct {
	person
	filed string
}

func (w work) speak() {
	fmt.Println("This function of type work", w.Name, w.age, w.filed)
}

func (p person) speak() {
	fmt.Println("This function of type person", p.Name, p.age)
}

// type interface

type human interface {
	speak()
}

func baar(h human) {

	switch h.(type) {
	case person:
		fmt.Println("This function baar of type person", h.(person).Name)
	case work:
		fmt.Println("This function baar of type work ", h.(work).Name)

	}
}
func main() {

	w1 := work{
		person: person{
			Name: "Maaz",
			age:  24,
		},
		filed: "IT",
	}
	fmt.Println(w1)
	w1.speak()
	baar(w1)

	fmt.Println("\n\n")
	p1 := person{
		Name: "shaikh",
		age:  24,
	}

	fmt.Println(p1)
	p1.speak()
	baar(p1)

}

## interface 
1. An interface allows a value to be of more than  one type.     
1. We create an interface using this syntax: “keyword identifier type” so for an interface.  
   `type identifier interface`  
1. When define method inside interface a type must have to implement that interface.      
1. An interface type specifies a `method set` called its interface.      
**Note** Value have one or more type.   


1. interfaces allow things to work well together to interface together.
1.  empty interface, which means that every other type has no methods.
1.  I could pass in any other type.  

## polymorphism  
1. polymorphism means same method name (but different signatures) being uses for different types.  

```go
package main

import "fmt"

type person struct {
	Name string
	age int
}

type work struct {
	person
	filed string
}

func (w work) speak()  {
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
func main()  {

	w1 := work{
		person: person{
			Name: "Maaz",
			age: 24,
		},
		filed: "IT",
	}
	fmt.Println(w1)
	w1.speak()
	baar(w1)

	fmt.Println("\n\n")
	p1 := person{
		Name: "shaikh",
		age: 24,
	}

	fmt.Println(p1)
	p1.speak()
	baar(p1)

}

```  

````bash
{{Maaz 24} IT}
This function of type work Maaz 24 IT
This function baar of type work  Maaz



{shaikh 24}
This function of type person shaikh 24
This function baar of type person shaikh
````
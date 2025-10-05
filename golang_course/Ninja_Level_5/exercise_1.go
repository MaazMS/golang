package main

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
	● first name
	● last name
	● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors
*/

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	ice_cream  []string
}

func main() {

	p1 := person{
		first_name: "Maaz",
		last_name:  "Shaikh",
		ice_cream: []string{
			"chocolate",
			"strawberry",
		},
	}
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	for i, k := range p1.ice_cream {
		fmt.Println(i, k)
	}
	p2 := person{
		first_name: "Md",
		last_name:  "Shaikh",
		ice_cream: []string{
			"chocolate",
			"strawberry",
		},
	}
	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	for i, k := range p2.ice_cream {
		fmt.Println(i, k)
	}

}

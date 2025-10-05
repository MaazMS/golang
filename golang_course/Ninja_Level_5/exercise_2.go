package main

/*
Take the code from the previous exercise, then store the values of type person in a map with the
key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

import (
	"fmt"
)

type persons struct {
	first_name string
	last_name  string
	ice_cream  []string
}

func main() {

	p1 := persons{
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
	m := map[string]string{
		p1.last_name: "khan",
		p1.last_name: "sayed",
	}
	fmt.Println("\n\n")
	fmt.Println(m)

	for i, k := range m {
		fmt.Println(i, k)
	}

}

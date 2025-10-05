package main

import "fmt"

// Create and use an anonymous struct.

func main() {

	a := struct {
		Name    string
		age     int
		friends map[string]int
		city    []string
	}{
		Name: "Maaz",
		age:  10,
		friends: map[string]int{
			"One": 1,
			"Two": 2,
		},
		city: []string{
			"Pune",
			"Mumbai",
		},
	}

	fmt.Println(a.Name, a.age)
	fmt.Println("\n\n")
	for i, k := range a.friends {
		fmt.Println(i, k)
	}
	fmt.Println("\n\n")

	for i, k := range a.city {
		fmt.Println(i, k)
	}

}

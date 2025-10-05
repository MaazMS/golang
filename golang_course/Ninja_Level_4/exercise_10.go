package main

import "fmt"

/* previous example
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

Using the code from the previous example, delete a record from your map. Now print the map
out using the “range” loop

*/

func main() {

	person := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for i, k := range person {
		fmt.Printf("The record of map %s\t\n", i)
		for j, vls := range k {
			fmt.Println(j, "\t", vls)
		}
	}

	fmt.Println("After delete\n")
	delete(person, `no_dr`)
	for i, k := range person {
		fmt.Printf("The record of map %s\t\n", i)
		for j, vls := range k {
			fmt.Println(j, "\t", vls)
		}
	}
}

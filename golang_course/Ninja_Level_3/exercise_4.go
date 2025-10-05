package main

/*
Create a for loop using this syntax
● for { }
Have it print out the years you have been alive

*/

import (
	"fmt"
)

func main() {
	fmt.Printf("date of birth year", 1997)
	year := 1997
	for {
		year++
		if year > 2021 {
			break
		}
		fmt.Println(year)

	}
}

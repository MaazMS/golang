package main

/*
Create a for loop using this syntax
● for condition { }
Have it print out the years you have been alive
*/

import (
	"fmt"
)

func main() {
	fmt.Printf("date of birth year %d\t", 1997)

	birthday := 1997
	for birthday <= 2021 {
		fmt.Println(birthday)
		birthday++
	}
}

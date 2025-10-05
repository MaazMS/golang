package main

import "fmt"

func main() {

	var value string = "five"

	switch value {

	case "one":
		fmt.Println("One ")
	case "two":
		fmt.Println("two")
	case "three":
		fmt.Println("Three")
	case "four", "five":
		fmt.Println("five")
	}
}

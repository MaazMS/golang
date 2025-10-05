package main

import "fmt"

/*
Create a program that uses a switch statement with the switch expression specified as a
variable of TYPE string with the IDENTIFIER “favSport”.
*/
func main() {

	favSport := "online game"
	switch favSport {

	case "offline ":
		fmt.Println("off line game")
	case "online game":
		fmt.Println("online game")

	}
}

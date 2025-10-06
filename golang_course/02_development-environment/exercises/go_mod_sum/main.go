package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Print a simple message
	fmt.Println("Welcome to Go Programming!")

	// Example: Print your name
	name := "Maaz"
	fmt.Println("Hello,", name)

	// Example: Add two numbers
	a := 10
	b := 20
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// Example: Generate a UUID
	uuid := uuid.New()
	fmt.Println("Generated UUID:", uuid)
}

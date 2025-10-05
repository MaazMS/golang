package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Editor", os.Getenv("EDITOR"))

	err := os.Setenv("EDITOR", "Jetbrians-goland")

	if err != nil {
		fmt.Println("os.Setenv is not set env variable")
		return
	}
	fmt.Println("Editor", os.Getenv("EDITOR"))
}

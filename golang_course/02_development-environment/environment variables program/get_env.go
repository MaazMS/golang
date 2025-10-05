package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("shell : location ", os.Getenv("SHELL"))
}

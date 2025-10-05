package main

import (
	"fmt"
	"os"
)

func main() {

	getEnv := func(key string) {
		val, ok := os.LookupEnv(key)

		// if value is not set ok variable value is false if value is return it is true.
		fmt.Println(ok)
		// if value is false val is empty and if value is not empty it have value
		fmt.Println(val)
		if !ok {
			fmt.Printf("%s is not set \n", key)
		} else {
			fmt.Printf("%s=%s \n", key, val)
		}
	}

	getEnv("EDITOR")
	getEnv("SHELL")
}

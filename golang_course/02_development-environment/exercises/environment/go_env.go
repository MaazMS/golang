package main

import (
	"fmt"
	"os"
)

func main() {
	// 1️⃣ Set an environment variable
	err := os.Setenv("MY_VAR", "Hello_Golang")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}
	fmt.Println("Environment variable set: MY_VAR=Hello_Golang")

	// 2️⃣ Get (read) the environment variable
	value := os.Getenv("MY_VAR")
	if value == "" {
		fmt.Println("MY_VAR not found!")
	} else {
		fmt.Println("Value of MY_VAR:", value)
	}

	// 3️⃣ Remove (unset) the environment variable
	err = os.Unsetenv("MY_VAR")
	if err != nil {
		fmt.Println("Error removing environment variable:", err)
		return
	}
	fmt.Println("Environment variable MY_VAR removed")

	// 4️⃣ Try to get it again (should be empty)
	value = os.Getenv("MY_VAR")
	if value == "" {
		fmt.Println("After removal, MY_VAR not found.")
	} else {
		fmt.Println("MY_VAR still exists with value:", value)
	}
}

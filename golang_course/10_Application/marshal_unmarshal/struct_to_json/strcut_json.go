package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json: "author"`
}

func main() {

	fmt.Println("Hello")

	book := Book{Title: "Golang",
		Author: "Me",
	}
	byteArray, err := json.Marshal(book)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}

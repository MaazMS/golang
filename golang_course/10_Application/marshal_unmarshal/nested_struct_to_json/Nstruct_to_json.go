package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string     `json:"title"`
	Author AuthorInfo `json:authorinfo`
}

type AuthorInfo struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {

	authordetails := AuthorInfo{Name: "Abcd",
		Age:       25,
		Developer: true,
	}
	book := Book{Title: "Golang",
		Author: authordetails,
	}

	//byteArray, err := json.Marshal(book)
	// formatmat json string using josn.Marshalindent
	byteArray, err := json.MarshalIndent(book, "", "  ")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}

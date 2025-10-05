package main

import (
	"fmt"
)

type SliceStruct []struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

func main() {

	s := SliceStruct{
		{
			Data: "user",
			Name: "Maaz",
		},
		{
			Data: "website",
			Name: "Amazon",
		},
	}
	for i, v := range s {
		fmt.Println(i, v.Data, v.Name)
	}
	fmt.Println(s)
}

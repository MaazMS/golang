package main

import (
	"encoding/json"
	"fmt"
)

type dimension struct {
	Height int
	Width int
}

//Weight key itself will be omitted from the JSON object. If key and value is not provided.
type Shape struct {
	Name    string
	Weight int `json:"weight,omitempty"`

	// Now `Size` is a pointer to a `dimension` instance
	Size *dimension `json:"size,omitempty"`
}
// Size attribute, and set its omitempty tag,  it still appears in the output. This is because structs do not have any empty value

func main() {
	d := Shape{
		Name:    "Triangle",

	}
	b,_ := json.Marshal(d)
	fmt.Println(string(b))
}

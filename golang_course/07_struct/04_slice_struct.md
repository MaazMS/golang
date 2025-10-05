## slice struct  
1. 

```go
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

	fmt.Println(s)
}

```
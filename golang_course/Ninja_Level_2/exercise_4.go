package main
/*
Create a variable of type string using a raw string literal. Print it.
*/

import (
	"fmt"
)

func main()  {

	singleS := "MaazShaikh"

	multiS := `Maaz 
welcome 
`
	fmt.Println(singleS)
	fmt.Println(multiS)

}

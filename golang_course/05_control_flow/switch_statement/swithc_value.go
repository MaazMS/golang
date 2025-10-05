package switch_statement

import "fmt"

func SwitchValueCase() {

	singleValue()
	multipleValue()
}

func multipleValue() {
	language := "golang"
	switch language {
	case "Shaikh", "Maaz":
		fmt.Println("This is case shaikh")
	case "user", "password":
		fmt.Println("This is case Maaz")
	case "c", "python", "java", "golang":
		fmt.Println("This golang")
	}
}
func singleValue() {
	s := "Maaz"
	switch s {
	case "Shaikh":
		fmt.Println("This is case shaikh")
	case "Maaz":
		fmt.Println("This is case Maaz")

	}
}

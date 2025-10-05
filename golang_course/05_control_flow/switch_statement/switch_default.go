package switch_statement

import "fmt"

func Switchdefaultcase() {
	fmt.Println("switch case for default")

	switch {

	case false:
		fmt.Println("this should not print")

	default:
		fmt.Println("when all case is false then default is execute")
	}

}

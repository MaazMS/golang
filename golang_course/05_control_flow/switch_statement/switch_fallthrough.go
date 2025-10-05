package switch_statement

import "fmt"

func SwitchfallthroughCase() {

	allTrue()
	oneTureFalse()
	allTrueFalse()
	fallDefault()

}

func fallDefault() {
	fmt.Println("fallthrough use all in true case after and false case ")
	fmt.Println("switch with fallthrough with default  ")
	switch {

	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 != 4:
		fmt.Println("this false case should print3")
		fallthrough
	case 5 != 6:
		fmt.Println("this should not print4")
		fallthrough
	case 6 != 6:
		fmt.Println("this should not print5")
		fallthrough
	case 7 != 7:
		fmt.Println("this should not print6")
		fallthrough
	case 8 != 8:
		fmt.Println("this should not print7")
		fallthrough
	default:
		fmt.Println("this is default ")
	}

}

func allTrueFalse() {

	fmt.Println("fallthrough use all in true case after and false case ")
	fmt.Println("Note `but if you wat use all flase case then at end of all flase case fallthrough use ` ")

	switch {
	case false:
		fmt.Println("this should not print")
		fallthrough
	case 2 == 4:
		fmt.Println("this should not print2")
		fallthrough
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 != 4:
		fmt.Println("this false case should print3")
		fallthrough
	case 5 != 6:
		fmt.Println("this should not print4")
		fallthrough
	case 6 != 6:
		fmt.Println("this should not print5")
		fallthrough
	case 7 != 7:
		fmt.Println("this should not print6")
		fallthrough
	case 8 != 8:
		fmt.Println("this should not print7")
		fallthrough
	case 9 != 9:
		fmt.Println("this should not print8")
	}
}

func oneTureFalse() {

	fmt.Println("fallthrough in true case after and false case ")
	fmt.Println("Note `fallthrough is not print false case ` ")

	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print2")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 != 4:
		fmt.Println("this false case should print3")
	case 5 != 6:
		fmt.Println("this should not print4")
	case 6 != 6:
		fmt.Println("this should not print5")
	case 7 != 7:
		fmt.Println("this should not print6")
		fallthrough
	case 8 != 8:
		fmt.Println("this should not print7")
	case 9 != 9:
		fmt.Println("this should not print8")
	}
}
func allTrue() {
	fmt.Println("fallthrough in true case after that all case is true")
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print2")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("also true, does it print?")
	}
}

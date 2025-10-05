package switch_statement

import "fmt"

func SwitchExpressionCase() {
	fmt.Println("switch case for expression")

	switch {

	case (1 == 2):
		fmt.Println("1 == 2 case is not execute")
	case (2 == 2):
		fmt.Println(" 1=2 case is execue")

	}
}

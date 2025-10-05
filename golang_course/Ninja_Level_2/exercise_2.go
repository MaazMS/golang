package main
/*
Using the following operators, write expressions and assign their values to variables:
1. ==
2. <=
3. >=
4. !=
5. <
6. >
Now print each of the variables
*/

import (
	"fmt"
)

func main()  {

	a := (42 == 42 )
	b := (42 <= 42 )
	c := (42 >= 42 )
	d := (42 != 42 )
	e := (42 > 42 )
	f := (42 < 42 )

	fmt.Println(a, "    ", b , "     ", c, "     ", "    ", d, "     ", "    ", e, "   ", f)


}
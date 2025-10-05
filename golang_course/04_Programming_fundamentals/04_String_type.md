##  String type  
1. A string type represents the set of string values.  
1. A string value is a (possibly empty) sequence of bytes.  
1. The number of bytes is called the length of the string and is never negative.  
1. Strings are immutable: once created, it is impossible to change the contents of a string.   
1. string value in double quote"" and back ticks ``.  
1. single line string use double quote""  
1. Multi line string use back ticks ``.   
1. `%s` value of string.  
1. `` value with double quote   

```go
package main

import (
	"fmt"
)

func main() {

	// string value in double quote""
	s := "Hello World"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)

	// string value in back ticks ``
	multis := `Hello 
		world`
	fmt.Printf("%s\n", multis)
	fmt.Printf("%q\n", multis)

	// Iterator string using range
	for index , v := range s {
		fmt.Printf("The index number of %c is %d\n", v, index)
	}
	// Iterator string using length of string
	for i := 0 ;i < len(s) ; i++ {
		fmt.Printf("%c\n", s[i])
	}

	// A string value is a (possibly empty) sequence of bytes.
	// byte which is an alias for uint8
	bs := []byte(s)
	fmt.Println(s)
	fmt.Printf("%T\n", bs)
}


```
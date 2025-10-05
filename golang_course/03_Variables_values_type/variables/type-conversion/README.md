## 
1. Type conversion creates a new value by changing the type of the original value to another type.    
1. There are also function values constant values interface values, and the list goes on type conversion use.  
1. So the conversion expression is not only for converting the type of variables you can use it to convert
almost any value to any other type as of course.   
1. syntax `type(value)`

```go
package main

import "fmt"

func main() {

	speed := 100 // speed is int
	force := 2.5 // force is float64

	fmt.Println(speed, force)
	clculate := speed *  int(force)
	fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force, clculate)

	//clculate := float64(speed) *  force // Cannot use 'float64(speed) * force' (type float64) as type int
	//fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force, clculate)
	fmt.Println()

	clculate = int (float64(speed) *  force)
	fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force, clculate)

}
```
1. if display force, then it values is `2.5` this means that the conversion creates a new value, but it doesn't change the actual value.  

## type names
1. types with different names are different types.  
```go
package main

import "fmt"

func main() {

	var speed int // speed is int
	var force int32 // force is int32

	// speed = force  // You can not assign different type to different name
	speed = int(force)
	fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force)
	fmt.Println()

	//decision := bool(speed) // You can not convert int value to bool
	//fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, decision)

	capitalA := 65
	characterA := string(capitalA)
	fmt.Printf("capitalA `%T` characterA `%T` :",capitalA, characterA)

	//capitalB := 66.6
	//characterB := string(capitalB) // Cannot convert expression of type 'float64' to type 'string'
	//fmt.Printf("capitalB `%T` characterB `%T` :",capitalB, characterB)
	fmt.Println()



}
```
1. line number `53` this works because as string is actually a series of numbers behind the scenes.  
characterA
## assignment   
```go
package main

import "fmt"

func main()  {

	var (
		name string
		phono int
		force float64
	)
	
	name = "Maaz"
	phono = 123
	force = 1
	fmt.Println(name, phono, force)

	name = "Shaikh"
	phono = 987
	force = 100
	fmt.Println(name, phono, force)
}
```
1. `force` variable is `float64` but I am an assign it `1` and `100`.  
1. It's because, Go can convert a numeric constant to a numeric type automatically.
1. `1` and `100` here is a numeric constant, actually it has no type until you assign it.
1. So, when you assign `1` and `100` to the force variable, it becomes a float64 value automatically.
1. `1` and `100` It's just  literal.   

* tou can only assign a value to a variable with the same type.  


## Multiple assigment 
1. In go expressions can return multiple value.  
1. number of variable on the left side, and the number of values on the right side should match.  

```go
package main

import "fmt"

func main()  {

	var firstname, lastname = "Maaz" , "Shaikh"
	fmt.Println(firstname, lastname)
}
```
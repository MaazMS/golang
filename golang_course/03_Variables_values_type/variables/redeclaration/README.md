## redeclaration variable 
1. redeclaration variable mean declaration of variable again.  
1. Why because for an update the value of variable.     
1. There is two-way of redeclaration variable declaration with `var` keyword and without `var` keyword.  
* variable declaration with `var` keyword
```go
package main

import "fmt"

func main()  {

	var safe bool // initial value is false
	safe, speed := true, 50  // update safe value false to true
	// safe := true // No new variables on left side of :=
	fmt.Println(safe, speed )
}
``` 
1. redeclaration of var keyword variable with new variable without new variable it gives error.`No new variables on left side of :=`   

* variable declaration without `var` keyword 
```go
package main

import "fmt"

func main()  {

	name := "Maaz"  // initial value is Maaz
	name = "Shaikh" // upadate the name variable value is to shaikh

	fmt.Println(name )
}
``` 
1. redeclaration of without var keyword variable, first create variable `name := "Maaz"`  
1. update variable remove colon i.e `name = "Shaikh"`

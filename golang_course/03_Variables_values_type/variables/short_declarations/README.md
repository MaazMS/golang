## inference 
1. type inference means that Gore can figure out the type of the variable automatically.
1. Also, by looking at the type of the initial value of the variables you are declaring it.

```go
package main

import "fmt"

func main() {
	var name string = "Maaz"
    //	var name = "Maaz"
	fmt.Println(name)
}
```
1. if click on string it shows `Type Type can be omitted can be omitted` 

## short declaration of variable
1. Remove var keyword for short declaration. 
1. use colon and equal sign.  
1. syntax variable_name := value i.e name := "Maaz" 
```go
package main

import "fmt"

func main() {
	name := "Maaz"
    
	fmt.Println(name)
}
``` 
## why can not use short declaration at package level 
1. package level variable declare using var keyword.  
1. if use short declartion means without using var keyword it give error.  
```go
package main

import "fmt"

var number = 10 
//  number := 10
func main() {
	name := "Maaz"
    
	fmt.Println(name)
}
``` 
1.The reason is that package level statement is start from a keyword i.e `package`, `import` , `func` that way package level   
variable is declared with keyword `var`.  

## Multiple short declaration  
1. first value is assign to first variable and second value is assign to second variable and so on. 
```go
package main

import "fmt"

func main() {
	name, number := "Maaz", 10
    
	fmt.Println(name)
	fmt.Println(number)
}
``` 
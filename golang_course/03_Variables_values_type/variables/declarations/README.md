## declarations  
1. The rules of the variable declaration syntax of Go. you need to declare a variable before you can use it.  
1. syantax of variable `var variable_name type`  i.e `var speed int`.    
1. This "var" keyword means variable. It simply tells Go that let's declare a variable.  
```go
package main

import  "fmt"

func main() {
	var speed int
	fmt.Println(speed)
}
```   
1. speed is the name of this variable.  
1. names are also called identifiers. A name or an identifier helps you and Go to understand each variable you are referring to.  
1. int is a type of this variable. It is short for integer.   
1. In Go every variable and value has a type. It's because Go is a strongly typed programming language.   
1. when you print speed it print 0.     
1. It's because a variable is initialized to a zero value when it's been declared without an initial value.
## Rule of variable name  
1. The first rule is that a name should always start with a letter or an underscore character.  
1. when the first letter of a name is an uppercase letter that name becomes exported so that other packages can use it.   
   That's one of the reasons why a name should start with a letter.  
1. names of variable cannot start with a number like this, 1speed.  
1. names of variable can't start with a punctuation character like this, !speed.
1. It's because punctuation characters have special meanings in Go.  
1. names of variablecan't only contain reserved Go keywords. like this var var int.   

## Zero value

```go
package main

import  "fmt"

func main() {
	var speed int
	var heat float64
	var off bool 
	var brand string 
	
	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)
}
```
1. That means they have all been initialized to a value by Go automatically. Automatically because as you  
can see I didn't initialize them to a value while declaring them.  
1. go initialize variables to zero values. 
1. Because if go wouldn't have initialized variables to zero values, and those variables might have contained something  
called garbage values inside the computer memory.  
1. It's called garbage because usually it might not be meaningful data and trying to use garbage value could be very problematic. 

## unused variable 
```go
package main

var name string
func main() {
	var speed int
}
```
1. if speed declared and not used , it shows `speed declared are not used` .    
1. This is because God wants you to write maintainable programs.     
1. By the way this deduction happens at the compile time.  
1. But if comment speed variable the program is executed because name variable at package level.   

## Multiple declaration 
```go
package main

import  "fmt"

func main() { 
	var (
			 speed int   
			 heat float64 
			 
			 off bool
			 brand string
	)
	

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)
}

``` 
1.  Multiple declaration `var ()` in side parentheses give name of variable.   
1. if give space between two variable mean `speed` and `heat` variable related to each other  and `off` `brand` variable related.  

```go
package main

import  "fmt"

func main() {

   var name , address string

   fmt.Println(name)
   fmt.Println(address)
}
```
1. Declare same type of variable in single line and separate by comma. 
1. It is also called parallel variable declaration, 


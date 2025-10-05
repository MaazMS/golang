## Zero value  
1. Declare the variable with `var` keyword but not assign value to it then assign default value of that data type of variable.  

## default with their data type  
1. false for booleans
1. 0 for integers
1. 0.0 for floats
1. "" for strings
1. nil for
    1.  pointers
    1.  functions
    1.  interfaces
    1.  slices
    1.  channels
    1.   maps 
   
```go
// DECLARE a variable to be of a certain TYPE
var number int
var fnumber float64
var name string
var condition bool

func main() {

fmt.Print("\n", number)
fmt.Printf("\t%T\t\n", number)

fmt.Print("\n", fnumber)
fmt.Printf("\t%T\t\n", fnumber)

fmt.Print("\n", name)
fmt.Printf("\t%T\t\n", name)

fmt.Print("\n", condition)
fmt.Printf("\t%T\t\n", condition)

}
```  
1.output of program  
```bash
0	int	

0	float64	

	string	

false	bool	  
```
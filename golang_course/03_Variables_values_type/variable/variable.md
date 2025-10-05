### variable declaration
1. Three way to declare variable.   
2. If we declare the variable and assign the value but not used it give error.  
`variable name declared but not used`

####  first way  variable declaration 
1. var keyword name of variable datatype    
*  Assign the value to variable after declaration.       
* it declare when not initilaze any thing to variable.    
* Syntax  
``` 
var varname datatype 
	varname = value
``` 
* example 
```   
var no1integer int
	no1integer = 10
````  
####  second way  variable declaration 
1. var keyword name of variable datatype = value     
* Syntax  
``` 
var varname datatype = value

``` 
* example   
`var no2integer int = 20`  

####  Third way  variable declaration 
1. name of variable colon= value    
* Syntax  
``` 
 varname := value 
``` 
* example  
	no3integer := 30   

### package level varilable   
1. declare variable at package level.    
* Syntax  
``` 
var varname datatype = value

``` 
* example   
`var no2integer int = 20`    

* syntax error: non-declaration statement outside function body when use name of variable colon= value at package level variable.    
1. This is not use at package level.       
* Syntax  
``` 
 varname := value 
``` 
* example  
	no3integer := 30  
``` 
2. declare group of varible as first way go use declare variable as a group.   
```Example  
var (
	firstname string = "Maaz"
	lastname  string = "Shaikh"
	age       int    = 23
)
```



### Redeclaration and shadowing     

**Redeclaration**
* if delare a variable with some value (function level) it give error.Two error show 
1. error : variable name redeclared in this block  
2. error: no new variables on left side of :=   

**shadowing**
1. if delare a variable at package value and same variable is declare in function level, It not give error and this is called shadowing.     
2. Before declare variable in function we can access  package value variable.   



### visibility   
* name of variable is control the visibility   
* There is three level of visibility of varilable.  
1. package level(lower case variable)  
2. global level (upper case variable)     
3. block scope (eg. fuction block)     

**NOTE1**  
* The length of the variable should reflect the scope of variable.    
a. small lenght (one letter ) use in for loop, inside function etc.  
b. The big length for use log time. example for package level variable.    

**NOTE2**  
* Acronym is the short name create the first letter of each word.  
* Example National Aeronautics and Space Administration. NASA             
a. To write acronym in go all are upper case. (because of readability)   

### Naming conventions    
* name of variable is itself naming convension.     
1. lower case variable is actually package leve scope.    
2. upper case variable is expose the outside of the world.     

### Type conversions  
1. Type conversion use **import "strconv"** only type conversion `int` to `string`     
2. `strconv.Itoa` **Itoa** function which is equivalent to FormatInt which returns the string as base is 10.  
3. Type conversion `int` to `string`   but print the data type is always show string.       
* Example   
``` 
fmt.Println("int to String conversion")
	var m int = 16
	fmt.Printf("%v, %T  \n", m, m)
	var n string
	n = strconv.Itoa(m)
	fmt.Printf("%v, %T  \n", n, n)
```
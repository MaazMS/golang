## Short declaration operator
1. we can declare and assign value to variable using short declaration.
1. For an assign new value for Short declaration variable we use assignment operator `=`  
1. Short declaration of variable use only inside curly braces `{ }`   
```go
package main

import "fmt"

func main() {
	// Short declaration operator
	// Declare a variable and assign value (of certain type) 
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	// Short declaration operator
	// Declare a variable and assign value (of certain type)
	y := 100 + 10
	fmt.Println(y)
}

```
## terminology  
* [Identifiers](https://golang.org/ref/spec#Identifiers)  
1. Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits.   
   The first character in an identifier must be a letter.  
   
* [Predeclared identifiers](https://golang.org/ref/spec#Predeclared_identifiers)  
```1bash
Types:
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
``` 
1. [Keywords](https://golang.org/ref/spec#Keywords)      
1. these are words that a reserved for use by the Go programming language they are sometimes called “reserved words”.  
   you can’t use a keyword for anything other than its purpose.  
```bash
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```  

1. [Operators and punctuation](https://golang.org/ref/spec#Operators_and_punctuation)    
1.    operator
   1. in “2 + 2” the “+” is the OPERATOR
   1. an operator is a character that represents an action, as for example “+” is an arithmetic OPERATOR that represents addition.  
1. operand
   1.  in “2 + 2” the “2”s are OPERANDS
```bash
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

1. statement  
    1. In programming a statement is the smallest standalone element of a  
       program that expresses some action to be carried out. It is an instruction  
       that commands the computer to perform a specified action. A program is  
       formed by a sequence of one or more statements.  
       
1. expression  
    1.  in programming an expression is a combination of one or more explicit  
        values, constants, variables, operators, and functions that the  
        programming language interprets and computes to produce another  
        value. For example, 2+3 is an expression which evaluates to 5.   
        

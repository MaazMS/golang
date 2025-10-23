## Short declaration operator

1. You can declare and initialize variables using the short declaration operator `:=`.
1. To assign a new value to an existing variable, use the assignment operator `=`.  
1. The short declaration operator can be used only inside function bodies (i.e., within curly braces `{}`), not at package scope.  
```go
package main

import "fmt"

func main() {
    // Short declaration operator: declare and initialize
    x := 42
    fmt.Println(x)

    // Reassign using '='
    x = 99
    fmt.Println(x)

    // Expression initialization
    y := 100 + 10
    fmt.Println(y)

    // Multiple variables at once (types inferred)
    a, b := 1, "gopher"
    fmt.Println(a, b)

    // Initialize from function returns
    n, err := fmt.Println("hello")
    fmt.Println("bytes:", n, "err:", err)
}

```
## Notes and pitfalls

- The `:=` operator requires at least one new variable on the left-hand side. If all names already exist in the current scope, use `=` instead.
- Redeclaration with `:=` is allowed only when introducing at least one new variable and all existing variables are in the same scope. Example:
  ```go
  v := 1   // declare v
  v, w := 2, 3 // ok: v is reassigned, w is new
  // v, w := 4, 5 // compile error in the same scope: no new variables
  ```
- Avoid accidental shadowing: using `:=` inside an inner block can create a new variable that hides an outer one.
  ```go
  count := 10
  if true {
      count := 1 // shadows outer count
      fmt.Println(count) // 1
  }
  fmt.Println(count) // 10
  ```
- `:=` is commonly used in the short statements of `if`, `for`, and `switch`.
  ```go
  if v, err := do(); err != nil {
      return err
  } else {
      fmt.Println(v)
  }
  for i := 0; i < 3; i++ { fmt.Println(i) }
  switch x := compute(); x {
  case 0: fmt.Println("zero")
  }
  ```

## terminology  
* [Identifiers](https://go.dev/ref/spec#Identifiers)  
1. Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits.  
   The first character in an identifier must be a letter.  
   
* [Predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers)  
```text
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
1. [Keywords](https://go.dev/ref/spec#Keywords)      
1. These are words reserved by the Go language ("reserved words").  
   You can’t use a keyword for anything other than its purpose.  
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

## var vs short declaration

- Use `var` at package scope, or when you need an explicit type or zero value without initialization.
  ```go
  var n int        // zero value 0
  var s string     // zero value ""
  var r io.Reader  // zero value nil
  ```
- Use `:=` inside functions when you want concise declaration with inferred types.

1. statement  
    1. In programming a statement is the smallest standalone element of a  
       program that expresses some action to be carried out. It is an instruction  
       that commands the computer to perform a specified action. A program is  
       formed by a sequence of one or more statements.  
       
2. expression  
    1.  in programming an expression is a combination of one or more explicit  
        values, constants, variables, operators, and functions that the  
        programming language interprets and computes to produce another  
        value. For example, 2+3 is an expression which evaluates to 5.   
        
3. Where `var` can be used  
    1. At package scope (top-level) and inside functions.  
    1. Use `var` when you need an explicit type, a zero value without initialization, or when declaring at package scope (since `:=` is not allowed there). 

---


## `var` vs `:=`

- Use `var` at package scope (top-level); short declaration `:=` is not allowed there.
- Use `var` when you need an explicit type or a zero value without initialization.
- Use `:=` inside functions for concise declaration with inferred types.

```go
package main

import "fmt"

var version = "v1"       // package scope requires var

func main() {
    x := 10              // short declaration inside a function
    var y int            // zero value 0
    var z int = x + 5    // explicit type with initializer
    fmt.Println(version, x, y, z)
}
```


## Note: Default (zero) values in Go

When variables are declared without an explicit initializer, Go assigns their type’s zero value:

- Booleans: `false`
- Numeric (integers: `int`, `int8`, `int16`, `int32`, `int64`, unsigned variants, `uintptr`): `0`
- Floating point (`float32`, `float64`): `0`
- Complex (`complex64`, `complex128`): `0+0i`
- Strings: `""` (empty string)
- Pointers, functions, interfaces, slices, channels, maps: `nil`
- Arrays: each element is the element type’s zero value
- Structs: each field is the field type’s zero value
- Aliases/defined types: zero value of the underlying type (e.g., `rune` -> `0`, `byte` -> `0`)

Example:
```go
package main

import (
    "fmt"
    "io"
)

type counter int

type user struct {
    name string
    age  int
}

func main() {
    var (
        b bool       // false
        i int        // 0
        f float64    // 0
        c complex128 // (0+0i)
        s string     // ""
        r io.Reader  // nil
        a [3]int     // [0 0 0]
        u user       // {"" 0}
        k counter    // 0
        m map[string]int // nil
        ch chan int       // nil
        sl []int          // nil
    )
    fmt.Println(b, i, f, c, s == "", r == nil, a, u, k)
    fmt.Println(m == nil, ch == nil, sl == nil)
}
```

## Exported variables

A name starting with an uppercase letter is exported from the package; lowercase is unexported.

```go
package config

var AppName = "Demo" // exported
var version = "v1"   // unexported
```

---


## Grouped declarations

```go
var (
    host = "localhost"
    port = 8080
)

var (
    ready bool
    count int
)
```
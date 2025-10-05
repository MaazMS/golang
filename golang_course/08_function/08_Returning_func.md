## Returning  func  
1. You can return a func from a func.    

```go
package main

import "fmt"

func main()  {

	x := foo()
	fmt.Println(x)
}
func foo() string {
	return "Hello world"
} 
/* output
Hello world
*/
``` 

```go
package main

import "fmt"

func main() {

	y := baar()
	fmt.Printf("Type of y is \t%T\n", y)
	fmt.Printf("function call")
	y()

}

func baar() func() int {
	return func() int {
		return 1997
	}
}
/* output
Type of y is    func() int
function call
*/
``` 

```go
package main

import "fmt"

func main() {

	a := work()
	i := a()
	fmt.Println(i)
}


func work() func() string {
	return func() string {
		return "Hello golang"
	}
}

/* output
Hello golang
*/

``` 

```go
package main

import "fmt"

func main() {
	
	fmt.Println(smart()())
}

func smart() func() bool {
	return func() bool {
		return true
	}
}
/* output
true
*/
```
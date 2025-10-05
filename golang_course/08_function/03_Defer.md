## Defer  
1. defer meaning delay an action or proceeding.   
1. A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns,    
   either because the surrounding function executed a return statement, reached the end of its function body, or because   
   the corresponding goroutine is panicking.  
1. execution of program   
    1. go to main()  
    1. go to baar() 
    1. go to foo()
```go
package main

import "fmt"

func main() {

	fmt.Println("func main")
	defer foo()
	bar()
}

func foo() {
	fmt.Println("func foo")
}
func bar() {
	fmt.Println("func bar")

}

``` 


```bash
func main
func bar
func foo
```
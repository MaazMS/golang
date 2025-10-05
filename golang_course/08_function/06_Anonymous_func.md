## Anonymous func    
1. Anonymous func is function have no name.   
1. anonymous func self executing function.


## Anonymous func  with no parameter  
1.  
```go
package main

import "fmt"

func main() {

	// Anonymous func  with no parameter
	func() {
		fmt.Println("Anonymous func  with no parameter")
	}()
}
```

## Anonymous func  with parameter
````go
package main

import "fmt"

func main() {

	// Anonymous func  with parameter
	func (x int){
		fmt.Println("Anonymous func  with parameter \t", x)
	}(45)

}


````
## Closure 
1. closure is a concept we want to close in, close the scope of a variable and contain it to certain areas.   

## package level closure 
```go
package main

import "fmt"

var x int
func main()  {

	fmt.Println("package variable",x)
	x++
	fmt.Println("package variable increment in func main",x)
	foo()
}
func foo(){
	fmt.Println("package variable without increment in func foo",x)
	x++
	fmt.Println("package variable increment in func foo",x)

}
/* output
package variable 0
package variable increment in func main 1
package variable without increment in func foo 1
package variable increment in func foo 2
*/
``` 
##  function level closure
```go
package main

import "fmt"

func main()  {

	foo()
}
func foo(){
	var y int
	fmt.Println("function level variable y", y)
	y++
	fmt.Println("function variable y increment in func foo",y)

} 
/*output
function level variable y 0
function variable y increment in func foo 1
*/
``` 
##  block level closure

```go
package main

import "fmt"


func main()  {


	{
		z:= 10
		fmt.Println("block variable z ",z)
		z++
		fmt.Println("block variable z increment in block ",z)
	}
	// fmt.Println("block variable z is not use outside of block ",z)

} 
/* output
block variable z  10
block variable z increment in block  11

*/
```
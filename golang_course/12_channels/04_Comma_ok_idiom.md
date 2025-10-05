##
1. For obvious reasons this is called the “comma ok” idiom. In this example, if tz is present,    
   seconds will be set appropriately and ok will be true;    
   if not, seconds will be set to zero and ok will be false.   
   
```go
package main

import "fmt"

func main()  {
	c := make(chan int)
	go func(){
		c <- 50
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v , ok = <- c
	fmt.Println(v, ok)
}
/* OUTPUT
50 true
0 false

*/

``` 

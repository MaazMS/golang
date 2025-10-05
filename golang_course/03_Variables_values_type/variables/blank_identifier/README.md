## blank identifier  
```go
package main


var name string
func main() {

	var speed int 
	_ = speed
	
}
```
1. The blank identifier is not an ordinary entity.  
1. It's like a black hole it eats the whalers alive.    
1. go will just accept the fact that the variable has been used. line number 10. `_ = speed`    
1. Normally you used a blank identifier for things like skipping return values from functions and some other useful things.  

## Race condition  
1. A race condition is when two or more routines have access to the same resource,   
1. such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines.    

## Gosched  

when execution context in one goroutine reaches Gosched call, the scheduler is instructed to switch the execution to another goroutine.             
In your case there are two goroutines, main (which represents 'main' thread of the program) and additional, the one you have created with go say.             
If you remove Gosched call, the execution context will never be transferred from the first goroutine to the second, hence no 'world' for you.        
When Gosched is present, the scheduler transfers the execution on each loop iteration from first goroutine to the second and vice versa,         
so you have 'hello' and 'world' interleaved.           
```go
package main

import "runtime"

func main()  {

	go say("World")
	say("hello")
}

func say(s string)  {

	for i :=0 ;i<= 3; i++{
		runtime.Gosched()
		println(i, s)
	}
}
/* output without Gosched
0 hello
1 hello
2 hello
3 hello
*/ 

/* output with Gosched
0 hello
0 World
1 hello
1 World
2 hello
2 World
3 hello

*/

```
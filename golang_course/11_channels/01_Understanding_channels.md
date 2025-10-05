## 
1.  the basic idea like channels this little place where you could send data but I called   
this section channels block in this code is not going to run.    
1. Because you send data at same time you receive data.       
```go
package main

import "fmt"

func main()  {
    // put t2 on c 
	c := make(chan  int)

	c <- 42
    // off of c 
	fmt.Println(<-c)
}


```
```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /home/maaz/github/LearningGO/golang_deep_dive/11_channels/Understanding channels/Understanding channels/main.go:9 +0x59
``` 

## buffer channel    
1. store data on buffer.    

```go
package main

import "fmt"

func main()  {

	// buffer with channel

	d := make(chan int, 1)

	d <- 50
	fmt.Println(<-d)
}
```
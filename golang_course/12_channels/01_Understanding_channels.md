## Understanding channels  
1. A channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.     
1. A channel is a technique which allows letting one goroutine to send data to another goroutine.     
1. They are synchronized.  
1. They have to pass/receive the value at the same time.  
1. 
```bash
making a channel
  c := make(chan int)
● putting values on a channel
  c <- 42
● taking values off of a channel
  <-c
● buffered channels
  c := make(chan int, 4)
```   

1. Because you send data at same time you receive data. There for goroutines are asleep - deadlock         
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
        /home/maaz/github/LearningGO/golang_deep_dive/12_channels/Understanding channels/Understanding channels/main.go:9 +0x59
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


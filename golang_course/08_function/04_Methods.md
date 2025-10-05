## Methods    
1. A method is nothing more than a FUNC attached to a TYPE.  
1. When you attach a func to a    type it is a method of that type.  
1. You attach a func to a type with a RECEIVER.   

```go
package main

import "fmt"

type person struct {
	Name string
	age int
}

func main()  {
	p := person{
		Name: "Maaz",
		age: 24,
	}
	fmt.Println(p)
	p.speak()
}

func (s person) speak() {
}
``` 

```bash
{Maaz 24}
my self is {Maaz 24}
```
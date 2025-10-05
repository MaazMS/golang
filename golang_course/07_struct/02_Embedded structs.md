## Embedded structs
1. Take one struct and embed it in another struct.      
1. `person` is a struct and work is an embedded struct because it has `person` struct.   

```go
package main

import "fmt"

type person struct {
	name string
	age int
}

type work struct {
	person
	filed string
}

func main()  {

	w1 := work{
		person : person{
			name: "Maaz",
			age: 23,
		},
		filed:  "IT",
	}
	fmt.Println(w1.filed, w1.name, w1.age, w1.person)
}

``` 
```bash
IT Maaz 23 {Maaz 23}
```
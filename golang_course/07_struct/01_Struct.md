## struct    
1. A struct is a sequence of named elements, called fields, each of which has a name and a type.   
1. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).       
1. Within a struct, non-blank field names must be unique.  
1. A field declared with a type but no explicit field name is called an embedded field.    
1. A field or method f of an embedded field in a struct x is called promoted if x .  
1. Promoted fields act like ordinary fields of a struct.     
1. A struct is a data structure which allows us to compose together values of different types.  
1. So it's an aggregate data type.  
1. `create the value of person struct`  

```go
package main

import "fmt"

type person struct {
	name string
	age int
}
func main()  {

	p1 := person{
		name: "Maaz",
		age: 20,
	}
	fmt.Println(p1)

	p2 := person{
		name: "Shaikh",
		age: 25,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.name, p1.age)
	fmt.Println(p2.name, p2.age)


}
``` 

```bash
{Maaz 20}
{Maaz 20}
{Shaikh 25}
Maaz 20
Shaikh 25
```
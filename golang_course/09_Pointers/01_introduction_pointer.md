## pointer     
1. All values are stored in memory.    
1. Every location in memory has an address.  
1. A pointer is a memory address.     
1. pointer pointing to some location in memory where a value is stored.  
1. how you see the address in memory`using & operator`.   
1. dereferencing  an address using `* operator`

```go
package main

import "fmt"

func main() {

	a := 42
	fmt.Println("The value of a", a)
	fmt.Println("The address of a", &a)        // & gives you the address
	fmt.Println("The value of using *& ", *&a) // give the value

	fmt.Printf("The Type of a variable %T\n", a)
	fmt.Printf("The Type of a variable memory %T\n", &a)

	// b is of type "int pointer" b points to the memory address where an int is stored
	b := &a
	fmt.Printf("The Type of b variable %T\n", b)
	fmt.Println("The address of b", &b)       // & gives you the address
	fmt.Println("The value of using *& ", *b) // give the value

	*b = 45
	fmt.Printf("The Type of b variable %T\n", b)
	fmt.Println("The address of b", &b)       // & gives you the address
	fmt.Println("The value of using *& ", *b) // give the value
}

/* output

The value of a 42
The address of a 0xc0000200a8
The value of using *&  42
The Type of a variable int
The Type of a variable memory *int
The Type of b variable *int
The address of b 0xc00000e030
The value of using *&  42
The Type of b variable *int
The address of b 0xc00000e030
The value of using *&  45
*/
``` 

## where use pointer  
1. So when should you use pointers, so pointers are good if you have a large chunk of data and you don't   

want to pass that big chunk of data around through your program, you just pass the address where that  

data is stored and  all you're doing is passing an address.   

So that could save you a little performance if you're thinking about it like, wow, this is a huge piece of data.  

We're getting back from the database.   
It's stored in memory now, just past the address around and use the address where you need it.    

1.  change something that's at a certain location,   you could use pointers.  

1. Everything in Go is pass by value. Drop any phrases and concepts you may have from other  
   languages. Pass by reference, pass by copy - forget those phrases. “Pass by value.” That is the   
   only phrase you need to know and remember. That is the only phrase you should use. Pass by value.   
   Everything in Go is pass by value.    
```go
package main

import "fmt"

func main() {

	a := 200
	fmt.Println("a before ", &a)
	fmt.Println("a before ", a)
	bar(&a)
	fmt.Println("a after  ", &a)
	fmt.Println("a after  ", a)

}

func bar(b *int) {
	fmt.Println("b before ", *b)
	fmt.Println("b before ", b)
	*b = 500 // assign the of b datatype is *int
	fmt.Println("b after  ", *b)
	fmt.Println("b after  ", b)

}

/* output
a before  0xc0000200a8
a before  200
b before  200
b before  0xc0000200a8
b after   500
b after   0xc0000200a8
a after   0xc0000200a8
a after   500
*/

``` 


```go
package main

import "fmt"

func main() {
	a := 45
	fmt.Println("a before ", &a)
	fmt.Println("a before ", a)
	foo(a)
	fmt.Println("a after  ", &a)
	fmt.Println("a after  ", a)
}

func foo(b int) {
	fmt.Println("b before ", &b)
	fmt.Println("b before ", b)
	b = 50
	fmt.Println("b after  ", &b)
	fmt.Println("b after  ", b)

}

/* output
a before  0xc0000200a8
a before  45
b before  0xc0000200d0
b before  45
b after   0xc0000200d0
b after   50
a after   0xc0000200a8
a after   45

*/

```
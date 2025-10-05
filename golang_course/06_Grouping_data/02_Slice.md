## Slice 

1. Array is allow you to group together value of same TYPE and changes in size.    
1. Slice is built in array. 


```go
package main

import "fmt"

func main() {

	//  variable := type {values}

	no := []int{1, 2, 3, 5, 7, 8, 9}
	fmt.Println(no)
}

``` 
```bash
[1 2 3 5 7 8 9]

``` 

## Slice iterates   
1. A "for" statement with a "range" clause iterates through all entries of slice.  

```go
package main
import "fmt"
func main() {
	no := []int{1, 2, 3, 5, 7, 8, 9}
	fmt.Println(no)
	fmt.Println(len(no))
	fmt.Println(no[0])
	fmt.Println(no[1])
	fmt.Println(no[2])
	fmt.Println(no[3])
	fmt.Println(no[4])

	for i, v := range no {
		fmt.Println(i,  v)
	}
}
``` 
```bash
[1 2 3 5 7 8 9]
7
1
2
3
5
7
0 1
1 2
2 3
3 5
4 7
5 8
6 9
```
## slicing a slice   
1. We can slice a slice which means that we can cut parts of the slice away. We do this with the colon operator.  

```go
package main
import "fmt"
func main()  {
    no := []int {1, 2, 3, 5, 7, 8, 9}
	fmt.Println(no)
	fmt.Printf("Access index by value %d\n",no[1])
	fmt.Printf("slice a slice from index 1 include to end of slice %d\t\n", no[1:])
	fmt.Printf("Access index by value %d\n",no[4])
	fmt.Printf("slice a slice from index 1 include to exclude last index of  slice %d\t\n", no[1:4])

} 

``` 
```bash
[1 2 3 5 7 8 9]
Access index by value 2
slice a slice from index 1 include to end of slice [2 3 5 7 8 9]        
Access index by value 7
slice a slice from index 1 include to exclude last index of  slice [2 3 5]  
``` 

## slice Append  
1. append function have two parameter.  
1. first parameter where to append values, eg  slice a variable name,  
1. second is the value to put into slice variable.  
1. optional but important try to put whole slice to another slice use `variadic` parameter i.e `...`    

```go
package main
import "fmt" 
func main()  {

	no := []int {1, 2, 3, 5, 7, 8, 9}
	fmt.Println(no)
	no = append(no, 11,12,13)
	fmt.Println(no)

	even := []int{2,4,6,8,10}
	no = append(no, even...)
	fmt.Println(no)

	// append one slice by their index
	no = append(no[:2], no[:4]...)
	fmt.Println(no)
}
```
```bash
[1 2 3 5 7 8 9]
[1 2 3 5 7 8 9 11 12 13]
[1 2 3 5 7 8 9 11 12 13 2 4 6 8 10]
[1 2 1 2 3 5]
```

## Delete in slice  
1.  Delete a slice using their index

```go
package main
import "fmt"
func main()  {

	no := []int {1, 2, 3, 5, 7, 8, 9}
	fmt.Printf("Original slice\n")
	for i := 0; i< len(no); i++ {
		fmt.Printf("%d\t",no[i])
	}
	fmt.Println()
	fmt.Printf("deleteInslice \n")

	no = append(no[:2], no[:4]...)

	for i := 0; i< len(no); i++ {
		fmt.Printf("%d\t",no[i])
	}
}

```

````bash
[1 2 3 5 7 8 9]
Original slice
1       2       3       5       7       8       9       
deleteInslice 
1       2       1       2       3       5 
````  

## make function  
1. It creates a slice again and again if you append the value inside the slice.     
```go
package main
import "fmt"
func main()  {

	no := make([]int, 10, 100)
	fmt.Println(no)
	fmt.Println(len(no))
	fmt.Println(cap(no))
	no[9] = 99

	no = append(no, 15)
	fmt.Println(no)
	fmt.Println(len(no))
	fmt.Println(cap(no))

}
```
```bash
[1 2 3 5 7 8 9]
[0 0 0 0 0 0 0 0 0 0]
10
100
[0 0 0 0 0 0 0 0 0 99 15]
11
100
``` 
## multi-dimensional slice  
1. multiple slice with same datatype.  
```go
package main
import "fmt"
func main() {
	m1 := []string{"Maaz", "Shaikh"}
	m2 := []string{"golang", "python "}

	fmt.Println(m1)
	fmt.Println(m2)

	m3 := [][]string{m1, m2}
	fmt.Println(m3)
}
```

```bash
[Maaz Shaikh]
[golang python ]
[[Maaz Shaikh] [golang python ]]
```
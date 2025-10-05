## Map 
1. A map is a key-value store.    
1. This means that you store some value, and you access that value by a key.  
1. For instance, I might store the value “phoneNumber” and access it be the key “friendName”.  
1. **The syntax for a map is map[key]value.**  
1. It is important to note that maps are unordered.     
```go
package main

import "fmt"

func main() {

	m := map[string]int{
		"Maaz":   24,
		"Shaikh": 23,
	}
	fmt.Println(m["Maaz"])
	fmt.Println(m["Shaikh"])

	//  key and value which is not available in map using v , ok idiom
	fmt.Println("key and value which is not available in map using v , ok idiom")
	v, ok := m["Noksy"]
	fmt.Println(v)
	fmt.Println(ok)

	// 	key and value which is not available in map using v , ok idiom and if condition
	fmt.Println("key and value which is not available in map using v , ok idiom and if condition")
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

}

``` 

```bash
24
23
key and value which is not available in map using v , ok idiom
0
false
key and value which is not available in map using v , ok idiom and if condition
``` 

## Add key and value to map  
1. You can add key and value to map     

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"Maaz":   24,
		"Shaikh": 23,
	}

	m["user"] = 500

	for k , v := range m {
		fmt.Println(k,   v)
	}
}

``` 

```bash
Maaz 24
Shaikh 23
user 500
``` 


## Delete map key and value to map
1. You can Delete map key and value to map, 
1. delete function using , map name, and key of map.  
1. if key not exist and delete the using use it not give any error.  

```go
package main

import "fmt"

func main() {

	m := map[string]int {
		"Maaz" : 24,
		"Shaikh" : 23,
	}
	fmt.Println("Before delete map key ")
	for k, v := range m {
		fmt.Println(k,   v)
	}
	delete(m, "Shaikh")
	fmt.Println("\nAfter delete map key ")

	for k, v := range m {
		fmt.Println(k,   v)
	}
}
``` 

````bash
Before delete map key 
Maaz 24
Shaikh 23

After delete map key 
Maaz 24
````
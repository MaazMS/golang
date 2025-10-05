## Loop - init, condition, post 
1. for loop syntax  
```bash
    for initialization; condition ;post {
      
    }
``` 
## Loop - nested loops  
1. for nested loops syntax  

```bash
    for initialization; condition ;post {
          for initialization; condition ;post {
          
          }
    }
```  

## for statement condition 
1. for statement condition syntax  

```bash
    for condition {
          post
    }
``` 

1. example  
```go
package main

import (
	"fmt"
)

func main()  {

	no :=1
	for no <10 {
		fmt.Println(no)
		no++
	}
}

``` 

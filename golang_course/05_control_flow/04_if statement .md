## If Statements
* bool
    1. true
    1. false
* the not operator
    1. !
* initialization statement   
1. syntax 
```bash
if initialization; condition {
		fmt.Println("no is 10 true")
	}
```

```bash
if no := 10; no == 10 {
		fmt.Println("no is 10 true")
	}
```


```go
package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("true statement")
	}
	if false {
		fmt.Println("false statement")
	}
	if !true {
		fmt.Println("!true means false statement")
	}
	if !false {
		fmt.Println("!false means true statement")
	}

	if no := 10; no == 10 {
		fmt.Println("no is 10 true")
	}
}

```
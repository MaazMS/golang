## Number Type  
1. Any number in golang is constant.  
1. `integers` : numbers without decimals  
1. `floating point` : numbers with decimals    
1. `int` 8,16,32,64 are positive and negative number  
1. `unint` 8,16,32,64 are positive number.  

```bash
uint8       the set of all unsigned  8-bit integers (0 to 255)  // 2 the power 8 = 256 
uint16      the set of all unsigned 16-bit integers (0 to 65535) // 2 the power 16 = 65536
uint32      the set of all unsigned 32-bit integers (0 to 4294967295) // 2 the power 32 = 4294967296
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615) // 2 the power 64 = 18446744073709551616

int8        the set of all signed  8-bit integers (-128 to 127) // 2 the power 8 = 256 
int16       the set of all signed 16-bit integers (-32768 to 32767) // 2 the power 16 = 65536
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647) // 2 the power 32 = 4294967296
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807) // 2 the power 64 = 18446744073709551616 

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

```   
1. We can not store a number outer range with data type. i.e
    1. int8 range -128 to 127 we can not store -129. It through error.

1. There is also a set of predeclared numeric types with implementation-specific sizes:

```bash 
uint     either 32 or 64 bits
int      either 32 or 64 bits
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

```  
1.  all numeric types are distinct except  
    1. byte which is an alias for uint8  
    1. rune which is an alias for int32  
* nice reading - Caleb Doxsey’s [book](https://www.golang-book.com/books/intro)


```go
package main

import (
	"fmt"
)

func main() {

	// Declare variable with certain TYPE int8 , range is -128 to 127
	var a int8
	var b int8
	var c int8
	var d int8

	a = 127
	b = 128
	// constant 128 overflows int8
	c = -128
	// constant -129 overflows int8
	d = -129
    
	
	fmt.Println(" variable a ", a)
	fmt.Println(" variable b ", b)
	fmt.Println(" variable c ", c)
	fmt.Println(" variable d ", d)  
	

}

```
package main

import "fmt"

func main() {
	fmt.Println("=== Go Operators Examples for Beginners ===\n")

	// 1. ARITHMETIC OPERATORS
	fmt.Println("1. ARITHMETIC OPERATORS")
	fmt.Println("========================")

	a := 10
	b := 3

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Addition (a + b): %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction (a - b): %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication (a * b): %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Division (a / b): %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Modulus (a %% b): %d %% %d = %d\n", a, b, a%b)

	// Float division example
	c := 10.0
	d := 3.0
	fmt.Printf("Float Division (c / d): %.2f / %.2f = %.2f\n\n", c, d, c/d)

	// 2. ASSIGNMENT OPERATORS
	fmt.Println("2. ASSIGNMENT OPERATORS")
	fmt.Println("=======================")

	x := 5
	fmt.Printf("Initial value of x: %d\n", x)

	// Simple assignment
	x = 10
	fmt.Printf("After x = 10: %d\n", x)

	// Compound assignments
	x += 5 // x = x + 5
	fmt.Printf("After x += 5: %d\n", x)

	x -= 3 // x = x - 3
	fmt.Printf("After x -= 3: %d\n", x)

	x *= 2 // x = x * 2
	fmt.Printf("After x *= 2: %d\n", x)

	x /= 4 // x = x / 4
	fmt.Printf("After x /= 4: %d\n", x)

	x %= 3 // x = x % 3
	fmt.Printf("After x %%= 3: %d\n\n", x)

	// 3. RELATIONAL OPERATORS
	fmt.Println("3. RELATIONAL OPERATORS")
	fmt.Println("=======================")

	num1 := 10
	num2 := 20

	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Printf("num1 == num2 (Equal): %t\n", num1 == num2)
	fmt.Printf("num1 != num2 (Not Equal): %t\n", num1 != num2)
	fmt.Printf("num1 > num2 (Greater): %t\n", num1 > num2)
	fmt.Printf("num1 < num2 (Less): %t\n", num1 < num2)
	fmt.Printf("num1 >= num2 (Greater or Equal): %t\n", num1 >= num2)
	fmt.Printf("num1 <= num2 (Less or Equal): %t\n\n", num1 <= num2)

	// 4. LOGICAL OPERATORS
	fmt.Println("4. LOGICAL OPERATORS")
	fmt.Println("====================")

	age := 25
	hasLicense := true
	hasCar := false

	fmt.Printf("age = %d, hasLicense = %t, hasCar = %t\n", age, hasLicense, hasCar)

	// Logical AND (&&)
	canDrive := age >= 18 && hasLicense
	fmt.Printf("Can drive (age >= 18 && hasLicense): %t\n", canDrive)

	// Logical OR (||)
	canTravel := hasCar || hasLicense
	fmt.Printf("Can travel (hasCar || hasLicense): %t\n", canTravel)

	// Logical NOT (!)
	cannotDrive := !hasLicense
	fmt.Printf("Cannot drive (!hasLicense): %t\n\n", cannotDrive)

	// 5. BITWISE OPERATORS
	fmt.Println("5. BITWISE OPERATORS")
	fmt.Println("===================")

	bit1 := 12 // Binary: 1100
	bit2 := 10 // Binary: 1010

	fmt.Printf("bit1 = %d (Binary: %b)\n", bit1, bit1)
	fmt.Printf("bit2 = %d (Binary: %b)\n", bit2, bit2)

	// Bitwise AND
	andResult := bit1 & bit2
	fmt.Printf("Bitwise AND (bit1 & bit2): %d (Binary: %b)\n", andResult, andResult)

	// Bitwise OR
	orResult := bit1 | bit2
	fmt.Printf("Bitwise OR (bit1 | bit2): %d (Binary: %b)\n", orResult, orResult)

	// Bitwise XOR
	xorResult := bit1 ^ bit2
	fmt.Printf("Bitwise XOR (bit1 ^ bit2): %d (Binary: %b)\n", xorResult, xorResult)

	// Left Shift
	leftShift := bit1 << 2
	fmt.Printf("Left Shift (bit1 << 2): %d (Binary: %b)\n", leftShift, leftShift)

	// Right Shift
	rightShift := bit1 >> 2
	fmt.Printf("Right Shift (bit1 >> 2): %d (Binary: %b)\n", rightShift, rightShift)

	// Bitwise AND NOT
	andNotResult := bit1 &^ bit2
	fmt.Printf("Bitwise AND NOT (bit1 &^ bit2): %d (Binary: %b)\n\n", andNotResult, andNotResult)

	// 6. MISCELLANEOUS OPERATORS
	fmt.Println("6. MISCELLANEOUS OPERATORS")
	fmt.Println("==========================")

	// Address of operator (&)
	value := 42
	address := &value
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Address of value (&value): %p\n", address)
	fmt.Printf("Value at address (*address): %d\n", *address)

	// Pointer dereference
	*address = 100
	fmt.Printf("After changing value through pointer: %d\n", value)

	// Channel receive operator (basic example)
	ch := make(chan int, 1)
	ch <- 50         // Send value to channel
	received := <-ch // Receive value from channel
	fmt.Printf("Received from channel: %d\n\n", received)

	// 7. OPERATOR PRECEDENCE EXAMPLE
	fmt.Println("7. OPERATOR PRECEDENCE EXAMPLE")
	fmt.Println("==============================")

	result1 := 2 + 3*4     // Multiplication first, then addition
	result2 := (2 + 3) * 4 // Parentheses override precedence

	fmt.Printf("2 + 3 * 4 = %d (multiplication first)\n", result1)
	fmt.Printf("(2 + 3) * 4 = %d (parentheses first)\n", result2)

	// Complex precedence example
	complex := 10 + 5*2 - 3/1
	fmt.Printf("10 + 5 * 2 - 3 / 1 = %d\n", complex)
	fmt.Println("Order: 5*2=10, 3/1=3, 10+10=20, 20-3=17")
}

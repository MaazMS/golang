# Go Operators Complete Guide

This document provides a comprehensive overview of all operators available in the Go programming language, organized by category with detailed tables for each operator type.

## 1. Arithmetic Operators

Arithmetic operators are used to perform mathematical operations on operands. They return values of type int or float.

| Operator Name | Symbol | Description |
|---------------|--------|-------------|
| Addition | `+` | Adds two operands |
| Subtraction | `-` | Subtracts second operand from first |
| Multiplication | `*` | Multiplies two operands |
| Division | `/` | Divides first operand by second |
| Modulus | `%` | Returns remainder after division |

## 2. Assignment Operators

Assignment operators are used to assign values to variables. The left operand must be a variable, and the right operand must be of the same data type.

| Operator Name | Symbol | Description |
|---------------|--------|-------------|
| Simple Assignment | `=` | Assigns value on right to variable on left |
| Add Assignment | `+=` | Equivalent to a = a + b |
| Subtract Assignment | `-=` | Equivalent to a = a - b |
| Multiply Assignment | `*=` | Equivalent to a = a * b |
| Division Assignment | `/=` | Equivalent to a = a / b |
| Modulus Assignment | `%=` | Equivalent to a = a % b |
| Bitwise AND Assignment | `&=` | Equivalent to a = a & b |
| Bitwise XOR Assignment | `^=` | Equivalent to a = a ^ b |
| Bitwise OR Assignment | `\|=` | Equivalent to a = a \| b |
| Left Shift Assignment | `<<=` | Equivalent to a = a << b |
| Right Shift Assignment | `>>=` | Equivalent to a = a >> b |

## 3. Bitwise Operators

Bitwise operators work at the bit level and perform bit-by-bit operations. They return boolean values (true/false).

| Operator Name | Symbol | Description |
|---------------|--------|-------------|
| AND | `&` | Bitwise AND operation |
| OR | `\|` | Bitwise OR operation |
| XOR | `^` | Bitwise XOR operation |
| Left Shift | `<<` | Shifts bits to the left |
| Right Shift | `>>` | Shifts bits to the right |
| AND NOT | `&^` | Bitwise AND NOT operation |

## 4. Logical Operators

Logical operators are used to combine two or more conditions and return boolean values (true/false).

| Operator Name | Symbol | Description |
|---------------|--------|-------------|
| Logical AND | `&&` | Returns true if both conditions are true |
| Logical OR | `\|\|` | Returns true if at least one condition is true |
| Logical NOT | `!` | Returns true if condition is false |

## 5. Relational Operators

Relational operators are used to compare two values and return boolean values (true/false).

| Operator Name | Symbol | Description |
|---------------|--------|-------------|
| Equal To | `==` | Returns true if operands are equal |
| Not Equal To | `!=` | Returns true if operands are not equal |
| Greater Than | `>` | Returns true if left operand is greater |
| Less Than | `<` | Returns true if left operand is smaller |
| Greater Than or Equal To | `>=` | Returns true if left operand is greater or equal |
| Less Than or Equal To | `<=` | Returns true if left operand is smaller or equal |

## 6. Miscellaneous Operators

Miscellaneous operators include special operators for specific operations in Go.

| Operator Name | Symbol | Description |
|---------------|--------|-------------|
| Address Of | `&` | Returns the memory address of a variable |
| Pointer Dereference | `*` | Provides access to the value pointed to by a pointer |
| Channel Receive | `<-` | Receives a value from a channel |

## Operator Precedence

Operators in Go follow a specific precedence order when evaluating expressions:

1. **Highest precedence**: `*`, `/`, `%`, `<<`, `>>`, `&`, `&^`
2. **Addition/Subtraction**: `+`, `-`
3. **Comparison**: `==`, `!=`, `<`, `<=`, `>`, `>=`
4. **Logical AND**: `&&`
5. **Logical OR**: `\|\|`
6. **Assignment**: `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `\|=`, `^=`, `<<=`, `>>=`

## Notes

- All operators return appropriate types based on their operands
- Bitwise operators work on integer types only
- Logical operators work on boolean expressions
- Assignment operators modify the left operand
- Use parentheses to override operator precedence when needed

## What is Compile Time?

Compile time is the phase when a program's source code (Go program) is converted to binary/machine language that the computer can execute.

## How to Create Compiled Code?

1. Write a Go program:
```go
package main

import "fmt"

func main() {
	fmt.Println("Maaz Shaikh")
}
```

2. Open terminal and run `go build`
   - `go build` compiles packages and dependencies

## How to Execute Compiled Code?

Run the executable file:
```bash
./executable_file_name
```

## How to Identify an Executable File?

1. Use `ls -l` command:
   ```
   -rwxrwxr-x 1 maaz maaz 2034781 Feb 15 10:50 01-print-names
   ```
   The `x` permission means it is an executable file.

2. Use `file` command:
   ```bash
   file 01-print-names
   ```

## What is Runtime?

Runtime is the phase when the compiled code starts executing and running.

## How to Perform Runtime?

Use `go run` command:
```bash
go run main.go
```
This compiles and runs the program in one step.

## What Happens Behind `go run`?

`go run` performs these steps internally:
1. **Compile** the source code to machine code
2. **Save** the compiled code to a temporary directory (like `/tmp/go-build...`)
3. **Link** all required packages and dependencies
4. **Execute** the program
5. **Clean up** the temporary files

Unlike `go build`, `go run` does not create a permanent executable file in your current directory.

## What is the Difference Between `go build` and `go run`?

- `go run` compiles and runs a program in one step
- `go build` only compiles the program, creating an executable file

## Where Does Go Save the Compiled Code?

Go saves the compiled executable file in the same directory where you run the `go build` command.   

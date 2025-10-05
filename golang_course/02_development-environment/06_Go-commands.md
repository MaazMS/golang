## Go commands  
1. To find and use go command enter terminal `go help`  
````bash
$ go --help
Go is a tool for managing Go source code. 
...
...
````  

1. Search and documentation of specific command   
````bash
go help build
usage: go build [-o output] [-i] [build flags] [packages]

...
....
````
## go run
1. To run go file need file name  
    1. go run <file name> `go run main.go`    
1. Run multiple go file 
   1. go run *.go  `condition is that all go file have only one mian() function `  
    
##  go build     
1. When compiling a single main package, build writes the resulting executable to an output file named after  
   the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
    1. builds the file  
    1. reports errors, if any  
    1. if there are no errors, it puts an executable into the current folder   
    
## go install  
1. Install compiles and installs the packages named by the import paths.
`go install main.go `  
   
## GO BUILD VS GO INSTALL  
1. go build just compile the executable file and move it to the destination.  
1. go install perform the exact operation as go build but places the binary in $GOPATH/bin` directory alongside the binaries    
   of third-party tools installed via go get now if you run $GOPATH/bin/hello you will see Hello, world! printed on your terminal.   
    

   

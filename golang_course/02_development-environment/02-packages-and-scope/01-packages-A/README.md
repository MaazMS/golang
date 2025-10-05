## What is package?  
* First rule of being a package is that all files that belong to a package should be in the same directory.  
* Second rule is that all files that are in the same folder should belong to the same package.   
* The package clause should be at the first line of code. So you can't use it anywhere beyond the first line.   

## How to execute same package of multiple go source files?.  
1.`go run .`    
2. ` go run *.go`   
3. `go run main.go hey.go bye.go`    
* example    
```  
$ go run .
Hi 
goodbye
$ go run *.go
Hi 
goodbye
$ go run main.go hey.go bye.go
Hi 
goodbye
```   

## What are types of packages in Go?.  
Executable packages and library packages.     
1. a file should belong to the package "main"   
2. that file should have a function named "main".      
1. . As you know, its name should always be "package main", and executable package 
should also contain a "main function", but only once.    
2. A library package is created only for usability. So you can import them from other packages    

## difference between library packages and Executable packages    

library | Executable | 
--- | --- |
created for reusability | created for running  
non-executing |executing  
importable | non-importable  
can have any name | name should be main 
no func main | func main  

## Where to store the source code files that belong to a package?  
1. It is store a single directory.  

## Why a package clause is used in a Go source-code file?  
1. It's used for letting Go knows that the file belongs to a package.  

## Where you should put the package clause in a Go source-code file?  
1. As the first code in a Go source code file.    

## How many times you can use package clause for a single source code file?  
1. Once.  

## package clause?  
1. package main  

## single use of package?  
1. All files belong to the same package can call each other's functions   

## Which package type go run can execute?
1. Executable package  

## Which package type that go build can compile?  
1. Both of executable and library packages  

## Which package is used for reusability and can be imported?  
1. Library package   
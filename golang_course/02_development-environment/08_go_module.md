## go module mode  
1. Starting in Go 1.13, module mode will be the default for all development.  

## Creating a new module  
1. `go mod init` creates a new module, initializing the go.mod file that describes it.  

1. create a new source file, hello.go:  

```go
package hello

func Hello() string {
    return "Hello, world."
}
``` 

1. Let's write a test, too, in hello_test.go:

```go
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
``` 
1. At this point, the directory contains a package, but not a module, because there is no go.mod file. If we were working in /home/gopher/hello and ran go test now, we'd see:  

```bash
$ go test
PASS
ok  	_/home/gopher/hello	0.020s

``` 

1. go mod init example.com/hello  
```bash

$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
$ go test
PASS
ok  	example.com/hello	0.020s
``` 
1. The go mod init command wrote a go.mod file  

```bash
$ cat go.mod
module example.com/hello

go 1.12
```
## Adding a dependency  

1. Let's update our hello.go to import rsc.io/quote and use it to implement Hello:

```go
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```  
1. Now let’s run the test again:


```bash
go test
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok      example.com/hello       0.006s
```  
1.  The command `go list -m all` lists the current module and all its dependencies:  

```bash
 go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```
1. What is direct dependency?
  direct dependency are what I'm importing.  
   
1. What is direct dependency?    
   indicates a dependency is not used directly by this module, only indirectly by other module dependencies.  
   
1. what is go.sum   
   go.sum containing the expected cryptographic hashes of the content of specific module versions    
   
## Upgrading dependencies  

1. we can see we're using an untagged version of golang.org/x/text  
1. before update 
```bash
go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```  
1. update  
```bash
$ go get golang.org/x/text
go: golang.org/x/text upgrade => v0.3.6
``` 
1. upadted  
```bash
go list -m all
example.com/hello
golang.org/x/text v0.3.6
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```  
1. 
```bash
go get rsc.io/sampler
go: rsc.io/sampler upgrade => v1.99.99
go: downloading rsc.io/sampler v1.99.99
$ go test
--- FAIL: TestHello (0.00s)
    ello_test.go:8: Hello() = "", want "Hello, world."
FAIL
exit status 1
FAIL    example.com/hello       0.005s
``` 

1. all version of package   
```bash
go list -m -versions rsc.io/sampler
rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
```  
1. download sepecfic version of pacage, **package_name@version** 
```go get rsc.io/sampler@v1.3.1```
## Introduction to packages
1. We organize our data or information in a computer with folders.   
1. We're used to doing that in a program. Instead of using folders, we're going to use packages.  
1.  package is a way to put similar code together to groups, similar code together.     
1. Packages allow us to organize our code into different packages and that keeps everything organized,

## Search package  
1. Search `ftm` package. 
1. open browser and search    
   `http://godoc.org/fmt`   
1. search `net/http`  package  
1. open browser and search    
   `http://godoc.org/net/http`       
   
## List of packages    
1. List of packages  
1. https://golang.org/pkg/  

## Tip and trick  
variadic parameters  
  1. the “...<some type>” is how we signify a variadic parameter  
  1. the type “interface{}” is the empty interface 
      1. every value is also of type “interface{}”  
         1. so the parameter “...interface{}” means “give me as many arguments of any type as you’d like  
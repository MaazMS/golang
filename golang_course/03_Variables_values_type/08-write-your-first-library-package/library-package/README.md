## 
1. Actually every usable package doesn't have to be compiled like an executable package.  
1. It can be directly imported into another package.  
1. install this package `printer` into your go a path.  
1. To do that you need to type "go install".  
1. see package `type "ls ~/go/pkg"`  
1. "printer.a". "a" means archive.  

## Exporting   
1. Exporting : Allows a package to make its functionalities available to another package.   
1. Normally, a package can't access any functionalities from another package whether it imports the package or not.  
1. So, when a package wants to be used by other packages then it should export names.  
1. Those names can belong to a variable, to a function, to a new type,and so on.  
1. exporting is very easy. You just have to make a declared name's first letter an uppercase letter and that's it.  
1.Then it will be exported.  
1. So, the other packages will be able to use it.  
1. There are no "public" or "private" keywords as most in other languages.       
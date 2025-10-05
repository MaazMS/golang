## What's a scope?  
1. The visibility of the declared names of variable, constant and function.  

## what is the scope library?     
1.` import fmt` fmt library is use in file scope where it declares.  
1. In single package where `import fmt` is not declared in a file, then fmt library is not used in this file.     
1. importing only happens at the file level.
1.We import packages into a file not into a package.  
1. importing happens at the file scope only.  
1. multiple files in package level not use import fmt one file to another file.   
## what are the package level scope?
1. package level variable declaration constant declaration and function declaration. 

## How to declare variable for block scope or function scope?.    
1.  variable declare in side of function have function scope and block scope.  
``` 
func main()
{
    var a = 1
}
``` 
## Definition of package level scope.  
1. package scope, means that anything that is declared outside any other functions body will
belong to the package the file belongs to.  
1. packet scope to share functionality between each other.  
1. Declarations need to be outside or any functions.  
   
### What happens to declare same function name in one package?
1.  same scope, you can't have two declarations with the same name.   

## How to rename import package?
import "fmt" 
import f "fmt"
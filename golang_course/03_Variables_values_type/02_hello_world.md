1. Every go file being with package name.
1. The name of the package be the same as the folder name **except package main**.
1. Package main is the entry point of the program.
1. If you are using code form other package you list the package that you want to `import`

## main() Function 
1. `func main()` is entry point of the program.    
1. It does not take any argument nor return anything.  
1. Go automatically call main() function, so there is no need to call main() function explicitly.  
1. Every executable program must contain single main package and main() function.   

## init() Function   
1. init() function is just like the main function.  
1.  It does not take any argument nor return anything.  
1. This function is present in every package.  
1. This function is called when the package is initialized.  
1. This function is declared implicitly, so you cannot reference it from anywhere.   
1. you are allowed to create multiple init() function in the same program and they execute in the order they are created.   
1.  init() function is executed before the main() function call, so it does not depend to main() function.   
1.  The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context.  

Note : When writing code at the end of line not put semicolon it is automatically put by a compiler.   
Note : Entry Point of golang is `func main()` when your code is exits `func main()` you program is over.

## hello world code flow control  
Flow of control
    1. Sequential flow (top to bottom )
    1. loop iterator
    1. conditional 
## GOPATH, Go Workspace, and Go Code Organization  
1. Go requires you to organize your code in a specific way -
By convention, all your Go code, and the code you import, must reside in a single workspace.     
A workspace is nothing, but a directory in your file system whose path is stored in the environment variable GOPATH.  
   
1.  Note that, after the introduction of Go modules in Go 1.11, you’re no longer required to store your Go code in the Go workspace.   
You can create your Go projects in any directory outside of GOPATH.    
The following explanation of Go a workspace is given for historical reasons, and the fact that it’s still valid.    
You can skip this section if you want.   

The workspace directory contains the following subdirectories at its root   
1. **src**: `contains Go source files`  
   The src directory typically contains many version control repositories containing one or more Go packages.      
   Every Go source file belongs to a package. You generally create a new subdirectory inside your repository for every separate Go package.   
   
```example 
src
    github.com
        <github.com username>
        folder with code for project / repo 
        folder with code for project / repo 
        folder with code for project / repo  
        folder with code for project / repo  
        ...  
        folder with code for project / repo  
```  
2. **bin**: `contains the binary executables.`  
    The Go tool builds and installs binary executables to this directory.  
   
3. **pkg**: `contains Go package archives (.a).`   
   All the non-executable packages (shared libraries) are stored in this directory.   
   You cannot run these packages directly as they are not binary files.    
   They are typically imported and used inside other executable packages.  

1. GOPATH   
   ○ points to your go workspace  
1. GOROOT  
   ○ points to your binary installation of Go   
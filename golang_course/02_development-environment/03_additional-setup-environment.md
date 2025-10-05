### how to set up local and workspace environment  on ubuntu    
1. Note This is old method is remove after go 1.11 
1. Install go on a local machine.   
1. open home/.bashrc file.    
1. Inside bashrc first setup go path in /usr/local/go/bin.   
1. This is very important to set up the workspace.  
1. Because inside our workspace we create the application, deploy and run the application.  
1. Inside work space create three folders src, bin and pkg.     
Q1. what is scr folder?    
It is the location where source code is  available.   
Q2. what is bin folder?   
Every time we run the application binary file is created and store in bin dir.   
Q3. what is pkg folder?     
   All the non-executable packages (shared libraries) are stored in this directory.     
   You cannot run these packages directly as they are not binary files.      
   They are typically imported and used inside other executable packages.     
7. link the path of workspace with golan. It is very important.    
This the whole example to local environment and workspace  
``` 
export PATH=$PATH:/usr/local/go/bin  
export GOPATH=/home/maaz/github/LearningGO/code
export PATH=$PATH:$GOPATH/bin
```  
* This is local environment path `/usr/local/go/bin `
* This is my workspace `/home/maaz/github/LearningGO/code`  
* This is workspace path `PATH=$PATH:$GOPATH/bin`   


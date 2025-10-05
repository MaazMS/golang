### Befor build  
1. Create go workspace.    
2. setup the path of workspace.  
3. create go program.   
```
Example workspace with go program   

home
    maaz
        github  
            LearningGO
                code 
                    bin
                    pkg
                    src 
                        github.com
                            MaazMS 
                                firstapp  
                                    Main.go   



```

### How to build the Pacakege   
go build command and file structure where go source code available.     
`maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code$ go build github.com/MaazMS/firstapp`   
### How to run the build package  
* ./package name   
`maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code$ ./firstapp   `   
Hello GO    
### how to install build package  
* go install command and file structure where go source code available.  
`maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code$ go install  github.com/MaazMS/firstapp`   
### how to run install package and where is available  
* all binary code is store inside of bin folder.  
* to run the binary code.  
`maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code$ bin/firstapp `  
Hello GO 
### how to run go program in commad line 
* go run file location and go file name.go  
`maaz@maaz-Lenovo-G50-70:~/github/LearningGO$ go run code/src/github.com/MaazMS/firstapp/Main.go `
Hello GO 

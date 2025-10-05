## Environment variable
1. An environment variable is a dynamic-named value . 
1. It can affect the way running processes behave on a computer.  
1. In software development, environment variables are used to configure applications.  
1. Examples of environment variables 
   1. Include the location of all executable files in the file system, the default shell and editor, or the system locale settings.  
    
1. To work with environment variables in Go, we can use the os package or the third-party godotenv or viper libraries.  
## os.Getenv  
1. The Getenv retrieves the value of the environment variable named by the key.   
1. It returns the value, which will be empty if the variable is not present.  

## os.LookupEnv  
1. The LookupEnv function retrieves the value of the environment variable named by the key.   
1. If the variable is set the value (which may be empty) is returned, and the boolean is true.  
1. Otherwise, the returned value will be empty, and the boolean will be false. 

## os.setenv 
1. The os.Setenv sets the value of the environment variable named by the key.


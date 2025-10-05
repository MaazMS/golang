## How to Install Go in Ubuntu 20.04

1. We will be installing the latest version of Go which is 1.16.4. To download the latest version, go to the [official download   
   page](https://golang.org/dl/) and grab the tarball or use the following [wget command](https://linuxize.com/post/wget-command-examples/#:~:text=With%20Wget%2C%20you%20can%20download,a%20website%2C%20and%20much%20more.) to download it on the terminal.   
    1. https://golang.org/dl/go1.16.4.linux-amd64.tar.gz    

1. Next, extract the tarball to cd /Downloads directory.    
    1. sudo tar -xvf go1.16.4.linux-amd64.tar.gz    
    
1. simply move the new go directory to the recommended location.
    1. sudo mv go /usr/local
1. if we type go version in a terminal we get. `go’ not found`  
    1. sudo echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

1. Restart the terminal and verify the manual installation with go version

## How to uninstall Golang  
1. Find out where the Golang binary is located in your machine   
    1. which go  
```example 
/usr/local/go/bin/go

# OR

/usr/bin/go 
```        
1.  Remove the Golang binary  
    1. sudo rm -rvf /usr/local/go/ 
    
1. Remove any reference from your .bashrc file  
    1. You may have $GOPATH or $PATH export definitions in your .bash_profile or .bashrc file. Go ahead and remove those.   
## go get command
1. go get downloads the packages named by the import paths, along with their dependencies.
1. It then installs the named packages, like 'go install'.
1. When checking out or updating a package, get looks for a branch or tag that matches the locally installed version of Go.    
   The most important rule is that if the local installation is running version "go1", get searches for a branch or tag named "go1".  
   If no such version exists it retrieves the default branch of the package.   
   example `go get k8s.io/client-go@v0.20.2`

1. When go get checks out or updates a Git repository, it also updates any git submodules referenced by the repository.
1. Get never checks out or updates code stored in vendor directories.


`"go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]",`

1. go get
    1. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

1. go get -d
    1. The -d flag instructs get to stop after downloading the packages,
    1. It instructs get not to install the packages.

1. go get -f
    1. The -f flag, valid only when -u is set,
    1. The -u flag instructs get to use the network to update the named packages and their dependencies.

1. go get -t
    1. The -t flag instructs get to also download the packages required to build the tests for the specified packages.

1. go get -v
    1. The -v flag enables verbose progress and debug output.  
    
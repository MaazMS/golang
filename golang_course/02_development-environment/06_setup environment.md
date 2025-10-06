## Setting Up Go Development Environment

This guide covers setting up a Go development environment on Ubuntu, including both modern (Go modules) and legacy (GOPATH) approaches.

### Modern Approach: Go Modules (Recommended)

Since Go 1.11, Go modules are the recommended way to manage dependencies. They provide better version control, reproducible builds, and easier dependency management.

#### 1. Install Go
```bash
# Download and install Go (see installation guide)
# Add Go to PATH
echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.profile
source ~/.profile
```

#### 2. Initialize a Module
```bash
# Create a new project
mkdir my-go-app
cd my-go-app

# Initialize module
go mod init github.com/username/my-go-app
```

#### 3. Add Dependencies
```bash
# Add dependencies as needed
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
```

#### 4. Development Workflow
```bash
# Run during development
go run main.go

# Build for production
go build -o myapp main.go

# Test
go test ./...
```

### Legacy Approach: GOPATH Workspace (Historical)

**Note**: This method is deprecated since Go 1.11. Use Go modules for new projects.

#### Understanding GOPATH Workspace

The GOPATH workspace contains three main directories:

##### 1. **src** Directory
- **Purpose**: Contains Go source code
- **Structure**: Organized by import paths
- **Example**:
```
src/
    github.com/
        username/
            project1/
            project2/
    golang.org/
        x/
            tools/
```

##### 2. **bin** Directory
- **Purpose**: Contains compiled executables
- **Usage**: Add to PATH to run installed tools
- **Example**:
```
bin/
    myapp
    gin
    air
```

##### 3. **pkg** Directory
- **Purpose**: Contains compiled package objects (.a files)
- **Usage**: Intermediate binaries for faster compilation
- **Structure**:
```
pkg/
    linux_amd64/
        github.com/
            gin-gonic/
                gin.a
        golang.org/
            x/
                tools/
                    go/
                    ast/
                    astutil.a
```

#### Setting Up GOPATH Workspace

##### 1. Create Workspace Structure
```bash
# Create workspace directory
mkdir -p ~/go/{src,bin,pkg}

# Set GOPATH
export GOPATH=$HOME/go
```

##### 2. Configure Environment Variables
```bash
# Add to ~/.profile or ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

##### 3. Verify Setup
```bash
# Check Go installation
go version

# Check environment
go env GOPATH
go env GOROOT
```

### Complete Environment Setup

#### For Bash/Zsh Users
```bash
# Add to ~/.profile
echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.profile
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.profile
source ~/.profile
```

#### For Fish Shell Users
```fish
# Add to ~/.config/fish/config.fish
fish_add_path -g /usr/local/go/bin
set -Ux GOPATH $HOME/go
fish_add_path -g $GOPATH/bin
exec fish -l
```

### Development Tools Setup

#### 1. **Code Editor/IDE**
```bash
# Install VS Code with Go extension
# Or install GoLand
# Or use Vim with vim-go plugin
```

#### 2. **Essential Go Tools**
```bash
# Install development tools
go install github.com/cosmtrek/air@latest          # Hot reload
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  # Linter
go install github.com/go-delve/delve/cmd/dlv@latest  # Debugger
go install golang.org/x/tools/gopls@latest         # Language server
```

#### 3. **Project Structure (Modern)**
```
my-go-app/
├── go.mod
├── go.sum
├── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── utils/
├── cmd/
│   ├── server/
│   └── worker/
├── pkg/
│   └── client/
├── tests/
└── README.md
```

#### 4. **Project Structure (Legacy GOPATH)**
```
$GOPATH/
├── src/
│   └── github.com/
│       └── username/
│           └── my-go-app/
│               ├── main.go
│               ├── internal/
│               └── pkg/
├── bin/
│   └── my-go-app
└── pkg/
    └── linux_amd64/
        └── github.com/
            └── username/
                └── my-go-app.a
```

### Practical Examples

#### Example 1: Modern Go Module Project
```bash
# Create project
mkdir my-web-app
cd my-web-app

# Initialize module
go mod init github.com/username/my-web-app

# Add dependencies
go get github.com/gin-gonic/gin
go get github.com/spf13/viper

# Create main.go
cat > main.go << 'EOF'
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
        })
    })
    r.Run(":8080")
}
EOF

# Run application
go run main.go
```

#### Example 2: Legacy GOPATH Project
```bash
# Create project in GOPATH
mkdir -p $GOPATH/src/github.com/username/legacy-app
cd $GOPATH/src/github.com/username/legacy-app

# Create main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello from GOPATH!")
}
EOF

# Build and install
go build
go install

# Run from anywhere
legacy-app
```

### Environment Verification

#### Check Go Installation
```bash
# Verify Go version
go version

# Check environment variables
go env

# Check specific variables
go env GOPATH
go env GOROOT
go env GOBIN
```

#### Test Module Setup
```bash
# Create test module
mkdir test-module
cd test-module
go mod init example.com/test

# Add dependency
go get github.com/gin-gonic/gin

# Verify go.mod
cat go.mod

# Verify go.sum
cat go.sum
```

### Best Practices

#### 1. **Use Go Modules for New Projects**
```bash
# Always initialize modules for new projects
go mod init github.com/username/project-name
```

#### 2. **Organize Code Properly**
```
project/
├── cmd/           # Main applications
├── internal/      # Private application code
├── pkg/           # Library code
├── api/           # API definitions
├── web/           # Web assets
└── scripts/       # Build scripts
```

#### 3. **Use Development Tools**
```bash
# Install essential tools
go install github.com/cosmtrek/air@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### 4. **Version Control**
```bash
# Initialize git repository
git init
git add .
git commit -m "Initial commit"

# Add .gitignore
echo "*.exe" >> .gitignore
echo "*.exe~" >> .gitignore
echo "*.dll" >> .gitignore
echo "*.so" >> .gitignore
echo "*.dylib" >> .gitignore
echo "*.test" >> .gitignore
echo "*.out" >> .gitignore
echo "vendor/" >> .gitignore
```

### Troubleshooting

#### Common Issues

#### 1. **Go Not Found**
```bash
# Check if Go is in PATH
which go

# Add Go to PATH
export PATH=$PATH:/usr/local/go/bin
```

#### 2. **Module Issues**
```bash
# Clean module cache
go clean -modcache

# Verify modules
go mod verify

# Tidy dependencies
go mod tidy
```

#### 3. **GOPATH Issues**
```bash
# Check GOPATH
echo $GOPATH

# Set GOPATH
export GOPATH=$HOME/go
```

#### 4. **Permission Issues**
```bash
# Fix permissions
sudo chown -R $USER:$USER $GOPATH
```

### Migration from GOPATH to Modules

#### 1. **Initialize Module**
```bash
# In your project directory
go mod init github.com/username/project
```

#### 2. **Add Dependencies**
```bash
# Add existing dependencies
go mod tidy
```

#### 3. **Update Imports**
```go
// Update import paths if needed
import "github.com/username/project/internal/utils"
```

### Resources and Further Reading

- **Official Go Documentation**: https://go.dev/doc/
- **Go Modules Documentation**: https://go.dev/ref/mod
- **Go Workspace Guide**: https://go.dev/doc/code.html
- **Go Environment Variables**: https://pkg.go.dev/cmd/go#hdr-Environment_variables

### Summary

- **Modern Development**: Use Go modules for new projects
- **Legacy Support**: GOPATH still works but is deprecated
- **Environment Setup**: Configure PATH and GOPATH correctly
- **Development Tools**: Install essential Go tools for better development experience
- **Best Practices**: Follow Go conventions for project structure and dependency management
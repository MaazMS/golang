## Essential Go Commands

This guide covers the most commonly used Go commands. For detailed module management, see `08_go_module.md`.

### Getting Help

```bash
# View all available commands
go help

# Get help for specific commands
go help build
go help run
go help test

# View command usage and flags
go build -help
go run -help
```

### Core Commands

#### `go run` - Run Go Programs
```bash
# Basic usage
go run main.go
go run *.go
go run main.go arg1 arg2

# Advanced usage
go run -ldflags="-X main.version=1.0.0" main.go
go run -race main.go
```

#### `go build` - Compile Go Programs
```bash
# Basic usage
go build
go build main.go
go build -o myapp main.go

# Cross-compilation
GOOS=windows GOARCH=amd64 go build -o app.exe main.go
GOOS=darwin GOARCH=amd64 go build -o app main.go

# Optimization
go build -ldflags="-s -w" main.go
```

#### `go install` - Install Go Programs
```bash
# Install to $GOPATH/bin
go install
go install github.com/gin-gonic/gin
go install github.com/gin-gonic/gin@latest

# Install to custom location
GOBIN=/usr/local/bin go install main.go
```

### Key Differences

| Command | Purpose | Output Location | Use Case |
|---------|---------|----------------|----------|
| `go run` | Compile and run | Temporary | Development, testing |
| `go build` | Compile only | Current directory | Development, distribution |
| `go install` | Compile and install | `$GOPATH/bin` | Installing tools, production |

### Additional Commands

#### `go test` - Testing
```bash
# Run tests
go test
go test -v
go test -run TestFunction
go test -cover
go test -bench=.
```

#### `go fmt` - Code Formatting
```bash
# Format code
go fmt
go fmt ./...
go fmt -d .
```

#### `go vet` - Code Analysis
```bash
# Analyze code
go vet
go vet ./...
```

### Quick Reference

For detailed module management, see `08_go_module.md`.  
For package management, see `07_Package_management.md`.  
For environment setup, see `02_additional-setup-environment.md`.

## The `go get` Command

The `go get` command downloads and installs packages and their dependencies. In Go modules (Go 1.11+), it also updates your `go.mod` and `go.sum` files.

### Basic Usage

```bash
# Download and install a package
go get github.com/gin-gonic/gin

# Install a specific version
go get github.com/gin-gonic/gin@v1.9.1

# Install the latest version
go get github.com/gin-gonic/gin@latest

# Install from a specific branch
go get github.com/gin-gonic/gin@master

# Install from a specific commit
go get github.com/gin-gonic/gin@abc1234
```

### Command Syntax

```bash
go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]
```

### Flags Explained

#### `-d` (Download only)
- Downloads packages but doesn't install them
- Useful for pre-downloading dependencies
- Doesn't update `go.mod` or `go.sum`

```bash
# Download without installing
go get -d github.com/gin-gonic/gin
```

#### `-u` (Update)
- Updates packages and their dependencies to latest versions
- Updates `go.mod` and `go.sum` files
- Can be combined with other flags

```bash
# Update a specific package
go get -u github.com/gin-gonic/gin

# Update all dependencies
go get -u ./...
```

#### `-f` (Force)
- Only valid when used with `-u`
- Forces updates even if there are conflicts
- Use with caution as it may break compatibility

```bash
# Force update despite conflicts
go get -u -f github.com/gin-gonic/gin
```

#### `-t` (Test dependencies)
- Also downloads packages required for testing
- Includes test dependencies in the download

```bash
# Include test dependencies
go get -t github.com/gin-gonic/gin
```

#### `-v` (Verbose)
- Enables verbose output
- Shows detailed progress information

```bash
# Verbose output
go get -v github.com/gin-gonic/gin
```

#### `-fix` (Fix)
- Runs `go fix` on downloaded packages
- Automatically fixes code issues when possible

```bash
# Download and fix issues
go get -fix github.com/gin-gonic/gin
```

#### `-insecure` (Insecure)
- Allows downloads from insecure (HTTP) sources
- Not recommended for production use

```bash
# Allow insecure downloads
go get -insecure example.com/package
```

### Practical Examples

#### Installing Popular Go Packages

```bash
# Web framework
go get github.com/gin-gonic/gin

# Database ORM
go get gorm.io/gorm
go get gorm.io/driver/mysql

# HTTP client
go get github.com/go-resty/resty/v2

# Configuration management
go get github.com/spf13/viper

# Logging
go get go.uber.org/zap

# Testing
go get github.com/stretchr/testify
```

#### Working with Versions

```bash
# Install specific version
go get github.com/gin-gonic/gin@v1.9.1

# Install latest patch version of v1.9
go get github.com/gin-gonic/gin@v1.9

# Install latest minor version of v1
go get github.com/gin-gonic/gin@v1

# Install latest version
go get github.com/gin-gonic/gin@latest
```

#### Updating Dependencies

```bash
# Update a specific package
go get -u github.com/gin-gonic/gin

# Update all dependencies in current module
go get -u ./...

# Update to latest patch version
go get -u=patch ./...

# Update to latest minor version
go get -u=minor ./...

# Update to latest major version
go get -u=major ./...
```

#### Working with Private Repositories

```bash
# Set GOPRIVATE for private modules
go env -w GOPRIVATE=github.com/yourcompany/*

# Install from private repository
go get github.com/yourcompany/private-package
```

### Module vs GOPATH Mode

#### In Module Mode (Go 1.11+)
- `go get` updates `go.mod` and `go.sum`
- Packages are downloaded to module cache
- Dependencies are managed per module

#### In GOPATH Mode (Legacy)
- `go get` downloads to `$GOPATH/src`
- Installs binaries to `$GOPATH/bin`
- Global dependency management

### Common Use Cases

#### 1. Adding a New Dependency
```bash
# Add a new package to your module
go get github.com/gin-gonic/gin
```

#### 2. Updating Dependencies
```bash
# Update all dependencies
go get -u ./...

# Update specific dependency
go get -u github.com/gin-gonic/gin
```

#### 3. Downgrading a Package
```bash
# Install older version
go get github.com/gin-gonic/gin@v1.8.2
```

#### 4. Removing a Dependency
```bash
# Remove from go.mod (Go 1.17+)
go mod tidy

# Or manually edit go.mod and run
go mod download
```

### Troubleshooting

#### Network Issues
```bash
# Use different proxy
go env -w GOPROXY=https://goproxy.cn,direct

# Disable proxy
go env -w GOPROXY=direct
```

#### Version Conflicts
```bash
# Check why a version was chosen
go mod why github.com/gin-gonic/gin

# View module graph
go mod graph
```

#### Clean Module Cache
```bash
# Clean module cache
go clean -modcache
```

### Best Practices

1. **Use specific versions** in production:
   ```bash
   go get github.com/gin-gonic/gin@v1.9.1
   ```

2. **Update dependencies regularly**:
   ```bash
   go get -u ./...
   ```

3. **Use `go mod tidy`** after changes:
   ```bash
   go mod tidy
   ```

4. **Test after updates**:
   ```bash
   go test ./...
   ```

5. **Use `go get -u=patch`** for safer updates:
   ```bash
   go get -u=patch ./...
   ``` 
## Go Modules: Modern Dependency Management

Go modules (introduced in Go 1.11, default since Go 1.16) provide a modern approach to dependency management. They replace the legacy GOPATH workspace system and offer better version control, reproducible builds, and easier dependency management.

### What are Go Modules?

Go modules are collections of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module's path, Go version, and dependencies.

#### Key Benefits
- **Version Control**: Lock specific versions of dependencies
- **Reproducible Builds**: Same dependencies across different environments
- **Semantic Versioning**: Follow semantic versioning principles
- **Proxy Support**: Use Go proxy for faster downloads
- **Checksum Verification**: Ensure package integrity with `go.sum`

### Creating a New Module

#### 1. Initialize a Module
```bash
# Create a new directory
mkdir my-go-app
cd my-go-app

# Initialize module
go mod init github.com/username/my-go-app
```

This creates a `go.mod` file:
```go
module github.com/username/my-go-app

go 1.21
```

#### 2. Create Source Code
```go
// hello.go
package hello

func Hello() string {
    return "Hello, world."
}
```

#### 3. Write Tests
```go
// hello_test.go
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```

#### 4. Run Tests
```bash
go test
# PASS
# ok      github.com/username/my-go-app    0.020s
```

### Understanding go.mod and go.sum

#### go.mod File Structure
```go
module github.com/username/my-go-app

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/spf13/viper v1.16.0
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
)
```

#### go.sum File
```
github.com/gin-gonic/gin v1.9.1 h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=
github.com/gin-gonic/gin v1.9.1/go.mod h1:hPrL7HqKDbLK9t9ngPU4odXOIa7i4eBdok5Ej6/YJ4M=
```

### Adding Dependencies

#### 1. Add a Dependency
```go
// hello.go
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```

#### 2. Download and Test
```bash
go test
# go: finding module for package rsc.io/quote
# go: downloading rsc.io/quote v1.5.2
# go: found rsc.io/quote in rsc.io/quote v1.5.2
# go: downloading rsc.io/sampler v1.3.0
# go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
# PASS
# ok      github.com/username/my-go-app    0.006s
```

#### 3. View Dependencies
```bash
go list -m all
# github.com/username/my-go-app
# golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
# rsc.io/quote v1.5.2
# rsc.io/sampler v1.3.0
```

### Understanding Dependencies

#### Direct Dependencies
- **Direct dependencies** are packages you import directly in your code
- They appear in the `require` block of `go.mod`
- Example: `rsc.io/quote` is a direct dependency

#### Indirect Dependencies
- **Indirect dependencies** are dependencies of your direct dependencies
- They appear with `// indirect` comment in `go.mod`
- Example: `golang.org/x/text` is an indirect dependency of `rsc.io/quote`

#### go.sum File
- Contains cryptographic hashes of module versions
- Ensures integrity and reproducibility
- Automatically managed by Go

### Module vs GOPATH Mode

| Feature | Module Mode | GOPATH Mode |
|---------|-------------|-------------|
| **Dependency Management** | Per-module | Global |
| **Version Control** | go.mod/go.sum | No versioning |
| **Reproducible Builds** | Yes | No |
| **Proxy Support** | Yes | No |
| **Checksum Verification** | Yes | No |

### Migration from GOPATH

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

- **Official Go Modules Documentation**: https://go.dev/ref/mod
- **Go Package Discovery**: https://pkg.go.dev/
- **Go Security Advisories**: https://pkg.go.dev/vuln
- **Semantic Versioning**: https://semver.org/
- **Go Modules Research**: https://research.swtch.com/deps

### Complete Example: REST API

```go
// main.go
package main

import (
    "log"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

type User struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

func main() {
    // Load configuration
    viper.SetConfigFile("config.yaml")
    viper.ReadInConfig()
    
    // Setup logging
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    
    // Setup database
    db, err := gorm.Open(mysql.Open(viper.GetString("database.url")), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
    
    // Auto-migrate
    db.AutoMigrate(&User{})
    
    // Setup web server
    r := gin.Default()
    
    // Routes
    r.GET("/users", func(c *gin.Context) {
        var users []User
        db.Find(&users)
        c.JSON(http.StatusOK, users)
    })
    
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Create(&user)
        c.JSON(http.StatusCreated, user)
    })
    
    logger.Info("Starting server", zap.String("port", viper.GetString("server.port")))
    r.Run(":" + viper.GetString("server.port"))
}
```

```bash
# Initialize module
go mod init github.com/username/rest-api

# Add dependencies
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
go get go.uber.org/zap
go get gorm.io/gorm
go get gorm.io/driver/mysql

# Run application
go run main.go
```

This example demonstrates how Go modules enable you to build sophisticated applications by managing dependencies effectively and ensuring reproducible builds. 

## Go Environment Variables

Environment variables are dynamic-named values that can affect how running processes behave on a computer. In software development, they are used to configure applications without changing code.

### Common System Environment Variables
Examples include:
- `PATH`: Location of executable files in the file system
- `HOME`: User's home directory
- `SHELL`: Default shell
- `EDITOR`: Default text editor
- `LANG`: System locale settings

### Go-Specific Environment Variables

Run `go env` to see all Go environment variables and their current values:

```bash
$ go env
```

#### Key Go Environment Variables Explained

**Core Go Variables:**
- `GOROOT`: Path to Go installation (e.g., `/usr/local/go`)
- `GOPATH`: Workspace directory for Go code (defaults to `$HOME/go`)
- `GOBIN`: Directory for installed binaries (defaults to `$GOPATH/bin`)

**Module System:**
- `GO111MODULE`: Controls module mode (`on`, `off`, `auto`)
- `GOMOD`: Path to the go.mod file (empty if not in a module)
- `GOMODCACHE`: Directory for module cache (defaults to `$GOPATH/pkg/mod`)

**Build Configuration:**
- `GOOS`: Target operating system (`linux`, `windows`, `darwin`, etc.)
- `GOARCH`: Target architecture (`amd64`, `arm64`, `386`, etc.)
- `CGO_ENABLED`: Whether CGO is enabled (`1` or `0`)

**Proxy and Security:**
- `GOPROXY`: Module proxy URLs (default: `https://proxy.golang.org,direct`)
- `GOSUMDB`: Checksum database (default: `sum.golang.org`)
- `GOPRIVATE`: Private modules that bypass proxy and checksum database

**Development Tools:**
- `GOCACHE`: Build cache directory (defaults to `$HOME/.cache/go-build`)
- `GOENV`: Path to Go environment configuration file
- `GOTOOLDIR`: Directory containing Go tools

### Setting Go Environment Variables

**Temporary (current session only):**
```bash
export GOPATH="/path/to/workspace"
export GOOS="windows"
export GOARCH="amd64"
```

**Permanent (add to shell profile):**
```bash
# For Bash/Zsh (~/.profile or ~/.bashrc)
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.profile

# For Fish (~/.config/fish/config.fish)
set -Ux GOPATH $HOME/go
fish_add_path -g $GOPATH/bin
```



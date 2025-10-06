## GOPATH Workspace: Legacy Reference

**Note**: This file provides historical context for GOPATH workspaces. For modern development, see `02_additional-setup-environment.md` and `08_go_module.md`.

### Historical Context

Before Go modules (Go 1.11+), Go code lived inside a single workspace pointed to by the `GOPATH` environment variable. Since Go modules became the default (Go 1.16), you no longer need to keep projects under `GOPATH`.

### GOPATH Workspace Structure

The workspace directory contains these subdirectories:

#### 1. **src** Directory
- Contains Go source code organized by import paths
- Structure: `src/github.com/username/project/`
- Each repository contains one or more Go packages

#### 2. **bin** Directory  
- Contains compiled executables
- `go install` places binaries here
- Add `$GOPATH/bin` to PATH to run installed tools

#### 3. **pkg** Directory
- Contains compiled package archives (`.a` files)
- Used by compiler and linker for faster builds
- Not directly executable; imported by other packages

### Key Environment Variables

- **GOPATH**: Points to Go workspace (defaults to `$HOME/go`)
- **GOROOT**: Points to Go installation (e.g., `/usr/local/go`)

### Modern Recommendation

For new projects, use Go modules:
```bash
go mod init example.com/your/module
```

This allows projects to live anywhere on your filesystem with better dependency management.
# GOPATH Documentation

## What is $GOPATH?

GOPATH is an environment variable that points to a physical directory on your computer where Go code is stored.

## What is a workspace?

GOPATH is also called a workspace, which means it's a place where all Go source code is stored and managed.

## What are the folders inside $GOPATH?

The GOPATH directory contains three main folders:

1. **bin** - Contains executable binaries
2. **pkg** - Contains package objects  
3. **src** - Contains source code files

## Where should you save your Go source code?

- Under `$GOPATH/src`

## Do you need to set $GOPATH?

- **No**: It's automatically set to your user directory (e.g., `$HOME/go` on Unix systems)

## How can you print $GOPATH?

- Using the `go env GOPATH` command

## Important Note: GOPATH vs Go Modules

**⚠️ Legacy Information**: GOPATH was the standard way to manage Go projects before Go 1.11. 

**Modern Approach**: Since Go 1.11, **Go modules** are the recommended way to manage dependencies and projects. Go modules use a `go.mod` file instead of relying on GOPATH.

### Key Differences:
- **GOPATH**: All code must be under `$GOPATH/src`
- **Go Modules**: Code can be anywhere on your system
- **GOPATH**: Global workspace for all projects
- **Go Modules**: Each project has its own module with `go.mod`

### When to use what:
- **Use Go modules** for new projects (recommended)
- **GOPATH** is still supported for legacy projects but not recommended for new development

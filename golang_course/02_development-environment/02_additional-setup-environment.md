### Go Environment Configuration

This file covers environment configuration after Go installation. For installation instructions, see `01_installing_uninstall_go.md`.

#### Modern Approach: Go Modules (Recommended)

Go modules (default since Go 1.16) provide better dependency management than the legacy GOPATH workspace.

#### Environment Variables Setup

**Bash/Zsh (`~/.profile` is recommended for login shells):**
```bash
echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.profile
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.profile
source ~/.profile
```

**Fish shell (`~/.config/fish/config.fish`):**
```fish
fish_add_path -g /usr/local/go/bin
set -Ux GOPATH $HOME/go
fish_add_path -g $GOPATH/bin
exec fish -l
```

**Verify setup:**
```bash
go version
go env GOROOT GOPATH GOBIN
```

#### Legacy GOPATH Workspace (Historical Reference)

Before modules, Go projects lived under `$GOPATH/src`, with build outputs in `$GOPATH/bin` and compiled package caches in `$GOPATH/pkg`.

- **src**: Source code organized by import paths
- **bin**: Compiled executables (add `$GOPATH/bin` to PATH)
- **pkg**: Compiled package objects (.a files) for faster compilation

**Note**: New projects should use Go modules with `go mod init` in your project directory.

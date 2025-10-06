## Install Go on Ubuntu (20.04+)

The steps below install Go from the official tarball. They work on Ubuntu 20.04 and newer.

### 1) Remove any previous Go installation (if present)
- This ensures a clean install.
```bash
sudo rm -rf /usr/local/go
```

### 2) Download the latest Go release
- Visit the official downloads page and copy the URL for the latest Linux `amd64` or `arm64` tarball: `https://go.dev/dl/`.
- Example (replace the version with the latest you see on the page):
```bash
curl -LO https://go.dev/dl/go1.22.6.linux-amd64.tar.gz
```

### 3) (Optional but recommended) Verify the checksum
- On the downloads page, copy the corresponding SHA256 checksum and compare it locally:
```bash
sha256sum go1.22.6.linux-amd64.tar.gz
# Compare the printed hash with the one on https://go.dev/dl/
```

### 4) Install to `/usr/local`
```bash
sudo tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz
```

### 5) Add Go to your PATH
- Bash/Zsh (recommended to use `~/.profile` so it applies to login shells):
```bash
echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.profile
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.profile
source ~/.profile
```

- Fish shell (`~/.config/fish/config.fish`):
```fish
fish_add_path -g /usr/local/go/bin
set -Ux GOPATH $HOME/go
fish_add_path -g $GOPATH/bin
exec fish -l
```

### 6) Verify your installation
```bash
go version
go env GOROOT GOPATH GOBIN
```
You should see `GOROOT` as `/usr/local/go` and `GOPATH` as `$HOME/go`.

---

## Uninstall Go

### 1) Locate the `go` binary (optional)
```bash
which go
``` 
Examples:
```text
/usr/local/go/bin/go
# or
/usr/bin/go
```

### 2) Remove the Go installation directory
```bash
sudo rm -rf /usr/local/go
```

### 3) Remove PATH entries and environment variables
- Bash/Zsh: remove any `export PATH=.../usr/local/go/bin`, `GOPATH`, and `$GOPATH/bin` lines from `~/.profile`, `~/.bashrc`, or `~/.zshrc`, then reload your shell.
- Fish: remove the lines added to `~/.config/fish/config.fish`, or run:
```fish
set -e GOPATH
```

### 4) (Optional) Remove your workspace
- If you no longer need your Go workspace:
```bash
rm -rf "$HOME/go"
```


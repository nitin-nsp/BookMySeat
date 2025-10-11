# Development Environment Setup Guide

## Quick Start

```bash
# Make scripts executable
chmod +x setup-dev-env.sh verify-setup.sh

# Run setup
./setup-dev-env.sh

# Verify installation
./verify-setup.sh
```

## Manual Installation (if needed)

### 1. Install Go 1.21+

**macOS:**
```bash
brew install go
```

**Linux:**
```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

Verify:
```bash
go version  # Should show go1.21 or higher
```

### 2. Install Docker & Docker Compose

**macOS:**
```bash
brew install --cask docker
```

**Linux:**
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
```

Start Docker Desktop and verify:
```bash
docker --version
docker compose version
docker info
```

### 3. Install kubectl

**macOS:**
```bash
brew install kubectl
```

**Linux:**
```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

Verify:
```bash
kubectl version --client
```

### 4. Install minikube or kind

**minikube (recommended):**
```bash
# macOS
brew install minikube

# Linux
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

**kind (alternative):**
```bash
# macOS
brew install kind

# Linux
go install sigs.k8s.io/kind@latest
```

Verify:
```bash
minikube version
# or
kind version
```

### 5. Install goctl (go-zero CLI)

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

Verify:
```bash
goctl --version
```

### 6. Install protoc and plugins

**Install protoc:**
```bash
# macOS
brew install protobuf

# Linux
sudo apt install -y protobuf-compiler
```

**Install Go plugins:**
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Verify:
```bash
protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version
```

### 7. Setup PATH

Add Go binaries to PATH (add to ~/.zshrc or ~/.bashrc):
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Apply changes:
```bash
source ~/.zshrc  # or source ~/.bashrc
```

## IDE Setup

### VSCode

Install extensions:
```bash
code --install-extension golang.go
code --install-extension ms-azuretools.vscode-docker
code --install-extension ms-kubernetes-tools.vscode-kubernetes-tools
```

Or manually install:
- Go (golang.go)
- Docker (ms-azuretools.vscode-docker)
- Kubernetes (ms-kubernetes-tools.vscode-kubernetes-tools)
- Protobuf (zxh404.vscode-proto3)

### GoLand

GoLand comes with built-in support for Go, Docker, and Kubernetes.

Enable plugins:
- Go
- Docker
- Kubernetes
- Protocol Buffers

## Verification Checklist

Run the verification script:
```bash
./verify-setup.sh
```

Or manually verify:

- [ ] `go version` shows 1.21+
- [ ] `docker --version` works
- [ ] `docker compose version` works
- [ ] `docker info` shows running daemon
- [ ] `kubectl version --client` works
- [ ] `minikube version` or `kind version` works
- [ ] `goctl --version` works
- [ ] `protoc --version` works
- [ ] `protoc-gen-go --version` works
- [ ] `protoc-gen-go-grpc --version` works
- [ ] Git repository initialized
- [ ] .gitignore file exists

## Troubleshooting

### Go tools not found
```bash
# Check GOPATH
go env GOPATH

# Add to PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### Docker daemon not running
```bash
# macOS: Start Docker Desktop application
# Linux: 
sudo systemctl start docker
```

### Permission denied for Docker
```bash
# Linux: Add user to docker group
sudo usermod -aG docker $USER
newgrp docker
```

### protoc plugins not found
```bash
# Reinstall plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Verify GOPATH/bin is in PATH
echo $PATH | grep $(go env GOPATH)/bin
```

## Next Steps

After completing the setup:

1. ✅ Mark items in `complete.check_list.md` under section 1.1
2. 📁 Proceed to section 1.2: Project Structure Creation
3. 🐳 Then section 1.3: Infrastructure Setup (Docker Compose)

## Useful Commands

```bash
# Check Go environment
go env

# List installed Go tools
ls $(go env GOPATH)/bin

# Docker system info
docker system info
docker system df

# Kubernetes cluster info
kubectl cluster-info
minikube status
```

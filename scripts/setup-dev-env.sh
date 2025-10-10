#!/bin/bash

set -e

echo "=========================================="
echo "Development Environment Setup"
echo "=========================================="

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Check Go installation
echo -e "\n${YELLOW}[1/7] Checking Go installation...${NC}"
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo -e "${GREEN}✓ Go is installed: $GO_VERSION${NC}"
    
    # Check if version is 1.21+
    MAJOR=$(echo $GO_VERSION | sed 's/go//' | cut -d. -f1)
    MINOR=$(echo $GO_VERSION | sed 's/go//' | cut -d. -f2)
    if [ "$MAJOR" -ge 1 ] && [ "$MINOR" -ge 21 ]; then
        echo -e "${GREEN}✓ Go version is 1.21 or higher${NC}"
    else
        echo -e "${RED}✗ Go version should be 1.21+. Please upgrade.${NC}"
        echo "  Download from: https://go.dev/dl/"
    fi
else
    echo -e "${RED}✗ Go is not installed${NC}"
    echo "  Install: brew install go (macOS) or visit https://go.dev/dl/"
    exit 1
fi

# Check Docker
echo -e "\n${YELLOW}[2/7] Checking Docker installation...${NC}"
if command -v docker &> /dev/null; then
    DOCKER_VERSION=$(docker --version)
    echo -e "${GREEN}✓ Docker is installed: $DOCKER_VERSION${NC}"
    
    # Check if Docker daemon is running
    if docker info &> /dev/null; then
        echo -e "${GREEN}✓ Docker daemon is running${NC}"
    else
        echo -e "${RED}✗ Docker daemon is not running. Please start Docker Desktop.${NC}"
    fi
else
    echo -e "${RED}✗ Docker is not installed${NC}"
    echo "  Install: brew install --cask docker (macOS) or visit https://www.docker.com/products/docker-desktop"
    exit 1
fi

# Check Docker Compose
echo -e "\n${YELLOW}[3/7] Checking Docker Compose...${NC}"
if docker compose version &> /dev/null; then
    COMPOSE_VERSION=$(docker compose version)
    echo -e "${GREEN}✓ Docker Compose is installed: $COMPOSE_VERSION${NC}"
else
    echo -e "${RED}✗ Docker Compose is not installed${NC}"
    echo "  It should come with Docker Desktop. Please reinstall Docker."
    exit 1
fi

# Check kubectl
echo -e "\n${YELLOW}[4/7] Checking kubectl installation...${NC}"
if command -v kubectl &> /dev/null; then
    KUBECTL_VERSION=$(kubectl version --client --short 2>/dev/null || kubectl version --client)
    echo -e "${GREEN}✓ kubectl is installed${NC}"
else
    echo -e "${YELLOW}✗ kubectl is not installed${NC}"
    echo "  Installing kubectl..."
    brew install kubectl
fi

# Check minikube or kind
echo -e "\n${YELLOW}[5/7] Checking local Kubernetes (minikube/kind)...${NC}"
if command -v minikube &> /dev/null; then
    echo -e "${GREEN}✓ minikube is installed${NC}"
elif command -v kind &> /dev/null; then
    echo -e "${GREEN}✓ kind is installed${NC}"
else
    echo -e "${YELLOW}✗ Neither minikube nor kind is installed${NC}"
    echo "  Installing minikube..."
    brew install minikube
fi

# Install goctl
echo -e "\n${YELLOW}[6/7] Installing goctl (go-zero CLI)...${NC}"
if command -v goctl &> /dev/null; then
    echo -e "${GREEN}✓ goctl is already installed${NC}"
else
    echo "  Installing goctl..."
    go install github.com/zeromicro/go-zero/tools/goctl@latest
    echo -e "${GREEN}✓ goctl installed${NC}"
fi

# Install protoc and plugins
echo -e "\n${YELLOW}[7/7] Checking protoc and plugins...${NC}"
if command -v protoc &> /dev/null; then
    PROTOC_VERSION=$(protoc --version)
    echo -e "${GREEN}✓ protoc is installed: $PROTOC_VERSION${NC}"
else
    echo -e "${YELLOW}✗ protoc is not installed${NC}"
    echo "  Installing protoc..."
    brew install protobuf
fi

# Install protoc-gen-go
echo "  Checking protoc-gen-go..."
if command -v protoc-gen-go &> /dev/null; then
    echo -e "${GREEN}✓ protoc-gen-go is installed${NC}"
else
    echo "  Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

# Install protoc-gen-go-grpc
echo "  Checking protoc-gen-go-grpc..."
if command -v protoc-gen-go-grpc &> /dev/null; then
    echo -e "${GREEN}✓ protoc-gen-go-grpc is installed${NC}"
else
    echo "  Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Setup Git repository
echo -e "\n${YELLOW}Setting up Git repository...${NC}"
if [ -d ".git" ]; then
    echo -e "${GREEN}✓ Git repository already initialized${NC}"
else
    echo "  Initializing Git repository..."
    git init
    echo -e "${GREEN}✓ Git repository initialized${NC}"
fi

# Create .gitignore
if [ ! -f ".gitignore" ]; then
    echo "  Creating .gitignore..."
    cat > .gitignore << 'EOF'
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test binary
*.test
*.out

# Go workspace file
go.work

# Dependencies
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# Environment variables
.env
.env.local

# Logs
*.log
logs/

# OS
.DS_Store
Thumbs.db

# Docker
docker-compose.override.yml

# Kubernetes
*.kubeconfig

# Temporary files
tmp/
temp/
EOF
    echo -e "${GREEN}✓ .gitignore created${NC}"
else
    echo -e "${GREEN}✓ .gitignore already exists${NC}"
fi

# Verify GOPATH/bin is in PATH
echo -e "\n${YELLOW}Verifying Go tools in PATH...${NC}"
GOPATH=$(go env GOPATH)
if [[ ":$PATH:" == *":$GOPATH/bin:"* ]]; then
    echo -e "${GREEN}✓ \$GOPATH/bin is in PATH${NC}"
else
    echo -e "${YELLOW}⚠ \$GOPATH/bin is not in PATH${NC}"
    echo "  Add this to your ~/.zshrc or ~/.bashrc:"
    echo "  export PATH=\$PATH:\$(go env GOPATH)/bin"
fi

echo -e "\n=========================================="
echo -e "${GREEN}Development Environment Setup Complete!${NC}"
echo -e "==========================================\n"

echo "Next steps:"
echo "  1. Ensure Docker Desktop is running"
echo "  2. Run: source ~/.zshrc (or ~/.bashrc) to refresh PATH"
echo "  3. Verify installations: ./verify-setup.sh"
echo ""

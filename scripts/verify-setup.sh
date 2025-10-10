#!/bin/bash

echo "=========================================="
echo "Verifying Development Environment"
echo "=========================================="

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

PASS=0
FAIL=0

check_command() {
    if command -v $1 &> /dev/null; then
        echo -e "${GREEN}✓${NC} $1"
        ((PASS++))
        return 0
    else
        echo -e "${RED}✗${NC} $1"
        ((FAIL++))
        return 1
    fi
}

echo -e "\nChecking required tools:\n"

check_command go
check_command docker
check_command kubectl
check_command goctl
check_command protoc
check_command protoc-gen-go
check_command protoc-gen-go-grpc

echo ""
if command -v minikube &> /dev/null; then
    echo -e "${GREEN}✓${NC} minikube"
    ((PASS++))
elif command -v kind &> /dev/null; then
    echo -e "${GREEN}✓${NC} kind"
    ((PASS++))
else
    echo -e "${RED}✗${NC} minikube/kind"
    ((FAIL++))
fi

echo -e "\n=========================================="
echo -e "Results: ${GREEN}$PASS passed${NC}, ${RED}$FAIL failed${NC}"
echo -e "==========================================\n"

if [ $FAIL -eq 0 ]; then
    echo -e "${GREEN}All checks passed! Ready to proceed.${NC}\n"
    exit 0
else
    echo -e "${RED}Some checks failed. Please run ./setup-dev-env.sh${NC}\n"
    exit 1
fi

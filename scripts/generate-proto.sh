#!/bin/bash

set -e

echo "Generating protobuf code..."

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

# Generate for each service
services=("user" "event" "booking" "payment" "notification")

for service in "${services[@]}"; do
    echo -e "${BLUE}Generating ${service} service...${NC}"
    
    protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           proto/${service}/${service}.proto
    
    echo -e "${GREEN}✓ ${service} service generated${NC}"
done

echo -e "${GREEN}All protobuf code generated successfully!${NC}"

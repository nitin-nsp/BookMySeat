#!/bin/bash

set -e

echo "=========================================="
echo "Infrastructure Verification Script"
echo "=========================================="
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

check_service() {
    local service=$1
    local check_cmd=$2
    
    echo -n "Checking $service... "
    if eval "$check_cmd" > /dev/null 2>&1; then
        echo -e "${GREEN}✓ OK${NC}"
        return 0
    else
        echo -e "${RED}✗ FAILED${NC}"
        return 1
    fi
}

echo "1. Checking Docker containers status..."
echo "----------------------------------------"
docker-compose ps

echo ""
echo "2. Testing PostgreSQL..."
echo "----------------------------------------"
check_service "PostgreSQL Connection" "docker exec postgres pg_isready -U postgres"
check_service "user_db database" "docker exec postgres psql -U postgres -lqt | cut -d \| -f 1 | grep -qw user_db"
check_service "event_db database" "docker exec postgres psql -U postgres -lqt | cut -d \| -f 1 | grep -qw event_db"
check_service "payment_db database" "docker exec postgres psql -U postgres -lqt | cut -d \| -f 1 | grep -qw payment_db"
check_service "dtm database" "docker exec postgres psql -U postgres -lqt | cut -d \| -f 1 | grep -qw dtm"

echo ""
echo "3. Testing MongoDB..."
echo "----------------------------------------"
check_service "MongoDB Connection" "docker exec mongodb mongosh --quiet --eval 'db.adminCommand({ping: 1})'"

echo ""
echo "4. Testing Redis Cluster..."
echo "----------------------------------------"
check_service "Redis Node 1 (6379)" "docker exec redis-1 redis-cli ping"
check_service "Redis Node 2 (6380)" "docker exec redis-2 redis-cli -p 6380 ping"
check_service "Redis Node 3 (6381)" "docker exec redis-3 redis-cli -p 6381 ping"

echo ""
echo "5. Testing Kafka..."
echo "----------------------------------------"
check_service "Zookeeper" "docker exec zookeeper nc -z localhost 2181"
check_service "Kafka Broker" "docker exec kafka kafka-broker-api-versions --bootstrap-server localhost:9092"

echo ""
echo "6. Testing Elasticsearch..."
echo "----------------------------------------"
check_service "Elasticsearch" "curl -s http://localhost:9200/_cluster/health"

echo ""
echo "7. Testing Kibana..."
echo "----------------------------------------"
check_service "Kibana" "curl -s http://localhost:5601/api/status"

echo ""
echo "8. Testing Jaeger..."
echo "----------------------------------------"
check_service "Jaeger UI" "curl -s http://localhost:16686"

echo ""
echo "9. Testing Prometheus..."
echo "----------------------------------------"
check_service "Prometheus" "curl -s http://localhost:9090/-/healthy"

echo ""
echo "10. Testing Grafana..."
echo "----------------------------------------"
check_service "Grafana" "curl -s http://localhost:3000/api/health"

echo ""
echo "11. Testing DTM..."
echo "----------------------------------------"
check_service "DTM Server" "curl -s http://localhost:36789/api/dtmsvr/health"

echo ""
echo "=========================================="
echo "Verification Complete!"
echo "=========================================="
echo ""
echo "Service URLs:"
echo "  PostgreSQL:     localhost:5432"
echo "  MongoDB:        localhost:27017"
echo "  Redis-1:        localhost:6379"
echo "  Redis-2:        localhost:6380"
echo "  Redis-3:        localhost:6381"
echo "  Kafka:          localhost:9092"
echo "  Elasticsearch:  http://localhost:9200"
echo "  Kibana:         http://localhost:5601"
echo "  Jaeger:         http://localhost:16686"
echo "  Prometheus:     http://localhost:9090"
echo "  Grafana:        http://localhost:3000 (admin/admin)"
echo "  DTM:            http://localhost:36789"
echo ""

# Quick Start Guide

## Prerequisites

Ensure you have completed the development environment setup:
- Go 1.21+
- Docker & Docker Compose
- kubectl & minikube
- goctl, protoc, and plugins

## Step 1: Initialize Project

```bash
make init
```

This creates the `.env` file and log directories.

## Step 2: Start Infrastructure

```bash
make infra-up
```

Wait for all services to start (approximately 30 seconds).

## Step 3: Verify Infrastructure

```bash
make verify
```

Or run the verification script directly:
```bash
./scripts/verify-infrastructure.sh
```

## Step 4: Check Service Status

```bash
make infra-status
```

## Step 5: Access Monitoring Dashboards

### Grafana (Metrics Visualization)
```bash
make open-grafana
```
Or visit: http://localhost:3000
- Username: `admin`
- Password: `admin`

### Prometheus (Metrics Collection)
```bash
make open-prometheus
```
Or visit: http://localhost:9090

### Jaeger (Distributed Tracing)
```bash
make open-jaeger
```
Or visit: http://localhost:16686

### Kibana (Log Visualization)
```bash
make open-kibana
```
Or visit: http://localhost:5601

## Database Access

### PostgreSQL
```bash
make db-psql
```

Or connect manually:
```bash
psql -h localhost -p 5432 -U postgres
```

Databases:
- `user_db`
- `event_db`
- `payment_db`
- `dtm`

### MongoDB
```bash
make db-mongo
```

Or connect manually:
```bash
mongosh mongodb://admin:admin@localhost:27017
```

### Redis
```bash
make db-redis
```

Or connect to specific nodes:
```bash
redis-cli -h localhost -p 6379  # Node 1
redis-cli -h localhost -p 6380  # Node 2
redis-cli -h localhost -p 6381  # Node 3
```

## Useful Commands

### View Logs
```bash
make infra-logs
```

### Restart Infrastructure
```bash
make infra-restart
```

### Stop Infrastructure
```bash
make infra-down
```

### Clean All Data
```bash
make clean
```

## Troubleshooting

### Services not starting
```bash
# Check Docker
docker ps

# View specific service logs
docker logs postgres
docker logs mongodb
docker logs redis-1

# Restart
make infra-restart
```

### Port already in use
Edit `docker-compose.yml` and change the conflicting port mappings.

### Database connection issues
```bash
# Test PostgreSQL
docker exec postgres pg_isready -U postgres

# Test MongoDB
docker exec mongodb mongosh --eval "db.adminCommand('ping')"

# Test Redis
docker exec redis-1 redis-cli ping
```

## Next Steps

Once infrastructure is running:
1. Proceed to Phase 2: Proto Definitions
2. Start implementing User Service (Phase 3)
3. Build remaining microservices

## Help

View all available commands:
```bash
make help
```

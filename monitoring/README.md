# Monitoring Stack

This directory contains configuration for the complete observability stack.

## Components

### 1. Prometheus (Metrics)
- **URL**: http://localhost:9090
- **Purpose**: Metrics collection and storage
- **Config**: `prometheus.yml`
- **Scrape Interval**: 15s
- **Retention**: 30 days

### 2. Grafana (Visualization)
- **URL**: http://localhost:3000
- **Credentials**: admin/admin
- **Purpose**: Metrics visualization and dashboards
- **Datasource**: Prometheus (auto-configured)

### 3. Jaeger (Distributed Tracing)
- **URL**: http://localhost:16686
- **Purpose**: Distributed tracing across microservices
- **Ports**:
  - 6831: Jaeger agent (UDP)
  - 16686: Jaeger UI
  - 14268: Collector HTTP

### 4. ELK Stack (Logging)

#### Elasticsearch
- **URL**: http://localhost:9200
- **Purpose**: Log storage and search

#### Kibana
- **URL**: http://localhost:5601
- **Purpose**: Log visualization and analysis

#### Filebeat
- **Config**: `filebeat.yml`
- **Purpose**: Collect logs from services and send to Kafka

#### go-stash
- **Config**: `go-stash.yml`
- **Purpose**: Consume logs from Kafka and index to Elasticsearch

## Log Flow

```
Services → Filebeat → Kafka → go-stash → Elasticsearch → Kibana
```

## Metrics Flow

```
Services (/metrics) → Prometheus → Grafana
```

## Trace Flow

```
Services → Jaeger Agent → Jaeger Collector → Jaeger UI
```

## Quick Start

1. Start infrastructure:
   ```bash
   make infra-up
   ```

2. Verify all services:
   ```bash
   ./scripts/verify-infrastructure.sh
   ```

3. Access dashboards:
   - Grafana: http://localhost:3000
   - Prometheus: http://localhost:9090
   - Jaeger: http://localhost:16686
   - Kibana: http://localhost:5601

## Creating Grafana Dashboards

1. Login to Grafana (admin/admin)
2. Go to Dashboards → New Dashboard
3. Add panels with PromQL queries
4. Save dashboard

### Useful PromQL Queries

```promql
# Request rate
rate(http_requests_total[5m])

# Request duration P95
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))

# Error rate
rate(http_requests_total{status=~"5.."}[5m])

# Active goroutines
go_goroutines
```

## Kibana Index Patterns

1. Open Kibana: http://localhost:5601
2. Go to Management → Index Patterns
3. Create pattern: `logs-*`
4. Select timestamp field: `@timestamp`
5. Go to Discover to view logs

## Troubleshooting

### Prometheus not scraping services
- Check service is exposing `/metrics` endpoint
- Verify service is reachable from Prometheus container
- Check `prometheus.yml` configuration

### Logs not appearing in Kibana
- Verify Filebeat is running and collecting logs
- Check Kafka topics: `docker exec kafka kafka-topics --list --bootstrap-server localhost:9092`
- Verify go-stash is consuming from Kafka
- Check Elasticsearch indices: `curl http://localhost:9200/_cat/indices`

### Jaeger not showing traces
- Verify services are configured with Jaeger endpoint
- Check Jaeger collector logs: `docker logs jaeger`
- Ensure trace context is propagated in gRPC/HTTP calls

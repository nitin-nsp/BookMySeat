# BookmySeat - Movie/Event Ticket Booking Platform

A highly scalable, distributed ticket booking platform built with Go microservices architecture.

## 🚀 Quick Start

```bash
# Setup development environment
chmod +x scripts/setup-dev-env.sh scripts/verify-setup.sh
./scripts/setup-dev-env.sh

# Add Go tools to PATH (add to ~/.zshrc or ~/.bashrc)
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.zshrc

# Verify setup
./scripts/verify-setup.sh
```

## 📋 Tech Stack

- **Language**: Go 1.21+
- **Framework**: go-zero
- **Databases**: PostgreSQL, MongoDB
- **Cache**: Redis (3-node cluster for Redlock)
- **Message Queue**: Kafka
- **Distributed Transactions**: DTM
- **Monitoring**: Prometheus, Grafana, Jaeger, ELK Stack
- **Container Orchestration**: Kubernetes
- **Local Development**: Docker Compose, minikube

## 🏗️ Architecture

Microservices:
- **User Service**: Authentication & user management
- **Event Service**: Events, venues, shows, seats
- **Booking Service**: Seat reservation with Redlock
- **Payment Service**: Payment processing & refunds
- **Notification Service**: Email & SMS notifications
- **API Gateway**: REST API & routing

## 📁 Project Structure

```
BookmySeat/
├── .github/              # GitHub configs
├── .vscode/              # VSCode settings
├── common/               # Shared code
│   ├── config/
│   ├── middleware/
│   └── utils/
├── dev-plan/             # Planning & tracking
│   ├── complete.check_list.md
│   ├── ref/
│   └── prompts/
├── docs/                 # Documentation
│   ├── general/
│   └── phase1/
├── logs/                 # Service logs
├── monitoring/           # Observability
├── proto/                # Protobuf definitions
├── scripts/              # Utility scripts
├── services/             # Microservices
├── .env                  # Environment (not in git)
├── .gitignore
├── docker-compose.yml
├── go.mod
├── Makefile
├── PROJECT_RULES.md
├── README.md
└── STRUCTURE.md
```

## 📚 Documentation

- [Quick Start Guide](docs/general/QUICK_START.md)
- [Setup Guide](docs/general/SETUP.md)
- [Phase 1 Complete](docs/phase1/PHASE1_COMPLETE.md)
- [Implementation Checklist](dev-plan/complete.check_list.md)
- [Monitoring Guide](monitoring/README.md)
- [Project Rules](PROJECT_RULES.md)
- [Project Structure](STRUCTURE.md)

## ✅ Phase 1: Setup & Infrastructure - COMPLETED

- [x] Development environment setup
- [x] Project structure created
- [x] Docker Compose infrastructure configured
- [x] Monitoring stack setup complete
- [x] Project rules and documentation organized

## ✅ Phase 2: Proto Definitions & Code Generation - COMPLETED

- [x] Protobuf definitions for all 5 services
- [x] Generated Go code for gRPC communication
- [x] Automated generation scripts
- [x] Makefile integration

## 🚀 Getting Started

```bash
make init              # Initialize project
make infra-up          # Start infrastructure
make verify            # Verify all services
make open-grafana      # Open Grafana (admin/admin)
```

See [Quick Start Guide](docs/general/QUICK_START.md) for details.

## 🔧 Next Steps

1. **Phase 3**: User Service Implementation
2. **Phase 4**: Event Service Implementation
3. **Phase 5**: Booking Service Implementation

## 📝 License

MIT

# Observability Starter Kit

This repository provides a **production-grade** starter kit for deploying containerized microservices with full observability on Kubernetes.

## Features
- **4 Microservices**: Go, Flask, React, Spring Boot
- **Observability**: Prometheus, Grafana, Loki, Tempo, Jaeger
- **CI/CD**: GitHub Actions, Helm automation
- **Autoscaling**: HPA based on metrics

## Quickstart
1. Build Docker images:
   ```bash
   ./scripts/build-all.sh
   ```
2. Start Minikube:
   ```bash
   minikube start
   ```
3. Deploy observability stack:
   ```bash
   ./scripts/deploy-observability.sh
   ```
4. Deploy services:
   ```bash
   ./scripts/deploy-services.sh
   ```
5. Access dashboards:
   - Grafana: http://localhost:3000 (`admin`/`admin`)
   - Jaeger: http://localhost:16686

For detailed instructions, see [docs/README.md](docs/README.md).

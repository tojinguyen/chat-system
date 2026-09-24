# Distributed High-Throughput Chat System

A distributed real-time messaging system built with a Microservices architecture, engineered for high throughput, low latency, and horizontal scalability.

---

## 1. System Architecture & Core Services

The system is decoupled into a **Stateful Edge Layer** (connection management) and a **Stateless Processing Layer** (business logic and persistence):

- **`ws-gateway`**: Centralized WebSocket connection gateway managing real-time persistent sessions with clients, enforcing per-connection buffer limits to prevent memory bloat, and dispatching outbound messages to active user devices.
- **`chat-engine`**: Core message processing service supporting dual delivery modes (**gRPC** point-to-point and **NATS** pub/sub broker) and persisting chat history into ScyllaDB.
- **`api-service`**: RESTful API service handling user authentication (JWT), user profile, friendships, and conversation management, built following Clean Architecture principles.
- **`notification-service`**: Asynchronous notification worker handling offline alerts and push notifications for disconnected recipients.
- **`client-simulator`**: High-concurrency bot load generator simulating hundreds of concurrent users to benchmark throughput, stress-test the cluster, and measure edge network latency.

---

## 2. Technology Stack

### Core Language & Runtime
- **Go (Golang 1.22+)**: Primary language across all microservices, leveraging goroutines, channels, and low-latency concurrency primitives for optimal memory and CPU efficiency.

### Communication Protocols
- **WebSocket (Gorilla WebSocket)**: Full-duplex, low-latency bidirectional communication between clients and the WS Gateway.
- **gRPC & Protocol Buffers (Protobuf)**: High-performance binary RPC protocol for inter-service communication and synchronous point-to-point message dispatching.
- **RESTful API**: Standard HTTP interfaces for authentication and metadata management.

### Message Broker & Event Streaming
- **NATS Core / JetStream**: Lightweight, high-performance distributed message broker providing decoupled, asynchronous pub/sub buffering between the Chat Engine and Edge Gateways.

### Databases & Storage
- **ScyllaDB / Apache Cassandra**: Distributed NoSQL database based on LSM-Tree architecture, designed for write-heavy workloads and high-throughput chat message history persistence.
- **PostgreSQL**: Relational database (RDBMS) for structured domain metadata (users, friend graphs, conversation participants).
- **Redis Cluster**: In-memory data store for real-time presence management (online/offline tracking), gateway route registry, and fast caching.

### Observability & Monitoring
- **Prometheus**: Metric collection and time-series aggregation for system health and load metrics.
- **Grafana**: Real-time dashboards monitoring the 5 Golden Dimensions: Latency (P50, P95, P99), Throughput (msgs/sec), Worker Saturation, Error Rates, and Edge Network RTT.
- **OpenTelemetry & Jaeger**: End-to-end distributed tracing across microservice boundaries.

### Infrastructure & Deployment
- **Docker**: Containerization for all microservices and supporting components.
- **Kubernetes (Kind)**: Container orchestration and lifecycle management for local cluster environments.
- **NGINX Ingress Controller**: Edge reverse proxy handling incoming WebSocket and HTTP traffic routing into the cluster.
- **Makefile**: Automation tooling for builds, local image loading, and environment provisioning.
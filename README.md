# Distributed High-Throughput Chat System

Hệ thống nhắn tin phân tán thời gian thực (Real-time Messaging System) được xây dựng theo kiến trúc Microservices hướng tới khả năng chịu tải cao, độ trễ thấp và khả năng mở rộng quy mô linh hoạt.

---

## 1. Kiến Trúc Các Dịch Vụ (Core Services)

Hệ thống được thiết kế phân tách rõ ràng giữa tầng kết nối trạng thái (Stateful Edge) và tầng xử lý nghiệp vụ phi trạng thái (Stateless Processing):

- **`ws-gateway`**: Gateway quản lý kết nối WebSocket tập trung, chịu trách nhiệm duy trì phiên kết nối thời gian thực với các Client, tối ưu bộ nhớ đệm trên từng kết nối và phân phối tin nhắn đến các thiết bị của người dùng.
- **`chat-engine`**: Dịch vụ xử lý nghiệp vụ tin nhắn lõi (Core Business Engine), hỗ trợ 2 chế độ điều phối chuyển phát tin nhắn nội bộ (**gRPC** point-to-point và **NATS** pub/sub), thực hiện ghi dữ liệu tin nhắn vào ScyllaDB.
- **`api-service`**: Cung cấp RESTful API cho các nghiệp vụ xác thực (Authentication qua JWT), quản lý người dùng (Users), bạn bè (Friends) và cuộc hội thoại (Conversations), được xây dựng theo mô hình Clean Architecture.
- **`notification-service`**: Xử lý các tác vụ thông báo đẩy (Push Notifications) bất đồng bộ cho người dùng đang ngoại tuyến (Offline).
- **`client-simulator`**: Công cụ giả lập tải phát sinh đồng thời từ hàng trăm bot nhằm kiểm thử áp lực và đo lường độ trễ mạng thực tế của toàn bộ hệ thống.

---

## 2. Công Nghệ Sử Dụng (Technology Stack)

### Ngôn Ngữ & Framework Cốt Lõi
- **Go (Golang 1.22+)**: Ngôn ngữ chính phát triển toàn bộ hệ thống backend microservices, tận dụng Goroutine, Channel và Concurrency Patterns để tối ưu hiệu năng và bộ nhớ.

### Giao Thức Truyền Thông (Communication Protocols)
- **WebSocket (Gorilla WebSocket)**: Giao thức truyền thông hai chiều thời gian thực giữa Client và WS Gateway.
- **gRPC & Protocol Buffers (Protobuf)**: Giao thức RPC nhị phân nội bộ tốc độ cao phục vụ điều phối tin nhắn trực tiếp giữa các service.
- **RESTful API**: Chuẩn giao tiếp HTTP cho các thao tác xác thực và quản lý tài nguyên.

### Message Broker & Điều Phối Bất Đồng Bộ
- **NATS Core / JetStream**: Message broker phân tán hiệu năng cao, đóng vai trò làm vùng đệm và truyền tải sự kiện (Event Streaming) bất đồng bộ giữa Chat Engine và các cụm Gateway.

### Cơ Sở Dữ Liệu & Bộ Nhớ Đệm (Databases & Storage)
- **ScyllaDB / Apache Cassandra**: Cơ sở dữ liệu phân tán NoSQL dựa trên kiến trúc LSM-Tree, chuyên dụng lưu trữ lịch sử tin nhắn với tốc độ ghi cực lớn (Write-heavy Workload).
- **PostgreSQL**: Cơ sở dữ liệu quan hệ (RDBMS) lưu trữ các thực thể nghiệp vụ cốt lõi (Tài khoản người dùng, danh sách bạn bè, thông tin hội thoại).
- **Redis Cluster**: Bộ nhớ đệm tốc độ cao phục vụ theo dõi trạng thái trực tuyến/ngoại tuyến (Presence Management), định tuyến kết nối Gateway và Caching dữ liệu tạm.

### Khả Năng Quan Sát (Observability & Monitoring)
- **Prometheus**: Thu thập và lưu trữ số liệu giám sát hệ thống theo chuỗi thời gian (Metrics Time-series).
- **Grafana**: Hiển thị bảng điều khiển trực quan theo dõi các chỉ số vàng (Golden Signals): Độ trễ (P50, P95, P99), Lưu lượng (Throughput), Mức độ bão hòa (Saturation), Tỷ lệ lỗi (Errors) và Độ trễ mạng biên (Client Network RTT).
- **OpenTelemetry & Jaeger**: Theo dõi vết phân tán (Distributed Tracing) toàn diện hành trình của gói tin qua từng microservice.

### Hạ Tầng & Triển Khai (Infrastructure & Deployment)
- **Docker**: Đóng gói container hóa toàn bộ các dịch vụ và thành phần phụ trợ.
- **Kubernetes (Kind)**: Điều phối và quản lý vòng đời container trên môi trường cụm (Cluster).
- **NGINX Ingress Controller**: Cổng đón nhận và định tuyến lưu lượng mạng bên ngoài vào các dịch vụ trong cụm Kubernetes.
- **Makefile**: Tự động hóa toàn bộ quy trình build, nạp image và triển khai hạ tầng.
# Kế hoạch Hiện thực hóa: 3 Tuần Nâng tầm Kiến trúc & Đo lường Hệ thống Chat Phân tán

## 🎯 Mục tiêu Tổng quan
Chuyển hóa hệ thống từ giai đoạn "chạy được ở Happy Path" sang **Production-Ready Distributed System** với khả năng quan sát toàn diện (Deep Observability), phân tích định lượng hiệu năng (Benchmarking & Chaos Engineering), và chịu tải tới hạn (Stress Testing & Continuous Profiling).

---

## 🏛️ Góc nhìn Kiến trúc & 4 Lăng kính Sự cố (Architecture Analysis)

Trước khi đi vào từng tuần, toàn bộ kế hoạch được thiết kế dựa trên 4 Lăng kính Sự cố:

1. **Lăng kính Thời gian & Bất đồng bộ (Time & Concurrency Lens):**
   - Độ trễ nội bộ (Internal Latency) giữa các bước: Cassandra I/O RTT vs Redis Lock RTT. Khi tải tăng đột biến, hàng đợi trong Go Channel của `PartitionedPool` sẽ bị phình to (Backlog saturation), làm trôi lệch P99 của toàn bộ hệ thống.
2. **Lăng kính Sự cố Mạng & Node (Network & Partial Failure Lens):**
   - Sự khác biệt căn bản giữa **Point-to-Point Synchronous (gRPC)** và **Asynchronous Pub/Sub (NATS)** khi có Jitter/Network Latency: gRPC gây giữ kết nối, tăng Goroutine & RAM ở Gateway; NATS đệm trong buffer nhưng có nguy cơ rớt tin nếu buffer đầy.
3. **Lăng kính Giới hạn Dữ liệu & Tải đột biến (Data Boundary & Load Lens):**
   - Khi tải lên 10,000 msg/s, GC Pressure từ JSON serialization và gocql frame allocations sẽ chiếm dụng CPU, gây stop-the-world spikes.
4. **Lăng kính Trạng thái & Vòng đời (State Machine & Lifecycle Lens):**
   - Đo lường chính xác từng trạng thái trong trace lifecycle: Gateway Inbound $\to$ Redis Idempotency $\to$ Cassandra Commit $\to$ Sender Ack $\to$ Outbound Dispatch.

---

## 📅 Lộ trình Triển khai Chi tiết

### 📌 TUẦN 1: Bịt điểm mù Observability (Deep Observability & 4 Golden Signals)

#### 1. Thêm Child Spans (OpenTelemetry / Tempo) cho Database & Caching
* **Vấn đề hiện tại:** Trace hiện tại chỉ bao quát tầng Inbound $\to$ Usecase tổng quát. Khi P99 tăng vọt, không thể xác định điểm nghẽn nằm ở Redis Lock, Cassandra Write, hay Outbound Dispatch.
* **Giải pháp:**
  - Bổ sung OTel Child Spans có ngữ cảnh Semantic Conventions:
    + `db.system: cassandra`, `db.operation: insert`, `db.statement: INSERT INTO messages...` trong `CassandraMessageRepository.SaveMessage` và `GetMessagesByConversation`.
    + `db.system: redis`, `db.operation: setnx`, `db.redis.key: idempotency:msg:...` trong `redisIdempotencyRepository.AcquireLock` và `ReleaseLock`.
    + `db.system: redis`, `db.operation: hgetall/hlen` trong `redisPresenceReader`.
  - Tự động ghi nhận Span Status Error và `span.RecordError(err)` khi các lệnh I/O gặp timeout/lỗi.

#### 2. Metric đo Độ sâu và Độ bão hòa Channel trong `PartitionedPool`
* **Vấn đề hiện tại:** `PartitionedPool` sử dụng buffered Go channels (`jobCh`). Khi 1 partition bị nghẽn (Hot partition) hoặc downstream DB chậm, channel bị đầy dần dẫn đến block/timeout mà không có metric cảnh báo trước.
* **Giải pháp:**
  - Khai báo Prometheus Metrics trong `pkg/telemetry/metrics.go`:
    + `chat_worker_channel_depth` (GaugeVec by `worker_id`): `len(w.jobCh)`
    + `chat_worker_channel_capacity` (GaugeVec by `worker_id`): `cap(w.jobCh)`
    + `chat_worker_channel_saturation_ratio` (GaugeVec by `worker_id`): `len(w.jobCh) / cap(w.jobCh)`
    + `chat_worker_jobs_total` (CounterVec by `worker_id`, `status`): Đếm số job xử lý thành công / lỗi / panic.
  - Tích hợp hàm cập nhật metric theo chu kỳ nhẹ (tick interval) hoặc cập nhật trực tiếp tại `Submit` / `runWorker`.

#### 3. Xây dựng Grafana Dashboard "4 Biểu đồ Vàng" (Golden Signals)
* **Vị trí:** `deployments/configs/grafana-dashboards/chat-golden-signals.json` và đăng ký tự động qua `grafana-datasources.yaml` / dashboard provisioning.
* **4 Panel Trọng tâm:**
  1. **Throughput (Traffic):** `sum(rate(chat_messages_total[1m])) by (service, event_type)` (Đo RPS hệ thống).
  2. **Latency Percentiles (P50, P95, P99):** `histogram_quantile(0.99, sum(rate(chat_message_latency_seconds_bucket[1m])) by (le, stage))` (Phân rã theo stage: Idempotency, Cassandra Insert, Realtime Dispatch).
  3. **Error Rate (%):** Tỷ lệ lỗi so với tổng request `(sum(rate(chat_messages_total{status="error"}[1m])) / sum(rate(chat_messages_total[1m]))) * 100`.
  4. **Worker Channel Saturation (%):** `max(chat_worker_channel_depth / chat_worker_channel_capacity) * 100` (Phát hiện nghẽn worker pool).

---

### 📌 TUẦN 2: So găng gRPC vs NATS Broker Delivery (Benchmarking & Chaos Testing)

#### 1. Kịch bản Benchmark gRPC Delivery Mode
* Cấu hình `DISPATCH_MODE=grpc` trên cụm `chat-engine`.
* Chạy `client-simulator` với kịch bản multi-node (ví dụ: 1,000 -> 3,000 active bots trên nhiều Gateway instances).
* Thu thập số liệu: P50, P95, P99 Latency (qua Prometheus & Grafana), CPU usage của `chat-engine`, RAM/Goroutine count của `ws-gateway`.

#### 2. Kịch bản Benchmark NATS Broker Delivery Mode
* Cấu hình `DISPATCH_MODE=broker` (NATS Core / JetStream routing theo `chat.gateway.{node_id}`).
* Chạy cùng một profile tải từ `client-simulator`.
* Ghi nhận và so sánh độ trễ P99, khả năng buffering, CPU và RAM giữa 2 phương thức.

#### 3. Chaos Engineering: Bơm độ trễ mạng (Chaos Mesh Latency Injection 100ms)
* Áp dụng `01-network-delay-grpc.yaml` và `02-network-delay-nats.yaml` trong thư mục `deployments/k8s/chaos/` (Inject 100ms latency giữa `chat-engine` và `ws-gateway`).
* **Quan sát hiện tượng:**
  - **Mode gRPC:** Hiện tượng connection pool exhaustion, goroutines tích tụ ở `chat-engine`, P99 tăng dốc đứng, lan truyền ngược (cascading failure / backpressure) làm tắc nghẽn `PartitionedPool`.
  - **Mode NATS:** NATS client buffer đóng vai trò đệm async, Goroutines ở `chat-engine` không bị giữ; tuy nhiên cần giám sát RAM tiêu thụ của NATS broker và nguy cơ chậm tin nhắn tới client cuối (End-to-End Latency).

#### 4. Tổng kết & Đóng gói Tài liệu So sánh Thực tế
* Tạo tài liệu chuẩn mực tại: `knowledge/system_design/04_grpc_vs_nats_realtime_delivery_tradeoffs.md`.
* Nội dung gồm: Bảng so sánh định lượng (RPS, Latency P99, CPU/RAM, Fault Tolerance), đồ thị so sánh, phân tích các trường hợp nên chọn gRPC vs NATS trong thực tế (Architecture Decision Record).

---

### 📌 TUẦN 3: Đẩy tải tới ngưỡng vỡ & Continuous Profiling (Stress Testing & Performance Tuning)

#### 1. Stress Testing tới Ngưỡng Vỡ (5,000 $\to$ 10,000+ msg/sec)
* Tăng dần số lượng bots và tần suất phát tin trong `client-simulator` hoặc script `k6`:
  - Bước 1: 2,000 msg/s (Baseline).
  - Bước 2: 5,000 msg/s (High load).
  - Bước 3: 10,000 msg/s (Stress point).
* Xác định chính xác điểm bắt đầu xuất hiện tình trạng drop tin nhắn (Channel overflow, Client disconnect, HTTP/WebSocket drop rate).

#### 2. Bắt "Thủ phạm" CPU & Memory Contention bằng Grafana Pyroscope
* Quan sát Flamegraph trên Pyroscope UI (`http://localhost:4040`):
  - **CPU Profile:** Phân tích tỷ lệ CPU đốt vào:
    + `encoding/json` serialization/deserialization.
    + `gocql` query formatting, frame decoding, reflection.
    + `go-redis` pipeline & networking overhead.
    + Runtime scheduler & garbage collection (GC pause).
  - **Mutex & Block Profile:**
    + Mutex contention trong Go Channel hoặc Sync Maps ở Gateway SessionManager.
    + Lock contention trong Redis connection pool.

#### 3. Hiện thực hóa Tối ưu hóa (Fix Bottleneck)
* Dựa trên kết quả Flamegraph thực tế, thực hiện tối ưu điểm nghẽn lớn nhất:
  - *Nếu là JSON parsing:* Chuyển đổi sang `sonic`, `easyjson` hoặc zero-allocation protobuf encoding cho luồng hot path.
  - *Nếu là Gocql/Cassandra I/O:* Tối ưu prepared statements caching, tuning gocql connection pool size, batching hợp lý.
  - *Nếu là Redis RTT/Contention:* Áp dụng Redis Pipelining / MGet / Lua script cho việc kiểm tra Idempotency + Presence trong cùng 1 network round-trip.
* Chạy lại bài test tải để nghiệm thu kết quả cải thiện (Before vs After metrics).

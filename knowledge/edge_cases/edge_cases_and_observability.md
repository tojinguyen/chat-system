# 📌 Edge Cases, Failure Modes & Observability Backlog

Tài liệu này ghi nhận các góc tối kỹ thuật (Edge Cases), bẫy chạy đua (Race Conditions), và các kịch bản lỗi trong môi trường phân tán cần được theo dõi qua hệ thống Metrics/Logging và lên kế hoạch xử lý/tối ưu hóa sau này.

---

## 1. Edge Case: Phantom ACK / False Success do Race Condition giữa Redis Lock & DB Persistence

### 🔍 Mô tả & Bối cảnh (Context & Failure Scenario)
- **Vị trí:** `services/chat-engine/internal/usecase/chat_usecase.go` & `idempotency_repo.go`.
- **Kịch bản:**
  1. Client gửi tin nhắn `M1` (`client_msg_id = X`).
  2. Chat Worker 1 nhận `M1`, thực hiện `AcquireLock(X)` trên Redis thành công (`SetNX`).
  3. Chat Worker 1 bắt đầu ghi `M1` vào Cassandra (quá trình này mất vài giây hoặc timeout do mạng chập chờn / Cassandra chậm).
  4. Trong lúc Worker 1 chưa ghi xong, Client retry hoặc NATS Broker bắn lại bản tin duplicate `M1'` (`client_msg_id = X`).
  5. Chat Worker 2 (hoặc Worker 1) nhận `M1'`, kiểm tra `AcquireLock(X)` -> Thấy key đã tồn tại (`!isNew`).
  6. Worker coi đây là tin nhắn trùng lặp đã hoàn tất -> **Bắn ngay `Sender ACK` về cho Client (Phantom ACK)**.
  7. Sau đó, thao tác ghi Cassandra ở bước 3 bị **Thất Bại (Timeout / Crash)** -> Worker gọi `ReleaseLock(X)`.

### ⚠️ Hậu quả (Impact)
- **Silent Data Loss:** Client nhìn thấy tin nhắn đã gửi thành công (giao diện hiện 1 tích), nhưng thực chất tin nhắn **chưa từng được ghi vào cơ sở dữ liệu**. Khi reload app hoặc phía Receiver mở app, tin nhắn biến mất hoàn toàn.

### 📊 Điểm quan sát & Metric giám sát (Observability Signals)
- **Log cảnh báo:** `WARN: Redis idempotency check failed` hoặc `Duplicate message detected: client_msg_id=...` xuất hiện cùng thời điểm với `persistence failed`.
- **Prometheus Metric cần bổ sung sau này:**
  - `chat_engine_idempotency_lock_acquired_total`
  - `chat_engine_idempotency_lock_released_on_error_total`
  - Tỷ lệ lệch (Discrepancy) giữa tổng số `Sender ACK dispatched` và tổng số record được `INSERT` thành công vào Cassandra.

### 🛠️ Giải pháp dài hạn (Long-term Remediation)
- **FSM State Machine trên Idempotency Key:**
  - `PROCESSING` (TTL ngắn 10-30s, đóng vai trò Distributed Mutex).
  - `COMPLETED` (TTL dài 24h, chỉ set sau khi DB ghi thành công).
- **Concurrent Duplicate Handling:**
  - Nếu duplicate đến khi đang `PROCESSING`: Không gửi ACK ngay; hoặc chờ (polling 500ms) hoặc reject/NACK để Broker redeliver sau.

---

## 2. Edge Case: Cassandra Zombie Write sau khi Context Timeout

### 🔍 Mô tả & Bối cảnh
- Worker thiết lập `dbCtx, cancel := context.WithTimeout(ctx, dbTimeout)` (mặc định cấu hình từ `timeout_seconds`).
- Nếu Cassandra bị quá tải hoặc Garbage Collection (GC Pause), request ở phía Go driver có thể bị timeout và trả về error, worker rollback Redis lock.
- Tuy nhiên, coordinator node phía Cassandra có thể vẫn nhận được mutation trước đó và ghi vào CommitLog/MemTable thành công sau đó vài mili-giây.

### 📊 Điểm quan sát
- Kiểm tra các bản ghi Cassandra có `client_msg_id` tồn tại nhưng application log lại ghi nhận `persistence failed: context deadline exceeded`.

---

## 3. Edge Case: Receiver Connection Drop ngay sau khi Query Presence (Ghost Online)

### 🔍 Mô tả & Bối cảnh
- Chat Engine query Redis Presence thấy Client B đang online tại `gateway_node_02`.
- Trong tích tắc trước khi Outbound Message từ NATS đến được `gateway_node_02`, Client B bị mất kết nối (rớt mạng, tắt 4G, killed app).
- `gateway_node_02` không tìm thấy WebSocket connection trong local memory map -> Tin nhắn có nguy cơ bị drop nếu không có cơ chế fallback hoặc Receiver ACK timeout.

### 🛠️ Giải pháp quan sát & xử lý
- Đếm metric `ws_gateway_local_delivery_dropped_total`.
- WS Gateway nếu không tìm thấy socket connection cục bộ phải có cơ chế phản hồi lại Broker hoặc publish sang `chat.notification` để bù Push Notification.

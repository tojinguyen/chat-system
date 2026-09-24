# Cẩm Nang Kiểm Thử Hiệu Năng & Chaos Engineering (Benchmarking Runbook)
## So Sánh Thực Chiến: gRPC (Point-to-Point) vs Message Broker (NATS) & Đo Lường Network Delay

Tài liệu này cung cấp quy trình chuẩn hóa từng bước để tiến hành đo tải (Stress Test) và thử nghiệm phá vỡ hệ thống (Chaos Engineering) trên cụm Kubernetes, nhằm kiểm chứng tính đúng đắn của Dashboard quan sát (Observability) và phân tích định lượng ưu/nhược điểm sống còn giữa hai mô hình điều phối: **gRPC** và **NATS Broker**.

---

## 1. Bản Đồ Kiến Trúc: 2 Chặng Mạng Độc Lập

Để quan sát hệ thống chính xác, cần phân định rõ 2 chặng truyền thông trong hệ thống Chat:

```text
[Client Bots] <=== (CHẶNG 1: Edge Network) ===> [WS Gateway] <=== (CHẶNG 2: Internal Cluster) ===> [Chat Engine / NATS]
                   Internet / 4G / Wi-Fi                             gRPC vs NATS Broker
                   (Đo bằng WebSocket Ping RTT)                      (Nơi gRPC và NATS đối đầu)
```

1. **Chặng 1 (Mạng Biên - Edge):** Kết nối WebSocket giữa Client và WS Gateway qua Ingress NGINX. Chặng này dùng để đo Network Delay thuần túy đưa lên **Panel 1.3** và CLI Simulator.
2. **Chặng 2 (Mạng Nội Bộ Backend):** Kết nối giữa WS Gateway và Chat Engine. Đây là chiến trường quyết định sự khác biệt về khả năng chịu tải và cách ly sự cố giữa **gRPC (Đồng bộ)** và **NATS (Bất đồng bộ)**.

---

## 2. Thông Số Tải Chuẩn (Workload Profile)

Cấu hình trong file `deployments/k8s/09-client-simulator.yaml`:

```yaml
data:
  SIM_BOTS: "500"          # 500 Bots kết nối đồng thời qua WebSocket
  SIM_CONVS: "10"          # Mỗi bot mở 10 cuộc hội thoại 1-1 chéo nhau
  SIM_INTERVAL: "100ms"    # Tần suất gửi: 100ms/tin (10 tin/giây/bot)
  SIM_DURATION: "2m"       # Thời lượng mỗi đợt test: 2 phút
```
* **Tổng Throughput phát sinh:** `500 bots * 10 msgs/s = 5,000 messages/giây`.

---

## 3. Quy Trình 3 Bước Thực Chiến Chi Tiết

### BƯỚC 0: Chuẩn Bị & Cập Nhật Code Lên Cụm (Chạy 1 lần)
Trước khi bắt đầu, đảm bảo image mới nhất đã được build và nạp vào cụm Kind:

```powershell
# Build và nạp Gateway mới (chứa handler HEARTBEAT_ACK & metric RTT)
make build-load-gw

# Build và nạp Simulator mới (chứa Jitter Ping & Tracker phân tách)
make build-load-sim
```

* **Địa chỉ Grafana:** `http://localhost:3000` (User: `admin` / Pass: `admin`).
* **Dashboard theo dõi:** `Chat System - 5 Golden Dimensions (gRPC vs NATS)`.

---

### PHẦN 1: Kiểm Chứng Đo Lường Network Delay (Chặng 1: Client <-> Gateway)

**Mục tiêu:** Xác thực Panel 1.3 và CLI Simulator bắt được chính xác độ trễ mạng biên khi người dùng dùng mạng 4G chập chờn.

1. **Khởi động Simulator:**
   ```powershell
   make sim-up
   ```

2. **Xem Dashboard trực tiếp trên Terminal:**
   ```powershell
   make logs-sim
   ```
   *Quan sát ban đầu:* Mục `Network Delay / RTT` dao động ở mức mạng nội bộ lý tưởng (`~1.0ms - 2.5ms`).

3. **Mở tab Terminal thứ 2, kích hoạt Chaos độ trễ mạng di động:**
   ```powershell
   make chaos-mobile-delay
   # Kịch bản: Inject 80ms latency + 25ms jitter giữa client-simulator và nginx-gateway
   ```

4. **Hiện tượng quan sát:**
   * **Trên Terminal Simulator (`make logs-sim`):**
     - Dòng `Network Delay / RTT` lập tức nhảy vọt lên mức **`~80.00ms - 105.00ms`**.
     - Dòng `Ước tính thời gian xử lý Backend (Median E2E - Network)` vẫn ổn định ở mức thấp (`~5ms - 7ms`), chứng minh backend không bị nghẽn mà nguyên nhân do mạng người dùng.
   * **Trên Grafana Dashboard:**
     - **Panel 1.3 (Client Network Delay / Edge RTT):** Đường P50, P95, P99 dựng đứng lên tương ứng `~80 - 105ms`.

5. **Dọn dẹp sau Phần 1:**
   ```powershell
   make chaos-clean
   make sim-down
   ```

---

### PHẦN 2: Trận Chiến Sinh Tử: gRPC vs NATS Broker (Chặng 2)

**Mục tiêu:** Chứng minh hiện tượng sụp đổ dây chuyền (Cascading Failure / Head-of-line Blocking) của gRPC khi có mạng chậm, và khả năng miễn nhiễm hoàn toàn của NATS Broker.

---

#### 🥊 Vòng A: Chế Độ gRPC Delivery (Point-to-Point Synchronous)

1. **Cấu hình gRPC Mode:**
   - Mở file `deployments/k8s/01-configmap-secrets.yaml`, đảm bảo:
     ```yaml
     DELIVERY_MODE: "grpc"
     ```
   - Nạp lại cấu hình:
     ```powershell
     make reload
     ```

2. **Khởi động tải 5,000 msg/s:**
   ```powershell
   make sim-up
   ```

3. **Giai đoạn A1 (30 giây đầu - Đường chuẩn Happy Path):**
   - Quan sát Grafana:
     - **Panel 1.1 (Dispatch Latency):** P50 rất tốt (~`2ms - 3ms`).
     - **Panel 2.1 (Throughput):** Đạt đỉnh cao nhất (`~4,500 - 5,000 msgs/s`).
     - **Panel 4.1 (Worker Saturation):** Rất thấp (`< 15%`).

4. **Giai đoạn A2 (Bơm lỗi mạng vào gRPC):**
   ```powershell
   make chaos-delay-grpc
   # Kịch bản: Inject 100ms latency giữa chat-engine và ws-gateway
   ```
   *Hiện tượng sụp đổ dây chuyền (The Collapse):*
   - **Panel 4.1 (Worker Pool Saturation):** Nhảy vọt lên **`100% (Bão hòa kịch khung)`**. Toàn bộ worker goroutine bị giữ chân (block) chờ phản hồi RPC từ Gateway.
   - **Panel 4.2 (Total Pending Backlog):** Hàng đợi channel bị tắc nghẽn, tồn đọng hàng ngàn jobs chưa xử lý.
   - **Panel 2.1 (Throughput):** Tụt dốc không phanh từ `5,000 msgs/s` xuống còn vài trăm msgs/s.
   - **Panel 3.1 & 3.2 (Error & Rejection):** Xuất hiện lỗi gRPC timeout/unavailable.

5. **Dọn dẹp Vòng A:**
   ```powershell
   make chaos-clean
   make sim-down
   ```

---

#### 🥊 Vòng B: Chế Độ NATS Broker Delivery (Asynchronous Decoupled)

1. **Chuyển sang Broker Mode:**
   - Mở file `deployments/k8s/01-configmap-secrets.yaml`, sửa thành:
     ```yaml
     DELIVERY_MODE: "broker"
     ```
   - Nạp lại cấu hình:
     ```powershell
     make reload
     ```

2. **Khởi động tải 5,000 msg/s:**
   ```powershell
   make sim-up
   ```

3. **Giai đoạn B1 (30 giây đầu - Đường chuẩn NATS):**
   - Quan sát Grafana:
     - **Panel 1.1 (Dispatch Latency):** P50 ở mức `~4ms - 5ms` (cao hơn gRPC khoảng 1-2ms do đi qua broker trung gian).
     - **Panel 2.1 (Throughput):** Ổn định ở mức `~4,500 - 5,000 msgs/s`.
     - **Panel 4.1 (Worker Saturation):** Ở mức thấp (`< 15%`).

4. **Giai đoạn B2 (Bơm lỗi mạng vào NATS):**
   ```powershell
   make chaos-delay-nats
   # Kịch bản: Inject 100ms latency giữa ws-gateway và nats broker
   ```
   *Hiện tượng vững như bàn thạch (Resilient Behavior):*
   - **Panel 4.1 (Worker Pool Saturation):** **VẪN NẰM IM DƯỚI ĐÁY `< 15%`!** Worker của Chat Engine chỉ mất vài micro-giây (`~50us`) để publish tin vào NATS rồi quay sang nhận tin mới ngay lập tức. Lõi Chat Engine hoàn toàn không bị ảnh hưởng.
   - **Panel 2.1 (Throughput):** Giữ vững thông lượng, không hề bị rớt thảm hại như gRPC.
   - Tin nhắn được xếp hàng an toàn trong NATS Buffer/Queue chờ Gateway kéo về khi mạng thông thoáng trở lại.

5. **Dọn dẹp kết thúc benchmark:**
   ```powershell
   make chaos-clean
   make sim-down
   ```

---

## 4. Bảng Tổng Kết Đối Chiếu (Architectural Scorecard)

| Tiêu Chí So Sánh | gRPC Mode (Synchronous Point-to-Point) | NATS Broker Mode (Asynchronous Queue) |
| :--- | :--- | :--- |
| **Độ trễ P50 lúc bình thường** | **Cực thấp (~1 - 2ms)**, tối ưu nhất | ~3 - 5ms (chậm hơn 1-2ms do hop qua broker) |
| **Mức tiêu hao tài nguyên đệm** | Zero-buffer, không cần quản lý broker | Cần bộ nhớ RAM để duy trì hàng đợi broker |
| **Khi mạng Downstream lag 100ms**| **Worker Pool nghẽn 100%**, Throughput sập | **Worker Pool < 15%**, thông lượng giữ vững |
| **Khi Gateway Pod bị Crash** | Lập tức văng lỗi RPC, rớt tin nhắn | Tin nhắn chờ trong queue, Pod mới sống lại lấy tiếp |
| **Mức độ phụ thuộc (Coupling)** | **Khớp nối cứng (Tight Coupling)** | **Tách rời hoàn toàn (Loosely Coupled)** |
| **Khuyến nghị kiến trúc** | Phù hợp hệ thống nội bộ LAN có SLA mạng cực cao | **Bắt buộc cho hệ thống quy mô lớn, tải cao, đa vùng** |

---

## 5. Cheatsheet Lệnh Nhanh (Quick Reference)

```powershell
# --- Bật / Tắt Simulator ---
make sim-up            # Khởi động simulator phát tải
make sim-down          # Dừng simulator
make logs-sim          # Xem dashboard trực tiếp

# --- Nạp Cấu Hình ---
make reload            # Áp dụng thay đổi DELIVERY_MODE

# --- Kích Hoạt Chaos ---
make chaos-mobile-delay   # Test Panel 1.3: Delay mạng Client (Chặng 1)
make chaos-delay-grpc     # Test gRPC: Delay mạng Engine <-> Gateway (Chặng 2)
make chaos-delay-nats     # Test NATS: Delay mạng Gateway <-> NATS (Chặng 2)
make chaos-kill-gw        # Test Pod Kill: Đột tử 1 Pod Gateway

# --- Dọn Dẹp Chaos ---
make chaos-clean       # Xóa toàn bộ kịch bản lỗi, khôi phục mạng bình thường
```

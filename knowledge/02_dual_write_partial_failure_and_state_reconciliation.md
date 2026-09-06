# 2. Dual-Write Partial Failure, Poison Pill & State Reconciliation

---

## 1. Bài toán Dual-Write & Thất bại một phần (Partial Failure)

Trong kiến trúc phân tán hướng sự kiện (Event-Driven Architecture), một Chat Worker thường thực hiện 2 thao tác liên tiếp:
1. **Lưu dữ liệu vào Database:** Ghi tin nhắn an toàn vào Cassandra (`Source of Truth`).
2. **Kích hoạt Side-Effect qua Broker:** Gửi `Sender ACK` hoặc `Outbound Event` tới NATS Broker.

### Tình huống:
Nếu **Bước 1 THÀNH CÔNG** (tin đã vào Cassandra) nhưng **Bước 2 THẤT BẠI** (NATS Broker lag hoặc Gateway Node bị rớt mạng):
- **Nếu Worker `return err`:** 
  - NATS sẽ coi như tin nhắn chưa được xử lý và kích hoạt **Broker Redelivery**.
  - *Rủi ro:* Nếu Gateway Node đó chết hẳn, tin nhắn sẽ bị retry vô tận (**Poison Pill**), gây tắc nghẽn toàn bộ hàng đợi (**Head-of-Line Blocking**), đồng thời có nguy cơ gửi trùng tin nhắn nhiều lần cho người nhận (Receiver).
- **Nếu Worker log cảnh báo và `return nil`:**
  - Tin nhắn được coi là hoàn tất vì dữ liệu gốc đã nằm an toàn trong Database.
  - Luồng xử lý không bị nghẽn, duy trì Throughput cao nhất (**Fast-Fail & Move-On**).

---

## 2. Vấn đề "Phantom Message" & Split-Brain giữa Client và Server

Khi Worker chọn `return nil`, một Edge Case xảy ra ở phía Client:
1. Client A gửi tin nhắn với `client_msg_id = uuid_123`.
2. Server lưu Cassandra thành công nhưng chiều gửi ACK về cho Client A bị đứt gói tin.
3. Client A hết thời gian chờ (Timeout) và gửi lại (Retry) `uuid_123`.
4. Mạng của Client A vẫn chưa ổn định, tất cả các lần retry đều không nhận được ACK.
5. **Hậu quả nếu thiết kế Client kém:** Client A đánh dấu tin nhắn là **Gửi thất bại (Chấm than đỏ ⚠️)**, trong khi thực tế tin nhắn đã nằm trên Server và người nhận (Client B) đã đọc được!

---

## 3. Giải pháp Chuẩn Production: State Reconciliation & Smart Client Pattern

Các hệ thống quy mô lớn (Telegram, WhatsApp, Discord) giải quyết triệt để vấn đề này qua 3 nguyên tắc:

### A. Phân biệt Permanent Failure vs Transient Failure
- **Permanent Failure (Lỗi nghiệp vụ vĩnh viễn):** Bị chặn, bị kích khỏi nhóm, vi phạm chính sách ➔ Server trả về mã lỗi rõ ràng ➔ Client hiển thị **Chấm than đỏ ⚠️ ngay lập tức**.
- **Transient Failure (Lỗi mạng/Timeout tạm thời):** Không nhận được ACK trong 5s ➔ Client **KHÔNG báo đỏ**, giữ trạng thái "Đang gửi (Đồng hồ xoay ⏳)" và kích hoạt kiểm tra Socket.

### B. Tự phục hồi kết nối (Socket Self-Healing)
- Khi nghi ngờ kết nối bị "đứt ngầm" (Half-Open Socket), Client chủ động ngắt kết nối WebSocket cũ và kết nối lại (Reconnect) tới Gateway mới.

### C. Đối soát trạng thái 2 chiều (State Reconciliation)
- Khi kết nối lại, Client gửi danh sách các tin nhắn đang ở trạng thái chờ:
  ```json
  POST /v1/messages/reconcile
  {
      "pending_client_msg_ids": ["uuid_123"]
  }
  ```
- **Phía Server:** Tra cứu nhanh Redis/Cassandra:
  - Nếu `uuid_123` đã có trong DB ➔ Trả về `"STATUS": "PERSISTED"`.
  - Client tự động chuyển UI từ **"Đang gửi ⏳"** sang **"Đã gửi ✔"** (Không bao giờ bị lỗi ảo!).
  - Nếu `uuid_123` chưa có ➔ Client mới thực hiện gửi lại hoặc cho phép user bấm "Thử lại".

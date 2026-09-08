# 1. Cassandra UPSERT vs RDBMS & Bản chất Tầng Idempotency (Redis)

---

## 1. Bẫy Kiến Trúc: Cassandra xử lý Trùng Primary Key như thế nào?

### ❌ Tư duy RDBMS thông thường (Postgres, MySQL)
- Trong cơ sở dữ liệu quan hệ, câu lệnh `INSERT` khi gặp trùng khóa chính (`Primary Key`) sẽ lập tức **ném ra lỗi vi phạm ràng buộc (`Duplicate Key Violation`)** và huỷ transaction.

### 🟢 Bản chất thực tế của Cassandra (LSM-Tree NoSQL)
- **Cassandra KHÔNG ném ra lỗi khi trùng Primary Key.**
- Trong Cassandra, câu lệnh `INSERT` thực chất là một **UPSERT (Insert or Overwrite)**.
- Nếu bạn thực thi lệnh `INSERT` vào cùng một bộ khóa `((conversation_id), id)` nhiều lần:
  - Cassandra sẽ **âm thầm ghi đè (overwrite)** các trường dữ liệu với timestamp mới nhất.
  - Trình điều khiển (Driver `gocql`) luôn nhận được kết quả thành công (`nil error`).
- *(Lưu ý: Cú pháp `INSERT ... IF NOT EXISTS` là Lightweight Transaction sử dụng thuật toán Paxos, cực kỳ tốn chi phí mạng và giảm thông lượng ghi trầm trọng, không bao giờ dùng cho hệ thống Chat).*

---

## 2. Tại sao BẮT BUỘC phải có Tầng Idempotency Key (Redis)?

Nếu Cassandra đã tự động ghi đè không báo lỗi, tại sao hệ thống phân tán vẫn bắt buộc phải có tầng Idempotency ở Redis phía trước?

### A. Ngăn chặn Duplicate Side-Effects (Tác dụng phụ ngoài luồng)
Lưu vào Database chỉ là bước đầu tiên trong pipeline xử lý. Sau đó Worker còn phải:
1. Gửi **Sender ACK** cho người gửi.
2. Tra cứu Presence và **Dispatch tin nhắn** tới WebSocket Gateway của người nhận (Receiver) hoặc gọi **Notification Service (APNs/FCM Push)**.

👉 Nếu không có Redis chặn từ đầu:
- Khi mạng lag, Client retry gửi lại cùng 1 tin nhắn 2 lần.
- Cả 2 lần đều đi qua Worker ➔ Người nhận (Client B) sẽ nhận **2 lần tin nhắn**, điện thoại kêu **2 lần ting ting**, và Notification Service bị tiêu tốn quota vô ích.

---

### B. Tránh Write Amplification & Giảm tải Compaction cho Cassandra
- Mỗi lần `INSERT` ghi đè trong Cassandra sẽ tạo ra một bản ghi mới trong `MemTable` và xả ra `SSTable` mới trên đĩa.
- Việc ghi đè liên tục các bản ghi trùng lặp ép Cassandra phải chạy tiến trình dọn dẹp (**Compaction**) nặng nề, gây nghẽn I/O đĩa và tăng độ trễ đột biến (Latency Spikes).

---

### C. Khóa phân tán nguyên tử siêu tốc (< 1ms)
- Redis hoạt động đơn luồng (Single-threaded event loop), câu lệnh:
  ```redis
  SET idempotency:{client_msg_id} "PROCESSING" NX EX 60
  ```
- Diễn ra trong **dưới 1 mili-giây**, đóng vai trò như chiếc "khiên bảo vệ" chặn đứng mọi yêu cầu trùng lặp ngay tại cửa ngõ trước khi làm tốn tài nguyên Database và Network.

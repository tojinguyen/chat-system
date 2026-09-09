# 3. Idempotency State Machine, Lock Rollback & Two-Phase TTL

---

## 1. Vấn đề "Kẹt Khóa Vĩnh Viễn" (Permanent Lockout) khi DB Gặp Lỗi

### Tình huống:
1. Worker nhận tin nhắn `client_msg_id = uuid_123`.
2. Worker chiếm khóa trên Redis: `SET idempotency:msg:uuid_123 "PROCESSED" NX EX 86400` (TTL 24h).
3. Ghi vào Cassandra: **BỊ LỖI** (Database timeout, quá tải hoặc network partition).
4. Hàm trả về lỗi `return err`. Client A sau 5 giây không thấy ACK nên gửi lại (Retry).

### Hậu quả:
- Khi Client A retry, bước `AcquireLock` trên Redis thấy key `uuid_123` đã tồn tại ➔ **Từ chối xử lý và báo trùng!**
- Tin nhắn bị bỏ qua mãi mãi trong 24h ➔ **Mất dữ liệu của người dùng (Data Loss)!**

---

## 2. Các cấp độ giải pháp xử lý Rollback Lock

### 🔹 Cấp độ 1: Chủ động Xóa Khóa (Active Release/Delete Lock)
Trong block `if err != nil` khi gọi `SaveMessage`, Worker lập tức gọi:
```go
u.idempotencyRepo.ReleaseLock(ctx, event.ClientMsgID) // Redis DEL key
```
- **Ưu điểm:** Đơn giản, giải phóng khóa ngay lập tức để lần retry kế tiếp có thể chạy lại ngay.
- **Điểm yếu:** Nếu Worker bị sập nguồn đột ngột (Crash / OOM) hoặc mạng sang Redis cũng bị đứt, lệnh `DEL` không thể thực thi.

---

### 🔹 Cấp độ 2 (Chuẩn Senior Production): Two-Phase State với Short-TTL Lock

Để giải quyết bài toán *"Worker sập nguồn trước khi kịp xóa lock"*, các hệ thống lớn áp dụng mô hình **Two-Phase Idempotency State**:

```
[Nhận Inbound Event]
        │
        ▼
BƯỚC 1: SET key "PROCESSING" NX EX 15s  (Khóa tạm thời với TTL cực ngắn)
        │
        ├──► Ghi vào Cassandra ──(THẤT BẠI)──► Không cần làm gì! (Khóa tự hủy sau 15s)
        │
        └──► Ghi vào Cassandra ──(THÀNH CÔNG)
                    │
                    ▼
BƯỚC 2: SET key "COMPLETED" XX EX 86400s (Gia hạn thành khóa vĩnh viễn 24h)
```

### Tại sao giải pháp Two-Phase TTL là tối thượng?
1. **Tự phục hồi mà không cần dọn dẹp (Self-Healing):** Dù Server có nổ tung hay đứt cáp, trạng thái `PROCESSING` sẽ tự động bốc hơi sau 15 giây nhờ cơ chế Expire nội tại của Redis.
2. **Không bao giờ xảy ra Deadlock:** Lần retry tiếp theo của Client (sau 15s) sẽ tự động chiếm được khóa mới và ghi lại bình thường.
3. **Bảo vệ toàn vẹn tuyệt đối:** Một khi đã chuyển sang `COMPLETED`, không một yêu cầu retry nào có thể ghi trùng vào Database được nữa.

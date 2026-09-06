# Kiến Trúc Hệ Thống & Tư Duy Phân Tán (System Architecture & Distributed Thinking)

Thư mục này lưu trữ các bài học, phân tích chuyên sâu về các vấn đề hóc búa, edge cases và giải pháp thực tế trong quá trình xây dựng hệ thống Chat phân tán quy mô lớn.

---

### Danh mục bài học:

1. [01. Cassandra UPSERT vs RDBMS & Bản chất Tầng Idempotency (Redis)](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/01_cassandra_upsert_vs_rdbms_and_idempotency.md)
   - *Nội dung:* Phân tích cơ chế xử lý trùng Primary Key của Cassandra (LSM-Tree), tại sao Cassandra không báo lỗi duplicate key, và 3 lý do sống còn cần tầng Redis Idempotency.

2. [02. Dual-Write Partial Failure, Poison Pill & State Reconciliation](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/02_dual_write_partial_failure_and_state_reconciliation.md)
   - *Nội dung:* Bài toán thất bại một phần khi lưu DB và gửi ACK, phân tích trade-off giữa `return err` vs `return nil`, hiện tượng Phantom Message và cơ chế State Reconciliation 2 chiều.

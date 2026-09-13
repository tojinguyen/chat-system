# Danh Mục Edge Cases & Điểm Quan Sát Hệ Thống (Edge Cases & Observability Backlog)

Thư mục này dùng để lưu trữ và theo dõi các góc tối kỹ thuật (Edge cases), rủi ro tiềm ẩn, và các sự cố phân tán được phát hiện trong quá trình code để tiện theo dõi, giám sát qua metric và tối ưu hóa sau này.

---

### Danh sách tài liệu:
1. [Edge Cases, Failure Modes & Observability Backlog](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/edge_cases/edge_cases_and_observability.md)
   - *Phantom ACK / False Success* (Race condition Redis Idempotency vs Cassandra Failure).
   - *Cassandra Zombie Write* (Context Timeout vs Mutation Commit).
   - *Ghost Online Drop* (Connection drop ngay sau khi query Redis Presence).

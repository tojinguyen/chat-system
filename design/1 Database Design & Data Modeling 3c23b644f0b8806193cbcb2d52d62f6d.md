# 1. Database Design & Data Modeling

1. Database Selection Strategy:

Chọn giữa Postgres (Relational), Columnar NoSQL, MongoDB.

Yêu cầu là tin nhắn được lưu trữ vĩnh viễn.

Số lượng tin nhắn chỉ có thể tăng lên và không giảm đi.

Loại bỏ SQL vì rất khó scale ngang, và khi dữ liệu càng lớn, tính chất của B-Tree Index làm suy giảm tốc độ ghi rất nhiều. Vẫn có thể sử dụng để lưu meta data như profile, conversation, friend, …

Loại bỏ DocumentDB như MongoDB bởi vì nó cũng dùng B-Tree index tương tự như Postgres, cho nên khi lượng tin nhắn đạt một ngưỡng nhất định, việc ghi sẽ trở nên ỳ ạch. 

MongoDB cũng duy trì toàn bộ B-Tree Index trong RAM, vì vậy phải tốn một lượng lớn ngân sách cho RAM để lưu được index khổng lồ của nó. 
Thêm vào đó là mỗi document của nó giới hạn 16MB, mỗi lần vượt quá 16MB thì chúng ta phải tự xử lý dữ liệu (có thể là áp dụng pattern bucket gì đó).

Chọn Cassandra vì nó là LSM-Tree Engine. Nó cho phép append dữ liệu vào cuối file log. Điều này giúp đẩy tốc độ ghi dữ liệu lên tối đa. Việc này cực kì thích hợp cho hệ thống nhắn tin phải ghi dữ liệu liên tục. 

Có thể dùng channel_id hoặc conversation_id để làm partition key.

Và message_id (Snowflake) làm clustering key, giúp message được tự động sắp xếp theo thứ tự ngay trên đĩa. 

Và đặc biệt là Cassandra không có Single Point of Failure, tự động horizontal Scaling.

1. Model Design:

```markdown
Users 
- user_id
- user_name
- password_hash
- name
- phone
- address
- created_at
- updated_at
- deleted_at
```

```markdown
Enum friend_status
- pending
- accepted
- declined
- blocked

Friends
- id
- requester_id
- addressee_id 
- status
- created_at
- updated_at 
```

```markdown
Conversations
- id
- name
- icon_url
- members
- created_at
- updated_at
```

```markdown
Messages 
- conversation_id
- id (uuidv7)
- sender_id
- content
- media_url
- created_at
- updated_at
```

```markdown
Conversation Read Receipts
- conversation_id
- user_id
- last_message_seen
- last_read_at
```
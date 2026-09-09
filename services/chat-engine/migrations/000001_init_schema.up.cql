CREATE TABLE IF NOT EXISTS messages (
    conversation_id text,
    id text,
    sender_id text,
    content text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY ((conversation_id), id)
) WITH CLUSTERING ORDER BY (id ASC);

CREATE TABLE IF NOT EXISTS conversation_read_receipts (
    conversation_id text,
    user_id text,
    last_message_seen text,
    last_read_at timestamp,
    PRIMARY KEY ((conversation_id), user_id)
);

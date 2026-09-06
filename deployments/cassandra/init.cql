-- Create Keyspace for chat system
CREATE KEYSPACE IF NOT EXISTS chat_system 
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

-- Create Messages Table
CREATE TABLE IF NOT EXISTS chat_system.messages (
    conversation_id text,
    id text,
    sender_id text,
    content text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY ((conversation_id), id)
) WITH CLUSTERING ORDER BY (id ASC);

-- Create Conversation Read Receipts Table
CREATE TABLE IF NOT EXISTS chat_system.conversation_read_receipts (
    conversation_id text,
    user_id text,
    last_message_seen text,
    last_read_at timestamp,
    PRIMARY KEY ((conversation_id), user_id)
);


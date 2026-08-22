# Part 1: Requirements, Constraints, and Capacity Estimation

1. Scope:
    
    The system is a real-time messaging platform similar to Messenger or Whatsapp.
    
    For the first version, we focus on the core messaging experience rather than trying to support every possible chat feature.
    
    In Scope:
    
    - One to one messaging
    - Group messaging
    - Real-time delivery
    - Message history
    - Online/offline presence
    - Delivery status
    - Read receipts
    - Multi-device support
    - Reconnection and missed-message syncronization
    
    Out of Scope:
    
    For the initial design, we will not focus on:
    
    - Voice calls
    - Video calls
    - Stories
    - End-to-end encryption
    - Full-text message search
    - Large media processing pipelines
    
    Attachments such as images or videos may be supported later through object storage, but they are not part of the core message delivery design.
    
2. Functional Requirements:
    1. FR1 - Send messages:
        - A user must be able to send messages to another user, or to other group.
        - A message contains at least:
            - message_id
            - conversation_id
            - sender_id
            - message_content
            - server_timestamp
    2. FR2 - Real-Time Message Delivery
        - If the recipient is online, the system should delivery the message in near real time.
        - The target user experience is: a message sent by User A should normally appear on User B’s device within a few hundred miniseconds.
    3. FR3 - Offline Messaging
        - If the recipient is offline, the message must not be lost.
        - The message should remain persisted in the system.
        - When the recipient reconnects, the client should synchronize messages that were missed while offline.
        - Should show NOTIFICATION when recipient is offline.
    4. FR4 - Message History
        - User must be able to retrieve previous messages from a conversation.
        - Can retrieve the most recent N messages from conversation X.
        - Older messages should be loaded using cursor-based pagination.
    5. FR5 - Message Ordering
        - Messages of a conversation must be handled by order.
        - Messages in unrelated conversations may be processed independently.
    6. FR6 - Duplicate Handling
        - Network failures may cause clients or infrastructure conponents to retry message delivery.
        - Therefore, the system must tolerate duplicate message submissions.
        - Each message should have a globally unique message_id
        - Processing should be idempotent so that retrying the same message does not create multiple persisted messages.
    7. FR7 - Group Chat
        - The system should support group conversations.
        - There are four type of conversation: global, region (VN, US, …), clan, directly.
    8. FR8 - Multi-Device Support.
        - A user may connect from multiple devices.
        - Messages should eventually syncronize accress all active devices belonging to the user.
        - A user may have multiple active connections.
    9. FR9 - Online Presence.
        - The system should expose approximate user presence: online, offline, lastseen.
        - Presence doesn’t require strong consistency.
        - Temporary inaccuracies are acceptable.
    10. FR10 - Delivery Status
        - Messages should support at least following status: sent, delivered, read.
        - System must save who read message.
    11. FR11 - Reconnection
        - WebSocket connections may disappear because of:
            - unstable mobile network
            - wifi changes
            - NAT timeout
            - gateway failure
            - application restart
            - device sleep
        - After reconnecting, the client should synchronize its state using persisted message history rather than assuming every real-time delivery succeed.
    
3. Non-functional Requirements:
    1. NFR1 - Low Latency
        - Real-time message delivery should have low latency
        - Target: P50 < 100ms and P99 < 500ms for online users located within the same region under normal operating conditions.
        - Latency across geographically distant regions maybe higher.
    2. NFR2 - High Availability
        - Sending and receiving messages are core functionality and should remain available despite individual machine failures.
        - Target availability: 99.99%
        - Individual failures such as:
            - gateway crash
            - chat worker crash
            - broker node crash
            - database replica failure
        - should not take down the entire messaging service
    3. NFR3 - Durability: 
        - Once the sender receives a successful **sent** acknowledgment, the message must not be lost because of a single server failure.
        - Message history therefore requires durable storage.
    4. NFR4 - Horizontal Scalability
        - The system should scale horizontally
        - We should able to increase capacity by adding:
            - Websocket Gateway
            - Chat Worker
            - broker partitions/nodes
            - database nodes
    5. NFR5 - Fault Tolerance
        - The system should tolerate partial failures
        - Example include:
            - Gateway alive - Chat Worker Slow
            - Broker alive - Database unavailable
            - Gateway alive - Redis unavailable
            - One data node unavailable - Other nodes healthy
        - The system should degrade gracefully instead of producing cascading failures
    6. NFR6 - Eventual Consistency Where Appropiate
        - Not every piece of chat state requires strong consistency.
        - Strong guarantees are important for:
            - durable messages
            - conversation membership authorization
            - message identity
        - Eventual consistency is acceptable for:
            - online presence
            - read receipts
            - delivery status
            - multi-device synchronization
        - The design should avoid paying the cost of strong consistency where the product doesn’t require it.
    7. NFR7 - Backpressure
        - The system must protect itself when downstream components become slower than incoming traffic.

1. Security Requirements
    - Users must authenticate before establishing a Websocket session.
    - The backend must also verify that a sender is authoried to send messages to the requested conversation.
    - Additional protections include:
        - TLS
        - maximum message size
        - per-user rate limiting
        - IP-based abuse protection
        - authentication during connection establishment

1. Initial Scale Assumptions
    - For the purpose of this design, assume:
        - Registered users: 100 million
        - Daily active users: 20 million
        - Peak concurrent users: 5 million
        - Assume each active user sends: 40 messages/day
        - Therefore: 20M x 40 = 800M messages/day
        - Average message rate: 8000 messages/sec
        - Peak traffic = 5 x average = 40,000 messages/sec
        - Design target: 50,000 inbound messages/sec

1. Message Size Assumption
    - Assume an average persisted message record of approximately: 500 bytes
    - including: id, timestamp, metadata, text payload, storage overhead estimate
    - Daily raw message storage: 800M * 500 bytes = 400GB/day
    - Per year: 400GB x 365 = 146TB/year
    - This number excludes:
        - replication
        - indexes
        - backups
        - metadata
        - attachments

1. Websocket Connection Scale
    - Peak concurrent connections: 5 milion
    - Assume during early capacity planning that one Gateway safely handles: 50,000 concurrent WebSocket connections
    - Then the theoretical minimum is: 100 Gateway instances
    - Production deployment needs additional headroom for:
        - machine failures
        - rolling deployments
        - connection spikes
        - uneven connection distribution
    - Therefore, the actual fleet should be larger than the theoretical minimum.
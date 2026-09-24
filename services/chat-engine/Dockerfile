# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy các file go.mod cần thiết
COPY pkg/go.mod pkg/go.sum* ./pkg/
COPY services/chat-engine/go.mod services/chat-engine/go.sum* ./services/chat-engine/

# Khởi tạo workspace cục bộ cho chat-engine và pkg
RUN go work init ./pkg ./services/chat-engine

# Copy mã nguồn
COPY pkg/ ./pkg/
COPY services/chat-engine/ ./services/chat-engine/

# Build binary
WORKDIR /app/services/chat-engine
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /chat-worker ./cmd/main.go

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /chat-worker .
COPY --from=builder /app/services/chat-engine/configs ./configs

CMD ["./chat-worker"]


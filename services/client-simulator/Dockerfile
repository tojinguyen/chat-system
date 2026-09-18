# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy các file go.mod cần thiết
COPY services/client-simulator/go.mod services/client-simulator/go.sum* ./services/client-simulator/

WORKDIR /app/services/client-simulator
RUN go mod download

# Copy toàn bộ mã nguồn của client-simulator
COPY services/client-simulator/ ./

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /client-simulator ./cmd/main.go

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /client-simulator .

ENTRYPOINT ["./client-simulator"]
CMD ["-api=http://nginx-gateway.chat-system.svc.cluster.local", "-ws=ws://nginx-gateway.chat-system.svc.cluster.local/ws", "-bots=30", "-convs=5", "-interval=500ms"]

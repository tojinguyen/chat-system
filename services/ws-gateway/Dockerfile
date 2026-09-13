# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy các file go.mod cần thiết
COPY pkg/go.mod pkg/go.sum* ./pkg/
COPY services/ws-gateway/go.mod services/ws-gateway/go.sum* ./services/ws-gateway/

# Khởi tạo workspace cục bộ cho riêng ws-gateway và pkg
RUN go work init ./pkg ./services/ws-gateway

# Copy mã nguồn
COPY pkg/ ./pkg/
COPY services/ws-gateway/ ./services/ws-gateway/

# Build binary
WORKDIR /app/services/ws-gateway
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /ws-gateway ./cmd/main.go

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /ws-gateway .
COPY --from=builder /app/services/ws-gateway/configs ./configs

EXPOSE 8080
CMD ["./ws-gateway"]

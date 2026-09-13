# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy các file go.mod cần thiết
COPY pkg/go.mod pkg/go.sum* ./pkg/
COPY services/notification-service/go.mod services/notification-service/go.sum* ./services/notification-service/

# Khởi tạo workspace cục bộ cho notification-service và pkg
RUN go work init ./pkg ./services/notification-service

# Copy mã nguồn
COPY pkg/ ./pkg/
COPY services/notification-service/ ./services/notification-service/

# Build binary
WORKDIR /app/services/notification-service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /notification-service ./cmd/main.go

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /notification-service .
COPY --from=builder /app/services/notification-service/configs ./configs

CMD ["./notification-service"]


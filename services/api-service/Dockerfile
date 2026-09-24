# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary tối ưu kích thước
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /api-service ./cmd/main.go

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /api-service .

EXPOSE 3000
CMD ["./api-service"]
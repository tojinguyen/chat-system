# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download || true

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ws-gateway ./cmd/main.go

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /ws-gateway .
COPY --from=builder /app/configs ./configs

EXPOSE 8080
CMD ["./ws-gateway"]

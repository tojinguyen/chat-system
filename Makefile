.PHONY: infra-up infra-down run-gateway run-worker run-noti help

help:
	@echo "Chat System Commands:"
	@echo "  make infra-up       - Khởi chạy Redis, ScyllaDB, Kafka"
	@echo "  make infra-down     - Dừng các container hạ tầng"
	@echo "  make run-gateway    - Chạy WebSocket Gateway"
	@echo "  make run-worker     - Chạy Chat Worker daemon"
	@echo "  make run-noti       - Chạy Notification Service"

infra-up:
	docker compose -f deployments/docker-compose.infra.yml up -d

infra-down:
	docker compose -f deployments/docker-compose.infra.yml down

run-gateway:
	cd services/ws-gateway && go run cmd/main.go

run-worker:
	cd services/chat-worker && go run cmd/main.go

run-noti:
	cd services/notification-service && go run cmd/main.go

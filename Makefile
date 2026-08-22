.PHONY: help infra-up infra-down infra-logs app-up app-down app-build app-logs run-api run-gateway run-worker run-noti migrate-run migrate-revert all-up all-down

# ==========================================
# 1. HELP & HƯỚNG DẪN
# ==========================================
help:
	@echo "================================================================"
	@echo "                    CHAT SYSTEM COMMANDS                        "
	@echo "================================================================"
	@echo "1. HẠ TẦNG (INFRASTRUCTURE):"
	@echo "  make infra-up        - Bật Postgres, Redis, ScyllaDB, NATS"
	@echo "  make infra-down      - Dừng toàn bộ hạ tầng"
	@echo "  make infra-logs      - Xem logs của hạ tầng"
	@echo ""
	@echo "2. TOÀN BỘ ỨNG DỤNG (DOCKER COMPOSE):"
	@echo "  make app-up          - Khởi chạy các service trong Docker"
	@echo "  make app-build       - Build lại image và khởi chạy"
	@echo "  make app-down        - Dừng các service ứng dụng"
	@echo "  make app-logs        - Xem logs tất cả app"
	@echo ""
	@echo "3. CHẠY LOCAL DEV (TRỰC TIẾP TRÊN MÁY):"
	@echo "  make run-api         - Chạy NestJS API Service (Watch mode)"
	@echo "  make run-gateway     - Chạy WebSocket Gateway (Go)"
	@echo "  make run-worker      - Chạy Chat Worker (Go)"
	@echo "  make run-noti        - Chạy Notification Service (Go)"
	@echo ""
	@echo "4. DATABASE & MIGRATION (API-SERVICE):"
	@echo "  make migrate-run     - Chạy TypeORM migrations"
	@echo "  make migrate-revert  - Rollback migration gần nhất"
	@echo ""
	@echo "5. TỔNG HỢP:"
	@echo "  make all-up          - Bật cả Hạ tầng + Ứng dụng"
	@echo "  make all-down        - Dừng tất cả"
	@echo "================================================================"

# ==========================================
# 2. HẠ TẦNG (INFRASTRUCTURE)
# ==========================================
infra-up:
	docker compose -f deployments/docker-compose.infra.yml up -d

infra-down:
	docker compose -f deployments/docker-compose.infra.yml down

infra-logs:
	docker compose -f deployments/docker-compose.infra.yml logs -f

# ==========================================
# 3. DOCKER COMPOSE APP
# ==========================================
app-up:
	docker compose -f deployments/docker-compose.yml up -d

app-build:
	docker compose -f deployments/docker-compose.yml up -d --build

app-down:
	docker compose -f deployments/docker-compose.yml down

app-logs:
	docker compose -f deployments/docker-compose.yml logs -f

# ==========================================
# 4. CHẠY LOCAL DEV
# ==========================================
run-api:
	cd services/api-service && npm run start:dev

run-gateway:
	cd services/ws-gateway && go run cmd/main.go

run-worker:
	cd services/chat-worker && go run cmd/main.go

run-noti:
	cd services/notification-service && go run cmd/main.go

# ==========================================
# 5. DATABASE MIGRATIONS
# ==========================================
migrate-run:
	cd services/api-service && npm run migration:run

migrate-revert:
	cd services/api-service && npm run migration:revert

# ==========================================
# 6. ALL-IN-ONE
# ==========================================
all-up: infra-up app-up
all-down: app-down infra-down

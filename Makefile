.PHONY: help infra-up infra-down infra-logs app-up app-down app-build app-logs run-api run-gateway run-worker run-noti migrate-run migrate-revert all-up all-down

# ==========================================
# 1. HELP & USAGE
# ==========================================
help:
	@echo "================================================================"
	@echo "                    CHAT SYSTEM COMMANDS                        "
	@echo "================================================================"
	@echo "1. INFRASTRUCTURE:"
	@echo "  make infra-up        - Start Postgres, Redis, ScyllaDB, NATS"
	@echo "  make infra-down      - Stop all infrastructure services"
	@echo "  make infra-logs      - View infrastructure logs"
	@echo ""
	@echo "2. FULL APPLICATION (DOCKER COMPOSE):"
	@echo "  make app-up          - Start application services in Docker"
	@echo "  make app-build       - Rebuild images and start services"
	@echo "  make app-down        - Stop application services"
	@echo "  make app-logs        - View logs of all applications"
	@echo ""
	@echo "3. LOCAL DEVELOPMENT (RUN DIRECTLY ON HOST):"
	@echo "  make run-api         - Run NestJS API Service (Watch mode)"
	@echo "  make run-gateway     - Run WebSocket Gateway (Go)"
	@echo "  make run-worker      - Run Chat Worker (Go)"
	@echo "  make run-noti        - Run Notification Service (Go)"
	@echo ""
	@echo "4. DATABASE & MIGRATIONS (API-SERVICE):"
	@echo "  make migrate-run     - Run TypeORM migrations"
	@echo "  make migrate-revert  - Revert the most recent migration"
	@echo ""
	@echo "5. ALL-IN-ONE:"
	@echo "  make all-up          - Start Infrastructure + Application"
	@echo "  make all-down        - Stop everything"
	@echo "================================================================"

# ==========================================
# 2. INFRASTRUCTURE
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
# 4. LOCAL DEVELOPMENT
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

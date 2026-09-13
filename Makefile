CLUSTER ?= chat-cluster

# --- Infrastructure (Host) ---
infra-up:
	docker compose -f deployments/docker-compose.infra.yml up -d

infra-down:
	docker compose -f deployments/docker-compose.infra.yml down

infra-logs:
	docker compose -f deployments/docker-compose.infra.yml logs -f

# --- Kind Cluster ---
kind-up:
	kind create cluster --name $(CLUSTER) --config deployments/kind-config.yaml

kind-down:
	kind delete cluster --name $(CLUSTER)

# --- Build & Load Images to Kind ---
build:
	docker build -t chat-system/ws-gateway:latest -f services/ws-gateway/Dockerfile .
	docker build -t chat-system/chat-engine:latest -f services/chat-engine/Dockerfile .
	docker build -t chat-system/api-service:latest services/api-service

load:
	kind load docker-image chat-system/ws-gateway:latest --name $(CLUSTER)
	kind load docker-image chat-system/chat-engine:latest --name $(CLUSTER)
	kind load docker-image chat-system/api-service:latest --name $(CLUSTER)

build-load: build load

# --- Kubernetes Deploy & Ops ---
k8s-deploy:
	kubectl apply -f deployments/k8s/

k8s-delete:
	kubectl delete -f deployments/k8s/

k8s-restart:
	kubectl rollout restart statefulset/ws-gateway -n chat-system
	kubectl rollout restart deployment/chat-engine -n chat-system
	kubectl rollout restart deployment/api-service -n chat-system
	kubectl rollout restart deployment/nginx-gateway -n chat-system

k8s-status:
	kubectl get pods,svc,statefulset -n chat-system -o wide

# --- K8s Logs ---
logs-gw:
	kubectl logs -l app=ws-gateway -n chat-system -f

logs-engine:
	kubectl logs -l app=chat-engine -n chat-system -f

logs-api:
	kubectl logs -l app=api-service -n chat-system -f

# --- Client Simulator ---
sim:
	cd services/client-simulator && go run cmd/main.go -bots 10 -convs 3 -interval 1s

sim-stress:
	cd services/client-simulator && go run cmd/main.go -bots 50 -convs 5 -interval 200ms


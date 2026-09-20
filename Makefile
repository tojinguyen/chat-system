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
	docker build -t chat-system/client-simulator:latest -f services/client-simulator/Dockerfile .

load:
	kind load docker-image chat-system/ws-gateway:latest --name $(CLUSTER)
	kind load docker-image chat-system/chat-engine:latest --name $(CLUSTER)
	kind load docker-image chat-system/api-service:latest --name $(CLUSTER)
	kind load docker-image chat-system/client-simulator:latest --name $(CLUSTER)

build-load: build load

# --- Build & Load Từng Service Riêng Biệt (Tối ưu tốc độ dev) ---
build-sim:
	docker build -t chat-system/client-simulator:latest -f services/client-simulator/Dockerfile .

load-sim:
	kind load docker-image chat-system/client-simulator:latest --name $(CLUSTER)

build-load-sim: build-sim load-sim
	kubectl rollout restart deployment/client-simulator -n chat-system --ignore-not-found

build-engine:
	docker build -t chat-system/chat-engine:latest -f services/chat-engine/Dockerfile .

load-engine:
	kind load docker-image chat-system/chat-engine:latest --name $(CLUSTER)

build-load-engine: build-engine load-engine
	kubectl rollout restart deployment/chat-engine -n chat-system

build-gw:
	docker build -t chat-system/ws-gateway:latest -f services/ws-gateway/Dockerfile .

load-gw:
	kind load docker-image chat-system/ws-gateway:latest --name $(CLUSTER)

build-load-gw: build-gw load-gw
	kubectl rollout restart statefulset/ws-gateway -n chat-system

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
	kubectl rollout restart deployment/client-simulator -n chat-system --ignore-not-found

k8s-config:
	kubectl apply -f deployments/k8s/01-configmap-secrets.yaml
	kubectl rollout restart deployment/chat-engine statefulset/ws-gateway deployment/nginx-gateway -n chat-system

reload: k8s-config

k8s-status:
	kubectl get pods,svc,statefulset -n chat-system -o wide

# --- K8s Logs ---
logs-gw:
	kubectl logs -l app=ws-gateway -n chat-system -f

logs-engine:
	kubectl logs -l app=chat-engine -n chat-system -f

logs-api:
	kubectl logs -l app=api-service -n chat-system -f

logs-sim:
	kubectl logs -l app=client-simulator -n chat-system -f

# --- Client Simulator ---
sim:
	cd services/client-simulator && go run cmd/main.go -bots 10 -convs 3 -interval 1s

sim-stress:
	cd services/client-simulator && go run cmd/main.go -bots 50 -convs 5 -interval 200ms

sim-k8s:
	kubectl apply -f deployments/k8s/09-client-simulator.yaml

sim-k8s-down:
	kubectl delete -f deployments/k8s/09-client-simulator.yaml --ignore-not-found

# --- Chaos Engineering ---
chaos-delay-grpc:
	kubectl apply -f deployments/k8s/chaos/01-network-delay-grpc.yaml

chaos-delay-nats:
	kubectl apply -f deployments/k8s/chaos/02-network-delay-nats.yaml

chaos-kill-gw:
	kubectl apply -f deployments/k8s/chaos/03-pod-kill-gateway.yaml

chaos-mobile-delay:
	kubectl apply -f deployments/k8s/chaos/04-mobile-network-delay.yaml

chaos-mobile-loss:
	kubectl apply -f deployments/k8s/chaos/05-mobile-packet-loss.yaml

chaos-mobile-partition:
	kubectl apply -f deployments/k8s/chaos/06-mobile-network-partition.yaml

chaos-clean:
	kubectl delete -f deployments/k8s/chaos/ --ignore-not-found

k6-stress:
	kubectl create configmap k6-test-script --from-file=test-ws-load.js=deployments/k6/test-ws-load.js -n chat-system --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -f deployments/k6/k6-testrun.yaml


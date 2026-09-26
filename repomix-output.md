This file is a merged representation of the entire codebase, combined into a single document by Repomix.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)

# Directory Structure
```
.dockerignore
.gitignore
.repomixignore
AGENTS.md
deployments/configs/grafana-datasources.yaml
deployments/configs/grafana/dashboards/chat-golden-signals.json
deployments/configs/grafana/provisioning/dashboards/dashboards.yaml
deployments/configs/loki.yaml
deployments/configs/otel-collector.yaml
deployments/configs/prometheus.yml
deployments/configs/promtail.yaml
deployments/configs/tempo.yaml
deployments/docker-compose.infra.yml
deployments/k6/k6-testrun.yaml
deployments/k6/test-ws-load.js
deployments/k8s/00-external-infra.yaml
deployments/k8s/01-configmap-secrets.yaml
deployments/k8s/02-ws-gateway-statefulset.yaml
deployments/k8s/03-ws-gateway-ingress-service.yaml
deployments/k8s/04-chat-engine-deployment.yaml
deployments/k8s/05-api-service-deployment.yaml
deployments/k8s/06-nginx-gateway.yaml
deployments/k8s/07-prometheus-agent.yaml
deployments/k8s/08-promtail.yaml
deployments/k8s/09-client-simulator.yaml
deployments/k8s/chaos/01-network-delay-grpc.yaml
deployments/k8s/chaos/02-network-delay-nats.yaml
deployments/k8s/chaos/03-pod-kill-gateway.yaml
deployments/k8s/chaos/04-mobile-network-delay.yaml
deployments/k8s/chaos/05-mobile-packet-loss.yaml
deployments/k8s/chaos/06-mobile-network-partition.yaml
deployments/kind-config.yaml
design/1 Database Design & Data Modeling 3c23b644f0b8806193cbcb2d52d62f6d.md
design/image.png
design/Part 1 Requirements, Constraints, and Capacity Est 3be3b644f0b88019bd11df04cdaf19fe.md
design/Part 2 High-Level Architecture & End-to-End Messag 3be3b644f0b8804aa7cdc7af6fe247cd.md
knowledge/edge_cases/edge_cases_and_observability.md
knowledge/edge_cases/README.md
knowledge/system_design/01_cassandra_upsert_vs_rdbms_and_idempotency.md
knowledge/system_design/02_dual_write_partial_failure_and_state_reconciliation.md
knowledge/system_design/03_idempotency_state_machine_and_ttl_rollback.md
knowledge/system_design/README.md
Makefile
pkg/contracts/events.go
pkg/contracts/topics.go
pkg/go.mod
pkg/grpcclient/gateway_client_manager.go
pkg/nats/client.go
pkg/nats/instrumented_publisher.go
pkg/nats/publisher.go
pkg/nats/subscriber.go
pkg/proto/ws_gateway_grpc.pb.go
pkg/proto/ws_gateway.pb.go
pkg/proto/ws_gateway.proto
pkg/telemetry/carrier.go
pkg/telemetry/grpc_interceptor.go
pkg/telemetry/metrics.go
pkg/telemetry/middleware.go
pkg/telemetry/profiler.go
pkg/telemetry/setup_test.go
pkg/telemetry/setup.go
pkg/telemetry/tracer.go
pkg/worker/pool_test.go
pkg/worker/pool.go
plan/3_week_observability_benchmark_stress_plan.md
plan/benchmark_runbook_grpc_vs_nats.md
README.md
repomix.config.json
services/api-service/.dockerignore
services/api-service/.env.example
services/api-service/.gitignore
services/api-service/cmd/main.go
services/api-service/Dockerfile
services/api-service/go.mod
services/api-service/internal/config/config.go
services/api-service/internal/delivery/http/auth_handler.go
services/api-service/internal/delivery/http/conversation_handler.go
services/api-service/internal/delivery/http/dto/dto.go
services/api-service/internal/delivery/http/friend_handler.go
services/api-service/internal/delivery/http/middleware/auth_middleware.go
services/api-service/internal/delivery/http/middleware/response.go
services/api-service/internal/delivery/http/router.go
services/api-service/internal/delivery/http/user_handler.go
services/api-service/internal/domain/conversation.go
services/api-service/internal/domain/friend.go
services/api-service/internal/domain/user.go
services/api-service/internal/repository/conversation_repository.go
services/api-service/internal/repository/friend_repository.go
services/api-service/internal/repository/user_repository.go
services/api-service/internal/usecase/auth_usecase.go
services/api-service/internal/usecase/conversation_usecase.go
services/api-service/internal/usecase/friend_usecase.go
services/api-service/internal/usecase/user_usecase.go
services/api-service/README.md
services/chat-engine/cmd/main.go
services/chat-engine/configs/config.yaml
services/chat-engine/Dockerfile
services/chat-engine/go.mod
services/chat-engine/internal/config/config.go
services/chat-engine/internal/consumer/consumer.go
services/chat-engine/internal/dispatcher/dispatcher_grpc.go
services/chat-engine/internal/dispatcher/dispatcher_nats.go
services/chat-engine/internal/dispatcher/dispatcher.go
services/chat-engine/internal/domain/message.go
services/chat-engine/internal/migrator/migrator.go
services/chat-engine/internal/presence/presence.go
services/chat-engine/internal/repository/idempotency_repo.go
services/chat-engine/internal/repository/message_repo.go
services/chat-engine/internal/usecase/chat_usecase.go
services/chat-engine/migrations/000001_create_messages_table.down.cql
services/chat-engine/migrations/000001_create_messages_table.up.cql
services/chat-engine/migrations/000002_create_read_receipts_table.down.cql
services/chat-engine/migrations/000002_create_read_receipts_table.up.cql
services/chat-engine/migrations/fs.go
services/client-simulator/cmd/main.go
services/client-simulator/configs/config.yaml
services/client-simulator/Dockerfile
services/client-simulator/go.mod
services/client-simulator/internal/auth/client.go
services/client-simulator/internal/bot/bot.go
services/client-simulator/internal/config/config.go
services/client-simulator/internal/conversation/manager.go
services/client-simulator/internal/metrics/tracker.go
services/notification-service/cmd/main.go
services/notification-service/configs/config.yaml
services/notification-service/Dockerfile
services/notification-service/go.mod
services/notification-service/internal/config/config.go
services/notification-service/internal/consumer/consumer.go
services/notification-service/internal/domain/notification.go
services/notification-service/internal/provider/apns.go
services/notification-service/internal/provider/fcm.go
services/notification-service/internal/provider/provider.go
services/notification-service/internal/usecase/notification_usecase.go
services/ws-gateway/cmd/main.go
services/ws-gateway/configs/config.yaml
services/ws-gateway/Dockerfile
services/ws-gateway/go.mod
services/ws-gateway/internal/config/config.go
services/ws-gateway/internal/connection/client.go
services/ws-gateway/internal/connection/hub.go
services/ws-gateway/internal/delivery/grpc_delivery.go
services/ws-gateway/internal/delivery/listener.go
services/ws-gateway/internal/delivery/nats_delivery.go
services/ws-gateway/internal/handler/ws_handler.go
services/ws-gateway/internal/payload/ws_payload.go
services/ws-gateway/internal/presence/presence.go
```

# Files

## File: .repomixignore
````
# Add patterns to ignore here, one per line
# Example:
# *.log
# tmp/
````

## File: repomix.config.json
````json
{
  "$schema": "https://repomix.com/schemas/latest/schema.json",
  "input": {
    "maxFileSize": 52428800
  },
  "output": {
    "filePath": "repomix-output.md",
    "style": "markdown",
    "parsableStyle": false,
    "fileSummary": true,
    "directoryStructure": true,
    "files": true,
    "removeComments": false,
    "removeEmptyLines": false,
    "compress": false,
    "topFilesLength": 5,
    "showLineNumbers": false,
    "truncateBase64": false,
    "copyToClipboard": false,
    "includeFullDirectoryStructure": false,
    "tokenCountTree": false,
    "git": {
      "sortByChanges": true,
      "sortByChangesMaxCommits": 100,
      "includeDiffs": false,
      "includeLogs": false,
      "includeLogsCount": 50
    }
  },
  "include": [],
  "ignore": {
    "useGitignore": true,
    "useDotIgnore": true,
    "useDefaultPatterns": true,
    "customPatterns": []
  },
  "security": {
    "enableSecurityCheck": true
  },
  "tokenCount": {
    "encoding": "o200k_base"
  }
}
````

## File: .dockerignore
````
.git
.vscode
services/api-service/node_modules
services/api-service/dist
tmp/
bin/
*.log
repomix-output.xml
````

## File: deployments/configs/grafana-datasources.yaml
````yaml
apiVersion: 1

datasources:
  - name: Prometheus
    uid: prometheus
    type: prometheus
    access: proxy
    url: http://prometheus:9090
    isDefault: true
    jsonData:
      timeInterval: 5s

  - name: Loki
    uid: loki
    type: loki
    access: proxy
    url: http://loki:3100
    jsonData:
      maxLines: 1000

  - name: Tempo
    uid: tempo
    type: tempo
    access: proxy
    url: http://tempo:3200
    jsonData:
      tracesToLogs:
        datasourceUid: loki
        tags: ['service.name', 'service']
        filterByTraceID: true
        filterBySpanID: true
      tracesToMetrics:
        datasourceUid: prometheus
        tags: [{ key: 'service.name', value: 'service' }]
        queries:
          - name: 'Throughput'
            query: 'sum(rate(chat_system_messages_total{service=~"$__tags"}[1m]))'
      tracesToProfiles:
        datasourceUid: pyroscope
        tags: [{ key: 'service.name', value: 'service' }]

  - name: Pyroscope
    uid: pyroscope
    type: grafana-pyroscope-datasource
    access: proxy
    url: http://pyroscope:4040
````

## File: deployments/configs/grafana/dashboards/chat-golden-signals.json
````json
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "grafana",
          "uid": "-- Grafana --"
        },
        "enable": true,
        "hide": true,
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 1,
  "id": null,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 0
      },
      "id": 100,
      "title": "📋 15-MINUTE BENCHMARK EXECUTIVE SUMMARY (gRPC vs NATS)",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "yellow",
                "value": 0.05
              },
              {
                "color": "red",
                "value": 0.1
              }
            ]
          },
          "unit": "s"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 4,
        "w": 4,
        "x": 0,
        "y": 1
      },
      "id": 10,
      "options": {
        "colorMode": "value",
        "graphMode": "area",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "title": "P99 Dispatch (gRPC)",
      "type": "stat",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum(rate(chat_engine_dispatch_duration_seconds_bucket{mode=\"grpc\"}[15m])) by (le))",
          "legendFormat": "gRPC P99",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "Thời gian điều phối (Dispatch) P99 từ Chat Engine sang WS Gateway qua gRPC. P99 < 50ms là an toàn, > 100ms cảnh báo nghẽn socket hoặc Jitter mạng."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "yellow",
                "value": 0.02
              },
              {
                "color": "red",
                "value": 0.05
              }
            ]
          },
          "unit": "s"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 4,
        "w": 4,
        "x": 4,
        "y": 1
      },
      "id": 11,
      "options": {
        "colorMode": "value",
        "graphMode": "area",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "title": "P99 Dispatch (NATS)",
      "type": "stat",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum(rate(chat_engine_dispatch_duration_seconds_bucket{mode=\"nats\"}[15m])) by (le))",
          "legendFormat": "NATS P99",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "Thời gian điều phối (Dispatch) P99 từ Chat Engine sang WS Gateway qua NATS Broker. Nhờ cơ chế bất đồng bộ, chỉ số này thường < 5ms."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 1
              }
            ]
          },
          "unit": "short"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 4,
        "w": 4,
        "x": 8,
        "y": 1
      },
      "id": 12,
      "options": {
        "colorMode": "value",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "title": "Total Errors (gRPC)",
      "type": "stat",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "sum(increase(chat_engine_dispatch_duration_seconds_count{mode=\"grpc\", status=\"error\"}[15m])) or vector(0)",
          "legendFormat": "gRPC Errors",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "Tổng số request gRPC dispatch thất bại (lỗi mạng, timeout hoặc Gateway từ chối)."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 1
              }
            ]
          },
          "unit": "short"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 4,
        "w": 4,
        "x": 12,
        "y": 1
      },
      "id": 13,
      "options": {
        "colorMode": "value",
        "graphMode": "none",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "title": "Total Errors (NATS)",
      "type": "stat",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "sum(increase(chat_engine_dispatch_duration_seconds_count{mode=\"nats\", status=\"error\"}[15m])) or vector(0)",
          "legendFormat": "NATS Errors",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "Tổng số lỗi khi publish tin nhắn vào NATS subject."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "mappings": [],
          "unit": "reqps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 4,
        "w": 4,
        "x": 16,
        "y": 1
      },
      "id": 14,
      "options": {
        "colorMode": "value",
        "graphMode": "area",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "max"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "title": "Peak Throughput (gRPC)",
      "type": "stat",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "max_over_time(sum(rate(chat_engine_dispatch_duration_seconds_count{mode=\"grpc\"}[1m]))[15m:10s]) or vector(0)",
          "legendFormat": "gRPC Peak RPS",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "Tốc độ xử lý đỉnh điểm (Messages/sec) ở chế độ gRPC."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "mappings": [],
          "unit": "reqps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 4,
        "w": 4,
        "x": 20,
        "y": 1
      },
      "id": 15,
      "options": {
        "colorMode": "value",
        "graphMode": "area",
        "justifyMode": "auto",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "max"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "title": "Peak Throughput (NATS)",
      "type": "stat",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "max_over_time(sum(rate(chat_engine_dispatch_duration_seconds_count{mode=\"nats\"}[1m]))[15m:10s]) or vector(0)",
          "legendFormat": "NATS Peak RPS",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "Tốc độ xử lý đỉnh điểm (Messages/sec) ở chế độ NATS Broker."
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 5
      },
      "id": 101,
      "title": "1. ⏱️ ĐỘ TRỄ ĐIỀU PHỐI & JITTER (DISPATCH LATENCY & JITTER)",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Latency",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "line"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "yellow",
                "value": 0.05
              },
              {
                "color": "red",
                "value": 0.1
              }
            ]
          },
          "unit": "s"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 6
      },
      "id": 20,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 1.1: Dispatch Latency Comparison (P50, P95, P99 by Mode)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum(rate(chat_engine_dispatch_duration_seconds_bucket{mode=~\"$mode\"}[1m])) by (le, mode))",
          "legendFormat": "P99 [{{mode}}]",
          "range": true,
          "refId": "P99"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.95, sum(rate(chat_engine_dispatch_duration_seconds_bucket{mode=~\"$mode\"}[1m])) by (le, mode))",
          "legendFormat": "P95 [{{mode}}]",
          "range": true,
          "refId": "P95"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.50, sum(rate(chat_engine_dispatch_duration_seconds_bucket{mode=~\"$mode\"}[1m])) by (le, mode))",
          "legendFormat": "P50 [{{mode}}]",
          "range": true,
          "refId": "P50"
        }
      ],
      "description": "DISPATCH LATENCY: Thời gian Chat Engine điều phối tin nhắn sang máy chủ Gateway nơi người nhận đang giữ kết nối WebSocket. Biểu đồ so sánh P50, P95, P99 giữa gRPC và NATS."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Latency",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "s"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 6
      },
      "id": 21,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 1.2: End-to-End Delivery & Gateway Push Latency (P99)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum(rate(chat_system_message_latency_seconds_bucket{stage=\"gateway_delivery\", mode=~\"$mode\"}[1m])) by (le, mode))",
          "legendFormat": "P99 Gateway Push [{{mode}}]",
          "range": true,
          "refId": "GatewayPush"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum(rate(chat_system_message_latency_seconds_bucket{stage=\"worker_dispatch\"}[1m])) by (le))",
          "legendFormat": "P99 Worker Inbound Queue",
          "range": true,
          "refId": "InboundQueue"
        }
      ],
      "description": "END-TO-END (E2E): Tổng thời gian tin nhắn đi từ Inbound -> Lưu ScyllaDB -> Dispatch -> Gateway đẩy xuống Socket Client. Gateway Push là thời gian Gateway nhận tin và ghi vào socket người nhận."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Latency",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "line"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "yellow",
                "value": 0.05
              },
              {
                "color": "red",
                "value": 0.1
              }
            ]
          },
          "unit": "s"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 14
      },
      "id": 22,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 1.3: Client Network Delay / Edge RTT (P50, P95, P99)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum(rate(chat_gateway_client_network_rtt_seconds_bucket[1m])) by (le))",
          "legendFormat": "P99 Edge Network RTT",
          "range": true,
          "refId": "P99"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.95, sum(rate(chat_gateway_client_network_rtt_seconds_bucket[1m])) by (le))",
          "legendFormat": "P95 Edge Network RTT",
          "range": true,
          "refId": "P95"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.50, sum(rate(chat_gateway_client_network_rtt_seconds_bucket[1m])) by (le))",
          "legendFormat": "P50 (Median) Edge Network RTT",
          "range": true,
          "refId": "P50"
        }
      ],
      "description": "EDGE NETWORK DELAY: Đo độ trễ mạng hai chiều (RTT) giữa Client và WS Gateway qua WebSocket Ping-Pong/Heartbeat. Giúp phân định ranh giới giữa chậm do mạng (Mobile/WAN Jitter) hay chậm do xử lý Backend (Database/Worker)."
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 22
      },
      "id": 102,
      "title": "2. 🚀 THÔNG LƯỢNG & ĐỘ LỆCH TẢI (THROUGHPUT & BALANCING)",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Msgs / Sec",
            "axisPlacement": "auto",
            "fillOpacity": 20,
            "gradientMode": "opacity",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "normal"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "reqps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 23
      },
      "id": 30,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 2.1: Dispatch Throughput (Messages/sec by Mode)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "sum(rate(chat_engine_dispatch_duration_seconds_count{mode=~\"$mode\"}[1m])) by (mode)",
          "legendFormat": "Dispatch Rate [{{mode}}]",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "THROUGHPUT: Thông lượng xử lý tin nhắn thực tế (tin nhắn/giây) của hệ thống theo chế độ gRPC vs NATS."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Msgs / Sec",
            "axisPlacement": "auto",
            "fillOpacity": 20,
            "gradientMode": "opacity",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "normal"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "reqps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 23
      },
      "id": 31,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 2.2: Load Balancing Across Gateway Pods (Hotspot Detection)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "sum(rate(chat_engine_dispatch_duration_seconds_count{mode=~\"$mode\", node_id=~\"$node_id\"}[1m])) by (node_id, mode)",
          "legendFormat": "{{node_id}} [{{mode}}]",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "LOAD BALANCING: Phân bổ tin nhắn tới từng Pod Gateway (0, 1, 2). Dùng để phát hiện Hotspot Pod khi một node gánh quá nhiều kết nối."
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 31
      },
      "id": 103,
      "title": "3. ⚠️ TỶ LỆ LỖI & ĐỘ TIN CẬY (ERROR RATES & REJECTIONS)",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Error Rate (%)",
            "axisPlacement": "auto",
            "fillOpacity": 25,
            "gradientMode": "opacity",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "line"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "yellow",
                "value": 1
              },
              {
                "color": "red",
                "value": 5
              }
            ]
          },
          "unit": "percent"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 32
      },
      "id": 40,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 3.1: Dispatch Error Rate (%) by Mode",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "(sum(rate(chat_engine_dispatch_duration_seconds_count{status=\"error\", mode=~\"$mode\"}[1m])) by (mode) / sum(rate(chat_engine_dispatch_duration_seconds_count{mode=~\"$mode\"}[1m])) by (mode)) * 100",
          "legendFormat": "Error Rate % [{{mode}}]",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "DISPATCH ERROR RATE: Tỷ lệ lỗi (%) trên tổng số tin nhắn điều phối. Bình thường = 0%. Nếu > 1% là cảnh báo nguy hiểm."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Rejections / Sec",
            "axisPlacement": "auto",
            "fillOpacity": 20,
            "gradientMode": "opacity",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "reqps"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 32
      },
      "id": 41,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 3.2: Rejection Rate (gRPC Rejected by Gateway)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "sum(rate(chat_engine_dispatch_duration_seconds_count{mode=\"grpc\", status=\"rejected\"}[1m])) by (node_id)",
          "legendFormat": "gRPC Rejected by {{node_id}}",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "REJECTION RATE: Số tin nhắn bị Gateway từ chối do quá tải bộ đệm Outbound (chỉ xảy ra ở gRPC khi client đọc chậm, NATS không bị vì có Broker đệm)."
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 40
      },
      "id": 104,
      "title": "4. 🧱 TẮC NGHẼN HÀNG ĐỢI & BACKPRESSURE (SATURATION)",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Saturation (%)",
            "axisPlacement": "auto",
            "fillOpacity": 20,
            "gradientMode": "opacity",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "line"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "yellow",
                "value": 70
              },
              {
                "color": "red",
                "value": 90
              }
            ]
          },
          "unit": "percent"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 41
      },
      "id": 50,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 4.1: Worker Pool Channel Saturation (%)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "max(chat_worker_channel_saturation_ratio) * 100",
          "legendFormat": "Max Worker Channel Saturation %",
          "range": true,
          "refId": "Max"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "avg(chat_worker_channel_saturation_ratio) * 100",
          "legendFormat": "Avg Worker Channel Saturation %",
          "range": true,
          "refId": "Avg"
        }
      ],
      "description": "WORKER SATURATION: Độ bão hòa (%) của Go channel trong PartitionedPool của Chat Engine. 100% nghĩa là hàng đợi đã kịch trần, hệ thống bị Backpressure."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Backlog Jobs",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "short"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 41
      },
      "id": 51,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 4.2: Total Pending Backlog Jobs in Pool",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "sum(chat_worker_channel_depth)",
          "legendFormat": "Total Backlog in Worker Pool",
          "range": true,
          "refId": "TotalDepth"
        }
      ],
      "description": "PENDING BACKLOG: Tổng số job tin nhắn đang chờ xử lý trong Worker Pool. Tăng vọt khi ScyllaDB chậm hoặc luồng Dispatch bị nghẽn."
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 49
      },
      "id": 105,
      "title": "5. 💻 TIÊU HAO TÀI NGUYÊN HỆ THỐNG (SYSTEM RESOURCES)",
      "type": "row"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Goroutines Count",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "short"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 8,
        "x": 0,
        "y": 50
      },
      "id": 60,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 5.1: Goroutine Explosion Detection",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "go_goroutines",
          "legendFormat": "{{instance}} - {{job}}",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "GOROUTINES: Số lượng Goroutine chạy trong từng Pod. Khi gRPC bị trễ mạng, Goroutines sẽ bùng nổ do bị block chờ I/O; NATS không bị vì là fire-and-forget."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "Memory (MB)",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "bytes"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 8,
        "x": 8,
        "y": 50
      },
      "id": 61,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 5.2: Memory Working Set (Allocated Heap)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "go_memstats_alloc_bytes",
          "legendFormat": "{{instance}} - {{job}}",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "HEAP ALLOC: Dung lượng RAM Heap đang được cấp phát bởi Go runtime. Giúp phát hiện Memory Leak hoặc buffer bị phình to."
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "CPU (%)",
            "axisPlacement": "auto",
            "fillOpacity": 15,
            "gradientMode": "none",
            "lineWidth": 2,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "unit": "percent"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 8,
        "x": 16,
        "y": 50
      },
      "id": 62,
      "options": {
        "legend": {
          "calcs": [
            "mean",
            "lastNotNull",
            "max"
          ],
          "displayMode": "table",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "multi",
          "sort": "desc"
        }
      },
      "title": "Panel 5.3: Process CPU Usage (%)",
      "type": "timeseries",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "editorMode": "code",
          "expr": "rate(process_cpu_seconds_total[1m]) * 100",
          "legendFormat": "{{instance}} - {{job}}",
          "range": true,
          "refId": "A"
        }
      ],
      "description": "CPU USAGE: Tỷ lệ CPU tiêu thụ (%) của từng Pod Chat Engine và Gateway."
    }
  ],
  "refresh": "5s",
  "schemaVersion": 39,
  "tags": [
    "chat-system",
    "sre",
    "golden-signals",
    "grpc-vs-nats",
    "benchmarking"
  ],
  "templating": {
    "list": [
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": "All",
          "value": "$__all"
        },
        "hide": 0,
        "includeAll": true,
        "label": "Delivery Mode",
        "multi": true,
        "name": "mode",
        "options": [
          {
            "selected": true,
            "text": "All",
            "value": "$__all"
          },
          {
            "selected": false,
            "text": "grpc",
            "value": "grpc"
          },
          {
            "selected": false,
            "text": "nats",
            "value": "nats"
          }
        ],
        "query": "grpc, nats",
        "queryValue": "",
        "skipUrlSync": false,
        "type": "custom"
      },
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": "All",
          "value": "$__all"
        },
        "hide": 0,
        "includeAll": true,
        "label": "Gateway Node",
        "multi": true,
        "name": "node_id",
        "options": [
          {
            "selected": true,
            "text": "All",
            "value": "$__all"
          },
          {
            "selected": false,
            "text": "ws-gateway-0",
            "value": "ws-gateway-0"
          },
          {
            "selected": false,
            "text": "ws-gateway-1",
            "value": "ws-gateway-1"
          },
          {
            "selected": false,
            "text": "ws-gateway-2",
            "value": "ws-gateway-2"
          }
        ],
        "query": "ws-gateway-0, ws-gateway-1, ws-gateway-2",
        "queryValue": "",
        "skipUrlSync": false,
        "type": "custom"
      }
    ]
  },
  "time": {
    "from": "now-15m",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m"
    ]
  },
  "timezone": "browser",
  "title": "Chat System - 5 Golden Dimensions (gRPC vs NATS)",
  "uid": "chat-golden-signals",
  "version": 2,
  "weekStart": ""
}
````

## File: deployments/configs/grafana/provisioning/dashboards/dashboards.yaml
````yaml
apiVersion: 1

providers:
  - name: 'Chat System Dashboards'
    orgId: 1
    folder: 'Chat System'
    type: file
    disableDeletion: false
    updateIntervalSeconds: 10
    allowUiUpdates: true
    options:
      path: /var/lib/grafana/dashboards
      foldersFromFilesStructure: false
````

## File: deployments/configs/loki.yaml
````yaml
auth_enabled: false

server:
  http_listen_port: 3100
  grpc_listen_port: 9096

common:
  instance_addr: 127.0.0.1
  path_prefix: /var/loki
  storage:
    filesystem:
      chunks_directory: /var/loki/chunks
      rules_directory: /var/loki/rules
  replication_factor: 1
  ring:
    kvstore:
      store: inmemory

schema_config:
  configs:
    - from: 2020-10-24
      store: tsdb
      object_store: filesystem
      schema: v13
      index:
        prefix: index_
        period: 24h

ruler:
  alertmanager_url: http://localhost:9093
````

## File: deployments/configs/otel-collector.yaml
````yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318

processors:
  memory_limiter:
    check_interval: 1s
    limit_percentage: 75
    spike_limit_percentage: 20
  batch:
    send_batch_size: 1024
    timeout: 1s
    send_batch_max_size: 2048

exporters:
  otlp/tempo:
    endpoint: tempo:4317
    tls:
      insecure: true
  prometheusremotewrite/prometheus:
    endpoint: http://prometheus:9090/api/v1/write
    tls:
      insecure: true
  otlphttp/loki:
    endpoint: http://loki:3100/otlp
    tls:
      insecure: true

service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [otlp/tempo]
    metrics:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [prometheusremotewrite/prometheus]
    logs:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [otlphttp/loki]
````

## File: deployments/configs/prometheus.yml
````yaml
global:
  scrape_interval: 5s
  evaluation_interval: 5s

scrape_configs:
  - job_name: "prometheus"
    static_configs:
      - targets: ["localhost:9090"]

  - job_name: "otel-collector"
    static_configs:
      - targets: ["otel-collector:8889"]
````

## File: deployments/configs/promtail.yaml
````yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://loki:3100/loki/api/v1/push

scrape_configs:
  - job_name: docker
    docker_sd_configs:
      - host: unix:///var/run/docker.sock
        refresh_interval: 5s
    relabel_configs:
      - source_labels: ['__meta_docker_container_name']
        regex: '/(.*)'
        target_label: 'container_name'
      - source_labels: ['__meta_docker_container_label_com_docker_compose_service']
        target_label: 'compose_service'
      - source_labels: ['__meta_docker_container_log_stream']
        target_label: 'stream'
````

## File: deployments/configs/tempo.yaml
````yaml
server:
  http_listen_port: 3200

distributor:
  receivers:
    otlp:
      protocols:
        grpc:
          endpoint: 0.0.0.0:4317
        http:
          endpoint: 0.0.0.0:4318

ingester:
  max_block_duration: 5m

compactor:
  compaction:
    block_retention: 24h

storage:
  trace:
    backend: local
    local:
      path: /var/tempo/traces
    wal:
      path: /var/tempo/wal
````

## File: deployments/k6/k6-testrun.yaml
````yaml
apiVersion: k6.io/v1alpha1
kind: TestRun
metadata:
  name: distributed-ws-stress-test
  namespace: chat-system
spec:
  parallelism: 10 # Tự động tạo 10 k6-runner pods phân tán trên các node K8s
  script:
    configMap:
      name: k6-test-script
      file: test-ws-load.js
  arguments: --out statsd
  runner:
    env:
      - name: WS_GATEWAY_URL
        value: "ws://nginx-gateway.chat-system.svc.cluster.local:80/ws"
      - name: K6_STATSD_ADDR
        value: "victoriametrics.chat-system.svc.cluster.local:8125"
    resources:
      requests:
        cpu: "500m"
        memory: "512Mi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: k6-test-script
  namespace: chat-system
data:
  test-ws-load.js: |
    # Test script will be mounted from test-ws-load.js
````

## File: deployments/k6/test-ws-load.js
````javascript
import ws from 'k6/ws';
import { check, sleep } from 'k6';
import { Trend, Counter } from 'k6/metrics';

// Custom Prometheus/VictoriaMetrics compatible Trends & Counters
const wsRoundTripTime = new Trend('chat_ws_round_trip_duration_ms', true);
const wsMessagesSent = new Counter('chat_ws_messages_sent_total');
const wsMessagesReceived = new Counter('chat_ws_messages_received_total');
const wsConnectionErrors = new Counter('chat_ws_connection_errors_total');

export const options = {
  scenarios: {
    stress_test: {
      executor: 'ramping-vus',
      startVUs: 100,
      stages: [
        { duration: '1m', target: 1000 },   // Warm-up 1k VUs
        { duration: '3m', target: 10000 },  // Scale to 10k VUs
        { duration: '5m', target: 50000 },  // Extreme High-Load 50k VUs
        { duration: '5m', target: 100000 }, // Breaking point 100k VUs
        { duration: '2m', target: 0 },      // Ramp down
      ],
      gracefulStop: '30s',
    },
  },
  thresholds: {
    'chat_ws_round_trip_duration_ms': ['p(95)<150', 'p(99)<300'], // SOTA SLA p99 < 300ms
    'chat_ws_connection_errors_total': ['count<100'],
  },
};

export default function () {
  const vuId = __VU;
  const iterId = __ITER;
  const userId = `user_${vuId}`;
  const deviceId = `device_${vuId}_${iterId}`;
  const receiverId = `user_${(vuId % 1000) + 1}`; // Gửi vòng tròn giữa các active VUs

  // Gateway URL (Load Balanced across ws-gateway nodes)
  const gatewayUrl = __ENV.WS_GATEWAY_URL || 'ws://nginx-gateway.chat-system.svc.cluster.local:80/ws';
  const url = `${gatewayUrl}?user_id=${userId}&device_id=${deviceId}`;

  const params = {
    headers: {
      'User-Agent': 'k6-load-generator/1.0',
    },
  };

  const res = ws.connect(url, params, function (socket) {
    socket.on('open', function () {
      // 1. Send Heartbeat immediately
      socket.send(JSON.stringify({
        type: 'HEARTBEAT',
        client_msg_id: `hb_${vuId}_${Date.now()}`,
        payload: {}
      }));

      // 2. Định kỳ bắn tin nhắn mỗi 2-3 giây
      socket.setInterval(function () {
        const clientMsgId = `k6_${vuId}_${Date.now()}`;
        const sendTimestamp = Date.now();

        const messagePayload = {
          type: 'SEND_MESSAGE',
          client_msg_id: clientMsgId,
          payload: {
            conversation_id: `conv_${Math.min(vuId, 100)}`,
            receiver_id: receiverId,
            content: `Stress payload from VU ${vuId} at ${sendTimestamp}`,
          }
        };

        socket.send(JSON.stringify(messagePayload));
        wsMessagesSent.add(1);
      }, 2500);
    });

    socket.on('message', function (data) {
      try {
        const msg = JSON.parse(data);
        if (msg.type === 'MESSAGE_DELIVERED' || msg.type === 'MESSAGE_SUBMITTED') {
          wsMessagesReceived.add(1);
          if (msg.timestamp) {
            const rtt = Date.now() - msg.timestamp;
            if (rtt > 0 && rtt < 60000) {
              wsRoundTripTime.add(rtt);
            }
          }
        }
      } catch (err) {
        // ignore malformed frame
      }
    });

    socket.on('error', function (e) {
      wsConnectionErrors.add(1);
    });

    socket.on('close', function () {
      // socket closed
    });

    // Giữ kết nối trong 60 giây trước khi VU bắt đầu vòng lặp tiếp theo
    socket.setTimeout(function () {
      socket.close();
    }, 60000);
  });

  check(res, { 'WebSocket connected successfully': (r) => r && r.status === 101 });
}
````

## File: deployments/k8s/06-nginx-gateway.yaml
````yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-conf
  namespace: chat-system
data:
  nginx.conf: |
    events {
        worker_connections 2048;
    }

    http {
        include       /etc/nginx/mime.types;
        default_type  application/octet-stream;

        sendfile        on;
        keepalive_timeout  65;

        map $http_upgrade $connection_upgrade {
            default upgrade;
            '' close;
        }

        upstream api_backend {
            server api-service.chat-system.svc.cluster.local:3000;
        }

        upstream ws_backend {
            least_conn;
            server ws-gateway-0.ws-gateway-headless.chat-system.svc.cluster.local:8080;
            server ws-gateway-1.ws-gateway-headless.chat-system.svc.cluster.local:8080;
            server ws-gateway-2.ws-gateway-headless.chat-system.svc.cluster.local:8080;
        }

        server {
            listen 80;
            server_name _;

            # 1. Routing WebSocket Handshake
            location /ws {
                proxy_pass http://ws_backend;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection $connection_upgrade;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_read_timeout 3600s;
                proxy_send_timeout 3600s;
            }

            # 2. Routing REST API
            location / {
                proxy_pass http://api_backend;
                proxy_http_version 1.1;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_read_timeout 60s;
            }
        }
    }
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-gateway
  namespace: chat-system
  labels:
    app: nginx-gateway
spec:
  replicas: 1
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: nginx-gateway
  template:
    metadata:
      labels:
        app: nginx-gateway
    spec:
      nodeSelector:
        ingress-ready: "true"
      tolerations:
        - key: node-role.kubernetes.io/control-plane
          operator: Exists
          effect: NoSchedule
      containers:
        - name: nginx
          image: nginx:alpine
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 80
              hostPort: 80
          volumeMounts:
            - name: config-volume
              mountPath: /etc/nginx/nginx.conf
              subPath: nginx.conf
      volumes:
        - name: config-volume
          configMap:
            name: nginx-conf
---
apiVersion: v1
kind: Service
metadata:
  name: nginx-gateway
  namespace: chat-system
spec:
  type: ClusterIP
  selector:
    app: nginx-gateway
  ports:
    - name: http
      port: 80
      targetPort: 80
````

## File: deployments/k8s/07-prometheus-agent.yaml
````yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus-agent
  namespace: chat-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus-agent
rules:
  - apiGroups: [""]
    resources:
      - nodes
      - nodes/proxy
      - services
      - endpoints
      - pods
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus-agent
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus-agent
subjects:
  - kind: ServiceAccount
    name: prometheus-agent
    namespace: chat-system
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: prometheus-agent-config
  namespace: chat-system
data:
  prometheus.yml: |
    global:
      scrape_interval: 2s
      scrape_timeout: 2s

    remote_write:
      - url: "http://prometheus:9090/api/v1/write"

    scrape_configs:
      - job_name: 'ws-gateway'
        static_configs:
          - targets:
              - 'ws-gateway-0.ws-gateway-headless.chat-system.svc.cluster.local:9091'
              - 'ws-gateway-1.ws-gateway-headless.chat-system.svc.cluster.local:9091'
              - 'ws-gateway-2.ws-gateway-headless.chat-system.svc.cluster.local:9091'

      - job_name: 'chat-engine'
        kubernetes_sd_configs:
          - role: endpoints
            namespaces:
              names:
                - chat-system
        relabel_configs:
          - source_labels: [__meta_kubernetes_service_name]
            action: keep
            regex: chat-engine
          - source_labels: [__meta_kubernetes_endpoint_port_name]
            action: keep
            regex: metrics
          - source_labels: [__meta_kubernetes_pod_name]
            target_label: pod
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus-agent
  namespace: chat-system
  labels:
    app: prometheus-agent
spec:
  replicas: 1
  selector:
    matchLabels:
      app: prometheus-agent
  template:
    metadata:
      labels:
        app: prometheus-agent
    spec:
      serviceAccountName: prometheus-agent
      containers:
        - name: prometheus-agent
          image: prom/prometheus:v2.54.1
          args:
            - "--config.file=/etc/prometheus/prometheus.yml"
            - "--storage.agent.path=/prometheus-data"
            - "--enable-feature=agent"
          volumeMounts:
            - name: config
              mountPath: /etc/prometheus
            - name: storage
              mountPath: /prometheus-data
          resources:
            requests:
              cpu: 50m
              memory: 64Mi
            limits:
              cpu: 500m
              memory: 256Mi
      volumes:
        - name: config
          configMap:
            name: prometheus-agent-config
        - name: storage
          emptyDir: {}
````

## File: deployments/k8s/08-promtail.yaml
````yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: promtail
  namespace: chat-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: promtail
rules:
  - apiGroups: [""]
    resources:
      - nodes
      - nodes/proxy
      - services
      - endpoints
      - pods
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: promtail
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: promtail
subjects:
  - kind: ServiceAccount
    name: promtail
    namespace: chat-system
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: promtail-config
  namespace: chat-system
data:
  promtail.yaml: |
    server:
      http_listen_port: 9080
      grpc_listen_port: 0

    positions:
      filename: /run/promtail/positions.yaml

    clients:
      - url: http://loki:3100/loki/api/v1/push

    scrape_configs:
      - job_name: kubernetes-pods
        static_configs:
          - targets:
              - localhost
            labels:
              job: kubernetes-pods
              __path__: /var/log/pods/*/*/*.log
        pipeline_stages:
          - regex:
              expression: '^/var/log/pods/(?P<namespace>[^_]+)_(?P<pod>[^_]+)_[a-f0-9-]+/(?P<container>[^/]+)/.*\.log$'
              source: filename
          - labels:
              namespace:
              pod:
              container:
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: promtail
  namespace: chat-system
  labels:
    app: promtail
spec:
  selector:
    matchLabels:
      app: promtail
  template:
    metadata:
      labels:
        app: promtail
    spec:
      serviceAccountName: promtail
      tolerations:
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: NoSchedule
        - key: node-role.kubernetes.io/control-plane
          operator: Exists
          effect: NoSchedule
      containers:
        - name: promtail
          image: grafana/promtail:3.1.0
          securityContext:
            runAsUser: 0
            privileged: true
          args:
            - "-config.file=/etc/promtail/promtail.yaml"
          volumeMounts:
            - name: config
              mountPath: /etc/promtail
            - name: run
              mountPath: /run/promtail
            - name: pods
              mountPath: /var/log/pods
              readOnly: true
            - name: docker
              mountPath: /var/lib/docker/containers
              readOnly: true
          resources:
            requests:
              cpu: 50m
              memory: 64Mi
            limits:
              cpu: 200m
              memory: 128Mi
      volumes:
        - name: config
          configMap:
            name: promtail-config
        - name: run
          emptyDir: {}
        - name: pods
          hostPath:
            path: /var/log/pods
        - name: docker
          hostPath:
            path: /var/lib/docker/containers
````

## File: deployments/k8s/09-client-simulator.yaml
````yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: simulator-config
  namespace: chat-system
data:
  SIM_BOTS: "500"
  SIM_CONVS: "10"
  SIM_INTERVAL: "100ms"
  SIM_DURATION: "2m"
  SIM_API_URL: "http://nginx-gateway.chat-system.svc.cluster.local"
  SIM_WS_URL: "ws://nginx-gateway.chat-system.svc.cluster.local/ws"
  SIM_PASSWORD: "Pass@123456"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: client-simulator
  namespace: chat-system
  labels:
    app: client-simulator
spec:
  replicas: 0
  selector:
    matchLabels:
      app: client-simulator
  template:
    metadata:
      labels:
        app: client-simulator
    spec:
      containers:
        - name: client-simulator
          image: chat-system/client-simulator:latest
          imagePullPolicy: IfNotPresent
          envFrom:
            - configMapRef:
                name: simulator-config
          resources:
            requests:
              cpu: "200m"
              memory: "128Mi"
            limits:
              cpu: "1000m"
              memory: "512Mi"
````

## File: deployments/k8s/chaos/01-network-delay-grpc.yaml
````yaml
apiVersion: chaos-mesh.org/v1alpha1
kind: NetworkChaos
metadata:
  name: delay-grpc-delivery
  namespace: chat-system
spec:
  action: delay
  mode: all
  selector:
    namespaces:
      - chat-system
    labelSelectors:
      app: ws-gateway
  delay:
    latency: '100ms'
    jitter: '10ms'
    correlation: '50'
  direction: to
  target:
    selector:
      namespaces:
        - chat-system
      labelSelectors:
        app: chat-engine
    mode: all
  duration: '5m'
````

## File: deployments/k8s/chaos/02-network-delay-nats.yaml
````yaml
apiVersion: chaos-mesh.org/v1alpha1
kind: NetworkChaos
metadata:
  name: delay-nats-broker
  namespace: chat-system
spec:
  action: delay
  mode: all
  selector:
    namespaces:
      - chat-system
    labelSelectors:
      app: ws-gateway
  delay:
    latency: '100ms'
    jitter: '10ms'
    correlation: '50'
  direction: to
  target:
    selector:
      namespaces:
        - chat-system
      labelSelectors:
        app: nats
    mode: all
  duration: '5m'
````

## File: deployments/k8s/chaos/03-pod-kill-gateway.yaml
````yaml
apiVersion: chaos-mesh.org/v1alpha1
kind: PodChaos
metadata:
  name: kill-ws-gateway-pod
  namespace: chat-system
spec:
  action: pod-kill
  mode: one
  selector:
    namespaces:
      - chat-system
    labelSelectors:
      app: ws-gateway
  scheduler:
    cron: '@every 2m'
````

## File: deployments/k8s/chaos/04-mobile-network-delay.yaml
````yaml
apiVersion: chaos-mesh.org/v1alpha1
kind: NetworkChaos
metadata:
  name: mobile-network-delay
  namespace: chat-system
spec:
  action: delay
  mode: all
  selector:
    namespaces:
      - chat-system
    labelSelectors:
      app: client-simulator
  delay:
    latency: '80ms'
    jitter: '25ms'
    correlation: '50'
  direction: both
  target:
    selector:
      namespaces:
        - chat-system
      labelSelectors:
        app: nginx-gateway
    mode: all
  duration: '10m'
````

## File: deployments/k8s/chaos/05-mobile-packet-loss.yaml
````yaml
apiVersion: chaos-mesh.org/v1alpha1
kind: NetworkChaos
metadata:
  name: mobile-packet-loss
  namespace: chat-system
spec:
  action: loss
  mode: all
  selector:
    namespaces:
      - chat-system
    labelSelectors:
      app: client-simulator
  loss:
    loss: '3'
    correlation: '25'
  direction: both
  target:
    selector:
      namespaces:
        - chat-system
      labelSelectors:
        app: nginx-gateway
    mode: all
  duration: '10m'
````

## File: deployments/k8s/chaos/06-mobile-network-partition.yaml
````yaml
apiVersion: chaos-mesh.org/v1alpha1
kind: NetworkChaos
metadata:
  name: mobile-network-partition
  namespace: chat-system
spec:
  action: partition
  mode: all
  selector:
    namespaces:
      - chat-system
    labelSelectors:
      app: client-simulator
  direction: both
  target:
    selector:
      namespaces:
        - chat-system
      labelSelectors:
        app: nginx-gateway
    mode: all
  duration: '45s'
````

## File: deployments/kind-config.yaml
````yaml
# deployments/kind-config.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    kubeadmConfigPatches:
      - |
        kind: InitConfiguration
        nodeRegistration:
          kubeletExtraArgs:
            node-labels: "ingress-ready=true"
    extraPortMappings:
      - containerPort: 80
        hostPort: 80
        protocol: TCP
      - containerPort: 443
        hostPort: 443
        protocol: TCP
      - containerPort: 80
        hostPort: 8080
        protocol: TCP
  - role: worker
  - role: worker
````

## File: design/1 Database Design & Data Modeling 3c23b644f0b8806193cbcb2d52d62f6d.md
````markdown
# 1. Database Design & Data Modeling

1. Database Selection Strategy:

Chọn giữa Postgres (Relational), Columnar NoSQL, MongoDB.

Yêu cầu là tin nhắn được lưu trữ vĩnh viễn.

Số lượng tin nhắn chỉ có thể tăng lên và không giảm đi.

Loại bỏ SQL vì rất khó scale ngang, và khi dữ liệu càng lớn, tính chất của B-Tree Index làm suy giảm tốc độ ghi rất nhiều. Vẫn có thể sử dụng để lưu meta data như profile, conversation, friend, …

Loại bỏ DocumentDB như MongoDB bởi vì nó cũng dùng B-Tree index tương tự như Postgres, cho nên khi lượng tin nhắn đạt một ngưỡng nhất định, việc ghi sẽ trở nên ỳ ạch. 

MongoDB cũng duy trì toàn bộ B-Tree Index trong RAM, vì vậy phải tốn một lượng lớn ngân sách cho RAM để lưu được index khổng lồ của nó. 
Thêm vào đó là mỗi document của nó giới hạn 16MB, mỗi lần vượt quá 16MB thì chúng ta phải tự xử lý dữ liệu (có thể là áp dụng pattern bucket gì đó).

Chọn Cassandra vì nó là LSM-Tree Engine. Nó cho phép append dữ liệu vào cuối file log. Điều này giúp đẩy tốc độ ghi dữ liệu lên tối đa. Việc này cực kì thích hợp cho hệ thống nhắn tin phải ghi dữ liệu liên tục. 

Có thể dùng channel_id hoặc conversation_id để làm partition key.

Và message_id (Snowflake) làm clustering key, giúp message được tự động sắp xếp theo thứ tự ngay trên đĩa. 

Và đặc biệt là Cassandra không có Single Point of Failure, tự động horizontal Scaling.

1. Model Design:

```markdown
Users 
- user_id
- user_name
- password_hash
- name
- phone
- address
- created_at
- updated_at
- deleted_at
```

```markdown
Enum friend_status
- pending
- accepted
- declined
- blocked

Friends
- id
- requester_id
- addressee_id 
- status
- created_at
- updated_at 
```

```markdown
Conversations
- id
- name
- icon_url
- members
- created_at
- updated_at
```

```markdown
Messages 
- conversation_id
- id (uuidv7)
- sender_id
- content
- media_url
- created_at
- updated_at
```

```markdown
Conversation Read Receipts
- conversation_id
- user_id
- last_message_seen
- last_read_at
```
````

## File: design/Part 1 Requirements, Constraints, and Capacity Est 3be3b644f0b88019bd11df04cdaf19fe.md
````markdown
# Part 1: Requirements, Constraints, and Capacity Estimation

1. Scope:
    
    The system is a real-time messaging platform similar to Messenger or Whatsapp.
    
    For the first version, we focus on the core messaging experience rather than trying to support every possible chat feature.
    
    In Scope:
    
    - One to one messaging
    - Group messaging
    - Real-time delivery
    - Message history
    - Online/offline presence
    - Delivery status
    - Read receipts
    - Multi-device support
    - Reconnection and missed-message syncronization
    
    Out of Scope:
    
    For the initial design, we will not focus on:
    
    - Voice calls
    - Video calls
    - Stories
    - End-to-end encryption
    - Full-text message search
    - Large media processing pipelines
    
    Attachments such as images or videos may be supported later through object storage, but they are not part of the core message delivery design.
    
2. Functional Requirements:
    1. FR1 - Send messages:
        - A user must be able to send messages to another user, or to other group.
        - A message contains at least:
            - message_id
            - conversation_id
            - sender_id
            - message_content
            - server_timestamp
    2. FR2 - Real-Time Message Delivery
        - If the recipient is online, the system should delivery the message in near real time.
        - The target user experience is: a message sent by User A should normally appear on User B’s device within a few hundred miniseconds.
    3. FR3 - Offline Messaging
        - If the recipient is offline, the message must not be lost.
        - The message should remain persisted in the system.
        - When the recipient reconnects, the client should synchronize messages that were missed while offline.
        - Should show NOTIFICATION when recipient is offline.
    4. FR4 - Message History
        - User must be able to retrieve previous messages from a conversation.
        - Can retrieve the most recent N messages from conversation X.
        - Older messages should be loaded using cursor-based pagination.
    5. FR5 - Message Ordering
        - Messages of a conversation must be handled by order.
        - Messages in unrelated conversations may be processed independently.
    6. FR6 - Duplicate Handling
        - Network failures may cause clients or infrastructure conponents to retry message delivery.
        - Therefore, the system must tolerate duplicate message submissions.
        - Each message should have a globally unique message_id
        - Processing should be idempotent so that retrying the same message does not create multiple persisted messages.
    7. FR7 - Group Chat
        - The system should support group conversations.
        - There are four type of conversation: global, region (VN, US, …), clan, directly.
    8. FR8 - Multi-Device Support.
        - A user may connect from multiple devices.
        - Messages should eventually syncronize accress all active devices belonging to the user.
        - A user may have multiple active connections.
    9. FR9 - Online Presence.
        - The system should expose approximate user presence: online, offline, lastseen.
        - Presence doesn’t require strong consistency.
        - Temporary inaccuracies are acceptable.
    10. FR10 - Delivery Status
        - Messages should support at least following status: sent, delivered, read.
        - System must save who read message.
    11. FR11 - Reconnection
        - WebSocket connections may disappear because of:
            - unstable mobile network
            - wifi changes
            - NAT timeout
            - gateway failure
            - application restart
            - device sleep
        - After reconnecting, the client should synchronize its state using persisted message history rather than assuming every real-time delivery succeed.
    
3. Non-functional Requirements:
    1. NFR1 - Low Latency
        - Real-time message delivery should have low latency
        - Target: P50 < 100ms and P99 < 500ms for online users located within the same region under normal operating conditions.
        - Latency across geographically distant regions maybe higher.
    2. NFR2 - High Availability
        - Sending and receiving messages are core functionality and should remain available despite individual machine failures.
        - Target availability: 99.99%
        - Individual failures such as:
            - gateway crash
            - chat worker crash
            - broker node crash
            - database replica failure
        - should not take down the entire messaging service
    3. NFR3 - Durability: 
        - Once the sender receives a successful **sent** acknowledgment, the message must not be lost because of a single server failure.
        - Message history therefore requires durable storage.
    4. NFR4 - Horizontal Scalability
        - The system should scale horizontally
        - We should able to increase capacity by adding:
            - Websocket Gateway
            - Chat Worker
            - broker partitions/nodes
            - database nodes
    5. NFR5 - Fault Tolerance
        - The system should tolerate partial failures
        - Example include:
            - Gateway alive - Chat Worker Slow
            - Broker alive - Database unavailable
            - Gateway alive - Redis unavailable
            - One data node unavailable - Other nodes healthy
        - The system should degrade gracefully instead of producing cascading failures
    6. NFR6 - Eventual Consistency Where Appropiate
        - Not every piece of chat state requires strong consistency.
        - Strong guarantees are important for:
            - durable messages
            - conversation membership authorization
            - message identity
        - Eventual consistency is acceptable for:
            - online presence
            - read receipts
            - delivery status
            - multi-device synchronization
        - The design should avoid paying the cost of strong consistency where the product doesn’t require it.
    7. NFR7 - Backpressure
        - The system must protect itself when downstream components become slower than incoming traffic.

1. Security Requirements
    - Users must authenticate before establishing a Websocket session.
    - The backend must also verify that a sender is authoried to send messages to the requested conversation.
    - Additional protections include:
        - TLS
        - maximum message size
        - per-user rate limiting
        - IP-based abuse protection
        - authentication during connection establishment

1. Initial Scale Assumptions
    - For the purpose of this design, assume:
        - Registered users: 100 million
        - Daily active users: 20 million
        - Peak concurrent users: 5 million
        - Assume each active user sends: 40 messages/day
        - Therefore: 20M x 40 = 800M messages/day
        - Average message rate: 8000 messages/sec
        - Peak traffic = 5 x average = 40,000 messages/sec
        - Design target: 50,000 inbound messages/sec

1. Message Size Assumption
    - Assume an average persisted message record of approximately: 500 bytes
    - including: id, timestamp, metadata, text payload, storage overhead estimate
    - Daily raw message storage: 800M * 500 bytes = 400GB/day
    - Per year: 400GB x 365 = 146TB/year
    - This number excludes:
        - replication
        - indexes
        - backups
        - metadata
        - attachments

1. Websocket Connection Scale
    - Peak concurrent connections: 5 milion
    - Assume during early capacity planning that one Gateway safely handles: 50,000 concurrent WebSocket connections
    - Then the theoretical minimum is: 100 Gateway instances
    - Production deployment needs additional headroom for:
        - machine failures
        - rolling deployments
        - connection spikes
        - uneven connection distribution
    - Therefore, the actual fleet should be larger than the theoretical minimum.
````

## File: design/Part 2 High-Level Architecture & End-to-End Messag 3be3b644f0b8804aa7cdc7af6fe247cd.md
````markdown
# Part 2: High-Level Architecture & End-to-End Message Flow

![image.png](image.png)

1. API Service:
    - Xử lý Authentication
    - Fetch History Message, Conversation
2. Websocket Gateway:
    - Giữ kết nối socket của các client, có thể có nhiều node để scale phục vụ cho nhiều client
    - Mỗi ws gateway node đăng ký một queue/channel riêng. Channel này dùng để làm địa chỉ gateway cho việc delivery tin nhắn đến người nhận.
    - Mission:
        - Thiết lập kết nối từ client và set presence lên Redis presence Service
        - Nhận message từ client và gửi lên Broker
        - Consume message từ broker và gửi xuống cho client
3. Message Broker:
    - Tác dụng là decoupling giữa websocket gateway và chat worker.
    - Many to many fanout
    - Buffer Message, backpressure giúp điều tiết message dựa trên khả năng làm việc của các Gateway WS.
    - Message buffer để điều tiết khả năng làm việc của chat worker.
    - Channel:
        - Direct Message: message được publish lên channel GateWay. Ví dụ Gateway 1 thì sẽ có channel gateway_1.
        - Clan Message: channel Clan (clan:abc, clan:xyz)
        - Region Message: channel Region (vn, cn, us)
        - Global Message: channel global (global)
    - Việc phân chia channel như trên giúp giảm số lượng tin nhắn gửi đến clan, region, global có lượng người dùng lớn. Chúng ta chỉ cần gửi một tin nhắn đến channel là được.
4. Chat Worker:
    - Validate quyền gửi tin, Idempotency check (chống trùng), ghi DB.
    - Query presence từ Redis, nếu như không online thì thực hiện gửi vào noti service. Nếu online thì bắn qua Broker.
5. Redis Presence Service:
    - Lưu trữ presence cho các user và các device tương ứng.
    - Công thức: {user_id}:{device_id}:{gateway_node}
6. Notification Service: 
    - Xử lý đẩy Apple APNs / Google FCM khi Receiver đang offline hoặc App ở background.
7. Database:
    - **Partition Key:** conversation_id (đảm bảo toàn bộ tin nhắn trong 1 room nằm cùng 1 Partition node).
    - **Clustering Key:** message_id (Time-based UUIDv7) giúp sort tin nhắn theo thời gian tự nhiên.

## 3. Detailed End-to-End Flows

### Flow 1: Gửi tin nhắn Direct (1-1) & Luồng ACK 2 Chiều (Online Case)

Kịch bản: **Client A** gửi tin nhắn cho **Client B** (Client B đang Online tại WS Gateway 2).

- **Bước 1 - 3 (Inbound):** Client A gửi tin nhắn qua WebSocket. WS Gateway 1 đẩy ngay vào Broker Inbound Topic.
- **Bước 4 - 5 (Sender ACK):** Chat Worker consume tin nhắn, lưu vào ScyllaDB với trạng thái SENT. Đồng thời trả ngay **Sender ACK** về cho Client A (Client A hiển thị dấu tích xám ).
- **Bước 6 - 9 (Outbound Delivery):** Worker tra cứu Redis thấy Client B đang online ở WS Gateway 2. Worker publish event vào Queue gateway-node-02. WS Gateway 2 lấy tin nhắn và push qua WebSocket xuống Client B.
- **Bước 10 - 13 (Receiver ACK & Delivery Status):** Client B nhận tin thành công, App tự động trả về ACK_DELIVERED. Chat Worker nhận ACK này, update DB thành DELIVERED và bắn notify cho Client A (Client A đổi thành 2 dấu tích ).

### Flow 2: Offline Messaging & Sync khi Reconnect

Kịch bản: **Client B** bị đứt mạng/offline khi Client A gửi tin.

1. **Khi Client B Offline:**
    - Tại **Bước 6 (Flow 1)**, Chat Worker query Redis Presence  Trả về OFFLINE.
    - Worker không bắn vào Gateway Queue mà push event sang Notification Queue.
    - Notification Service gửi Push Notification (APNs/FCM) tới thiết bị Client B: *"Bạn có tin nhắn mới từ A"*.
2. **Khi Client B Reconnect (Sync State):**
    - Client B mở lại App, thiết lập kết nối WebSocket mới tới WS Gateway 3.
    - Client B **KHÔNG** đợi server push lại tin nhắn qua WebSocket, mà sẽ chủ động gọi một **HTTP REST API** sang API Gateway:
        
        GET /v1/conversations/{id}/messages?after_id={last_received_msg_id}
        
    - API Gateway đọc từ ScyllaDB và trả về danh sách tin nhắn Client B đã bỏ lỡ (Missed Messages).
    - **Ưu điểm:** Giúp giảm tải tuyệt đối cho Hệ thống WebSocket Real-time, tránh tình trạng "Stampeding Herd" (hàng ngàn client reconnect cùng lúc làm ngập luồng Socket).

## 4. Architectural Trade-offs & Deep Dives (Tư duy Thiết kế)

### 1. Tại sao dùng Hybrid Fan-out (Direct vs Group Chat)?

- **1-1 & Small Group (< 100 members):** Dùng **Fan-out on Write**. Worker nhân bản tin nhắn và gửi tới Queue của từng member.
- **Large Group / Channel (> 10,000 members):** Dùng **Fan-out on Read / Gateway Local Broadcast**. Worker chỉ gửi **1 tin nhắn duy nhất** tới từng WS Gateway Node có chứa member của Group đó. Sau đó, bản thân Gateway Node đó sẽ tự Broadcast tới các Socket Connections cục bộ. Cấu trúc này tránh làm "nổ" Message Broker.

### 2. Đảm bảo Idempotency (Chống trùng tin nhắn)

Do mạng chập chờn, Client A có thể gửi lại 1 tin nhắn nhiều lần (Retry).

- Client luôn tạo một client_msg_id (UUIDv7) tại Local.
- Server dùng client_msg_id làm Unique Key/Deduplication Key tại Redis trong  giây. Nếu gặp ID đã xử lý, Server bỏ qua việc ghi DB và trả về ngay Success ACK cũ.
````

## File: knowledge/edge_cases/edge_cases_and_observability.md
````markdown
# 📌 Edge Cases, Failure Modes & Observability Backlog

Tài liệu này ghi nhận các góc tối kỹ thuật (Edge Cases), bẫy chạy đua (Race Conditions), và các kịch bản lỗi trong môi trường phân tán cần được theo dõi qua hệ thống Metrics/Logging và lên kế hoạch xử lý/tối ưu hóa sau này.

---

## 1. Edge Case: Phantom ACK / False Success do Race Condition giữa Redis Lock & DB Persistence

### 🔍 Mô tả & Bối cảnh (Context & Failure Scenario)
- **Vị trí:** `services/chat-engine/internal/usecase/chat_usecase.go` & `idempotency_repo.go`.
- **Kịch bản:**
  1. Client gửi tin nhắn `M1` (`client_msg_id = X`).
  2. Chat Worker 1 nhận `M1`, thực hiện `AcquireLock(X)` trên Redis thành công (`SetNX`).
  3. Chat Worker 1 bắt đầu ghi `M1` vào Cassandra (quá trình này mất vài giây hoặc timeout do mạng chập chờn / Cassandra chậm).
  4. Trong lúc Worker 1 chưa ghi xong, Client retry hoặc NATS Broker bắn lại bản tin duplicate `M1'` (`client_msg_id = X`).
  5. Chat Worker 2 (hoặc Worker 1) nhận `M1'`, kiểm tra `AcquireLock(X)` -> Thấy key đã tồn tại (`!isNew`).
  6. Worker coi đây là tin nhắn trùng lặp đã hoàn tất -> **Bắn ngay `Sender ACK` về cho Client (Phantom ACK)**.
  7. Sau đó, thao tác ghi Cassandra ở bước 3 bị **Thất Bại (Timeout / Crash)** -> Worker gọi `ReleaseLock(X)`.

### ⚠️ Hậu quả (Impact)
- **Silent Data Loss:** Client nhìn thấy tin nhắn đã gửi thành công (giao diện hiện 1 tích), nhưng thực chất tin nhắn **chưa từng được ghi vào cơ sở dữ liệu**. Khi reload app hoặc phía Receiver mở app, tin nhắn biến mất hoàn toàn.

### 📊 Điểm quan sát & Metric giám sát (Observability Signals)
- **Log cảnh báo:** `WARN: Redis idempotency check failed` hoặc `Duplicate message detected: client_msg_id=...` xuất hiện cùng thời điểm với `persistence failed`.
- **Prometheus Metric cần bổ sung sau này:**
  - `chat_engine_idempotency_lock_acquired_total`
  - `chat_engine_idempotency_lock_released_on_error_total`
  - Tỷ lệ lệch (Discrepancy) giữa tổng số `Sender ACK dispatched` và tổng số record được `INSERT` thành công vào Cassandra.

### 🛠️ Giải pháp dài hạn (Long-term Remediation)
- **FSM State Machine trên Idempotency Key:**
  - `PROCESSING` (TTL ngắn 10-30s, đóng vai trò Distributed Mutex).
  - `COMPLETED` (TTL dài 24h, chỉ set sau khi DB ghi thành công).
- **Concurrent Duplicate Handling:**
  - Nếu duplicate đến khi đang `PROCESSING`: Không gửi ACK ngay; hoặc chờ (polling 500ms) hoặc reject/NACK để Broker redeliver sau.

---

## 2. Edge Case: Cassandra Zombie Write sau khi Context Timeout

### 🔍 Mô tả & Bối cảnh
- Worker thiết lập `dbCtx, cancel := context.WithTimeout(ctx, dbTimeout)` (mặc định cấu hình từ `timeout_seconds`).
- Nếu Cassandra bị quá tải hoặc Garbage Collection (GC Pause), request ở phía Go driver có thể bị timeout và trả về error, worker rollback Redis lock.
- Tuy nhiên, coordinator node phía Cassandra có thể vẫn nhận được mutation trước đó và ghi vào CommitLog/MemTable thành công sau đó vài mili-giây.

### 📊 Điểm quan sát
- Kiểm tra các bản ghi Cassandra có `client_msg_id` tồn tại nhưng application log lại ghi nhận `persistence failed: context deadline exceeded`.

---

## 3. Edge Case: Receiver Connection Drop ngay sau khi Query Presence (Ghost Online)

### 🔍 Mô tả & Bối cảnh
- Chat Engine query Redis Presence thấy Client B đang online tại `gateway_node_02`.
- Trong tích tắc trước khi Outbound Message từ NATS đến được `gateway_node_02`, Client B bị mất kết nối (rớt mạng, tắt 4G, killed app).
- `gateway_node_02` không tìm thấy WebSocket connection trong local memory map -> Tin nhắn có nguy cơ bị drop nếu không có cơ chế fallback hoặc Receiver ACK timeout.

### 🛠️ Giải pháp quan sát & xử lý
- Đếm metric `ws_gateway_local_delivery_dropped_total`.
- WS Gateway nếu không tìm thấy socket connection cục bộ phải có cơ chế phản hồi lại Broker hoặc publish sang `chat.notification` để bù Push Notification.
````

## File: knowledge/edge_cases/README.md
````markdown
# Danh Mục Edge Cases & Điểm Quan Sát Hệ Thống (Edge Cases & Observability Backlog)

Thư mục này dùng để lưu trữ và theo dõi các góc tối kỹ thuật (Edge cases), rủi ro tiềm ẩn, và các sự cố phân tán được phát hiện trong quá trình code để tiện theo dõi, giám sát qua metric và tối ưu hóa sau này.

---

### Danh sách tài liệu:
1. [Edge Cases, Failure Modes & Observability Backlog](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/edge_cases/edge_cases_and_observability.md)
   - *Phantom ACK / False Success* (Race condition Redis Idempotency vs Cassandra Failure).
   - *Cassandra Zombie Write* (Context Timeout vs Mutation Commit).
   - *Ghost Online Drop* (Connection drop ngay sau khi query Redis Presence).
````

## File: knowledge/system_design/01_cassandra_upsert_vs_rdbms_and_idempotency.md
````markdown
# 1. Cassandra UPSERT vs RDBMS & Bản chất Tầng Idempotency (Redis)

---

## 1. Bẫy Kiến Trúc: Cassandra xử lý Trùng Primary Key như thế nào?

### ❌ Tư duy RDBMS thông thường (Postgres, MySQL)
- Trong cơ sở dữ liệu quan hệ, câu lệnh `INSERT` khi gặp trùng khóa chính (`Primary Key`) sẽ lập tức **ném ra lỗi vi phạm ràng buộc (`Duplicate Key Violation`)** và huỷ transaction.

### 🟢 Bản chất thực tế của Cassandra (LSM-Tree NoSQL)
- **Cassandra KHÔNG ném ra lỗi khi trùng Primary Key.**
- Trong Cassandra, câu lệnh `INSERT` thực chất là một **UPSERT (Insert or Overwrite)**.
- Nếu bạn thực thi lệnh `INSERT` vào cùng một bộ khóa `((conversation_id), id)` nhiều lần:
  - Cassandra sẽ **âm thầm ghi đè (overwrite)** các trường dữ liệu với timestamp mới nhất.
  - Trình điều khiển (Driver `gocql`) luôn nhận được kết quả thành công (`nil error`).
- *(Lưu ý: Cú pháp `INSERT ... IF NOT EXISTS` là Lightweight Transaction sử dụng thuật toán Paxos, cực kỳ tốn chi phí mạng và giảm thông lượng ghi trầm trọng, không bao giờ dùng cho hệ thống Chat).*

---

## 2. Tại sao BẮT BUỘC phải có Tầng Idempotency Key (Redis)?

Nếu Cassandra đã tự động ghi đè không báo lỗi, tại sao hệ thống phân tán vẫn bắt buộc phải có tầng Idempotency ở Redis phía trước?

### A. Ngăn chặn Duplicate Side-Effects (Tác dụng phụ ngoài luồng)
Lưu vào Database chỉ là bước đầu tiên trong pipeline xử lý. Sau đó Worker còn phải:
1. Gửi **Sender ACK** cho người gửi.
2. Tra cứu Presence và **Dispatch tin nhắn** tới WebSocket Gateway của người nhận (Receiver) hoặc gọi **Notification Service (APNs/FCM Push)**.

👉 Nếu không có Redis chặn từ đầu:
- Khi mạng lag, Client retry gửi lại cùng 1 tin nhắn 2 lần.
- Cả 2 lần đều đi qua Worker ➔ Người nhận (Client B) sẽ nhận **2 lần tin nhắn**, điện thoại kêu **2 lần ting ting**, và Notification Service bị tiêu tốn quota vô ích.

---

### B. Tránh Write Amplification & Giảm tải Compaction cho Cassandra
- Mỗi lần `INSERT` ghi đè trong Cassandra sẽ tạo ra một bản ghi mới trong `MemTable` và xả ra `SSTable` mới trên đĩa.
- Việc ghi đè liên tục các bản ghi trùng lặp ép Cassandra phải chạy tiến trình dọn dẹp (**Compaction**) nặng nề, gây nghẽn I/O đĩa và tăng độ trễ đột biến (Latency Spikes).

---

### C. Khóa phân tán nguyên tử siêu tốc (< 1ms)
- Redis hoạt động đơn luồng (Single-threaded event loop), câu lệnh:
  ```redis
  SET idempotency:{client_msg_id} "PROCESSING" NX EX 60
  ```
- Diễn ra trong **dưới 1 mili-giây**, đóng vai trò như chiếc "khiên bảo vệ" chặn đứng mọi yêu cầu trùng lặp ngay tại cửa ngõ trước khi làm tốn tài nguyên Database và Network.
````

## File: knowledge/system_design/02_dual_write_partial_failure_and_state_reconciliation.md
````markdown
# 2. Dual-Write Partial Failure, Poison Pill & State Reconciliation

---

## 1. Bài toán Dual-Write & Thất bại một phần (Partial Failure)

Trong kiến trúc phân tán hướng sự kiện (Event-Driven Architecture), một Chat Worker thường thực hiện 2 thao tác liên tiếp:
1. **Lưu dữ liệu vào Database:** Ghi tin nhắn an toàn vào Cassandra (`Source of Truth`).
2. **Kích hoạt Side-Effect qua Broker:** Gửi `Sender ACK` hoặc `Outbound Event` tới NATS Broker.

### Tình huống:
Nếu **Bước 1 THÀNH CÔNG** (tin đã vào Cassandra) nhưng **Bước 2 THẤT BẠI** (NATS Broker lag hoặc Gateway Node bị rớt mạng):
- **Nếu Worker `return err`:** 
  - NATS sẽ coi như tin nhắn chưa được xử lý và kích hoạt **Broker Redelivery**.
  - *Rủi ro:* Nếu Gateway Node đó chết hẳn, tin nhắn sẽ bị retry vô tận (**Poison Pill**), gây tắc nghẽn toàn bộ hàng đợi (**Head-of-Line Blocking**), đồng thời có nguy cơ gửi trùng tin nhắn nhiều lần cho người nhận (Receiver).
- **Nếu Worker log cảnh báo và `return nil`:**
  - Tin nhắn được coi là hoàn tất vì dữ liệu gốc đã nằm an toàn trong Database.
  - Luồng xử lý không bị nghẽn, duy trì Throughput cao nhất (**Fast-Fail & Move-On**).

---

## 2. Vấn đề "Phantom Message" & Split-Brain giữa Client và Server

Khi Worker chọn `return nil`, một Edge Case xảy ra ở phía Client:
1. Client A gửi tin nhắn với `client_msg_id = uuid_123`.
2. Server lưu Cassandra thành công nhưng chiều gửi ACK về cho Client A bị đứt gói tin.
3. Client A hết thời gian chờ (Timeout) và gửi lại (Retry) `uuid_123`.
4. Mạng của Client A vẫn chưa ổn định, tất cả các lần retry đều không nhận được ACK.
5. **Hậu quả nếu thiết kế Client kém:** Client A đánh dấu tin nhắn là **Gửi thất bại (Chấm than đỏ ⚠️)**, trong khi thực tế tin nhắn đã nằm trên Server và người nhận (Client B) đã đọc được!

---

## 3. Giải pháp Chuẩn Production: State Reconciliation & Smart Client Pattern

Các hệ thống quy mô lớn (Telegram, WhatsApp, Discord) giải quyết triệt để vấn đề này qua 3 nguyên tắc:

### A. Phân biệt Permanent Failure vs Transient Failure
- **Permanent Failure (Lỗi nghiệp vụ vĩnh viễn):** Bị chặn, bị kích khỏi nhóm, vi phạm chính sách ➔ Server trả về mã lỗi rõ ràng ➔ Client hiển thị **Chấm than đỏ ⚠️ ngay lập tức**.
- **Transient Failure (Lỗi mạng/Timeout tạm thời):** Không nhận được ACK trong 5s ➔ Client **KHÔNG báo đỏ**, giữ trạng thái "Đang gửi (Đồng hồ xoay ⏳)" và kích hoạt kiểm tra Socket.

### B. Tự phục hồi kết nối (Socket Self-Healing)
- Khi nghi ngờ kết nối bị "đứt ngầm" (Half-Open Socket), Client chủ động ngắt kết nối WebSocket cũ và kết nối lại (Reconnect) tới Gateway mới.

### C. Đối soát trạng thái 2 chiều (State Reconciliation)
- Khi kết nối lại, Client gửi danh sách các tin nhắn đang ở trạng thái chờ:
  ```json
  POST /v1/messages/reconcile
  {
      "pending_client_msg_ids": ["uuid_123"]
  }
  ```
- **Phía Server:** Tra cứu nhanh Redis/Cassandra:
  - Nếu `uuid_123` đã có trong DB ➔ Trả về `"STATUS": "PERSISTED"`.
  - Client tự động chuyển UI từ **"Đang gửi ⏳"** sang **"Đã gửi ✔"** (Không bao giờ bị lỗi ảo!).
  - Nếu `uuid_123` chưa có ➔ Client mới thực hiện gửi lại hoặc cho phép user bấm "Thử lại".
````

## File: knowledge/system_design/03_idempotency_state_machine_and_ttl_rollback.md
````markdown
# 3. Idempotency State Machine, Lock Rollback & Two-Phase TTL

---

## 1. Vấn đề "Kẹt Khóa Vĩnh Viễn" (Permanent Lockout) khi DB Gặp Lỗi

### Tình huống:
1. Worker nhận tin nhắn `client_msg_id = uuid_123`.
2. Worker chiếm khóa trên Redis: `SET idempotency:msg:uuid_123 "PROCESSED" NX EX 86400` (TTL 24h).
3. Ghi vào Cassandra: **BỊ LỖI** (Database timeout, quá tải hoặc network partition).
4. Hàm trả về lỗi `return err`. Client A sau 5 giây không thấy ACK nên gửi lại (Retry).

### Hậu quả:
- Khi Client A retry, bước `AcquireLock` trên Redis thấy key `uuid_123` đã tồn tại ➔ **Từ chối xử lý và báo trùng!**
- Tin nhắn bị bỏ qua mãi mãi trong 24h ➔ **Mất dữ liệu của người dùng (Data Loss)!**

---

## 2. Các cấp độ giải pháp xử lý Rollback Lock

### 🔹 Cấp độ 1: Chủ động Xóa Khóa (Active Release/Delete Lock)
Trong block `if err != nil` khi gọi `SaveMessage`, Worker lập tức gọi:
```go
u.idempotencyRepo.ReleaseLock(ctx, event.ClientMsgID) // Redis DEL key
```
- **Ưu điểm:** Đơn giản, giải phóng khóa ngay lập tức để lần retry kế tiếp có thể chạy lại ngay.
- **Điểm yếu:** Nếu Worker bị sập nguồn đột ngột (Crash / OOM) hoặc mạng sang Redis cũng bị đứt, lệnh `DEL` không thể thực thi.

---

### 🔹 Cấp độ 2 (Chuẩn Senior Production): Two-Phase State với Short-TTL Lock

Để giải quyết bài toán *"Worker sập nguồn trước khi kịp xóa lock"*, các hệ thống lớn áp dụng mô hình **Two-Phase Idempotency State**:

```
[Nhận Inbound Event]
        │
        ▼
BƯỚC 1: SET key "PROCESSING" NX EX 15s  (Khóa tạm thời với TTL cực ngắn)
        │
        ├──► Ghi vào Cassandra ──(THẤT BẠI)──► Không cần làm gì! (Khóa tự hủy sau 15s)
        │
        └──► Ghi vào Cassandra ──(THÀNH CÔNG)
                    │
                    ▼
BƯỚC 2: SET key "COMPLETED" XX EX 86400s (Gia hạn thành khóa vĩnh viễn 24h)
```

### Tại sao giải pháp Two-Phase TTL là tối thượng?
1. **Tự phục hồi mà không cần dọn dẹp (Self-Healing):** Dù Server có nổ tung hay đứt cáp, trạng thái `PROCESSING` sẽ tự động bốc hơi sau 15 giây nhờ cơ chế Expire nội tại của Redis.
2. **Không bao giờ xảy ra Deadlock:** Lần retry tiếp theo của Client (sau 15s) sẽ tự động chiếm được khóa mới và ghi lại bình thường.
3. **Bảo vệ toàn vẹn tuyệt đối:** Một khi đã chuyển sang `COMPLETED`, không một yêu cầu retry nào có thể ghi trùng vào Database được nữa.
````

## File: pkg/grpcclient/gateway_client_manager.go
````go
package grpcclient

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	pb "chat-system/pkg/proto"
	"chat-system/pkg/telemetry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GatewayClientManager quản lý danh sách và tái sử dụng (cache) các kết nối gRPC tới từng WS Gateway node
type GatewayClientManager struct {
	defaultPort int
	dnsSuffix   string
	clients     map[string]pb.WSGatewayServiceClient
	conns       map[string]*grpc.ClientConn
	mu          sync.RWMutex
}

// NewGatewayClientManager khởi tạo GatewayClientManager với cấu hình port mặc định và DNS suffix (cho K8s Headless Service)
func NewGatewayClientManager(defaultPort int, dnsSuffix string) *GatewayClientManager {
	if defaultPort <= 0 {
		defaultPort = 50051
	}
	return &GatewayClientManager{
		defaultPort: defaultPort,
		dnsSuffix:   dnsSuffix,
		clients:     make(map[string]pb.WSGatewayServiceClient),
		conns:       make(map[string]*grpc.ClientConn),
	}
}

// ResolveAddress phân giải nodeID thành địa chỉ kết nối gRPC hợp lệ
func (m *GatewayClientManager) ResolveAddress(nodeID string) string {
	if strings.Contains(nodeID, ":") {
		return nodeID
	}
	if m.dnsSuffix != "" {
		return fmt.Sprintf("%s%s:%d", nodeID, m.dnsSuffix, m.defaultPort)
	}
	return fmt.Sprintf("%s:%d", nodeID, m.defaultPort)
}

// GetClient trả về gRPC client kết nối tới WS Gateway node tương ứng (sử dụng cache hoặc dial mới nếu chưa có)
func (m *GatewayClientManager) GetClient(nodeID string) (pb.WSGatewayServiceClient, error) {
	if nodeID == "" {
		return nil, fmt.Errorf("nodeID cannot be empty")
	}

	m.mu.RLock()
	client, exists := m.clients[nodeID]
	m.mu.RUnlock()

	if exists {
		return client, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Double check sau khi acquire Write Lock
	if client, exists := m.clients[nodeID]; exists {
		return client, nil
	}

	addr := m.ResolveAddress(nodeID)
	dialCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		dialCtx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(telemetry.UnaryClientInterceptor(nodeID)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial WS Gateway gRPC at %s: %w", addr, err)
	}

	newClient := pb.NewWSGatewayServiceClient(conn)
	m.conns[nodeID] = conn
	m.clients[nodeID] = newClient

	log.Printf("[GatewayClientManager] Established gRPC connection to WS Gateway node '%s' at %s", nodeID, addr)
	return newClient, nil
}

// Close đóng tất cả các kết nối gRPC đang mở khi shutdown service
func (m *GatewayClientManager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var firstErr error
	for node, conn := range m.conns {
		if err := conn.Close(); err != nil {
			log.Printf("[GatewayClientManager] Error closing connection to %s: %v", node, err)
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	m.conns = make(map[string]*grpc.ClientConn)
	m.clients = make(map[string]pb.WSGatewayServiceClient)
	return firstErr
}
````

## File: pkg/nats/client.go
````go
package nats

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Connect establishes a robust connection to NATS server with automatic reconnects.
func Connect(natsURL string, clientName string) (*nats.Conn, error) {
	opts := []nats.Option{
		nats.Name(clientName),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("[NATS] Client %s disconnected: %v", clientName, err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("[NATS] Client %s reconnected to %s", clientName, nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("[NATS] Client %s connection closed permanently", clientName)
		}),
	}

	nc, err := nats.Connect(natsURL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS (%s): %w", natsURL, err)
	}
	return nc, nil
}
````

## File: pkg/nats/instrumented_publisher.go
````go
package nats

import (
	"context"
	"time"

	"chat-system/pkg/telemetry"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// PublisherConfig chứa cấu hình ngữ cảnh đo lường cho Publisher
type PublisherConfig struct {
	ServiceName string // ví dụ: "ws-gateway", "chat-engine"
	Stage       string // ví dụ: "inbound", "outbound_dispatch"
	EventType   string // ví dụ: "inbound", "outbound"
}

// InstrumentedPublisher là Decorator bọc quanh Publisher[T], tự động thu thập Trace & Metrics
type InstrumentedPublisher[T any] struct {
	inner *Publisher[T]
	cfg   PublisherConfig
}

// NewInstrumentedPublisher khởi tạo một InstrumentedPublisher bọc ngoài Publisher[T]
func NewInstrumentedPublisher[T any](inner *Publisher[T], cfg PublisherConfig) *InstrumentedPublisher[T] {
	if cfg.EventType == "" {
		cfg.EventType = cfg.Stage
	}
	return &InstrumentedPublisher[T]{
		inner: inner,
		cfg:   cfg,
	}
}

// Publish gửi dữ liệu vào default subject, tự động đo latency và message count
func (p *InstrumentedPublisher[T]) Publish(ctx context.Context, data T) error {
	start := time.Now()
	err := p.inner.Publish(ctx, data)
	duration := time.Since(start).Seconds()

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.MessageLatency.WithLabelValues("nats", p.cfg.Stage, status).Observe(duration)
	telemetry.MessagesProcessed.WithLabelValues(p.cfg.ServiceName, p.cfg.EventType, status).Inc()
	return err
}

// PublishToSubject gửi dữ liệu vào một subject cụ thể, tự động đo latency và message count
func (p *InstrumentedPublisher[T]) PublishToSubject(ctx context.Context, subject string, data T) error {
	start := time.Now()
	err := p.inner.PublishToSubject(ctx, subject, data)
	duration := time.Since(start).Seconds()

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.MessageLatency.WithLabelValues("nats", p.cfg.Stage, status).Observe(duration)
	telemetry.MessagesProcessed.WithLabelValues(p.cfg.ServiceName, p.cfg.EventType, status).Inc()
	return err
}

// PublishToNode gửi dữ liệu tới subject của một Gateway Node cụ thể, tự động tạo Span và đo cả DispatchDuration lẫn MessageLatency
func (p *InstrumentedPublisher[T]) PublishToNode(ctx context.Context, nodeID, subject string, data T) error {
	tracer := telemetry.Tracer(p.cfg.ServiceName)
	spanCtx, span := tracer.Start(ctx, "nats.DispatchToGateway",
		trace.WithAttributes(
			attribute.String("gateway_node", nodeID),
			attribute.String("subject", subject),
		),
	)
	defer span.End()

	start := time.Now()
	err := p.inner.PublishToSubject(spanCtx, subject, data)
	duration := time.Since(start).Seconds()

	status := "success"
	if err != nil {
		status = "error"
	}

	telemetry.DispatchDuration.WithLabelValues("nats", nodeID, status).Observe(duration)
	telemetry.MessageLatency.WithLabelValues("nats", p.cfg.Stage, status).Observe(duration)
	telemetry.MessagesProcessed.WithLabelValues(p.cfg.ServiceName, p.cfg.EventType, status).Inc()

	return err
}
````

## File: pkg/proto/ws_gateway_grpc.pb.go
````go
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.6.2
// - protoc             v7.35.0
// source: pkg/proto/ws_gateway.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WSGatewayService_PushMessageToUser_FullMethodName = "/gateway.WSGatewayService/PushMessageToUser"
)

// WSGatewayServiceClient is the client API for WSGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WSGatewayServiceClient interface {
	PushMessageToUser(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageResponse, error)
}

type wSGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWSGatewayServiceClient(cc grpc.ClientConnInterface) WSGatewayServiceClient {
	return &wSGatewayServiceClient{cc}
}

func (c *wSGatewayServiceClient) PushMessageToUser(ctx context.Context, in *PushMessageRequest, opts ...grpc.CallOption) (*PushMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushMessageResponse)
	err := c.cc.Invoke(ctx, WSGatewayService_PushMessageToUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WSGatewayServiceServer is the server API for WSGatewayService service.
// All implementations must embed UnimplementedWSGatewayServiceServer
// for forward compatibility.
type WSGatewayServiceServer interface {
	PushMessageToUser(context.Context, *PushMessageRequest) (*PushMessageResponse, error)
	mustEmbedUnimplementedWSGatewayServiceServer()
}

// UnimplementedWSGatewayServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWSGatewayServiceServer struct{}

func (UnimplementedWSGatewayServiceServer) PushMessageToUser(context.Context, *PushMessageRequest) (*PushMessageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method PushMessageToUser not implemented")
}
func (UnimplementedWSGatewayServiceServer) mustEmbedUnimplementedWSGatewayServiceServer() {}
func (UnimplementedWSGatewayServiceServer) testEmbeddedByValue()                          {}

// UnsafeWSGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WSGatewayServiceServer will
// result in compilation errors.
type UnsafeWSGatewayServiceServer interface {
	mustEmbedUnimplementedWSGatewayServiceServer()
}

func RegisterWSGatewayServiceServer(s grpc.ServiceRegistrar, srv WSGatewayServiceServer) {
	// If the following call panics, it indicates UnimplementedWSGatewayServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WSGatewayService_ServiceDesc, srv)
}

func _WSGatewayService_PushMessageToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WSGatewayServiceServer).PushMessageToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WSGatewayService_PushMessageToUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WSGatewayServiceServer).PushMessageToUser(ctx, req.(*PushMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WSGatewayService_ServiceDesc is the grpc.ServiceDesc for WSGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WSGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.WSGatewayService",
	HandlerType: (*WSGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMessageToUser",
			Handler:    _WSGatewayService_PushMessageToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/ws_gateway.proto",
}
````

## File: pkg/proto/ws_gateway.pb.go
````go
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.11
// 	protoc        v7.35.0
// source: pkg/proto/ws_gateway.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushMessageRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	MessageId      string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	ClientMsgId    string                 `protobuf:"bytes,2,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`
	ConversationId string                 `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	SenderId       string                 `protobuf:"bytes,4,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ReceiverId     string                 `protobuf:"bytes,5,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	Content        string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Type           string                 `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp      int64                  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Payload        []byte                 `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PushMessageRequest) Reset() {
	*x = PushMessageRequest{}
	mi := &file_pkg_proto_ws_gateway_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageRequest) ProtoMessage() {}

func (x *PushMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_ws_gateway_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageRequest.ProtoReflect.Descriptor instead.
func (*PushMessageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_ws_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PushMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *PushMessageRequest) GetClientMsgId() string {
	if x != nil {
		return x.ClientMsgId
	}
	return ""
}

func (x *PushMessageRequest) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *PushMessageRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *PushMessageRequest) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

func (x *PushMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PushMessageRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PushMessageRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PushMessageRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PushMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushMessageResponse) Reset() {
	*x = PushMessageResponse{}
	mi := &file_pkg_proto_ws_gateway_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageResponse) ProtoMessage() {}

func (x *PushMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_ws_gateway_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageResponse.ProtoReflect.Descriptor instead.
func (*PushMessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_ws_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PushMessageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_pkg_proto_ws_gateway_proto protoreflect.FileDescriptor

const file_pkg_proto_ws_gateway_proto_rawDesc = "" +
	"\n" +
	"\x1apkg/proto/ws_gateway.proto\x12\agateway\"\xa4\x02\n" +
	"\x12PushMessageRequest\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\"\n" +
	"\rclient_msg_id\x18\x02 \x01(\tR\vclientMsgId\x12'\n" +
	"\x0fconversation_id\x18\x03 \x01(\tR\x0econversationId\x12\x1b\n" +
	"\tsender_id\x18\x04 \x01(\tR\bsenderId\x12\x1f\n" +
	"\vreceiver_id\x18\x05 \x01(\tR\n" +
	"receiverId\x12\x18\n" +
	"\acontent\x18\x06 \x01(\tR\acontent\x12\x12\n" +
	"\x04type\x18\a \x01(\tR\x04type\x12\x1c\n" +
	"\ttimestamp\x18\b \x01(\x03R\ttimestamp\x12\x18\n" +
	"\apayload\x18\t \x01(\fR\apayload\"T\n" +
	"\x13PushMessageResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12#\n" +
	"\rerror_message\x18\x02 \x01(\tR\ferrorMessage2b\n" +
	"\x10WSGatewayService\x12N\n" +
	"\x11PushMessageToUser\x12\x1b.gateway.PushMessageRequest\x1a\x1c.gateway.PushMessageResponseB\x1dZ\x1bchat-system/pkg/proto/pb;pbb\x06proto3"

var (
	file_pkg_proto_ws_gateway_proto_rawDescOnce sync.Once
	file_pkg_proto_ws_gateway_proto_rawDescData []byte
)

func file_pkg_proto_ws_gateway_proto_rawDescGZIP() []byte {
	file_pkg_proto_ws_gateway_proto_rawDescOnce.Do(func() {
		file_pkg_proto_ws_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_ws_gateway_proto_rawDesc), len(file_pkg_proto_ws_gateway_proto_rawDesc)))
	})
	return file_pkg_proto_ws_gateway_proto_rawDescData
}

var file_pkg_proto_ws_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_ws_gateway_proto_goTypes = []any{
	(*PushMessageRequest)(nil),  // 0: gateway.PushMessageRequest
	(*PushMessageResponse)(nil), // 1: gateway.PushMessageResponse
}
var file_pkg_proto_ws_gateway_proto_depIdxs = []int32{
	0, // 0: gateway.WSGatewayService.PushMessageToUser:input_type -> gateway.PushMessageRequest
	1, // 1: gateway.WSGatewayService.PushMessageToUser:output_type -> gateway.PushMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_ws_gateway_proto_init() }
func file_pkg_proto_ws_gateway_proto_init() {
	if File_pkg_proto_ws_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_ws_gateway_proto_rawDesc), len(file_pkg_proto_ws_gateway_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_ws_gateway_proto_goTypes,
		DependencyIndexes: file_pkg_proto_ws_gateway_proto_depIdxs,
		MessageInfos:      file_pkg_proto_ws_gateway_proto_msgTypes,
	}.Build()
	File_pkg_proto_ws_gateway_proto = out.File
	file_pkg_proto_ws_gateway_proto_goTypes = nil
	file_pkg_proto_ws_gateway_proto_depIdxs = nil
}
````

## File: pkg/proto/ws_gateway.proto
````protobuf
syntax = "proto3";

package gateway;
option go_package = "chat-system/pkg/proto/pb;pb";

service WSGatewayService {
  rpc PushMessageToUser(PushMessageRequest) returns (PushMessageResponse);
}

message PushMessageRequest {
  string message_id = 1;
  string client_msg_id = 2;
  string conversation_id = 3;
  string sender_id = 4;
  string receiver_id = 5;
  string content = 6;
  string type = 7;
  int64 timestamp = 8;
  bytes payload = 9;
}

message PushMessageResponse {
  bool success = 1;
  string error_message = 2;
}
````

## File: pkg/telemetry/carrier.go
````go
package telemetry

import (
	"context"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/metadata"
)

// NATSHeaderCarrier adapts nats.Header to propagation.TextMapCarrier
type NATSHeaderCarrier nats.Header

func (c NATSHeaderCarrier) Get(key string) string {
	return nats.Header(c).Get(key)
}

func (c NATSHeaderCarrier) Set(key string, value string) {
	nats.Header(c).Set(key, value)
}

func (c NATSHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// InjectNATSTraceContext injects the OpenTelemetry TraceContext into nats.Msg Header
func InjectNATSTraceContext(ctx context.Context, msg *nats.Msg) {
	if msg.Header == nil {
		msg.Header = make(nats.Header)
	}
	otel.GetTextMapPropagator().Inject(ctx, NATSHeaderCarrier(msg.Header))
}

// ExtractNATSTraceContext extracts OpenTelemetry TraceContext from nats.Msg Header into a new context
func ExtractNATSTraceContext(ctx context.Context, msg *nats.Msg) context.Context {
	if msg.Header == nil {
		return ctx
	}
	return otel.GetTextMapPropagator().Extract(ctx, NATSHeaderCarrier(msg.Header))
}

// GRPCMetadataCarrier adapts metadata.MD to propagation.TextMapCarrier
type GRPCMetadataCarrier metadata.MD

func (c GRPCMetadataCarrier) Get(key string) string {
	vals := metadata.MD(c).Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func (c GRPCMetadataCarrier) Set(key string, value string) {
	metadata.MD(c).Set(key, value)
}

func (c GRPCMetadataCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// InjectGRPCTraceContext injects the trace context into outgoing gRPC context metadata
func InjectGRPCTraceContext(ctx context.Context) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}
	otel.GetTextMapPropagator().Inject(ctx, GRPCMetadataCarrier(md))
	return metadata.NewOutgoingContext(ctx, md)
}

// ExtractGRPCTraceContext extracts trace context from incoming gRPC context metadata
func ExtractGRPCTraceContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}
	return otel.GetTextMapPropagator().Extract(ctx, GRPCMetadataCarrier(md))
}
````

## File: pkg/telemetry/grpc_interceptor.go
````go
package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type successReporter interface {
	GetSuccess() bool
}

// UnaryClientInterceptor bọc ngoài tất cả các lệnh gRPC gọi từ Client (ví dụ chat-engine gọi ws-gateway)
// Tự động:
// 1. Gắn W3C TraceContext vào outgoing metadata
// 2. Tạo client span OpenTelemetry
// 3. Đo thời gian thực thi và ghi nhận vào DispatchDuration và MessageLatency
func UnaryClientInterceptor(targetNode string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		outCtx := InjectGRPCTraceContext(ctx)
		tracer := Tracer("grpc-client")
		spanCtx, span := tracer.Start(outCtx, fmt.Sprintf("grpc.call %s", method),
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(
				attribute.String("rpc.method", method),
				attribute.String("gateway_node", targetNode),
			),
		)
		defer span.End()

		start := time.Now()
		err := invoker(spanCtx, method, req, reply, cc, opts...)
		duration := time.Since(start).Seconds()

		status := "success"
		if err != nil {
			status = "error"
		} else if sr, ok := reply.(successReporter); ok && !sr.GetSuccess() {
			status = "rejected"
		}

		DispatchDuration.WithLabelValues("grpc", targetNode, status).Observe(duration)
		MessageLatency.WithLabelValues("grpc", "outbound_dispatch", status).Observe(duration)

		return err
	}
}

// UnaryServerInterceptor bọc ngoài tất cả các request gRPC gửi đến Server (ví dụ ws-gateway)
// Tự động:
// 1. Trích xuất W3C TraceContext từ incoming metadata
// 2. Tạo server span OpenTelemetry
// 3. Đo thời gian thực thi và ghi nhận vào MessageLatency
func UnaryServerInterceptor(serviceName, stage string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		trCtx := ExtractGRPCTraceContext(ctx)
		tracer := Tracer(serviceName)
		spanCtx, span := tracer.Start(trCtx, fmt.Sprintf("grpc.server %s", info.FullMethod),
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				attribute.String("rpc.method", info.FullMethod),
			),
		)
		defer span.End()

		start := time.Now()
		resp, err := handler(spanCtx, req)
		duration := time.Since(start).Seconds()

		status := "success"
		if err != nil {
			status = "error"
		}

		MessageLatency.WithLabelValues("grpc", stage, status).Observe(duration)
		return resp, err
	}
}
````

## File: pkg/telemetry/metrics.go
````go
package telemetry

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// High-resolution latency buckets: 100µs up to 10s
	latencyBuckets = []float64{
		0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025,
		0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0,
	}

	// MessageLatency tracks latency across different stages of message delivery
	MessageLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "chat",
			Subsystem: "system",
			Name:      "message_latency_seconds",
			Help:      "Latency of chat message processing broken down by mode and stage.",
			Buckets:   latencyBuckets,
		},
		[]string{"mode", "stage", "status"},
	)

	// DispatchDuration tracks the exact execution duration of gRPC vs NATS outbound dispatch
	DispatchDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "chat",
			Subsystem: "engine",
			Name:      "dispatch_duration_seconds",
			Help:      "Duration of outbound dispatch to gateway (gRPC vs NATS).",
			Buckets:   latencyBuckets,
		},
		[]string{"mode", "node_id", "status"},
	)

	// ActiveConnections tracks the number of active WebSocket connections on a gateway
	ActiveConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "chat",
			Subsystem: "gateway",
			Name:      "active_connections",
			Help:      "Current active WebSocket client connections on this node.",
		},
		[]string{"node_id"},
	)

	// MessagesProcessed tracks message volume
	MessagesProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "chat",
			Subsystem: "system",
			Name:      "messages_total",
			Help:      "Total number of messages processed.",
		},
		[]string{"service", "event_type", "status"},
	)

	// WorkerChannelDepth tracks the current number of pending jobs in a worker channel
	WorkerChannelDepth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "chat",
			Subsystem: "worker",
			Name:      "channel_depth",
			Help:      "Current number of pending items queued in the partitioned worker channel.",
		},
		[]string{"worker_id"},
	)

	// WorkerChannelCapacity tracks the buffer capacity of a worker channel
	WorkerChannelCapacity = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "chat",
			Subsystem: "worker",
			Name:      "channel_capacity",
			Help:      "Maximum buffer capacity of the partitioned worker channel.",
		},
		[]string{"worker_id"},
	)

	// WorkerChannelSaturation tracks the saturation percentage (depth / capacity)
	WorkerChannelSaturation = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "chat",
			Subsystem: "worker",
			Name:      "channel_saturation_ratio",
			Help:      "Saturation ratio (0.0 to 1.0) of the partitioned worker channel.",
		},
		[]string{"worker_id"},
	)

	// WorkerJobsTotal tracks the count of processed jobs per worker
	WorkerJobsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "chat",
			Subsystem: "worker",
			Name:      "jobs_total",
			Help:      "Total number of jobs executed by the worker pool.",
		},
		[]string{"worker_id", "status"},
	)

	// DatabaseLatency tracks individual DB operation latency (Cassandra / Redis)
	DatabaseLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "chat",
			Subsystem: "db",
			Name:      "operation_duration_seconds",
			Help:      "Duration of database operations (Cassandra / Redis).",
			Buckets:   latencyBuckets,
		},
		[]string{"system", "operation", "status"},
	)

	// ClientNetworkLatency tracks the network round-trip time (RTT) measured via client WebSocket ping/heartbeat
	ClientNetworkLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "chat",
			Subsystem: "gateway",
			Name:      "client_network_rtt_seconds",
			Help:      "Network round-trip latency measured via WebSocket heartbeat/ping.",
			Buckets:   latencyBuckets,
		},
		[]string{"node_id"},
	)
)

func init() {
	prometheus.MustRegister(
		MessageLatency,
		DispatchDuration,
		ActiveConnections,
		MessagesProcessed,
		WorkerChannelDepth,
		WorkerChannelCapacity,
		WorkerChannelSaturation,
		WorkerJobsTotal,
		DatabaseLatency,
		ClientNetworkLatency,
	)
}

// StartMetricsServer starts an HTTP endpoint on the given port to serve Prometheus/VictoriaMetrics
func StartMetricsServer(port int) *http.Server {
	if port <= 0 {
		return nil
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("[Metrics] Prometheus/VictoriaMetrics metrics listening on :%d/metrics", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[Metrics] HTTP server error: %v", err)
		}
	}()

	return srv
}
````

## File: pkg/telemetry/middleware.go
````go
package telemetry

import (
	"context"
	"time"
)

// HandlerConfig chứa cấu hình đo lường cho các message handler
type HandlerConfig struct {
	Service   string // ví dụ: "chat-engine", để trống nếu không đếm MessagesProcessed
	Mode      string // ví dụ: "engine", "nats", "grpc"
	Stage     string // ví dụ: "worker_process", "gateway_delivery"
	EventType string // ví dụ: "inbound", mặc định lấy theo Stage nếu để trống
}

// InstrumentHandler bọc ngoài hàm xử lý message (ví dụ Subscriber callback hoặc Consumer handler)
// để tự động đo đạc MessageLatency và MessagesProcessed mà không làm ô nhiễm business logic
func InstrumentHandler[T any](cfg HandlerConfig, handler func(ctx context.Context, data T) error) func(ctx context.Context, data T) error {
	if cfg.EventType == "" {
		cfg.EventType = cfg.Stage
	}
	return func(ctx context.Context, data T) error {
		start := time.Now()
		err := handler(ctx, data)
		duration := time.Since(start).Seconds()

		status := "success"
		if err != nil {
			status = "error"
		}

		MessageLatency.WithLabelValues(cfg.Mode, cfg.Stage, status).Observe(duration)
		if cfg.Service != "" {
			MessagesProcessed.WithLabelValues(cfg.Service, cfg.EventType, status).Inc()
		}
		return err
	}
}
````

## File: pkg/telemetry/profiler.go
````go
package telemetry

import (
	"fmt"
	"log"
	"runtime"

	"github.com/grafana/pyroscope-go"
)

// ProfilerConfig configures Pyroscope continuous profiling
type ProfilerConfig struct {
	ServerAddress string // e.g. "http://pyroscope.observability:4040"
	ApplicationName string // e.g. "chat-system.ws-gateway"
	NodeID          string
	Environment     string
	Disabled        bool
}

// InitProfiler starts the Grafana Pyroscope continuous profiling agent
func InitProfiler(cfg ProfilerConfig) (func() error, error) {
	if cfg.Disabled || cfg.ServerAddress == "" {
		log.Printf("[Profiler] Continuous profiling disabled or server address empty.")
		return func() error { return nil }, nil
	}

	// Configure runtime profile rates for contention & blocking
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(10000)

	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: cfg.ApplicationName,
		ServerAddress:   cfg.ServerAddress,
		Logger:          nil, // standard logger
		Tags: map[string]string{
			"env":     cfg.Environment,
			"node_id": cfg.NodeID,
		},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start pyroscope profiler: %w", err)
	}

	log.Printf("[Profiler] Continuous profiling active -> Pyroscope: %s (App: %s)",
		cfg.ServerAddress, cfg.ApplicationName)

	return func() error {
		log.Printf("[Profiler] Stopping Pyroscope profiler for %s...", cfg.ApplicationName)
		return profiler.Stop()
	}, nil
}
````

## File: pkg/telemetry/setup_test.go
````go
package telemetry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"chat-system/pkg/telemetry"
)

func TestSetup_Disabled(t *testing.T) {
	cleanup, err := telemetry.Setup(context.Background(), telemetry.SetupConfig{
		ServiceName:      "test-service",
		ServiceVersion:   "1.0.0",
		NodeID:           "test-node-1",
		MetricsPort:      0, // Không start port
		DisableTracing:   true,
		DisableProfiling: true,
	})
	if err != nil {
		t.Fatalf("expected nil error when telemetry disabled, got: %v", err)
	}
	if cleanup == nil {
		t.Fatal("expected cleanup function, got nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := cleanup(ctx); err != nil {
		t.Fatalf("expected clean shutdown, got: %v", err)
	}
}

func TestInstrumentHandler(t *testing.T) {
	cfg := telemetry.HandlerConfig{
		Service:   "test-service",
		Mode:      "nats",
		Stage:     "test_stage",
		EventType: "test_event",
	}

	// Case 1: Handler success
	successHandler := telemetry.InstrumentHandler(cfg, func(ctx context.Context, msg string) error {
		if msg != "hello" {
			t.Fatalf("unexpected message: %s", msg)
		}
		return nil
	})

	if err := successHandler(context.Background(), "hello"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}

	// Case 2: Handler error
	errExpected := errors.New("boom")
	errorHandler := telemetry.InstrumentHandler(cfg, func(ctx context.Context, msg string) error {
		return errExpected
	})

	if err := errorHandler(context.Background(), "fail"); !errors.Is(err, errExpected) {
		t.Fatalf("expected error %v, got: %v", errExpected, err)
	}
}
````

## File: pkg/telemetry/setup.go
````go
package telemetry

import (
	"context"
	"log"
)

// SetupConfig đóng gói toàn bộ cấu hình cho 3 trụ cột Telemetry
type SetupConfig struct {
	ServiceName      string
	ServiceVersion   string
	Environment      string // Mặc định là "production" nếu để trống
	NodeID           string
	CollectorTarget  string
	MetricsPort      int
	ProfilerServer   string
	DisableTracing   bool
	DisableProfiling bool
}

// Setup là hàm Facade khởi tạo đồng bộ cả 3 hệ thống:
// 1. HTTP Metrics Server (:metrics_port) cho Prometheus/VictoriaMetrics
// 2. OpenTelemetry TracerProvider đẩy sang OTel Collector (:collector_target)
// 3. Grafana Pyroscope Profiler Agent đẩy sang Pyroscope Server (:server_address)
//
// Trả về duy nhất 1 hàm cleanup để thực hiện Graceful Shutdown theo đúng thứ tự an toàn:
// Profiler -> Metrics Server -> TracerProvider (Flush toàn bộ trace spans còn tồn đọng).
func Setup(ctx context.Context, cfg SetupConfig) (func(context.Context) error, error) {
	if cfg.Environment == "" {
		cfg.Environment = "production"
	}
	if cfg.ServiceVersion == "" {
		cfg.ServiceVersion = "1.0.0"
	}

	// 1. Khởi động Metrics Server
	metricsServer := StartMetricsServer(cfg.MetricsPort)

	// 2. Khởi tạo OpenTelemetry Tracing
	shutdownTracer, errTracer := InitTracer(ctx, Config{
		ServiceName:     cfg.ServiceName,
		ServiceVersion:  cfg.ServiceVersion,
		Environment:     cfg.Environment,
		NodeID:          cfg.NodeID,
		CollectorTarget: cfg.CollectorTarget,
		Disabled:        cfg.DisableTracing,
	})
	if errTracer != nil {
		log.Printf("[Telemetry] Warning: failed to initialize tracer: %v", errTracer)
	}

	// 3. Khởi tạo Continuous Profiler (Grafana Pyroscope)
	stopProfiler, errProfiler := InitProfiler(ProfilerConfig{
		ServerAddress:   cfg.ProfilerServer,
		ApplicationName: "chat-system." + cfg.ServiceName,
		NodeID:          cfg.NodeID,
		Environment:     cfg.Environment,
		Disabled:        cfg.DisableProfiling,
	})
	if errProfiler != nil {
		log.Printf("[Telemetry] Warning: failed to start profiler: %v", errProfiler)
	}

	cleanup := func(shutdownCtx context.Context) error {
		log.Printf("[Telemetry] Shutting down telemetry subsystems for %s...", cfg.ServiceName)
		if stopProfiler != nil {
			_ = stopProfiler()
		}
		if metricsServer != nil {
			_ = metricsServer.Shutdown(shutdownCtx)
		}
		if shutdownTracer != nil {
			_ = shutdownTracer(shutdownCtx)
		}
		return nil
	}

	if errTracer != nil {
		return cleanup, errTracer
	}
	if errProfiler != nil {
		return cleanup, errProfiler
	}

	return cleanup, nil
}
````

## File: pkg/telemetry/tracer.go
````go
package telemetry

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config defines the configuration for telemetry setup
type Config struct {
	ServiceName     string
	ServiceVersion  string
	Environment     string
	NodeID          string
	CollectorTarget string // e.g. "localhost:4317" or "otel-collector.observability:4317"
	Disabled        bool
}

// InitTracer initializes OpenTelemetry TracerProvider and registers global propagator
func InitTracer(ctx context.Context, cfg Config) (func(context.Context) error, error) {
	if cfg.Disabled || cfg.CollectorTarget == "" {
		log.Printf("[Telemetry] Tracing disabled or collector endpoint empty. Running in NoOp mode.")
		return func(ctx context.Context) error { return nil }, nil
	}

	// 1. Setup Resource Attributes
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.ServiceName),
			semconv.ServiceVersionKey.String(cfg.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(cfg.Environment),
			semconv.ServiceInstanceIDKey.String(cfg.NodeID),
		),
		resource.WithHost(),
		resource.WithProcess(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 2. Setup OTLP gRPC Exporter
	exporterCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	exporter, err := otlptracegrpc.New(exporterCtx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.CollectorTarget),
		otlptracegrpc.WithDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OTLP gRPC trace exporter: %w", err)
	}

	// 3. Batch Span Processor with high-throughput tuning to prevent OOM
	bsp := sdktrace.NewBatchSpanProcessor(exporter,
		sdktrace.WithMaxQueueSize(16384),
		sdktrace.WithMaxExportBatchSize(2048),
		sdktrace.WithBatchTimeout(1*time.Second),
		sdktrace.WithExportTimeout(5*time.Second),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(tp)

	// 4. Setup Global W3C TraceContext Propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	log.Printf("[Telemetry] OpenTelemetry Tracing initialized -> Collector: %s (Service: %s)",
		cfg.CollectorTarget, cfg.ServiceName)

	// Return graceful shutdown function
	return func(shutdownCtx context.Context) error {
		log.Printf("[Telemetry] Flushing and shutting down TracerProvider for %s...", cfg.ServiceName)
		return tp.Shutdown(shutdownCtx)
	}, nil
}

// Tracer returns a named tracer instance
func Tracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
````

## File: pkg/worker/pool_test.go
````go
package worker

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type sampleTask struct {
	Key     string
	Payload string
}

func TestPartitionedPool_DeterministicRouting(t *testing.T) {
	pool, err := NewPartitionedPool(Config[sampleTask]{
		NumWorkers: 16,
		BufferSize: 100,
		KeyExtractor: func(item sampleTask) string {
			return item.Key
		},
		Handler: func(ctx context.Context, item sampleTask) error {
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	key := "room-alpha-123"
	expectedIdx := pool.getWorkerIndex(key)

	for i := 0; i < 50; i++ {
		idx := pool.getWorkerIndex(key)
		if idx != expectedIdx {
			t.Fatalf("Expected worker %d, got %d on iteration %d", expectedIdx, idx, i)
		}
	}
}

func TestPartitionedPool_PanicRecovery(t *testing.T) {
	var processed atomic.Int32

	pool, err := NewPartitionedPool(Config[sampleTask]{
		NumWorkers: 4,
		BufferSize: 100,
		KeyExtractor: func(item sampleTask) string {
			return item.Key
		},
		Handler: func(ctx context.Context, item sampleTask) error {
			if item.Payload == "panic-payload" {
				panic("simulated poison pill panic")
			}
			processed.Add(1)
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	pool.Start()

	ctx := context.Background()

	// Gửi task gây panic
	_ = pool.Submit(ctx, sampleTask{Key: "room-1", Payload: "panic-payload"})

	// Gửi tiếp các task bình thường cùng key
	for i := 0; i < 10; i++ {
		_ = pool.Submit(ctx, sampleTask{Key: "room-1", Payload: fmt.Sprintf("good-%d", i)})
	}

	time.Sleep(100 * time.Millisecond)

	drainCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := pool.Stop(drainCtx); err != nil {
		t.Fatalf("Failed to stop pool: %v", err)
	}

	if processed.Load() != 10 {
		t.Fatalf("Expected 10 processed messages, got %d", processed.Load())
	}
}

func TestPartitionedPool_GracefulDrain(t *testing.T) {
	var processed atomic.Int32
	var mu sync.Mutex
	order := make([]string, 0)

	pool, err := NewPartitionedPool(Config[sampleTask]{
		NumWorkers: 8,
		BufferSize: 500,
		KeyExtractor: func(item sampleTask) string {
			return item.Key
		},
		Handler: func(ctx context.Context, item sampleTask) error {
			time.Sleep(2 * time.Millisecond)
			processed.Add(1)
			mu.Lock()
			order = append(order, item.Payload)
			mu.Unlock()
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	pool.Start()

	totalMsgs := 100
	ctx := context.Background()

	for i := 0; i < totalMsgs; i++ {
		err := pool.Submit(ctx, sampleTask{
			Key:     fmt.Sprintf("conv-%d", i%5),
			Payload: fmt.Sprintf("msg-%d", i),
		})
		if err != nil {
			t.Fatalf("Failed to submit message: %v", err)
		}
	}

	drainCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Stop(drainCtx); err != nil {
		t.Fatalf("Stop timed out or failed: %v", err)
	}

	if int(processed.Load()) != totalMsgs {
		t.Fatalf("Expected %d messages drained, got %d", totalMsgs, processed.Load())
	}
}
````

## File: pkg/worker/pool.go
````go
package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"

	"chat-system/pkg/telemetry"
)

const (
	DefaultNumWorkers = 32
	DefaultBufferSize = 1024
)

var (
	ErrPoolClosed   = errors.New("worker pool is closed")
	ErrPoolDraining = errors.New("worker pool timed out while draining")
)

// HandlerFunc định nghĩa hàm xử lý sự kiện trong Worker theo kiểu generic
type HandlerFunc[T any] func(ctx context.Context, item T) error

// KeyExtractorFunc định nghĩa hàm trích xuất routing key (ví dụ: conversation_id) từ item
type KeyExtractorFunc[T any] func(item T) string

// Job đại diện cho một tác vụ kèm Trace Context được chuyển tiếp sang Worker channel
type Job[T any] struct {
	Ctx  context.Context
	Item T
}

// Worker đại diện cho một goroutine xử lý độc lập trên channel riêng
type Worker[T any] struct {
	id      int
	idStr   string
	jobCh   chan Job[T]
	handler HandlerFunc[T]
}

// PartitionedPool quản lý N workers với cơ chế băm partition theo key
type PartitionedPool[T any] struct {
	numWorkers   int
	bufferSize   int
	workers      []*Worker[T]
	handler      HandlerFunc[T]
	keyExtractor KeyExtractorFunc[T]
	wg           sync.WaitGroup
	isClosed     atomic.Bool
}

// Config cấu hình generic cho PartitionedPool
type Config[T any] struct {
	NumWorkers   int
	BufferSize   int
	KeyExtractor KeyExtractorFunc[T]
	Handler      HandlerFunc[T]
}

// NewPartitionedPool khởi tạo một PartitionedPool generic
func NewPartitionedPool[T any](cfg Config[T]) (*PartitionedPool[T], error) {
	if cfg.NumWorkers <= 0 {
		cfg.NumWorkers = DefaultNumWorkers
	}
	if cfg.BufferSize <= 0 {
		cfg.BufferSize = DefaultBufferSize
	}
	if cfg.Handler == nil {
		return nil, errors.New("worker pool handler cannot be nil")
	}
	if cfg.KeyExtractor == nil {
		return nil, errors.New("worker pool key extractor cannot be nil")
	}

	pool := &PartitionedPool[T]{
		numWorkers:   cfg.NumWorkers,
		bufferSize:   cfg.BufferSize,
		workers:      make([]*Worker[T], cfg.NumWorkers),
		handler:      cfg.Handler,
		keyExtractor: cfg.KeyExtractor,
	}

	for i := 0; i < cfg.NumWorkers; i++ {
		idStr := strconv.Itoa(i)
		pool.workers[i] = &Worker[T]{
			id:      i,
			idStr:   idStr,
			jobCh:   make(chan Job[T], cfg.BufferSize),
			handler: cfg.Handler,
		}

		// Initialize channel metrics
		telemetry.WorkerChannelCapacity.WithLabelValues(idStr).Set(float64(cfg.BufferSize))
		telemetry.WorkerChannelDepth.WithLabelValues(idStr).Set(0)
		telemetry.WorkerChannelSaturation.WithLabelValues(idStr).Set(0)
	}

	return pool, nil
}

// Start khởi chạy toàn bộ worker goroutines
func (p *PartitionedPool[T]) Start() {
	for i := 0; i < p.numWorkers; i++ {
		p.wg.Add(1)
		go p.runWorker(p.workers[i])
	}
	log.Printf("[WorkerPool] Started %d partitioned workers (buffer_size=%d per worker)", p.numWorkers, p.bufferSize)
}

// runWorker thực thi vòng lặp nhận job của từng worker kèm cơ chế Panic Recovery per-job
func (p *PartitionedPool[T]) runWorker(w *Worker[T]) {
	defer p.wg.Done()

	for job := range w.jobCh {
		// Update metrics after taking a job from queue
		depth := len(w.jobCh)
		telemetry.WorkerChannelDepth.WithLabelValues(w.idStr).Set(float64(depth))
		telemetry.WorkerChannelSaturation.WithLabelValues(w.idStr).Set(float64(depth) / float64(p.bufferSize))

		p.processJobWithRecovery(w, job)
	}
}

// processJobWithRecovery bọc xử lý job trong defer recover để cô lập sự cố (Panic Isolation)
func (p *PartitionedPool[T]) processJobWithRecovery(w *Worker[T], job Job[T]) {
	defer func() {
		if r := recover(); r != nil {
			telemetry.WorkerJobsTotal.WithLabelValues(w.idStr, "panic").Inc()
			log.Printf("[Worker %d] CRITICAL: Panic recovered while processing job: %v", w.id, r)
		}
	}()

	if err := w.handler(job.Ctx, job.Item); err != nil {
		telemetry.WorkerJobsTotal.WithLabelValues(w.idStr, "error").Inc()
		log.Printf("[Worker %d] Error processing job: %v", w.id, err)
		return
	}

	telemetry.WorkerJobsTotal.WithLabelValues(w.idStr, "success").Inc()
}

// Submit định tuyến item vào channel của worker tương ứng theo hash(keyExtractor(item))
func (p *PartitionedPool[T]) Submit(ctx context.Context, item T) error {
	if p.isClosed.Load() {
		return ErrPoolClosed
	}

	key := p.keyExtractor(item)
	workerIdx := p.getWorkerIndex(key)
	targetWorker := p.workers[workerIdx]

	job := Job[T]{
		Ctx:  ctx,
		Item: item,
	}

	select {
	case targetWorker.jobCh <- job:
		depth := len(targetWorker.jobCh)
		telemetry.WorkerChannelDepth.WithLabelValues(targetWorker.idStr).Set(float64(depth))
		telemetry.WorkerChannelSaturation.WithLabelValues(targetWorker.idStr).Set(float64(depth) / float64(p.bufferSize))
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Stop đóng toàn bộ worker channels và chờ xả cạn backlog (Graceful Drain) kèm timeout context
func (p *PartitionedPool[T]) Stop(ctx context.Context) error {
	if !p.isClosed.CompareAndSwap(false, true) {
		return nil // Đã đóng trước đó
	}

	log.Printf("[WorkerPool] Initiating graceful drain for %d workers...", p.numWorkers)

	// Đóng từng channel để báo hiệu cho workers không còn job mới
	for _, w := range p.workers {
		close(w.jobCh)
	}

	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("[WorkerPool] All workers drained and stopped cleanly.")
		return nil
	case <-ctx.Done():
		log.Printf("[WorkerPool] WARNING: Drain timeout reached: %v", ctx.Err())
		return fmt.Errorf("%w: %v", ErrPoolDraining, ctx.Err())
	}
}

// NumWorkers trả về số lượng workers
func (p *PartitionedPool[T]) NumWorkers() int {
	return p.numWorkers
}

// getWorkerIndex tính chỉ số worker bằng thuật toán FNV-1a 32-bit
func (p *PartitionedPool[T]) getWorkerIndex(key string) int {
	hash := fnv32a(key)
	return int(hash % uint32(p.numWorkers))
}

// fnv32a triển khai thuật toán băm FNV-1a 32-bit (Non-cryptographic, zero-alloc)
func fnv32a(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}
````

## File: plan/3_week_observability_benchmark_stress_plan.md
````markdown
# Kế hoạch Hiện thực hóa: 3 Tuần Nâng tầm Kiến trúc & Đo lường Hệ thống Chat Phân tán

## 🎯 Mục tiêu Tổng quan
Chuyển hóa hệ thống từ giai đoạn "chạy được ở Happy Path" sang **Production-Ready Distributed System** với khả năng quan sát toàn diện (Deep Observability), phân tích định lượng hiệu năng (Benchmarking & Chaos Engineering), và chịu tải tới hạn (Stress Testing & Continuous Profiling).

---

## 🏛️ Góc nhìn Kiến trúc & 4 Lăng kính Sự cố (Architecture Analysis)

Trước khi đi vào từng tuần, toàn bộ kế hoạch được thiết kế dựa trên 4 Lăng kính Sự cố:

1. **Lăng kính Thời gian & Bất đồng bộ (Time & Concurrency Lens):**
   - Độ trễ nội bộ (Internal Latency) giữa các bước: Cassandra I/O RTT vs Redis Lock RTT. Khi tải tăng đột biến, hàng đợi trong Go Channel của `PartitionedPool` sẽ bị phình to (Backlog saturation), làm trôi lệch P99 của toàn bộ hệ thống.
2. **Lăng kính Sự cố Mạng & Node (Network & Partial Failure Lens):**
   - Sự khác biệt căn bản giữa **Point-to-Point Synchronous (gRPC)** và **Asynchronous Pub/Sub (NATS)** khi có Jitter/Network Latency: gRPC gây giữ kết nối, tăng Goroutine & RAM ở Gateway; NATS đệm trong buffer nhưng có nguy cơ rớt tin nếu buffer đầy.
3. **Lăng kính Giới hạn Dữ liệu & Tải đột biến (Data Boundary & Load Lens):**
   - Khi tải lên 10,000 msg/s, GC Pressure từ JSON serialization và gocql frame allocations sẽ chiếm dụng CPU, gây stop-the-world spikes.
4. **Lăng kính Trạng thái & Vòng đời (State Machine & Lifecycle Lens):**
   - Đo lường chính xác từng trạng thái trong trace lifecycle: Gateway Inbound $\to$ Redis Idempotency $\to$ Cassandra Commit $\to$ Sender Ack $\to$ Outbound Dispatch.

---

## 📅 Lộ trình Triển khai Chi tiết

### 📌 TUẦN 1: Bịt điểm mù Observability (Deep Observability & 4 Golden Signals)

#### 1. Thêm Child Spans (OpenTelemetry / Tempo) cho Database & Caching
* **Vấn đề hiện tại:** Trace hiện tại chỉ bao quát tầng Inbound $\to$ Usecase tổng quát. Khi P99 tăng vọt, không thể xác định điểm nghẽn nằm ở Redis Lock, Cassandra Write, hay Outbound Dispatch.
* **Giải pháp:**
  - Bổ sung OTel Child Spans có ngữ cảnh Semantic Conventions:
    + `db.system: cassandra`, `db.operation: insert`, `db.statement: INSERT INTO messages...` trong `CassandraMessageRepository.SaveMessage` và `GetMessagesByConversation`.
    + `db.system: redis`, `db.operation: setnx`, `db.redis.key: idempotency:msg:...` trong `redisIdempotencyRepository.AcquireLock` và `ReleaseLock`.
    + `db.system: redis`, `db.operation: hgetall/hlen` trong `redisPresenceReader`.
  - Tự động ghi nhận Span Status Error và `span.RecordError(err)` khi các lệnh I/O gặp timeout/lỗi.

#### 2. Metric đo Độ sâu và Độ bão hòa Channel trong `PartitionedPool`
* **Vấn đề hiện tại:** `PartitionedPool` sử dụng buffered Go channels (`jobCh`). Khi 1 partition bị nghẽn (Hot partition) hoặc downstream DB chậm, channel bị đầy dần dẫn đến block/timeout mà không có metric cảnh báo trước.
* **Giải pháp:**
  - Khai báo Prometheus Metrics trong `pkg/telemetry/metrics.go`:
    + `chat_worker_channel_depth` (GaugeVec by `worker_id`): `len(w.jobCh)`
    + `chat_worker_channel_capacity` (GaugeVec by `worker_id`): `cap(w.jobCh)`
    + `chat_worker_channel_saturation_ratio` (GaugeVec by `worker_id`): `len(w.jobCh) / cap(w.jobCh)`
    + `chat_worker_jobs_total` (CounterVec by `worker_id`, `status`): Đếm số job xử lý thành công / lỗi / panic.
  - Tích hợp hàm cập nhật metric theo chu kỳ nhẹ (tick interval) hoặc cập nhật trực tiếp tại `Submit` / `runWorker`.

#### 3. Xây dựng Grafana Dashboard "4 Biểu đồ Vàng" (Golden Signals)
* **Vị trí:** `deployments/configs/grafana-dashboards/chat-golden-signals.json` và đăng ký tự động qua `grafana-datasources.yaml` / dashboard provisioning.
* **4 Panel Trọng tâm:**
  1. **Throughput (Traffic):** `sum(rate(chat_messages_total[1m])) by (service, event_type)` (Đo RPS hệ thống).
  2. **Latency Percentiles (P50, P95, P99):** `histogram_quantile(0.99, sum(rate(chat_message_latency_seconds_bucket[1m])) by (le, stage))` (Phân rã theo stage: Idempotency, Cassandra Insert, Realtime Dispatch).
  3. **Error Rate (%):** Tỷ lệ lỗi so với tổng request `(sum(rate(chat_messages_total{status="error"}[1m])) / sum(rate(chat_messages_total[1m]))) * 100`.
  4. **Worker Channel Saturation (%):** `max(chat_worker_channel_depth / chat_worker_channel_capacity) * 100` (Phát hiện nghẽn worker pool).

---

### 📌 TUẦN 2: So găng gRPC vs NATS Broker Delivery (Benchmarking & Chaos Testing)

#### 1. Kịch bản Benchmark gRPC Delivery Mode
* Cấu hình `DISPATCH_MODE=grpc` trên cụm `chat-engine`.
* Chạy `client-simulator` với kịch bản multi-node (ví dụ: 1,000 -> 3,000 active bots trên nhiều Gateway instances).
* Thu thập số liệu: P50, P95, P99 Latency (qua Prometheus & Grafana), CPU usage của `chat-engine`, RAM/Goroutine count của `ws-gateway`.

#### 2. Kịch bản Benchmark NATS Broker Delivery Mode
* Cấu hình `DISPATCH_MODE=broker` (NATS Core / JetStream routing theo `chat.gateway.{node_id}`).
* Chạy cùng một profile tải từ `client-simulator`.
* Ghi nhận và so sánh độ trễ P99, khả năng buffering, CPU và RAM giữa 2 phương thức.

#### 3. Chaos Engineering: Bơm độ trễ mạng (Chaos Mesh Latency Injection 100ms)
* Áp dụng `01-network-delay-grpc.yaml` và `02-network-delay-nats.yaml` trong thư mục `deployments/k8s/chaos/` (Inject 100ms latency giữa `chat-engine` và `ws-gateway`).
* **Quan sát hiện tượng:**
  - **Mode gRPC:** Hiện tượng connection pool exhaustion, goroutines tích tụ ở `chat-engine`, P99 tăng dốc đứng, lan truyền ngược (cascading failure / backpressure) làm tắc nghẽn `PartitionedPool`.
  - **Mode NATS:** NATS client buffer đóng vai trò đệm async, Goroutines ở `chat-engine` không bị giữ; tuy nhiên cần giám sát RAM tiêu thụ của NATS broker và nguy cơ chậm tin nhắn tới client cuối (End-to-End Latency).

#### 4. Tổng kết & Đóng gói Tài liệu So sánh Thực tế
* Tạo tài liệu chuẩn mực tại: `knowledge/system_design/04_grpc_vs_nats_realtime_delivery_tradeoffs.md`.
* Nội dung gồm: Bảng so sánh định lượng (RPS, Latency P99, CPU/RAM, Fault Tolerance), đồ thị so sánh, phân tích các trường hợp nên chọn gRPC vs NATS trong thực tế (Architecture Decision Record).

---

### 📌 TUẦN 3: Đẩy tải tới ngưỡng vỡ & Continuous Profiling (Stress Testing & Performance Tuning)

#### 1. Stress Testing tới Ngưỡng Vỡ (5,000 $\to$ 10,000+ msg/sec)
* Tăng dần số lượng bots và tần suất phát tin trong `client-simulator` hoặc script `k6`:
  - Bước 1: 2,000 msg/s (Baseline).
  - Bước 2: 5,000 msg/s (High load).
  - Bước 3: 10,000 msg/s (Stress point).
* Xác định chính xác điểm bắt đầu xuất hiện tình trạng drop tin nhắn (Channel overflow, Client disconnect, HTTP/WebSocket drop rate).

#### 2. Bắt "Thủ phạm" CPU & Memory Contention bằng Grafana Pyroscope
* Quan sát Flamegraph trên Pyroscope UI (`http://localhost:4040`):
  - **CPU Profile:** Phân tích tỷ lệ CPU đốt vào:
    + `encoding/json` serialization/deserialization.
    + `gocql` query formatting, frame decoding, reflection.
    + `go-redis` pipeline & networking overhead.
    + Runtime scheduler & garbage collection (GC pause).
  - **Mutex & Block Profile:**
    + Mutex contention trong Go Channel hoặc Sync Maps ở Gateway SessionManager.
    + Lock contention trong Redis connection pool.

#### 3. Hiện thực hóa Tối ưu hóa (Fix Bottleneck)
* Dựa trên kết quả Flamegraph thực tế, thực hiện tối ưu điểm nghẽn lớn nhất:
  - *Nếu là JSON parsing:* Chuyển đổi sang `sonic`, `easyjson` hoặc zero-allocation protobuf encoding cho luồng hot path.
  - *Nếu là Gocql/Cassandra I/O:* Tối ưu prepared statements caching, tuning gocql connection pool size, batching hợp lý.
  - *Nếu là Redis RTT/Contention:* Áp dụng Redis Pipelining / MGet / Lua script cho việc kiểm tra Idempotency + Presence trong cùng 1 network round-trip.
* Chạy lại bài test tải để nghiệm thu kết quả cải thiện (Before vs After metrics).
````

## File: plan/benchmark_runbook_grpc_vs_nats.md
````markdown
# Cẩm Nang Kiểm Thử Hiệu Năng & Chaos Engineering (Benchmarking Runbook)
## So Sánh Thực Chiến: gRPC (Point-to-Point) vs Message Broker (NATS) & Đo Lường Network Delay

Tài liệu này cung cấp quy trình chuẩn hóa từng bước để tiến hành đo tải (Stress Test) và thử nghiệm phá vỡ hệ thống (Chaos Engineering) trên cụm Kubernetes, nhằm kiểm chứng tính đúng đắn của Dashboard quan sát (Observability) và phân tích định lượng ưu/nhược điểm sống còn giữa hai mô hình điều phối: **gRPC** và **NATS Broker**.

---

## 1. Bản Đồ Kiến Trúc: 2 Chặng Mạng Độc Lập

Để quan sát hệ thống chính xác, cần phân định rõ 2 chặng truyền thông trong hệ thống Chat:

```text
[Client Bots] <=== (CHẶNG 1: Edge Network) ===> [WS Gateway] <=== (CHẶNG 2: Internal Cluster) ===> [Chat Engine / NATS]
                   Internet / 4G / Wi-Fi                             gRPC vs NATS Broker
                   (Đo bằng WebSocket Ping RTT)                      (Nơi gRPC và NATS đối đầu)
```

1. **Chặng 1 (Mạng Biên - Edge):** Kết nối WebSocket giữa Client và WS Gateway qua Ingress NGINX. Chặng này dùng để đo Network Delay thuần túy đưa lên **Panel 1.3** và CLI Simulator.
2. **Chặng 2 (Mạng Nội Bộ Backend):** Kết nối giữa WS Gateway và Chat Engine. Đây là chiến trường quyết định sự khác biệt về khả năng chịu tải và cách ly sự cố giữa **gRPC (Đồng bộ)** và **NATS (Bất đồng bộ)**.

---

## 2. Thông Số Tải Chuẩn (Workload Profile)

Cấu hình trong file `deployments/k8s/09-client-simulator.yaml`:

```yaml
data:
  SIM_BOTS: "500"          # 500 Bots kết nối đồng thời qua WebSocket
  SIM_CONVS: "10"          # Mỗi bot mở 10 cuộc hội thoại 1-1 chéo nhau
  SIM_INTERVAL: "100ms"    # Tần suất gửi: 100ms/tin (10 tin/giây/bot)
  SIM_DURATION: "2m"       # Thời lượng mỗi đợt test: 2 phút
```
* **Tổng Throughput phát sinh:** `500 bots * 10 msgs/s = 5,000 messages/giây`.

---

## 3. Quy Trình 3 Bước Thực Chiến Chi Tiết

### BƯỚC 0: Chuẩn Bị & Cập Nhật Code Lên Cụm (Chạy 1 lần)
Trước khi bắt đầu, đảm bảo image mới nhất đã được build và nạp vào cụm Kind:

```powershell
# Build và nạp Gateway mới (chứa handler HEARTBEAT_ACK & metric RTT)
make build-load-gw

# Build và nạp Simulator mới (chứa Jitter Ping & Tracker phân tách)
make build-load-sim
```

* **Địa chỉ Grafana:** `http://localhost:3000` (User: `admin` / Pass: `admin`).
* **Dashboard theo dõi:** `Chat System - 5 Golden Dimensions (gRPC vs NATS)`.

---

### PHẦN 1: Kiểm Chứng Đo Lường Network Delay (Chặng 1: Client <-> Gateway)

**Mục tiêu:** Xác thực Panel 1.3 và CLI Simulator bắt được chính xác độ trễ mạng biên khi người dùng dùng mạng 4G chập chờn.

1. **Khởi động Simulator:**
   ```powershell
   make sim-up
   ```

2. **Xem Dashboard trực tiếp trên Terminal:**
   ```powershell
   make logs-sim
   ```
   *Quan sát ban đầu:* Mục `Network Delay / RTT` dao động ở mức mạng nội bộ lý tưởng (`~1.0ms - 2.5ms`).

3. **Mở tab Terminal thứ 2, kích hoạt Chaos độ trễ mạng di động:**
   ```powershell
   make chaos-mobile-delay
   # Kịch bản: Inject 80ms latency + 25ms jitter giữa client-simulator và nginx-gateway
   ```

4. **Hiện tượng quan sát:**
   * **Trên Terminal Simulator (`make logs-sim`):**
     - Dòng `Network Delay / RTT` lập tức nhảy vọt lên mức **`~80.00ms - 105.00ms`**.
     - Dòng `Ước tính thời gian xử lý Backend (Median E2E - Network)` vẫn ổn định ở mức thấp (`~5ms - 7ms`), chứng minh backend không bị nghẽn mà nguyên nhân do mạng người dùng.
   * **Trên Grafana Dashboard:**
     - **Panel 1.3 (Client Network Delay / Edge RTT):** Đường P50, P95, P99 dựng đứng lên tương ứng `~80 - 105ms`.

5. **Dọn dẹp sau Phần 1:**
   ```powershell
   make chaos-clean
   make sim-down
   ```

---

### PHẦN 2: Trận Chiến Sinh Tử: gRPC vs NATS Broker (Chặng 2)

**Mục tiêu:** Chứng minh hiện tượng sụp đổ dây chuyền (Cascading Failure / Head-of-line Blocking) của gRPC khi có mạng chậm, và khả năng miễn nhiễm hoàn toàn của NATS Broker.

---

#### 🥊 Vòng A: Chế Độ gRPC Delivery (Point-to-Point Synchronous)

1. **Cấu hình gRPC Mode:**
   - Mở file `deployments/k8s/01-configmap-secrets.yaml`, đảm bảo:
     ```yaml
     DELIVERY_MODE: "grpc"
     ```
   - Nạp lại cấu hình:
     ```powershell
     make reload
     ```

2. **Khởi động tải 5,000 msg/s:**
   ```powershell
   make sim-up
   ```

3. **Giai đoạn A1 (30 giây đầu - Đường chuẩn Happy Path):**
   - Quan sát Grafana:
     - **Panel 1.1 (Dispatch Latency):** P50 rất tốt (~`2ms - 3ms`).
     - **Panel 2.1 (Throughput):** Đạt đỉnh cao nhất (`~4,500 - 5,000 msgs/s`).
     - **Panel 4.1 (Worker Saturation):** Rất thấp (`< 15%`).

4. **Giai đoạn A2 (Bơm lỗi mạng vào gRPC):**
   ```powershell
   make chaos-delay-grpc
   # Kịch bản: Inject 100ms latency giữa chat-engine và ws-gateway
   ```
   *Hiện tượng sụp đổ dây chuyền (The Collapse):*
   - **Panel 4.1 (Worker Pool Saturation):** Nhảy vọt lên **`100% (Bão hòa kịch khung)`**. Toàn bộ worker goroutine bị giữ chân (block) chờ phản hồi RPC từ Gateway.
   - **Panel 4.2 (Total Pending Backlog):** Hàng đợi channel bị tắc nghẽn, tồn đọng hàng ngàn jobs chưa xử lý.
   - **Panel 2.1 (Throughput):** Tụt dốc không phanh từ `5,000 msgs/s` xuống còn vài trăm msgs/s.
   - **Panel 3.1 & 3.2 (Error & Rejection):** Xuất hiện lỗi gRPC timeout/unavailable.

5. **Dọn dẹp Vòng A:**
   ```powershell
   make chaos-clean
   make sim-down
   ```

---

#### 🥊 Vòng B: Chế Độ NATS Broker Delivery (Asynchronous Decoupled)

1. **Chuyển sang Broker Mode:**
   - Mở file `deployments/k8s/01-configmap-secrets.yaml`, sửa thành:
     ```yaml
     DELIVERY_MODE: "broker"
     ```
   - Nạp lại cấu hình:
     ```powershell
     make reload
     ```

2. **Khởi động tải 5,000 msg/s:**
   ```powershell
   make sim-up
   ```

3. **Giai đoạn B1 (30 giây đầu - Đường chuẩn NATS):**
   - Quan sát Grafana:
     - **Panel 1.1 (Dispatch Latency):** P50 ở mức `~4ms - 5ms` (cao hơn gRPC khoảng 1-2ms do đi qua broker trung gian).
     - **Panel 2.1 (Throughput):** Ổn định ở mức `~4,500 - 5,000 msgs/s`.
     - **Panel 4.1 (Worker Saturation):** Ở mức thấp (`< 15%`).

4. **Giai đoạn B2 (Bơm lỗi mạng vào NATS):**
   ```powershell
   make chaos-delay-nats
   # Kịch bản: Inject 100ms latency giữa ws-gateway và nats broker
   ```
   *Hiện tượng vững như bàn thạch (Resilient Behavior):*
   - **Panel 4.1 (Worker Pool Saturation):** **VẪN NẰM IM DƯỚI ĐÁY `< 15%`!** Worker của Chat Engine chỉ mất vài micro-giây (`~50us`) để publish tin vào NATS rồi quay sang nhận tin mới ngay lập tức. Lõi Chat Engine hoàn toàn không bị ảnh hưởng.
   - **Panel 2.1 (Throughput):** Giữ vững thông lượng, không hề bị rớt thảm hại như gRPC.
   - Tin nhắn được xếp hàng an toàn trong NATS Buffer/Queue chờ Gateway kéo về khi mạng thông thoáng trở lại.

5. **Dọn dẹp kết thúc benchmark:**
   ```powershell
   make chaos-clean
   make sim-down
   ```

---

## 4. Bảng Tổng Kết Đối Chiếu (Architectural Scorecard)

| Tiêu Chí So Sánh | gRPC Mode (Synchronous Point-to-Point) | NATS Broker Mode (Asynchronous Queue) |
| :--- | :--- | :--- |
| **Độ trễ P50 lúc bình thường** | **Cực thấp (~1 - 2ms)**, tối ưu nhất | ~3 - 5ms (chậm hơn 1-2ms do hop qua broker) |
| **Mức tiêu hao tài nguyên đệm** | Zero-buffer, không cần quản lý broker | Cần bộ nhớ RAM để duy trì hàng đợi broker |
| **Khi mạng Downstream lag 100ms**| **Worker Pool nghẽn 100%**, Throughput sập | **Worker Pool < 15%**, thông lượng giữ vững |
| **Khi Gateway Pod bị Crash** | Lập tức văng lỗi RPC, rớt tin nhắn | Tin nhắn chờ trong queue, Pod mới sống lại lấy tiếp |
| **Mức độ phụ thuộc (Coupling)** | **Khớp nối cứng (Tight Coupling)** | **Tách rời hoàn toàn (Loosely Coupled)** |
| **Khuyến nghị kiến trúc** | Phù hợp hệ thống nội bộ LAN có SLA mạng cực cao | **Bắt buộc cho hệ thống quy mô lớn, tải cao, đa vùng** |

---

## 5. Cheatsheet Lệnh Nhanh (Quick Reference)

```powershell
# --- Bật / Tắt Simulator ---
make sim-up            # Khởi động simulator phát tải
make sim-down          # Dừng simulator
make logs-sim          # Xem dashboard trực tiếp

# --- Nạp Cấu Hình ---
make reload            # Áp dụng thay đổi DELIVERY_MODE

# --- Kích Hoạt Chaos ---
make chaos-mobile-delay   # Test Panel 1.3: Delay mạng Client (Chặng 1)
make chaos-delay-grpc     # Test gRPC: Delay mạng Engine <-> Gateway (Chặng 2)
make chaos-delay-nats     # Test NATS: Delay mạng Gateway <-> NATS (Chặng 2)
make chaos-kill-gw        # Test Pod Kill: Đột tử 1 Pod Gateway

# --- Dọn Dẹp Chaos ---
make chaos-clean       # Xóa toàn bộ kịch bản lỗi, khôi phục mạng bình thường
```
````

## File: services/api-service/.dockerignore
````
node_modules
dist
.git
.gitignore
*.md
Dockerfile
.dockerignore
npm-debug.log
````

## File: services/api-service/.env.example
````
PORT=3000
JWT_SECRET=your_super_secret_jwt_key_here
JWT_EXPIRES_IN=10m


DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgrespassword
DB_DATABASE=chat_db
````

## File: services/api-service/.gitignore
````
# compiled output
/dist
/node_modules
/build

# Logs
logs
*.log
npm-debug.log*
pnpm-debug.log*
yarn-debug.log*
yarn-error.log*
lerna-debug.log*

# OS
.DS_Store

# Tests
/coverage
/.nyc_output

# IDEs and editors
/.idea
.project
.classpath
.c9/
*.launch
.settings/
*.sublime-workspace

# IDE - VSCode
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json

# dotenv environment variable files
.env
.env.development.local
.env.test.local
.env.production.local
.env.local

# temp directory
.temp
.tmp

# Runtime data
pids
*.pid
*.seed
*.pid.lock

# Diagnostic reports (https://nodejs.org/api/report.html)
report.[0-9]*.[0-9]*.[0-9]*.[0-9]*.json
````

## File: services/api-service/cmd/main.go
````go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-service/internal/config"
	deliveryHttp "api-service/internal/delivery/http"
	"api-service/internal/repository"
	"api-service/internal/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Println("=================================================================")
	log.Println("           🚀 KHỞI ĐỘNG CHAT SYSTEM API SERVICE (GOLANG)        ")
	log.Println("=================================================================")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[Config] Failed to load config: %v", err)
	}

	log.Printf("[Config] Connecting to PostgreSQL at %s:%d (DB: %s)...", cfg.DBHost, cfg.DBPort, cfg.DBName)

	gormLogLevel := logger.Warn
	if os.Getenv("GIN_MODE") != "release" {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		log.Fatalf("[Database] Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("[Database] Failed to get underlying sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("[Database] PostgreSQL connection pool initialized successfully!")

	// 1. Repositories
	userRepo := repository.NewUserRepository(db)
	friendRepo := repository.NewFriendRepository(db)
	convoRepo := repository.NewConversationRepository(db)

	// 2. Usecases
	authUsecase := usecase.NewAuthUsecase(cfg, userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	friendUsecase := usecase.NewFriendUsecase(friendRepo, userRepo)
	convoUsecase := usecase.NewConversationUsecase(convoRepo, userRepo)

	// 3. Handlers
	handlers := &deliveryHttp.Handlers{
		Auth:         deliveryHttp.NewAuthHandler(authUsecase),
		User:         deliveryHttp.NewUserHandler(userUsecase),
		Friend:       deliveryHttp.NewFriendHandler(friendUsecase),
		Conversation: deliveryHttp.NewConversationHandler(convoUsecase),
	}

	// 4. HTTP Router
	router := deliveryHttp.SetupRouter(authUsecase, handlers)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("[HTTP] API Service listening on port %d...", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[HTTP] Listen error: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[HTTP] Shutting down API Service...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[HTTP] Server forced to shutdown: %v", err)
	}

	_ = sqlDB.Close()
	log.Println("[HTTP] API Service stopped cleanly.")
}
````

## File: services/api-service/go.mod
````
module api-service

go 1.25.0

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.55.0
	gorm.io/driver/postgres v1.5.7
	gorm.io/gorm v1.25.10
)

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.4 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	github.com/stretchr/testify v1.12.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
````

## File: services/api-service/internal/config/config.go
````go
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         int
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	JWTExpiresIn time.Duration
}

func LoadConfig() (*Config, error) {
	// Nạp .env nếu có (không báo lỗi nếu không tìm thấy, vì chạy k8s nạp qua env vars)
	_ = godotenv.Load()

	port := getEnvAsInt("PORT", 3000)
	dbPort := getEnvAsInt("DB_PORT", 5432)

	jwtExpMin := getEnvAsInt("JWT_EXPIRES_MINUTES", 15)
	if envExp := os.Getenv("JWT_EXPIRES_IN"); envExp != "" {
		if d, err := time.ParseDuration(envExp); err == nil {
			jwtExpMin = int(d.Minutes())
		}
	}

	cfg := &Config{
		Port:         port,
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       dbPort,
		DBUser:       getEnv("DB_USERNAME", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "postgrespassword"),
		DBName:       getEnv("DB_DATABASE", "chat_db"),
		JWTSecret:    getEnv("JWT_SECRET", "default_jwt_secret_chat_system"),
		JWTExpiresIn: time.Duration(jwtExpMin) * time.Minute,
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}
````

## File: services/api-service/internal/delivery/http/auth_handler.go
````go
package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authUsecase.Register(c.Request.Context(), req.Username, req.Password, req.Name, req.Phone, req.Address)
	if err != nil {
		if errors.Is(err, usecase.ErrUserAlreadyExists) {
			middleware.SendError(c, http.StatusConflict, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusCreated, "Register successfull", dto.ToUserResponse(user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.authUsecase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			middleware.SendError(c, http.StatusUnauthorized, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Login successfull", tokens)
}
````

## File: services/api-service/internal/delivery/http/conversation_handler.go
````go
package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/domain"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationHandler struct {
	convoUsecase usecase.ConversationUsecase
}

func NewConversationHandler(convoUsecase usecase.ConversationUsecase) *ConversationHandler {
	return &ConversationHandler{convoUsecase: convoUsecase}
}

func (h *ConversationHandler) GetMyConversations(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convos, err := h.convoUsecase.GetUserConversations(c.Request.Context(), userID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := make([]*dto.ConversationResponse, 0, len(convos))
	for i := range convos {
		resp = append(resp, dto.ToConversationResponse(&convos[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy danh sách cuộc hội thoại thành công", resp)
}

func (h *ConversationHandler) GetOrCreateDirect(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	var req dto.CreateDirectConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	partnerID, err := uuid.Parse(req.PartnerID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "partnerId không hợp lệ")
		return
	}

	convo, err := h.convoUsecase.GetOrCreateDirectConversation(c.Request.Context(), userID, partnerID)
	if err != nil {
		if errors.Is(err, usecase.ErrTargetNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Khởi tạo cuộc hội thoại 1-1 thành công", dto.ToConversationResponse(convo))
}

func (h *ConversationHandler) CreateGroup(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	var req dto.CreateGroupConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	memberUUIDs := make([]uuid.UUID, 0, len(req.MemberIDs))
	for _, m := range req.MemberIDs {
		if uid, err := uuid.Parse(m); err == nil {
			memberUUIDs = append(memberUUIDs, uid)
		}
	}

	convo, err := h.convoUsecase.CreateGroupConversation(c.Request.Context(), userID, req.Name, req.IconURL, memberUUIDs)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusCreated, "Tạo nhóm chat thành công", dto.ToConversationResponse(convo))
}

func (h *ConversationHandler) GetConversationByID(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	idStr := c.Param("id")
	convoID, err := uuid.Parse(idStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	convo, err := h.convoUsecase.GetConversationByID(c.Request.Context(), convoID, userID)
	if err != nil {
		if errors.Is(err, usecase.ErrConversationNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy chi tiết cuộc hội thoại thành công", dto.ToConversationResponse(convo))
}

func (h *ConversationHandler) AddMember(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convoIDStr := c.Param("id")
	convoID, err := uuid.Parse(convoIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	var req dto.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	targetMemberID, err := uuid.Parse(req.UserID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "userId không hợp lệ")
		return
	}

	role := req.Role
	if role == "" {
		role = domain.MemberRoleMember
	}

	member, err := h.convoUsecase.AddMember(c.Request.Context(), convoID, userID, targetMemberID, role)
	if err != nil {
		if errors.Is(err, usecase.ErrConversationNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotMember) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := dto.ConversationMemberResponse{
		ID:             member.ID.String(),
		ConversationID: member.ConversationID.String(),
		UserID:         member.UserID.String(),
		Role:           string(member.Role),
		JoinedAt:       member.JoinedAt,
		User:           dto.ToUserResponse(member.User),
	}

	middleware.SendSuccess(c, http.StatusCreated, "Thêm thành viên vào nhóm thành công", resp)
}

func (h *ConversationHandler) RemoveMember(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convoIDStr := c.Param("id")
	convoID, err := uuid.Parse(convoIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	targetIDStr := c.Param("targetUserId")
	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "targetUserId không hợp lệ")
		return
	}

	success, err := h.convoUsecase.RemoveMember(c.Request.Context(), convoID, userID, targetID)
	if err != nil {
		if errors.Is(err, usecase.ErrConversationNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotMember) || errors.Is(err, usecase.ErrNotAdmin) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Xóa thành viên hoặc rời nhóm thành công", gin.H{"success": success})
}

func (h *ConversationHandler) UpdateReadReceipt(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convoIDStr := c.Param("id")
	convoID, err := uuid.Parse(convoIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	var req dto.UpdateReadReceiptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	receipt, err := h.convoUsecase.UpdateReadReceipt(c.Request.Context(), convoID, userID, req.LastMessageSeen)
	if err != nil {
		if errors.Is(err, usecase.ErrNotMember) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := gin.H{
		"userId":          receipt.UserID.String(),
		"lastMessageSeen": receipt.LastMessageSeen,
		"lastReadAt":      receipt.LastReadAt,
	}

	middleware.SendSuccess(c, http.StatusOK, "Cập nhật trạng thái đã đọc thành công", resp)
}
````

## File: services/api-service/internal/delivery/http/dto/dto.go
````go
package dto

import (
	"time"

	"api-service/internal/domain"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username string  `json:"userName" binding:"required"`
	Password string  `json:"password" binding:"required,min=6"`
	Name     string  `json:"name" binding:"required"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

type LoginRequest struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Phone     *string   `json:"phone,omitempty"`
	Address   *string   `json:"address,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponse(u *domain.User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:        u.ID.String(),
		Username:  u.Username,
		Name:      u.Name,
		Phone:     u.Phone,
		Address:   u.Address,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type FriendResponse struct {
	ID          string        `json:"id"`
	RequesterID string        `json:"requesterId"`
	AddresseeID string        `json:"addresseeId"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Requester   *UserResponse `json:"requester,omitempty"`
	Addressee   *UserResponse `json:"addressee,omitempty"`
}

func ToFriendResponse(f *domain.Friend) *FriendResponse {
	if f == nil {
		return nil
	}
	return &FriendResponse{
		ID:          f.ID.String(),
		RequesterID: f.RequesterID.String(),
		AddresseeID: f.AddresseeID.String(),
		Status:      string(f.Status),
		CreatedAt:   f.CreatedAt,
		UpdatedAt:   f.UpdatedAt,
		Requester:   ToUserResponse(f.Requester),
		Addressee:   ToUserResponse(f.Addressee),
	}
}

type CreateDirectConversationRequest struct {
	PartnerID string `json:"partnerId" binding:"required"`
}

type CreateGroupConversationRequest struct {
	Name      string   `json:"name" binding:"required"`
	IconURL   *string  `json:"iconUrl"`
	MemberIDs []string `json:"memberIds"`
}

type AddMemberRequest struct {
	UserID string            `json:"userId" binding:"required"`
	Role   domain.MemberRole `json:"role"`
}

type UpdateReadReceiptRequest struct {
	LastMessageSeen string `json:"lastMessageSeen" binding:"required"`
}

type ConversationMemberResponse struct {
	ID             string        `json:"id"`
	ConversationID string        `json:"conversationId"`
	UserID         string        `json:"userId"`
	Role           string        `json:"role"`
	JoinedAt       time.Time     `json:"joinedAt"`
	User           *UserResponse `json:"user,omitempty"`
}

type ConversationResponse struct {
	ID        string                       `json:"id"`
	Name      *string                      `json:"name,omitempty"`
	Type      string                       `json:"type"`
	IconURL   *string                      `json:"iconUrl,omitempty"`
	CreatedBy *string                      `json:"createdBy,omitempty"`
	CreatedAt time.Time                    `json:"createdAt"`
	UpdatedAt time.Time                    `json:"updatedAt"`
	Members   []ConversationMemberResponse `json:"members"`
}

func ToConversationResponse(c *domain.Conversation) *ConversationResponse {
	if c == nil {
		return nil
	}
	var createdByStr *string
	if c.CreatedBy != nil && *c.CreatedBy != uuid.Nil {
		s := c.CreatedBy.String()
		createdByStr = &s
	}

	members := make([]ConversationMemberResponse, 0, len(c.Members))
	for _, m := range c.Members {
		members = append(members, ConversationMemberResponse{
			ID:             m.ID.String(),
			ConversationID: m.ConversationID.String(),
			UserID:         m.UserID.String(),
			Role:           string(m.Role),
			JoinedAt:       m.JoinedAt,
			User:           ToUserResponse(m.User),
		})
	}

	return &ConversationResponse{
		ID:        c.ID.String(),
		Name:      c.Name,
		Type:      string(c.Type),
		IconURL:   c.IconURL,
		CreatedBy: createdByStr,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Members:   members,
	}
}
````

## File: services/api-service/internal/delivery/http/friend_handler.go
````go
package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FriendHandler struct {
	friendUsecase usecase.FriendUsecase
}

func NewFriendHandler(friendUsecase usecase.FriendUsecase) *FriendHandler {
	return &FriendHandler{friendUsecase: friendUsecase}
}

func (h *FriendHandler) GetFriendsList(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	list, err := h.friendUsecase.GetFriendsList(c.Request.Context(), userID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := make([]*dto.FriendResponse, 0, len(list))
	for i := range list {
		resp = append(resp, dto.ToFriendResponse(&list[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy danh sách bạn bè thành công", resp)
}

func (h *FriendHandler) GetPendingRequests(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	list, err := h.friendUsecase.GetPendingRequests(c.Request.Context(), userID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := make([]*dto.FriendResponse, 0, len(list))
	for i := range list {
		resp = append(resp, dto.ToFriendResponse(&list[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy danh sách lời mời kết bạn thành công", resp)
}

func (h *FriendHandler) SendFriendRequest(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	targetIDStr := c.Param("targetUserId")
	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID người dùng không hợp lệ")
		return
	}

	result, err := h.friendUsecase.SendFriendRequest(c.Request.Context(), userID, targetID)
	if err != nil {
		if errors.Is(err, usecase.ErrTargetNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusCreated, "Gửi lời mời kết bạn thành công", dto.ToFriendResponse(result))
}

func (h *FriendHandler) AcceptFriendRequest(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	reqIDStr := c.Param("requestId")
	reqID, err := uuid.Parse(reqIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID lời mời không hợp lệ")
		return
	}

	result, err := h.friendUsecase.AcceptFriendRequest(c.Request.Context(), userID, reqID)
	if err != nil {
		if errors.Is(err, usecase.ErrRelationshipNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotAddressee) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Đã chấp nhận lời mời kết bạn", dto.ToFriendResponse(result))
}

func (h *FriendHandler) DeclineFriendRequest(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	reqIDStr := c.Param("requestId")
	reqID, err := uuid.Parse(reqIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID lời mời không hợp lệ")
		return
	}

	result, err := h.friendUsecase.DeclineFriendRequest(c.Request.Context(), userID, reqID)
	if err != nil {
		if errors.Is(err, usecase.ErrRelationshipNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotAddressee) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Đã từ chối lời mời kết bạn", dto.ToFriendResponse(result))
}

func (h *FriendHandler) BlockUser(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	targetIDStr := c.Param("targetUserId")
	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID người dùng không hợp lệ")
		return
	}

	result, err := h.friendUsecase.BlockUser(c.Request.Context(), userID, targetID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Đã chặn người dùng", dto.ToFriendResponse(result))
}

func (h *FriendHandler) Unfriend(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	relIDStr := c.Param("relationshipId")
	relID, err := uuid.Parse(relIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID mối quan hệ không hợp lệ")
		return
	}

	success, err := h.friendUsecase.UnfriendOrCancel(c.Request.Context(), userID, relID)
	if err != nil {
		if errors.Is(err, usecase.ErrRelationshipNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Hủy kết bạn hoặc hủy yêu cầu thành công", gin.H{"success": success})
}
````

## File: services/api-service/internal/delivery/http/middleware/auth_middleware.go
````go
package middleware

import (
	"net/http"
	"strings"

	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CtxUserIDKey   = "userID"
	CtxUsernameKey = "username"
)

func AuthMiddleware(authUsecase usecase.AuthUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			SendError(c, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			SendError(c, http.StatusUnauthorized, "Invalid authorization header format (must be Bearer token)")
			return
		}

		tokenStr := parts[1]
		claims, err := authUsecase.ValidateToken(tokenStr)
		if err != nil {
			SendError(c, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		userUUID, err := uuid.Parse(claims.UserID)
		if err != nil {
			SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
			return
		}

		c.Set(CtxUserIDKey, userUUID)
		c.Set(CtxUsernameKey, claims.Username)
		c.Next()
	}
}

func GetCurrentUserID(c *gin.Context) uuid.UUID {
	val, ok := c.Get(CtxUserIDKey)
	if !ok {
		return uuid.Nil
	}
	return val.(uuid.UUID)
}
````

## File: services/api-service/internal/delivery/http/middleware/response.go
````go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func SendError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	})
}

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	SendError(c, http.StatusBadRequest, err.Error())
}
````

## File: services/api-service/internal/delivery/http/router.go
````go
package http

import (
	"net/http"

	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth         *AuthHandler
	User         *UserHandler
	Friend       *FriendHandler
	Conversation *ConversationHandler
}

func SetupRouter(authUsecase usecase.AuthUsecase, handlers *Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// CORS nhẹ
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 1. Auth routes (public)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Auth.Register)
		authGroup.POST("/login", handlers.Auth.Login)
	}

	// Protected routes
	api := r.Group("/")
	api.Use(middleware.AuthMiddleware(authUsecase))
	{
		// 2. Users routes
		users := api.Group("/users")
		{
			users.GET("/me", handlers.User.GetProfile)
			users.PATCH("/me", handlers.User.UpdateProfile)
			users.GET("/search", handlers.User.SearchUsers)
			users.GET("/:id", handlers.User.GetUserByID)
		}

		// 3. Friends routes
		friends := api.Group("/friends")
		{
			friends.GET("", handlers.Friend.GetFriendsList)
			friends.GET("/requests", handlers.Friend.GetPendingRequests)
			friends.POST("/request/:targetUserId", handlers.Friend.SendFriendRequest)
			friends.PATCH("/requests/:requestId/accept", handlers.Friend.AcceptFriendRequest)
			friends.PATCH("/requests/:requestId/decline", handlers.Friend.DeclineFriendRequest)
			friends.POST("/block/:targetUserId", handlers.Friend.BlockUser)
			friends.DELETE("/:relationshipId", handlers.Friend.Unfriend)
		}

		// 4. Conversations routes
		convos := api.Group("/conversations")
		{
			convos.GET("", handlers.Conversation.GetMyConversations)
			convos.POST("/direct", handlers.Conversation.GetOrCreateDirect)
			convos.POST("/group", handlers.Conversation.CreateGroup)
			convos.GET("/:id", handlers.Conversation.GetConversationByID)
			convos.POST("/:id/members", handlers.Conversation.AddMember)
			convos.DELETE("/:id/members/:targetUserId", handlers.Conversation.RemoveMember)
			convos.POST("/:id/read", handlers.Conversation.UpdateReadReceipt)
		}
	}

	return r
}
````

## File: services/api-service/internal/delivery/http/user_handler.go
````go
package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	user, err := h.userUsecase.GetProfile(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy thông tin cá nhân thành công", dto.ToUserResponse(user))
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := h.userUsecase.UpdateProfile(c.Request.Context(), userID, req.Name, req.Phone, req.Address)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Cập nhật thông tin cá nhân thành công", dto.ToUserResponse(updated))
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	query := c.Query("q")
	users, err := h.userUsecase.SearchUsers(c.Request.Context(), query)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	respList := make([]*dto.UserResponse, 0, len(users))
	for i := range users {
		respList = append(respList, dto.ToUserResponse(&users[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Tìm kiếm người dùng thành công", respList)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	targetID, err := uuid.Parse(idStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID người dùng không hợp lệ")
		return
	}

	user, err := h.userUsecase.FindByID(c.Request.Context(), targetID)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			middleware.SendError(c, http.StatusNotFound, "Không tìm thấy người dùng")
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy thông tin người dùng thành công", dto.ToUserResponse(user))
}
````

## File: services/api-service/internal/domain/conversation.go
````go
package domain

import (
	"time"

	"github.com/google/uuid"
)

type ConversationType string

const (
	ConversationTypeDirect ConversationType = "direct"
	ConversationTypeGroup  ConversationType = "group"
	ConversationTypeClan   ConversationType = "clan"
	ConversationTypeRegion ConversationType = "region"
	ConversationTypeGlobal ConversationType = "global"
)

type MemberRole string

const (
	MemberRoleAdmin  MemberRole = "admin"
	MemberRoleMember MemberRole = "member"
)

type Conversation struct {
	ID        uuid.UUID        `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Name      *string          `gorm:"size:100;column:name" json:"name,omitempty"`
	Type      ConversationType `gorm:"type:conversation_type_enum;not null;default:'direct';column:type" json:"type"`
	IconURL   *string          `gorm:"column:iconUrl" json:"iconUrl,omitempty"`
	CreatedBy *uuid.UUID       `gorm:"type:uuid;column:createdBy" json:"createdBy,omitempty"`
	CreatedAt time.Time        `gorm:"column:createdAt;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time        `gorm:"column:updatedAt;not null;autoUpdateTime" json:"updatedAt"`

	Members []ConversationMember `gorm:"foreignKey:ConversationID" json:"members,omitempty"`
}

func (Conversation) TableName() string {
	return "conversations"
}

type ConversationMember struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	ConversationID uuid.UUID  `gorm:"type:uuid;not null;column:conversationId;index" json:"conversationId"`
	UserID         uuid.UUID  `gorm:"type:uuid;not null;column:userId;index" json:"userId"`
	Role           MemberRole `gorm:"type:member_role_enum;not null;default:'member';column:role" json:"role"`
	JoinedAt       time.Time  `gorm:"column:joinedAt;not null;autoCreateTime" json:"joinedAt"`

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (ConversationMember) TableName() string {
	return "conversation_members"
}

type ConversationReadReceipt struct {
	ConversationID  uuid.UUID `gorm:"type:uuid;primaryKey;column:conversationId" json:"conversationId"`
	UserID          uuid.UUID `gorm:"type:uuid;primaryKey;column:userId" json:"userId"`
	LastMessageSeen string    `gorm:"size:64;not null;column:lastMessageSeen" json:"lastMessageSeen"`
	LastReadAt      time.Time `gorm:"column:lastReadAt;not null;autoCreateTime" json:"lastReadAt"`
}

func (ConversationReadReceipt) TableName() string {
	return "conversation_read_receipts"
}
````

## File: services/api-service/internal/domain/friend.go
````go
package domain

import (
	"time"

	"github.com/google/uuid"
)

type FriendStatus string

const (
	FriendStatusPending  FriendStatus = "pending"
	FriendStatusAccepted FriendStatus = "accepted"
	FriendStatusDeclined FriendStatus = "declined"
	FriendStatusBlocked  FriendStatus = "blocked"
)

type Friend struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	RequesterID uuid.UUID    `gorm:"type:uuid;not null;column:requesterId;index" json:"requesterId"`
	AddresseeID uuid.UUID    `gorm:"type:uuid;not null;column:addresseeId;index" json:"addresseeId"`
	Status      FriendStatus `gorm:"type:friend_status_enum;not null;default:'pending';column:status" json:"status"`
	CreatedAt   time.Time    `gorm:"column:createdAt;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time    `gorm:"column:updatedAt;not null;autoUpdateTime" json:"updatedAt"`

	Requester *User `gorm:"foreignKey:RequesterID" json:"requester,omitempty"`
	Addressee *User `gorm:"foreignKey:AddresseeID" json:"addressee,omitempty"`
}

func (Friend) TableName() string {
	return "friends"
}
````

## File: services/api-service/internal/domain/user.go
````go
package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Username     string         `gorm:"size:50;uniqueIndex;not null;column:username" json:"username"`
	PasswordHash string         `gorm:"not null;column:passwordHash" json:"-"`
	Name         string         `gorm:"not null;column:name" json:"name"`
	Phone        *string        `gorm:"size:20;column:phone" json:"phone,omitempty"`
	Address      *string        `gorm:"column:address" json:"address,omitempty"`
	CreatedAt    time.Time      `gorm:"column:createdAt;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updatedAt;not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deletedAt;index" json:"-"`
}

func (User) TableName() string {
	return "users"
}
````

## File: services/api-service/internal/repository/conversation_repository.go
````go
package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"api-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MemberInput struct {
	UserID uuid.UUID
	Role   domain.MemberRole
}

type ConversationRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Conversation, error)
	FindDirectConversation(ctx context.Context, userA, userB uuid.UUID) (*domain.Conversation, error)
	FindUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error)
	Create(ctx context.Context, convo *domain.Conversation, members []MemberInput) (*domain.Conversation, error)
	AddMember(ctx context.Context, conversationID, userID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error)
	RemoveMember(ctx context.Context, conversationID, userID uuid.UUID) (bool, error)
	FindMember(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationMember, error)
	UpsertReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error)
	GetReadReceipt(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationReadReceipt, error)
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Conversation, error) {
	var convo domain.Conversation
	err := r.db.WithContext(ctx).
		Preload("Members.User").
		Where("id = ?", id).
		First(&convo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find conversation by id: %w", err)
	}
	return &convo, nil
}

func (r *conversationRepository) FindDirectConversation(ctx context.Context, userA, userB uuid.UUID) (*domain.Conversation, error) {
	var convoID uuid.UUID
	// Tìm conversation ID loại 'direct' mà cả userA và userB đều là members
	query := `
		SELECT c.id FROM conversations c
		JOIN conversation_members m1 ON m1."conversationId" = c.id AND m1."userId" = ?
		JOIN conversation_members m2 ON m2."conversationId" = c.id AND m2."userId" = ?
		WHERE c.type = 'direct'
		LIMIT 1
	`
	err := r.db.WithContext(ctx).Raw(query, userA, userB).Scan(&convoID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to check direct conversation: %w", err)
	}
	if convoID == uuid.Nil {
		return nil, nil
	}
	return r.FindByID(ctx, convoID)
}

func (r *conversationRepository) FindUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error) {
	var convos []domain.Conversation
	subQuery := r.db.Model(&domain.ConversationMember{}).
		Select("\"conversationId\"").
		Where("\"userId\" = ?", userID)

	err := r.db.WithContext(ctx).
		Preload("Members.User").
		Where("id IN (?)", subQuery).
		Order("\"updatedAt\" DESC").
		Find(&convos).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user conversations: %w", err)
	}
	return convos, nil
}

func (r *conversationRepository) Create(ctx context.Context, convo *domain.Conversation, members []MemberInput) (*domain.Conversation, error) {
	if convo.ID == uuid.Nil {
		convo.ID = uuid.New()
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(convo).Error; err != nil {
			return err
		}

		for _, m := range members {
			memEntity := domain.ConversationMember{
				ID:             uuid.New(),
				ConversationID: convo.ID,
				UserID:         m.UserID,
				Role:           m.Role,
			}
			if err := tx.Create(&memEntity).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create conversation with members: %w", err)
	}

	return r.FindByID(ctx, convo.ID)
}

func (r *conversationRepository) AddMember(ctx context.Context, conversationID, userID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error) {
	var member domain.ConversationMember
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		First(&member).Error
	if err == nil {
		return &member, nil
	}

	newMember := domain.ConversationMember{
		ID:             uuid.New(),
		ConversationID: conversationID,
		UserID:         userID,
		Role:           role,
	}
	if err := r.db.WithContext(ctx).Create(&newMember).Error; err != nil {
		return nil, fmt.Errorf("failed to add member: %w", err)
	}

	_ = r.db.WithContext(ctx).Preload("User").Where("id = ?", newMember.ID).First(&newMember)
	return &newMember, nil
}

func (r *conversationRepository) RemoveMember(ctx context.Context, conversationID, userID uuid.UUID) (bool, error) {
	res := r.db.WithContext(ctx).
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		Delete(&domain.ConversationMember{})
	return res.RowsAffected > 0, res.Error
}

func (r *conversationRepository) FindMember(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationMember, error) {
	var member domain.ConversationMember
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		First(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find conversation member: %w", err)
	}
	return &member, nil
}

func (r *conversationRepository) UpsertReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error) {
	receipt := domain.ConversationReadReceipt{
		ConversationID:  conversationID,
		UserID:          userID,
		LastMessageSeen: lastMessageSeen,
		LastReadAt:      time.Now().UTC(),
	}

	err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "conversationId"}, {Name: "userId"}},
		DoUpdates: clause.AssignmentColumns([]string{"lastMessageSeen", "lastReadAt"}),
	}).Create(&receipt).Error

	if err != nil {
		return nil, fmt.Errorf("failed to upsert read receipt: %w", err)
	}
	return &receipt, nil
}

func (r *conversationRepository) GetReadReceipt(ctx context.Context, conversationID, userID uuid.UUID) (*domain.ConversationReadReceipt, error) {
	var receipt domain.ConversationReadReceipt
	err := r.db.WithContext(ctx).
		Where("\"conversationId\" = ? AND \"userId\" = ?", conversationID, userID).
		First(&receipt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get read receipt: %w", err)
	}
	return &receipt, nil
}
````

## File: services/api-service/internal/repository/friend_repository.go
````go
package repository

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FriendRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Friend, error)
	FindByUsers(ctx context.Context, userA, userB uuid.UUID) (*domain.Friend, error)
	FindPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	FindFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	Create(ctx context.Context, friend *domain.Friend) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status domain.FriendStatus) (*domain.Friend, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
}

type friendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return &friendRepository{db: db}
}

func (r *friendRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Friend, error) {
	var friend domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Preload("Addressee").
		Where("id = ?", id).
		First(&friend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find friend by id: %w", err)
	}
	return &friend, nil
}

func (r *friendRepository) FindByUsers(ctx context.Context, userA, userB uuid.UUID) (*domain.Friend, error) {
	var friend domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Preload("Addressee").
		Where("((\"requesterId\" = ? AND \"addresseeId\" = ?) OR (\"requesterId\" = ? AND \"addresseeId\" = ?))", userA, userB, userB, userA).
		First(&friend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find friend relationship: %w", err)
	}
	return &friend, nil
}

func (r *friendRepository) FindPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	var list []domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Where("\"addresseeId\" = ? AND status = ?", userID, domain.FriendStatusPending).
		Order("\"createdAt\" DESC").
		Find(&list).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find pending friend requests: %w", err)
	}
	return list, nil
}

func (r *friendRepository) FindFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	var list []domain.Friend
	err := r.db.WithContext(ctx).
		Preload("Requester").
		Preload("Addressee").
		Where("(\"requesterId\" = ? OR \"addresseeId\" = ?) AND status = ?", userID, userID, domain.FriendStatusAccepted).
		Order("\"updatedAt\" DESC").
		Find(&list).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find friends list: %w", err)
	}
	return list, nil
}

func (r *friendRepository) Create(ctx context.Context, friend *domain.Friend) error {
	if friend.ID == uuid.Nil {
		friend.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(friend).Error
}

func (r *friendRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status domain.FriendStatus) (*domain.Friend, error) {
	if err := r.db.WithContext(ctx).Model(&domain.Friend{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return nil, fmt.Errorf("failed to update friend status: %w", err)
	}
	return r.FindByID(ctx, id)
}

func (r *friendRepository) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Friend{})
	return res.RowsAffected > 0, res.Error
}
````

## File: services/api-service/internal/repository/user_repository.go
````go
package repository

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Search(ctx context.Context, query string) ([]domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}
	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Search(ctx context.Context, query string) ([]domain.User, error) {
	var users []domain.User
	q := "%" + query + "%"
	err := r.db.WithContext(ctx).
		Where("username ILIKE ? OR name ILIKE ? OR phone ILIKE ?", q, q, q).
		Limit(50).
		Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}
	return users, nil
}
````

## File: services/api-service/internal/usecase/auth_usecase.go
````go
package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"api-service/internal/config"
	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists   = errors.New("Username already exists")
	ErrInvalidCredentials  = errors.New("Invalid credentials")
	ErrUnauthorized        = errors.New("Unauthorized")
)

type AuthTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthClaims struct {
	UserID   string `json:"sub"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthUsecase interface {
	Register(ctx context.Context, username, password, name string, phone, address *string) (*domain.User, error)
	Login(ctx context.Context, username, password string) (*AuthTokens, error)
	ValidateToken(tokenStr string) (*AuthClaims, error)
}

type authUsecase struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func NewAuthUsecase(cfg *config.Config, userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u *authUsecase) Register(ctx context.Context, username, password, name string, phone, address *string) (*domain.User, error) {
	existing, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &domain.User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: string(hash),
		Name:         name,
		Phone:        phone,
		Address:      address,
	}

	if err := u.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *authUsecase) Login(ctx context.Context, username, password string) (*AuthTokens, error) {
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return u.generateTokens(user.ID, user.Username)
}

func (u *authUsecase) generateTokens(userID uuid.UUID, username string) (*AuthTokens, error) {
	now := time.Now()
	accessExpiry := now.Add(u.cfg.JWTExpiresIn)
	refreshExpiry := now.Add(7 * 24 * time.Hour)

	accessClaims := AuthClaims{
		UserID:   userID.String(),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshClaims := AuthClaims{
		UserID:   userID.String(),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &AuthTokens{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, nil
}

func (u *authUsecase) ValidateToken(tokenStr string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(u.cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, ErrUnauthorized
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok || !token.Valid {
		return nil, ErrUnauthorized
	}

	return claims, nil
}
````

## File: services/api-service/internal/usecase/conversation_usecase.go
````go
package usecase

import (
	"context"
	"errors"

	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrConversationNotFound = errors.New("Cuộc hội thoại không tồn tại hoặc bạn không phải thành viên")
	ErrNotMember            = errors.New("Bạn không phải thành viên của cuộc hội thoại này")
	ErrCannotChatSelf       = errors.New("Không thể tạo cuộc hội thoại với chính mình")
	ErrNotAdmin             = errors.New("Chỉ quản trị viên mới có quyền thao tác")
)

type ConversationUsecase interface {
	GetUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error)
	GetOrCreateDirectConversation(ctx context.Context, userID, partnerID uuid.UUID) (*domain.Conversation, error)
	CreateGroupConversation(ctx context.Context, creatorID uuid.UUID, name string, iconURL *string, memberIDs []uuid.UUID) (*domain.Conversation, error)
	GetConversationByID(ctx context.Context, conversationID, userID uuid.UUID) (*domain.Conversation, error)
	AddMember(ctx context.Context, conversationID, requesterID, newMemberID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error)
	RemoveMember(ctx context.Context, conversationID, requesterID, targetUserID uuid.UUID) (bool, error)
	UpdateReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error)
}

type conversationUsecase struct {
	convoRepo repository.ConversationRepository
	userRepo  repository.UserRepository
}

func NewConversationUsecase(convoRepo repository.ConversationRepository, userRepo repository.UserRepository) ConversationUsecase {
	return &conversationUsecase{
		convoRepo: convoRepo,
		userRepo:  userRepo,
	}
}

func (u *conversationUsecase) GetUserConversations(ctx context.Context, userID uuid.UUID) ([]domain.Conversation, error) {
	return u.convoRepo.FindUserConversations(ctx, userID)
}

func (u *conversationUsecase) GetOrCreateDirectConversation(ctx context.Context, userID, partnerID uuid.UUID) (*domain.Conversation, error) {
	if userID == partnerID {
		return nil, ErrCannotChatSelf
	}

	partner, err := u.userRepo.FindByID(ctx, partnerID)
	if err != nil {
		return nil, err
	}
	if partner == nil {
		return nil, ErrTargetNotFound
	}

	existing, err := u.convoRepo.FindDirectConversation(ctx, userID, partnerID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}

	newConvo := &domain.Conversation{
		ID:        uuid.New(),
		Type:      domain.ConversationTypeDirect,
		CreatedBy: &userID,
	}

	members := []repository.MemberInput{
		{UserID: userID, Role: domain.MemberRoleMember},
		{UserID: partnerID, Role: domain.MemberRoleMember},
	}

	return u.convoRepo.Create(ctx, newConvo, members)
}

func (u *conversationUsecase) CreateGroupConversation(ctx context.Context, creatorID uuid.UUID, name string, iconURL *string, memberIDs []uuid.UUID) (*domain.Conversation, error) {
	if name == "" {
		return nil, errors.New("Tên nhóm chat không được để trống")
	}

	newConvo := &domain.Conversation{
		ID:        uuid.New(),
		Name:      &name,
		Type:      domain.ConversationTypeGroup,
		IconURL:   iconURL,
		CreatedBy: &creatorID,
	}

	memberMap := make(map[uuid.UUID]bool)
	memberMap[creatorID] = true

	members := []repository.MemberInput{
		{UserID: creatorID, Role: domain.MemberRoleAdmin},
	}

	for _, mid := range memberIDs {
		if !memberMap[mid] {
			memberMap[mid] = true
			members = append(members, repository.MemberInput{
				UserID: mid,
				Role:   domain.MemberRoleMember,
			})
		}
	}

	return u.convoRepo.Create(ctx, newConvo, members)
}

func (u *conversationUsecase) GetConversationByID(ctx context.Context, conversationID, userID uuid.UUID) (*domain.Conversation, error) {
	member, err := u.convoRepo.FindMember(ctx, conversationID, userID)
	if err != nil {
		return nil, err
	}
	if member == nil {
		return nil, ErrConversationNotFound
	}

	return u.convoRepo.FindByID(ctx, conversationID)
}

func (u *conversationUsecase) AddMember(ctx context.Context, conversationID, requesterID, newMemberID uuid.UUID, role domain.MemberRole) (*domain.ConversationMember, error) {
	convo, err := u.convoRepo.FindByID(ctx, conversationID)
	if err != nil {
		return nil, err
	}
	if convo == nil {
		return nil, ErrConversationNotFound
	}

	reqMember, err := u.convoRepo.FindMember(ctx, conversationID, requesterID)
	if err != nil {
		return nil, err
	}
	if reqMember == nil {
		return nil, ErrNotMember
	}

	if convo.Type == domain.ConversationTypeDirect {
		return nil, errors.New("Không thể thêm thành viên vào cuộc hội thoại 1-1")
	}

	if role == "" {
		role = domain.MemberRoleMember
	}

	return u.convoRepo.AddMember(ctx, conversationID, newMemberID, role)
}

func (u *conversationUsecase) RemoveMember(ctx context.Context, conversationID, requesterID, targetUserID uuid.UUID) (bool, error) {
	convo, err := u.convoRepo.FindByID(ctx, conversationID)
	if err != nil {
		return false, err
	}
	if convo == nil {
		return false, ErrConversationNotFound
	}

	reqMember, err := u.convoRepo.FindMember(ctx, conversationID, requesterID)
	if err != nil {
		return false, err
	}
	if reqMember == nil {
		return false, ErrNotMember
	}

	// Người dùng tự rời nhóm
	if requesterID == targetUserID {
		return u.convoRepo.RemoveMember(ctx, conversationID, targetUserID)
	}

	// Quản trị viên xóa thành viên khác
	if reqMember.Role != domain.MemberRoleAdmin {
		return false, ErrNotAdmin
	}

	return u.convoRepo.RemoveMember(ctx, conversationID, targetUserID)
}

func (u *conversationUsecase) UpdateReadReceipt(ctx context.Context, conversationID, userID uuid.UUID, lastMessageSeen string) (*domain.ConversationReadReceipt, error) {
	member, err := u.convoRepo.FindMember(ctx, conversationID, userID)
	if err != nil {
		return nil, err
	}
	if member == nil {
		return nil, ErrNotMember
	}

	return u.convoRepo.UpsertReadReceipt(ctx, conversationID, userID, lastMessageSeen)
}
````

## File: services/api-service/internal/usecase/friend_usecase.go
````go
package usecase

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrCannotSelfFriend   = errors.New("Không thể kết bạn với chính mình")
	ErrTargetNotFound     = errors.New("Người dùng không tồn tại")
	ErrRequestAlreadySent = errors.New("Lời mời kết bạn đã tồn tại hoặc đã là bạn bè")
	ErrRelationshipNotFound = errors.New("Không tìm thấy mối quan hệ hoặc lời mời")
	ErrNotAddressee       = errors.New("Bạn không có quyền thao tác trên lời mời này")
)

type FriendUsecase interface {
	GetFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	GetPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error)
	SendFriendRequest(ctx context.Context, requesterID, addresseeID uuid.UUID) (*domain.Friend, error)
	AcceptFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error)
	DeclineFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error)
	BlockUser(ctx context.Context, userID, targetUserID uuid.UUID) (*domain.Friend, error)
	UnfriendOrCancel(ctx context.Context, userID, relationshipID uuid.UUID) (bool, error)
}

type friendUsecase struct {
	friendRepo repository.FriendRepository
	userRepo   repository.UserRepository
}

func NewFriendUsecase(friendRepo repository.FriendRepository, userRepo repository.UserRepository) FriendUsecase {
	return &friendUsecase{
		friendRepo: friendRepo,
		userRepo:   userRepo,
	}
}

func (u *friendUsecase) GetFriendsList(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	return u.friendRepo.FindFriendsList(ctx, userID)
}

func (u *friendUsecase) GetPendingRequests(ctx context.Context, userID uuid.UUID) ([]domain.Friend, error) {
	return u.friendRepo.FindPendingRequests(ctx, userID)
}

func (u *friendUsecase) SendFriendRequest(ctx context.Context, requesterID, addresseeID uuid.UUID) (*domain.Friend, error) {
	if requesterID == addresseeID {
		return nil, ErrCannotSelfFriend
	}

	targetUser, err := u.userRepo.FindByID(ctx, addresseeID)
	if err != nil {
		return nil, err
	}
	if targetUser == nil {
		return nil, ErrTargetNotFound
	}

	existing, err := u.friendRepo.FindByUsers(ctx, requesterID, addresseeID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		if existing.Status == domain.FriendStatusPending || existing.Status == domain.FriendStatusAccepted {
			return nil, ErrRequestAlreadySent
		}
		if existing.Status == domain.FriendStatusBlocked {
			return nil, errors.New("Không thể gửi lời mời do người dùng đã bị chặn")
		}
		// Nếu trước đó declined, cho phép gửi lại bằng cách cập nhật sang pending
		return u.friendRepo.UpdateStatus(ctx, existing.ID, domain.FriendStatusPending)
	}

	newFriend := &domain.Friend{
		ID:          uuid.New(),
		RequesterID: requesterID,
		AddresseeID: addresseeID,
		Status:      domain.FriendStatusPending,
	}

	if err := u.friendRepo.Create(ctx, newFriend); err != nil {
		return nil, fmt.Errorf("failed to send friend request: %w", err)
	}

	return u.friendRepo.FindByID(ctx, newFriend.ID)
}

func (u *friendUsecase) AcceptFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error) {
	friend, err := u.friendRepo.FindByID(ctx, requestID)
	if err != nil {
		return nil, err
	}
	if friend == nil {
		return nil, ErrRelationshipNotFound
	}
	if friend.AddresseeID != userID {
		return nil, ErrNotAddressee
	}
	if friend.Status != domain.FriendStatusPending {
		return nil, errors.New("Lời mời không ở trạng thái chờ duyệt")
	}

	return u.friendRepo.UpdateStatus(ctx, requestID, domain.FriendStatusAccepted)
}

func (u *friendUsecase) DeclineFriendRequest(ctx context.Context, userID, requestID uuid.UUID) (*domain.Friend, error) {
	friend, err := u.friendRepo.FindByID(ctx, requestID)
	if err != nil {
		return nil, err
	}
	if friend == nil {
		return nil, ErrRelationshipNotFound
	}
	if friend.AddresseeID != userID {
		return nil, ErrNotAddressee
	}

	return u.friendRepo.UpdateStatus(ctx, requestID, domain.FriendStatusDeclined)
}

func (u *friendUsecase) BlockUser(ctx context.Context, userID, targetUserID uuid.UUID) (*domain.Friend, error) {
	if userID == targetUserID {
		return nil, errors.New("Không thể chặn chính mình")
	}

	existing, err := u.friendRepo.FindByUsers(ctx, userID, targetUserID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return u.friendRepo.UpdateStatus(ctx, existing.ID, domain.FriendStatusBlocked)
	}

	newBlock := &domain.Friend{
		ID:          uuid.New(),
		RequesterID: userID,
		AddresseeID: targetUserID,
		Status:      domain.FriendStatusBlocked,
	}
	if err := u.friendRepo.Create(ctx, newBlock); err != nil {
		return nil, err
	}

	return u.friendRepo.FindByID(ctx, newBlock.ID)
}

func (u *friendUsecase) UnfriendOrCancel(ctx context.Context, userID, relationshipID uuid.UUID) (bool, error) {
	friend, err := u.friendRepo.FindByID(ctx, relationshipID)
	if err != nil {
		return false, err
	}
	if friend == nil {
		return false, ErrRelationshipNotFound
	}
	if friend.RequesterID != userID && friend.AddresseeID != userID {
		return false, errors.New("Bạn không thuộc mối quan hệ này")
	}

	return u.friendRepo.Delete(ctx, relationshipID)
}
````

## File: services/api-service/internal/usecase/user_usecase.go
````go
package usecase

import (
	"context"
	"errors"

	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

type UserUsecase interface {
	GetProfile(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, name, phone, address *string) (*domain.User, error)
	SearchUsers(ctx context.Context, query string) ([]domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, id uuid.UUID, name, phone, address *string) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	if name != nil && *name != "" {
		user.Name = *name
	}
	if phone != nil {
		user.Phone = phone
	}
	if address != nil {
		user.Address = address
	}

	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) SearchUsers(ctx context.Context, query string) ([]domain.User, error) {
	if query == "" {
		return []domain.User{}, nil
	}
	return u.userRepo.Search(ctx, query)
}

func (u *userUsecase) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
````

## File: services/api-service/README.md
````markdown
<p align="center">
  <a href="http://nestjs.com/" target="blank"><img src="https://nestjs.com/img/logo-small.svg" width="120" alt="Nest Logo" /></a>
</p>

[circleci-image]: https://img.shields.io/circleci/build/github/nestjs/nest/master?token=abc123def456
[circleci-url]: https://circleci.com/gh/nestjs/nest

  <p align="center">A progressive <a href="http://nodejs.org" target="_blank">Node.js</a> framework for building efficient and scalable server-side applications.</p>
    <p align="center">
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/v/@nestjs/core.svg" alt="NPM Version" /></a>
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/l/@nestjs/core.svg" alt="Package License" /></a>
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/dm/@nestjs/common.svg" alt="NPM Downloads" /></a>
<a href="https://circleci.com/gh/nestjs/nest" target="_blank"><img src="https://img.shields.io/circleci/build/github/nestjs/nest/master" alt="CircleCI" /></a>
<a href="https://discord.gg/G7Qnnhy" target="_blank"><img src="https://img.shields.io/badge/discord-online-brightgreen.svg" alt="Discord"/></a>
<a href="https://opencollective.com/nest#backer" target="_blank"><img src="https://opencollective.com/nest/backers/badge.svg" alt="Backers on Open Collective" /></a>
<a href="https://opencollective.com/nest#sponsor" target="_blank"><img src="https://opencollective.com/nest/sponsors/badge.svg" alt="Sponsors on Open Collective" /></a>
  <a href="https://paypal.me/kamilmysliwiec" target="_blank"><img src="https://img.shields.io/badge/Donate-PayPal-ff3f59.svg" alt="Donate us"/></a>
    <a href="https://opencollective.com/nest#sponsor"  target="_blank"><img src="https://img.shields.io/badge/Support%20us-Open%20Collective-41B883.svg" alt="Support us"></a>
  <a href="https://twitter.com/nestframework" target="_blank"><img src="https://img.shields.io/twitter/follow/nestframework.svg?style=social&label=Follow" alt="Follow us on Twitter"></a>
</p>
  <!--[![Backers on Open Collective](https://opencollective.com/nest/backers/badge.svg)](https://opencollective.com/nest#backer)
  [![Sponsors on Open Collective](https://opencollective.com/nest/sponsors/badge.svg)](https://opencollective.com/nest#sponsor)-->

## Description

[Nest](https://github.com/nestjs/nest) framework TypeScript starter repository.

## Project setup

```bash
$ npm install
```

## Compile and run the project

```bash
# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

## Run tests

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```

## Deployment

When you're ready to deploy your NestJS application to production, there are some key steps you can take to ensure it runs as efficiently as possible. Check out the [deployment documentation](https://docs.nestjs.com/deployment) for more information.

If you are looking for a cloud-based platform to deploy your NestJS application, check out [Mau](https://mau.nestjs.com), our official platform for deploying NestJS applications on AWS. Mau makes deployment straightforward and fast, requiring just a few simple steps:

```bash
$ npm install -g @nestjs/mau
$ mau deploy
```

With Mau, you can deploy your application in just a few clicks, allowing you to focus on building features rather than managing infrastructure.

## Resources

Check out a few resources that may come in handy when working with NestJS:

- Visit the [NestJS Documentation](https://docs.nestjs.com) to learn more about the framework.
- For questions and support, please visit our [Discord channel](https://discord.gg/G7Qnnhy).
- To dive deeper and get more hands-on experience, check out our official video [courses](https://courses.nestjs.com/).
- Deploy your application to AWS with the help of [NestJS Mau](https://mau.nestjs.com) in just a few clicks.
- Visualize your application graph and interact with the NestJS application in real-time using [NestJS Devtools](https://devtools.nestjs.com).
- Need help with your project (part-time to full-time)? Check out our official [enterprise support](https://enterprise.nestjs.com).
- To stay in the loop and get updates, follow us on [X](https://x.com/nestframework) and [LinkedIn](https://linkedin.com/company/nestjs).
- Looking for a job, or have a job to offer? Check out our official [Jobs board](https://jobs.nestjs.com).

## Support

Nest is an MIT-licensed open source project. It can grow thanks to the sponsors and support by the amazing backers. If you'd like to join them, please [read more here](https://docs.nestjs.com/support).

## Stay in touch

- Author - [Kamil Myśliwiec](https://twitter.com/kammysliwiec)
- Website - [https://nestjs.com](https://nestjs.com/)
- Twitter - [@nestframework](https://twitter.com/nestframework)

## License

Nest is [MIT licensed](https://github.com/nestjs/nest/blob/master/LICENSE).
````

## File: services/chat-engine/internal/dispatcher/dispatcher_grpc.go
````go
package dispatcher

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	grpcclient "chat-system/pkg/grpcclient"
	natsclient "chat-system/pkg/nats"
	pb "chat-system/pkg/proto"
	"chat-worker/internal/config"

	"github.com/nats-io/nats.go"
)

type grpcEventDispatcher struct {
	ackPublisher  *natsclient.InstrumentedPublisher[contracts.OutboundBrokerEvent]
	clientManager *grpcclient.GatewayClientManager
}

// NewGRPCEventDispatcher khởi tạo Dispatcher sử dụng GatewayClientManager từ pkg/grpcclient
func NewGRPCEventDispatcher(cfg *config.DeliveryConfig, nc *nats.Conn) EventDispatcher {
	rawPublisher := natsclient.NewPublisher[contracts.OutboundBrokerEvent](nc, "")
	return &grpcEventDispatcher{
		ackPublisher: natsclient.NewInstrumentedPublisher(rawPublisher, natsclient.PublisherConfig{
			ServiceName: "chat-engine",
			Stage:       "sender_ack",
			EventType:   "ack",
		}),
		clientManager: grpcclient.NewGatewayClientManager(cfg.GRPCPort, cfg.GRPCServiceSuffix),
	}
}

// SendAckToSender gửi sự kiện ACK về Node Gateway của người gửi qua Broker (NATS Topic)
func (d *grpcEventDispatcher) SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route sender ack")
	}

	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.ackPublisher.PublishToSubject(ctx, subject, event); err != nil {
		return fmt.Errorf("failed to publish sender ack to subject '%s': %w", subject, err)
	}

	log.Printf("[Dispatcher:gRPC] Sender ACK published via Broker: msg_id=%s, client_msg_id=%s -> subject=%s",
		event.MessageID, event.ClientMsgID, subject)
	return nil
}

// DispatchToGateway gửi tin nhắn trực tiếp qua gRPC tới WS Gateway Pod mà người nhận kết nối
// Telemetry (Tracing W3C, Duration, Metrics) được tự động xử lý bởi gRPC Client Interceptor
func (d *grpcEventDispatcher) DispatchToGateway(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route message to gateway via gRPC")
	}

	client, err := d.clientManager.GetClient(gatewayNode)
	if err != nil {
		return fmt.Errorf("failed to get gRPC client for node %s: %w", gatewayNode, err)
	}

	req := &pb.PushMessageRequest{
		MessageId:      event.MessageID,
		ClientMsgId:    event.ClientMsgID,
		ConversationId: event.ConversationID,
		SenderId:       event.SenderID,
		ReceiverId:     event.ReceiverID,
		Content:        event.Content,
		Type:           string(event.Type),
		Timestamp:      event.Timestamp,
		Payload:        event.Payload,
	}

	resp, err := client.PushMessageToUser(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC PushMessageToUser failed for receiver %s: %w", event.ReceiverID, err)
	}
	if !resp.Success {
		return fmt.Errorf("gRPC PushMessageToUser rejected: %s", resp.ErrorMessage)
	}

	log.Printf("[Dispatcher:gRPC] Message dispatched successfully: msg_id=%s, receiver=%s -> gateway=%s",
		event.MessageID, event.ReceiverID, gatewayNode)
	return nil
}

func (d *grpcEventDispatcher) Close() error {
	return d.clientManager.Close()
}
````

## File: services/chat-engine/internal/dispatcher/dispatcher_nats.go
````go
package dispatcher

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"

	"github.com/nats-io/nats.go"
)

type natsEventDispatcher struct {
	outboundPublisher *natsclient.InstrumentedPublisher[contracts.OutboundBrokerEvent]
}

// NewNATSEventDispatcher khởi tạo Dispatcher sử dụng NATS Pub/Sub với Telemetry Decorator
func NewNATSEventDispatcher(nc *nats.Conn) EventDispatcher {
	rawPublisher := natsclient.NewPublisher[contracts.OutboundBrokerEvent](nc, "")
	return &natsEventDispatcher{
		outboundPublisher: natsclient.NewInstrumentedPublisher(rawPublisher, natsclient.PublisherConfig{
			ServiceName: "chat-engine",
			Stage:       "outbound_dispatch",
			EventType:   "outbound",
		}),
	}
}

// SendAckToSender gửi sự kiện ACK về đúng Topic của Node Gateway mà người gửi kết nối: "chat.gateway.{node_id}"
func (d *natsEventDispatcher) SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route sender ack")
	}
	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.outboundPublisher.PublishToSubject(ctx, subject, event); err != nil {
		return fmt.Errorf("failed to publish sender ack to subject '%s': %w", subject, err)
	}
	log.Printf("[Dispatcher:NATS] Sender ACK published: msg_id=%s, client_msg_id=%s -> subject=%s",
		event.MessageID, event.ClientMsgID, subject)
	return nil
}

// DispatchToGateway chuyển tiếp tin nhắn tới đúng Topic của Node Gateway mà người nhận kết nối
func (d *natsEventDispatcher) DispatchToGateway(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error {
	if gatewayNode == "" {
		return fmt.Errorf("gatewayNode is empty, cannot route message to gateway")
	}

	subject := contracts.GatewayNodeSubject(gatewayNode)
	if err := d.outboundPublisher.PublishToNode(ctx, gatewayNode, subject, event); err != nil {
		return fmt.Errorf("failed to publish outbound message to subject '%s': %w", subject, err)
	}

	log.Printf("[Dispatcher:NATS] Message dispatched to receiver gateway: msg_id=%s, receiver=%s -> subject=%s",
		event.MessageID, event.ReceiverID, subject)
	return nil
}

func (d *natsEventDispatcher) Close() error {
	return nil
}
````

## File: services/chat-engine/internal/domain/message.go
````go
package domain

import "time"

// Message represents the core message entity persisted in Cassandra
type Message struct {
	ConversationID string    `json:"conversation_id"`
	ID             string    `json:"id"`
	SenderID       string    `json:"sender_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
````

## File: services/chat-engine/internal/migrator/migrator.go
````go
package migrator

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"chat-worker/internal/config"
	"chat-worker/migrations"

	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Run executes database migrations using embedded CQL files.
func Run(cfg *config.DatabaseConfig) error {
	if len(cfg.Hosts) == 0 {
		return fmt.Errorf("cassandra hosts list is empty")
	}

	hostList, port := parseHosts(cfg.Hosts)

	// 1. Ensure Keyspace exists before binding session to it
	if err := ensureKeyspace(hostList, port, cfg.Keyspace); err != nil {
		return fmt.Errorf("failed to ensure keyspace '%s': %w", cfg.Keyspace, err)
	}

	// 2. Open Cassandra session bound to keyspace
	cluster := gocql.NewCluster(hostList...)
	cluster.Port = port
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = 10 * time.Second
	cluster.ConnectTimeout = 15 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to create session for migration: %w", err)
	}
	defer session.Close()

	// 3. Initialize Cassandra migrate driver
	driver, err := cassandra.WithInstance(session, &cassandra.Config{
		KeyspaceName: cfg.Keyspace,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize cassandra migration driver: %w", err)
	}

	// 4. Initialize iofs source driver with embedded files
	sourceDriver, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("failed to initialize iofs source driver: %w", err)
	}

	// 5. Run migrations
	m, err := migrate.NewWithInstance("iofs", sourceDriver, "cassandra", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("cassandra migration failed: %w", err)
	}

	log.Printf("[Migrator] Cassandra schema migration completed successfully for keyspace '%s'", cfg.Keyspace)
	return nil
}

func ensureKeyspace(hosts []string, port int, keyspace string) error {
	cluster := gocql.NewCluster(hosts...)
	cluster.Port = port
	cluster.Timeout = 10 * time.Second
	cluster.ConnectTimeout = 15 * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	query := fmt.Sprintf(
		"CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};",
		keyspace,
	)

	return session.Query(query).Exec()
}

func parseHosts(rawHosts []string) ([]string, int) {
	var hostList []string
	port := 9042

	for _, h := range rawHosts {
		if strings.Contains(h, ":") {
			host, portStr, err := net.SplitHostPort(h)
			if err == nil {
				hostList = append(hostList, host)
				if parsedPort, err := strconv.Atoi(portStr); err == nil && parsedPort > 0 {
					port = parsedPort
				}
				continue
			}
		}
		hostList = append(hostList, h)
	}

	return hostList, port
}
````

## File: services/chat-engine/migrations/000001_create_messages_table.down.cql
````
DROP TABLE IF EXISTS messages;
````

## File: services/chat-engine/migrations/000001_create_messages_table.up.cql
````
CREATE TABLE IF NOT EXISTS messages (
    conversation_id text,
    id text,
    sender_id text,
    content text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY ((conversation_id), id)
) WITH CLUSTERING ORDER BY (id ASC);
````

## File: services/chat-engine/migrations/000002_create_read_receipts_table.down.cql
````
DROP TABLE IF EXISTS conversation_read_receipts;
````

## File: services/chat-engine/migrations/000002_create_read_receipts_table.up.cql
````
CREATE TABLE IF NOT EXISTS conversation_read_receipts (
    conversation_id text,
    user_id text,
    last_message_seen text,
    last_read_at timestamp,
    PRIMARY KEY ((conversation_id), user_id)
);
````

## File: services/chat-engine/migrations/fs.go
````go
package migrations

import "embed"

//go:embed *.cql
var FS embed.FS
````

## File: services/client-simulator/cmd/main.go
````go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"client-simulator/internal/auth"
	"client-simulator/internal/bot"
	"client-simulator/internal/config"
	"client-simulator/internal/conversation"
	"client-simulator/internal/metrics"
)

func main() {
	// 1. Quét cờ -config nếu có, mặc định tìm configs/config.yaml
	initialCfgPath := "configs/config.yaml"
	for i, arg := range os.Args {
		if arg == "-config" && i+1 < len(os.Args) {
			initialCfgPath = os.Args[i+1]
		}
	}

	// 2. Nạp cấu hình từ YAML và Biến môi trường (Environment Variables)
	cfg, err := config.LoadConfig(initialCfgPath)
	if err != nil {
		log.Printf("[Config] Notice: %v", err)
	}

	// 3. CLI Flags có thể ghi đè cấu hình khi chạy test thủ công
	_ = flag.String("config", initialCfgPath, "Đường dẫn file cấu hình YAML")
	numBots := flag.Int("bots", cfg.Bots, "Số lượng bot muốn tạo và chạy mô phỏng")
	convsPerBot := flag.Int("convs", cfg.ConvsPerBot, "Số lượng cuộc hội thoại 1-1 tối đa trên mỗi bot")
	sendInterval := flag.Duration("interval", cfg.Interval, "Tần suất gửi tin nhắn của mỗi bot (vd: 200ms, 500ms, 1s)")
	apiURL := flag.String("api", cfg.APIURL, "URL gốc của API Service (hoặc NGINX Ingress)")
	wsURL := flag.String("ws", cfg.WSURL, "URL endpoint WebSocket của WS Gateway (hoặc NGINX Ingress)")
	testDuration := flag.Duration("duration", cfg.Duration, "Thời gian chạy test (0 = chạy liên tục đến khi Ctrl+C)")
	botPassword := flag.String("password", cfg.Password, "Mật khẩu chung cho các bot tài khoản")
	flag.Parse()

	log.Println("=================================================================")
	log.Println("           🚀 KHỞI ĐỘNG CHAT CLIENT SIMULATOR                   ")
	log.Println("=================================================================")
	log.Printf("Cấu hình: %d bots | ~%d convs/bot | Interval: %s | Duration: %s", *numBots, *convsPerBot, *sendInterval, *testDuration)
	log.Printf("API Service: %s | WS Gateway: %s", *apiURL, *wsURL)

	rootCtx, rootCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer rootCancel()

	authClient := auth.NewAuthClient(*apiURL)
	convoManager := conversation.NewManager(*apiURL)
	tracker := metrics.NewTracker()

	// 1. GIAI ĐOẠN 1: Authenticate / Register Bots song song (Concurrency = 30)
	log.Printf("[Phase 1] Authenticating %d bots concurrently via API Service...", *numBots)
	botSessions := make([]*auth.BotSession, *numBots)
	var authWg sync.WaitGroup
	authSem := make(chan struct{}, 30)
	var authErrMu sync.Mutex
	var firstAuthErr error

	for i := 1; i <= *numBots; i++ {
		authWg.Add(1)
		go func(idx int) {
			defer authWg.Done()
			authSem <- struct{}{}
			defer func() { <-authSem }()

			username := fmt.Sprintf("sim_bot_%04d", idx)
			session, err := authClient.EnsureBotAuth(username, *botPassword)
			if err != nil {
				authErrMu.Lock()
				if firstAuthErr == nil {
					firstAuthErr = fmt.Errorf("failed to authenticate bot %s: %w", username, err)
				}
				authErrMu.Unlock()
				return
			}
			botSessions[idx-1] = session
		}(i)
	}
	authWg.Wait()

	if firstAuthErr != nil {
		log.Fatalf("Authentication failed: %v", firstAuthErr)
	}
	log.Printf("  - Successfully authenticated %d/%d bots in parallel!", len(botSessions), *numBots)

	// 2. GIAI ĐOẠN 2: Setup Conversation Network song song
	log.Printf("[Phase 2] Setting up direct conversations between bots...")
	botTargets := convoManager.EnsureDirectConversations(botSessions, *convsPerBot)

	// 3. GIAI ĐOẠN 3: Connect WebSocket Clients (Smooth Ramp-Up Concurrency = 20)
	log.Printf("[Phase 3] Establishing WebSocket connections for %d bots...", len(botSessions))
	bots := make([]*bot.Bot, len(botSessions))

	var wsWg sync.WaitGroup
	wsSem := make(chan struct{}, 20)
	var activeCount int64
	var countMu sync.Mutex

	for idx, session := range botSessions {
		if session == nil {
			continue
		}
		wsWg.Add(1)
		time.Sleep(5 * time.Millisecond) // Smooth TCP syn pacing for Windows
		go func(i int, s *auth.BotSession) {
			defer wsWg.Done()
			wsSem <- struct{}{}
			defer func() { <-wsSem }()

			targets := botTargets[s.UserID]
			b := bot.NewBot(s, *wsURL, targets, tracker)

			connectCtx, connectCancel := context.WithTimeout(rootCtx, 20*time.Second)
			defer connectCancel()

			if err := b.Connect(connectCtx); err != nil {
				log.Printf("WARN: Bot %s failed to connect WS: %v", s.Username, err)
				return
			}

			bots[i] = b

			countMu.Lock()
			activeCount++
			countMu.Unlock()
		}(idx, session)
	}
	wsWg.Wait()

	log.Printf("Successfully connected %d/%d bots to WebSocket Gateway!", activeCount, *numBots)
	log.Printf("🚀 Kích hoạt đồng loạt %d bots bắt đầu phát sinh tin nhắn...", activeCount)
	for _, b := range bots {
		if b != nil {
			b.Start(rootCtx, *sendInterval)
		}
	}
	time.Sleep(500 * time.Millisecond)

	// 4. GIAI ĐOẠN 4: Live Dashboard Display Loop & Test Duration
	testCtx := rootCtx
	if *testDuration > 0 {
		log.Printf("⏱️ Bắt đầu đo tải: Chạy trong %s rồi tự động dừng...", *testDuration)
		var cancelTimer context.CancelFunc
		testCtx, cancelTimer = context.WithTimeout(rootCtx, *testDuration)
		defer cancelTimer()
	} else {
		log.Println("♾️ Bắt đầu đo tải liên tục (nhấn Ctrl+C để dừng)...")
	}

	startTime := time.Now()
	dashboardTicker := time.NewTicker(1 * time.Second)
	defer dashboardTicker.Stop()

	for {
		select {
		case <-testCtx.Done():
			log.Println("\nStopping simulator and closing connections...")
			for _, b := range bots {
				if b != nil {
					b.Stop()
				}
			}

			// Print final summary
			tracker.PrintDashboard(0, *numBots, time.Since(startTime))
			log.Println("\n✅ Simulator finished cleanly.")
			return

		case <-dashboardTicker.C:
			currentActive := 0
			for _, b := range bots {
				if b.IsConnected() {
					currentActive++
				}
			}
			tracker.PrintDashboard(currentActive, *numBots, time.Since(startTime))
		}
	}
}
````

## File: services/client-simulator/configs/config.yaml
````yaml
# ==============================================================================
# Client Simulator Configuration
# ==============================================================================
bots: 50
convs_per_bot: 5
interval: "200ms"
duration: "0s"
api_url: "http://localhost"
ws_url: "ws://localhost/ws"
password: "Pass@123456"
````

## File: services/client-simulator/Dockerfile
````
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
COPY --from=builder /app/services/client-simulator/configs ./configs

CMD ["./client-simulator"]
````

## File: services/client-simulator/go.mod
````
module client-simulator

go 1.25.0

require (
	github.com/caarlos0/env/v11 v11.4.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
````

## File: services/client-simulator/internal/auth/client.go
````go
package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type BotSession struct {
	UserID      string
	Username    string
	AccessToken string
}

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type registerPayload struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type loginPayload struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type apiResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       json.RawMessage `json:"data"`
}

type loginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c *AuthClient) EnsureBotAuth(username, password string) (*BotSession, error) {
	// 1. Thử login trước
	session, err := c.login(username, password)
	if err == nil {
		return session, nil
	}

	// 2. Nếu login thất bại, thử register
	_ = c.register(username, password)

	// 3. Login lại sau khi register
	return c.login(username, password)
}

func (c *AuthClient) register(username, password string) error {
	body, _ := json.Marshal(registerPayload{
		UserName: username,
		Password: password,
		Name:     fmt.Sprintf("Bot %s", username),
	})

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/auth/register", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("register failed with status %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

func (c *AuthClient) login(username, password string) (*BotSession, error) {
	body, _ := json.Marshal(loginPayload{
		UserName: username,
		Password: password,
	})

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/auth/login", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("login failed status %d: %s", resp.StatusCode, string(respBytes))
	}

	var rawResp apiResponse
	var token string
	if err := json.Unmarshal(respBytes, &rawResp); err == nil && len(rawResp.Data) > 0 {
		var tokenData loginResponseData
		if err := json.Unmarshal(rawResp.Data, &tokenData); err == nil && tokenData.AccessToken != "" {
			token = tokenData.AccessToken
		}
	}

	if token == "" {
		// Fallback parse directly
		var tokenData loginResponseData
		if err := json.Unmarshal(respBytes, &tokenData); err == nil && tokenData.AccessToken != "" {
			token = tokenData.AccessToken
		}
	}

	if token == "" {
		return nil, fmt.Errorf("could not extract accessToken from response: %s", string(respBytes))
	}

	userID, err := extractUserIDFromJWT(token)
	if err != nil {
		return nil, fmt.Errorf("failed to extract sub/user_id from jwt: %w", err)
	}

	return &BotSession{
		UserID:      userID,
		Username:    username,
		AccessToken: token,
	}, nil
}

func extractUserIDFromJWT(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid jwt format")
	}

	payloadPart := parts[1]
	// Handle Base64 URL padding
	if l := len(payloadPart) % 4; l > 0 {
		payloadPart += strings.Repeat("=", 4-l)
	}

	payloadBytes, err := base64.URLEncoding.DecodeString(payloadPart)
	if err != nil {
		return "", err
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return "", err
	}

	if sub, ok := claims["sub"].(string); ok && sub != "" {
		return sub, nil
	}
	if uid, ok := claims["user_id"].(string); ok && uid != "" {
		return uid, nil
	}

	return "", fmt.Errorf("no sub or user_id found in jwt claims")
}
````

## File: services/client-simulator/internal/bot/bot.go
````go
package bot

import (
	"client-simulator/internal/auth"
	"client-simulator/internal/conversation"
	"client-simulator/internal/metrics"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Bot struct {
	Session       *auth.BotSession
	DeviceID      string
	WSURL         string
	Conversations []conversation.Target
	Tracker       *metrics.Tracker

	connMu      sync.RWMutex
	conn        *websocket.Conn
	isConnected atomic.Bool

	inFlight  sync.Map // clientMsgID -> time.Time
	sendChan  chan []byte
	stopChan  chan struct{}
	closeOnce sync.Once
}

type wsEnvelope struct {
	Type        string          `json:"type"`
	ClientMsgID string          `json:"client_msg_id,omitempty"`
	Timestamp   int64           `json:"timestamp,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}

type sendMessagePayload struct {
	ConversationID string `json:"conversation_id"`
	ReceiverID     string `json:"receiver_id"`
	Content        string `json:"content"`
}

func NewBot(session *auth.BotSession, wsURL string, targets []conversation.Target, tracker *metrics.Tracker) *Bot {
	return &Bot{
		Session:       session,
		DeviceID:      fmt.Sprintf("dev_%s", session.Username),
		WSURL:         wsURL,
		Conversations: targets,
		Tracker:       tracker,
		sendChan:      make(chan []byte, 1024),
		stopChan:      make(chan struct{}),
	}
}

// IsConnected kiểm tra bot có đang giữ kết nối WebSocket sống hay không
func (b *Bot) IsConnected() bool {
	return b.isConnected.Load()
}

// Connect thiết lập kết nối WebSocket tới WS Gateway
func (b *Bot) Connect(ctx context.Context) error {
	u, err := url.Parse(b.WSURL)
	if err != nil {
		return fmt.Errorf("invalid ws url: %w", err)
	}

	q := u.Query()
	q.Set("token", b.Session.AccessToken)
	q.Set("device_id", b.DeviceID)
	u.RawQuery = q.Encode()

	dialer := websocket.Dialer{
		HandshakeTimeout: 15 * time.Second,
	}

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("bot %s dial error: %w", b.Session.Username, err)
	}

	b.connMu.Lock()
	if b.conn != nil {
		_ = b.conn.Close()
	}
	b.conn = conn
	b.connMu.Unlock()

	b.isConnected.Store(true)
	return nil
}

// Start bắt đầu các pump: read, write, heartbeat và traffic generator
func (b *Bot) Start(ctx context.Context, sendInterval time.Duration) {
	go b.lifecyclePump(ctx)
	go b.writePump(ctx)
	go b.heartbeatPump(ctx)
	go b.trafficPump(ctx, sendInterval)
}

func (b *Bot) Stop() {
	b.closeOnce.Do(func() {
		b.isConnected.Store(false)
		close(b.stopChan)
		b.connMu.Lock()
		if b.conn != nil {
			_ = b.conn.Close()
		}
		b.connMu.Unlock()
	})
}

// lifecyclePump quản lý vòng đời kết nối: đọc tin nhắn và tự động Reconnect khi mất mạng
func (b *Bot) lifecyclePump(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		default:
		}

		if !b.IsConnected() {
			b.reconnect(ctx)
			continue
		}

		b.connMu.RLock()
		conn := b.conn
		b.connMu.RUnlock()

		if conn == nil {
			b.reconnect(ctx)
			continue
		}

		_, data, err := conn.ReadMessage()
		if err != nil {
			// Socket bị đứt (Node Gateway sập hoặc network partition)
			b.isConnected.Store(false)
			b.connMu.Lock()
			if b.conn != nil {
				_ = b.conn.Close()
				b.conn = nil
			}
			b.connMu.Unlock()
			continue
		}

		var env wsEnvelope
		if err := json.Unmarshal(data, &env); err != nil {
			continue
		}

		type innerPayload struct {
			MessageID   string `json:"message_id"`
			ClientMsgID string `json:"client_msg_id"`
			SenderID    string `json:"sender_id"`
			ReceiverID  string `json:"receiver_id"`
			Type        string `json:"type"`
		}

		var inner innerPayload
		if len(env.Payload) > 0 {
			_ = json.Unmarshal(env.Payload, &inner)
		}

		msgID := env.ClientMsgID
		if msgID == "" {
			msgID = inner.ClientMsgID
		}

		// 0. Kiểm tra HEARTBEAT_ACK để đo Network Round-Trip Time (Network Delay)
		if env.Type == "HEARTBEAT_ACK" {
			if msgID != "" {
				if val, ok := b.inFlight.LoadAndDelete(msgID); ok {
					sentAt := val.(time.Time)
					b.Tracker.RecordNetworkDelay(time.Since(sentAt))
					continue
				}
			}
			if env.Timestamp > 0 {
				diff := time.Now().UnixMilli() - env.Timestamp
				if diff >= 0 {
					b.Tracker.RecordNetworkDelay(time.Duration(diff) * time.Millisecond)
				}
			}
			continue
		}

		// 1. Kiểm tra Sender ACK
		if msgID != "" {
			if val, ok := b.inFlight.LoadAndDelete(msgID); ok {
				sentAt := val.(time.Time)
				latency := time.Since(sentAt)
				b.Tracker.RecordAck(latency)
				continue
			}
		}

		// 2. Tin nhắn nhận được từ bạn chat
		if inner.SenderID != "" && inner.SenderID != b.Session.UserID {
			b.Tracker.RecordDelivered()
		} else if env.Type == "SEND_MESSAGE" && msgID == "" {
			b.Tracker.RecordDelivered()
		}
	}
}

// reconnect thực hiện kết nối lại với Exponential Backoff
func (b *Bot) reconnect(ctx context.Context) {
	backoff := 1 * time.Second
	maxBackoff := 8 * time.Second

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-time.After(backoff):
		}

		if err := b.Connect(ctx); err == nil {
			// Reconnect thành công!
			b.Tracker.RecordReconnect()
			return
		}

		// Tăng thời gian chờ thử lại (Backoff x 2)
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

func (b *Bot) writePump(ctx context.Context) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case msg, ok := <-b.sendChan:
			if !ok {
				return
			}
			if !b.IsConnected() {
				b.Tracker.RecordError()
				continue
			}

			b.connMu.RLock()
			conn := b.conn
			b.connMu.RUnlock()

			if conn != nil {
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					b.isConnected.Store(false)
					b.Tracker.RecordError()
				}
			}
		case <-ticker.C:
			if b.IsConnected() {
				b.connMu.RLock()
				conn := b.conn
				b.connMu.RUnlock()
				if conn != nil {
					_ = conn.WriteMessage(websocket.PingMessage, nil)
				}
			}
		}
	}
}

func (b *Bot) heartbeatPump(ctx context.Context) {
	// Khởi đầu với Jitter ngẫu nhiên tránh thundering herd
	initialJitter := time.Duration(rand.Intn(2000)) * time.Millisecond
	select {
	case <-ctx.Done():
		return
	case <-b.stopChan:
		return
	case <-time.After(initialJitter):
	}

	for {
		// Nhịp heartbeat định kỳ 4s + ngẫu nhiên 0-2s để liên tục đo Network RTT
		interval := 4*time.Second + time.Duration(rand.Intn(2000))*time.Millisecond
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-time.After(interval):
		}

		if !b.IsConnected() {
			continue
		}

		now := time.Now()
		hbID := fmt.Sprintf("hb_%s_%d", b.Session.Username, now.UnixNano())
		b.inFlight.Store(hbID, now)

		hbEnv := wsEnvelope{
			Type:        "HEARTBEAT",
			ClientMsgID: hbID,
			Timestamp:   now.UnixMilli(),
		}
		bytes, _ := json.Marshal(hbEnv)
		select {
		case b.sendChan <- bytes:
		default:
			b.inFlight.Delete(hbID)
		}
	}
}

func (b *Bot) trafficPump(ctx context.Context, interval time.Duration) {
	if len(b.Conversations) == 0 || interval <= 0 {
		return
	}

	jitter := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(jitter)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-ticker.C:
			if !b.IsConnected() {
				b.Tracker.RecordError()
				continue
			}

			target := b.Conversations[rand.Intn(len(b.Conversations))]

			clientMsgID, err := uuid.NewV7()
			var msgIDStr string
			if err == nil {
				msgIDStr = clientMsgID.String()
			} else {
				msgIDStr = uuid.NewString()
			}

			payloadBytes, _ := json.Marshal(sendMessagePayload{
				ConversationID: target.ConversationID,
				ReceiverID:     target.PartnerID,
				Content:        fmt.Sprintf("Hello from %s to %s at %s", b.Session.Username, target.PartnerID, time.Now().Format("15:04:05.000")),
			})

			env := wsEnvelope{
				Type:        "SEND_MESSAGE",
				ClientMsgID: msgIDStr,
				Timestamp:   time.Now().UnixMilli(),
				Payload:     payloadBytes,
			}

			envBytes, err := json.Marshal(env)
			if err != nil {
				log.Printf("[Bot] Marshal error: %v", err)
				continue
			}

			b.inFlight.Store(msgIDStr, time.Now())
			b.Tracker.RecordSent()

			select {
			case b.sendChan <- envBytes:
			default:
				b.Tracker.RecordError()
			}
		}
	}
}
````

## File: services/client-simulator/internal/config/config.go
````go
package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Bots        int           `yaml:"bots" env:"SIM_BOTS" envDefault:"50"`
	ConvsPerBot int           `yaml:"convs_per_bot" env:"SIM_CONVS" envDefault:"5"`
	Interval    time.Duration `yaml:"interval" env:"SIM_INTERVAL" envDefault:"200ms"`
	Duration    time.Duration `yaml:"duration" env:"SIM_DURATION" envDefault:"0s"`
	APIURL      string        `yaml:"api_url" env:"SIM_API_URL" envDefault:"http://localhost"`
	WSURL       string        `yaml:"ws_url" env:"SIM_WS_URL" envDefault:"ws://localhost/ws"`
	Password    string        `yaml:"password" env:"SIM_PASSWORD" envDefault:"Pass@123456"`
}

// LoadConfig nạp cấu hình theo thứ tự ưu tiên: Defaults -> File YAML -> Environment Variables
func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}

	// 1. Đọc file YAML nếu file tồn tại
	if path != "" {
		if file, err := os.Open(path); err == nil {
			defer file.Close()
			if err := yaml.NewDecoder(file).Decode(cfg); err != nil {
				return nil, fmt.Errorf("failed to decode yaml config: %w", err)
			}
		}
	}

	// 2. Ghi đè bằng biến môi trường (Environment Variables từ K8s ConfigMap)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return cfg, nil
}
````

## File: services/client-simulator/internal/conversation/manager.go
````go
package conversation

import (
	"bytes"
	"client-simulator/internal/auth"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Target struct {
	ConversationID string
	PartnerID      string
}

type ConversationResponse struct {
	ID      string `json:"id"`
	Members []struct {
		UserID string `json:"userId"`
	} `json:"members"`
}

type apiResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       json.RawMessage `json:"data"`
}

type Manager struct {
	baseURL    string
	httpClient *http.Client
}

func NewManager(baseURL string) *Manager {
	return &Manager{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// EnsureDirectConversations đảm bảo mỗi bot có đủ số lượng direct conversation với các bot khác song song
func (m *Manager) EnsureDirectConversations(bots []*auth.BotSession, desiredConvsPerBot int) map[string][]Target {
	botTargets := make(map[string][]Target)
	totalBots := len(bots)

	if desiredConvsPerBot >= totalBots {
		desiredConvsPerBot = totalBots - 1
	}
	if desiredConvsPerBot <= 0 {
		desiredConvsPerBot = 1
	}

	log.Printf("[Conversation] Setting up conversation graph concurrently: %d bots, target ~%d convs/bot...",
		totalBots, desiredConvsPerBot)

	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, 30) // 30 concurrent queries

	for i, currentBot := range bots {
		if currentBot == nil {
			continue
		}
		wg.Add(1)
		go func(idx int, b *auth.BotSession) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			existingMap := make(map[string]string)

			// 1. Lấy danh sách conversation hiện có của bot
			existingConvos, err := m.fetchUserConversations(b.AccessToken)
			if err == nil {
				for _, convo := range existingConvos {
					for _, member := range convo.Members {
						if member.UserID != b.UserID && member.UserID != "" {
							existingMap[member.UserID] = convo.ID
						}
					}
				}
			}

			// 2. Nếu chưa đủ desiredConvsPerBot, tạo thêm với các bot khác
			for step := 1; step <= desiredConvsPerBot; step++ {
				partnerIdx := (idx + step) % totalBots
				if partnerIdx == idx {
					continue
				}
				partnerBot := bots[partnerIdx]
				if partnerBot == nil {
					continue
				}

				if _, exists := existingMap[partnerBot.UserID]; exists {
					continue
				}

				// Gọi POST /conversations/direct để tạo
				convoID, err := m.createDirectConversation(b.AccessToken, partnerBot.UserID)
				if err == nil && convoID != "" {
					existingMap[partnerBot.UserID] = convoID
				}
			}

			// Gom danh sách target cho bot hiện tại
			targets := make([]Target, 0, len(existingMap))
			for partnerID, convoID := range existingMap {
				targets = append(targets, Target{
					ConversationID: convoID,
					PartnerID:      partnerID,
				})
			}

			mu.Lock()
			botTargets[b.UserID] = targets
			mu.Unlock()
		}(i, currentBot)
	}
	wg.Wait()

	log.Printf("[Conversation] Conversation graph established successfully in parallel.")
	return botTargets
}

func (m *Manager) fetchUserConversations(token string) ([]ConversationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/conversations", m.baseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var convos []ConversationResponse
	var rawResp apiResponse
	if err := json.Unmarshal(respBytes, &rawResp); err == nil && len(rawResp.Data) > 0 {
		_ = json.Unmarshal(rawResp.Data, &convos)
	}

	if len(convos) == 0 {
		_ = json.Unmarshal(respBytes, &convos)
	}

	return convos, nil
}

func (m *Manager) createDirectConversation(token, partnerID string) (string, error) {
	body, _ := json.Marshal(map[string]string{
		"partnerId": partnerID,
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/conversations/direct", m.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var created ConversationResponse
	var rawResp apiResponse
	if err := json.Unmarshal(respBytes, &rawResp); err == nil && len(rawResp.Data) > 0 {
		_ = json.Unmarshal(rawResp.Data, &created)
	}

	if created.ID == "" {
		_ = json.Unmarshal(respBytes, &created)
	}

	return created.ID, nil
}
````

## File: services/client-simulator/internal/metrics/tracker.go
````go
package metrics

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// Tracker quản lý số liệu đo lường hiệu năng và độ trễ (Latency)
type Tracker struct {
	totalSent         atomic.Int64
	totalAck          atomic.Int64
	totalDelivered    atomic.Int64
	totalErrors       atomic.Int64
	totalReconnects   atomic.Int64
	totalLatencyNanos atomic.Int64

	// Sample buffer để tính Percentiles cho Sender ACK (E2E Latency)
	samplesMu sync.Mutex
	samples   []float64 // Lưu latency bằng millisecond

	// Thống kê độ trễ mạng thuần túy (Network Delay / Ping-Pong RTT)
	netPingCount          atomic.Int64
	totalNetLatencyNanos  atomic.Int64
	netSamplesMu          sync.Mutex
	netSamples            []float64 // Lưu latency bằng millisecond
}

func NewTracker() *Tracker {
	return &Tracker{
		samples:    make([]float64, 0, 10000),
		netSamples: make([]float64, 0, 10000),
	}
}

func (t *Tracker) RecordSent() {
	t.totalSent.Add(1)
}

func (t *Tracker) RecordAck(d time.Duration) {
	t.totalAck.Add(1)
	nanos := d.Nanoseconds()
	t.totalLatencyNanos.Add(nanos)

	ms := float64(nanos) / 1e6
	t.samplesMu.Lock()
	if len(t.samples) < 50000 {
		t.samples = append(t.samples, ms)
	} else {
		// Ring-buffer giữ tối đa 50.000 samples gần nhất
		t.samples = t.samples[len(t.samples)-25000:]
		t.samples = append(t.samples, ms)
	}
	t.samplesMu.Unlock()
}

func (t *Tracker) RecordNetworkDelay(d time.Duration) {
	t.netPingCount.Add(1)
	nanos := d.Nanoseconds()
	t.totalNetLatencyNanos.Add(nanos)

	ms := float64(nanos) / 1e6
	t.netSamplesMu.Lock()
	if len(t.netSamples) < 50000 {
		t.netSamples = append(t.netSamples, ms)
	} else {
		t.netSamples = t.netSamples[len(t.netSamples)-25000:]
		t.netSamples = append(t.netSamples, ms)
	}
	t.netSamplesMu.Unlock()
}

func (t *Tracker) RecordDelivered() {
	t.totalDelivered.Add(1)
}

func (t *Tracker) RecordError() {
	t.totalErrors.Add(1)
}

func (t *Tracker) RecordReconnect() {
	t.totalReconnects.Add(1)
}

func (t *Tracker) GetStats() (sent, ack, delivered, errs, reconnects int64, avgMs float64, p50, p90, p95, p99, maxMs float64) {
	sent = t.totalSent.Load()
	ack = t.totalAck.Load()
	delivered = t.totalDelivered.Load()
	errs = t.totalErrors.Load()
	reconnects = t.totalReconnects.Load()

	if ack > 0 {
		avgMs = float64(t.totalLatencyNanos.Load()) / float64(ack) / 1e6
	}

	t.samplesMu.Lock()
	n := len(t.samples)
	if n > 0 {
		sorted := make([]float64, n)
		copy(sorted, t.samples)
		t.samplesMu.Unlock()

		sort.Float64s(sorted)
		p50 = sorted[int(float64(n)*0.50)]
		p90 = sorted[int(float64(n)*0.90)]
		p95 = sorted[int(float64(n)*0.95)]
		p99 = sorted[int(float64(n)*0.99)]
		maxMs = sorted[n-1]
	} else {
		t.samplesMu.Unlock()
	}

	return
}

func (t *Tracker) GetNetworkStats() (count int64, avgMs float64, p50, p90, p95, p99, maxMs float64) {
	count = t.netPingCount.Load()
	if count > 0 {
		avgMs = float64(t.totalNetLatencyNanos.Load()) / float64(count) / 1e6
	}

	t.netSamplesMu.Lock()
	n := len(t.netSamples)
	if n > 0 {
		sorted := make([]float64, n)
		copy(sorted, t.netSamples)
		t.netSamplesMu.Unlock()

		sort.Float64s(sorted)
		p50 = sorted[int(float64(n)*0.50)]
		p90 = sorted[int(float64(n)*0.90)]
		p95 = sorted[int(float64(n)*0.95)]
		p99 = sorted[int(float64(n)*0.99)]
		maxMs = sorted[n-1]
	} else {
		t.netSamplesMu.Unlock()
	}

	return
}

func (t *Tracker) PrintDashboard(activeBots, totalBots int, uptime time.Duration) {
	sent, ack, delivered, errs, reconnects, ackAvgMs, ackP50, ackP90, ackP95, ackP99, ackMaxMs := t.GetStats()
	pingCount, netAvgMs, netP50, netP90, netP95, netP99, netMaxMs := t.GetNetworkStats()

	secs := uptime.Seconds()
	var sendTps, ackTps float64
	if secs > 0 {
		sendTps = float64(sent) / secs
		ackTps = float64(ack) / secs
	}

	var ackRate float64
	if sent > 0 {
		ackRate = float64(ack) / float64(sent) * 100
	}

	fmt.Printf("\033[H\033[2J") // Clear console screen
	fmt.Println("========================= CHAT STRESS TEST DASHBOARD =========================")
	fmt.Printf("Active Bots:      %d / %d Online (WS Connected)\n", activeBots, totalBots)
	fmt.Printf("Elapsed Time:     %s\n", uptime.Round(time.Second))
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Printf("Messages Sent:    %-8d | Send Rate:       %.1f msgs/sec\n", sent, sendTps)
	fmt.Printf("ACKs Received:    %-8d | ACK Rate:        %.1f%% (%.1f acks/sec)\n", ack, ackRate, ackTps)
	fmt.Printf("Msg Delivered:    %-8d | Errors:          %d\n", delivered, errs)
	fmt.Printf("Reconnects:       %-8d | Ping Checks:     %d\n", reconnects, pingCount)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("📡 Network Delay / RTT (Client <-> Gateway Ping):")
	if pingCount > 0 {
		fmt.Printf("  - Average:      %6.2f ms | p50 (Median): %6.2f ms\n", netAvgMs, netP50)
		fmt.Printf("  - p90:          %6.2f ms | p95:          %6.2f ms\n", netP90, netP95)
		fmt.Printf("  - p99:          %6.2f ms | Max:          %6.2f ms\n", netP99, netMaxMs)
	} else {
		fmt.Println("  (Đang thu thập mẫu ping heartbeat...)")
	}
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("⚡ End-to-End Latency (Sender ACK):")
	if ack > 0 {
		fmt.Printf("  - Average:      %6.2f ms | p50 (Median): %6.2f ms\n", ackAvgMs, ackP50)
		fmt.Printf("  - p90:          %6.2f ms | p95:          %6.2f ms\n", ackP90, ackP95)
		fmt.Printf("  - p99:          %6.2f ms | Max:          %6.2f ms\n", ackP99, ackMaxMs)
		if pingCount > 0 && ackP50 >= netP50 {
			estProcessing := ackP50 - netP50
			fmt.Printf("  -> Ước tính thời gian xử lý Backend (Median E2E - Network): ~%.2f ms\n", estProcessing)
		}
	} else {
		fmt.Println("  (Chưa có message ACK)")
	}
	fmt.Println("==============================================================================")
	fmt.Println("Nhấn Ctrl+C để dừng kiểm thử.")
}
````

## File: services/notification-service/cmd/main.go
````go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"notification-service/internal/config"
	"notification-service/internal/consumer"
	"notification-service/internal/provider"
	"notification-service/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Starting Notification Service worker...")

	fcmProvider, _ := provider.NewFCMProvider(cfg.FCM.CredentialsFile)
	apnsProvider, _ := provider.NewAPNsProvider(
		cfg.APNs.KeyFile,
		cfg.APNs.KeyID,
		cfg.APNs.TeamID,
		cfg.APNs.Topic,
		cfg.APNs.Production,
	)

	uc := usecase.NewNotificationUsecase(fcmProvider, apnsProvider)

	notiConsumer, err := consumer.NewNotificationConsumer(
		cfg.NATS.URL,
		cfg.NATS.NotificationSubject,
		cfg.NATS.QueueGroup,
	)
	if err != nil {
		log.Fatalf("Failed to initialize NATS consumer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := notiConsumer.Start(ctx, uc.HandleNotification); err != nil {
			log.Printf("Notification consumer error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Notification Service...")
	cancel()
	notiConsumer.Close()
	log.Println("Notification Service exited cleanly")
}
````

## File: services/notification-service/configs/config.yaml
````yaml
nats:
  url: "nats://localhost:4222"
  notification_subject: "chat.notifications"
  queue_group: "notification-worker-group"

fcm:
  credentials_file: "configs/fcm-service-account.json"

apns:
  key_file: "configs/apns-key.p8"
  key_id: "KEY_ID"
  team_id: "TEAM_ID"
  topic: "com.chat.app"
  production: false
````

## File: services/notification-service/go.mod
````
module notification-service

go 1.23
````

## File: services/notification-service/internal/config/config.go
````go
package config

type Config struct {
	NATS NATSConfig `yaml:"nats"`
	FCM  FCMConfig  `yaml:"fcm"`
	APNs APNsConfig `yaml:"apns"`
}

type NATSConfig struct {
	URL                 string `yaml:"url"`
	NotificationSubject string `yaml:"notification_subject"`
	QueueGroup          string `yaml:"queue_group"`
}

type FCMConfig struct {
	CredentialsFile string `yaml:"credentials_file"`
}

type APNsConfig struct {
	KeyFile    string `yaml:"key_file"`
	KeyID      string `yaml:"key_id"`
	TeamID     string `yaml:"team_id"`
	Topic      string `yaml:"topic"`
	Production bool   `yaml:"production"`
}

func LoadConfig(path string) (*Config, error) {
	// TODO: Implement YAML/ENV config loader
	return &Config{
		NATS: NATSConfig{
			URL:                 "nats://localhost:4222",
			NotificationSubject: "chat.notifications",
			QueueGroup:          "notification-worker-group",
		},
	}, nil
}
````

## File: services/notification-service/internal/consumer/consumer.go
````go
package consumer

import (
	"context"
	"notification-service/internal/domain"
)

type NotificationHandler func(ctx context.Context, event *domain.NotificationEvent) error

type NotificationConsumer interface {
	Start(ctx context.Context, handler NotificationHandler) error
	Close() error
}

type natsNotificationConsumer struct {
	// TODO: nats.Conn / QueueSubscriber reference
}

func NewNotificationConsumer(natsURL string, subject string, queueGroup string) (NotificationConsumer, error) {
	return &natsNotificationConsumer{}, nil
}

func (c *natsNotificationConsumer) Start(ctx context.Context, handler NotificationHandler) error {
	// TODO: natsConn.QueueSubscribe(subject, queueGroup, func(msg *nats.Msg) { ... })
	return nil
}

func (c *natsNotificationConsumer) Close() error {
	return nil
}
````

## File: services/notification-service/internal/domain/notification.go
````go
package domain

import "time"

type DevicePlatform string

const (
	PlatformIOS     DevicePlatform = "IOS"
	PlatformAndroid DevicePlatform = "ANDROID"
	PlatformWeb     DevicePlatform = "WEB"
)

// NotificationEvent represents the message payload received from Kafka topic
type NotificationEvent struct {
	RecipientID    string    `json:"recipient_id"`
	SenderID       string    `json:"sender_id"`
	ConversationID string    `json:"conversation_id"`
	MessageID      string    `json:"message_id"`
	PreviewContent string    `json:"preview_content"`
	CreatedAt      time.Time `json:"created_at"`
}

// PushPayload represents normalized payload passed to APNs/FCM providers
type PushPayload struct {
	DeviceToken string
	Title       string
	Body        string
	Data        map[string]string
}
````

## File: services/notification-service/internal/provider/apns.go
````go
package provider

import (
	"context"
	"notification-service/internal/domain"
)

type APNsProvider struct {
	// TODO: Apple Push Notification HTTP/2 client
}

func NewAPNsProvider(keyFile, keyID, teamID, topic string, production bool) (PushProvider, error) {
	return &APNsProvider{}, nil
}

func (p *APNsProvider) SendPush(ctx context.Context, payload *domain.PushPayload) error {
	// TODO: Send push notification via APNs
	return nil
}
````

## File: services/notification-service/internal/provider/fcm.go
````go
package provider

import (
	"context"
	"notification-service/internal/domain"
)

type FCMProvider struct {
	// TODO: Firebase messaging client
}

func NewFCMProvider(credentialsFile string) (PushProvider, error) {
	return &FCMProvider{}, nil
}

func (p *FCMProvider) SendPush(ctx context.Context, payload *domain.PushPayload) error {
	// TODO: Send push notification via Firebase Cloud Messaging
	return nil
}
````

## File: services/notification-service/internal/provider/provider.go
````go
package provider

import (
	"context"
	"notification-service/internal/domain"
)

type PushProvider interface {
	SendPush(ctx context.Context, payload *domain.PushPayload) error
}
````

## File: services/notification-service/internal/usecase/notification_usecase.go
````go
package usecase

import (
	"context"
	"notification-service/internal/domain"
	"notification-service/internal/provider"
)

type NotificationUsecase struct {
	fcmProvider  provider.PushProvider
	apnsProvider provider.PushProvider
}

func NewNotificationUsecase(fcm, apns provider.PushProvider) *NotificationUsecase {
	return &NotificationUsecase{
		fcmProvider:  fcm,
		apnsProvider: apns,
	}
}

func (u *NotificationUsecase) HandleNotification(ctx context.Context, event *domain.NotificationEvent) error {
	return nil
}
````

## File: services/ws-gateway/internal/delivery/listener.go
````go
package delivery

import "context"

type DeliveryListener interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
````

## File: services/ws-gateway/internal/payload/ws_payload.go
````go
package payload

import (
	"encoding/json"

	"chat-system/pkg/contracts"
)

// WSEventType represents the type of WebSocket event
type WSEventType string

const (
	WSEventSendMessage   WSEventType = "SEND_MESSAGE"
	WSEventFailedToSend  WSEventType = "FAILED_TO_SEND"
	WSEventHeartbeat     WSEventType = "HEARTBEAT"
	WSEventHeartbeatAck  WSEventType = "HEARTBEAT_ACK"
)

func (t WSEventType) ToBrokerMessageType() (contracts.BrokerMessageType, bool) {
	switch t {
	case WSEventSendMessage:
		return contracts.BrokerEventMessageSubmitted, true
	default:
		return "", false
	}
}

// WSMessage represents the generic envelope exchanged with client over WebSocket
type WSMessage struct {
	Type        WSEventType     `json:"type"`
	ClientMsgID string          `json:"client_msg_id,omitempty"`
	Timestamp   int64           `json:"timestamp,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}

type SendMessagePayload struct {
	ConversationID string `json:"conversation_id"`
	ReceiverID     string `json:"receiver_id,omitempty"`
	Content        string `json:"content"`
}

type FailedToSendPayload struct {
	Error string `json:"error"`
}

type MessageDeliveryPayload struct {
	MessageID      string                      `json:"message_id"`
	ClientMsgID    string                      `json:"client_msg_id,omitempty"`
	ConversationID string                      `json:"conversation_id"`
	SenderID       string                      `json:"sender_id"`
	ReceiverID     string                      `json:"receiver_id,omitempty"`
	Content        string                      `json:"content"`
	Type           contracts.BrokerMessageType `json:"type"`
	Timestamp      int64                       `json:"timestamp"`
}

func (msDelivery *MessageDeliveryPayload) NewWSMessageFromDelivery() (*WSMessage, error) {
	bytes, err := json.Marshal(msDelivery)
	if err != nil {
		return nil, err
	}
	return &WSMessage{
		Type:        WSEventSendMessage,
		ClientMsgID: msDelivery.ClientMsgID,
		Timestamp:   msDelivery.Timestamp,
		Payload:     bytes,
	}, nil
}
````

## File: .gitignore
````
# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Code coverage profiles and other test artifacts
*.out
coverage.*
*.coverprofile
profile.cov

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
go.work.sum

# env file
.env

# Editor/IDE
# .idea/
# .vscode/

# Node.js & NestJS
node_modules/
dist/
build/
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pnpm-debug.log*
lerna-debug.log*
.nyc_output/
.temp/
.tmp/
pids/
*.pid
*.seed
*.pid.lock
````

## File: AGENTS.md
````markdown
# AI Backend Mentor & Distributed Systems Architect Guidelines

Bạn đóng vai trò là một **Staff/Senior Backend Architect & Systems Mentor**. Mục tiêu của bạn không chỉ là sinh mã nguồn mà là **đào tạo và định hình tư duy của người dùng trở thành một Senior Backend Engineer thực thụ**, có khả năng thiết kế và vận hành các hệ thống phân tán quy mô lớn (High Throughput, Low Latency, High Availability).

---

## 1. Phong cách giảng dạy & Phản hồi (Teaching & Communication Style)
- **Giao tiếp hoàn toàn bằng Tiếng Việt.**
- **Quy chuẩn Định dạng Văn bản (Formatting Rule):**
  - **TUYỆT ĐỐI KHÔNG** sử dụng các ký hiệu công thức toán LaTeX (ví dụ: `$...$`, `$$...$$`, `\to`, `\approx`, `\mu s`, `\mathbf`).
  - Luôn sử dụng ký tự văn bản thuần (plain text) hoặc Markdown tiêu chuẩn: dùng `->` thay cho `\to`, dùng `~` thay cho `\approx`, dùng `us` hoặc `micro-giây` thay cho `\mu s`, dùng text in đậm `**text**` thay cho `\mathbf{text}`.
- **Tư duy "Why trước How":** Trước khi đưa ra bất kỳ đoạn code nào, luôn giải thích bản chất kiến trúc, lý do tại sao chọn giải pháp này (Architecture Trade-offs), và các lựa chọn thay thế (Alternative Approaches).
- **Tư duy phản biện sắc bén (Devil's Advocate & Critical Thinking):**
  - Không đồng ý dễ dãi với các giải pháp chỉ chạy đúng ở "Happy Path". Luôn chỉ ra điểm yếu (flaws), rủi ro tiềm ẩn (hidden risks), và chi phí kỹ thuật (technical debt) của mỗi quyết định.
  - Đóng vai trò là người "chất vấn kỹ thuật" (Socratic questioning) để rèn luyện bản lĩnh kiến trúc cho người học.
- **Hướng dẫn từng bước (Step-by-Step Mentoring):** Cung cấp các đoạn code mẫu sạch (Clean Code), chuẩn mực, có comment giải thích rõ ràng để người dùng tự tay đọc hiểu và tích hợp vào dự án.
- **Không viết code thay toàn bộ khi người dùng muốn học:** Chia nhỏ vấn đề thành các bài học/mô-đun logic để người dùng nắm chắc từng phần.

---

## 2. Phương pháp luận Đào tạo Tư duy săn "Edge Cases" (Edge Case Discovery Framework)

Huấn luyện người dùng thói quen nhìn hệ thống qua **4 Lăng kính Sự cố (Failure Lenses)**:

### 🔍 Lăng kính 1: Thời gian & Bất đồng bộ (Time & Concurrency Lens)
- *Câu hỏi rèn luyện:* "Điều gì xảy ra nếu 2 request đến cùng mili-giây?", "Nếu message B đến trước message A thì sao?", "Nếu client ngắt kết nối đúng lúc server chuẩn bị ghi DB?".
- *Chủ đề:* Race conditions, Out-of-order delivery, Thundering Herd, Clock skew.

### 🔍 Lăng kính 2: Sự cố Mạng & Node (Network & Partial Failure Lens)
- *Câu hỏi rèn luyện:* "Nếu gọi downstream bị timeout nhưng thực chất downstream đã insert thành công thì sao?", "Nếu broker crash đúng lúc ack message?", "Network bị nghẽn (jitter/high latency) thì queue có bị phình to làm OOM không?".
- *Chủ đề:* Dual-write problem, At-least-once duplicates, Backpressure, Zombie requests, Split-brain.

### 🔍 Lăng kính 3: Giới hạn Dữ liệu & Tải đột biến (Data Boundary & Load Lens)
- *Câu hỏi rèn luyện:* "Một group chat có 1 triệu user gửi 1 tin nhắn thì fan-out thế nào?", "Nếu một partition key nhận 90% lượng traffic (Hotspot) thì xử lý sao?", "Nếu payload chứa ký tự độc hoặc payload rỗng/quá khổ?".
- *Chủ đề:* Hot partition, Tombstone saturation, Poison pill message, Memory exhaustion.

### 🔍 Lăng kính 4: Trạng thái & Vòng đời (State Machine & Lifecycle Lens)
- *Câu hỏi rèn luyện:* "Entity có thể nhảy từ State X sang State Z mà bỏ qua State Y không?", "Nếu user retry 5 lần thì có bị charge tiền / gửi tin nhắn 5 lần không?".
- *Chủ đề:* Idempotency Key, Finite State Machine (FSM) validation, Distributed Locking pitfalls.

---

## 3. Tiêu chuẩn kiến thức cần đào sâu (Deep-Dive Focus Areas)

Mọi giải pháp kỹ thuật cần được phân tích dưới lăng kính của một kỹ sư backend cấp cao:

### A. Hệ thống phân tán (Distributed Systems)
- **Đảm bảo tính nhất quán dữ liệu:** Mô hình Consistency (Eventual Consistency, Strong Consistency, Linearizability), CAP/PACELC Theorem.
- **Ngữ cảnh bất đồng bộ & Broker:** Out-of-order delivery, At-least-once delivery, Idempotency patterns, Message deduplication, Poison pill messages.
- **Xử lý sự cố mạng & Node failure:** Network partitions, Split-brain, Circuit Breakers, Exponential Backoff, Dead Letter Queues (DLQ).

### B. Cơ sở dữ liệu & Storage Engine
- **NoSQL / LSM-Tree vs B-Tree:** Hiểu sâu cơ chế ghi của Cassandra/ScyllaDB (CommitLog, MemTable, SSTable, Compaction), cách thiết kế Partition Key / Clustering Key để tránh Hotspot Node và Tombstone issues.
- **Caching & In-memory (Redis):** Cache Stampede, Cache Penetration, Cache Breakdown, Atomic operations (`SETNX`, Lua script), TTL Strategy.

### C. Concurrency & Performance trong Go
- **Concurrency Patterns:** Worker Pools, Fan-in/Fan-out, Context cancellation & Timeout propagation, Race conditions, Deadlocks, Memory leak do Goroutine leak.
- **Clean Architecture & Idiomatic Go:** Tách bạch Domain, Repository, Usecase, Delivery; Dependency Injection; Defensive programming; Error wrapping (`%w`).

---

## 4. Quy trình hướng dẫn tính năng mới (Feature Workflow)

Khi hướng dẫn một tính năng:
1. **Phân tích bài toán & "Bẫy" Edge Cases:** Liệt kê các kịch bản lỗi, race condition, hoặc điểm nghẽn có thể xảy ra ở quy mô production; đặt câu hỏi để người dùng tự nhận ra bẫy trước khi đưa ra đáp án.
2. **Thiết kế Contract / Data Flow:** Vẽ luồng dữ liệu (Data Flow) và cấu trúc dữ liệu.
3. **Cung cấp Code mẫu chuẩn Clean Code:** Code có type-safety, xử lý lỗi triệt để, idiomatic Go.
4. **Thử thách Phản biện (Challenge):** Đưa ra ít nhất 1 câu hỏi "phản biện góc tối" (Edge case challenge) để người dùng tự suy ngẫm, phản biện và củng cố kiến thức.
````

## File: deployments/k8s/00-external-infra.yaml
````yaml
apiVersion: v1
kind: Namespace
metadata:
  name: chat-system
  labels:
    name: chat-system
    tier: backend
---
apiVersion: v1
kind: Service
metadata:
  name: nats
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: redis
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: cassandra
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: postgres
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: otel-collector
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: pyroscope
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: loki
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
---
apiVersion: v1
kind: Service
metadata:
  name: prometheus
  namespace: chat-system
spec:
  type: ExternalName
  externalName: host.docker.internal
````

## File: deployments/k8s/01-configmap-secrets.yaml
````yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: chat-config
  namespace: chat-system
data:
  # Cấu hình kết nối Hạ tầng
  NATS_URL: "nats://nats.chat-system.svc.cluster.local:4222"
  REDIS_ADDR: "redis.chat-system.svc.cluster.local:6379"
  REDIS_DB: "0"

  # Cấu hình NATS Subjects
  NATS_INBOUND_SUBJECT: "inbound.messages"
  NATS_INBOUND_STREAM: "CHAT_INBOUND"
  NATS_INBOUND_CONSUMER_GROUP: "chat_workers"
  NATS_OUTBOUND_SUBJECT_PREFIX: "gateway.outbound"

  # Cấu hình Ports & Delivery Mode
  WS_PORT: "8080"
  GRPC_PORT: "50051"
  DELIVERY_MODE: "grpc" # "grpc" or "broker"
  GRPC_SERVICE_SUFFIX: ".ws-gateway-headless.chat-system.svc.cluster.local"
  PRESENCE_TTL: "60"

  # Cấu hình Chat Engine Worker Pool
  CHAT_WORKER_COUNT: "128"
  CHAT_WORKER_BUFFER_SIZE: "2048"
  CHAT_WORKER_DRAIN_TIMEOUT_SECONDS: "10"

  # Cấu hình Observability SOTA Stack
  OTEL_COLLECTOR_TARGET: "otel-collector.chat-system.svc.cluster.local:4317"
  PYROSCOPE_SERVER: "http://pyroscope.chat-system.svc.cluster.local:4040"
---
apiVersion: v1
kind: Secret
metadata:
  name: chat-secrets
  namespace: chat-system
type: Opaque
data:
  # Base64 of secret key / password (eg. secret "my-jwt-secret")
  # You can create by: echo -n 'my-jwt-secret' | base64
  JWT_ACCESS_TOKEN_SECRET: bXktand0LXNlY3JldA==
  REDIS_PASSWORD: ""
````

## File: deployments/k8s/02-ws-gateway-statefulset.yaml
````yaml
apiVersion: v1
kind: Service
metadata:
  name: ws-gateway-headless
  namespace: chat-system
  labels:
    app: ws-gateway
spec:
  clusterIP: None # ◄── Bắt buộc là None để K8s CoreDNS phân giải trực tiếp từng Pod
  selector:
    app: ws-gateway
  ports:
    - name: grpc-delivery
      port: 50051
      targetPort: 50051
    - name: ws-port
      port: 8080
      targetPort: 8080
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: ws-gateway
  namespace: chat-system
spec:
  serviceName: ws-gateway-headless # ◄── Liên kết với Service ở trên
  replicas: 3 # ◄── Sẽ tạo: ws-gateway-0, ws-gateway-1, ws-gateway-2
  podManagementPolicy: Parallel # ◄── Khởi động các Pod cùng lúc để scale nhanh
  selector:
    matchLabels:
      app: ws-gateway
  template:
    metadata:
      labels:
        app: ws-gateway
    spec:
      terminationGracePeriodSeconds: 30 # Cho phép 30s để đóng kết nối WebSocket êm ái
      containers:
        - name: ws-gateway
          image: chat-system/ws-gateway:latest
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 8080
              name: ws-port
            - containerPort: 50051
              name: grpc-delivery
            - containerPort: 9091
              name: metrics
          envFrom:
            - configMapRef:
                name: chat-config
            - secretRef:
                name: chat-secrets
          env:
            # Tự động gán tên Pod (ws-gateway-0, ws-gateway-1...) vào biến NODE_ID
            - name: SERVER_NODE_ID
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
          readinessProbe:
            httpGet:
              path: /health
              port: 8080
            initialDelaySeconds: 5
            periodSeconds: 10
          livenessProbe:
            httpGet:
              path: /health
              port: 8080
            initialDelaySeconds: 15
            periodSeconds: 20
          resources:
            requests:
              cpu: "250m"
              memory: "256Mi"
            limits:
              cpu: "1000m"
              memory: "1Gi"
````

## File: deployments/k8s/03-ws-gateway-ingress-service.yaml
````yaml
apiVersion: v1
kind: Service
metadata:
  name: ws-gateway-public
  namespace: chat-system
  labels:
    app: ws-gateway
spec:
  type: ClusterIP
  selector:
    app: ws-gateway
  ports:
    - name: websocket
      port: 80
      targetPort: 8080
      protocol: TCP
````

## File: deployments/k8s/04-chat-engine-deployment.yaml
````yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: chat-engine
  namespace: chat-system
  labels:
    app: chat-engine
spec:
  replicas: 2
  selector:
    matchLabels:
      app: chat-engine
  template:
    metadata:
      labels:
        app: chat-engine
    spec:
      containers:
        - name: chat-engine
          image: chat-system/chat-engine:latest
          imagePullPolicy: IfNotPresent
          envFrom:
            - configMapRef:
                name: chat-config
            - secretRef:
                name: chat-secrets
          env:
            # Tự động gán Pod name làm Worker ID
            - name: WORKER_ID
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: CASSANDRA_HOSTS
              value: "cassandra.chat-system.svc.cluster.local"
            - name: CASSANDRA_KEYSPACE
              value: "chat_system"
            - name: CASSANDRA_TABLE
              value: "messages"
          ports:
            - containerPort: 9092
              name: metrics
          resources:
            requests:
              cpu: "500m"
              memory: "512Mi"
            limits:
              cpu: "2000m"
              memory: "2Gi"
---
apiVersion: v1
kind: Service
metadata:
  name: chat-engine
  namespace: chat-system
spec:
  selector:
    app: chat-engine
  ports:
    - name: metrics
      port: 9092
      targetPort: 9092
````

## File: deployments/k8s/05-api-service-deployment.yaml
````yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-service
  namespace: chat-system
  labels:
    app: api-service
spec:
  replicas: 2
  selector:
    matchLabels:
      app: api-service
  template:
    metadata:
      labels:
        app: api-service
    spec:
      containers:
        - name: api-service
          image: chat-system/api-service:latest
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 3000
              name: http
          env:
            - name: NODE_ENV
              value: "production"
            - name: PORT
              value: "3000"
            - name: DB_HOST
              value: "postgres.chat-system.svc.cluster.local"
            - name: DB_PORT
              value: "5432"
            - name: DB_USERNAME
              value: "postgres"
            - name: DB_PASSWORD
              value: "postgrespassword"
            - name: DB_DATABASE
              value: "chat_db"
            - name: REDIS_HOST
              value: "redis.chat-system.svc.cluster.local"
            - name: REDIS_PORT
              value: "6379"
            - name: JWT_SECRET
              valueFrom:
                secretKeyRef:
                  name: chat-secrets
                  key: JWT_ACCESS_TOKEN_SECRET
          resources:
            requests:
              cpu: "100m"
              memory: "64Mi"
            limits:
              cpu: "1000m"
              memory: "256Mi"
---
apiVersion: v1
kind: Service
metadata:
  name: api-service
  namespace: chat-system
spec:
  type: ClusterIP
  selector:
    app: api-service
  ports:
    - name: http
      port: 3000
      targetPort: 3000
````

## File: knowledge/system_design/README.md
````markdown
# Kiến Trúc Hệ Thống & Tư Duy Phân Tán (System Architecture & Distributed Thinking)

Thư mục này lưu trữ các bài học, phân tích chuyên sâu về các vấn đề hóc búa, edge cases và giải pháp thực tế trong quá trình xây dựng hệ thống Chat phân tán quy mô lớn.

---

### Danh mục bài học:

1. [01. Cassandra UPSERT vs RDBMS & Bản chất Tầng Idempotency (Redis)](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/01_cassandra_upsert_vs_rdbms_and_idempotency.md)
   - *Nội dung:* Phân tích cơ chế xử lý trùng Primary Key của Cassandra (LSM-Tree), tại sao Cassandra không báo lỗi duplicate key, và 3 lý do sống còn cần tầng Redis Idempotency.

2. [02. Dual-Write Partial Failure, Poison Pill & State Reconciliation](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/02_dual_write_partial_failure_and_state_reconciliation.md)
   - *Nội dung:* Bài toán thất bại một phần khi lưu DB và gửi ACK, phân tích trade-off giữa `return err` vs `return nil`, hiện tượng Phantom Message và cơ chế State Reconciliation 2 chiều.

3. [03. Idempotency State Machine, Lock Rollback & Two-Phase TTL](file:///d:/BACKEND/PROJECTS/chat-system/knowledge/system_design/03_idempotency_state_machine_and_ttl_rollback.md)
   - *Nội dung:* Vấn đề kẹt khóa vĩnh viễn (Permanent Lockout), xử lý rollback khi DB lỗi, và giải pháp Two-Phase Idempotency State (Short TTL vs Long TTL) tự phục hồi khi server crash.
````

## File: pkg/contracts/events.go
````go
package contracts

import (
	"encoding/json"
	"time"
)

// BrokerMessageType represents standardized event types published across the message broker
type BrokerMessageType string

const (
	BrokerEventMessageSubmitted BrokerMessageType = "MESSAGE_SUBMITTED"
	BrokerEventMessageDelivered BrokerMessageType = "MESSAGE_DELIVERED"
	BrokerEventNotification     BrokerMessageType = "NOTIFICATION_DISPATCH"
)

// InboundBrokerEvent represents the message forwarded from WebSocket Gateway to Chat Engine via NATS
type InboundBrokerEvent struct {
	Type        BrokerMessageType `json:"type"`
	ClientMsgID string            `json:"client_msg_id"`
	SenderID    string            `json:"sender_id"`
	DeviceID    string            `json:"device_id"`
	GatewayNode string            `json:"gateway_node"`
	Payload     json.RawMessage   `json:"payload"`
	SentAt      time.Time         `json:"sent_at"`
}

// SendMessagePayload is the unmarshaled payload of a MESSAGE_SUBMITTED event
type SendMessagePayload struct {
	ConversationID string `json:"conversation_id"`
	ReceiverID     string `json:"receiver_id,omitempty"`
	Content        string `json:"content"`
}

// OutboundBrokerEvent represents the message routed from Chat Engine to a specific Gateway Node
type OutboundBrokerEvent struct {
	MessageID      string          `json:"message_id"`
	ClientMsgID    string          `json:"client_msg_id,omitempty"`
	ConversationID string          `json:"conversation_id"`
	SenderID       string          `json:"sender_id"`
	ReceiverID     string          `json:"receiver_id,omitempty"`
	Content        string            `json:"content"`
	Type           BrokerMessageType `json:"type"`
	Timestamp      int64             `json:"timestamp"`
	Payload        json.RawMessage `json:"payload,omitempty"`
}

// NotificationBrokerEvent represents payload routed to Notification Service when receiver is offline
type NotificationBrokerEvent struct {
	RecipientID    string `json:"recipient_id"`
	SenderID       string `json:"sender_id"`
	ConversationID string `json:"conversation_id"`
	ContentSnippet string `json:"content_snippet"`
	Timestamp      int64  `json:"timestamp"`
}
````

## File: pkg/contracts/topics.go
````go
package contracts

import "fmt"

// GatewayNodeSubject returns the dedicated subject for a specific WebSocket Gateway node
func GatewayNodeSubject(nodeID string) string {
	return fmt.Sprintf("chat.gateway.%s", nodeID)
}
````

## File: pkg/go.mod
````
module chat-system/pkg

go 1.25.0

require (
	github.com/grafana/pyroscope-go v1.4.2
	github.com/nats-io/nats.go v1.53.1
	github.com/prometheus/client_golang v1.24.1
	go.opentelemetry.io/otel v1.46.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.46.0
	go.opentelemetry.io/otel/sdk v1.46.0
	go.opentelemetry.io/otel/trace v1.46.0
	google.golang.org/grpc v1.83.2
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.11 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0 // indirect
	github.com/klauspost/compress v1.19.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.15 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.70.1 // indirect
	github.com/prometheus/procfs v0.21.1 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.46.0 // indirect
	go.opentelemetry.io/otel/metric v1.46.0 // indirect
	go.opentelemetry.io/proto/otlp v1.11.0 // indirect
	golang.org/x/crypto v0.55.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260819154853-08b0e4226688 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260819154853-08b0e4226688 // indirect
)
````

## File: pkg/nats/publisher.go
````go
package nats

import (
	"context"
	"encoding/json"
	"fmt"

	"chat-system/pkg/telemetry"

	"github.com/nats-io/nats.go"
)

// Publisher provides a generic, type-safe interface for publishing structured messages to NATS
type Publisher[T any] struct {
	conn           *nats.Conn
	defaultSubject string
}

// NewPublisher creates a new generic Publisher for the specified default subject
func NewPublisher[T any](conn *nats.Conn, defaultSubject string) *Publisher[T] {
	return &Publisher[T]{
		conn:           conn,
		defaultSubject: defaultSubject,
	}
}

// Publish serializes data to JSON and publishes it to the default subject
func (p *Publisher[T]) Publish(ctx context.Context, data T) error {
	return p.PublishToSubject(ctx, p.defaultSubject, data)
}

// PublishToSubject serializes data to JSON and publishes it to a specific custom subject
func (p *Publisher[T]) PublishToSubject(ctx context.Context, subject string, data T) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled before publishing: %w", err)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal message for subject %s: %w", subject, err)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    payload,
		Header:  make(nats.Header),
	}
	telemetry.InjectNATSTraceContext(ctx, msg)

	if err := p.conn.PublishMsg(msg); err != nil {
		return fmt.Errorf("failed to publish message to subject %s: %w", subject, err)
	}

	return nil
}
````

## File: pkg/nats/subscriber.go
````go
package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"chat-system/pkg/telemetry"

	"github.com/nats-io/nats.go"
)

// Handler represents a type-safe message handler function
type Handler[T any] func(ctx context.Context, data T) error

// Subscriber provides a generic, type-safe interface for consuming structured messages from NATS
type Subscriber[T any] struct {
	conn       *nats.Conn
	subject    string
	queueGroup string
	sub        *nats.Subscription
}

// NewSubscriber creates a new generic Subscriber for a subject (direct fan-out or unicast)
func NewSubscriber[T any](conn *nats.Conn, subject string) *Subscriber[T] {
	return &Subscriber[T]{
		conn:    conn,
		subject: subject,
	}
}

// NewQueueSubscriber creates a new generic Subscriber with a queue group for worker load balancing
func NewQueueSubscriber[T any](conn *nats.Conn, subject string, queueGroup string) *Subscriber[T] {
	return &Subscriber[T]{
		conn:       conn,
		subject:    subject,
		queueGroup: queueGroup,
	}
}

// Start begins consuming messages asynchronously and invokes the handler for each received message
func (s *Subscriber[T]) Start(ctx context.Context, handler Handler[T]) error {
	msgHandler := func(msg *nats.Msg) {
		msgCtx := telemetry.ExtractNATSTraceContext(ctx, msg)
		tracer := telemetry.Tracer("nats-subscriber")
		spanCtx, span := tracer.Start(msgCtx, fmt.Sprintf("nats.consume %s", s.subject))
		defer span.End()

		var data T
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("[NATS Subscriber] Failed to unmarshal message from %s: %v", s.subject, err)
			return
		}

		if err := handler(spanCtx, data); err != nil {
			log.Printf("[NATS Subscriber] Error handling message from %s: %v", s.subject, err)
		}
	}

	var sub *nats.Subscription
	var err error

	if s.queueGroup != "" {
		sub, err = s.conn.QueueSubscribe(s.subject, s.queueGroup, msgHandler)
	} else {
		sub, err = s.conn.Subscribe(s.subject, msgHandler)
	}

	if err != nil {
		return fmt.Errorf("failed to subscribe to subject %s: %w", s.subject, err)
	}

	s.sub = sub

	// Auto cleanup when context is canceled
	go func() {
		<-ctx.Done()
		_ = s.Close()
	}()

	return nil
}

// Close gracefully drains and removes the subscription
func (s *Subscriber[T]) Close() error {
	if s.sub != nil && s.sub.IsValid() {
		return s.sub.Drain()
	}
	return nil
}
````

## File: services/api-service/Dockerfile
````
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
````

## File: services/chat-engine/cmd/main.go
````go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
	"chat-worker/internal/config"
	"chat-worker/internal/consumer"
	"chat-worker/internal/dispatcher"
	"chat-worker/internal/migrator"
	"chat-worker/internal/presence"
	"chat-worker/internal/repository"
	"chat-worker/internal/usecase"

	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting Chat Engine worker: %s (workers=%d, buffer=%d)",
		cfg.Worker.ID, cfg.Worker.Workers, cfg.Worker.BufferSize)

	// Initialize Unified Telemetry (Tracing + Profiling + Metrics Server)
	shutdownTelemetry, err := telemetry.Setup(context.Background(), telemetry.SetupConfig{
		ServiceName:      "chat-engine",
		ServiceVersion:   "1.0.0",
		NodeID:           cfg.Worker.ID,
		CollectorTarget:  cfg.Telemetry.CollectorTarget,
		MetricsPort:      cfg.Telemetry.MetricsPort,
		ProfilerServer:   cfg.Profiler.ServerAddress,
		DisableTracing:   cfg.Telemetry.Disabled,
		DisableProfiling: cfg.Profiler.Disabled,
	})
	if err != nil {
		log.Printf("[Warning] Telemetry setup: %v", err)
	}
	defer shutdownTelemetry(context.Background())

	// Run Database Schema Migrations
	if err := migrator.Run(&cfg.Database); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Connect to NATS Broker
	nc, err := natsclient.Connect(cfg.NATS.URL, "chat-worker-"+cfg.Worker.ID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize Cassandra Database Session
	dbSession, err := repository.NewCassandraSession(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra session: %v", err)
	}
	defer dbSession.Close()

	// Initialize Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	defer redisClient.Close()

	// Initialize Repositories, Presence & Dispatcher
	idempotencyTTL := time.Duration(cfg.Redis.IdempotencyTTLSeconds) * time.Second
	idempotencyRepo := repository.NewRedisIdempotencyRepository(redisClient, idempotencyTTL)
	presenceReader := presence.NewPresenceReader(redisClient)
	messageRepo := repository.NewCassandraMessageRepository(dbSession, cfg.Database.Table)

	eventDispatcher, err := dispatcher.NewEventDispatcher(&cfg.Delivery, nc)
	if err != nil {
		log.Fatalf("Failed to initialize event dispatcher: %v", err)
	}
	defer eventDispatcher.Close()

	// Initialize Usecase
	dbTimeout := time.Duration(cfg.Database.TimeoutSeconds) * time.Second
	chatUsecase := usecase.NewChatUsecase(messageRepo, idempotencyRepo, presenceReader, eventDispatcher, dbTimeout)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize Inbound Consumer with Partitioned Worker Pool
	inboundConsumer, err := consumer.NewConsumer(consumer.Config{
		NC:          nc,
		Subject:     cfg.NATS.InboundSubject,
		QueueGroup:  cfg.NATS.InboundConsumerGroup,
		NumWorkers:  cfg.Worker.Workers,
		BufferSize:  cfg.Worker.BufferSize,
		ChatUsecase: chatUsecase,
	})
	if err != nil {
		log.Fatalf("Failed to initialize partitioned consumer: %v", err)
	}

	if err := inboundConsumer.Start(ctx); err != nil {
		log.Fatalf("Failed to start inbound consumer: %v", err)
	}

	log.Printf("Consumer started on subject '%s' with queue group '%s' and %d partitioned workers",
		cfg.NATS.InboundSubject, cfg.NATS.InboundConsumerGroup, cfg.Worker.Workers)

	// Wait for termination signal
	<-ctx.Done()

	log.Println("Received termination signal, shutting down Chat Engine...")

	// Graceful shutdown sequence with drain timeout
	drainTimeout := time.Duration(cfg.Worker.DrainTimeoutSeconds) * time.Second
	if drainTimeout <= 0 {
		drainTimeout = 10 * time.Second
	}
	drainCtx, cancel := context.WithTimeout(context.Background(), drainTimeout)
	defer cancel()

	if err := inboundConsumer.Stop(drainCtx); err != nil {
		log.Printf("Error during graceful consumer drain: %v", err)
	}

	log.Println("Chat Engine exited cleanly")
}
````

## File: services/chat-engine/configs/config.yaml
````yaml
worker:
  id: "chat-worker-01"
  workers: 32
  buffer_size: 1024
  drain_timeout_seconds: 10

nats:
  url: "nats://localhost:4222"
  inbound_subject: "chat.inbound"
  inbound_stream: "CHAT_INBOUND"
  inbound_consumer_group: "chat-worker-group"
  outbound_subject_prefix: "chat.gateway."
  notification_subject: "chat.notifications"

database:
  hosts:
    - "cassandra:9042"
  keyspace: "chat_system"
  table: "messages"
  timeout_seconds: 5

redis:
  addr: "localhost:6379"
  password: ""
  db: 0
  idempotency_ttl_seconds: 86400

delivery:
  mode: "grpc" # "grpc" or "broker"
  grpc_port: 50051
  grpc_service_suffix: ""
````

## File: services/chat-engine/Dockerfile
````
# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy các file go.mod cần thiết
COPY pkg/go.mod pkg/go.sum* ./pkg/
COPY services/chat-engine/go.mod services/chat-engine/go.sum* ./services/chat-engine/

# Khởi tạo workspace cục bộ cho chat-engine và pkg
RUN go work init ./pkg ./services/chat-engine

# Copy mã nguồn
COPY pkg/ ./pkg/
COPY services/chat-engine/ ./services/chat-engine/

# Build binary
WORKDIR /app/services/chat-engine
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /chat-worker ./cmd/main.go

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /chat-worker .
COPY --from=builder /app/services/chat-engine/configs ./configs

CMD ["./chat-worker"]
````

## File: services/chat-engine/go.mod
````
module chat-worker

go 1.25.0

require (
	github.com/caarlos0/env/v11 v11.4.1
	github.com/gocql/gocql v1.7.0
	github.com/golang-migrate/migrate/v4 v4.19.1
	github.com/nats-io/nats.go v1.53.1
	github.com/redis/go-redis/v9 v9.22.0
	go.opentelemetry.io/otel v1.46.0
	go.opentelemetry.io/otel/trace v1.46.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/klauspost/compress v1.19.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/nats-io/nkeys v0.4.15 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.70.0 // indirect
	go.opentelemetry.io/otel/metric v1.46.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.55.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
````

## File: services/chat-engine/internal/config/config.go
````go
package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Worker    WorkerConfig    `yaml:"worker"`
	NATS      NATSConfig      `yaml:"nats"`
	Database  DatabaseConfig  `yaml:"database"`
	Redis     RedisConfig     `yaml:"redis"`
	Delivery  DeliveryConfig  `yaml:"delivery"`
	Telemetry TelemetryConfig `yaml:"telemetry"`
	Profiler  ProfilerConfig  `yaml:"profiler"`
}

type TelemetryConfig struct {
	CollectorTarget string `yaml:"collector_target" env:"OTEL_COLLECTOR_TARGET" envDefault:"localhost:4317"`
	MetricsPort     int    `yaml:"metrics_port" env:"METRICS_PORT" envDefault:"9092"`
	Disabled        bool   `yaml:"disabled" env:"OTEL_DISABLED" envDefault:"false"`
}

type ProfilerConfig struct {
	ServerAddress string `yaml:"server_address" env:"PYROSCOPE_SERVER" envDefault:"http://localhost:4040"`
	Disabled      bool   `yaml:"disabled" env:"PYROSCOPE_DISABLED" envDefault:"false"`
}

type WorkerConfig struct {
	ID                  string `yaml:"id" env:"WORKER_ID" envDefault:"chat-worker-01"`
	Workers             int    `yaml:"workers" env:"CHAT_WORKER_COUNT" envDefault:"32"`
	BufferSize          int    `yaml:"buffer_size" env:"CHAT_WORKER_BUFFER_SIZE" envDefault:"1024"`
	DrainTimeoutSeconds int    `yaml:"drain_timeout_seconds" env:"CHAT_WORKER_DRAIN_TIMEOUT_SECONDS" envDefault:"10"`
}

type NATSConfig struct {
	URL                   string `yaml:"url" env:"NATS_URL" envDefault:"nats://localhost:4222"`
	InboundSubject        string `yaml:"inbound_subject" env:"NATS_INBOUND_SUBJECT" envDefault:"chat.inbound"`
	InboundStream         string `yaml:"inbound_stream" env:"NATS_INBOUND_STREAM" envDefault:"CHAT_INBOUND"`
	InboundConsumerGroup  string `yaml:"inbound_consumer_group" env:"NATS_INBOUND_CONSUMER_GROUP" envDefault:"chat_workers"`
	OutboundSubjectPrefix string `yaml:"outbound_subject_prefix" env:"NATS_OUTBOUND_SUBJECT_PREFIX" envDefault:"chat.gateway."`
	NotificationSubject   string `yaml:"notification_subject" env:"NATS_NOTIFICATION_SUBJECT" envDefault:"chat.notification"`
}

type DatabaseConfig struct {
	Hosts          []string `yaml:"hosts" env:"CASSANDRA_HOSTS" envSeparator:","`
	Keyspace       string   `yaml:"keyspace" env:"CASSANDRA_KEYSPACE" envDefault:"chat_system"`
	Table          string   `yaml:"table" env:"CASSANDRA_TABLE" envDefault:"messages"`
	TimeoutSeconds int      `yaml:"timeout_seconds" env:"CASSANDRA_TIMEOUT_SECONDS" envDefault:"5"`
}

type RedisConfig struct {
	Addr                  string `yaml:"addr" env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password              string `yaml:"password" env:"REDIS_PASSWORD"`
	DB                    int    `yaml:"db" env:"REDIS_DB" envDefault:"0"`
	IdempotencyTTLSeconds int    `yaml:"idempotency_ttl_seconds" env:"REDIS_IDEMPOTENCY_TTL" envDefault:"86400"`
}

type DeliveryConfig struct {
	Mode              string `yaml:"mode" env:"DELIVERY_MODE" envDefault:"grpc"` // "grpc" or "broker"
	GRPCPort          int    `yaml:"grpc_port" env:"GRPC_PORT" envDefault:"50051"`
	GRPCServiceSuffix string `yaml:"grpc_service_suffix" env:"GRPC_SERVICE_SUFFIX" envDefault:""`
}

var Cfg *Config

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	// 1. Đọc file YAML nếu có (làm baseline/default)
	if path != "" {
		data, err := os.ReadFile(path)
		if err == nil {
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("không thể parse file yaml config: %w", err)
			}
		}
	}

	// 2. Tự động đọc và ghi đè từ Environment Variables (Type-Safe qua caarlos0/env)
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("không thể parse biến môi trường: %w", err)
	}

	Cfg = &cfg

	if err := Cfg.Validate(); err != nil {
		return nil, fmt.Errorf("cấu hình không hợp lệ: %w", err)
	}

	return Cfg, nil
}

func (c *Config) Validate() error {
	if c.Worker.ID == "" {
		return fmt.Errorf("worker.id không được để trống")
	}
	if c.NATS.URL == "" {
		return fmt.Errorf("nats.url không được để trống")
	}
	if c.NATS.InboundSubject == "" {
		return fmt.Errorf("nats.inbound_subject không được để trống")
	}
	if len(c.Database.Hosts) == 0 {
		return fmt.Errorf("database.hosts không được để trống")
	}
	if c.Redis.Addr == "" {
		return fmt.Errorf("redis.addr không được để trống")
	}
	return nil
}
````

## File: services/chat-engine/internal/consumer/consumer.go
````go
package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
	"chat-system/pkg/worker"
	"chat-worker/internal/usecase"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Consumer wraps NATS Queue Subscriber and dispatches to a generic Key-based Partitioned Worker Pool
type Consumer struct {
	subscriber  *natsclient.Subscriber[contracts.InboundBrokerEvent]
	chatUsecase usecase.ChatUsecase
	pool        *worker.PartitionedPool[contracts.InboundBrokerEvent]
}

// Config chứa cấu hình cần thiết để khởi tạo Consumer
type Config struct {
	NC          *nats.Conn
	Subject     string
	QueueGroup  string
	NumWorkers  int
	BufferSize  int
	ChatUsecase usecase.ChatUsecase
}

// NewConsumer creates a new Inbound Consumer with generic partitioned worker pool support
func NewConsumer(cfg Config) (*Consumer, error) {
	c := &Consumer{
		subscriber:  natsclient.NewQueueSubscriber[contracts.InboundBrokerEvent](cfg.NC, cfg.Subject, cfg.QueueGroup),
		chatUsecase: cfg.ChatUsecase,
	}

	pool, err := worker.NewPartitionedPool(worker.Config[contracts.InboundBrokerEvent]{
		NumWorkers:   cfg.NumWorkers,
		BufferSize:   cfg.BufferSize,
		KeyExtractor: extractConversationKey,
		Handler:      c.handleWorkerProcess,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create worker pool: %w", err)
	}

	c.pool = pool
	return c, nil
}

// Start begins consuming messages and starting worker goroutines
func (c *Consumer) Start(ctx context.Context) error {
	// 1. Khởi động worker pool
	c.pool.Start()

	// 2. Wrap handler với telemetry stage "worker_dispatch"
	instrumentedHandler := telemetry.InstrumentHandler(telemetry.HandlerConfig{
		Service:   "chat-engine",
		Mode:      "engine",
		Stage:     "worker_dispatch",
		EventType: "inbound",
	}, c.dispatchMessage)

	return c.subscriber.Start(ctx, instrumentedHandler)
}

// dispatchMessage trích xuất span context và định tuyến message vào worker channel
func (c *Consumer) dispatchMessage(ctx context.Context, event contracts.InboundBrokerEvent) error {
	if err := c.pool.Submit(ctx, event); err != nil {
		log.Printf("[Consumer] Failed to submit message to worker pool msg_id=%s: %v", event.ClientMsgID, err)
		return err
	}
	return nil
}

// handleWorkerProcess được thực thi bên trong Worker Goroutine được phân bổ
func (c *Consumer) handleWorkerProcess(ctx context.Context, event contracts.InboundBrokerEvent) error {
	tracer := telemetry.Tracer("chat-engine")
	spanCtx, span := tracer.Start(ctx, "chat_engine.process",
		trace.WithAttributes(
			attribute.String("client_msg_id", event.ClientMsgID),
			attribute.String("sender_id", event.SenderID),
			attribute.String("gateway_node", event.GatewayNode),
		),
	)
	defer span.End()

	if err := c.chatUsecase.ProcessInboundMessage(spanCtx, event); err != nil {
		log.Printf("[Consumer] Error processing event client_msg_id=%s: %v", event.ClientMsgID, err)
		return err
	}
	return nil
}

// Stop gracefully closes NATS subscription and drains the worker pool
func (c *Consumer) Stop(drainCtx context.Context) error {
	log.Println("[Consumer] Stopping NATS subscription...")
	if err := c.subscriber.Close(); err != nil {
		log.Printf("[Consumer] Warning: error closing NATS subscriber: %v", err)
	}

	log.Println("[Consumer] Draining worker pool...")
	return c.pool.Stop(drainCtx)
}

// extractConversationKey trích xuất nhanh conversation_id để băm định tuyến
func extractConversationKey(event contracts.InboundBrokerEvent) string {
	type quickPayload struct {
		ConversationID string `json:"conversation_id"`
	}

	var pLoad quickPayload
	if len(event.Payload) > 0 {
		if err := json.Unmarshal(event.Payload, &pLoad); err == nil && pLoad.ConversationID != "" {
			return pLoad.ConversationID
		}
	}

	if event.ClientMsgID != "" {
		return event.ClientMsgID
	}
	return event.SenderID
}
````

## File: services/chat-engine/internal/dispatcher/dispatcher.go
````go
package dispatcher

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	"chat-worker/internal/config"

	"github.com/nats-io/nats.go"
)

// EventDispatcher định nghĩa giao diện chung cho việc định tuyến outbound event (gRPC hoặc NATS Broker)
type EventDispatcher interface {
	// SendAckToSender gửi sự kiện ACK về đúng Node Gateway mà người gửi đang kết nối
	SendAckToSender(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error

	// DispatchToGateway chuyển tiếp tin nhắn tới đúng Node Gateway mà người nhận đang kết nối
	DispatchToGateway(ctx context.Context, gatewayNode string, event contracts.OutboundBrokerEvent) error

	// Close giải phóng tài nguyên kết nối (gRPC conns, etc.)
	Close() error
}

// NewEventDispatcher khởi tạo dispatcher tương ứng với DeliveryMode đã cấu hình
func NewEventDispatcher(cfg *config.DeliveryConfig, nc *nats.Conn) (EventDispatcher, error) {
	switch cfg.Mode {
	case "grpc":
		log.Printf("[Dispatcher] Initializing gRPC Event Dispatcher (default port: %d, suffix: '%s')",
			cfg.GRPCPort, cfg.GRPCServiceSuffix)
		return NewGRPCEventDispatcher(cfg, nc), nil
	case "broker":
		log.Println("[Dispatcher] Initializing NATS Broker Event Dispatcher")
		return NewNATSEventDispatcher(nc), nil
	default:
		return nil, fmt.Errorf("unsupported delivery mode '%s', must be 'grpc' or 'broker'", cfg.Mode)
	}
}
````

## File: services/chat-engine/internal/presence/presence.go
````go
package presence

import (
	"context"
	"fmt"
	"time"

	"chat-system/pkg/telemetry"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("chat-worker/presence")

type PresenceReader interface {
	GetUserRoutes(ctx context.Context, userID string) (map[string]string, error)
	IsUserOnline(ctx context.Context, userID string) (bool, error)
}

type redisPresenceReader struct {
	client *redis.Client
}

func NewPresenceReader(client *redis.Client) PresenceReader {
	return &redisPresenceReader{
		client: client,
	}
}

func presenceKey(userID string) string {
	return fmt.Sprintf("presence:%s", userID)
}

// GetUserRoutes retrieves all active device -> gatewayNode mappings for a user.
// Returns an empty map if user is offline or has no active devices.
func (r *redisPresenceReader) GetUserRoutes(ctx context.Context, userID string) (map[string]string, error) {
	key := presenceKey(userID)
	ctx, span := tracer.Start(ctx, "Redis.GetUserRoutes",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "hgetall"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.user_id", userID),
		),
	)
	defer span.End()

	startTime := time.Now()
	res, err := r.client.HGetAll(ctx, key).Result()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "hgetall", status).Observe(duration)
		return nil, err
	}

	span.SetAttributes(attribute.Int("chat.devices_count", len(res)))
	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "hgetall", status).Observe(duration)

	return res, nil
}

// IsUserOnline returns true if user has at least one active device online in Redis.
func (r *redisPresenceReader) IsUserOnline(ctx context.Context, userID string) (bool, error) {
	key := presenceKey(userID)
	ctx, span := tracer.Start(ctx, "Redis.IsUserOnline",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "hlen"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.user_id", userID),
		),
	)
	defer span.End()

	startTime := time.Now()
	count, err := r.client.HLen(ctx, key).Result()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "hlen", status).Observe(duration)
		return false, err
	}

	online := count > 0
	span.SetAttributes(attribute.Bool("chat.is_online", online))
	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "hlen", status).Observe(duration)

	return online, nil
}
````

## File: services/chat-engine/internal/repository/idempotency_repo.go
````go
package repository

import (
	"context"
	"fmt"
	"time"

	"chat-system/pkg/telemetry"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type IdempotencyRepository interface {
	AcquireLock(ctx context.Context, clientMsgID string) (bool, error)
	ReleaseLock(ctx context.Context, clientMsgID string) error
}

type redisIdempotencyRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisIdempotencyRepository(client *redis.Client, ttl time.Duration) IdempotencyRepository {
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return &redisIdempotencyRepository{
		client: client,
		ttl:    ttl,
	}
}

func (r *redisIdempotencyRepository) AcquireLock(ctx context.Context, clientMsgID string) (bool, error) {
	if clientMsgID == "" {
		return false, fmt.Errorf("clientMsgID cannot be empty")
	}

	key := fmt.Sprintf("idempotency:msg:%s", clientMsgID)

	ctx, span := tracer.Start(ctx, "Redis.AcquireLock",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "setnx"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.client_msg_id", clientMsgID),
		),
	)
	defer span.End()

	startTime := time.Now()
	// Lệnh SET key value NX EX ttl là thao tác nguyên tử (Atomic) trong Redis
	acquired, err := r.client.SetNX(ctx, key, "PROCESSED", r.ttl).Result()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "setnx", status).Observe(duration)
		return false, fmt.Errorf("redis setnx failed for key %s: %w", key, err)
	}

	span.SetAttributes(attribute.Bool("chat.lock_acquired", acquired))
	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "setnx", status).Observe(duration)

	return acquired, nil
}

func (r *redisIdempotencyRepository) ReleaseLock(ctx context.Context, clientMsgID string) error {
	if clientMsgID == "" {
		return nil
	}

	key := fmt.Sprintf("idempotency:msg:%s", clientMsgID)

	ctx, span := tracer.Start(ctx, "Redis.ReleaseLock",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "redis"),
			attribute.String("db.operation", "del"),
			attribute.String("db.redis.key", key),
			attribute.String("chat.client_msg_id", clientMsgID),
		),
	)
	defer span.End()

	startTime := time.Now()
	err := r.client.Del(ctx, key).Err()
	duration := time.Since(startTime).Seconds()

	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("redis", "del", status).Observe(duration)
		return fmt.Errorf("redis del failed for key %s: %w", key, err)
	}

	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("redis", "del", status).Observe(duration)

	return nil
}
````

## File: services/chat-engine/internal/repository/message_repo.go
````go
package repository

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"chat-system/pkg/telemetry"
	"chat-worker/internal/config"
	"chat-worker/internal/domain"

	"github.com/gocql/gocql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("chat-worker/repository")

// MessageRepository defines the data access contract for messages
type MessageRepository interface {
	SaveMessage(ctx context.Context, msg *domain.Message) error
	GetMessagesByConversation(ctx context.Context, conversationID string, limit int) ([]*domain.Message, error)
}

// CassandraMessageRepository implements MessageRepository using Apache Cassandra
type CassandraMessageRepository struct {
	session *gocql.Session
	table   string
}

// NewCassandraSession creates and initializes a Cassandra session based on configuration
func NewCassandraSession(cfg *config.DatabaseConfig) (*gocql.Session, error) {
	if len(cfg.Hosts) == 0 {
		return nil, fmt.Errorf("cassandra hosts list is empty")
	}

	var hostList []string
	port := 9042

	for _, h := range cfg.Hosts {
		if strings.Contains(h, ":") {
			host, portStr, err := net.SplitHostPort(h)
			if err == nil {
				hostList = append(hostList, host)
				if parsedPort, err := strconv.Atoi(portStr); err == nil && parsedPort > 0 {
					port = parsedPort
				}
				continue
			}
		}
		hostList = append(hostList, h)
	}

	cluster := gocql.NewCluster(hostList...)
	cluster.Port = port
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 10 * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cassandra cluster: %w", err)
	}

	return session, nil
}

// NewCassandraMessageRepository creates a new instance of CassandraMessageRepository
func NewCassandraMessageRepository(session *gocql.Session, table string) *CassandraMessageRepository {
	return &CassandraMessageRepository{
		session: session,
		table:   table,
	}
}

// SaveMessage inserts a new message into Cassandra with OpenTelemetry Child Span
func (r *CassandraMessageRepository) SaveMessage(ctx context.Context, msg *domain.Message) error {
	ctx, span := tracer.Start(ctx, "Cassandra.SaveMessage",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "cassandra"),
			attribute.String("db.operation", "insert"),
			attribute.String("db.sql.table", r.table),
			attribute.String("chat.conversation_id", msg.ConversationID),
			attribute.String("chat.message_id", msg.ID),
			attribute.String("chat.sender_id", msg.SenderID),
		),
	)
	defer span.End()

	startTime := time.Now()
	query := fmt.Sprintf(`
		INSERT INTO %s (conversation_id, id, sender_id, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, r.table)

	err := r.session.Query(query,
		msg.ConversationID,
		msg.ID,
		msg.SenderID,
		msg.Content,
		msg.CreatedAt,
		msg.UpdatedAt,
	).WithContext(ctx).Exec()

	duration := time.Since(startTime).Seconds()
	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "OK")
	}

	telemetry.DatabaseLatency.WithLabelValues("cassandra", "insert", status).Observe(duration)
	return err
}

// GetMessagesByConversation retrieves recent messages for a conversation with OpenTelemetry Child Span
func (r *CassandraMessageRepository) GetMessagesByConversation(ctx context.Context, conversationID string, limit int) ([]*domain.Message, error) {
	if limit <= 0 {
		return nil, nil
	}

	ctx, span := tracer.Start(ctx, "Cassandra.GetMessagesByConversation",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("db.system", "cassandra"),
			attribute.String("db.operation", "select"),
			attribute.String("db.sql.table", r.table),
			attribute.String("chat.conversation_id", conversationID),
			attribute.Int("db.limit", limit),
		),
	)
	defer span.End()

	startTime := time.Now()
	query := fmt.Sprintf(`
		SELECT conversation_id, id, sender_id, content, created_at, updated_at
		FROM %s
		WHERE conversation_id = ?
		LIMIT ?
	`, r.table)

	iter := r.session.Query(query, conversationID, limit).WithContext(ctx).Iter()
	defer iter.Close()

	var messages []*domain.Message
	var convID, id, senderID, content string
	var createdAt, updatedAt time.Time

	for iter.Scan(&convID, &id, &senderID, &content, &createdAt, &updatedAt) {
		messages = append(messages, &domain.Message{
			ConversationID: convID,
			ID:             id,
			SenderID:       senderID,
			Content:        content,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		})
	}

	err := iter.Close()
	duration := time.Since(startTime).Seconds()
	status := "success"
	if err != nil {
		status = "error"
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		telemetry.DatabaseLatency.WithLabelValues("cassandra", "select", status).Observe(duration)
		return nil, fmt.Errorf("failed to query messages: %w", err)
	}

	span.SetStatus(codes.Ok, "OK")
	telemetry.DatabaseLatency.WithLabelValues("cassandra", "select", status).Observe(duration)
	return messages, nil
}
````

## File: services/chat-engine/internal/usecase/chat_usecase.go
````go
package usecase

import (
	"chat-system/pkg/contracts"
	"chat-worker/internal/dispatcher"
	"chat-worker/internal/domain"
	"chat-worker/internal/presence"
	"chat-worker/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ChatUsecase interface {
	ProcessInboundMessage(ctx context.Context, event contracts.InboundBrokerEvent) error
}
type chatUsecase struct {
	messageRepo     repository.MessageRepository
	idempotencyRepo repository.IdempotencyRepository
	presenceReader  presence.PresenceReader
	dispatcher      dispatcher.EventDispatcher
	dbTimeout       time.Duration
}

func NewChatUsecase(
	messageRepo repository.MessageRepository,
	idempotencyRepo repository.IdempotencyRepository,
	presenceReader presence.PresenceReader,
	dispatcher dispatcher.EventDispatcher,
	dbTimeout time.Duration,
) ChatUsecase {
	if dbTimeout <= 0 {
		dbTimeout = 5 * time.Second
	}
	return &chatUsecase{
		messageRepo:     messageRepo,
		idempotencyRepo: idempotencyRepo,
		presenceReader:  presenceReader,
		dispatcher:      dispatcher,
		dbTimeout:       dbTimeout,
	}
}

// ProcessInboundMessage xử lý sự kiện tin nhắn gửi đến từ WebSocket Gateway
func (u *chatUsecase) ProcessInboundMessage(ctx context.Context, event contracts.InboundBrokerEvent) error {
	// 1. Unmarshal payload tin nhắn từ InboundBrokerEvent
	var payload contracts.SendMessagePayload
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal message payload: %w", err)
	}
	if payload.ConversationID == "" || payload.Content == "" {
		return fmt.Errorf("invalid payload: conversation_id or content is empty")
	}

	now := time.Now().UTC()

	// 2. IDEMPOTENCY CHECK (Atomic SetNX trên Redis)
	isNew, err := u.idempotencyRepo.AcquireLock(ctx, event.ClientMsgID)
	if err != nil {
		// Nếu Redis lỗi (Redis down), log cảnh báo nhưng cho phép luồng tiếp tục (Graceful Degradation)
		log.Printf("[Usecase] WARN: Redis idempotency check failed: %v", err)
	} else if !isNew {
		// Tin nhắn TRÙNG LẶP: Bỏ qua ghi DB, gửi lại Sender ACK cho Client
		log.Printf("[Usecase] Duplicate message detected: client_msg_id=%s. Resending Sender ACK only.", event.ClientMsgID)

		ackEvent := contracts.OutboundBrokerEvent{
			MessageID:      event.ClientMsgID,
			ClientMsgID:    event.ClientMsgID,
			ConversationID: payload.ConversationID,
			SenderID:       event.SenderID,
			ReceiverID:     event.SenderID,
			Type:           contracts.BrokerEventMessageSubmitted,
			Timestamp:      now.UnixMilli(),
		}
		_ = u.dispatcher.SendAckToSender(ctx, event.GatewayNode, ackEvent)
		return nil
	}

	// 3. Chuyển đổi sang Domain Entity
	msg := &domain.Message{
		ConversationID: payload.ConversationID,
		ID:             event.ClientMsgID,
		SenderID:       event.SenderID,
		Content:        payload.Content,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// 4. Giới hạn timeout khi ghi vào Cassandra (Defensive Concurrency - Configurable)
	dbCtx, cancel := context.WithTimeout(ctx, u.dbTimeout)
	defer cancel()
	if err := u.messageRepo.SaveMessage(dbCtx, msg); err != nil {
		// Rollback/Release lock trên Redis nếu ghi DB thất bại để cho phép Client retry
		if relErr := u.idempotencyRepo.ReleaseLock(ctx, event.ClientMsgID); relErr != nil {
			log.Printf("[Usecase] ERROR: Failed to release idempotency lock for client_msg_id=%s: %v", event.ClientMsgID, relErr)
		}
		return fmt.Errorf("persistence failed: %w", err)
	}
	log.Printf("[Usecase] Message persisted successfully: msg_id=%s, conversation_id=%s, sender=%s",
		msg.ID, msg.ConversationID, msg.SenderID)

	// 5. Gửi Sender ACK về cho người gửi (Client A)
	ackEvent := contracts.OutboundBrokerEvent{
		MessageID:      msg.ID,
		ClientMsgID:    event.ClientMsgID,
		ConversationID: msg.ConversationID,
		SenderID:       msg.SenderID,
		ReceiverID:     msg.SenderID,
		Type:           contracts.BrokerEventMessageSubmitted,
		Timestamp:      now.UnixMilli(),
	}
	if err := u.dispatcher.SendAckToSender(ctx, event.GatewayNode, ackEvent); err != nil {
		// Log cảnh báo nhưng không làm fail cả luồng vì tin nhắn đã được lưu DB an toàn
		log.Printf("[Usecase] WARN: Failed to send Sender ACK: %v", err)
	}

	// 6. OUTBOUND DELIVERY: Tra cứu Presence và chuyển tiếp tin nhắn tới người nhận
	receiverID := payload.ReceiverID
	if receiverID != "" && receiverID != event.SenderID {
		routes, err := u.presenceReader.GetUserRoutes(ctx, receiverID)
		if err != nil {
			log.Printf("[Usecase] ERROR: Failed to query presence for receiver %s: %v", receiverID, err)
		}

		if len(routes) > 0 {
			// Gom nhóm theo unique gatewayNode để tránh gửi trùng lặp nếu 1 node có nhiều device
			dispatchedNodes := make(map[string]bool)
			for _, gatewayNode := range routes {
				if dispatchedNodes[gatewayNode] {
					continue
				}
				dispatchedNodes[gatewayNode] = true

				outboundMsg := contracts.OutboundBrokerEvent{
					MessageID:      msg.ID,
					ClientMsgID:    event.ClientMsgID,
					ConversationID: msg.ConversationID,
					SenderID:       msg.SenderID,
					ReceiverID:     receiverID,
					Content:        msg.Content,
					Type:           contracts.BrokerEventMessageSubmitted,
					Timestamp:      now.UnixMilli(),
				}

				if err := u.dispatcher.DispatchToGateway(ctx, gatewayNode, outboundMsg); err != nil {
					log.Printf("[Usecase] ERROR: Failed to dispatch message to receiver gateway %s: %v", gatewayNode, err)
				}
			}
		} else {
			log.Printf("[Usecase] Receiver %s is offline. Skipping realtime push.", receiverID)
		}
	}

	return nil
}
````

## File: services/notification-service/Dockerfile
````
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
````

## File: services/ws-gateway/configs/config.yaml
````yaml
server:
  port: 8080
  node_id: "gateway-node-01"
  delivery_mode: "grpc"

redis:
  addr: "localhost:6379"
  password: ""
  db: 0

nats:
  url: "nats://localhost:4222"
  inbound_subject: "chat.inbound"
  outbound_subject_prefix: "chat.gateway."

grpc:
  port: 50051
  advertised_addr: "localhost:50051"
````

## File: services/ws-gateway/Dockerfile
````
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
````

## File: services/ws-gateway/internal/connection/client.go
````go
package connection

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"chat-system/pkg/contracts"
	"chat-system/pkg/telemetry"
	"ws-gateway/internal/config"
	"ws-gateway/internal/payload"

	"github.com/gorilla/websocket"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Client represents a single active WebSocket connection
type Client struct {
	UserID    string
	DeviceID  string
	SendChan  chan *payload.WSMessage
	Conn      *websocket.Conn
	Hub       *Hub
	closeOnce sync.Once
}

// ReadPump handles reading messages from the WebSocket connection
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.UnregisterClient(c)
		c.closeConnSafely()
	}()

	c.Conn.SetReadLimit(int64(config.Cfg.Ws.MaxMessageSize))
	pongWait := time.Duration(config.Cfg.Ws.PongWait) * time.Second
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))

	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		var msg payload.WSMessage
		err := c.Conn.ReadJSON(&msg)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ReadPump] error reading message from user %s: %v", c.UserID, err)
			}
			break
		}

		c.handleIncomingMessage(&msg)
	}
}

func (c *Client) handleIncomingMessage(msg *payload.WSMessage) {
	switch msg.Type {
	case payload.WSEventHeartbeat:
		// 1. Phản hồi ngay HEARTBEAT_ACK về cho client đo Round-Trip Time (Network Delay)
		ackMsg := &payload.WSMessage{
			Type:        payload.WSEventHeartbeatAck,
			ClientMsgID: msg.ClientMsgID,
			Timestamp:   msg.Timestamp,
		}
		select {
		case c.SendChan <- ackMsg:
		default:
		}

		// 2. Đo Inbound Network Delay và ghi nhận lên Prometheus
		if msg.Timestamp > 0 {
			diff := time.Now().UnixMilli() - msg.Timestamp
			if diff >= 0 && diff < 60000 {
				telemetry.ClientNetworkLatency.WithLabelValues(config.Cfg.Server.NodeID).Observe(float64(diff) / 1000.0)
			}
		}

		// 3. Cập nhật presence bất đồng bộ
		go func(c *Client) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := c.Hub.presence.Heartbeat(ctx, c.UserID, c.DeviceID, config.Cfg.Server.NodeID, time.Duration(config.Cfg.Pres.TTL)*time.Second); err != nil {
				log.Printf("Error sending heartbeat for user %s: %v", c.UserID, err)
			}
		}(c)
	case payload.WSEventSendMessage:
		brokerMessageType, ok := msg.Type.ToBrokerMessageType()
		if !ok {
			log.Printf("Unhandled message type: %s", msg.Type)
			return
		}

		tracer := telemetry.Tracer("ws-gateway")
		traceCtx, span := tracer.Start(context.Background(), "ws.receive_message",
			trace.WithAttributes(
				attribute.String("client_msg_id", msg.ClientMsgID),
				attribute.String("sender_id", c.UserID),
				attribute.String("device_id", c.DeviceID),
				attribute.String("gateway_node", config.Cfg.Server.NodeID),
			),
		)
		defer span.End()

		inboundEvent := contracts.InboundBrokerEvent{
			Type:        brokerMessageType,
			ClientMsgID: msg.ClientMsgID,
			SenderID:    c.UserID,
			DeviceID:    c.DeviceID,
			GatewayNode: config.Cfg.Server.NodeID,
			Payload:     msg.Payload,
			SentAt:      time.Now().UTC(),
		}
		pubCtx, cancel := context.WithTimeout(traceCtx, 2*time.Second)
		defer cancel()

		if err := c.Hub.producer.Publish(pubCtx, inboundEvent); err != nil {
			c.sendErrorMessage(msg.ClientMsgID, "Failed to send message")
			return
		}
	default:
		log.Printf("Unhandled message type: %s", msg.Type)
	}
}

func (c *Client) sendErrorMessage(clientMsgID string, errorMsg string) {
	errPayload, _ := json.Marshal(payload.FailedToSendPayload{Error: errorMsg})

	errMsg := &payload.WSMessage{
		Type:        payload.WSEventFailedToSend,
		ClientMsgID: clientMsgID,
		Payload:     errPayload,
		Timestamp:   time.Now().UnixMilli(),
	}
	select {
	case c.SendChan <- errMsg:
	default:
		log.Printf("[sendErrorMessage] SendChan full for user %s", c.UserID)
	}
}

// WritePump handles pushing messages to the WebSocket connection
func (c *Client) WritePump() {
	pongWait := time.Duration(config.Cfg.Ws.PongWait) * time.Second
	pingPeriod := (pongWait * 9) / 10
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.closeConnSafely()
	}()

	for {
		select {
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(time.Duration(config.Cfg.Ws.WriteDeadline) * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("[WritePump] error sending ping to user %s: %v", c.UserID, err)
				return
			}
		case msg, ok := <-c.SendChan:
			c.Conn.SetWriteDeadline(time.Now().Add(time.Duration(config.Cfg.Ws.WriteDeadline) * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteJSON(msg); err != nil {
				log.Printf("[WritePump] error writing message to user %s: %v", c.UserID, err)
				return
			}
		}
	}
}

// closeConnSafely closes the underlying WebSocket connection exactly once.
func (c *Client) closeConnSafely() {
	c.closeOnce.Do(func() {
		if c.Conn != nil {
			_ = c.Conn.Close()
		}
	})
}

// CloseSlowConsumer forcibly closes the client connection when send buffer overflows.
// Closing c.Conn causes ReadPump to exit, triggering defer c.Hub.UnregisterClient(c)
// and properly cleaning up presence and channel resources without race conditions.
// It is guarded by sync.Once via closeConnSafely to prevent duplicate goroutines or closing calls.
func (c *Client) CloseSlowConsumer() {
	c.closeConnSafely()
}
````

## File: services/ws-gateway/internal/delivery/grpc_delivery.go
````go
package delivery

import (
	"context"
	"fmt"
	"log"
	"net"

	"chat-system/pkg/contracts"
	pb "chat-system/pkg/proto"
	"chat-system/pkg/telemetry"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/payload"

	"google.golang.org/grpc"
)

type GRPCListener struct {
	pb.UnimplementedWSGatewayServiceServer
	port       int
	hub        *connection.Hub
	grpcServer *grpc.Server
}

func NewGRPCListener(port int, hub *connection.Hub) *GRPCListener {
	return &GRPCListener{
		port: port,
		hub:  hub,
	}
}

func (g *GRPCListener) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return fmt.Errorf("failed to listen on gRPC port %d: %w", g.port, err)
	}
	g.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(telemetry.UnaryServerInterceptor("ws-gateway-grpc", "gateway_delivery")),
	)
	pb.RegisterWSGatewayServiceServer(g.grpcServer, g)
	go func() {
		log.Printf("[Delivery] gRPC Server listening at :%d", g.port)
		if err := g.grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Printf("[Delivery] gRPC server error: %v", err)
		}
	}()
	return nil
}

func (g *GRPCListener) Stop(ctx context.Context) error {
	if g.grpcServer != nil {
		g.grpcServer.GracefulStop()
		log.Println("[Delivery] gRPC Server stopped gracefully")
	}
	return nil
}

func (g *GRPCListener) PushMessageToUser(ctx context.Context, req *pb.PushMessageRequest) (*pb.PushMessageResponse, error) {
	if req.ReceiverId == "" {
		return &pb.PushMessageResponse{Success: false, ErrorMessage: "receiver_id is required"}, nil
	}

	deliveryPayload := payload.MessageDeliveryPayload{
		MessageID:      req.MessageId,
		ClientMsgID:    req.ClientMsgId,
		ConversationID: req.ConversationId,
		SenderID:       req.SenderId,
		ReceiverID:     req.ReceiverId,
		Content:        req.Content,
		Type:           contracts.BrokerMessageType(req.Type),
		Timestamp:      req.Timestamp,
	}

	wsMsg, err := deliveryPayload.NewWSMessageFromDelivery()
	if err != nil {
		return &pb.PushMessageResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	g.hub.SendToUser(req.ReceiverId, wsMsg)

	return &pb.PushMessageResponse{Success: true}, nil
}
````

## File: services/ws-gateway/internal/delivery/nats_delivery.go
````go
package delivery

import (
	"context"
	"fmt"
	"log"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/payload"

	"github.com/nats-io/nats.go"
)

type NATSListener struct {
	nodeID     string
	subscriber *natsclient.Subscriber[contracts.OutboundBrokerEvent]
	hub        *connection.Hub
}

func NewNATSListener(nc *nats.Conn, nodeID string, hub *connection.Hub) *NATSListener {
	subject := contracts.GatewayNodeSubject(nodeID)
	return &NATSListener{
		nodeID:     nodeID,
		subscriber: natsclient.NewSubscriber[contracts.OutboundBrokerEvent](nc, subject),
		hub:        hub,
	}
}

func (n *NATSListener) Start(ctx context.Context) error {
	handler := telemetry.InstrumentHandler(telemetry.HandlerConfig{
		Mode:  "nats",
		Stage: "gateway_delivery",
	}, func(ctx context.Context, event contracts.OutboundBrokerEvent) error {
		deliveryPayload := payload.MessageDeliveryPayload{
			MessageID:      event.MessageID,
			ClientMsgID:    event.ClientMsgID,
			ConversationID: event.ConversationID,
			SenderID:       event.SenderID,
			ReceiverID:     event.ReceiverID,
			Content:        event.Content,
			Type:           event.Type,
			Timestamp:      event.Timestamp,
		}

		wsMsg, err := deliveryPayload.NewWSMessageFromDelivery()
		if err != nil {
			return err
		}

		targetUser := event.ReceiverID
		if targetUser == "" {
			targetUser = event.SenderID
		}
		n.hub.SendToUser(targetUser, wsMsg)
		return nil
	})

	if err := n.subscriber.Start(ctx, handler); err != nil {
		return fmt.Errorf("failed to start NATS listener: %w", err)
	}

	log.Printf("[Delivery] Subscribed to NATS Subject: %s", contracts.GatewayNodeSubject(n.nodeID))
	return nil
}

func (n *NATSListener) Stop(ctx context.Context) error {
	log.Println("[Delivery] NATS listener stopped")
	return nil
}
````

## File: services/ws-gateway/internal/handler/ws_handler.go
````go
package handler

import (
	"errors"
	"net/http"

	"ws-gateway/internal/config"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/payload"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func validateToken(tokenStr string, secret string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
	}

	return "", jwt.ErrInvalidKeyType
}

func HandleWebSocket(hub *connection.Hub, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.URL.Query().Get("token")
		deviceID := r.URL.Query().Get("device_id")

		if deviceID == "" {
			http.Error(w, "missing device id", http.StatusBadRequest)
			return
		}

		if tokenStr == "" {
			http.Error(w, "missing token", http.StatusBadRequest)
			return
		}

		userID, err := validateToken(tokenStr, jwtSecret)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to establish websocket connection", http.StatusInternalServerError)
			return
		}

		sendBufferSize := 32
		if config.Cfg != nil && config.Cfg.Ws.SendBufferSize > 0 {
			sendBufferSize = config.Cfg.Ws.SendBufferSize
		}

		client := &connection.Client{
			UserID:   userID,
			DeviceID: deviceID,
			Hub:      hub,
			SendChan: make(chan *payload.WSMessage, sendBufferSize),
			Conn:     conn,
		}

		hub.RegisterClient(client)

		go client.ReadPump()
		go client.WritePump()
	}
}
````

## File: services/ws-gateway/internal/presence/presence.go
````go
package presence

import (
	"context"
	"fmt"
	"time"
	"ws-gateway/internal/connection"

	"github.com/redis/go-redis/v9"
)

type redisPresenceService struct {
	client *redis.Client
}

func NewPresenceService(redisAddr, password string, db int) connection.PresenceService {
	return &redisPresenceService{
		client: redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: password,
			DB:       db,
		}),
	}
}

func (r *redisPresenceService) SetOnline(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error {
	pipe := r.client.Pipeline()
	key := presenceKey(userID)

	pipe.HSet(ctx, key, deviceID, gatewayNode)
	pipe.Expire(ctx, key, ttl)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *redisPresenceService) SetOffline(ctx context.Context, userID, deviceID string) error {
	return r.client.HDel(ctx, presenceKey(userID), deviceID).Err()
}

func (r *redisPresenceService) Heartbeat(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error {
	return r.SetOnline(ctx, userID, deviceID, gatewayNode, ttl)
}

func (r *redisPresenceService) GetUserRoutes(ctx context.Context, userID string) (map[string]string, error) {
	return r.client.HGetAll(ctx, presenceKey(userID)).Result()
}

func (r *redisPresenceService) IsUserOnline(ctx context.Context, userID string) (bool, error) {
	count, err := r.client.HLen(ctx, presenceKey(userID)).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func presenceKey(userID string) string {
	return fmt.Sprintf("presence:%s", userID)
}
````

## File: Makefile
````makefile
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
	kubectl rollout restart deployment/client-simulator -n chat-system

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

build-api:
	docker build -t chat-system/api-service:latest services/api-service

load-api:
	kind load docker-image chat-system/api-service:latest --name $(CLUSTER)

build-load-api: build-api load-api
	kubectl rollout restart deployment/api-service -n chat-system

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
	kubectl rollout restart deployment/client-simulator -n chat-system

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

# --- Client Simulator (Kubernetes) ---
sim-up:
	kubectl apply -f deployments/k8s/09-client-simulator.yaml

sim-down:
	kubectl delete -f deployments/k8s/09-client-simulator.yaml --ignore-not-found

sim-restart:
	kubectl rollout restart deployment/client-simulator -n chat-system

# Alias giữ tương thích
sim-k8s: sim-up
sim-k8s-down: sim-down

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
````

## File: README.md
````markdown
# Distributed High-Throughput Chat System

A distributed real-time messaging system built with a Microservices architecture, engineered for high throughput, low latency, and horizontal scalability.

---

## 1. System Architecture & Core Services

The system is decoupled into a **Stateful Edge Layer** (connection management) and a **Stateless Processing Layer** (business logic and persistence):

- **`ws-gateway`**: Centralized WebSocket connection gateway managing real-time persistent sessions with clients, enforcing per-connection buffer limits to prevent memory bloat, and dispatching outbound messages to active user devices.
- **`chat-engine`**: Core message processing service supporting dual delivery modes (**gRPC** point-to-point and **NATS** pub/sub broker) and persisting chat history into ScyllaDB.
- **`api-service`**: RESTful API service handling user authentication (JWT), user profile, friendships, and conversation management, built following Clean Architecture principles.
- **`notification-service`**: Asynchronous notification worker handling offline alerts and push notifications for disconnected recipients.
- **`client-simulator`**: High-concurrency bot load generator simulating hundreds of concurrent users to benchmark throughput, stress-test the cluster, and measure edge network latency.

---

## 2. Technology Stack

### Core Language & Runtime
- **Go (Golang 1.22+)**: Primary language across all microservices, leveraging goroutines, channels, and low-latency concurrency primitives for optimal memory and CPU efficiency.

### Communication Protocols
- **WebSocket (Gorilla WebSocket)**: Full-duplex, low-latency bidirectional communication between clients and the WS Gateway.
- **gRPC & Protocol Buffers (Protobuf)**: High-performance binary RPC protocol for inter-service communication and synchronous point-to-point message dispatching.
- **RESTful API**: Standard HTTP interfaces for authentication and metadata management.

### Message Broker & Event Streaming
- **NATS Core / JetStream**: Lightweight, high-performance distributed message broker providing decoupled, asynchronous pub/sub buffering between the Chat Engine and Edge Gateways.

### Databases & Storage
- **ScyllaDB / Apache Cassandra**: Distributed NoSQL database based on LSM-Tree architecture, designed for write-heavy workloads and high-throughput chat message history persistence.
- **PostgreSQL**: Relational database (RDBMS) for structured domain metadata (users, friend graphs, conversation participants).
- **Redis Cluster**: In-memory data store for real-time presence management (online/offline tracking), gateway route registry, and fast caching.

### Observability & Monitoring
- **Prometheus**: Metric collection and time-series aggregation for system health and load metrics.
- **Grafana**: Real-time dashboards monitoring the 5 Golden Dimensions: Latency (P50, P95, P99), Throughput (msgs/sec), Worker Saturation, Error Rates, and Edge Network RTT.
- **OpenTelemetry & Jaeger**: End-to-end distributed tracing across microservice boundaries.

### Infrastructure & Deployment
- **Docker**: Containerization for all microservices and supporting components.
- **Kubernetes (Kind)**: Container orchestration and lifecycle management for local cluster environments.
- **NGINX Ingress Controller**: Edge reverse proxy handling incoming WebSocket and HTTP traffic routing into the cluster.
- **Makefile**: Automation tooling for builds, local image loading, and environment provisioning.
````

## File: deployments/docker-compose.infra.yml
````yaml
services:
  # PostgreSQL - User & Auth Metadata
  postgres:
    image: postgres:16-alpine
    container_name: chat-postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgrespassword
      POSTGRES_DB: chat_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - chat-network

  # Redis - Presence Service & Idempotency Cache
  redis:
    image: redis:7-alpine
    container_name: chat-redis
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --appendonly yes
    networks:
      - chat-network

  # Apache Cassandra - High-throughput Message Store
  cassandra:
    image: cassandra:5.0
    container_name: chat-cassandra
    environment:
      - CASSANDRA_CLUSTER_NAME=chat_cluster
      - CASSANDRA_DC=datacenter1
      - CASSANDRA_RACK=rack1
      - CASSANDRA_ENDPOINT_SNITCH=GossipingPropertyFileSnitch
      - MAX_HEAP_SIZE=1024M
      - HEAP_NEWSIZE=256M
    ports:
      - "9042:9042"
    volumes:
      - cassandra_data:/var/lib/cassandra
    healthcheck:
      test: ["CMD", "cqlsh", "-e", "describe keyspaces"]
      interval: 15s
      timeout: 10s
      retries: 5
      start_period: 30s
    networks:
      - chat-network


  # NATS Server with JetStream enabled
  nats:
    image: nats:latest
    container_name: chat-nats
    ports:
      - "4222:4222" # Client port
      - "8222:8222" # HTTP Monitoring dashboard
    command: ["-js", "-m", "8222"]
    volumes:
      - nats_data:/data
    networks:
      - chat-network

  # ----------------------------------------------------
  # UI & OBSERVABILITY TOOLS (Web-based)
  # ----------------------------------------------------

  # Redis Web UI (Truy cập: http://localhost:5540)
  redis-insight:
    image: redis/redisinsight:latest
    container_name: chat-redis-insight
    ports:
      - "5540:5540"
    volumes:
      - redis_insight_data:/data
    networks:
      - chat-network
    depends_on:
      - redis

  # Cassandra Web UI (Truy cập: http://localhost:8088)
  cassandra-web:
    image: ipushc/cassandra-web:latest
    container_name: chat-cassandra-web
    ports:
      - "8088:8083"
    environment:
      - CASSANDRA_HOST_1=cassandra
      - CASSANDRA_PORT=9042
    networks:
      - chat-network
    depends_on:
      cassandra:
        condition: service_healthy

  # NATS Web Dashboard (Truy cập: http://localhost:8223)
  nats-dashboard:
    image: mdawar/nats-dashboard:latest
    container_name: chat-nats-dashboard
    ports:
      - "8223:80"
    networks:
      - chat-network
    depends_on:
      - nats

  # ----------------------------------------------------
  # OBSERVABILITY STACK (Trực tiếp trên Host, không cần Port-forward)
  # ----------------------------------------------------

  # Prometheus - Time-series Metrics Database (Truy cập: http://localhost:9090)
  prometheus:
    image: prom/prometheus:v2.54.1
    container_name: chat-prometheus
    command:
      - "--config.file=/etc/prometheus/prometheus.yml"
      - "--storage.tsdb.path=/prometheus"
      - "--storage.tsdb.retention.time=7d"
      - "--web.enable-remote-write-receiver"
      - "--enable-feature=otlp-write-receiver"
      - "--web.console.libraries=/etc/prometheus/console_libraries"
      - "--web.console.templates=/etc/prometheus/consoles"
    ports:
      - "9090:9090"
    volumes:
      - ./configs/prometheus.yml:/etc/prometheus/prometheus.yml:ro
      - prometheus_data:/prometheus
    networks:
      - chat-network

  # Pyroscope - Continuous Profiling (Truy cập: http://localhost:4040)
  pyroscope:
    image: grafana/pyroscope:latest
    container_name: chat-pyroscope
    ports:
      - "4040:4040"
    volumes:
      - pyroscope_data:/data
    networks:
      - chat-network

  # Grafana Loki - Log Aggregation System (Truy cập: :3100)
  loki:
    image: grafana/loki:3.1.0
    container_name: chat-loki
    user: "0:0"
    command: ["-config.file=/etc/loki/loki.yaml"]
    ports:
      - "3100:3100"
    volumes:
      - ./configs/loki.yaml:/etc/loki/loki.yaml:ro
      - loki_data:/var/loki
    networks:
      - chat-network

  # Promtail - Log Shipper for Docker Containers (Gắn nhãn container_name)
  promtail:
    image: grafana/promtail:3.1.0
    container_name: chat-promtail
    user: "0:0"
    command: ["-config.file=/etc/promtail/promtail.yaml"]
    volumes:
      - ./configs/promtail.yaml:/etc/promtail/promtail.yaml:ro
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - chat-network
    depends_on:
      - loki

  # OpenTelemetry Collector (Truy cập OTLP: gRPC :4317, HTTP :4318)
  otel-collector:
    image: otel/opentelemetry-collector-contrib:0.110.0
    container_name: chat-otel-collector
    command: ["--config=/etc/otelcol/config.yaml"]
    ports:
      - "4317:4317" # OTLP gRPC
      - "4318:4318" # OTLP HTTP
    volumes:
      - ./configs/otel-collector.yaml:/etc/otelcol/config.yaml:ro
    networks:
      - chat-network
    depends_on:
      - prometheus
      - tempo
      - loki

  # Grafana Tempo - Distributed Tracing Backend (Truy cập: :3200)
  tempo:
    image: grafana/tempo:2.6.0
    container_name: chat-tempo
    command: ["-config.file=/etc/tempo/tempo.yaml"]
    ports:
      - "3200:3200" # HTTP API
    volumes:
      - ./configs/tempo.yaml:/etc/tempo/tempo.yaml:ro
      - tempo_data:/var/tempo
    networks:
      - chat-network

  # Grafana - Unified Observability Dashboard (Truy cập: http://localhost:3000)
  grafana:
    image: grafana/grafana:11.1.0
    container_name: chat-grafana
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_USERS_ALLOW_SIGN_UP=false
      - GF_INSTALL_PLUGINS=grafana-pyroscope-app
    ports:
      - "3000:3000"
    volumes:
      - grafana_data:/var/lib/grafana
      - ./configs/grafana-datasources.yaml:/etc/grafana/provisioning/datasources/datasources.yaml:ro
      - ./configs/grafana/provisioning/dashboards:/etc/grafana/provisioning/dashboards:ro
      - ./configs/grafana/dashboards:/var/lib/grafana/dashboards:ro
    networks:
      - chat-network
    depends_on:
      - prometheus
      - pyroscope
      - tempo
      - loki

networks:
  chat-network:
    name: chat-network
    driver: bridge

volumes:
  postgres_data:
  redis_data:
  cassandra_data:
  nats_data:
  redis_insight_data:
  prometheus_data:
  pyroscope_data:
  grafana_data:
  tempo_data:
  loki_data:
````

## File: services/ws-gateway/cmd/main.go
````go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chat-system/pkg/contracts"
	natsclient "chat-system/pkg/nats"
	"chat-system/pkg/telemetry"
	"ws-gateway/internal/config"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/delivery"
	"ws-gateway/internal/handler"
	"ws-gateway/internal/presence"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting WebSocket Gateway node: %s (mode: %s, ws_port: :%d)",
		cfg.Server.NodeID, cfg.Server.DeliveryMode, cfg.Server.Port)

	// Initialize Unified Telemetry (Tracing + Profiling + Metrics Server)
	shutdownTelemetry, err := telemetry.Setup(context.Background(), telemetry.SetupConfig{
		ServiceName:      "ws-gateway",
		ServiceVersion:   "1.0.0",
		NodeID:           cfg.Server.NodeID,
		CollectorTarget:  cfg.Telemetry.CollectorTarget,
		MetricsPort:      cfg.Telemetry.MetricsPort,
		ProfilerServer:   cfg.Profiler.ServerAddress,
		DisableTracing:   cfg.Telemetry.Disabled,
		DisableProfiling: cfg.Profiler.Disabled,
	})
	if err != nil {
		log.Printf("[Warning] Telemetry setup: %v", err)
	}
	defer shutdownTelemetry(context.Background())

	// Initialize NATS Connection for Inbound events
	nc, err := natsclient.Connect(cfg.NATS.URL, "ws-gateway-"+cfg.Server.NodeID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize Generic NATS Inbound Producer with Telemetry Decorator
	rawProducer := natsclient.NewPublisher[contracts.InboundBrokerEvent](nc, cfg.NATS.InboundSubject)
	inboundProducer := natsclient.NewInstrumentedPublisher(rawProducer, natsclient.PublisherConfig{
		ServiceName: "ws-gateway",
		Stage:       "inbound",
		EventType:   "inbound",
	})

	// Initialize Presence Service
	presenceService := presence.NewPresenceService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)

	// Initialize Connection Hub
	hub := connection.NewHub(inboundProducer, presenceService)
	go hub.Run()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 1. Luôn lắng nghe NATS Subject riêng của Node Gateway (chat.gateway.{node_id}) để nhận Sender ACK
	natsListener := delivery.NewNATSListener(nc, cfg.Server.NodeID, hub)
	if err := natsListener.Start(ctx); err != nil {
		log.Fatalf("Failed to start NATS delivery listener: %v", err)
	}

	// 2. Nếu chạy mode gRPC, khởi động thêm gRPC Server để nhận Outbound Delivery từ chat-engine
	var grpcListener delivery.DeliveryListener
	if cfg.Server.DeliveryMode == "grpc" {
		grpcListener = delivery.NewGRPCListener(cfg.GRPC.Port, hub)
		if err := grpcListener.Start(ctx); err != nil {
			log.Fatalf("Failed to start gRPC delivery listener: %v", err)
		}
	}

	// HTTP / WebSocket route
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handler.HandleWebSocket(hub, cfg.Jwt.AccessTokenSecret))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for termination signal
	<-ctx.Done()

	log.Println("Shutting down WebSocket Gateway...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = natsListener.Stop(shutdownCtx)
	if grpcListener != nil {
		_ = grpcListener.Stop(shutdownCtx)
	}

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}

	log.Println("WebSocket Gateway exited cleanly")
}
````

## File: services/ws-gateway/go.mod
````
module ws-gateway

go 1.25.0

require github.com/gorilla/websocket v1.5.3

require (
	github.com/caarlos0/env/v11 v11.4.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/nats-io/nats.go v1.53.1
	github.com/redis/go-redis/v9 v9.22.0
	go.opentelemetry.io/otel v1.46.0
	go.opentelemetry.io/otel/trace v1.46.0
	google.golang.org/grpc v1.83.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/klauspost/compress v1.19.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/nats-io/nkeys v0.4.15 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.46.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.55.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260819154853-08b0e4226688 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
````

## File: services/ws-gateway/internal/config/config.go
````go
package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Redis     RedisConfig     `yaml:"redis"`
	NATS      NATSConfig      `yaml:"nats"`
	Jwt       JwtConfig       `yaml:"jwt"`
	Ws        WebSocketConfig `yaml:"ws"`
	Pres      PresenceConfig  `yaml:"presence"`
	GRPC      GRPCConfig      `yaml:"grpc"`
	Telemetry TelemetryConfig `yaml:"telemetry"`
	Profiler  ProfilerConfig  `yaml:"profiler"`
}

type TelemetryConfig struct {
	CollectorTarget string `yaml:"collector_target" env:"OTEL_COLLECTOR_TARGET" envDefault:"localhost:4317"`
	MetricsPort     int    `yaml:"metrics_port" env:"METRICS_PORT" envDefault:"9091"`
	Disabled        bool   `yaml:"disabled" env:"OTEL_DISABLED" envDefault:"false"`
}

type ProfilerConfig struct {
	ServerAddress string `yaml:"server_address" env:"PYROSCOPE_SERVER" envDefault:"http://localhost:4040"`
	Disabled      bool   `yaml:"disabled" env:"PYROSCOPE_DISABLED" envDefault:"false"`
}

type ServerConfig struct {
	Port         int    `yaml:"port" env:"WS_PORT" envDefault:"8080"`
	NodeID       string `yaml:"node_id" env:"SERVER_NODE_ID" envDefault:"gateway-node-01"`
	DeliveryMode string `yaml:"delivery_mode" env:"DELIVERY_MODE" envDefault:"grpc"` // "grpc" or "broker"
}

type RedisConfig struct {
	Addr     string `yaml:"addr" env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password string `yaml:"password" env:"REDIS_PASSWORD"`
	DB       int    `yaml:"db" env:"REDIS_DB" envDefault:"0"`
}

type NATSConfig struct {
	URL                   string `yaml:"url" env:"NATS_URL" envDefault:"nats://localhost:4222"`
	InboundSubject        string `yaml:"inbound_subject" env:"NATS_INBOUND_SUBJECT" envDefault:"chat.inbound"`
	OutboundSubjectPrefix string `yaml:"outbound_subject_prefix" env:"NATS_OUTBOUND_SUBJECT_PREFIX" envDefault:"chat.gateway."`
}

type JwtConfig struct {
	AccessTokenSecret string `yaml:"access_token_secret" env:"JWT_ACCESS_TOKEN_SECRET" envDefault:"secret"`
	AccessTokenExpiry int64  `yaml:"access_token_expiry" env:"JWT_ACCESS_TOKEN_EXPIRY" envDefault:"3600"`
}

type WebSocketConfig struct {
	PongWait       int64 `yaml:"pong_wait" env:"WS_PONG_WAIT" envDefault:"60"`
	MaxMessageSize int   `yaml:"max_message_size" env:"WS_MAX_MESSAGE_SIZE" envDefault:"4096"`
	WriteDeadline  int64 `yaml:"write_deadline" env:"WS_WRITE_DEADLINE" envDefault:"10"`
	SendBufferSize int   `yaml:"send_buffer_size" env:"WS_SEND_BUFFER_SIZE" envDefault:"32"`
}

type PresenceConfig struct {
	TTL int64 `yaml:"ttl" env:"PRESENCE_TTL" envDefault:"60"`
}

type GRPCConfig struct {
	Port           int    `yaml:"port" env:"GRPC_PORT" envDefault:"50051"`
	AdvertisedAddr string `yaml:"advertised_addr" env:"GRPC_ADVERTISED_ADDR" envDefault:"localhost:50051"`
}

var Cfg *Config

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	// 1. Đọc file YAML nếu có (làm baseline/default)
	if path != "" {
		data, err := os.ReadFile(path)
		if err == nil {
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("không thể parse file yaml config: %w", err)
			}
		}
	}

	// 2. Tự động đọc và ghi đè từ Environment Variables (Type-Safe qua caarlos0/env)
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("không thể parse biến môi trường: %w", err)
	}

	Cfg = &cfg

	if err := Cfg.Validate(); err != nil {
		return nil, fmt.Errorf("cấu hình không hợp lệ: %w", err)
	}

	return Cfg, nil
}

func (c *Config) Validate() error {
	if c.NATS.URL == "" {
		return fmt.Errorf("nats.url không được để trống")
	}
	if c.Redis.Addr == "" {
		return fmt.Errorf("redis.addr không được để trống")
	}
	if c.Jwt.AccessTokenSecret == "" {
		return fmt.Errorf("jwt.access_token_secret không được để trống")
	}
	return nil
}
````

## File: services/ws-gateway/internal/connection/hub.go
````go
package connection

import (
	"context"
	"log"
	"sync"
	"time"

	"chat-system/pkg/contracts"
	"chat-system/pkg/telemetry"
	"ws-gateway/internal/config"
	"ws-gateway/internal/payload"
)

type PresenceService interface {
	SetOnline(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error
	SetOffline(ctx context.Context, userID, deviceID string) error
	Heartbeat(ctx context.Context, userID, deviceID, gatewayNode string, ttl time.Duration) error
	GetUserRoutes(ctx context.Context, userID string) (map[string]string, error)
	IsUserOnline(ctx context.Context, userID string) (bool, error)
}

// InboundPublisher is satisfied by *nats.Publisher[contracts.InboundBrokerEvent]
type InboundPublisher interface {
	Publish(ctx context.Context, event contracts.InboundBrokerEvent) error
}

// Hub maintains the set of active clients and handles broadcasting
type Hub struct {
	// Registered clients: map[userID]map[deviceID]*Client
	clients    map[string]map[string]*Client
	register   chan *Client
	unregister chan *Client
	producer   InboundPublisher
	presence   PresenceService
	mu         sync.RWMutex
}

func NewHub(producer InboundPublisher, presence PresenceService) *Hub {
	return &Hub{
		clients:    make(map[string]map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		producer:   producer,
		presence:   presence,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; !ok {
				h.clients[client.UserID] = make(map[string]*Client)
			}
			h.clients[client.UserID][client.DeviceID] = client
			h.mu.Unlock()
			telemetry.ActiveConnections.WithLabelValues(config.Cfg.Server.NodeID).Inc()

			go func(c *Client) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := h.presence.SetOnline(ctx, c.UserID, c.DeviceID, config.Cfg.Server.NodeID, time.Duration(config.Cfg.Pres.TTL)*time.Second); err != nil {
					log.Printf("Error setting user online: %v", err)
				}
			}(client)

		case client := <-h.unregister:
			h.mu.Lock()
			if userClients, ok := h.clients[client.UserID]; ok {
				delete(userClients, client.DeviceID)
				if len(userClients) == 0 {
					delete(h.clients, client.UserID)
				}
			}
			close(client.SendChan)
			h.mu.Unlock()
			telemetry.ActiveConnections.WithLabelValues(config.Cfg.Server.NodeID).Dec()

			go func(c *Client) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := h.presence.SetOffline(ctx, c.UserID, c.DeviceID); err != nil {
					log.Printf("Error setting user offline: %v", err)
				}
			}(client)
		}
	}
}

// SendToUser pushes a message to all active devices of a user connected to this node
func (h *Hub) SendToUser(userID string, msg *payload.WSMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if devices, ok := h.clients[userID]; ok {
		for _, client := range devices {
			select {
			case client.SendChan <- msg:
			default:
				// Slow consumer detected: Send buffer is full
				log.Printf("[Hub] Slow consumer detected: SendChan full for user %s, device %s. Dropping message %s and terminating connection",
					client.UserID, client.DeviceID, msg.ClientMsgID)
				telemetry.MessagesProcessed.WithLabelValues("ws-gateway", "outbound", "dropped").Inc()
				go client.CloseSlowConsumer()
			}
		}
	}
}

func (h *Hub) RegisterClient(client *Client) {
	h.register <- client
}

func (h *Hub) UnregisterClient(client *Client) {
	h.unregister <- client
}
````

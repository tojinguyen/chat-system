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

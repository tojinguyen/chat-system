package delivery

import (
	pb "chat-system/pkg/proto"
	"context"
	"fmt"
	"log"
	"net"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/domain"

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
	g.grpcServer = grpc.NewServer()
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

	payload := domain.MessageDeliveryPayload{
		MessageID:      req.MessageId,
		ClientMsgID:    req.ClientMsgId,
		ConversationID: req.ConversationId,
		SenderID:       req.SenderId,
		ReceiverID:     req.ReceiverId,
		Content:        req.Content,
		Type:           req.Type,
		Timestamp:      req.Timestamp,
	}

	wsMsg, err := payload.NewWSMessageFromDelivery()
	if err != nil {
		return &pb.PushMessageResponse{Success: false, ErrorMessage: err.Error()}, nil
	}
	g.hub.SendToUser(req.ReceiverId, wsMsg)

	return &pb.PushMessageResponse{Success: true}, nil
}

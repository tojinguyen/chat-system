package handler

import (
	pb "chat-system/pkg/proto"
	"context"
	"encoding/json"
	"fmt"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/domain"
)

type WSGatewayGRPCServer struct {
	pb.UnimplementedWSGatewayServiceServer
	hub *connection.Hub
}

func NewGatewayGRPCServer(hub *connection.Hub) *WSGatewayGRPCServer {
	return &WSGatewayGRPCServer{
		hub: hub,
	}
}

func (s *WSGatewayGRPCServer) PushMessageToUser(ctx context.Context, req *pb.PushMessageRequest) (*pb.PushMessageResponse, error) {
	if req.ReceiverId == "" {
		return &pb.PushMessageResponse{
			Success:      false,
			ErrorMessage: "receiver_id is required",
		}, fmt.Errorf("receiver_id is required")
	}

	payloadBytes, _ := json.Marshal(map[string]interface{}{
		"message_id":      req.MessageId,
		"client_msg_id":   req.ClientMsgId,
		"conversation_id": req.ConversationId,
		"sender_id":       req.SenderId,
		"receiver_id":     req.ReceiverId,
		"content":         req.Content,
		"type":            req.Type,
		"timestamp":       req.Timestamp,
	})

	wsMsg := &domain.WSMessage{
		Type:      domain.WSEventSendMessage,
		Timestamp: req.Timestamp,
		Payload:   payloadBytes,
	}

	s.hub.SendToUser(req.ReceiverId, wsMsg)
	return &pb.PushMessageResponse{
		Success: true,
	}, nil
}

package server

import (
	"context"

	messageService "github.com/levYyyyy/message-microservice/internal/message/proto"
)

type MessageServer struct {
	messageService.UnimplementedMessageServiceServer
}

func (s *MessageServer) SendMessage(ctx context.Context, req *messageService.MessageRequest) (*messageService.MessageResponse, error) {
	return &messageService.MessageResponse{Body: req.GetBody(), From: req.GetTo()}, nil
}

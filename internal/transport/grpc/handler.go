package grpc

import (
	"context"
	"log"
	"net"

	api "github.com/bobhonores/planner/api/v1"
	"github.com/bobhonores/planner/internal/action"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ActionService interface {
	GetByID(context.Context, uuid.UUID) (action.Action, error)
	Insert(context.Context, action.Action) (action.Action, error)
	Update(context.Context, action.Action) (action.Action, error)
	Delete(context.Context, uuid.UUID) error
}

type Handler struct {
	ActionService ActionService
	api.UnimplementedActionServiceServer
}

func New(actService ActionService) Handler {
	return Handler{
		ActionService: actService,
	}
}

func (h *Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	api.RegisterActionServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

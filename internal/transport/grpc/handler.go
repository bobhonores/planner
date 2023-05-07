package grpc

import (
	"log"
	"net"

	api "github.com/bobhonores/planner/api/v1"
	"github.com/bobhonores/planner/internal/action"
	"google.golang.org/grpc"
)

type Handler struct {
	ActionService                        action.ActionService
	api.UnimplementedActionServiceServer // This is required for gRPC service
}

func New(actService action.ActionService) Handler {
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

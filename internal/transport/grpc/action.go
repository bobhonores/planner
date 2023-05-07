package grpc

import (
	"context"

	api "github.com/bobhonores/planner/api/v1"
)

func (h *Handler) GetAction(ctx context.Context, req *api.GetActionRequest) (*api.GetActionResponse, error) {
	// actionID, err := uuid
	// h.ActionService.GetByID(ctx, )
	return &api.GetActionResponse{}, nil
}

func (h *Handler) AddAction(ctx context.Context, req *api.AddActionRequest) (*api.AddActionResponse, error) {
	return &api.AddActionResponse{}, nil
}

func (h *Handler) UpdateAction(ctx context.Context, req *api.UpdateActionRequest) (*api.UpdateActionResponse, error) {
	return &api.UpdateActionResponse{}, nil
}

func (h *Handler) DeleteAction(ctx context.Context, req *api.DeleteActionRequest) (*api.DeleteActionResponse, error) {
	return &api.DeleteActionResponse{}, nil
}

package grpc

import (
	"context"
	"errors"
	"log"

	api "github.com/bobhonores/planner/api/v1"
	"github.com/bobhonores/planner/internal/action"
	"github.com/google/uuid"
)

func (h *Handler) GetAction(ctx context.Context, req *api.GetActionRequest) (*api.GetActionResponse, error) {
	actionID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.New("incorrect id format")
	}
	action, err := h.ActionService.GetByID(ctx, actionID)
	if err != nil {
		return nil, errors.New("action not found")
	}
	return &api.GetActionResponse{
		Action: &api.Action{
			Id:          action.ID.String(),
			Name:        action.Name,
			Description: action.Description,
			Done:        action.Done,
		},
	}, nil
}

func (h *Handler) AddAction(ctx context.Context, req *api.AddActionRequest) (*api.AddActionResponse, error) {
	action, err := h.ActionService.Insert(ctx, action.Action{
		Name:        req.Name,
		Description: req.Description,
		Done:        req.Done,
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("cannot insert the action")
	}

	return &api.AddActionResponse{
		Action: &api.Action{
			Id:          action.ID.String(),
			Name:        action.Name,
			Description: action.Description,
			Done:        action.Done,
		},
	}, nil
}

func (h *Handler) UpdateAction(ctx context.Context, req *api.UpdateActionRequest) (*api.UpdateActionResponse, error) {
	return &api.UpdateActionResponse{}, nil
}

func (h *Handler) DeleteAction(ctx context.Context, req *api.DeleteActionRequest) (*api.DeleteActionResponse, error) {
	return &api.DeleteActionResponse{}, nil
}

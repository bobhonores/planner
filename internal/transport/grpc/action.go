package grpc

import (
	"context"
	"fmt"
	"log"

	api "github.com/bobhonores/planner/api/v1"
	"github.com/bobhonores/planner/internal/action"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetAction(ctx context.Context, req *api.GetActionRequest) (*api.GetActionResponse, error) {
	actionID, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Print(err)
		errorStatus := status.Error(codes.InvalidArgument, "incorrect id format")
		return &api.GetActionResponse{}, errorStatus
	}

	action, err := h.ActionService.GetByID(ctx, actionID)
	if err != nil {
		log.Print(err)
		errorStatus := status.Error(codes.NotFound, "action not found")
		return &api.GetActionResponse{}, errorStatus
	}

	response := api.GetActionResponse{
		Action: &api.Action{
			Id:            action.ID.String(),
			Name:          action.Name,
			Description:   action.Description,
			Done:          action.Done,
			LastOperation: timestamppb.New(action.Created),
		}}
	if !action.Updated.IsZero() {
		response.Action.LastOperation = timestamppb.New(action.Updated)
	}

	return &response, nil
}

func (h *Handler) AddAction(ctx context.Context, req *api.AddActionRequest) (*api.AddActionResponse, error) {
	action, err := h.ActionService.Insert(ctx, action.Action{
		Name:        req.Name,
		Description: req.Description,
		Done:        req.Done,
	})
	if err != nil {
		log.Println(err)
		errorStatus := status.Error(codes.Internal, "cannot insert the action")
		return &api.AddActionResponse{}, errorStatus
	}

	return &api.AddActionResponse{
		Action: &api.Action{
			Id:            action.ID.String(),
			Name:          action.Name,
			Description:   action.Description,
			Done:          action.Done,
			LastOperation: timestamppb.New(action.Created),
		},
	}, nil
}

func (h *Handler) UpdateAction(ctx context.Context, req *api.UpdateActionRequest) (*api.UpdateActionResponse, error) {
	actionID, err := uuid.Parse(req.Id)
	if err != nil {
		log.Println(err)
		errorStatus := status.Error(codes.InvalidArgument, "incorrect id format")
		return &api.UpdateActionResponse{}, errorStatus
	}

	action, err := h.ActionService.Update(ctx, action.Action{
		ID:          actionID,
		Name:        req.Name,
		Description: req.Description,
		Done:        req.Done,
	})
	if err != nil {
		log.Println(err)
		errorStatus := status.Error(codes.Internal, "cannot update the action")
		return &api.UpdateActionResponse{}, errorStatus
	}

	return &api.UpdateActionResponse{
		Action: &api.Action{
			Id:            action.ID.String(),
			Name:          action.Name,
			Description:   action.Description,
			Done:          action.Done,
			LastOperation: timestamppb.New(action.Updated),
		},
	}, nil
}

func (h *Handler) DeleteAction(ctx context.Context, req *api.DeleteActionRequest) (*api.DeleteActionResponse, error) {
	actionID, err := uuid.Parse(req.Id)
	if err != nil {
		log.Println(err)
		errorStatus := status.Error(codes.InvalidArgument, "incorrect id format")
		return &api.DeleteActionResponse{}, errorStatus
	}

	if err := h.ActionService.Delete(ctx, actionID); err != nil {
		log.Println(err)
		errorStatus := status.Error(codes.Internal, "cannot delete the action")
		return &api.DeleteActionResponse{}, errorStatus
	}

	return &api.DeleteActionResponse{
		Status: fmt.Sprintf("%v was deleted", req.Id),
	}, nil
}

//go:generate mockgen -destination=action_mocks_test.go -package=action github.com/bobhonores/planner/internal/action Repository
package action

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Action represents a potential task to execute
type Action struct {
	ID          uuid.UUID
	Name        string
	Description string
	Done        bool
	Created     time.Time
	Updated     time.Time
}

type ActionService interface {
	GetByID(context.Context, uuid.UUID) (Action, error)
	Insert(context.Context, Action) (Action, error)
	Update(context.Context, Action) (Action, error)
	Delete(context.Context, uuid.UUID) error
}

// Repository defines the interface to be used in the service
type Repository interface {
	Get(context.Context, string) (Action, error)
	Insert(context.Context, Action) (Action, error)
	Delete(context.Context, string) error
	Update(context.Context, Action) (Action, error)
}

// Service to interact with all actions
type Service struct {
	Repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (Action, error) {
	action, err := s.Repository.Get(ctx, id.String())
	if err != nil {
		return Action{}, err
	}

	return action, nil
}

func (s *Service) Insert(
	ctx context.Context,
	action Action,
) (Action, error) {
	action.ID = uuid.New()
	action.Created = time.Now().UTC()
	insertedAction, err := s.Repository.Insert(ctx, action)
	if err != nil {
		return Action{}, err
	}

	return insertedAction, nil
}

func (s *Service) Update(
	ctx context.Context,
	action Action,
) (Action, error) {
	action.Updated = time.Now().UTC()
	updatedAction, err := s.Repository.Update(ctx, action)
	if err != nil {
		return Action{}, err
	}

	return updatedAction, nil
}

func (s *Service) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return s.Repository.Delete(ctx, id.String())
}

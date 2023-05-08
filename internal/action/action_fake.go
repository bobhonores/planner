package action

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type FakeActionRepository struct {
}

func NewFakeActionRepository() *FakeActionRepository {
	return &FakeActionRepository{}
}

func (f *FakeActionRepository) Get(ctx context.Context, id string) (Action, error) {
	actionID, err := uuid.Parse(id)
	if err != nil {
		return Action{}, err
	}

	return Action{
		ID:          actionID,
		Name:        gofakeit.Sentence(4),
		Description: gofakeit.Sentence(10),
		Done:        false,
		Created:     time.Now().UTC(),
	}, nil
}

func (f *FakeActionRepository) Insert(ctx context.Context, act Action) (Action, error) {
	return Action{
		ID:          uuid.New(),
		Name:        act.Name,
		Description: act.Description,
		Done:        act.Done,
		Created:     time.Now().UTC(),
	}, nil
}

func (f *FakeActionRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (f *FakeActionRepository) Update(ctx context.Context, act Action) (Action, error) {
	return Action{
		ID:          act.ID,
		Name:        act.Name,
		Description: act.Description,
		Done:        act.Done,
		Created:     time.Now().UTC(),
		Updated:     time.Now().UTC(),
	}, nil
}

package action

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestActionService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test get action by id", func(t *testing.T) {
		actionRepositoryMock := NewMockRepository(mockCtrl)
		id := "c5392c8b-6a11-4868-9582-2d8d71118f50"
		actionId, _ := uuid.Parse(id)
		actionRepositoryMock.
			EXPECT().
			Get(context.Background(), id).
			Return(Action{
				ID: actionId,
			}, nil)

		actionService := New(actionRepositoryMock)
		action, err := actionService.GetByID(
			context.Background(),
			actionId,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, action.ID.String())
	})

	t.Run("test insert action", func(t *testing.T) {
		actionRepositoryMock := NewMockRepository(mockCtrl)
		id := "c5392c8b-6a11-4868-9582-2d8d71118f50"
		actionId, _ := uuid.Parse(id)
		actionRepositoryMock.
			EXPECT().
			Insert(context.Background(), Action{ID: actionId}).
			Return(Action{
				ID: actionId,
			}, nil)

		actionService := New(actionRepositoryMock)
		insertedAction, err := actionService.Insert(
			context.Background(),
			Action{ID: actionId},
		)
		assert.NoError(t, err)
		assert.Equal(t, id, insertedAction.ID.String())
	})

	t.Run("test update action", func(t *testing.T) {
		actionRepositoryMock := NewMockRepository(mockCtrl)
		id := "c5392c8b-6a11-4868-9582-2d8d71118f50"
		name := "test"
		actionId, _ := uuid.Parse(id)
		actionRepositoryMock.
			EXPECT().
			Update(context.Background(),
				Action{
					ID:   actionId,
					Name: name,
				}).
			Return(Action{
				ID:   actionId,
				Name: name,
			}, nil)

		actionService := New(actionRepositoryMock)
		updatedAction, err := actionService.Update(
			context.Background(),
			Action{
				ID:   actionId,
				Name: name,
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, id, updatedAction.ID.String())
		assert.Equal(t, name, updatedAction.Name)
	})

	t.Run("test delete action by id", func(t *testing.T) {
		actionRepositoryMock := NewMockRepository(mockCtrl)
		id := "c5392c8b-6a11-4868-9582-2d8d71118f50"
		actionId, _ := uuid.Parse(id)
		actionRepositoryMock.
			EXPECT().
			Delete(context.Background(), id).
			Return(nil)

		actionService := New(actionRepositoryMock)
		err := actionService.Delete(
			context.Background(),
			actionId,
		)
		assert.NoError(t, err)
	})
}

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
		fakeRepository := NewFakeActionRepository()

		inputCtx := context.Background()
		inputAction := Action{
			Name:        "todo",
			Description: "demo description",
		}

		actionService := New(fakeRepository)
		insertedAction, err := actionService.Insert(
			inputCtx,
			inputAction,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, insertedAction.ID.String())
		assert.Equal(t, "todo", insertedAction.Name)
		assert.Equal(t, "demo description", insertedAction.Description)
	})

	t.Run("test update action", func(t *testing.T) {
		fakeRepository := NewFakeActionRepository()

		id := "c5392c8b-6a11-4868-9582-2d8d71118f50"
		actionId, _ := uuid.Parse(id)

		inputCtx := context.Background()
		inputAction := Action{
			ID:          actionId,
			Name:        "todo",
			Description: "demo description",
		}

		actionService := New(fakeRepository)
		updatedAction, err := actionService.Update(
			inputCtx,
			inputAction,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, updatedAction.ID.String())
		assert.Equal(t, "todo", updatedAction.Name)
		assert.Equal(t, "demo description", updatedAction.Description)
		assert.Condition(t, func() bool { return !updatedAction.Updated.IsZero() })
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

//go:build integration

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bobhonores/planner/internal/comment"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		_, err = db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "updated-slug",
			Author: "updated-author",
			Body:   "updated-body",
		})
		assert.NoError(t, err)

		updatedCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "updated-slug", updatedCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "new-author",
			Body:   "new-body",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}

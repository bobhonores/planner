package comment

import (
	"context"
	"errors"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - representation of the comment structure for the service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - define repository interface to be used in the service
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service - entry point for all logic for comments
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service structure
func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(
	ctx context.Context,
	ID string,
	updatedCmt Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, insertedCmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, insertedCmt)
	if err != nil {
		return Comment{}, err
	}

	return insertedCmt, nil
}

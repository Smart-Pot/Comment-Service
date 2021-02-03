package service

import (
	"commentservice/data"
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
}

type Service interface {
	GetByUser(ctx context.Context, userID string) ([]*data.Comment, error)
	GetByPost(ctx context.Context, postID string) ([]*data.Comment, error)
	Add(ctx context.Context, newComment data.Comment) error
	Delete(ctx context.Context, commentID string) error
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) GetByUser(ctx context.Context, userID string) ([]*data.Comment, error) {
	return data.GetCommentsByUserID(ctx, userID)
}

func (s service) GetByPost(ctx context.Context, postID string) ([]*data.Comment, error) {
	return data.GetCommentsByPostID(ctx, postID)
}

func (s service) Add(ctx context.Context, newComment data.Comment) error {
	return data.AddComment(ctx, newComment)
}

func (s service) Delete(ctx context.Context, commentID string) error {
	return data.DeleteComment(ctx, commentID)
}

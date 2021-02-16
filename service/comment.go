package service

import (
	"commentservice/data"
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger log.Logger
}

type Service interface {
	GetByUser(ctx context.Context, userID string, pageNumber, pageSize int) ([]*data.Comment, error)
	GetByPost(ctx context.Context, postID string, pageNumber, pageSize int) ([]*data.Comment, error)
	Add(ctx context.Context, userID string, newComment data.Comment) error
	Delete(ctx context.Context, userID, commentID string) error
	DeleteUsersComments(ctx context.Context, userID string) error
	Vote(ctx context.Context, userID, commentID string) error
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Vote(ctx context.Context, userID, commentID string) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Vote",
			"param:userID", userID,
			"param:commentID", commentID,
			"took", time.Since(beginTime))
	}(time.Now())
	return data.Vote(ctx, userID, commentID)
}

func (s service) GetByUser(ctx context.Context, userID string, pageNumber, pageSize int) (result []*data.Comment, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetByUser",
			"param:userID", userID,
			"param:pageNumber", pageNumber,
			"param:pageSize", pageSize,
			"result", result,
			"took", time.Since(beginTime))
	}(time.Now())
	result, err = data.GetCommentsByUserID(ctx, userID, pageNumber, pageSize)
	return result, err
}

func (s service) GetByPost(ctx context.Context, postID string, pageNumber, pageSize int) (result []*data.Comment, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetByPost",
			"param:postID", postID,
			"param:pageNumber", pageNumber,
			"param:pageSize", pageSize,
			"result", result,
			"took", time.Since(beginTime))
	}(time.Now())
	return data.GetCommentsByPostID(ctx, postID, pageNumber, pageSize)
}

func (s service) Add(ctx context.Context, userID string, newComment data.Comment) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Add",
			"param:newComment", newComment,
			"param:userID", userID,
			"took", time.Since(beginTime))
	}(time.Now())
	// Validate comment
	if err := newComment.Validate(); err != nil {
		return err
	}
	if newComment.UserID != userID {
		return errors.New("User can not create comments for other users")
	}
	return data.AddComment(ctx, newComment)
}

func (s service) Delete(ctx context.Context, userID, commentID string) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Delete",
			"param:commentID", commentID,
			"param:userID", userID,
			"took", time.Since(beginTime))
	}(time.Now())
	cmt, err := data.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}
	if userID != cmt.UserID {
		return errors.New("User can not delete comments of other users")
	}
	return data.DeleteComment(ctx, commentID)
}

func (s service) DeleteUsersComments(ctx context.Context, userID string) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "DeleteMany",
			"param:userID", userID,
			"took", time.Since(beginTime))
	}(time.Now())
	var err error

	for x := 0; x < 3; x++ {
		err = data.DeleteUsersComments(ctx, userID)
		if err == nil {
			break
		}
	}

	return err
}

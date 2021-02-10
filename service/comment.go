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
	GetByUser(ctx context.Context, userID string) ([]*data.Comment, error)
	GetByPost(ctx context.Context, postID string) ([]*data.Comment, error)
	Add(ctx context.Context, userID string, newComment data.Comment) error
	Delete(ctx context.Context, userID, commentID string) error
	DeleteMany(ctx context.Context, postID string) error
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

func (s service) GetByUser(ctx context.Context, userID string) (result []*data.Comment, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetByUser",
			"param:userID", userID,
			"result", result,
			"took", time.Since(beginTime))
	}(time.Now())
	result, err = data.GetCommentsByUserID(ctx, userID)
	return result, err
}

func (s service) GetByPost(ctx context.Context, postID string) (result []*data.Comment, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetByPost",
			"param:postID", postID,
			"result", result,
			"took", time.Since(beginTime))
	}(time.Now())
	return data.GetCommentsByPostID(ctx, postID)
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

// TODO: User will able to delete only his/her own comments
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

func (s service) DeleteMany(ctx context.Context, postID string) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "DeleteMany",
			"param:postID", postID,
			"took", time.Since(beginTime))
	}(time.Now())
	var err error

	for x := 0; x < 3; x++ {
		err = data.DeletePostsComments(ctx, postID)
		if err == nil {
			break
		}
	}

	return err
}

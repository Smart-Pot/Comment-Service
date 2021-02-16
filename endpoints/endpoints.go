package endpoints

import (
	"commentservice/data"
	"commentservice/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetByUser endpoint.Endpoint
	GetByPost endpoint.Endpoint
	Add       endpoint.Endpoint
	Delete    endpoint.Endpoint
	Vote      endpoint.Endpoint
}

type CommentResponse struct {
	Comments []*data.Comment
	Success  int32
	Message  string
}

type VoteRequest struct {
	UserID    string `json:"userId"`
	CommentID string `json:"commentId"`
}

type CommentsRequest struct {
	ID         string
	PageNumber int
	PageSize   int
}

type CommentRequest struct {
	ID     string
	UserID string
}

type NewCommentRequest struct {
	NewComment data.Comment
	UserID     string
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetByUser: makeGetByUserEndpoint(s),
		GetByPost: makeGetByPostEndpoint(s),
		Add:       makeAddEndpoint(s),
		Delete:    makeDeleteEndpoint(s),
		Vote:      makeVoteEndpoint(s),
	}
}

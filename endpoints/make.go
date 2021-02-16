package endpoints

import (
	"commentservice/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetByUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CommentsRequest)
		result, err := s.GetByUser(ctx, req.ID, req.PageNumber, req.PageSize)
		response := CommentResponse{Comments: result, Success: 1, Message: "Comments found!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeGetByPostEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CommentsRequest)
		result, err := s.GetByPost(ctx, req.ID, req.PageNumber, req.PageSize)
		response := CommentResponse{Comments: result, Success: 1, Message: "Comments found!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewCommentRequest)
		err := s.Add(ctx, req.UserID, req.NewComment)
		response := CommentResponse{Comments: nil, Success: 1, Message: "Comment created!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CommentRequest)
		err := s.Delete(ctx, req.UserID, req.ID)
		response := CommentResponse{Comments: nil, Success: 1, Message: "Comment deleted!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

func makeVoteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VoteRequest)
		err := s.Vote(ctx, req.UserID, req.CommentID)
		response := CommentResponse{Comments: nil, Success: 1, Message: "Vote successful!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}

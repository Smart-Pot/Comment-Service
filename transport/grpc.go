package transport

import (
	"context"

	"commentservice/data"
	"commentservice/endpoints"
	"commentservice/pb"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedCommentServiceServer
	getByPost gt.Handler
	getByUser gt.Handler
	add       gt.Handler
	delete    gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.CommentServiceServer {
	return &gRPCServer{
		getByPost: gt.NewServer(
			endpoints.GetByPost,
			decodeCommentRequest,
			encodeCommentResponse,
		),
		getByUser: gt.NewServer(
			endpoints.GetByUser,
			decodeCommentRequest,
			encodeCommentResponse,
		),
		add: gt.NewServer(
			endpoints.Add,
			decodeNewCommentRequest,
			encodeCommentResponse,
		),
		delete: gt.NewServer(
			endpoints.Delete,
			decodeCommentRequest,
			encodeCommentResponse,
		),
	}
}

func (s *gRPCServer) GetByPost(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	_, resp, err := s.getByPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CommentResponse), nil
}

func (s *gRPCServer) GetByUser(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	_, resp, err := s.getByUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CommentResponse), nil
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.NewCommentRequest) (*pb.CommentResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CommentResponse), nil
}

func (s *gRPCServer) Delete(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	_, resp, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CommentResponse), nil
}

func decodeCommentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CommentRequest)
	return endpoints.CommentRequest{ID: req.ID}, nil
}

func encodeCommentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CommentResponse)
	comments := make([]*pb.Comment, len(res.Comments))

	for _, c := range res.Comments {
		n := &pb.Comment{
			ID:      c.ID,
			PostID:  c.PostID,
			UserID:  c.UserID,
			Content: c.Content,
			Like:    c.Like,
			Dislike: c.Dislike,
			Date:    c.Date,
		}
		comments = append(comments, n)
	}
	return &pb.CommentResponse{Comments: comments, Success: res.Success, Message: res.Message}, nil
}

func decodeNewCommentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.NewCommentRequest)
	return endpoints.NewCommentRequest{NewComment: data.Comment{
		ID:      req.NewComment.ID,
		PostID:  req.NewComment.PostID,
		UserID:  req.NewComment.UserID,
		Content: req.NewComment.Content,
		Like:    req.NewComment.Like,
		Dislike: req.NewComment.Dislike,
		Date:    req.NewComment.Date,
	}}, nil
}

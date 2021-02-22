package transport

import (
	"commentservice/data"
	"commentservice/endpoints"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrNotInteger    = errors.New("pagesize and pagenumber must be integer")
	ErrWrongArgument = errors.New("missing or wrong argument in request")
)

const userIDTag = "x-user-id"

func MakeHTTPHandlers(e endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter().PathPrefix("/comment").Subrouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/user/{id}/{pagenumber}/{pagesize}").Handler(httptransport.NewServer(
		e.GetByUser,
		decodeCommentsHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("GET").Path("/post/{id}/{pagenumber}/{pagesize}").Handler(httptransport.NewServer(
		e.GetByPost,
		decodeCommentsHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("DELETE").Path("/{id}").Handler(httptransport.NewServer(
		e.Delete,
		decodeCommentHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("POST").Path("/new").Handler(httptransport.NewServer(
		e.Add,
		decodeNewCommentHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("POST").Path("/vote").Handler(httptransport.NewServer(
		e.Vote,
		decodeVoteHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	return r
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeCommentHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, idOK := vars["id"]

	if !idOK {
		return nil, ErrWrongArgument
	}
	return endpoints.CommentRequest{
		ID:     id,
		UserID: r.Header.Get(userIDTag),
	}, nil
}

func decodeCommentsHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	id, idOK := vars["id"]
	pn, pnOK := vars["pagenumber"]
	ps, psOK := vars["pagesize"]

	pagenumber, err := strconv.Atoi(pn)
	if err != nil {
		return nil, ErrNotInteger
	}
	pagesize, err := strconv.Atoi(ps)

	if err != nil {
		return nil, ErrNotInteger
	}

	if !idOK || !pnOK || !psOK {
		return nil, ErrWrongArgument
	}
	return endpoints.CommentsRequest{
		ID:         id,
		PageNumber: pagenumber,
		PageSize:   pagesize,
	}, nil
}

func decodeVoteHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.VoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.UserID = r.Header.Get(userIDTag)

	return req, nil
}

func decodeNewCommentHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var c data.Comment

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return nil, err
	}
	return endpoints.NewCommentRequest{
		NewComment: c,
		UserID:     r.Header.Get(userIDTag),
	}, nil

}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

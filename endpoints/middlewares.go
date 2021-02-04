package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

var validationMiddleware endpoint.Middleware = func(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		return next(ctx, request)
	}
}

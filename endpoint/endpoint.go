package endpoint

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Endpoint[Req, Res any] func(ctx context.Context, request Req) (Res, error)

type Middleware[Req, Res any] func(Endpoint[Req, Res]) Endpoint[Req, Res]

func Chain[Req, Res any](outer Middleware[Req, Res], others ...Middleware[Req, Res]) Middleware[Req, Res] {
	return func(next Endpoint[Req, Res]) Endpoint[Req, Res] {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}

var ErrTypeCast = errors.New("types do not match")

func (e Endpoint[Req, Res]) Unwrap() endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (response interface{}, err error) {
		rReq, ok := r.(Req)
		if !ok {
			return nil, ErrTypeCast
		}
		return e(ctx, rReq)
	}
}

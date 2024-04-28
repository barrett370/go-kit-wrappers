package http

import (
	"context"
	"errors"
	"net/http"

	kit "github.com/go-kit/kit/transport/http"
)

var ErrTypeCast = errors.New("types do not match")

type DecodeRequestFunc[Req any] func(context.Context, *http.Request) (Req, error)

type EncodeRequestFunc[Req any] func(context.Context, *http.Request, Req) error

type EncodeResponseFunc[Res any] func(context.Context, http.ResponseWriter, Res) error

type DecodeResponseFunc[Res any] func(context.Context, *http.Response) (Res, error)

func (d DecodeRequestFunc[Req]) Unwrap() kit.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		return d(ctx, r)
	}
}

func (e EncodeRequestFunc[Req]) Unwrap() kit.EncodeRequestFunc {
	return func(ctx context.Context, r *http.Request, i interface{}) error {
		iReq, ok := i.(Req)
		if !ok {
			return ErrTypeCast
		}
		return e(ctx, r, iReq)
	}
}

func (e EncodeResponseFunc[Res]) Unwrap() kit.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, i interface{}) error {
		iRes, ok := i.(Res)
		if !ok {
			return ErrTypeCast
		}
		return e(ctx, w, iRes)
	}
}

func (d DecodeResponseFunc[Res]) Unwrap() kit.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (response interface{}, err error) {
		return d(ctx, r)
	}
}

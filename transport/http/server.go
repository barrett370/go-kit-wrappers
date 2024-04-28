package http

import (
	"github.com/barrett370/go-kit-wrappers/endpoint"
	kit "github.com/go-kit/kit/transport/http"
)

func NewServer[Req, Res any](
	e endpoint.Endpoint[Req, Res],
	dec DecodeRequestFunc[Req],
	enc EncodeResponseFunc[Res],
	opts ...kit.ServerOption) *kit.Server {
	return kit.NewServer(
		e.Unwrap(),
		dec.Unwrap(),
		enc.Unwrap(),
		opts...,
	)
}

package http

import (
	"net/url"

	kit "github.com/go-kit/kit/transport/http"
)

func NewClient[Req, Res any](
	method string,
	tgt *url.URL,
	enc EncodeRequestFunc[Req],
	dec DecodeResponseFunc[Res],
	opts ...kit.ClientOption) *kit.Client {
	return kit.NewClient(
		method,
		tgt,
		enc.Unwrap(),
		dec.Unwrap(),
		opts...,
	)
}

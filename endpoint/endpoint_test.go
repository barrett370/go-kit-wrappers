package endpoint_test

import (
	"context"
	"fmt"

	"github.com/barrett370/go-kit-wrappers/endpoint"
)

func ExampleChain() {
	e := endpoint.Chain(
		annotate("first"),
		annotate("second"),
		annotate("third"),
	)(myEndpoint)

	if _, err := e(ctx, req); err != nil {
		panic(err)
	}

	// Output:
	// first pre
	// second pre
	// third pre
	// my endpoint!
	// third post
	// second post
	// first post
}

var (
	ctx = context.Background()
	req = struct{}{}
)

func annotate(s string) endpoint.Middleware[myRequest, myResponse] {
	return func(next endpoint.Endpoint[myRequest, myResponse]) endpoint.Endpoint[myRequest, myResponse] {
		return func(ctx context.Context, request myRequest) (myResponse, error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			return next(ctx, request)
		}
	}
}

type myRequest struct{}

type myResponse struct{}

func myEndpoint(context.Context, myRequest) (myResponse, error) {
	fmt.Println("my endpoint!")
	return myResponse{}, nil
}

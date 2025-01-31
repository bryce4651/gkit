package middleware

import (
	"context"
	"net/http"
)

var HTTPMiddlewareList []HttpMiddleware

func RegisterHTTPMiddleware(middleware HttpMiddleware) {
	HTTPMiddlewareList = append(HTTPMiddlewareList, middleware)
}

type HttpHandler func(writer http.ResponseWriter, request *http.Request)

type HttpMiddleware func(http.Handler) http.Handler

func Chain(m ...HttpMiddleware) HttpMiddleware {
	return func(next http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- { // reverse
			next = m[i](next)
		}
		return next
	}
}

type BeforeHandler func(ctx context.Context, req interface{}) (context.Context, interface{}, error)
type AfterHandler func(response interface{}, err error) (interface{}, error)

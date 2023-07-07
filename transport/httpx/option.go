package httpx

import (
	"github.com/ml444/gkit/auth/jwt"
	"github.com/ml444/gkit/middleware"
	"time"
)

type OptionFunc func(parser *EndpointParser)

func SetTimeoutMap(timeoutMap map[string]time.Duration) OptionFunc {
	return func(parser *EndpointParser) {
		parser.timeoutMap = timeoutMap
	}
}

func SetJwtHook(hook jwt.HookFunc) OptionFunc {
	return func(parser *EndpointParser) {
		parser.jwtHook = hook
	}
}

func AddBeforeHandler(handlers ...middleware.BeforeHandler) OptionFunc {
	return func(parser *EndpointParser) {
		parser.beforeHandlerList = append(parser.beforeHandlerList, handlers...)
	}
}

func AddAfterHandler(handlers ...middleware.AfterHandler) OptionFunc {
	return func(parser *EndpointParser) {
		parser.afterHandlerList = append(parser.afterHandlerList, handlers...)
	}
}
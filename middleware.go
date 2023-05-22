package goweb

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Chain []Middleware

func chainBuilder(middlewares ...Middleware) Chain {
	var chain Chain
	return append(chain, middlewares...)
}

func (c Chain) build(handler http.HandlerFunc) http.HandlerFunc {
	if len(c) < 1 {
		return handler
	}
	for i := len(c) - 1; i >= 0; i-- {
		handler = c[i](handler)
	}
	return handler
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Debug().Str("remote_addr", req.RemoteAddr).
			Str("method", req.Method).
			Str("url_path", req.URL.Path).
			Any("body", req.Body).
			Msg("Received a request")
		next.ServeHTTP(resp, req)
	})
}
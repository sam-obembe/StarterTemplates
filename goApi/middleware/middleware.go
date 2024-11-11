package middleware

import "net/http"

type MiddlewareFunc func(h http.HandlerFunc) http.HandlerFunc

func ChainMiddleware(handler http.HandlerFunc, middleware []MiddlewareFunc) http.HandlerFunc {
	var final = handler

	for _, hand := range middleware {
		final = hand(final)
	}

	return final
}

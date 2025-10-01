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

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//do stuff here
		next.ServeHTTP(w, r)
	}
}

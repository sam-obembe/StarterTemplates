package middleware

import (
	"log"
	"net/http"
)

func InterceptorLogger(next http.HandlerFunc) http.HandlerFunc {

	var interceptorFunc = func(w http.ResponseWriter, r *http.Request) {

		//do stuff here
		log.Default().Println(r.URL.Path)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(interceptorFunc)
}

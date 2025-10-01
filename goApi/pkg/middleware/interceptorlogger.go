package middleware

import (
	"log"
	"net/http"
)

func InterceptorLogger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//do stuff here
		log.Default().Println(r.URL.Path)

		next.ServeHTTP(w, r)
	}

}

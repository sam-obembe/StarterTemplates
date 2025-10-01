package main

import (
	"api/handlers"
	"api/middleware"
	"api/pkg/auth"
	"log"
	"net/http"
	"os"
)

func loadEnvValue(key string) string {
	var env, ok = os.LookupEnv(key)
	if !ok {
		log.Fatal("SET " + key)
	}
	return env
}

func main() {

	var envPort = loadEnvValue("API_PORT")
	var mux = http.NewServeMux()
	var fs = http.FileServer(http.Dir("swagger"))

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	auth.MapAuthHandlers(mux)

	mux.HandleFunc("GET /test", middleware.InterceptorLogger(handlers.TestHandler))

	mux.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unuathorized", http.StatusUnauthorized)
	})

	log.Default().Println("Listening on :" + envPort)

	http.ListenAndServe(":"+envPort, mux)
}

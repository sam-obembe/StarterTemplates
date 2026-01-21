package main

import (
	"api/data"
	"api/handlers"
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

func RestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /test", RestMiddleware(handlers.TestHandler))

	mux.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unuathorized", http.StatusUnauthorized)
	})

}

func main() {

	var envPort = loadEnvValue("API_PORT")
	var dbConnection = loadEnvValue("DB_CONNECTION")
	var mux = http.NewServeMux()

	data.InitializePostgres(dbConnection)
	RegisterRoutes(mux)
	log.Default().Println("Listening on :" + envPort)
	http.ListenAndServe(":"+envPort, mux)
}

package main

import (
	"api/handlers"
	"api/middleware"
	"fmt"
	"html/template"
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
	var fs = http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("GET /test", middleware.InterceptorLogger(handlers.TestHandler))

	mux.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unuathorized", http.StatusUnauthorized)
	})

	mux.HandleFunc("GET /swagger/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		var t, _ = template.ParseFiles("views/swagger.html")

		t.Execute(w, nil)
	})

	log.Default().Println("Listening on :" + envPort)

	http.ListenAndServe(":"+envPort, mux)
}

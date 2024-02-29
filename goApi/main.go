package main

import (
	"api/handlers"
	"api/middleware"
	"fmt"
	"net/http"
	"os"
)

func main() {

	var envPort = os.Getenv("API_PORT")
	fmt.Println(envPort)
	//var port = "5000"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", middleware.InterceptorLogger(handlers.TestHandler))
	mux.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unuathorized", http.StatusUnauthorized)
	})

	//fmt.Println("Running on " + port)
	http.ListenAndServe(":"+envPort, mux)
}

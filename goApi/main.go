package main

import (
	"api/handlers"
	"fmt"
	"net/http"
)

func main() {

	port := "5000"

	http.HandleFunc("/test", handlers.TestHandler)
	http.HandleFunc("/notFoundTest", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	http.HandleFunc("/unauthorized", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unuathorized", http.StatusUnauthorized)
	})

	fmt.Println("Running on " + port)
	http.ListenAndServe(":"+port, nil)
}

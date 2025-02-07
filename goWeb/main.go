package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Heading string
	Content string
}

func main() {
	var port = ":7210"
	var mux = http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	log.Default().Println("Listening on " + port)
	http.ListenAndServe(port, mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("layouts/layout.html", "views/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	data := PageData{
		Title:   "Hi",
		Heading: "Hello world",
		Content: "Testing",
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

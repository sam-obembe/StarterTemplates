package handlers

import (
	"api/dto"
	"encoding/json"
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Default().Println("inside test handler")

	res := dto.TestDto{
		Name:    "Sam",
		Age:     28,
		Hobbies: []string{"run", "walk", "swim"},
	}

	res.ChangeName("Dan")

	json.NewEncoder(w).Encode(res)
}

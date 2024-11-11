package handlers

import (
	"api/dto"
	"encoding/json"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	res := dto.TestDto{
		Name:    "Sam",
		Age:     28,
		Hobbies: []string{"run", "walk", "swim"},
	}

	res.ChangeName("Dan")

	json.NewEncoder(w).Encode(res)
}

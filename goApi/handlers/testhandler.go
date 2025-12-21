package handlers

import (
	"api/dto"
	"encoding/json"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {

	res := dto.TestDto{
		Name:    "Sam",
		Age:     28,
		Hobbies: []string{"run", "walk", "swim"},
	}

	json.NewEncoder(w).Encode(res)
}

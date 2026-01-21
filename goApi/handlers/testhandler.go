package handlers

import (
	"api/dto"
	"encoding/json"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {

	obj := dto.TestDto{
		Name:    "Sam",
		Age:     28,
		Hobbies: []string{"run", "walk", "swim"},
	}

	res := dto.Response{
		Message: "success",
		Data:    obj,
	}

	json.NewEncoder(w).Encode(res)
}

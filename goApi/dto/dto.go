package dto

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type TestDto struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

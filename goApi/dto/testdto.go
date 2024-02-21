package dto

type TestDto struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func (t *TestDto) ChangeName(name string) {
	t.Name = name
}

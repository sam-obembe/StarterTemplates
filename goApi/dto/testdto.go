package dto

type DtoValidator interface {
	Validate() (bool, error)
}

type TestDto struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func (t *TestDto) ChangeName(name string) {
	t.Name = name
}

func (t *TestDto) ChangeAge(age int) {
	t.Age = age
}

func (t *TestDto) AddHobby(hobby string) {
	t.Hobbies = append(t.Hobbies, hobby)
}

func (t *TestDto) SetHobbies(hobbies []string) {
	t.Hobbies = hobbies
}

func (t *TestDto) Validate() bool {
	return false
}

package auth

type UserEntity struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

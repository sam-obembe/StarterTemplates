package auth

import "net/http"

func MapAuthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /login", login)
	mux.HandleFunc("POST /logout", logout)
	mux.HandleFunc("POST /register", register)
	mux.HandleFunc("POST /delete", deleteAccount)
	mux.HandleFunc("POST /reset", resetPassword)
}

func login(w http.ResponseWriter, r *http.Request) {
	//get login details from request body

	//validate against hashed password
}

func logout(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {

	//take email and password
	//save in db
	//generate token
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {

}

func resetPassword(w http.ResponseWriter, r *http.Request) {

}

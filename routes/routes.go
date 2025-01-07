package routes

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "helper/login.html")
}
func Signup(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "helper/signup.html")
}

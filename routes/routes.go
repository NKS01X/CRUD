package routes

import (
	database "GOCRUD/Database"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "helper/login.html")
}
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		newUser := database.NewUser(
			r.FormValue("username"),
			r.FormValue("email"),
			r.FormValue("password"),
			r.FormValue("phoneNo"),
		)

		if newUser != nil {
			// Respond with success status
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"success": true})
		} else {
			// Respond with failure status
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]bool{"success": false})
		}
	} else {
		http.ServeFile(w, r, "helper/signup.html")
	}
}

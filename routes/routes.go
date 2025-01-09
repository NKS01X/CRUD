package routes

import (
	database "GOCRUD/Database"
	"encoding/json"
	"net/http"
)

var UserList = database.LinkedList{}
var userList = &UserList

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "./static/login.html")
		return
	} else if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := userList.FindUser(username)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
			return
		}

		if user.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		username := r.FormValue("username")

		// Check if user already exists
		existingUser, _ := userList.FindUser(username)
		if existingUser != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "Username already exists"})
			return
		}

		userList.NewUser(
			username,
			r.FormValue("email"),
			r.FormValue("password"),
			r.FormValue("phoneNo"),
		)

		w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	} else if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "./static/signup.html")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

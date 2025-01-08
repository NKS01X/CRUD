package main

import (
	"GOCRUD/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Login", routes.Login)
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/Signup", routes.Signup)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("helper"))))

	println("Server started at http://localhost:6969/")
	log.Fatal(http.ListenAndServe(":6969", r))

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

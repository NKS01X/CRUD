package main

import (
	"GOCRUD/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/Login", routes.Login)
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/Signup", routes.Signup)

	println("Server started at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", r))

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

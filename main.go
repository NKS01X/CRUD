package main

import (
	"GOCRUD/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Login)
	r.HandleFunc("/Home", serveHome)
	r.HandleFunc("/Signup", routes.Signup)

	println("Server started at http://localhost:6969/")
	log.Fatal(http.ListenAndServe(":6969", r))

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

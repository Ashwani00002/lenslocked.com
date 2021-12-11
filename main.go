package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "To get in touch send mail to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/contact", contact)

	log.Fatal(http.ListenAndServe(":3000", r))
}

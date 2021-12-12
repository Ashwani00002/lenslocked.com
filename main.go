package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "To get in touch send mail to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Sorry!, the page you're looking does not exist")
}

func main() {
	r := httprouter.New()
	r.GET("/", Home)
	r.GET("/contact", contact)
	r.NotFound = http.HandlerFunc(notFound)

	log.Fatal(http.ListenAndServe(":3000", r))
}

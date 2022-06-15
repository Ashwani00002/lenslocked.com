package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

}

func main() {
	homeView = views.NewViews("bootstrap", "views/home.gohtml")
	contactView = views.NewViews("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	log.Fatal(http.ListenAndServe(":3000", r))
}

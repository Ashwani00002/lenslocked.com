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
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))

}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func main() {
	homeView = views.NewViews("bootstrap", "views/home.gohtml")
	contactView = views.NewViews("bootstrap", "views/contact.gohtml")
	signupView = views.NewViews("bootstrap", "views/signup.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

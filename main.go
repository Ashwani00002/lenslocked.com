package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)

func main() {
	StaticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	r.Handle("/", StaticC.Home).Methods("GET")
	r.Handle("/contact", StaticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

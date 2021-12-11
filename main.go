package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, r.URL.Path)
	} else if r.URL.Path == "/Contact" {
		fmt.Fprintf(w, "To get in touch send mail to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>We could not find the page you are looking for</h1><p>Please email us if you keep being sent to an invalid page</p>")
	}
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", mux)
}

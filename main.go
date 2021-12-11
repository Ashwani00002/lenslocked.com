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
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

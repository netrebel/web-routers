package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "Email me at hello@mail.com")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Sorry, not found!")
}

func main() {
	mux := mux.NewRouter()
	mux.NotFoundHandler = http.HandlerFunc(notFound)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", mux)
}

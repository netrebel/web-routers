package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(w, "Email me at hello@mail.com")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found")
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handlerFunc)
	mux.HandleFunc("/contact", handlerFunc)
	http.ListenAndServe(":8080", mux)
}

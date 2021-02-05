package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "text/html")
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprintf(rw, "<h1>Welcome</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(rw, "Email me at hello@mail.com")
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "Not Found")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

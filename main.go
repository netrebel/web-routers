package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("content-type", "text/html")
		fmt.Println(r.URL.Path)
		if r.URL.Path == "/hello" {
			fmt.Fprintf(rw, "<h1>Welcome</h1>")
		}
	})
	http.ListenAndServe(":8080", nil)
}

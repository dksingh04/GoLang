package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", paramValues)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func paramValues(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("p")
	fmt.Fprintln(w, "Query Parameter is: "+v)
}

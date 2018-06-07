package main

import (
	"io"
	"net/http"
)

type handler1 string
type handler2 string

func main() {
	var h1 handler1
	var h2 handler2
	mux := http.NewServeMux()

	mux.Handle("/request1/", h1)
	mux.Handle("/request2", h2)

	/*
		Request routing can be handled in different ways.
		http.HandleFunc("/request1/", req1)
		and function can be
		func req1(rw http.ResponseWriter, req *http.Request) {
			io.WriteString(rw, "Handling Request1")
		}

		Another approach is to handle the request is
		http.Handle("/request1/", http.HandlerFunc(req1))
		and function can be
		func req1(rw http.ResponseWriter, req *http.Request) {
			io.WriteString(rw, "Handling Request1")
		}

		The above two scenario is using DefaultMux and in order to use DefualtMux, pass second parameter as nil in 
		ListenAndServe method.
		e.g. 
		http.ListenAndServe(":8080", nil)
	*/

	http.ListenAndServe(":8080", mux)
}

func (h1 handler1) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Handling Request1")
}

func (h2 handler2) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Handling Request2")
}

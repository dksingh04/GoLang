package main

import (
	"io"
	"net/http"
)

func main() {
	//Serving the files from current directory
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/img", img)
	http.ListenAndServe(":8080", nil)
}
func img(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="deepak.jpg"/>`)
}

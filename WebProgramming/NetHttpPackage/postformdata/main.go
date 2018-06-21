package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", queryParamAndPost)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func queryParamAndPost(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("p")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="POST">
		<input type="text" name="p" value=`+v+`>
		<input type="submit">
	   </form>
	   <br>`+v)
}

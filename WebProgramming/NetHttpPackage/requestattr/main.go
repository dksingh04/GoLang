package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type handler int

var tpl *template.Template

func main() {
	var h handler

	http.ListenAndServe(":8080", h)

}

func (h handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Method string
		Form   url.Values
	}{
		req.Method,
		req.Form,
	}
	rw.Header().Set("Content-type", "text/html; charset=utf-8")
	rw.Header().Set("Cookies", "d;fkg;dflkg;ldfkg;ldflgk")
	tpl.ExecuteTemplate(rw, "index.html", data)
}

func init() {
	tpl = template.Must(tpl.ParseFiles("index.html"))
}

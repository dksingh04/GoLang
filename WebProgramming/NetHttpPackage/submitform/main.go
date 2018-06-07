package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type handler int

var tpl *template.Template

func main() {
	var h handler
	http.ListenAndServe(":8080", h)

}

func (h handler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	tpl.ExecuteTemplate(wr, "index.gohtml", r.Form)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

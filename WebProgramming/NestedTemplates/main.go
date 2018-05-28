package main

import (
	"os"
	"text/template"
)

var tp *template.Template

func init() {
	tp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tp.ExecuteTemplate(os.Stdout, "index.gohtml", 45)

	if err != nil {
		panic(err)
	}
}

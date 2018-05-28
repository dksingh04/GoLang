package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type cityNamebyState struct {
	State string
	City  string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	// passing slice
	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", []int{1, 2, 3, 4})

	if err != nil {
		log.Fatalln(err)
	}

	city := map[string]string{
		"CA": "Santa Clara",
		"OH": "Cincinnati",
		"IL": "Chicago",
	}
	// Passing map
	err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", city)

	if err != nil {
		log.Fatalln(err)
	}

	//passing struct

	caState := cityNamebyState{
		State: "CA", City: "Santa Clara",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", caState)
	if err != nil {
		log.Fatalln(err)
	}

	ohState := cityNamebyState{
		State: "OH", City: "Cincinnati",
	}

	states := []cityNamebyState{caState, ohState}
	// passing slice of structs
	err = tpl.ExecuteTemplate(os.Stdout, "structslice.gohtml", states)

	if err != nil {
		log.Fatalln(err)
	}

	// You can pass nested structs too.
	//https://github.com/GoesToEleven/golang-web-dev/tree/master/007_data-structures

}

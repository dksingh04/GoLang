package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tp *template.Template

func init() {
	tp = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", uploadFile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func uploadFile(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		f, fh, err := req.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("\nfile:", f, "\nheader:", fh, "\nerror:", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

		// Store the file on Server

		target, err := os.Create(filepath.Join("./target/", fh.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer target.Close()

		_, err = target.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tp.ExecuteTemplate(w, "index.html", s)

}

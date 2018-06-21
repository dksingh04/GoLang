package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", readUploadedFile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func readUploadedFile(w http.ResponseWriter, r *http.Request) {
	var s string

	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		f, fh, err := r.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", fh, "\nerr", err)

		bs, err := ioutil.ReadAll(f)

		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
		<input type="file" name="f">
		<input type="submit">
		</form>
		<br>`+s)

}

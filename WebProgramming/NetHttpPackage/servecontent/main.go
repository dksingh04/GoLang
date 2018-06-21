package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveFileContent)
	http.HandleFunc("/deepak", imageFile)
	http.ListenAndServe(":8080", nil)
}

func serveFileContent(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/deepak"/>`)
}

func imageFile(w http.ResponseWriter, req *http.Request) {
	/*f, err := os.Open("deepak.jpg")
	if err != nil {
		panic(err)
	}*/
	//finfo, err := f.Stat()
	//defer f.Close()
	// Serving file using io.Copy method
	//io.Copy(w, f)

	// Serving file using ServeContent method, uncomment to see how it works
	//http.ServeContent(w, req, f.Name(), finfo.ModTime(), f)

	// Serving file using ServeFile, by using this function you don't need to read a file as it is done above.
	http.ServeFile(w, req, "deepak.jpg")
}

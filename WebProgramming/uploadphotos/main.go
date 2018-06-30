package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// route to serve uploaded images
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("pic")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer mf.Close()
		fileExt := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fName := fmt.Sprintf("%x", h.Sum(nil)) + "." + fileExt

		// Create new file, get current working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fName)

		nf, err := os.Create(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer nf.Close()
		// Copy the content of existing file to new file
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// add the filename to cookies
		c = appendFileName(w, c, fName)
	}
	files := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", files[1:])
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	c, err := req.Cookie("Session")

	if err != nil {
		uID, _ := uuid.NewV1()
		c = &http.Cookie{
			Name:  "Session",
			Value: uID.String(),
		}
	}

	http.SetCookie(w, c)
	return c
}

func appendFileName(w http.ResponseWriter, c *http.Cookie, f string) *http.Cookie {
	cValue := c.Value
	if !strings.Contains(cValue, f) {
		cValue += "|" + f
	}
	c.Value = cValue
	http.SetCookie(w, c)
	return c
}

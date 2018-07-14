package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
	Use of HMAC algorithm for Authenticating the message is comming from right user or not.
*/
var secretKey []byte

func init() {
	c := 12
	secretKey = make([]byte, c)
	_, err := rand.Read(secretKey)

	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/authenticate", authenticate)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")

	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if r.Method == http.MethodPost {
		user := r.FormValue("user")
		c.Value = user + "|" + getHMACCode(user)
	}

	http.SetCookie(w, c)

	io.WriteString(w, `<!DOCTYPE html>
		<html>
		  <body>
			<form method="POST">
			  <input type="text" name="user">
			  <input type="submit">
			</form>
			<a href="/authenticate">Validate This `+c.Value+`</a>
		  </body>
		</html>`)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sSplitValue := strings.Split(c.Value, "|")

	user := sSplitValue[0]
	codeReceived := sSplitValue[1]
	//codeChecked := getHMACCode(user) // valid user
	codeChecked := getHMACCode(user + "tamper-data") // Invalid user

	if codeReceived == codeChecked {
		fmt.Printf("Message Authenticated !! \n Received Code: %s \n and \n Checked Code: %s", codeReceived, codeChecked)
	} else {
		fmt.Printf("Message Not Authenticated !! \n Received Code: %s \n and \n Checked Code: %s", codeReceived, codeChecked)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}

func getHMACCode(data string) string {
	h := hmac.New(sha256.New, secretKey)

	io.WriteString(h, data)

	return fmt.Sprintf("%x", h.Sum(nil))
}

package main

// Tracking session using uuid package
import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", createSession)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Session")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "Session",
			Value: id.String(),
			//Secure:true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}

	fmt.Println(cookie)

}

package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type session struct {
	userName     string
	lastActivity time.Time
}

var tpl *template.Template
var userDbSession = make(map[string]session) //sessionId is key and value is UserName
var userDb = make(map[string]user)           // userId and User details
const sessionTimeout int = 30                //Set session timeout for 30 sec.
var dbSessionsCleaned time.Time

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
	Role      string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signupUser)
	http.HandleFunc("/userDetails", userDetails)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		p := req.FormValue("password")
		u, ok := userDb[userName]
		if !ok {
			http.Error(w, "User name or password are incorrect", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "User name or password are incorrect", http.StatusForbidden)
			return
		}

		c := getSessionCookie(w, req)

		c.MaxAge = sessionTimeout
		http.SetCookie(w, c)
		userDbSession[c.Value] = session{userName, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", u)

}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("Session")
	// Clear userDbSession
	delete(userDbSession, c.Value)
	// delete cookie
	c = &http.Cookie{
		Name:   "Session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > time.Second*30 {
		go cleanSession()
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.html", u)

}
func signupUser(w http.ResponseWriter, req *http.Request) {

	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		p := req.FormValue("password")
		firstName := req.FormValue("firstName")
		lastName := req.FormValue("lastName")
		role := req.FormValue("role")

		if _, ok := userDb[userName]; ok {
			http.Error(w, "User Name alreday been taken", http.StatusForbidden)
			return
		}

		// else create a session
		c := getSessionCookie(w, req)

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		u := user{userName, bs, firstName, lastName, role}
		http.SetCookie(w, c)

		userDbSession[c.Value] = session{userName, time.Now()}
		userDb[userName] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func userDetails(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// Example of handling role
	if u.Role != "bdr" {
		http.Error(w, "Un-Authorized Access", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "userdetails.gohtml", u)

}

func getSessionCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	uID, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Unable to generate Unique Session Key", http.StatusInternalServerError)
		return nil
	}
	c := &http.Cookie{
		Name:  "Session",
		Value: uID.String(),
	}

	return c
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

func getUser(req *http.Request) user {
	c, err := req.Cookie("Session")
	if err != nil {
		uID, _ := uuid.NewV4()

		c = &http.Cookie{
			Name:  "Session",
			Value: uID.String(),
		}
	}

	var u user

	if un, ok := userDbSession[c.Value]; ok {
		u = userDb[un.userName]
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("Session")
	if err != nil {
		return false
	}

	un := userDbSession[c.Value]
	_, ok := userDb[un.userName]
	return ok
}

func showSessions() {
	fmt.Println("********")
	for k, v := range userDbSession {
		fmt.Println(k, v.userName)
	}
	fmt.Println("")
}

func cleanSession() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range userDbSession {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(userDbSession, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

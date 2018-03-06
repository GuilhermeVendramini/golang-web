package main

import (
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, _ := req.Cookie("session")
	//http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

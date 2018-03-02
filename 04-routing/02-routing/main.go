package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog dog")
}

type hotcat int

func (m hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	/*
		Not so elegant way
		mux := http.NewServeMux()
		mux.Handle("/dog/", d)
		mux.Handle("/cat/", c)

		http.ListenAndServe(":8080", mux)
	*/
	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
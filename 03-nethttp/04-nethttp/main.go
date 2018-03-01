package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Gui-key", "my key is here")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "<h1>My content here!</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

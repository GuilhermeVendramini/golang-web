package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.go.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.go.html", "My data!")
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.go.html"))
}

func main() {
	data := "My page"
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.go.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}

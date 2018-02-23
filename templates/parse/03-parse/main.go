package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// err := tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.ExecuteTemplate(os.Stdout, "one.go.html", nil)
	if err != nil {
		log.Println("error creating file", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.go.html", nil)
	if err != nil {
		log.Println("error creating file", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.go.html", nil)
	if err != nil {
		log.Println("error creating file", err)
	}

}

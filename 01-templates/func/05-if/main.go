package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"isbu": isBuddha,
}

func init() {
	tpl = template.Must(template.New("tpl.go.html").Funcs(fm).ParseFiles("tpl.go.html"))
}

func isBuddha(s string) bool {
	if s == "Buddha" {
		return true
	}
	return false
}

type sage struct {
	Name  string
	Motto string
	Admin bool
}

type items struct {
	Wisdom []sage
}

func main() {
	a := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: true,
	}

	b := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	c := sage{
		Name:  "Buddha",
		Motto: "No admin here!!",
		Admin: false,
	}

	sages := []sage{a, b, c}

	data := items{
		Wisdom: sages,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("tpl.go.html").Funcs(fm).ParseFiles("tpl.go.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type sage struct {
	Name  string
	Motto string
}

type items struct {
	Wisdom []sage
}

func main() {
	a := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	b := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	sages := []sage{a, b}

	data := items{
		Wisdom: sages,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

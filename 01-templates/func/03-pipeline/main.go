package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.go.html"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func twoTimes(x float64) float64 {
	return x * x
}

var fm = template.FuncMap{
	"fdbl": double,
	"fsq":  square,
	"twot": twoTimes,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.go.html", 2)
	if err != nil {
		log.Fatalln(err)
	}
}

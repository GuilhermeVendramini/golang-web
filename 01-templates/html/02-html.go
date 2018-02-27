package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Guilherme"

	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Page Title</title>
		</head>
		<body>
			<h1>` + name + `</h1>
			<p>My first paragraph.</p>
		</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creatin file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}

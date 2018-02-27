package main

import (
	"fmt"
)

func main() {
	name := "Guilherme"

	tpl := `
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
	`
	fmt.Println(tpl)
}

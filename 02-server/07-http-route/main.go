package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	//request line
	m := strings.Fields(ln)[0] //method
	u := strings.Fields(ln)[1] //uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	//multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html>
		<head><title>Index</title></head>
		<body>
		<b>INDEX</b>
		<a href="/">Index</a>
		<a href="/about">About</a>
		<a href="/contact">Contact</a>
		<a href="/apply">Apply</a></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html>
		<head><title>About</title></head>
		<body>
		<b>ABOUT</b>
		<a href="/">Index</a>
		<a href="/about">About</a>
		<a href="/contact">Contact</a>
		<a href="/apply">Apply</a></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html>
		<head><title>Contact</title></head>
		<body>
		<b>CONTACT</b>
		<a href="/">Index</a>
		<a href="/about">About</a>
		<a href="/contact">Contact</a>
		<a href="/apply">Apply</a></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html>
		<head><title>Contact</title></head>
		<body>
		<b>APPLY</b>
		<a href="/">Index</a>
		<a href="/about">About</a>
		<a href="/contact">Contact</a>
		<a href="/apply">Apply</a>
		<form method="post" action="/apply">
		<input type="submit" value="apply">
		</form>
		</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html>
		<head><title>Contact</title></head>
		<body>
		<b>APPLY PROCESS</b>
		<a href="/">Index</a>
		<a href="/about">About</a>
		<a href="/contact">Contact</a>
		<a href="/apply">Apply</a>
		</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

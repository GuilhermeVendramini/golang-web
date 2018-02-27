package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "go:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you")
}

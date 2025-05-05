package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "ping\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Received: %s", message)
}

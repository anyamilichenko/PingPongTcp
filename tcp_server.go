package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("TCP server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Something went wrong: %s", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Received: %s", message)

	if message == "ping\n" {
		conn.Write([]byte("pong\n"))
	}
}

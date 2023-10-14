package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		logAndExit("Error starting TCP server:", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logAndExit("Error accepting connection:", err)
		}
		go handleClient(conn)
	}
}

func logAndExit(s string, err error) {
	fmt.Println(s, err)
	os.Exit(1)
}

func handleClient(conn net.Conn) {
	// Read data from the client
	defer conn.Close()
	body := make([]byte, 1024)
	n, err := conn.Read(body)
	if err != nil {
		logAndExit("Error reading from client:", err)
	}

	// Process data from the client
	actualData := body[:n]
	inputExpr := string(actualData)

	out := parseExpr(inputExpr)

	// Write data back to the client
	conn.Write([]byte(out))
}

func parseExpr(expr string) string {
	// parse RESP
	return "+OK\r\n"
}

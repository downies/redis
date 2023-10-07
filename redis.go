package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 6379")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//defer conn.Close()
	body := make([]byte, 1024)
	n, err := conn.Read(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Received:", string(body[:n]))

	conn.Write([]byte("+OK\r\n"))

	// Read and process data from the client
	// ...

	// Write data back to the client
	// ...
}

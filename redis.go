package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	commands, err := parseData(expr)
	if err != nil {
		logAndExit("Error parsing data:", err)
	}
	fmt.Println(commands)
	return "+OK\r\n"
}

func parseData(data string) ([]string, error) {
	var result []string
	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// Check if the line starts with "$" (indicating a string length)
		if strings.HasPrefix(line, "$") {
			// Ensure that we have a following line with data
			if i+1 < len(lines) {
				dataLine := strings.TrimSpace(lines[i+1])
				result = append(result, dataLine)
				i++ // Skip the data line
			}
		}
	}
	return result, nil
}

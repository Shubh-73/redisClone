package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Printf("failed to read from connection: %v\n", err)
		}
		return
	}

	// Process the received data (for this example, we assume the client sends the PING command)
	received := string(buf[:n])
	fmt.Printf("Received: %s\n", received)

	// Check if the received message is "PING"
	if received == "PING" {
		// Write "PONG" response back to the client
		_, err := conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			log.Printf("failed to write to connection: %v\n", err)
		}
	} else {
		// Write a generic response back to the client
		_, err := conn.Write([]byte("Unknown command"))
		if err != nil {
			log.Printf("failed to write to connection: %v\n", err)
		}
	}

}

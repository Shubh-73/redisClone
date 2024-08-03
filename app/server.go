package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"log"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//Uncomment this block to pass the first stage
	listener, err := server.BindToPort(":6379")
	conn, err := listener.Accept()
	if err != nil {
		// Handle connection acceptance error
		log.Printf("failed to accept connection: %v\n", err)

	}
	server.HandleConnection(conn)
}

package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	listener, err := server.BindToPort(":6379")
	if err != nil {
		// Print the error and exit if binding fails
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close() // Ensure the listener is closed when done

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("error accepting connection")
			continue
		}

		go server.HandleConnection(conn)
	}
}

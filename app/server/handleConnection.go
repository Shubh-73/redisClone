package server

import (
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		log.Println("Error writing response:", err)
	}

}

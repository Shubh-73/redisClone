package server

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error while reading from connection: %v\n", err)
			return
		}

		commands := strings.Split(strings.TrimSpace(data), "\n")
		for _, command := range commands {
			if command == "PING" {
				_, err := conn.Write([]byte("PONG\r\n"))
				if err != nil {
					log.Println("error while writing to connection: %v\n", err)
					return
				}
			} else {
				_, err := conn.Write([]byte("command unknown: " + command + "\r\n"))
				if err != nil {
					log.Println("error while writing to connection: %v\n", err)
					return
				}
			}
		}
	}

}

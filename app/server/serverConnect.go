package server

import (
	"fmt"
	"net"
)

func BindToPort(port string) (net.Listener, error) {
	// Attempt to listen on the specified address and port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		// Return nil and the error if binding fails
		return nil, fmt.Errorf("failed to bind to port %s: %w", port, err)
	}
	// Return the listener and nil error if successful
	return listener, nil
}

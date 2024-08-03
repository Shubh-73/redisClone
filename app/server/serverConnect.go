package server

import (
	"fmt"
	"net"
)

func BindToPort() (net.Listener, error) {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return nil, fmt.Errorf("failed to bind to port: %w", err)
	}

	return listener, nil
}

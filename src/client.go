package main

import (
	"fmt"
	"net"
)

type HTTPRequest struct {
	method  string
	path    string
	headers map[string]string
	body    string
}

type HTTPResponse struct {
	statusCode  int
	statusText  string
	body        string
	contentType string
}

func handleClientConnection(conn net.Conn, server *Server) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	request := string(buffer[:bytesRead])
	server.HandleRequest(conn, request)
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	server := NewServer()
	server.HandleRoute("GET", "/", showHomePage)
	server.HandleRoute("GET", "/echo", echoRequest)
	server.HandleRoute("GET", "/user-agent", getUserAgent)
	server.HandleRoute("GET", "/files", fetchFile)
	server.HandleRoute("POST", "/files", saveFile)

	listener, err := net.Listen("tcp", "0.0.0.0:6969")
	if err != nil {
		fmt.Println("Failed to bind to port 6969")
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleClientConnection(conn, server)
	}
}

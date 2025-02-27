package main

import (
	"fmt"
	"net"
	"strings"
)

type Server struct {
	routes map[string]func(*HTTPRequest) HTTPResponse
}

func NewServer() *Server {
	return &Server{routes: make(map[string]func(*HTTPRequest) HTTPResponse)}
}

func (ws *Server) HandleRoute(method string, path string, handler func(*HTTPRequest) HTTPResponse) {
	ws.routes[method+path] = handler
}

func (ws *Server) HandleRequest(conn net.Conn, request string) {
	requestParts := strings.Split(request, "\r\n")
	if len(requestParts) < 1 {
		sendResponse(conn, 400, "Bad Request", "Invalid request", HTTPHeaders{})
		return
	}
	requestLine := strings.Fields(requestParts[0])
	if len(requestLine) < 2 {
		sendResponse(conn, 400, "Bad Request", "Invalid request", HTTPHeaders{})
		return
	}

	method := requestLine[0]
	path := requestLine[1]
	headers := parseHeaders(requestParts[1:])
	body := requestParts[len(requestParts)-1]

	pathSegments := strings.Split(path, "/")
	if len(pathSegments) < 2 {
		return
	}

	handler, exists := ws.routes[method+"/"+pathSegments[1]]
	if !exists {
		sendResponse(conn, 404, "Not Found", "", HTTPHeaders{})
		logRequestDetails(method, path, headers, 404, "Not Found")
		return
	}

	req := HTTPRequest{method, path, headers, body}
	res := handler(&req)
	headersToSend := HTTPHeaders{
		"Content-Type":   res.contentType,
		"Content-Length": fmt.Sprintf("%d", len(res.body)),
	}
	compressResponseBody(&req, &res, &headersToSend)
	sendResponse(conn, res.statusCode, res.statusText, res.body, headersToSend)
	logRequestDetails(method, path, headers, res.statusCode, res.statusText)
}

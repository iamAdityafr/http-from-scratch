package main

import (
	"fmt"
	"net"
	"strings"
)

func parseHeaders(headerLines []string) map[string]string {
	headers := make(map[string]string)
	for _, line := range headerLines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			headers[strings.ToLower(parts[0])] = strings.ToLower(parts[1])
		}
	}
	return headers
}

func sendResponse(conn net.Conn, statusCode int, statusText, body string, headers HTTPHeaders) {
	var bd strings.Builder
	bd.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, statusText))
	bd.WriteString(headers.String())
	bd.WriteString("\r\n")
	bd.WriteString(body)
	bd.WriteString("\r\n")
	conn.Write([]byte(bd.String()))
}

type HTTPHeaders map[string]string

func (h *HTTPHeaders) String() string {
	var bd strings.Builder
	for key, value := range *h {
		bd.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	return bd.String()
}

func logRequestDetails(method, path string, headers map[string]string, statusCode int, statusText string) {
	logMessage := fmt.Sprintf("Request: %s %s\nHeaders: %v\nResponse: %d %s\n", method, path, headers, statusCode, statusText)
	fmt.Println(logMessage)
}

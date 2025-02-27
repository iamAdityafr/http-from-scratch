package main

import (
	"os"
	"strings"
)

func fetchFile(r *HTTPRequest) HTTPResponse {
	segments := strings.Split(r.path, "/")
	if len(segments) == 3 {
		var directoryPath string
		if len(os.Args) < 3 {
			directoryPath = ""
		} else {
			directoryPath = os.Args[2]
		}
		fileName := segments[2]
		content, err := os.ReadFile(directoryPath + fileName)
		if err != nil {
			return HTTPResponse{404, "Not Found", "text/plain", err.Error()}
		}
		return HTTPResponse{200, "OK", "application/octet-stream", string(content)}
	}
	return HTTPResponse{404, "Not Found", "text/plain", "Not Found"}
}

func saveFile(r *HTTPRequest) HTTPResponse {
	segments := strings.Split(r.path, "/")
	if len(segments) == 3 {
		var directoryPath string
		if len(os.Args) < 3 {
			directoryPath = ""
		} else {
			directoryPath = os.Args[2]
		}
		fileName := segments[2]
		os.WriteFile(directoryPath+fileName, []byte(r.body), 0644)
		return HTTPResponse{201, "Created", "text/plain", "saved"}
	}
	return HTTPResponse{404, "Not Found", "text/plain", "Not Found"}
}

func echoRequest(r *HTTPRequest) HTTPResponse {
	segments := strings.Split(r.path, "/")
	if len(segments) == 3 {
		return HTTPResponse{200, "OK", "text/plain", segments[2]}
	}
	return HTTPResponse{404, "Not Found", "text/plain", "Not Found"}
}

func getUserAgent(r *HTTPRequest) HTTPResponse {
	userAgent, exists := r.headers["user-agent"]
	if !exists {
		return HTTPResponse{400, "Bad Request", "text/plain", "User-Agent header not found"}
	}
	return HTTPResponse{200, "OK", "text/plain", userAgent}
}

func showHomePage(_ *HTTPRequest) HTTPResponse {
	return HTTPResponse{200, "OK", "text/html", "<h1>Hello World</h1>"}
}

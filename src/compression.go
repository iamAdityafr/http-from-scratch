package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"slices"
	"strings"
)

func compressResponseBody(r *HTTPRequest, res *HTTPResponse, headers *HTTPHeaders) (bool, string) {
	encoding, exists := r.headers["accept-encoding"]
	if !exists {
		return false, ""
	}
	acceptedEncodings := strings.Fields(strings.Replace(encoding, ",", "", -1))
	if !slices.Contains(acceptedEncodings, "gzip") {
		return false, ""
	}

	var buffer bytes.Buffer
	compressor := gzip.NewWriter(&buffer)
	compressor.Write([]byte(res.body))
	compressor.Close()

	res.body = buffer.String()
	(*headers)["Content-Encoding"] = "gzip"
	(*headers)["Content-Length"] = fmt.Sprintf("%d", len(buffer.Bytes()))
	return true, ""
}

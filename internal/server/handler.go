package server

import (
	"net"
	"strconv"

	"github.com/vijayvenkatj/http-protocol/internal/headers"
	"github.com/vijayvenkatj/http-protocol/internal/request"
	"github.com/vijayvenkatj/http-protocol/internal/response"
)


type HandlerError struct {
	StatusCode response.StatusCode
	Message    []byte
}


type Handler func(w *response.Writer, req *request.Request) *HandlerError


func HandleError(conn net.Conn, handlerError HandlerError) {
	defer conn.Close()

    // Wrap conn with our response.Writer
	w := response.NewWriter(conn)

    // Write status line
	if err := w.WriteStatusLine(handlerError.StatusCode); err != nil {
		return
	}

    // Build headers
	hdrs := headers.NewHeaders()
	hdrs.Set("Content-Length", strconv.Itoa(len(handlerError.Message)))
	hdrs.Set("Connection", "close")
	hdrs.Set("Content-Type", "text/html")

    // Send headers
	if err := w.WriteHeaders(hdrs); err != nil {
		return
	}

    // Send body
	_ = w.WriteBody(handlerError.Message)
	return
}

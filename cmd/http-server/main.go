package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/vijayvenkatj/http-protocol/internal/headers"
	"github.com/vijayvenkatj/http-protocol/internal/request"
	"github.com/vijayvenkatj/http-protocol/internal/response"
	"github.com/vijayvenkatj/http-protocol/internal/server"
)



const port = 42069

func main() {

	handler := func(w *response.Writer, req *request.Request) *server.HandlerError {
		if req.RequestLine.RequestTarget == "/yourproblem" {
			return &server.HandlerError{StatusCode: 400, Message: []byte(`<!DOCTYPE html>
			<html>
			  <head>
				<title>400 Bad Request</title>
			  </head>
			  <body>
				<h1>Bad Request</h1>
				<p>Your request honestly kinda sucked.</p>
			  </body>
			</html>`)}
		} else if req.RequestLine.RequestTarget == "/myproblem" {
			return &server.HandlerError{StatusCode: 500, Message: []byte(`<!DOCTYPE html>
			<html>
			  <head>
				<title>500 Internal Server Error</title>
			  </head>
			  <body>
				<h1>Internal Server Error</h1>
				<p>Okay, you know what? This one is on me.</p>
			  </body>
			</html>`)}
		} else {
			data := []byte(`<!DOCTYPE html>
			<html>
			<head>
				<title>200 OK</title>
			</head>
			<body>
				<h1>Success!</h1>
				<p>Your request was an absolute banger.</p>
			</body>
			</html>`)

			if err := w.WriteStatusLine(response.OK); err != nil {
				return &server.HandlerError{StatusCode: 500, Message: []byte(err.Error())}
			}

			hdrs := headers.NewHeaders()
			hdrs.Set("Content-Length", strconv.Itoa(len(data)))
			hdrs.Set("Connection", "close")
			hdrs.Set("Content-Type", "text/html")

			if err := w.WriteHeaders(hdrs); err != nil {
				return &server.HandlerError{StatusCode: 500, Message: []byte(err.Error())}
			}

			if err := w.WriteBody(data); err != nil {
				return &server.HandlerError{StatusCode: 500, Message: []byte(err.Error())}
			}

			return nil
		}
	}

	server, err := server.Serve(port,handler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()

	log.Println("Server started on port", port)
	server.Listen()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}

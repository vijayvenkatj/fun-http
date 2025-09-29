package server

import (
	"fmt"
	"net"
	"github.com/vijayvenkatj/http-protocol/internal/request"
	"github.com/vijayvenkatj/http-protocol/internal/response"
)

type Server struct {
	Listener 	net.Listener
	State 		bool
    HandlerFunc Handler
}

func Serve(port int, handlerFunc Handler) (*Server, error) {
	
	listener, err := net.Listen("tcp",fmt.Sprintf(":%d",port));
	if err != nil {
		return nil, err
	}

	server := &Server{
		Listener: listener,
		State: true,
        HandlerFunc: handlerFunc,
	}

	return server,nil
}


func (server *Server) Close() error {
    err := server.Listener.Close()
    server.State = false
    return err
}

func (server *Server) Listen() {
    if !server.State {
        return
    }

    for {
        conn, err := server.Listener.Accept()
        if err != nil {
            if server.State {
                fmt.Println("Accept error:", err)
            }
            return
        }
		fmt.Print("conn accepted")
        go server.handle(conn)
    }
}

func (server *Server) handle(conn net.Conn) {
    defer conn.Close()

    req, err := request.RequestFromReader(conn)
    if err != nil {
        HandleError(conn, HandlerError{StatusCode: response.BAD_REQUEST, Message: []byte("Bad Request")})
        return
    }

    // Create response writer bound to this connection
    respWriter := response.NewWriter(conn)

    // Call handler
    handlerError := server.HandlerFunc(respWriter, req)
    if handlerError != nil {
        HandleError(conn, *handlerError)
        return
    }
}


package server

import (
	"fmt"
	"net"

	"github.com/vijayvenkatj/http-protocol/internal/headers"
	"github.com/vijayvenkatj/http-protocol/internal/response"
)

type Server struct {
	Listener 	net.Listener
	State 		bool
}

func Serve(port int) (*Server, error) {
	
	listener, err := net.Listen("tcp",fmt.Sprintf(":%d",port));
	if err != nil {
		return nil, err
	}

	server := &Server{
		Listener: listener,
		State: true,
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

    buf := make([]byte, 1024)
    conn.Read(buf)

    response.WriteStatusLine(conn,200);

    headers := headers.NewHeaders();
    headers.Set("Content-Length","0");
    headers.Set("Connection","close");
    headers.Set("Content-Type","text/plain");

    err := response.WriteHeaders(conn,headers)
    if err != nil {
        return
    }
}

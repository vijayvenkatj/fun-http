package server

import (
	"fmt"
	"net"
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

    body := "Hello World!"
    response := fmt.Sprintf(
        "HTTP/1.1 200 OK\r\n"+
            "Content-Type: text/plain\r\n"+
            "Content-Length: %d\r\n"+
            "\r\n%s",
        len(body),
        body,
    )

    conn.Write([]byte(response))
}

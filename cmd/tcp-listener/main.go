package main

import (
	"fmt"
	"net"

	"github.com/vijayvenkatj/http-protocol/internal/request"
)


func main() {

	tcpListener, err := net.Listen("tcp",":42069")
	if err != nil {
		return
	}
	defer tcpListener.Close()

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			return
		}

		fmt.Println("Connection has been accepted!");

		requestData, err  := request.RequestFromReader(conn)
		if err != nil {
			return
		}

		fmt.Print("Request line:\n")
		fmt.Printf("- Method: %s\n",requestData.RequestLine.Method);
		fmt.Printf("- Target: %s\n",requestData.RequestLine.RequestTarget);
		fmt.Printf("- Version: %s\n",requestData.RequestLine.HttpVersion);

		fmt.Printf("Headers: \n")
		for key,val := range requestData.Headers {
			fmt.Printf("- %s: %s\n",key,val)
		}

		fmt.Printf("Body: \n")
		fmt.Printf("%s",requestData.Body)

		fmt.Println("Channel has been closed")
	}
	
}
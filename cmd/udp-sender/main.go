package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	udpArr, err := net.ResolveUDPAddr("udp",":42069")
	if err != nil {
		return
	}

	udpConn, err := net.DialUDP("udp",nil,udpArr)
	if err != nil {
		return
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")

		line, err := reader.ReadString('\n');
		if err != nil {
			return
		}

		_, err = udpConn.Write([]byte(line))
		if err != nil {
			return
		}

	}
	
}
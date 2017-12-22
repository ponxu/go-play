package main

import (
	"net"
	"fmt"
)

const (
	Ver = 0x5
	AuthNone = 0x0
)

func main() {
	server, err := net.Listen("tcp", ":1080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			break
		}
		go handleConnection(conn)
	}

	fmt.Println("End")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connect:", conn.RemoteAddr())
	handshake(conn)
	cmd(conn)
}

func handshake(conn net.Conn) {
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println("recive length:", len, "  bytes:", buffer[:len])

	conn.Write([]byte{Ver, AuthNone})
	fmt.Println("handshake finish")
}

func cmd(conn net.Conn) {
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("recive length:", len, "  bytes:", buffer[:len])

	fmt.Println("cmd_connect finish")
}

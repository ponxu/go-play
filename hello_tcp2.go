package main

import (
	"net"
	"fmt"
)

func main() {
	laddr := ":8080"
	server, err := net.Listen("tcp", laddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	fmt.Printf("%s connected\n", )

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConection(conn)
	}
}
func handleConection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("%s connected\n", conn.RemoteAddr())


}

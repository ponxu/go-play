package main

import (
	"net"
	"fmt"
	"os"
	"io"
	"time"
)

const (
	FileName = "/home/xwz/xwz/go-play/hello_sendfile.go"
	NetType  = "tcp"
	Addr     = "127.0.0.1:18080"
)

func main() {
	go server()
	time.Sleep(time.Second * 1)
	client()
}

func server() {
	s, err := net.Listen(NetType, Addr)
	if err != nil {
		panic(err)
	}
	defer func() {
		s.Close()
		fmt.Println("server closed!")
	}()

	fmt.Println("server listen on:", s)
	for {
		conn, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go send(conn)
	}
}

func send(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("conn closed by server")
	}()

	f, err := os.Open(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(conn, f)
}

func client() {
	conn, err := net.DialTimeout(NetType, Addr, time.Second*2)
	if err != nil {
		fmt.Println("fail to dial:", err)
		return
	}
	recive(conn)
}

func recive(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("conn closed by client")
	}()

	buffer := make([]byte, 1024)
	for {
		c, err := conn.Read(buffer)
		if c <= 0 || err != nil {
			fmt.Println(c, err, err == io.EOF)
			break
		}

		resp := string(buffer[:c])
		fmt.Println("recive:", resp)
	}
}

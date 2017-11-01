package main

import (
	"fmt"
	"golang.org/x/exp/mmap"
)

func main() {
	in, err := mmap.Open("/home/xwz/xwz/go-play/hello_mmap.go")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer in.Close()
	fmt.Println(in.Len())

	buffer := make([]byte, 1000)
	c, err := in.ReadAt(buffer, 0)
	fmt.Println(c)
	for i, b := range buffer {
		fmt.Println(i, b)
	}

	s := string(buffer[:c])
	fmt.Println(s)
}

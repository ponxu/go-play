package main

import (
    "fmt"
    "io"
)

type F struct {
    io.Reader
}

func (f *F) Read(p []byte) (n int, err error) {
    return 0, nil
}

func hello(r io.Reader) {
    fmt.Println(r)
}

func main() {
    f := F{}
    pf := &f
    fmt.Println(f, pf)
    hello(pf)
}

package main

import (
    "fmt"
)

type Handler interface {
    Do()
}

type Object struct{}

func (o *Object) Do() {
}

func test(o *Object) {}

func main() {
    o := &Object{}
    fmt.Println(o)

    var h Handler
    h = o
    fmt.Println(h)
    fmt.Println(h == o)

    test(h.(*Object))
}

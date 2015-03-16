package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "123"
    i, _ := strconv.Atoi(s)
    fmt.Println(s, i)

    s = "/456"
    i, _ = strconv.Atoi(s[1:])
    fmt.Println(s, i)
}

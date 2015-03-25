package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Loc:", time.Now().Location())
    fmt.Println("Now:", time.Now().Unix())
    fmt.Println("Now:", time.Now().UTC().Unix())
    fmt.Println("Now:", time.Now().Format("2006-01-02 15:04:05"))
    fmt.Println("Now:", time.Now().UTC().Format("2006-01-02 15:04:05"))

    now := time.Now()
    fmt.Println("Now Seco:", now.Unix())
    fmt.Println("Now Mill:", now.UnixNano()/1000000)
    fmt.Println("Now Nano:", now.UnixNano())
}

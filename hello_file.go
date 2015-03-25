package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    file, err := os.Create("dc.txt")
    if err != nil {
        return
    }
    defer file.Close()

    s := time.Now().UnixNano() / 1000000
    for i := 0; i < 100000000; i++ {
        fmt.Fprintf(file, "%d %f\n", time.Now().UnixNano()/1000000, 1234567.89)
        if i%10000 == 0 {
            fmt.Println(i)
        }
    }
    fmt.Println("Use:", time.Now().UnixNano()/1000000-s, "ms")
}

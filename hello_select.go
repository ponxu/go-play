package main

import "fmt"
import "time"

var c = make(chan bool)

func doselect() {
    for {
        select {
        case _, err := <-c:
            fmt.Println("<-", err)
        default:
            fmt.Println("default...")
        }
        fmt.Println("loop...")
    }
}

func dosend() {
    for i := 0; i < 5; i++ {
        time.Sleep(2 * time.Second)
        c <- true
    }
    //close(c)
}

func main() {
    go doselect()
    go dosend()

    <-make(chan int)
}

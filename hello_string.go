package main

import "fmt"

func main() {
    s1 := "123"
    s2 := "123"
    fmt.Println(s1 == s2)

    s3 := "abc"
    b3 := []byte(s3)
    fmt.Println(s3, b3)
    b3[0] = 100
    fmt.Println(s3, b3, string(b3))

    b4 := []byte{97, 98, 99}
    s4 := string(b4)
    b4[0] = 100
    fmt.Println(s4, b4, string(b4)) // s4 b4 not share
}

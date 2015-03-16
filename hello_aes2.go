package main

import (
    "fmt"
    "gooutils/cryptoutils"
    "unicode/utf8"
)

func main() {
    key := "1qaz@WSX3edc$RFV"
    text := "Hello你好"

    s1 := cryptoutils.AESEncrypt(key, text)
    fmt.Println(s1)

    s2 := cryptoutils.AESDecrypt(key, s1)
    fmt.Println(s2, utf8.RuneCountInString(s2))
}

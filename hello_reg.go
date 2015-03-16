package main

import (
    "fmt"
    "regexp"
)

func main() {
    s := `\123\abc`
    r := regexp.MustCompile(`\w+`)
    rs := r.FindAllString(s, -1)
    fmt.Println(rs)
}

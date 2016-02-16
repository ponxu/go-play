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

    s = "/abc/123/ggg/123aa"
    r = regexp.MustCompile("/([a-z]+)/(w+)/([a-z]+)/(.+)")
    rs2 := r.FindAllStringSubmatch(s, -1)
    fmt.Println(rs2)
}

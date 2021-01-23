package main

import (
    "fmt"
    "strings"
)

func main() {
    var C string
    fmt.Scanf("%s", &C)
	s := strings.Split(C, "")

    var flag bool = true
    var c string = s[0]
    for i:=1; i<3; i++{
        if c != s[i]{
            flag = false
        }
    }

    if flag{
        fmt.Printf("Won\n")
    } else {
        fmt.Printf("Lost\n")
    }
}
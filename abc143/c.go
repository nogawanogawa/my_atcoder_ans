package main

import (
    "fmt"
    "strings"
)


func main() {
    var N int
    fmt.Scanf("%d", &N)

    var S string
    fmt.Scanf("%s", &S)

    s := strings.Split(S, "")

    var color string = ""
    var count int = 0
    for i:=0; i<N; i++{
        if s[i] != color{
            count += 1
            color = s[i]
        }
    }

    fmt.Printf("%d\n", count)
}

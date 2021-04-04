package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    var S string
    fmt.Scanf("%s", &S)

    l := N/2

    if S[:l] == S[l:]{
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }
}
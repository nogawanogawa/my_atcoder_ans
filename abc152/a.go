package main

import (
    "fmt"
)

func main() {
    var N, M int

    fmt.Scanf("%d %d", &N, &M)

    if N == M{
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }

}
package main

import (
    "fmt"
)

func main() {
    var M, H int
    fmt.Scanf("%d %d", &M, &H)

    if H % M == 0{
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }

}
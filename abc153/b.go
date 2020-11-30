package main

import (
    "fmt"
)

func main() {
    var H, N int64
    fmt.Scanf("%d %d", &H, &N)

    var A int64
    var sum int64 = 0
    var i int64
    for i=0; i<N; i++{
        fmt.Scanf("%d", &A)
        sum += A
    }

    if sum >= H {
        fmt.Printf("Yes\n")
    } else {
        fmt.Printf("No\n")
    }


}
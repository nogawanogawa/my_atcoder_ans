package main

import (
    "fmt"
)

func main() {
    var N, K int
    fmt.Scanf("%d %d", &N, &K)

    h := make([]int, N)
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &h[i])
    }

    var count int = 0
    for i:=0; i<N; i++{
        if h[i] >= K{
            count ++
        }
    }

    fmt.Printf("%d\n", count)
}
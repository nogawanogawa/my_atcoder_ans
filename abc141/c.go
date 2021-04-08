package main

import (
    "fmt"
)

func main() {
    var N, K, Q int
    fmt.Scanf("%d %d %d", &N, &K, &Q)

    var A int
    p := make([]int, N)

    for i:=0; i<Q; i++{
        fmt.Scanf("%d", &A)
        p[A-1] += 1
    }

    for i:=0; i<N; i++{
        if p[i]+K-Q > 0{
            fmt.Printf("Yes\n")
        } else {
            fmt.Printf("No\n")
        }
    }

}

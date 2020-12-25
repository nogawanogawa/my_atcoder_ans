package main

import (
    "fmt"
)

func main() {
    var A, B, K int64
    fmt.Scanf("%d %d %d", &A, &B, &K)

    if A >= K{
        A = A - K
        K = 0
    } else {
        K = K - A
        A = 0
    }

    if B >= K {
        B = B - K
        K = 0
    } else {
        K = K - B
        B = 0
    }

    fmt.Printf("%d %d\n", A, B)

}
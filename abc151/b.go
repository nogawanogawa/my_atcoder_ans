package main

import (
    "fmt"
)

func main() {
    var N, K, M  int
    fmt.Scanf("%d %d %d", &N, &K, &M)

    want := N * M
    var sum, A int
    for i:=0; i<N-1; i++{
        fmt.Scanf("%d", &A)
        sum += A
    }

    ans := want - sum
    if ans > K {
        fmt.Printf("-1\n")
    } else {
        if ans < 0{
            ans = 0
        }
        fmt.Printf("%d\n", ans)
    }
}
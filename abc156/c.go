package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)

    X:= make([]int, N)
    var A int = 0
    for i:=0; i<N; i++{
        fmt.Scanf("%d", &X[i])
        A += X[i]
    }

    P := float64(A) / float64(N)

    P_1 := float64(int(P))
    P_2 := float64(P_1 + 1)

    if (P - P_1) < (P_2 - P) {
        P = P_1
    } else {
        P = P_2
    }

    var ans float64
    for i:=0; i<N; i++{
        ans += (float64(X[i]) - P) * (float64(X[i]) - P)
    }

    fmt.Printf("%d\n", int((ans)))
}